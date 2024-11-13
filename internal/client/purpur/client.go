package purpur

import (
	"encoding/json"
	"fmt"
	"github.com/eldius/bubble-pocs/internal/config"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type Client struct {
	c *http.Client
}

func NewClient() *Client {
	return &Client{
		c: &http.Client{},
	}
}

func (c *Client) GetMinecraftVesions() (*GetMinecraftVersionsResponse, error) {

	res, err := c.c.Get("https://api.purpurmc.org/v2/purpur/")
	if err != nil {
		return nil, fmt.Errorf("getting minecraft versions: %v", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("getting minecraft versions: %v", res.Status)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading minecraft versions: %v", err)
	}
	return parseGetMineVersionsResponse(b, res.StatusCode)
}

func (c *Client) GetPurpurBuildsByMineVersion(ver string) (*GetPurpurVersionsResponse, error) {
	url := "https://api.purpurmc.org/v2/purpur/" + ver
	log := slog.With(slog.String("url", url))
	res, err := c.c.Get(url)
	if err != nil {
		return nil, fmt.Errorf("getting purpur builds for ver %s: %v", ver, err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	b, err := io.ReadAll(res.Body)
	if config.GetDebug() {
		log.With(
			slog.String("body", string(b)),
			slog.Int("status_code", res.StatusCode),
		).Debug("GetPurpurBuildsByMineVersion")
	}
	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("getting purpur builds: %v", res.Status)
	}
	if err != nil {
		return nil, fmt.Errorf("reading purpur builds: %v", err)
	}
	return parseGetPurpurVersionsResponse(b, res.StatusCode)
}

func (c *Client) Download(mineVer, purpurBuild string) (string, error) {
	url := fmt.Sprintf("https://api.purpurmc.org/v2/purpur/%s/%s/download", mineVer, purpurBuild)
	log := slog.With(slog.String("url", url))
	res, err := c.c.Get(url)
	if err != nil {
		return "", fmt.Errorf("getting purpur package for %s-%s: %v", mineVer, purpurBuild, err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	destDir, err := os.MkdirTemp(os.TempDir(), "downloads-*")
	if err != nil {
		err = fmt.Errorf("creating temp dir: %v", err)
		return "", err
	}
	out, err := os.CreateTemp(destDir, fmt.Sprintf("purpur-%s-%s.jar", mineVer, purpurBuild))
	if err != nil {
		err = fmt.Errorf("creating temp file: %v", err)
		return "", err
	}

	log.With(
		slog.String("dest_name", out.Name()),
		slog.String("dest_folder", destDir),
		slog.Int("status_code", res.StatusCode),
		"headers", res.Header,
	).Debug("DownloadingPackage")

	if res.StatusCode/100 != 2 {
		return "", fmt.Errorf("getting purpur builds: %v", res.Status)
	}

	if _, err := io.Copy(out, res.Body); err != nil {
		err = fmt.Errorf("copying package to temp file: %v", err)
		return "", err
	}

	return out.Name(), nil
}

func parseGetMineVersionsResponse(b []byte, statusCode int) (*GetMinecraftVersionsResponse, error) {
	if config.GetDebug() {
		slog.With(
			slog.String("body", string(b)),
			slog.Int("status_code", statusCode),
		).Debug("GetMinecraftVersions")
	}
	var versionsRes GetMinecraftVersionsResponse
	if err := json.Unmarshal(b, &versionsRes); err != nil {
		return nil, fmt.Errorf("decoding minecraft versions: %v", err)
	}
	return &versionsRes, nil
}

func parseGetPurpurVersionsResponse(b []byte, statusCode int) (*GetPurpurVersionsResponse, error) {
	if config.GetDebug() {
		slog.With(
			slog.String("body", string(b)),
			slog.Int("status_code", statusCode),
		).Debug("GetMinecraftVersions")
	}
	var versionsRes GetPurpurVersionsResponse
	if err := json.Unmarshal(b, &versionsRes); err != nil {
		return nil, fmt.Errorf("decoding minecraft versions: %v", err)
	}
	return &versionsRes, nil
}

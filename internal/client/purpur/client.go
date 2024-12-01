package purpur

import (
	"encoding/json"
	"fmt"
	"github.com/eldius/bubble-pocs/internal/config"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

const (
	purpurApiBaseEndpoint = "https://api.purpurmc.org/v2"
	purpurProject         = "purpur"
)

type Client struct {
	c *http.Client
}

func NewClient() *Client {
	return &Client{
		c: &http.Client{},
	}
}

func (c *Client) GetProjects() (*GetMinecraftVersionsResponse, error) {
	url := purpurApiBaseEndpoint
	var res GetMinecraftVersionsResponse
	handler, err := handleAPIResponse(c.c.Get(url))
	if err != nil {
		return nil, fmt.Errorf("error getting projects: %w", err)
	}
	if err := handler(&res); err != nil {
		return nil, fmt.Errorf("error getting projects: %w", err)
	}

	return &res, nil
}

func (c *Client) GetPurpurMinecraftVesions() (*GetMinecraftVersionsResponse, error) {
	var res GetMinecraftVersionsResponse
	handler, err := handleAPIResponse(c.c.Get(getProjectEndpoint(purpurProject)))
	if err != nil {
		return nil, fmt.Errorf("getting minecraft versions: %v", err)
	}
	if err := handler(&res); err != nil {
		return nil, fmt.Errorf("getting minecraft versions: %v", err)
	}
	return &res, nil
}

func (c *Client) GetPurpurBuildsByMineVersion(ver string) (*GetPurpurVersionsResponse, error) {
	url := fmt.Sprintf("%s/%s", getProjectEndpoint(purpurProject), ver)
	var res GetPurpurVersionsResponse
	handler, err := handleAPIResponse(c.c.Get(url))
	if err != nil {
		return nil, fmt.Errorf("getting minecraft versions: %w", err)
	}
	if err := handler(&res); err != nil {
		return nil, fmt.Errorf("getting minecraft versions: %w", err)
	}
	return &res, nil
}

func (c *Client) DownloadPurpur(mineVer, purpurBuild, destDir string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s/download", getProjectEndpoint(purpurProject), mineVer, purpurBuild)
	log := slog.With(slog.String("url", url))
	res, err := c.c.Get(url)
	if err != nil {
		return "", fmt.Errorf("getting purpur package for %s-%s: %v", mineVer, purpurBuild, err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	fileName := fmt.Sprintf("purpur-%s-%s.jar", mineVer, purpurBuild)

	out, err := os.Create(filepath.Join(destDir, fileName))
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

func getProjectEndpoint(project string) string {
	return purpurApiBaseEndpoint + "/" + project
}

func handleAPIResponse(r *http.Response, err error) (func(out any) error, error) {
	errFunc := func(_ any) error { return err }
	if err != nil {
		slog.With(
			"error", err,
		).Debug("ParsingAPIResponse")
		return errFunc, err
	}

	defer func() {
		_ = r.Body.Close()
	}()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return errFunc, fmt.Errorf("reading response body: %v", err)
	}
	if config.GetDebug() {
		slog.With(
			"request", map[string]any{
				"url":     r.Request.URL.String(),
				"headers": r.Request.Header,
				"method":  r.Request.Method,
				"response": map[string]any{
					"body":        string(b),
					"headers":     r.Header,
					"status_code": r.StatusCode,
				},
			}).Debug("ParsingAPIResponse")
	}

	if r.StatusCode/100 != 2 {
		return errFunc, fmt.Errorf("getting purpur api response: %v", r.Status)
	}

	return func(out any) error {
		if err := json.Unmarshal(b, out); err != nil {
			return fmt.Errorf("decoding minecraft versions: %v", err)
		}
		return nil
	}, nil
}

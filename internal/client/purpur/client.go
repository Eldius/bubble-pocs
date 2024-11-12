package purpur

import (
	"encoding/json"
	"fmt"
	"github.com/eldius/bubble-pocs/internal/config"
	"io"
	"log/slog"
	"net/http"
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
		err = fmt.Errorf("getting minecraft versions: %v", err)
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

func (c *Client) GetBuildsByMineVersion(ver string) (*GetPurpurVersionsResponse, error) {
	res, err := c.c.Get("https://api.purpurmc.org/v2/" + ver)
	if err != nil {
		err = fmt.Errorf("getting purpur builds for ver %s: %v", ver, err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("getting purpur builds: %v", res.Status)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading purpur builds: %v", err)
	}
	return parseGetPurpurVersionsResponse(b, res.StatusCode)
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

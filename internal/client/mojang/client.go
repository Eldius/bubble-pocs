package mojang

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Mojang struct {
	c *http.Client
}

type MojangUsers []MojangUser

type MojangUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewMojang() *Mojang {
	return &Mojang{
		c: &http.Client{},
	}
}

func (c *Mojang) FetchUsers(users ...string) (MojangUsers, error) {
	if len(users) == 0 {
		return nil, errors.New("no users")
	}

	b, _ := json.Marshal(users)
	body := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, "https://api.minecraftservices.com/minecraft/profile/lookup/bulk/byname", body)
	if err != nil {
		err = fmt.Errorf("creating request object: %w", err)
		return nil, err
	}

	res, err := c.c.Do(req)
	if err != nil {
		err = fmt.Errorf("making request: %w", err)
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	//resBody, err := io.ReadAll(res.Body)
	//if err != nil {
	//	err = fmt.Errorf("reading response body: %w", err)
	//	return nil, err
	//}
	//
	//slog.With(
	//	slog.String("status", res.Status),
	//	slog.String("body", string(resBody)),
	//).Info("got response")

	var resp MojangUsers
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		err = fmt.Errorf("decoding response: %w", err)
		return nil, err
	}

	return resp, nil
}

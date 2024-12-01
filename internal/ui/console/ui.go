package console

import (
	"context"
	"fmt"
	"github.com/gorcon/rcon"
	"time"
)

func Start(ctx context.Context, host string, port int, password string) error {
	conn, err := rcon.Dial(
		fmt.Sprintf("%s:%d", host, port),
		password,
		rcon.SetDialTimeout(100*time.Millisecond),
	)
	if err != nil {
		err = fmt.Errorf("could not connect to %s: %w", host, err)
		return err
	}
	defer func() {
		_ = conn.Close()
	}()

	response, err := conn.Execute("help")
	if err != nil {
		err = fmt.Errorf("could not get help: %w", err)
		return err
	}

	fmt.Println(response)
	return nil
}

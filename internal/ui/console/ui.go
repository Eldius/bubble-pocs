package console

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gorcon/rcon"
	"log/slog"
	"os"
	"strings"
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

	//response, err := conn.Execute("help")
	//if err != nil {
	//	err = fmt.Errorf("could not get help: %w", err)
	//	return err
	//}
	//
	//fmt.Println(response)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("#> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			err = fmt.Errorf("could not read input: %w", err)
			return err
		}
		command = strings.TrimSuffix(command, "\n")

		log := slog.With(slog.String("command", command))

		if command == "exit" {
			return nil
		}

		resp, err := conn.Execute(command)
		if err != nil {
			err = fmt.Errorf("could not read input: %w", err)
			return err
		}

		log.With("response", resp).Debug("ExecutingCommand")
		fmt.Println(resp)
	}

	return nil
}

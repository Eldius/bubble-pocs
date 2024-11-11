
run:
	go run ./cmd/tui hello

users:
	go run ./cmd/tui users Eldius foo bar

users-styled:
	go run ./cmd/tui users Eldius foo bar --styled

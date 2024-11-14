
run:
	go run ./cmd/tui hello --debug

users:
	go run ./cmd/tui users Eldius foo bar --debug

users-styled:
	go run ./cmd/tui users Eldius foo bar --styled --debug

purpur:
	go run ./cmd/tui purpur --debug

phone:
	go run ./cmd/tui phone --debug


TEST_SERVER ?= 192.168.0.42


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

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run

snapshot-local:
	goreleaser release --snapshot --clean

release-local:
	goreleaser release --clean --skip=publish

put:
	echo 'rm ~/.bin/bubbles' | sftp $(USER)@$(TEST_SERVER)
	echo 'put ./dist/bubbles_linux_arm64_v8.0/bubbles .bin/' | sftp $(USER)@$(TEST_SERVER)

console:
	go run ./cmd/tui console \
		--host $(TEST_SERVER) \
		--port 25575 \
		--password 'MyStrongP@ss#123'

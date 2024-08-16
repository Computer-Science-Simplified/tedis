build:
	go build -o bin/tedis-cli cmd/cli/main.go

build-all: build
	GOOS=windows GOARCH=amd64 go build -o bin/tedis-cli_win_amd64.exe cmd/cli/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/tedis-cli_linux_amd64 cmd/cli/main.go

run: build
	bin/tedis-cli
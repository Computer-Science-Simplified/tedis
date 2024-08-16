analyze:
	go vet ./...
	staticcheck ./...

build: analyze
	go build -o bin/tedis cmd/api/main.go

build-all: build
	GOOS=windows GOARCH=amd64 go build -o bin/tedis_win_amd64.exe cmd/api/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/tedis_linux_amd64 cmd/api/main.go

run: build
	bin/tedis

migrate:
	go run cmd/migrate/main.go
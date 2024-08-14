analyze:
	go vet ./...
	staticcheck ./...

build: analyze
	go build -o bin/tedis cmd/api/main.go

run: build
	bin/tedis
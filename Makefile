run:
	go run ./cmd/cli/main.go sample.mp3

test:
	go test ./...

test-details:
	go test -v ./...
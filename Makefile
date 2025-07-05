run:
	go run ./cmd/cli/main.go sample.mp3

test:
	go test ./...

test-details:
	go test -v ./...

setup:
	echo "Setup the project"
	
	echo "Create tmp directory"
	mkdir -p ./tmp

	echo "Fetching testing data..."
	scripts/download.sh
	
	echo "Copying quiet saturday file as a sample media file"
	cp ./tmp/excerpts/Quiet-Saturday_fma-169775_001_00-00-17.mp3 sample.mp3

clean:
	echo "Cleaning up after myself"
	rm -rf ./tmp
	rm -f sample.mp3
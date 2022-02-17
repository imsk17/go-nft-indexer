
# Build a Binary for all major Operating Systems
build:
	@echo "Compiling for every OS and Platform"
	# 386 variants
	GOOS=freebsd GOARCH=386 go build -o dist/go-indexer-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o dist/go-indexer-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o dist/go-indexer-windows-386 main.go
	# amd64 variants
	GOOS=freebsd GOARCH=amd64 go build -o dist/go-indexer-freebsd-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/go-indexer-darwin-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/go-indexer-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/go-indexer-windows-amd64 main.go
	# arm64 variants
	GOOS=darwin GOARCH=amd64 go build -o dist/go-indexer-darwin-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/go-indexer-linux-arm64 main.go
	@echo "## Build completed successfully ##"

# Run the Program
run:
	@echo "Starting the server on the specified port"
	go run main.go

# Clean the dist/ directory
clean:
	@echo "Cleaning the dist directory \n"
	rm -rf dist/
	@echo "## Cleaned 'dist/' Successfully ##"
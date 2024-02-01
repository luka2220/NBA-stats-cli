
tidy:
	go fmt ./...
	go tidy ./...

# Build the executable binary and store in ./bin
build:
	go build -o bin/ main.go

run:
	go run main.go



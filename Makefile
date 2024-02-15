# Check for any possible errors
# Detect any possible shadowed variables
vet:
	go fmt ./...
	go vet ./...
	shadow ./...

# Add missing and remove unsed modules from go.mod
tidy:
	go fmt ./...
	go mod tidy

# Build the executable binary and store in ./bin
build:
	go build -o bin/ 

# Run the test commands in the cmd folder
test commands:
	go test -v ./cmd

# Clean up packages and files
clean:
	go clean 
	rm ./bin/nba-stats

# Execute the schedule command
schedule:
	go run main.go schedule

# Execute the standings command
standings:
	go run main.go standings

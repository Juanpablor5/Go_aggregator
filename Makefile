# run the project
run:
	go run src/main.go

# build the project
build:
	go build src/main.go

# tidies dependencies
tidy-deps:
	go mod tidy

# vendors dependencies
deps:
	go mod tidy
	go mod vendor
	
# run all tests with coverage
test-cover:
	go test ./... -cover -coverprofile=coverage.out -race

	
fmt:
	gofmt -w -s ./src

check:
	staticcheck ./...

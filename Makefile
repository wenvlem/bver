default: build

build:
	@echo "Building bver"
	@go build -ldflags="-s -w"

test: 
	@go test -coverprofile=cover.prof

view: 
	@go tool cover -html=cover.prof

.PHONY: build test view

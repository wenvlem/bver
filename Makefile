default: build

build:
	@echo "Building bver"
	@CGO_ENABLED=0 go build -ldflags="-s -w"

docker:
	@echo "Building bver docker image"
	@docker > access.log
	@docker build -t bver .

run:
	@echo "Running bver docker container"
	@docker > access.log
	@docker run -v ${PWD}/access.log:/var/log/access.log --rm -d bver

test: 
	@go test -coverprofile=cover.prof

view: 
	@go tool cover -html=cover.prof

.PHONY: build docker run test view

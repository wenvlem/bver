default: build

build:
	@echo "Building bver"
	@CGO_ENABLED=0 go build -ldflags="-s -w"

docker:
	@echo "Building bver docker image"
	@> access.log
	@docker build -t bver .

save:
	@echo "Saving bver docker image"
	@> access.log
	@docker image save -o bver-docker.tgz bver

run: clean
	@echo "Running bver docker container"
	@> access.log
	@docker run -v ${PWD}/access.log:/var/log/access.log  --name bver --rm -d bver

clean:
	@echo "Cleaning bver docker container"
	@> access.log
	@docker rm -f bver || true

test: 
	@go test -coverprofile=cover.prof -race

view: 
	@go tool cover -html=cover.prof

.PHONY: build docker run clean test view

PROJECTNAME := $(shell basename "$(PWD)")

## deps: Dowload and check the dependencies
deps:
	go mod tidy

## build: Build the docker image
build:
	docker build -t ${PROJECTNAME} .

## run: Run the image of docker
run:
	docker run -it -p 3000:3000 ${PROJECTNAME}

## test: Run the test of project
test:
	go test ./... -coverpkg=./... -coverprofile=coverage.out

## coverage: Get coverage of all tests
coverage: test
	go tool cover -func=coverage.out

## coverage-html: Generate the report in HTML of test coverage
coverage-html: test
	go tool cover -html=coverage.out -o coverage.html 

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

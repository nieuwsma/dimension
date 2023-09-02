# Copyright

# Service
NAME ?= dimension
VERSION ?= $(shell cat .version)


all: image unittest  server

binary: server

#create server binary
server:
	go build -v -o bin/server github.com/nieuwsma/system-power-capping-service/cmd/spc

#create docker image for server
image:
	docker build --pull ${DOCKER_ARGS} --tag '${NAME}:${VERSION}' -f Dockerfile .

#run all tests
test: unittest

#run unit tests
unittest:
	./runUnitTest.sh

#empty the bin
clean:
	rm bin/*

#generate lines of code in project
cloc:
	cloc --exclude-dir=vendor,test,.docker,.DS_Store,.git,.idea,bin,configs,hms-simulation-environment --exclude-ext=.json,JSON --exclude-lang=JSON .
# Copyright

# Service
NAME ?= dimension
VERSION ?= $(shell cat .version)


all: image unittest  server

binary: server autoplayertest

#create server binary
server:
	go build -v -o bin/server github.com/nieuwsma/dimension/cmd/server

#create autoplayertest binary
autoplayertest:
	go build -v -o bin/autoplay github.com/nieuwsma/dimension/cmd/autoplay


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
	cloc --exclude-dir=vendor,data_processing,test,.docker,.DS_Store,.git,.idea,bin,configs --exclude-ext=.json,JSON --exclude-lang=JSON .
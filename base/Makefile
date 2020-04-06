.PHONY: build clean docker run

GO=CGO_ENABLED=0 GO111MODULE=on go
GOCGO=CGO_ENABLED=1 GO111MODULE=on go

DOCKERS=dockers
.PHONY: $(DOCKERS)

MICROSERVICES=microservices

.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)
DOCKER_TAG=$(VERSION)-dev

GOFLAGS=-ldflags "-X project.Version=$(VERSION)"

build: $(MICROSERVICES)
builds

clean:
	rm -f $(MICROSERVICES)

run:
	cd bin && ./launch.sh

docker: $(DOCKERS)
dockerbuilds

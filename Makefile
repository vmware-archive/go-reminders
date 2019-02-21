# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

# Check that the container name that gets build is defined.
ifndef CONTAINER
$(error CONTAINER, which specifies the docker container name to build, is not set)
endif

default: cmd/go-reminders/go-reminders cmd/go-reminders/go-reminders-darwin

all: container

container: cmd/go-reminders/go-reminders
	cd build/docker; ./build.sh
.PHONY: container

cmd/go-reminders/go-reminders: go.mod $(GOFILES)
	cd cmd/go-reminders; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a --installsuffix cgo go-reminders.go

cmd/go-reminders/go-reminders-darwin: go.mod $(GOFILES)
	cd cmd/go-reminders; GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -o go-reminders-darwin --installsuffix cgo go-reminders.go 

go.mod:
	go mod init github.com/vmware/go-reminders
	for m in $$(cat forcemodules); do go get "$$m"; done

test:
	go test ./...
.PHONY: test

clean:
	cd cmd/go-reminders && \
	go clean && \
	rm -f go-reminders go-reminders-darwin && \
	go clean -modcache
.PHONY: clean

run:
	docker run --name go-reminders -d -p 8080:8080 $(CONTAINER) /go-reminders -a 172.16.78.227
.PHONY: run

stop:
	-killall go-reminders
	-docker stop go-reminders
	-docker rm go-reminders
.PHONY: stop

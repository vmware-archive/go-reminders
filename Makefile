# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

# Check that the container name that gets build is defined.
ifndef CONTAINER
$(error CONTAINER, which specifies the docker container name to build, is not set)
endif

default: cmd/go-reminders/go-reminders

all: container

container: cmd/go-reminders/go-reminders
	cd build/docker; ./build.sh
.PHONY: container

cmd/go-reminders/go-reminders: go.mod $(GOFILES)
	cd cmd/go-reminders; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a --installsuffix cgo go-reminders.go

go.mod:
	go mod init github.com/vmware/go-reminders
	for m in $$(cat forcemodules); do go get "$$m"; done

test:
	go test ./...
.PHONY: test

clean:
	go clean
	rm -f go-reminders
	go clean -modcache
.PHONY: clean

run:
	docker run -d -p 8080:8080 go-reminders /go-reminders -a 172.16.78.227
.PHONY: run

stop:
	killall go-reminders
.PHONY: stop

#
all: docker

go-reminders: imports
	cd cmd/go-reminders; CGO_ENABLED=0 go build -a --installsuffix cgo go-reminders.go
.PHONY: go-reminders

imports:
	@#cd cmd/go-reminders; go get ./...
.PHONY: imports

.PHONY: imports

docker: go-reminders
	cd build/docker/continer; docker build -t opencloudtools/go-reminders --rm=true .
.PHONY: docker

test:
	go test ./...
.PHONY: test

clean:
	go clean
	rm -f go-reminders
	echo "Cleaning up Docker Engine before building."
	docker kill $$(docker ps -a | awk '/go-reminder/ { print $$1}') || echo -
	docker rm $$(docker ps -a | awk '/go-reminder/ { print $$1}') || echo -
	docker rmi go-reminders
.PHONY: clean

run:
	docker run -d -p 8080:8080 go-reminders /go-reminders -a 172.16.78.227
.PHONY: run

stop:
	killall go-reminders
.PHONY: stop

#
all: docker

go-reminders: imports
	CGO_ENABLED=0 go build -a --installsuffix cgo .

imports:
	go get ./...

docker: go-reminders
	docker build -t opencloudtools/go-reminders --rm=true .

test:
	go test ./...

clean:
	go clean
	rm -f go-reminders
	echo "Cleaning up Docker Engine before building."
	docker kill $$(docker ps -a | awk '/go-reminder/ { print $$1}') || echo -
	docker rm $$(docker ps -a | awk '/go-reminder/ { print $$1}') || echo -
	docker rmi go-reminders

run:
	docker run -d -p 8080:8080 go-reminders /go-reminders -a 172.16.78.227

stop:
	killall go-reminders

.PHONY: go-reminders docker clean clean run stop

#
go-reminders: imports
	CGO_ENABLED=0 go build -a --installsuffix cgo .

imports:
	go get ./...

clean:
	go clean

run:
	./go-reminders >go-reminders.log 2>&1 &

stop:
	killall go-reminders

.PHONY: go-reminders clean clean run stop

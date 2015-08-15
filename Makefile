#
vmworld: imports
	CGO_ENABLED=0 go build -a --installsuffix cgo .

imports:
	go get ./...

clean:
	go clean

run:
	./vmworld >vmworld.log 2>&1 &

stop:
	killall vmworld

.PHONY: vmworld clean clean run stop

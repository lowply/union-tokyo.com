run:
	go run main.go

deps:
	dep ensure -update

build:
	go build -o main


DEP := $(shell command -v dep 2> /dev/null)

run:
	go run main.go

deps:
ifdef DEP
	dep ensure -update
else
	$(error Please install golang/dep)
endif

build:
	go build -o bin/main

deploy:build
ifdef DEPDIR
	rsync -av --delete bin public templates ${DEPDIR}
else
	$(error Usage: "DEPDIR=/path/to/dir make deploy")
endif

.PHONY: deploy

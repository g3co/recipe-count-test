PREFIX ?= /usr/local
VERSION = "v0.0.1"

all: build

build:
	mkdir -p bin
	go mod download
	GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o bin/calc
	docker build -t recipe_calc .

run:
	docker run -v $(conf):$(conf) -v $(data):$(data) -ti recipe_calc -conf=$(conf) -data=$(data)

test:
	go test ./... -v

.PHONY: all buld run test

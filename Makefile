VERSION= 1.0.0
FLAGS= -v -mod=vendor

default: test

test:
	go test -v ./... -timeout 1m

build:
	swag init
	go build $(FLAGS)
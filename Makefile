BUILD_TIME := $(shell date "+%F %T")
COMMIT_SHA1 := $(shell git rev-parse HEAD )
VERSION=0.0.1
SOURCE := ./connection.go
BUILD_NAME := connection


.PHONY:gen-code gen-swagger-doc start-service start-swagger-ui build gen-md-doc

gen-code:
	goctl api go -api ./api/connection.api -dir ./  --style=goZero

start-service:
	go run ./connection.go

build:
	go build ${LDFLAGS} ./connection.go


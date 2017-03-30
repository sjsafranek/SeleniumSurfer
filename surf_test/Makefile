##=======================================================================##
## Makefile
## Created: Wed Aug 05 14:35:14 PDT 2015 @941 /Internet Time/
# :mode=makefile:tabSize=3:indentSize=3:
## Purpose:
##======================================================================##

SHELL=/bin/bash
PROJECT_NAME = SurfTest
GPATH = $(shell pwd)

.PHONY: fmt get-deps test install build scrape clean

install: fmt get-deps test
	@GOPATH=${GPATH} go build -o surf_test main.go

build: fmt get-deps test
	@GOPATH=${GPATH} go build -o surf_test main.go

get-deps:
	@GOPATH=${GPATH} go get gopkg.in/headzoo/surf.v1
	@GOPATH=${GPATH} go get github.com/PuerkitoBio/goquery

fmt:
	@GOPATH=${GPATH} gofmt -s -w *.go

test: fmt get-deps
	@GOPATH=${GPATH} go test -v -cover -bench=. -test.benchmem

scrape:
	@find src -type d -name '.hg' -or -type d -name '.git' | xargs rm -rf

clean:
	@GOPATH=${GPATH} go clean

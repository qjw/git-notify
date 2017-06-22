.PHONY: build

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

PACKAGES = $(shell go list ./... | grep -v /vendor/)

deps: deps_backend

deps_backend:
	go get -u github.com/kardianos/govendor
	go get -u github.com/jteeuwen/go-bindata/...
	govendor sync

gen: gen_template

gen_template:
	go generate github.com/qjw/git-notify/templates/

test:
	go test -cover $(PACKAGES)

build: build_static

build_static:
	CGO_ENABLED=0 GOOS=linux  go install -a -installsuffix cgo -ldflags '${EXTLDFLAGS}' github.com/qjw/git-notify
	mkdir -p release
	cp $(GOPATH)/bin/git-notify release/
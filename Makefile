PKG = github.com/deislabs/cnab-dashboard/collector/cmd/collector-api
SHELL = bash

COMMIT ?= $(shell git rev-parse --short HEAD)
VERSION ?= $(shell git describe --tags 2> /dev/null || echo v0)
PERMALINK ?= $(shell git name-rev --name-only --tags --no-undefined HEAD &> /dev/null && echo latest || echo canary)

LDFLAGS = -w -X $(PKG)/pkg.Version=$(VERSION) -X $(PKG)/pkg.Commit=$(COMMIT)
XBUILD = CGO_ENABLED=0 go build -a -tags netgo -ldflags '$(LDFLAGS)'

default: deploy

build:
	GOOS=linux GOARCH=amd64 $(XBUILD) -o collector/collector-api $(PKG)
	docker-compose build

deploy:
	docker-compose up -d --force-recreate
	open http://localhost:8080
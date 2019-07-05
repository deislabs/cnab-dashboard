PKG = github.com/deislabs/cnab-dashboard/claims-api
SHELL = bash

COMMIT ?= $(shell git rev-parse --short HEAD)
VERSION ?= $(shell git describe --tags 2> /dev/null || echo v0)
PERMALINK ?= $(shell git name-rev --name-only --tags --no-undefined HEAD &> /dev/null && echo latest || echo canary)

LDFLAGS = -w -X $(PKG)/pkg.Version=$(VERSION) -X $(PKG)/pkg.Commit=$(COMMIT)
XBUILD = CGO_ENABLED=0 go build -a -tags netgo -ldflags '$(LDFLAGS)'

dashboard:
	docker build -t cnab-dashboard -f smashing/Dockerfile smashing/
	-@docker rm -f cnab-dashboard
	docker run --name cnab-dashboard \
		-v `pwd`/data/smashing/dashboards:/dashboards \
		-v `pwd`/data/smashing/widgets:/widgets \
		-v `pwd`/data/smashing/jobs:/jobs \
		-d -p 8080:3030 cnab-dashboard
	open http://localhost:8080

api:
	GOOS=linux GOARCH=amd64 $(XBUILD) -o claims-api/claims-api $(PKG)
	docker build -t claims-api -f claims-api/Dockerfile claims-api
	-@docker rm -f claims-api
	docker run --name claims-api -d -p 8081:3031 claims-api
	open http://localhost:8081/stats
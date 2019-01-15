##
## functions
##

rwildcard = $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))

##
## vars
##

GOPATH ?= $(HOME)/go
BIN = $(GOPATH)/bin/calcbiz
SOURCES = $(call rwildcard, ./, *.go)
OUR_SOURCES = $(filter-out $(call rwildcard,./vendor,*.go),$(SOURCES))
PROTOS = $(call rwildcard, ./, *.proto)
OUR_PROTOS = $(filter-out $(call rwildcard,./vendor,*.proto),$(PROTOS))
GENERATED_FILES = \
        $(patsubst %.proto,%.pb.go,$(PROTOS)) \
        $(call rwildcard .//, *.gen.go) \
        $(call rwildcard .//, *.pb.gw.go) \
	apidocs.swagger.json

PROTOC_OPTS = -I/protobuf:vendor:.
RUN_OPTS ?=

##
## rules
##

.PHONY: help
help:
	@echo "Make commands:"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: \
	  '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | \
	  sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | grep -v / | \
	  sed 's/^/  $(HELP_MSG_PREFIX)make /'

.PHONY: run
run: $(BIN)
	$(BIN) server $(RUN_OPTS)

.PHONY: install
install: $(BIN)
$(BIN): .generated $(OUR_SOURCES)
	packr clean
	go install -v .

.PHONY: release
release:
	packr
	go install -v .

.PHONY: clean
clean:
	rm -f $(GENERATED_FILES) .generated
	packr clean

.PHONY: _ci_prepare
_ci_prepare:
	touch $(OUR_PROTOS) $(GENERATED_FILES)
	sleep 1
	touch .generated

.PHONY: generate
generate: .generated
.generated: $(OUR_PROTOS)
	rm -f $(GENERATED_FILES)
	go mod vendor
	docker run \
	  --user="$(shell id -u)" \
	  --volume="$(PWD):/go/src/ultre.me/calcbiz" \
	  --workdir="/go/src/ultre.me/calcbiz" \
	  --entrypoint="sh" \
	  --rm \
	  pathwar/protoc:v1 \
	  -xec "make _generate"
	touch $@

.PHONY: _generate
_generate: $(GENERATED_FILES)

.PHONY: test
test: .generated
	go test -v ./...

%.pb.go: %.proto
	protoc \
	  $(PROTOC_OPTS) \
	  --grpc-gateway_out=logtostderr=true:"$(GOPATH)/src" \
	  --gogofaster_out=plugins=grpc:"$(GOPATH)/src" \
	  "$(dir $<)"/*.proto

apidocs.swagger.json: $(OUR_PROTOS)
	protoc \
	  $(PROTOC_OPTS) \
	  --swagger_out=allow_merge=true:. \
	  ./api/api.proto

.PHONY: lint
lint:
	golangci-lint run --verbose ./...

##
## production
##

.PHONY: docker.up
docker.up:
	docker-compose up -d --force-recreate --remove-orphans

.PHONY: docker.build
docker.build:
	docker build -t ultreme/calcbiz .

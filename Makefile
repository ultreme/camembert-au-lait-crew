GOPKG ?= ultre.me/calcbiz
DOCKER_IMAGE ?= ultreme/calcbiz
GOBINS ?= ./cmd/calcbiz
RUN_OPTS ?=

PRE_INSTALL_STEPS += generate
PRE_UNITTEST_STEPS += generate
PRE_TEST_STEPS += generate
PRE_BUILD_STEPS += generate
PRE_LINT_STEPS += generate
PRE_TIDY_STEPS += generate
PRE_BUMPDDEPS_STEPS += generate

all: test install

include rules.mk

.PHONY: run
run: install
	calcbiz server $(RUN_OPTS)

##
## generate
##

PROTOS_SRC := $(wildcard ./api/*.proto)
GEN_SRC := $(PROTOS_SRC) Makefile
.PHONY: generate
generate: gen.sum
gen.sum: $(GEN_SRC)
	shasum $(GEN_SRC) | sort > gen.sum.tmp
	diff -q gen.sum gen.sum.tmp || ( \
	  set -e; \
	  GO111MODULE=on go mod vendor; \
	  docker run \
	    --user=`id -u` \
	    --volume="$(PWD)/.:/go/src/ultre.me" \
	    --workdir="/go/src/ultre.me" \
	    --entrypoint="sh" \
	    --rm \
	     pathwar/protoc:4 \
	    -xec 'make generate_local'; \
	    make tidy \
	)

.PHONY: generate_local
generate_local:
	@set -e; for proto in $(PROTOS_SRC); do ( set -xe; \
	  protoc -I ./vendor/github.com/grpc-ecosystem/grpc-gateway:./api:./vendor:/protobuf --grpc-gateway_out=logtostderr=true:"$(GOPATH)/src" --gogofaster_out="plugins=grpc:$(GOPATH)/src" "$$proto" \
	); done
	goimports -w ./pkg ./cmd ./internal
	shasum $(GEN_SRC) | sort > gen.sum.tmp
	mv gen.sum.tmp gen.sum

.PHONY: clean
clean:
	rm -f gen.sum $(wildcard */*/*.pb.go */*/*.pb.gw.go)

##
## production
##

.PHONY: prod.up
prod.up:
	docker-compose build --pull
	docker-compose up -d --force-recreate --remove-orphans

.PHONY: deploy
deploy:
	ssh zrwf.m.42.am -xec 'cd ~/go/src/ultre.me/calcbiz; git pull; make prod.up'

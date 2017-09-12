SOURCE :=	$(shell find . -name "*.go")
OWN_PACKAGES := $(shell go list ./... | grep -v vendor)


calc-www: $(SOURCE)
	go build -o ./calc-www ./cmd/calc-www/main.go


run: up

up:
	docker-compose up -d --force-recreate --remove-orphans

.PHONY: dev
dev: calc-www
	./calc-www server


.PHONY: docker
docker:
	docker build -t camembertaulaitcrew/camembert-au-lait-crew .


.PHONY: test
test:
	go test -v $(OWN_PACKAGES)

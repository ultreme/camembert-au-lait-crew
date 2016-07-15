SOURCE :=	$(shell find . -name "*.go")
OWN_PACKAGES := $(shell go list ./... | grep -v vendor)


calc-www: $(SOURCE)
	go build -o ./calc-www ./cmd/calc-www/main.go


.PHONY: run
run: calc-www
	./calc-www server


.PHONY: docker
docker:
	docker build -t ultreme/camembert-au-lait-crew .


.PHONY: test
test:
	go test -v $(OWN_PACKAGES)

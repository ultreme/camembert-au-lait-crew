SOURCE :=	$(shell find . -name "*.go")


install:
	go install ./cmd/calc-www


run: up

up:
	docker-compose up -d --force-recreate --remove-orphans

.PHONY: dev
dev: install
	calc-www server


.PHONY: docker
docker:
	docker build -t ultreme/calcbiz .


.PHONY: test
test:
	go test -v ./...

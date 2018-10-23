FROM golang:1.11
COPY go.* /go/src/ultre.me/calcbiz/
WORKDIR /go/src/ultre.me/calcbiz
RUN GO111MODULE=on go mod download
COPY . /go/src/ultre.me/calcbiz
RUN GO111MODULE=on go install -v ./cmd/calc-www
CMD ["calc-www", "server"]
EXPOSE 9000

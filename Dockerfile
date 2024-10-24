FROM            golang:1.23-alpine as builder
WORKDIR         /go/src/ultre.me/calcbiz
RUN             apk --no-cache --update add npm make gcc g++ musl-dev openssl-dev git
ENV             GO111MODULE=on
COPY            go.* /go/src/ultre.me/calcbiz/
RUN             go mod download
RUN             go install github.com/gobuffalo/packr/packr
COPY            . /go/src/ultre.me/calcbiz/
RUN             make packr
RUN             make install

FROM            alpine:3.20
RUN             apk --no-cache --update add ca-certificates && update-ca-certificates
COPY            --from=builder /go/bin/calcbiz /bin/calcbiz
#COPY            ./static .
ENTRYPOINT      ["/bin/calcbiz"]
CMD             ["server"]
EXPOSE          9000 9001

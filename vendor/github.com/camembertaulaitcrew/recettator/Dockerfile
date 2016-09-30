FROM golang:1.6
COPY . /go/src/github.com/camembertaulaitcrew/recettator
WORKDIR /go/src/github.com/camembertaulaitcrew/recettator
RUN go install -v ./cmd/recettator
ENTRYPOINT ["recettator"]
CMD ["-h"]

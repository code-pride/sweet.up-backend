FROM golang:1.15.4-alpine
WORKDIR $GOPATH/src/github.com/code-pride/sweet.up
ADD . .
RUN go build -o sweet-up ./cmd/

CMD [ "./sweet-up", "-conf-path", "./docker/" ]
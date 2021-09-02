FROM golang:latest
ENV GO111MODULE=off
COPY ./ /go/src/hello_world
WORKDIR /go/src/hello_world
ENTRYPOINT [ "go", "run", "main.go" ]
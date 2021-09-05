FROM golang:latest as builder
ENV GOOS=linux
ENV CGO_ENABLED=0
COPY ./ /go/src/hello_world
WORKDIR /go/src/hello_world
RUN go build .

FROM alpine:latest
WORKDIR /usr/home
COPY --from=builder /go/src/hello_world/hello_world /usr/home
ENTRYPOINT ["./hello_world"]
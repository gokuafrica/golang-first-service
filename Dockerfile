FROM golang:latest as builder
COPY ./ /go/src/hello_world
WORKDIR /go/src/hello_world
RUN go build .
# Using multistage build because the golang image is very large in size
FROM gcr.io/distroless/base-debian10
WORKDIR /usr/home
COPY --from=builder /go/src/hello_world/hello_world /usr/home
ENTRYPOINT ["./hello_world"]
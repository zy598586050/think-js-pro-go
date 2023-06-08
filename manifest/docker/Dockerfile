FROM golang:latest AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

RUN mkdir -p think-js-pro-go

WORKDIR /think-js-pro-go

COPY . .

RUN go build main.go

FROM alpine:latest

RUN mkdir -p think-js-pro-go

WORKDIR /think-js-pro-go

COPY --from=builder main .

COPY --from=builder manifest .

COPY --from=builder resource .

ENTRYPOINT ["./main"]
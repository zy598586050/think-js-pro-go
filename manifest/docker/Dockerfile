FROM golang:latest

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN mkdir -p think-js-pro-go

WORKDIR /think-js-pro-go

COPY ../../main .

COPY ../../manifest .

COPY ../../resource .

ENTRYPOINT ["/app/main"]
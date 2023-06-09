# 第一步
FROM golang:latest AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

RUN mkdir -p app-go

WORKDIR /app-go

COPY . .

RUN go build main.go

# 第二步
FROM scratch

RUN mkdir -p think-js-pro-go

WORKDIR /think-js-pro-go

COPY --from=builder /app-go/main /main

COPY --from=builder /app-go/manifest /manifest

COPY --from=builder /app-go/resource /resource

ENTRYPOINT ["./main"]
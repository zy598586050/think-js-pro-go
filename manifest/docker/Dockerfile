FROM alpine:latest

RUN mkdir -p think-js-pro-go

WORKDIR /think-js-pro-go

COPY main .

COPY manifest .

COPY resource .

ENTRYPOINT ["./main"]
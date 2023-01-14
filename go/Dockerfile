# 1st stage build for prod env
FROM golang:1.19-alpine as builder

WORKDIR /go/src/app

RUN apk update \
&&  apk add --no-cache bash gcc musl-dev \
&&  rm -rf /var/lib/apt/lists/* \
&&  export GO111MODULE=on

COPY ./blog_app/ .

RUN go mod tidy \
&&  go build -o /go_app .

# prod env; runs built image
FROM alpine as prd

COPY --from=builder /go_app .
COPY ./blog_app/.env .

EXPOSE 8080

ENTRYPOINT [ "/go_app" ]


# dev env; live reload enabled
FROM golang:1.19-alpine as dev

WORKDIR /go/src/app

RUN apk update \
&&  apk add --no-cache bash gcc musl-dev \
&&  rm -rf /var/lib/apt/lists/* \
&&  export GO111MODULE=on

COPY ./blog_app/ .

RUN go mod tidy

EXPOSE 8080

# CMD ["./start.sh"]
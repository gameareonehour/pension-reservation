FROM golang:1.21.4-alpine

RUN apk upgrade --update && \
    apk --no-cache add tzdata git gcc libc-dev && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

ENV TZ=Asia/Tokyo

RUN cd /tmp && \
    go install github.com/cosmtrek/air@v1.45.0 && \
    go install github.com/rakyll/gotest@latest && \
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

CMD ["sh"]

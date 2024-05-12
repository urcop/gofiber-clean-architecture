FROM golang:1.21.5-alpine as gogcc

ENV GOOS=linux
ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN apk update && apk add --no-cache \
        gcc \
        musl-dev \
        git \
        build-base

FROM gogcc as builder

WORKDIR /build

COPY . .
COPY .gitconfig /etc/gitconfig
RUN go mod download && go mod verify

RUN go build -ldflags="-s -w" -o app ./cmd/app

# production stage
FROM alpine:latest

RUN apk update && apk add --no-cache \
        gcc \
        musl-dev

WORKDIR /app/

COPY --from=builder /build/api .
COPY --from=builder /build/app .
COPY --from=builder /build/config/ /config/

#CMD ["ls", "config/"]
CMD ["/app/app", "s"]
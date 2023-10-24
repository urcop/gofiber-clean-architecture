# build stage
FROM golang:1.20 as builder

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o build/app ./cmd/app

# production stage
FROM ubuntu:20.04
COPY --from=builder /app/api ./api
COPY --from=builder /app/build .
CMD ["./app", "s"]
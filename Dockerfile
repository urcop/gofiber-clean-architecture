# build stage
FROM golang:1.21.5 as builder

WORKDIR /app
COPY . .

RUN go mod download

RUN make build

# production stage
FROM ubuntu:20.04
COPY --from=builder /app/api ./api
COPY --from=builder /app/build .
COPY --from=builder /app/config/ ./config/
CMD ["./app", "s"]
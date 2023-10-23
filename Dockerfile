# build stage
FROM golang:1.20 as builder

RUN mkdir /app
WORKDIR /app
COPY . .

RUN make build

# production stage
FROM ubuntu:20.04
COPY --from=builder /app/api ./api
COPY --from=builder /app/build .
CMD ["./app", "serve"]
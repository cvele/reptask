FROM golang:1.24.1-alpine AS builder
WORKDIR /app
RUN apk update && apk add --no-cache make gcc musl-dev
COPY . .
RUN make deps
RUN make build

FROM alpine:3.21.3
WORKDIR /app
COPY --from=builder /app/bin/pack-distributor .
EXPOSE 8080
ENTRYPOINT ["./pack-distributor"]
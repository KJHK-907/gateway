FROM golang:1.23-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/gateway

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/gateway /app/gateway
EXPOSE 8081
ENTRYPOINT [ "/app/gateway" ]
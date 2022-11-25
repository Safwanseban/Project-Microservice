FROM golang:alpine AS builder 

WORKDIR /broker-service

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o brokerapp .

FROM alpine:latest

WORKDIR /root

COPY --from=builder /broker-service/brokerapp .
     

COPY . .
EXPOSE 8086
CMD ["./brokerapp"]
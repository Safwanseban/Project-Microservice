FROM golang:alpine AS builder 

WORKDIR /logger-service

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o loggerApp .

FROM alpine:latest

WORKDIR /root

COPY --from=builder /logger-service/loggerApp .
     

COPY . .

CMD ["./loggerApp"]
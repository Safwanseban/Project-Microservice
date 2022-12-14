FROM golang:alpine AS builder 

WORKDIR /broker-service

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o authApp .

FROM alpine:latest

WORKDIR /root
# ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait

# RUN chmod +x /wait
COPY --from=builder /broker-service/authApp .
     

COPY . .
EXPOSE 8086
CMD ["./authApp"]
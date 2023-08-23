FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go build -o ssm-get

FROM nginx:alpine

COPY --from=builder /app/ssm-get /usr/local/bin/ssm-get
COPY 40-nginx-ssm.sh /docker-entrypoint.d/
RUN chmod +x /docker-entrypoint.d/40-nginx-ssm.sh\
    && chmod +x /usr/local/bin/ssm-get
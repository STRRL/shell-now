# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go mod download -x

RUN go build -o shell-now ./cmd/shell-now

# Run stage
FROM nicolaka/netshoot:v0.13

COPY --from=builder /app/shell-now /usr/local/bin/shell-now

ENTRYPOINT ["shell-now"]

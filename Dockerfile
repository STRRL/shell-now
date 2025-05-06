# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go mod download -x

RUN go build -o shell-now ./cmd/shell-now

# Run stage
FROM nicolaka/netshoot:v0.13

# download cloudflared
RUN curl -L https://github.com/cloudflare/cloudflared/releases/download/2025.4.2/cloudflared-linux-amd64 -o /usr/local/bin/cloudflared
RUN chmod +x /usr/local/bin/cloudflared

# download ttyd
RUN curl -L https://github.com/tsl0922/ttyd/releases/download/1.7.7/ttyd.x86_64 -o /usr/local/bin/ttyd
RUN chmod +x /usr/local/bin/ttyd

# copy shell-now binary
COPY --from=builder /app/shell-now /usr/local/bin/shell-now

ENTRYPOINT ["shell-now"]

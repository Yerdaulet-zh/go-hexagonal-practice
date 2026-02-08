# ---------- BUILD STAGE ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy all source
COPY . .

# Build binary from cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd

# ---------- RUNTIME STAGE ----------
FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/configs ./configs

EXPOSE 2112

USER nonroot:nonroot

ENTRYPOINT ["/app/app"]

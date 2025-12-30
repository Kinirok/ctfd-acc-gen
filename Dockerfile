FROM golang:1.25-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM build AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o /app/main \
    ./cmd/ctfd-app


FROM gcr.io/distroless/static-debian12

WORKDIR /
COPY --from=builder /app/main /main

USER nonroot:nonroot
ENTRYPOINT ["/main"]
FROM golang:1.25-alpine AS build
RUN apk update && apk add --no-cache bash
#RUN apt-get install -y bash
# Проверить установку
#RUN which bash && bash --versio
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Этап 2: Тестирование
FROM build AS tester
COPY . .
RUN go test ./... -short

# Этап 3: Сборка
FROM build AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o /app/main \
    ./cmd/ctfd-app

# Этап 4: Финальный образ
FROM gcr.io/distroless/static-debian12

WORKDIR /
COPY --from=builder /app/main /main

USER nonroot:nonroot
ENTRYPOINT ["/main"]
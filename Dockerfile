FROM golang:latest AS build

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/app .

# Запускаем приложение
CMD ["./app"]
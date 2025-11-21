# ----------- STAGE 1: builder -----------
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/app/main.go


FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache postgresql-client

COPY --from=builder /app/server .

COPY migrations ./migrations

EXPOSE 8080

CMD ["./server"]
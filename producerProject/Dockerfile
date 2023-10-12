FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o cmd/app cmd/main.go

FROM scratch

COPY --from=builder /app .

ENTRYPOINT ["/cmd/app"]
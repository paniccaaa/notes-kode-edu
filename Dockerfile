FROM golang:1.23.0-alpine AS builder

WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . ./
RUN go build -o ./bin/app cmd/notes-kode-edu/main.go

FROM alpine AS runner

WORKDIR /app

COPY --from=builder /usr/local/src/bin/app /app/app

COPY ./.env /app/.env

COPY ./config/local.yaml /app/config/local.yaml

EXPOSE 8080

CMD ["/app/app"]

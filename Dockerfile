FROM golang:1.17.3

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
EXPOSE 8000
CMD ["go", "run", "./cmd/kv-search-api/main.go"]
FROM golang:latest

WORKDIR /app

COPY muzyaka .

CMD ["go", "run", "./cmd/main.go"]
FROM golang:1.21

WORKDIR /app

COPY . .

CMD ["go", "run", "main.go"]
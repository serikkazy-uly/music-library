FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init -g cmd/main.go -o docs

RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main"]
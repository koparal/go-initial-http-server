FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o server cmd/server/main.go

EXPOSE 8080

CMD ["./server"]

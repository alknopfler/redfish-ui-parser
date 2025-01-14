FROM golang:1.18

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

EXPOSE 8081

CMD ["go", "run", "main.go"]

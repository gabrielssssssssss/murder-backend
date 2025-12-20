FROM golang:1.24.9

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o murder-backend ./cmd/murder/main.go

CMD ["./murder-backend"]
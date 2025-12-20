FROM golang:1.24.10

WORKDIR /app

ENV GOPROXY=direct

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o murder-backend ./cmd/murder/main.go

RUN curl -L https://install.meilisearch.com | sh

CMD sh -c "./murder-backend & exec ./meilisearch --master-key=murder"
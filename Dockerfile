FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD cmd ./cmd/
ADD internal ./internal/
COPY .env ./
RUN go build -o ./GohCMS/ ./cmd ./internal

CMD ["./GohCMS/cmd"]
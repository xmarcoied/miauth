FROM golang:1.13-alpine
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o main
ENTRYPOINT ["./main"]
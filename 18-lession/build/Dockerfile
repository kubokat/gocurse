FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/18-lession/cmd
RUN GOOS=linux go build -o app

EXPOSE 8121

CMD ["./app"]

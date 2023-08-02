FROM golang:1.20.5

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /surveyor-service

EXPOSE 8080

CMD ["/surveyor-service"]
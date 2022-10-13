FROM golang:1.18.7-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-go-server-ping

EXPOSE 8080

CMD [ "./docker-go-server-ping" ]
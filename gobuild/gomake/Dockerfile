##
## Build
##
FROM golang:1.18.7-alpine3.16 AS build

WORKDIR /app

COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main ./goinstall

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /main /main

EXPOSE 8080

ENTRYPOINT ["/main"]
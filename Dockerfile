## Build the application
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o monkey .

## deploy the application into a lean image
FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/monkey .

CMD ["./monkey"]

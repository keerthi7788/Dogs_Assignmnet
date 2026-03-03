
# Build Stage
FROM golang:1.25-alpine AS build-stage

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/dogs-service ./cmd/dogs-service/main.go



# Execution Stage
FROM alpine:latest AS execution-stage

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=build-stage /bin/dogs-service /bin/dogs-service

COPY dogs.json .

EXPOSE 8080

CMD ["/bin/dogs-service"]
FROM golang:1.18-alpine AS build
RUN apk update
RUN apk upgrade
RUN apk add git curl
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /go/bin/main

FROM alpine
COPY --from=build /go/bin/main /app
EXPOSE 80
CMD ["/app"]

FROM golang:buster as builder
WORKDIR /app
ADD . .
RUN go env -w GO111MODULE=auto
RUN go build -o app
EXPOSE 8080
CMD ["./app"]
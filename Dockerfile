FROM golang:buster as builder
#RUN useradd -ms /bin/bash newuser
#USER newuser
#WORKDIR /home/newuser
ADD . .
RUN go env -w GO111MODULE=auto
RUN go build -o app
RUN echo Hi > /data/demo/TestFile
EXPOSE 8080
CMD ["./app"]
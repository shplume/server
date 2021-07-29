FROM golang:1.16-alpine

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /go/release

ADD . .

RUN go mod download

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]

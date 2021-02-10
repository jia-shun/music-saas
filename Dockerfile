FROM golang:latest

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"

WORKDIR $GOPATH/src/music-saas
COPY . $GOPATH/src/music-saas
RUN go build .

EXPOSE 8888
ENTRYPOINT ["./music-saas"]
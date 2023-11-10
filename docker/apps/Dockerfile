FROM golang:1.21.1-alpine

LABEL author="yuluo" \
	email="1481556636@qq.com"

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

ARG APP_HOME
VOLUME ${APP_HOME}
WORKDIR ${APP_HOME}

RUN go build -o main ${APP_HOME}/main.go

ENTRYPOINT ["sh", "main"]

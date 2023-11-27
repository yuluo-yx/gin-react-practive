FROM centos:7.9.2009

LABEL author=yuluo \
    email=yuluo829@aliyun.com \
    version="2023.11.27"

ARG APP_HOME

VOLUME ${APP_HOME}
WORKDIR ${APP_HOME}

ENTRYPOINT ["./backend/main"]

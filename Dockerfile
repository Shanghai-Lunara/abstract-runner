FROM ubuntu:18.04

LABEL maintainer="1024769485@qq.com"

WORKDIR /data/
COPY abstract-runner /data/

CMD ["/data/abstract-runner", "-alsologtostderr=true", "-v", "4"]

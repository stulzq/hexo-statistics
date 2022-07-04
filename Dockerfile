FROM alpine:3.14

WORKDIR /app

COPY bin/linux-amd64/* .
ENV HEXO_DEBUG=false

ENTRYPOINT ["./hexo-stat"]
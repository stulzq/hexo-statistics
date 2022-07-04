FROM alpine:3.14

WORKDIR /app

COPY bin/hexo-statistics-docker .
ENV HEXO_DEBUG=false

ENTRYPOINT ["./hexo-stat"]
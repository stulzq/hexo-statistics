# Hexo Statistics

Hexo blog traffic statistics service, based on redis. Use hyperloglog to count UV.

Demo: https://xcmaster.com/

## Get Start

### Binary

````shell

````

### Docker

````shell
mkdir -p /data/hexo-stat/conf

curl https://raw.githubusercontent.com/stulzq/hexo-statistics/main/conf/config.yml -o /data/hexo/stat/config.yml

# update your config

docker run --name hexo-stat \
  -v /data/hexo-stat/conf:/app/conf \
  -v /data/hexo-stat/logs:/app/logs \
  stulzq/hexo-statistics:v0.1.0

````

## TODO

- Data Export
- Data Import

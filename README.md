# go-http-nats-proxy

depends on nats.io


build docker image:
```
docker build . -t nats-proxy
```


start nats.io docker:

```
docker run -it --rm --name nats nats:latest --debug
```

start nats-proxy:

```
docker run --rm -it --name nats-proxy --link nats nats-proxy
```


usage:
```
docker run -it --rm --link nats-proxy curlimages/curl:latest curl -X POST -H "Content-Type: text/plain" --data "hello there" "http://nats-proxy/?topic=test"
```
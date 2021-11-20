# go-http-nats-proxy

depends on nats.io

run build.sh

```
./build.sh
```

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
docker run --rm -it --link nats nats-proxy
```

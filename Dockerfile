from scratch

COPY dist/nats_proxy_linux_x86_64 /app

ENTRYPOINT ["/app"]


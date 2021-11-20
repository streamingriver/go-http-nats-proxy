#!/bin/bash

# Remove old binaries (if any)
rm -rf dist
mkdir dist

env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "dist/nats_proxy_linux_x86_64"          # Linux 64bit
env GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -o "dist/nats_proxy_linux_arm"       # Linux armv5/armel/arm (it also works on armv6)
env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o "dist/nats_proxy_linux_armhf"     # Linux armv7/armhf
env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o "dist/nats_proxy_linux_aarch64"         # Linux armv8/aarch64


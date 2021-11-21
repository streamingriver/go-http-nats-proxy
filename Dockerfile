FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /proxy


FROM scratch
WORKDIR /
COPY --from=0 /proxy /

EXPOSE 80

ENTRYPOINT ["/proxy"]


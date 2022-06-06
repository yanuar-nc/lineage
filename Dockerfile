FROM golang:alpine as builder

WORKDIR /app

# Add the source code
ENV SRC_DIR=/app

ADD . $SRC_DIR

# RUN go mod download

RUN  CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o go-boiler-plate .

EXPOSE 9100

ENTRYPOINT ["/app/go-boiler-plate"]
FROM jrottenberg/ffmpeg:3.3-alpine as ffmpeg

FROM golang:1.9-alpine as builder
COPY --from=ffmpeg /usr/local/ /usr/local/
RUN apk add --no-cache --update ca-certificates libcrypto1.0 libssl1.0 libgomp expat libgcc libstdc++
COPY . /go/src/github.com/flaneurtv/worker-ffmpeg/

RUN cd /go/src/github.com/flaneurtv/worker-ffmpeg \
    && apk add --no-cache \
        git \
        gettext \
        jq \
    && go get -d -v ./... \
    && GOOS=linux GOARCH=amd64 go build -o "/go/bin/worker-ffmpeg" ./main.go 

WORKDIR /go

CMD ash

FROM jrottenberg/ffmpeg:3.3-alpine as ffmpeg

FROM node:8-alpine
COPY --from=ffmpeg /usr/local/ /usr/local/
RUN apk  add --no-cache --update ca-certificates libcrypto1.0 libssl1.0 libgomp expat

WORKDIR /srv
CMD ash

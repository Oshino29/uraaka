# build layer
FROM golang:1.17-alpine AS build
RUN apk add build-base git
WORKDIR /app
COPY . .

RUN go mod download
RUN go build



# app layer
FROM alpine:latest
WORKDIR /app
EXPOSE 8080

ENV USER=uraaka
ENV GROUP=uraaka
ENV UID=1000
ENV GID=1000
RUN addgroup -g $GID $GROUP &&\
    adduser \
    --disabled-password \
    --gecos "" \
    --home "${pwd}" \
    --ingroup "$GROUP" \
    --no-create-home \
    --uid "$UID" \
    "$USER"
ENV TZ="Asia/Shanghai"
RUN apk add --no-cache tzdata

COPY --from=build --chown=uraaka:uraaka /app/uraaka /app/uraaka
COPY --from=build --chown=uraaka:uraaka /app/templates /app/templates

RUN mkdir /data && chown uraaka:uraaka /data
CMD [ "/app/uraaka" ]
FROM golang:alpine AS build
RUN apk add build-base git
WORKDIR /src
COPY . .
RUN go build -o uraaka

ARG user=uraaka
ARG group=uraaka
ARG uid=1000
ARG gid=1000
RUN groupadd -g ${gid} ${group} && useradd -u ${uid} -G ${group} -s /bin/sh -D ${user}


FROM alpine:latest
COPY --from=build /src/uraaka /app/uraaka
WORKDIR /app
EXPOSE 8080

USER ${user}
CMD ["/app/uraaka"]
FROM golang:1.16-alpine as backend

ARG GIT_BRANCH

ENV CGO_ENABLED=0

ADD . /build
WORKDIR /build

RUN apk add --no-cache --update git tzdata ca-certificates

RUN \
    version=${GIT_BRANCH}-$(date +%Y%m%dT%H:%M:%S) && \
    echo "version=$version" && \
    go build -o /build/api-service -ldflags "-X main.revision=${version} -s -w"

FROM ghcr.io/umputun/baseimage/app:v1.6.1 as base

FROM scratch
#todo: enable reproxy SPA mod

COPY --from=backend /build/api-service /srv/api-service
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

WORKDIR /srv
ENTRYPOINT ["/srv/api-service"]

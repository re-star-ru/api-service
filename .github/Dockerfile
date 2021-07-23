FROM golang:1.16-alpine as backend

ARG GIT_BRANCH
ARG GITHUB_SHA
ARG CI
ARG GIT_BRANCH

ENV CGO_ENABLED=0

ADD . /build
WORKDIR /build

RUN \
    version=${GIT_BRANCH}-$(date +%Y%m%dT%H:%M:%S) && \
    echo "version=$version" && \
    cd cmd/app && \
    go build -o /build/rp -ldfl ags "-X main.revision=${version} -s -w"

FROM scratch
COPY --from=backend /build /srv
WORKDIR /srv
ENTRYPOINT ["/srv/reproxy"]
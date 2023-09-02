# Copyright

# Dockerfile for building Dimension server
# Build base just has the packages installed we need.
FROM golang:1.20-alpine AS build-base
RUN echo "GOPATH is set to: $GOPATH"

RUN set -ex \
    && apk -U upgrade \
    && apk add build-base

# Base copies in the files we need to test/build.
FROM build-base AS base


RUN go env -w GO111MODULE=auto

# Copy all the necessary files to the image.
COPY cmd $GOPATH/src/github.com/nieuwsma/dimension/cmd
COPY vendor $GOPATH/src/github.com/nieuwsma/dimension/vendor
COPY pkg $GOPATH/src/github.com/nieuwsma/dimension/pkg
COPY internal $GOPATH/src/github.com/nieuwsma/dimension/internal
COPY .version $GOPATH/src/github.com/nieuwsma/dimension/.version

### Build Stage ###
FROM base AS builder

RUN set -ex && go build -v -o /usr/local/bin/server github.com/nieuwsma/dimension/cmd/server

### Final Stage ###

FROM alpine:3
LABEL maintainer="Nieuwsma"
EXPOSE 28007
STOPSIGNAL SIGTERM

RUN set -ex \
    && apk -U upgrade

# Get the system-power-capping-service from the builder stage.
COPY --from=builder /usr/local/bin/server /usr/local/bin/.
COPY .version /


ENV LOG_LEVEL "TRACE"
ENV PORT "8080"

#nobody 65534:65534
USER 65534:65534

CMD ["sh", "-c", "server $@"]

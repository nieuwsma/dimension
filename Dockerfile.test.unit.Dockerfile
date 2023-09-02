# Copyright 
# This file only exists as a means to run tests in an automated fashion.

FROM golang:1.20-alpine

RUN set -ex \
    && apk -U upgrade \
    && apk add build-base 
ENV LOG_LEVEL "WARN"
ENV TEST_DIR "$GOPATH/src/github.com/nieuwsma/dimension/pkg/logic/test_cases"


RUN go env -w GO111MODULE=auto

COPY cmd $GOPATH/src/github.com/nieuwsma/dimension/cmd
COPY vendor $GOPATH/src/github.com/nieuwsma/dimension/vendor
COPY pkg $GOPATH/src/github.com/nieuwsma/dimension/pkg
COPY internal $GOPATH/src/github.com/nieuwsma/dimension/internal
COPY .version $GOPATH/src/github.com/nieuwsma/dimension/.version

CMD set -ex \
    && go version \
    && go test -cover -v -o dimension github.com/nieuwsma/dimension/internal/api \
    && go test -cover -v -o dimension github.com/nieuwsma/dimension/pkg/logic

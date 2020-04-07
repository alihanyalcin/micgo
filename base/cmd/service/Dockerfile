FROM golang:1.14.1-alpine AS builder

ENV GO111MODULE=on

WORKDIR /go/src/{project}

RUN apk update && apk add make && apk add bash git

COPY go.mod .

RUN go mod download

COPY . .
RUN make cmd/{servicename}/{servicename}

FROM scratch

EXPOSE {portnumber}

COPY --from=builder /bin/bash /bin/bash
COPY --from=builder /go/src/{project}/cmd/{servicename}/{servicename} /
COPY --from=builder /go/src/{project}/cmd/{servicename}/res/docker/configuration.toml /res/configuration.toml

ENTRYPOINT ["/{servicename}","--confdir=/res"]
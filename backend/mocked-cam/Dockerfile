FROM golang:1.23-alpine as buildbase

RUN apk add git build-base

WORKDIR /mocked-cam

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go mod vendor
RUN go build  -o /usr/local/bin/mocked-cam /mocked-cam


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/mocked-cam /usr/local/bin/mocked-cam

RUN apk add --no-cache ca-certificates

ENTRYPOINT ["mocked-cam"]

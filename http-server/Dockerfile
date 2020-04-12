# STAGE 1: Build

FROM golang:alpine AS build

RUN apk update \
  && apk add git

COPY . /go/src/app

# Don't do this in production! Use vendoring instead.
RUN go get -v app

RUN go install app

RUN ls -l /go/bin


# STAGE 2: Deployment
FROM alpine

COPY --from=build /go/bin/app /go/bin/server

EXPOSE 8080


ENTRYPOINT ["/go/bin/server"]
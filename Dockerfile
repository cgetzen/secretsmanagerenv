FROM ruby AS base
ENV GOPATH=/go

RUN apt-get update

FROM base AS update
RUN mkdir /go
RUN apt-get -qqy install golang-go

FROM update AS go

RUN gem install fpm

FROM go AS gem

RUN go get github.com/urfave/cli
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/aws/session
RUN go get github.com/aws/aws-sdk-go/service/secretsmanager

FROM golang
MAINTAINER Jonathan Hurter <jonathan@hurter.fr>

RUN go get github.com/johnsudaar/gitngo
RUN go install github.com/johnsudaar/gitngo

WORKDIR /go/src/github.com/johnsudaar/gitngo
ENTRYPOINT ["/go/bin/gitngo"]

EXPOSE 8080

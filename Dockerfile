FROM golang
MAINTAINER jesse.miller@adops.com

ADD . /go/src/github.com/oaodev/acs-view
RUN go install github.com/oaodev/acs-view

EXPOSE 8080
CMD ["/go/bin/acs-view"]
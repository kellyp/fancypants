FROM golang

ADD . /go/fancy

WORKDIR /go/fancy

RUN go install

EXPOSE 8080/tcp

CMD ["/go/bin/fancypants"]
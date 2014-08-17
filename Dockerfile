FROM golang

RUN go get github.com/codegangsta/gin

ENV PORT 3000
EXPOSE 3000

ADD . /go/src/github.com/ddollar/tug-example
WORKDIR /go/src/github.com/ddollar/tug-example
RUN go get .

CMD ["gin", "-p", "$PORT"]

FROM golang

RUN go get github.com/codegangsta/gin

ADD . /go/src/github.com/ddollar/tug-example
WORKDIR /go/src/github.com/ddollar/tug-example
RUN go get .

EXPOSE 3000

CMD ["gin", "-p", "3000"]

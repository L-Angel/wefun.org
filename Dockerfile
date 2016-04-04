FROM golang
MAINTAINER astaxie xiemengjun@gmail.com

RUN go get github.com/astaxie/beego

RUN go get github.com/DeanChina/wefun.org

WORKDIR /go/src/github.com/DeanChina/wefun.org

RUN go build

EXPOSE 8080

CMD ["./wefun.org"]

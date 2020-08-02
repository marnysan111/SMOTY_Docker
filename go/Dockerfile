FROM golang:latest

RUN mkdir /go/src/test 

WORKDIR /go/src/test

add . /go/src/test

RUN go get github.com/gin-gonic/gin \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/gin-contrib/sessions
FROM golang:latest

RUN mkdir /go/src/smoty

add . /go/src/smoty

WORKDIR /go/src/smoty/code

RUN go get github.com/gin-gonic/gin \
    && go get gorm.io/gorm \
    && go get gorm.io/driver/mysql \
    && go get github.com/gin-contrib/sessions

    
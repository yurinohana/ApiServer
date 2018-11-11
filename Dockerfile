FROM golang:latest

WORKDIR /go/src/github.com/yurinohana/ApiServer

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/aws/aws-sdk-go/aws
RUN go get -u github.com/aws/aws-sdk-go/aws/session
RUN go get -u github.com/guregu/dynamo

CMD go build


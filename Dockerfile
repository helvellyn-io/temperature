FROM golang:1.14.9-alpine
RUN mkdir /go/src/demo
ADD . go.mod go.sum main.go /go/src/demo/
WORKDIR /go/src/demo
RUN go build 
ENTRYPOINT [ "./demo" ] 

FROM golang:1.5.3
MAINTAINER Gwenn Etourneau <gwenn.etourneau@gmail.com>
RUN mkdir -p /gopath/bin/
RUN go get github.com/tools/godep
RUN go get github.com/Pivotal-Japan/firehose-to-fluentd
RUN cd $GOPATH/src/github.com/Pivotal-Japan/firehose-to-fluentd  \
    ; CGO_ENABLED=0 godep go build  -a --installsuffix cgo --ldflags="-s"
RUN cp $GOPATH/src/github.com/Pivotal-Japan/firehose-to-fluentd/firehose-to-fluentd /gopath/bin/
COPY Dockerfile.final /gopath/bin/Dockerfile
RUN ls -lah /gopath/bin/

CMD docker build  -t getourneau/firehose-to-fluentd /gopath/bin

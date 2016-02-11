FROM golang:1.5.3
MAINTAINER Gwenn Etourneau <gwenn.etourneau@gmail.com>
RUN mkdir -p /gopath/bin/
RUN go get github.com/tools/godep
ADD firehose-to-fluentd.tgz   $GOPATH/src/github.com/shinji62/firehose-to-fluentd/
RUN cd $GOPATH/src/github.com/shinji62/firehose-to-fluentd \
    ; CGO_ENABLED=0 godep go build  -a --installsuffix cgo --ldflags="-s"
RUN cd $GOPATH/src/github.com/shinji62/firehose-to-fluentd 
RUN cp $GOPATH/src/github.com/shinji62/firehose-to-fluentd/firehose-to-fluentd /gopath/bin/
COPY Dockerfile.final /gopath/bin/Dockerfile

CMD docker build  --no-cache -t getourneau/firehose-to-fluentd-dev /gopath/bin/

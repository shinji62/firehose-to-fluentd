all: test linux32 linux64 darwin64

test:
	ginkgo -r .

linux32:
	GOARCH=386 GOOS=linux godep go build -o dist/linux/i386/firehose-to-fluentd

linux64:
	GOARCH=amd64 GOOS=linux godep go build  -o dist/linux/amd64/firehose-to-fluentd

darwin64:
	GOARCH=amd64 GOOS=darwin godep go build  -o dist/darwin/amd64/firehose-to-fluentd

docker-dev:
	$(SHELL) ./Docker/build-dev.sh

docker-final:
	$(SHELL) ./Docker/build.sh     

clean:
	$(RM) dist/*
	$(RM) *.prof

all: clean build

clean:
	rm -f chariot:

build: genproto
	go build .

syncbuild: syncproto genproto
	go build .

genproto:
	./proto/gen_go.sh

syncproto:
	cd proto && git pull origin master

init:
	git submodule update --init

install: init genproto
	glide install

apollotunnel:
	ssh -fNL 4834:apollo.marathon.mesos:4834 centos@52.42.186.11

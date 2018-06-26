DEPENDENCIES = cli.go config.go main.go
SHA := $(shell git rev-parse HEAD)

build: $(DEPENDENCIES)
	mkdir -p bin
	go build -o bin/smenv

deb:
	docker build -t ubuntu-smenv:${SHA} .
	docker run -v $(shell pwd):/smenv -it ubuntu-smenv:${SHA} /bin/bash -c \
	"cd /smenv && make && fpm -s dir -t deb -v 0.0.1 -p bin -n smenv bin/smenv"

help:
	# make - builds smenv in bin
	# make deb - builds a deb in bin

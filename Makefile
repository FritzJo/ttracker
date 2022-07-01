PREFIX ?= /usr/local
APPNAME ?= ttracker

build:
	echo "Building $(APPNAME) binary"
	go build -o bin/$(APPNAME) *.go

install: build
	echo "$(PREFIX)"
	cp ./bin/$(APPNAME) $(PREFIX)/bin/$(APPNAME)

uninstall:
	echo "Removing pacheck"
	rm -vf $(PREFIX)/bin/$(APPNAME)
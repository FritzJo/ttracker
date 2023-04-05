PREFIX ?= /usr/local
APPNAME ?= ttracker

test:
	echo "Running tests for $(APPNAME)"
	go test -race -vet=off ./tests

build:
	echo "Building $(APPNAME) binary"
	go build -o bin/$(APPNAME) *.go

install: build
	echo "Installing to $(PREFIX)"
	cp ./bin/$(APPNAME) $(PREFIX)/bin/$(APPNAME)

uninstall:
	echo "Removing $(APPNAME)"
	rm -vf $(PREFIX)/bin/$(APPNAME)
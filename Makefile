PREFIX ?= /usr/local
APPNAME ?= ttracker

SUPPORTED_ARCHS := amd64 arm64
SUPPORTED_OS := linux darwin windows

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

release:
	$(foreach arch,$(SUPPORTED_ARCHS),\
		$(foreach os,$(SUPPORTED_OS),\
			echo "Compiling: $(os) ($(arch))";\
			if [ "$(os)" = "windows" ]; then \
				GOOS=$(os) GOARCH=$(arch) go build -o bin/$(APPNAME)_$(os)_$(arch).exe;\
			fi; \
			if [ "$(os)" != "windows" ]; then \
				GOOS=$(os) GOARCH=$(arch) go build -o bin/$(APPNAME)_$(os)_$(arch);\
			fi; \
		)\
	)
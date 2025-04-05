# Variables
BINARY_NAME=trans

# Make sure to use tabs, not spaces
all: build

build:
	go build -o $(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)

distclean: clean

install:
	install -D -m 0755 $(BINARY_NAME) $(DESTDIR)/usr/bin/$(BINARY_NAME)

test:
	go test -v ./...

.PHONY: all build clean distclean install test
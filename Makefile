GOCMD=go
GOFLAGS=-ldflags "-s -w"

.PHONY: install build clean

all: install build

install:
	go get -u github.com/shirou/gopsutil/process
	go get -u github.com/shirou/gopsutil/net
	go get -u github.com/fatih/color

build: gohunting.go 
	$(GOCMD) build $(GOFLAGS) $<

clean:
	rm -f gohunting


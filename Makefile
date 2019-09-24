GO=go
GO_FLAGS=-ldflags "-s -w"

.PHONY: install build clean

all: build

install:
	go get -u github.com/shirou/gopsutil/process
	go get -u github.com/shirou/gopsutil/net
	go get -u github.com/fatih/color

build: install
	$(GO) build $(GO_FLAGS) gohunting.go

clean:
	rm -f gohunting


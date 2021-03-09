SHELL=PATH='$(PATH)' /bin/sh

GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s'



# enable second expansion
.SECONDEXPANSION:

.PHONY: all

BINDIR=./bin
NAME=did-lightnode

resdir=./webpages

all: lnx mac arm


staticfile:
	go-bindata -o $(resdir)/webfs/webfs.go -pkg=webfs $(resdir)/html/dist/...

mac: staticfile
	GOOS=darwin go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).mac
arm: staticfile
	GOOS=linux GOARM=7 GOARCH=arm go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).arm
lnx: staticfile
	GOOS=linux go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).lnx

clean:
	rm $(BINDIR)/$(NAME)

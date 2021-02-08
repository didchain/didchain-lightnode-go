SHELL=PATH='$(PATH)' /bin/sh

GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s'



# enable second expansion
.SECONDEXPANSION:

.PHONY: all

BINDIR=./bin
NAME=did-lightnode

all: lnx mac arm

mac:
	GOOS=darwin go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).mac
arm:
	GOOS=linux GOARM=7 GOARCH=arm go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).arm
lnx:
	GOOS=linux go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).lnx

clean:
	rm $(BINDIR)/$(NAME)

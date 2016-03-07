VERSION := $(shell bin/genversion --version)
COMMIT := $(shell bin/genversion --commit)
ldflags = "-X "github.com/ipfs/go-ipfs/repo/config".CurrentVersionNumber=$(VERSION) \
-X "github.com/ipfs/go-ipfs/repo/config".CurrentCommit=$(COMMIT)"

all: install

install:
	go install -ldflags=$(ldflags)

build:
	go build -i -ldflags=$(ldflags)

nofuse:
	go install -tags nofuse -ldflags=$(ldflags)

clean:
	go clean -ldflags=$(ldflags)

uninstall:
	go clean -i -ldflags=$(ldflags)

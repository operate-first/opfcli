EXE = opfcli-$(shell go env GOOS)-$(shell go env GOARCH)
SRCS = $(shell find . -type f -name '*.go')
PKG = $(shell go list)
VERSION = $(shell git describe --tags --exact-match 2> /dev/null || echo unknown)
COMMIT = $(shell git rev-parse --short=10 HEAD)
DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%S")

GOLDFLAGS = \
	    -X '$(PKG)/version.BuildVersion=$(VERSION)' \
	    -X '$(PKG)/version.BuildHash=$(COMMIT)' \
	    -X '$(PKG)/version.BuildDate=$(DATE)'

all: $(EXE)

$(EXE): $(SRCS)
	go build -o $@ -ldflags "$(GOLDFLAGS)"

clean:
	rm -f $(EXE)


DEFAULT: all

BLDDIR=build
APPS=main test

all: $(APPS)

$(BLDDIR)/%:
	@mkdir -p $(dir $@)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -a -tags netgo -ldflags '-w -extldflags "-static"' -o $@ ./bin/$*

$(APPS): %: $(BLDDIR)/%

DEFAULT: all

BLDDIR = build
APPS = main test

$(BLDDIR)/main: $(wildcard ./bin/main/*.go ./*/*.go)  # 这些文件变更，main才能重新构建
$(BLDDIR)/test: $(wildcard ./bin/test/*.go ./*/*.go)

all: $(APPS)

$(BLDDIR)/%:
	@mkdir -p $(dir $@)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -a -tags netgo -ldflags '-w -extldflags "-static"' -o $@ ./bin/$*

$(APPS): %: $(BLDDIR)/%
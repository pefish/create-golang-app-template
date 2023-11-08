
DEFAULT: build-cur

ifeq ($(GOPATH),)
  GOPATH = $(HOME)/go
endif

build-cur:
	GOPATH=$(GOPATH) go install github.com/pefish/go-build-tool/cmd/...@latest
	$(GOPATH)/bin/go-build-tool

install: build-cur
	sudo install -C ./build/bin/linux/template /usr/local/bin/template

install-service: install
	sudo mkdir -p /etc/systemd/system
	sudo install -C -m 0644 ./script/template.service /etc/systemd/system/template.service
	sudo systemctl daemon-reload
	@echo
	@echo "template service installed."


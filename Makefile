
DEFAULT: all

all:
	make build_linux

build_linux:
	rm -rf ./build/
	mkdir ./build/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -a -tags netgo -ldflags '-w -extldflags "-static"' -o build/ $(TARGET)
	@echo "Done building."

run:
	go run $(TARGET)

build_wasm:
	rm -rf ./build/
	mkdir ./build/
	GOARCH=wasm GOOS=js go build -o ./build/ $(TARGET)
	@echo "Done building."

run_wasm:
	GOOS=js GOARCH=wasm go run -exec="$(shell go env GOROOT)/misc/wasm/go_js_wasm_exec" $(TARGET)

node_exec_wasm:
	node $(shell go env GOROOT)/misc/wasm/wasm_exec.js $(TARGET)

test_wasm:
	GOOS=js GOARCH=wasm go test -exec="$(shell go env GOROOT)/misc/wasm/go_js_wasm_exec" $(TARGET)
	@echo "Done testing."
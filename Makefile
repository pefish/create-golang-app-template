
DEFAULT: all

all:
	rm -rf ./build/ && go build -o build/example ./bin/example/
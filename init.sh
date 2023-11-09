#!/bin/bash

cat go.mod | sed "s@template@${NAME}@g" > temp && rm -rf go.mod && mv temp go.mod

cat README.md | sed "s@XXX@${PROJECT_NAME}@g" > temp && rm -rf README.md && mv temp README.md

cat README_zh-cn.md | sed "s@XXX@${PROJECT_NAME}@g" > temp && rm -rf README_zh-cn.md && mv temp README_zh-cn.md

cat Dockerfile | sed "s@template@${PROJECT_NAME}@g" > temp && rm -rf Dockerfile && mv temp Dockerfile

mv ./cmd/template ./cmd/"${PROJECT_NAME}"

cat version/version.go | sed "s@template@${PROJECT_NAME}@g" > temp && rm -rf version/version.go && mv temp version/version.go

cat cmd/template/main.go | sed "s@template@${PROJECT_NAME}@g" > temp && rm -rf cmd/template/main.go && mv temp cmd/template/main.go

cat cmd/template/command/default.go | sed "s@template@${PROJECT_NAME}@g" > temp && rm -rf cmd/template/command/default.go && mv temp cmd/template/command/default.go

cat pkg/task/test.go | sed "s@template@${PROJECT_NAME}@g" > temp && rm -rf pkg/task/test.go && mv temp pkg/task/test.go

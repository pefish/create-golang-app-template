#!/bin/bash

cat go.mod | sed "s@create_golang_app_template@${NAME}@g" > temp && rm -rf go.mod && mv temp go.mod

cat README.md | sed "s@XXX@${PROJECT_NAME}@g" > temp && rm -rf README.md && mv temp README.md

cat README_zh-cn.md | sed "s@XXX@${PROJECT_NAME}@g" > temp && rm -rf README_zh-cn.md && mv temp README_zh-cn.md

cat Dockerfile | sed "s@template@${PROJECT_NAME}@g" > temp && rm -rf Dockerfile && mv temp Dockerfile


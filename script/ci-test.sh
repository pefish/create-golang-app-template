#!/bin/bash

set -euxo pipefail

projectName="username-app-name"
configPath="${HOME}/data/username/app-name-test"


git reset --hard && git pull && git checkout test && git pull

imageName="${projectName}:`git rev-parse --short HEAD`"

if [[ "`sudo docker images -q ${imageName} 2> /dev/null`" == "" ]]; then
  sudo docker build -t ${imageName} .
fi

sudo docker stop ${projectName}-test && sudo docker rm ${projectName}-test

sudo docker run --name ${projectName}-test -d -v ${configPath}:/app/config ${imageName}



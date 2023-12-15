#!/bin/bash

set -euxo pipefail

projectName="_username_-_app-name_"
configPath="${HOME}/data/_username_/_app-name_-prod"


git reset --hard && git pull && git checkout main && git pull

imageName="${projectName}:`git rev-parse --short HEAD`"

if [[ "`sudo docker images -q ${imageName} 2> /dev/null`" == "" ]]; then
  sudo docker build -t ${imageName} .
fi

sudo docker stop ${projectName}-prod && sudo docker rm ${projectName}-prod

sudo docker run --name ${projectName}-prod -d -v ${configPath}:/app/config ${imageName}

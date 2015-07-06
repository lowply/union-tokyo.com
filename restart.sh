#!/bin/bash

cd $(dirname $0)

logger(){
	echo "$(date): ${1}" | tee -a logs/restart.log
}

type docker > /dev/null 2>&1 || { logger "docker is not installed"; exit 1; }
type git > /dev/null 2>&1 || { logger "git is not installed"; exit 1; }

USERNAME="lowply"
APPNAME="union-tokyo.com"
CPORT="3000"
HPORT="127.0.0.1:3000"

logger "--- ${1} ---"
logger "restart.sh kicked by webhook.fixture.jp"

logger "git pull ..."
git pull

logger "docker build ..."
docker build -t ${USERNAME}/${APPNAME} .

logger "restarting container ..."
docker stop ${APPNAME}
docker rm ${APPNAME}
docker run -d --name ${APPNAME} -p ${HPORT}:${CPORT} ${USERNAME}/${APPNAME}

logger "done"

#!/bin/bash

docker build -t lowply/union-tokyo.com .
docker stop union-tokyo.com
docker rm union-tokyo.com
docker run -d --name union-tokyo.com --publish 3000:3000 lowply/union-tokyo.com

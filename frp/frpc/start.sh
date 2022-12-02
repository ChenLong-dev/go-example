#!/bin/bash

EXEC=frpc

ROOT=/home/frp_0.45.0_linux_amd64

echo "start ${EXEC} ..."

cd ${ROOT}

pwd

echo "====== exec ======"

./${EXEC} -c ${EXEC}.ini >/dev/null 2>&1 &

echo "====== after ======"

ps -ef | grep ${EXEC}
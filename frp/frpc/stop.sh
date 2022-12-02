#!/bin/bash

EXEC=frpc

echo "stop ${EXEC} ..."

echo "====== before ======"

ps -ef | grep ${EXEC} | grep -v grep

ps -ef | grep ${EXEC} | grep -v grep | awk '{print $1}'

echo "====== exec ======"

ps -ef | grep ${EXEC} | grep -v grep | awk '{print $1}' | xargs kill -9

echo "====== after ======"

ps -ef | grep ${EXEC}
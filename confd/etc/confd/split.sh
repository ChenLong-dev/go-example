#!/bin/bash

PREFIX_LINE="==>"
SUFFIX_LINE="<=="
READ_FLAG="read"
NOREAD_FLAG="noread"
ROOT_PATH="/tmp/impauth/authserver/etc"
BASE_DIR="base.templ"
APP_DIR="app.templ"

#echo -e "\033[30m ### 30:黑   ### \033[0m"
#echo -e "\033[31m ### 31:红   ### \033[0m"
#echo -e "\033[32m ### 32:绿   ### \033[0m"
#echo -e "\033[33m ### 33:黄   ### \033[0m"
#echo -e "\033[34m ### 34:蓝色 ### \033[0m"
#echo -e "\033[35m ### 35:紫色 ### \033[0m"
#echo -e "\033[36m ### 36:深绿 ### \033[0m"
#echo -e "\033[37m ### 37:白色 ### \033[0m"

function log(){
echo $1 >> /tmp/split.log
}

function check_file()
{
  file_path=$1
  if [ ! -f ${file_path} ]; then
    echo -e "\033[31m ### [check_file] ${file_path} is not exist!   ### \033[0m"
    log "[check_file] ${file_path} is not exist!"
    exit 1
  fi
}

function del_app_file()
{
  path=$(dirname $1)
  del_dir=${path}/${APP_DIR}
  if [ -d ${del_dir} ]; then
    rm -rf ${del_dir}/*
    echo -e "\033[32m ### [del_file] del file in ${del_dir}  ### \033[0m"
    log "[del_file] del file in ${del_dir}"
  fi
}

function check_dir_and_exist()
{
  dir_path=$1
  if [ x${dir_path} == x"" ]; then
    log "x${dir_path} == x''"
    return
  fi
  if [ ! -d ${dir_path} ]; then
    echo -e "\033[31m ### [check_dir_and_exist] ${dir_path} is not exist!   ### \033[0m"
    log "[check_dir_and_exist] ${dir_path} is not exist!"
    mkdir -p ${dir_path}
  fi
}

##############################################################################
function write_file()
{
  start=$1
  end=$2
  src_file=$3
  dest_file=$4
  echo -e "\033[32m ### [write_file] split start: ${start} end: ${end}, file_path:${dir}/${filename}   ### \033[0m"
  log "[write_file] split start: ${start} end: ${end}, file_path:${dir}/${filename}"

  if [ ! -f ${dest_file} ]; then
    rm -rf ${dest_file}
  fi
  sed -n "${start},${end}p" ${src_file} >> ${dest_file}
}

function read_file()
{
  line_num=0
  count=0
  dir=""
  filename=""
  is_read=${NOREAD_FLAG}
  # 逐行读取文件内容
  while read line; do
    line_num=$((line_num+1))
    if [ x$(echo ${line} | grep "${PREFIX_LINE}") != x"" ]; then
      is_read=${READ_FLAG}
    fi
    if [ x$(echo ${line} | grep "${SUFFIX_LINE}") != x"" ]; then
      start=$((line_num-count+2))
      end=$((line_num-1))
      write_file ${start} ${end} ${1} ${dir}/${filename}

      is_read=${NOREAD_FLAG}
      count=0
      dir=""
      filename=""
    fi

    if [[ x${is_read} == x${READ_FLAG} ]]; then
          count=$((count+1))
          #echo "+++ [${line_num}-${count}]:${line}"
          if [ ${count} -eq 2 ]; then
            filename=`echo $(basename ${line}) | sed 's/\r//g'`
            if [[ ${line} == "base@"* ]]; then
              dir=${ROOT_PATH}/${BASE_DIR}
            else
              dir=${ROOT_PATH}/${APP_DIR}
            fi
            check_dir_and_exist ${dir}
          fi
        fi
      done < "$1"
}

function main() {
    check_file $1
    del_app_file $1
    read_file $1
}

main $@
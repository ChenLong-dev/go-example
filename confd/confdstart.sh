#!/bin/sh
#@auth cl
#@time 20230911

#

WORKSPACE_PATH=`pwd`
CONF_DIR=/etc/confd
BACKEND_TYPE=etcdv3
BACKEND_NODE=http://localhost:2379

function check_confd() {
  if which confd >/dev/null; then
    echo -e "\033[32m ### confd is exist!  ### \033[0m"
  else
    echo -e "\033[31m ### confd not found!  ### \033[0m"
    exit 1
  fi
}

function check_result()
{
  result=$1
  object=$2
  if [ ${result} -ne 0 ]; then
    echo -e "\033[31m ### [check_result] result is failed! result:${result}, object:${object}   ### \033[0m"
    exit 1
  fi
  echo -e "\033[32m ### [check_result] result is success! result:${result}, object:${object}   ### \033[0m"
}

function check_dir() {
  file_path=$1
  if [ ! -d ${file_path} ] ; then
    echo -e "\033[31m ### [check_dir] ${file_path} is not exist!   ### \033[0m"
    exit 1
  fi
}
############################################################################


function start_confd() {
  confd_dir=$1
  if [ x${confd_dir} == x"" ]; then
    confd_dir=/etc/confd
  fi
  confd -watch -backend ${BACKEND_TYPE} -node ${BACKEND_NODE} -confdir ${confd_dir}
  check_result $? "start confd:${confd_dir}"

}

function main() {
  confd_dir=${WORKSPACE_PATH}/etc/confd
  check_dir ${confd_dir}/conf.d
  check_dir ${confd_dir}/templates
  echo -e "\033[32m ### confd_dir:${confd_dir}  ### \033[0m"

  check_confd
  start_confd ${confd_dir}
}

main $@
#!/bin/bash

function iam::test::demoFun1()
{
    echo "这是我的第一个 shell 函数!"
    return `expr 1 + 1`
}

# $* 指以一个单字符串显示所有向脚本传递的参数
if [[ "$*" =~ iam::test:: ]];then
  eval $*
fi


#demoFun1


#Header="-HContent-Type: application/json"
#CCURL="curl -f -s -XPOST"
#UCURL="curl -f -s -XPUT"
#RCURL="curl -f -s -XGET"
#DCURL="curl -f -s -XDELETE"
#
## iam::test:policy()
#a()
#{
#  echo "adsf"
#  ${DCURL} http://${IAM_APISERVER_URL}/v1/policies; echo
#}
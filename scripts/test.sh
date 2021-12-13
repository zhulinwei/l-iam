#!/bin/bash

IAM_APISERVER_URL=${IAM_APISERVER_HOST}:${IAM_APISERVER_PORT}

Header="-HContent-Type: application/json"
CCURL="curl -f -s --request POST" # Create
UCURL="curl -f -s --request PUT" # Update
RCURL="curl -f -s --request GET" # Retrieve
DCURL="curl -f -s --request DELETE" # Delete

iam::test::ping()
{
  ${RCURL} "${Header}" http://${IAM_APISERVER_URL}/v1/ping
}

iam::test::user()
{
  # 检查是否存在name=zlw的用户
  list=$(${RCURL} "${Header}" "http://${IAM_APISERVER_URL}/api/v1/users?page=1&size=100&name=zlw")
  userId=$(echo ${list} | jq '.data.list[0].id')

  if [[ -n "$userId" && "$userId" != "null" ]];then
    echo '删除历史数据--->'
    ${DCURL} "${Header}" "http://${IAM_APISERVER_URL}/api/v1/users/${userId}";echo;
    echo 'not null'
  fi

  echo '创建用户--->'
  user=$(${CCURL} "${Header}" http://${IAM_APISERVER_URL}/api/v1/users \
  -d '{"username": "zlw", "password": "passwd", "phone": "1234", "email": "zlw@com"}')
  echo "${user}"

  echo '查询用户--->'
  userId=$(echo ${user} | jq '.data.id')
  ${RCURL} "${Header}" "http://${IAM_APISERVER_URL}/api/v1/users/${userId}";echo;

  echo '更新用户--->'
  ${UCURL} "${Header}" http://${IAM_APISERVER_URL}/api/v1/users/${userId} \
  -d '{"username": "zlw-new", "password": "passwd", "phone": "1234", "email": "zlw@com"}';echo;
}

iam::test::policy()
{
  list=$(${RCURL} "${Header}" "http://${IAM_APISERVER_URL}/api/v1/policies?page=1&size=100&name=zlw")
  policyId=$(echo ${list} | jq '.data.list[0].id')

  # 如果policyId不为空，则代表name=zlw的策略存在，需要先删除
  if [[ -n "$policyId" && "$policyId" != "null" ]];then
    echo '删除历史数据--->'
    ${DCURL} "${Header}" "http://${IAM_APISERVER_URL}/api/v1/policies/${policyId}";echo;
  fi

  echo '创建策略--->'
  policy=$(${CCURL} "${Header}" http://${IAM_APISERVER_URL}/api/v1/policies \
  -d '{"name":"zlw","policy":{"description":"One policy to rule them all.","subjects":["users:<peter|ken>","users:maria","groups:admins"],"actions":["delete","<create|update>"],"effect":"allow","resources":["resources:articles:<.*>","resources:printer"],"conditions":{"remoteIPAddress":{"type":"CIDRCondition","options":{"cidr":"192.168.0.1/16"}}}}}')
  echo "${policy}"

  echo '查询用户--->'
  policyId=$(echo "${policy}" | jq '.data.id')
  ${RCURL} "${Header}" "http://${IAM_APISERVER_URL}/api/v1/policies/${policyId}";echo;

  echo '更新策略--->'
  policy=$(${UCURL} "${Header}" "http://${IAM_APISERVER_URL}/api/v1/policies/${policyId}" \
  -d '{"name":"zlw-new","policy":{"description":"One policy to rule them all.","subjects":["users:<peter|ken>","users:maria","groups:admins"],"actions":["delete","<create|update>"],"effect":"allow","resources":["resources:articles:<.*>","resources:printer"],"conditions":{"remoteIPAddress":{"type":"CIDRCondition","options":{"cidr":"192.168.0.1/16"}}}}}')
  echo "${policy}"
}

# $* 指以一个单字符串显示所有向脚本传递的参数
if [[ "$*" =~ iam::test:: ]];then
  eval $*
fi

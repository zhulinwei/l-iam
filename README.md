# l-iam

> l-iam 是极客专栏 [Go语言项目开发实战](https://time.geekbang.org/column/intro/100079601) 的练习代码，专栏官方的项目仓库请参考 [marmotedu/iam](https://github.com/marmotedu/iam)   


## 架构

1. 通过web端请求api server对密钥/策略进行CRUD操作；
2. api server收到请求后，除了自身会持久化数据，还会通过Redis发送SecretChanged/PolicyChanged消息
3. authz server在启动的过程中会通过grpc接口向api server获取所有密钥和策略，同时会监听Redis，收到Changed消息后会重新同步密钥和策略消息

## 附录

### 阿里云RAM

访问控制RAM（Resource Access Management）是阿里云提供的一项管理用户身份与资源访问权限的服务，详情可以参考[阿里云RAM](https://help.aliyun.com/product/28625.html) 。

### 腾讯云CAM

访问管理（Cloud Access Management，CAM）是腾讯云提供的一套 Web 服务，用于帮助客户安全地管理腾讯云账户的访问权限，资源管理和使用权限，详情可以参考[腾讯云CAM](https://cloud.tencent.com/document/product/598) 。

### 华为云IAM

统一身份认证（Identity and Access Management，简称IAM）是华为云提供权限管理的基础服务，可以帮助您安全地控制云服务和资源的访问权限，详情可以参考[华为云IAM](https://support.huaweicloud.com/productdesc-iam/iam_01_0026.html) 。
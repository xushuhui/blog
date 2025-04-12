---
title: Go 搭建内网代理
published: 2020-01-03 13:59:58
tags: ["Go"]
categories: ["Go"]
---
## 前言

gitlab地址：gitlab.example.com:82
代理地址： test.com

## 修改host

## nginx配置

### conf

```sh
server {
    listen 80;
    server_name www.test.com test.com;
  
    # ssl_certificate           ../cert/cert.crt; 
    # ssl_certificate_key       ../cert/cert.key; 

    error_log logs/error.log;   #指定错误日志文件路径
    access_log logs/access.log; #指定访问日志文件路径
      
    if ($args ~* "^go-get=1") {
         return 200 "<!DOCTYPE html><html><head><meta name='go-import' content='test.com$uri git http://gitlab.example.com:82$uri.git'></head></html>";
   }

 
}
```

## go环境变量配置

```sh
go env -w GONOPROXY=test.com
```

## 安装

go get test.com/grestful/session

---
title: "Redis 核心技术（1）基础架构"
date: 2020-08-10T14:44:08+08:00
tags: ["Redis","NoSQL"]
categories: ["Redis"]
draft: true
---

## 基础架构

一个键值数据库包括了访问框架、索引模块、操作模块和存储模块四部分

![cover](https://tvax2.sinaimg.cn/large/a616b9a4gy1gilk4ebrnlj21tk2xb1ik.jpg)

### 访问模块

访问模式通常有两种：一种是通过函数库调用的方式供外部应用使用，比如，上图中的 libsimplekv.so，就是以动态链接库的形式链接到我们自己的程序中，提供键值存储功能；

另一种是通过网络框架以 Socket 通信的形式对外提供键值对操作，这种形式可以提供广泛的键值存储服务

RocksDB 以动态链接库的形式使用，而 Memcached 和 Redis 则是通过网络框架访问

### 索引模块

数据库解析了客户端发来的请求，知道了要进行的键值对操作，此时，需要查找所要操作的键值对是否存在，这依赖于键值数据库的索引模块。

索引的作用是让键值数据库根据 key 找到相应 value 的存储位置，进而执行操作。

Memcached 和 Redis 采用哈希表作为 key-value 索引，而 RocksDB 则采用跳表作为内存中 key-value 的索引。

内存键值数据库（例如 Redis）采用哈希表作为索引，其键值数据基本都是保存在内存中的，而内存的高性能随机访问特性可以很好地与哈希表 O(1) 的操作复杂度相匹配。

Redis 通过索引找到一个 key 所对应的 value 后，仍然需要从 value 的复杂结构（例如集合和列表）中进一步找到我们实际需要的数据，这个操作的效率本身就依赖于它们的实现结构。

### 总结

![image](https://tva1.sinaimg.cn/large/a616b9a4gy1gill5foskyj243r2rie82.jpg)

1、访问模块接收到客户端的 socket 通信，得到客户端发的命令 get k1

2、索引模块用 hash 表做索引找到键 k1 的位置，从存储模块的各类数据结构中找到数据，返回给客户端

3、如果是 set 和 del 操作，需要分配和释放内存

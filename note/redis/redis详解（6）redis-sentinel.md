---
title: Redis 详解（6）redis sentinel

date: 2018-03-12 14:51:29
tags: ["redis", "NoSQL"]
categories: ["Redis"]
---

## 主从复制高可用的问题

### 手动故障转移

* 指定一个 slave 执行 slave of no one
* 指定其他 slave 执行 slave of new master
* 客户端连接新 master

### 写能力和存储能力受限

## 架构说明

### redis sentinel 故障转移

* 多个 sentinel 发现并确认 master 有问题
* 选举出一个 sentinel 作为领导
* 选出一个 slave 作为 master
* 通知其余 slave 成为新 master 的 slave
* 通知客户端主从变化
* 等待老的 master 复活成为新 master 的 slave

## 安装配置

* 配置开启主从节点
* 配置开启 sentinel 监控主节点
* 实际情况多机器部署
* 详细配置节点

### redis 主节点

#### 启动

```
redis-server redis-7000.conf
```

#### 配置

```
port 7000
daemonize yes
pidfile /var/run/redis-7000.pid
logfile "7000.log"
dir "opt/soft/redis/data"
```

### redis 从节点

#### 启动

```
redis-server redis-7001.conf
redis-server redis-7002.conf
```

#### 配置

```
port 7001
daemonize yes
pidfile /var/run/redis-7001.pid
logfile "7001.log"
dir "opt/soft/redis/data"
slaveof 127.0.0.1 7000

port 7002
daemonize yes
pidfile /var/run/redis-7002.pid
logfile "7002.log"
dir "opt/soft/redis/data"
slaveof 127.0.0.1 7000
```

### redis sentinel 主要配置

```
port ${port}
dir "opt/soft/redis/data"
logfile "${port}.log"
sentinel monitor mymaster 127.0.0.1 7000 2
sentinel down-after-milliseconds mymaster 30000
sentinel parallel-syncs mymaster 1
sentinel failover-timeout mymaster 180000
```

## 客户端连接

## 实现原理

## 常见开放运维问题

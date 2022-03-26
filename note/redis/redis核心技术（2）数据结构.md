---
title: "Redis 核心技术（2）数据结构"
date: 2020-09-10T15:53:33+08:00
tags: ["Redis","NoSQL"]
categories: ["Redis"]
draft: true
---

## 数据结构

|数据类型|数据结构|
|-|-|
|String（字符串）|简单动态字符串|
|List（列表）|双向链表，压缩列表|
|Hash（哈希）|哈希表，压缩列表|
|Set（集合） |哈希表，整数数组|
|Sorted Set（有序集合）|跳表、压缩列表|

## 全局哈希表

为了实现从键到值的快速访问，Redis 使用了一个哈希表来保存所有键值对。

一个哈希表，其实就是一个数组，数组的每个元素称为一个哈希桶。所以，一个哈希表是由多个哈希桶组成的，每个哈希桶中保存了键值对数据。

哈希桶中的元素保存的并不是值本身，而是指向具体值的指针。不管值是 String，还是集合类型，哈希桶中的元素都是指向它们的指针。

> 这个哈希表更像是一个用 key 来定位 value 的索引，真正存储 value 的数据结构不是这个哈希表。

![image](https://tva3.sinaimg.cn/large/a616b9a4gy1gilmfn8vu6j21d90obn4g.jpg)

## 哈希冲突

Redis 的 hash 表是全局的，所以当写入大量的 key 时，将会带来哈希冲突。

Redis 解决 hash 冲突的方式，是链式哈希：同一个哈希桶中的多个元素用一个链表来保存。

当哈希冲突链过长时，Redis 会对 hash 表进行 rehash 操作。rehash 就是增加现有的 hash 桶数量，分散 entry 元素。

## rehash

为了使 rehash 操作更高效，Redis 默认使用了两个全局哈希表：哈希表 1 和哈希表 2，起始时 hash2 没有分配空间。

随着数据增多，Redis 执行分三步执行 rehash：
- 给 hash2 分配更大的内存空间，如是 hash1 的两倍
- 把 hash1 中的数据重新映射并拷贝到哈希表 2 中
- 释放 hash1 的空间

## 渐进式 rehash

由于步骤 2 重新映射非常耗时，会阻塞其他命令。

redis 从集中迁移数据，改成每处理一个请求时，就从 hash1 中的第一个索引位置，顺带将这个索引位置上的所有 entries 拷贝到 hash2 中。

## 压缩列表，跳表的特点

![image](https://tvax3.sinaimg.cn/large/a616b9a4gy1gilnzzzbgxj217z08r0v6.jpg)

压缩列表类似于一个数组，不同的是：压缩列表在表头有三个字段 zlbytes,zltail 和 zllen 分别表示长度，列表尾的偏移量和列表中的 entry 的个数，压缩列表尾部还有一个 zlend，表示列表结束。所以压缩列表定位第一个和最后一个是 O(1), 但其他就是 O(n)。

![image](https://tvax2.sinaimg.cn/large/a616b9a4gy1gilo0xibb2j21g60va478.jpg)

跳表：是在链表的基础上增加了多级索引，通过索引的几次跳转，实现数据快速定位。

## 总结

- 数组和压缩列表可以提升内存利用率，因为他们的数据结构紧凑
- 数组对 CPU 高速缓存支持友好，当数据元素超过阈值时，会转为 hash 和跳表，保证查询效率

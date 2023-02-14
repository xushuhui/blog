---
title: 'Redis源码解析SDS'
tags: ["redis"]
categories: ["redis"]
date: "2023-02-13T16:12:55+08:00"
toc: true
---
## 前言
字符串底层由简单动态字符串（Simple Dynamic Strings, SDS）是Redis的基本数据结构之一，用于存储字符串和整型数据。
SDS兼容C语言标准字符串处理函数，且在此基础上保证了二进制安全。

SDS如何兼容C语言字符串？如何保证二进制安全？

SDS对象中的buf是一个柔性数组，上层调用时，SDS直接返回了buf。由于buf是直接指向内容的指针，故兼容C语言函数。而当真正读取内容时，SDS会通过len来限制读取长度，而非“\0”，保证了二进制安全

## 总结
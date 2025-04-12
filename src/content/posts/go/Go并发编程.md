---
title: Go并发编程
published: 2022-01-03 13:59:58
tags: ["Go"]
categories: ["Go"]
---
## Mutex
### Mutex几种状态
- mutexLocked 
互斥锁的锁定状态
- mutexWoken
从正常模式被唤醒
- mutexStarving
当前的互斥锁进入饥饿状态
- waitersCount
当前互斥锁上等待的Goroutine个数

### Mutex正常模式和饥饿模式
#### 正常模式（非公平锁）
正常模式下，使用等待锁的goroutine按照先进先出的顺序等待。唤醒的Gorotine不会直接拥有锁，而是会和新请求goroutine竞争锁。新请求的goroutine更容易抢占，因为它正在CPU上执行，所以刚刚唤醒的goroutine有很大可能在锁竞争中失败，在这种情况下，这个被唤醒的goroutine会加入到等待队列的前面

## gRPC

### gRPC 是什么

“A high-performance, open-source universal RPC framework”
- 多语言：语言中立，支持多种语言。
- 轻量级、高性能：序列化支持 PB(Protocol Buffer) 和 JSON，PB 是一种语言无关的高性能序列化框架。
- 可插拔
- IDL：基于文件定义服务，通过 proto3 工具生成指定语言的数据结构、服务端接口以及客户端 Stub。
![image](https://tva3.sinaimg.cn/large/a616b9a4ly1gmndyioxogj20va0kiadm.jpg)

### 设计理念

- 移动端：基于标准的 HTTP2 设计，支持双向流、消息头压缩、单 TCP 的多路复用、服务端推送等特性，这些特性使得 gRPC 在移动端设备上更加省电和节省网络流量。
- 服务而非对象、消息而非引用：促进微服务的系统间粗粒度消息交互设计理念。
- 负载无关的：不同的服务需要使用不同的消息类型和编码，例如 protocol buffers、JSON、XML 和 Thrift。
- 流：Streaming API。
- 阻塞式和非阻塞式：支持异步和同步处理在客户端和服务端间交互的消息序列。
- 元数据交换：常见的横切关注点，如认证或跟踪，依赖数据交换。
- 标准化状态码：客户端通常以有限的方式响应 API 调用返回的错误。
![image](https://tva1.sinaimg.cn/large/a616b9a4ly1gmndzc07eoj20u50pqgr0.jpg)

### HealthCheck

gRPC 有一个标准的健康检测协议，在 gRPC 的所有语言实现中基本都提供了生成代码和用于设置运行状态的功能。
主动健康检查 health check，可以在服务提供者服务不稳定时，被消费者所感知，临时从负载均衡中摘除，减少错误请求。当服务提供者重新稳定后，health check 成功，重新加入到消费者的负载均衡，恢复请求。health check，同样也被用于外挂方式的容器健康检测，或者流量检测 (k8s liveness & readiness)。
![image](https://tva1.sinaimg.cn/large/a616b9a4ly1gmne0jeaorj20tl0m8wgt.jpg)

## 服务发现

### 客户端发现

一个服务实例被启动时，它的网络地址会被写到注册表上；当服务实例终止时，再从注册表中删除；这个服务实例的注册表通过心跳机制动态刷新；
客户端使用一个负载均衡算法，去选择一个可用的服务实例，来响应这个请求。

>直连，比服务端服务发现少一次网络跳转，Consumer 需要内置特定的服务发现客户端和发现逻辑。

![image](https://tvax3.sinaimg.cn/large/a616b9a4ly1gmne1x6jskj21dw0o5gvp.jpg)

### 服务端发现

客户端通过负载均衡器向一个服务发送请求，这个负载均衡器会查询服务注册表，并将请求路由到可用的服务实例上。服务实例在服务注册表上被注册和注销 (Consul Template+Nginx，kubernetes+etcd)。
![image](https://tvax4.sinaimg.cn/large/a616b9a4ly1gmne4dytfbj21fs0npwo8.jpg)

> Consumer 无需关注服务发现具体细节，只需知道服务的 DNS 域名即可，支持异构语言开发，需要基础设施支撑，多了一次网络跳转，可能有性能损失。

微服务的核心是去中心化，我们使用客户端发现模式。

### 使用

早期我们使用最熟悉的 Zookeeper 作为服务发现，但是实际场景是海量服务发现和注册，服务状态可以弱一致，需要的是 AP 系统。
- 分布式协调服务（要求任何时刻对 ZooKeeper 的访问请求能得到一致的数据，从而牺牲可用性）。
- 网络抖动或网络分区会导致的 master 节点因为其他节点失去联系而重新选举或超过半数不可用导致服务注册发现瘫痪。
- 大量服务长连接导致性能瓶颈。
我们参考了 Eureka 实现了自己的 AP 发现服务，试想两个场景，牺牲一致性，最终一致性的情况：
- 注册的事件延迟
- 注销的事件延迟

**server 启动时需要互相加载**
通过 Family(appid) 和 Addr(IP:Port) 定位实例，除此之外还可以附加更多的元数据：权重、染色标签、集群等。
> appid: 使用三段式命名，business.service.xxx

Provider  注册后定期 (30s) 心跳一次，注册， 心跳，下线都需要进行同步，注册和下线需要进行长轮询推送。
> 新启动节点，需要 load cache，JVM 预热。
> 故障时，Provider 不建议重启和发布。

Consumer 启动时拉取实例，发起 30s 长轮询。
> 故障时，需要 client 侧 cache 节点信息。

Server 定期 (60s) 检测失效 (90s) 的实例，失效则剔除。短时间里丢失了大量的心跳连接 (15 分钟内心跳低于期望值*85%)，开启自我保护，保留过期服务不删除。
![image](https://tva1.sinaimg.cn/large/a616b9a4ly1gmne8rgg13j20zr0hltdd.jpg)

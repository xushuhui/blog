# 如何封装一个非阻塞式 Java 网络库

## 1. 前言

到目前为止，我们已经学完了 Java 网络编程的所有关键知识点，如果能将这些知识点灵活应用到项目中，那再好不过了。本小节将从实战的角度出发，展示如何基于 Java NIO 封装一个非阻塞式网络库。核心是采用事件反应器模型，将复杂的 SocketChannel、ServerSocketChannel、ByteBuffer、Selector 进行抽象，将繁琐的数据读写操作加以封装，以便应用程序调用。

本项目提供的完整代码路径：

[https://github.com/haska1025/imooc-sock-core-tech/tree/master/java_netprogramming](https://github.com/haska1025/imooc-sock-core-tech/tree/master/java_netprogramming)

## 2. 系统类图

在 OOD 的思想体系下，现实世界中无论是具体的、还是抽象的事物，都是**对象**，现实世界就是由**对象**组成的。**对象**有自己的**属性**和**行为**，对象之间可以产生联系。把具有相同**属性**和**行为**的对象叫做同一类对象，抽象出**类**。OOD 的思想非常符合人的思维方式，更容易建模和抽象。如果你对 OOD 不是很熟悉的话，可以查阅相关资料学习。

我们现在的工作是要把 Java NIO 中的各个组件加以抽象，抽象出一个或若干个类，通过 UML 展现出来，类图如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnywbf2kj60jp09eaaw02)

下来我们就解释一下以上几个接口和实现类的功能：

|类名|功能|
|----|----|
|Poller           |Java NIO Selector 事件多路复用机制的封装，实现事件循环机制，是一个事件反应器，是一个功能实现类                    |
|IOHandler        |是 Poller 的配套类，响应 accept、connect、read、write 事件，是一个 Java 接口                                        |
|SocketHandler    |是一个功能类，实现了 IOHandler 接口，将一些通用的逻辑加以封装                                                     |
|Acceptor         |实现 TCP 服务器监听功能，继承自 SocketHandler                                                                     |
|TcpHandler       |TCP 客户端、服务器逻辑的封装，继承自 SocketHandler。主要完成客户端连接，服务器接收新连接，数据收发，关闭连接的功能|
|IOAdapter        |事件循环机制向应用层提供的一个回调接口，一般由 IOHandler 调用                                                     |
|AbstractAdapter  |是一个 Java 接口，主要完成通用逻辑，实现了 IOAdapter 接口                                                         |
|Listener         |事件监听接口，主要实现线程切换的功能                                                                              |
|CustomEventObject|代表一个具体事件，是 Listener 的配套类                                                                            |
|IOThread         |对 Java 线程的封装，聚合了 Poller 功能。我们说过 Java 的 Selector 其实是同步的，需要一个线程调用 select 监听事件  |
|ThreadPool       |对 IOThread 的封装，提供线程池功能                                                                                |

## 3. 接口设计

软件的接口是指软件模块对外提供的一组函数或者方法，目的是让别的模块访问本模块的功能，以达到组件复用的目的。根据模块逻辑复杂度的不同，接口由分为：系统接口、子系统接口、模块接口、子模块接口、类接口。关于模块、子模块、类的应用都非常灵活，我们这里认为类就是最小的模块。

本小节所说的接口是指 Java 接口。我们抽象了三个 Java 接口：Listener、IOHandler、IOAdapter，还有一个功能类 Poller。现在对每个接口中的方法加以说明：

* Poller

|类名|接口名|描述|
|----|------|----|
|Poller|register|将 IOHandler 实例添加到 Poller|
|      |start   |启动一个 Poller 实例          |
|      |close   |停止一个 Poller 实例          |
|      |poll    |Poller 进入事件循环           |

* IOHandler

|类名|接口名|描述|
|----|------|----|
|IOHandler|handle_read     |是一个回调方法，当 Poller 监听到某个 IOHandler 注册的**读**事件触发时，调用 handle_read            |
|         |handle_write    |是一个回调方法，当 Poller 监听到某个 IOHandler 注册的**写**事件触发时，调用 handle_write           |
|         |handle_accept   |是一个回调方法，当 Poller 监听到某个 IOHandler 注册的**accept**事件触发时，调用 handle_accept      |
|         |handle_connected|是一个回调方法，当 Poller 监听到某个 IOHandler 注册的**connected**事件触发时，调用 handle_connected|
|         |getSocketChannel|用于获取 IOHandler 对应的 SocketChannel 对象                                                       |

IOHandler 是一个抽象接口，TcpHandler 需要实现此接口，当然你也可以实现其他协议，只需要扩展 IOHandler 的接口即可。

* IOAdapter

|类名|接口名|描述|
|----|------|----|
|IOAdapter|onAccept        |当 IOHandler 收到一个新的 TCP 连接时，在 handle_read 中回调此方法      |
|         |onConnected     |当 io_hanIOHandlerdler 完成异步连接时，在 handle_connected 中回调此方法|
|         |onRead          |IOHandler 会在 handle_read 中回调此方法                                |
|         |onWrite         |IOHandler 会在 handle_write 中回调此方法                               |
|         |onClose         |IOHandler 收到连接被关闭时，回调此方法                                 |
|         |setSocketHandler|向 IOAdapter 设置一个 IOHandler 对象                                   |

应用层需要实现 IOAdapter 的接口，并且要覆盖接口中的方法，完成数据收发。

* Listener

|类名|接口名|描述|
|----|------|----|
|Listener|process|处理异步事件|

## 4. 总结

本小节主要是将前面小节介绍的 Java NIO 相关模块进行一个抽象，形成一个非阻塞的 Java 网络库。采用的设计思想就是依赖倒转，将复杂的网络编程细节进行封装，让应用程序员不需要关注这些复杂的机制。

我们抽象出了一组接口和一组功能类，并对类的功能和接口中的方法进行了一一说明。本小节可以说是对整个系列内容的总结和实践。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

---
title: Java从零开始（147）Java Async SocketChannel
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java AsynchronousSocketChannel 介绍


## 1. 前言

Java NIO 可以编写高性能服务器，所依赖的 I/O 事件分发机制是 Selector。Selector 的工作原理就是有一个线程会调用 Selector 的 select 方法，然后进入阻塞状态，等待事件的发生。一旦有 I/O 事件发生，阻塞在 select 方法上的线程会返回，然后进行事件分发。其本质还是一个同步实现。

本小节将要介绍 Java 7 中引入的完全异步的编程方法，核心组件是 AsynchronousServerSocketChannel 和 AsynchronousSocketChannel 两个类，分别用来编写服务器和客户端程序。 AsynchronousServerSocketChannel 和 AsynchronousSocketChannel 是在 java.nio.channels 包中引入的。

## 2. 基于 Future 编写服务器程序

创建一个 AsynchronousServerSocketChannel 服务器的步骤如下：

* 创建 AsynchronousServerSocketChannel 的实例，需要通过它提供的工厂方法 open 完成。如下：

```java
AsynchronousServerSocketChannel server = AsynchronousServerSocketChannel.open()
```

* 将 AsynchronousServerSocketChannel 绑定在一个本地 **IP 地址**或者**端口**。

```java
server.bind(new InetSocketAddress("127.0.0.1", PORT));
```

* 向 AsynchronousServerSocketChannel 投递一个 accept 操作。accept 调用会立即返回，不会阻塞调用线程。accept 的返回值是一个 Future 对象。

```java
Future<AsynchronousSocketChannel> acceptFuture = server.accept();
```

* 通过 Future 对象的 get 方法获取新的连接对象，返回值是 AsynchronousSocketChannel 类型的对象。注意，Future 对象的 get 方法会阻塞调用线程。get 方法接收一个 timeout 参数。

```java
 AsynchronousSocketChannel client = acceptFuture.get(10, TimeUnit.SECONDS);
```

* 调用 AsynchronousSocketChannel 的 read 方法，投递一个 read 事件。注意：read 接收的参数是 ByteBuffer。read 是异步调用，不会阻塞线程。Future 的 get 调用会阻塞调用线程。

```java
ByteBuffer inBuffer = ByteBuffer.allocate(128);
Future<Integer> readResult = client.read(inBuffer);
System.out.println("Do something");
readResult.get();
```

* 调用 AsynchronousSocketChannel 的 write 方法，投递一个 write 事件。注意：write 接收的参数是 ByteBuffer。write 是异步调用，不会阻塞线程。Future 的 get 调用会阻塞调用线程。

```java
Future<Integer> writeResult = client.write(inBuffer);
writeResult.get();
```

服务器完整代码：

```java
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.AsynchronousServerSocketChannel;
import java.nio.channels.AsynchronousSocketChannel;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
public class AsyncServer {
    private static final int PORT =56002;

    public static void main(String[] args) {
        try (AsynchronousServerSocketChannel server = AsynchronousServerSocketChannel.open()){
            server.bind(new InetSocketAddress("127.0.0.1", PORT));

            Future<AsynchronousSocketChannel> acceptFuture = server.accept();
            AsynchronousSocketChannel client = acceptFuture.get(10, TimeUnit.SECONDS);

            if (client != null && client.isOpen()){
                ByteBuffer inBuffer = ByteBuffer.allocate(128);
                Future<Integer> readResult = client.read(inBuffer);
                System.out.println("Do something");
                readResult.get();

                inBuffer.flip();
                Future<Integer> writeResult = client.write(inBuffer);
                writeResult.get();
            }

            client.close();
        } catch (Exception e) {
            e.printStackTrace();
        }

    }
}
```

## 3. 基于 Future 编写客户端程序

编写客户端程序，首先是创建 AsynchronousSocketChannel 实例，通过它的 open 方法完成。然后调用 AsynchronousSocketChannel 的 connect 方法连接服务器，同样是异步调用，不会阻塞调用线程。调用 Future 的 get 方法获取连接结果。剩下客户端数据收发逻辑和服务器的数据收发逻辑一致。

客户端完整代码：

```java
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.AsynchronousSocketChannel;
import java.util.concurrent.Future;

public class AsyncClient {
    private static final int PORT =56002;
    public static void main(String[] args) {
        try (AsynchronousSocketChannel client = AsynchronousSocketChannel.open()) {
            Future<Void> result = client.connect(new InetSocketAddress("127.0.0.1", PORT));
            System.out.println("Async connect the server");
            result.get();

            String reqMessage = "Hello server!";
            ByteBuffer reqBuffer = ByteBuffer.wrap(reqMessage.getBytes());
            Future<Integer> writeResult = client.write(reqBuffer);
            System.out.println("Async send to server:" + reqMessage);
            writeResult.get();

            ByteBuffer inBuffer = ByteBuffer.allocate(128);
            Future<Integer> readResult = client.read(inBuffer);
            readResult.get();
            System.out.println("Async recv from server:" + new String(inBuffer.array()).trim());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```

## 4. 异步 I/O 操作说明

异步 Socket 编程的一个关键点是：AsynchronousServerSocketChannel 和 AsynchronousSocketChannel 提供的一组 I/O 操作是异步的，方法调用完后会立即返回，而不会关心操作是否完成，并不会阻塞调用线程。如果要想获取 I/O 操作的结果，可以通过 Future 的方式，或者是 CompletionHandler 的方式。

下面列举的 connect、accept、read、write 四组 I/O 方法，返回值是 Future 对象的 I/O 方法，前面已经介绍。还有就是需要传入一个 attachment 参数和一个 CompletionHandler 参数，这是基于完成例程的方式。

* connect 异步操作

```java
public abstract Future<Void> connect(SocketAddress remote);
public abstract <A> void connect(SocketAddress remote,
                                     A attachment,
                                     CompletionHandler<Void,? super A> handler);
```

* accept 异步操作

```java
public abstract Future<AsynchronousSocketChannel> accept();
public abstract <A> void accept(A attachment, CompletionHandler<AsynchronousSocketChannel,? super A> handler);
```

* read 异步操作

```java
public abstract Future<Integer> read(ByteBuffer dst);
public final <A> void read(ByteBuffer dst,
                               A attachment,
                               CompletionHandler<Integer,? super A> handler)
```

* write 异步操作

```java
public abstract Future<Integer> write(ByteBuffer src);
public final <A> void write(ByteBuffer src,
                                A attachment,
                                CompletionHandler<Integer,? super A> handler)
```

通过 Future 实现异步客户端、服务器程序，尽管 I/O 相关方法调用是异步的，但是还得通过 Future 的 get 方法获取操作的结果，而 Future 的 get 调用是同步的，所以还是没有做到完全异步。而通过 CompletionHandler 获取 I/O 结果，所有 I/O 操作的执行结果都会通过 CompletionHandler 回调返回。

## 5. 基于 CompletionHandler 编写服务器

基于 CompletionHandler 编写服务器，关键是两步：

* 需要给每一个 I/O 操作传入一个 attachment 参数，这是用来记录用户上下文信息的。在示例代码中，我们抽象了一个类表示上下文信息。

```java
private static class AsyncIOOP {
        private int op_type;
        private ByteBuffer read_buffer;
        private AsynchronousSocketChannel client;
}
```

* 还需要传入一个 CompletionHandler 参数，这需要你自定义一个类并且实现 CompletionHandler 接口。由于 accept 操作和其他三个操作不同，所以我们定义了两个实现 CompletionHandler 接口的类。

```java
private static class AsyncIOOPCompletionHandler implements CompletionHandler<Integer, AsyncIOOP>
{
}

private static class AsyncAcceptCompletionHandler implements CompletionHandler<AsynchronousSocketChannel, syncIOOP>
{
}
```

每一个 I/O 操作完成，系统都会回调 CompletionHandler 的 completed 方法，你需要覆盖此方法，然后处理返回结果。

示例代码实现的是一个 Echo 逻辑，关键步骤如下：

* 服务器启动的时候，投递一个 accept 操作。
* 当收到 accept 操作完成，首先投递一个 accept 操作，准备接收新客户端请求；然后为刚接收的客户端投递一个 read 操作，准备接收数据。
* 当收到 read 操作完成，向客户端投递一个 write 操作，发送响应给客户端；然后再次投递一个 read 操作，准备接收新的消息。
* 当收到 write 操作完成，我们没有处理逻辑，因为这是一个简单的 Echo 功能。

完整代码如下：

```java
import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.ByteBuffer;
import java.nio.channels.AsynchronousServerSocketChannel;
import java.nio.channels.AsynchronousSocketChannel;
import java.nio.channels.CompletionHandler;

public class AsyncServerCompletionHandler {
    private static final int PORT =56002;
    private AsynchronousServerSocketChannel server = null;
    private static final int ASYNC_READ = 1;
    private static final int ASYNC_WRITE = 2;
    private static final int ASYNC_ACCEPT = 3;
    private static final int ASYNC_CONNECT = 4;

    private static class AsyncIOOP {
        private int op_type;
        private ByteBuffer read_buffer;
        private AsynchronousSocketChannel client;

        public int getOp_type() {
            return op_type;
        }

        public void setOp_type(int op_type) {
            this.op_type = op_type;
        }

        public ByteBuffer getRead_buffer() {
            return read_buffer;
        }

        public void setRead_buffer(ByteBuffer read_buffer) {
            this.read_buffer = read_buffer;
        }

        public AsynchronousSocketChannel getClient() {
            return client;
        }

        public void setClient(AsynchronousSocketChannel client) {
            this.client = client;
        }

        public AsyncIOOP(int op) {
            this(op, null, null);
        }
        public AsyncIOOP(int op, ByteBuffer b) {
            this(op, b, null);
        }
        public AsyncIOOP(int op, ByteBuffer b, AsynchronousSocketChannel ch) {
            this.op_type = op;
            this.read_buffer = b;
            this.client = ch;
        }
    }
    private static class AsyncIOOPCompletionHandler implements CompletionHandler<Integer, AsyncIOOP>
    {
        private AsyncServerCompletionHandler server;

        public AsyncIOOPCompletionHandler(AsyncServerCompletionHandler server){
            this.server = server;
        }
        @Override
        public void completed(Integer result, AsyncIOOP attachment) {
            if (attachment.op_type == ASYNC_READ) {
                server.async_write(attachment.getClient(), "Hello Client!");

                ByteBuffer inBuffer = attachment.getRead_buffer();
                System.out.println("Recv message from client:" + new String(inBuffer.array()).trim());

                server.async_read(attachment.getClient());
            } else if (attachment.op_type == ASYNC_WRITE) {

            }
        }

        @Override
        public void failed(Throwable exc, AsyncIOOP attachment) {
            try {
                attachment.getClient().close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    private static class AsyncAcceptCompletionHandler implements CompletionHandler<AsynchronousSocketChannel, AsyncIOOP>
    {
        private AsyncServerCompletionHandler server;

        public AsyncAcceptCompletionHandler(AsyncServerCompletionHandler server) {
            this.server = server;
        }

        @Override
        public void completed(AsynchronousSocketChannel result, AsyncIOOP attachment) {
            if (attachment.op_type == ASYNC_ACCEPT) {
                server.accept_new_client();

                if (result != null && result.isOpen()) {
                    server.async_read(result);
                }
            }
        }

        @Override
        public void failed(Throwable exc, AsyncIOOP attachment) {

        }
    }

    public void start() {
        try {
            server = AsynchronousServerSocketChannel.open();
            server.bind(new InetSocketAddress("127.0.0.1", PORT));
            accept_new_client();
        } catch (Exception e) {
            e.printStackTrace();
            stop();
        }
    }

    public void accept_new_client() {
        server.accept(new AsyncIOOP(ASYNC_ACCEPT), new AsyncAcceptCompletionHandler(this));
    }

    public void async_read(AsynchronousSocketChannel client){
        ByteBuffer inBuffer = ByteBuffer.allocate(128);
        AsyncIOOP ioop = new AsyncIOOP(ASYNC_READ, inBuffer, client);
        client.read(inBuffer, ioop, new AsyncIOOPCompletionHandler(this));
    }
    public void async_write(AsynchronousSocketChannel client, String message){
        ByteBuffer outBuffer = ByteBuffer.wrap(message.getBytes());
        AsyncIOOP ioop = new AsyncIOOP(ASYNC_WRITE, outBuffer, client);
        client.write(outBuffer, ioop, new AsyncIOOPCompletionHandler(this));
    }
    public void stop(){
        if (server != null){
            try {
                server.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
    public static void main(String[] args) {
        AsyncServerCompletionHandler server = new AsyncServerCompletionHandler();
        server.start();

        try {
            Thread.sleep(1000*1000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
```

## 6. 总结

本小节重点是介绍 Java NIO2 中引入的异步 Socket 的功能。异步 Socket 的核心是每一个 I/O 方法（connect、accept、read、write）的调用只是向系统投递一个事件，方法执行完会立即返回。如果要获取 I/O 执行的结果，可以通过 Future 或者 CompletionHandler 获取。Java 的这个机制非常类似 Windows IOCP（完成端口）的功能，如果有兴趣可以参考[专栏](1) IOCP 一节，或者 [IOCP 相关实现代码](2)。

## 7. 参考

[1]:[imooc_nwp] [https://www.imooc.com/read/80](https://www.imooc.com/read/80)

[2]:[iocp] [https://github.com/haska1025/imooc-sock-core-tech/tree/master/04-25_Windows_IOCP机制/iocp_server](https://github.com/haska1025/imooc-sock-core-tech/tree/master/04-25_Windows_IOCP%E6%9C%BA%E5%88%B6/iocp_server)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

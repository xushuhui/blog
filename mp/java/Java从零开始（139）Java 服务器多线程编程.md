---
title: Java从零开始（139）Java 服务器多线程编程
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java 服务器多线程编程


## 1. 前言

前面小节介绍的 Java TCP Socket 程序是单线程模型，也是阻塞式模型。我们调用 java.net.ServerSocket 的 accept 方法，此时线程会被阻塞，等待客户端连接。当有新客户端连接到服务器以后，accept 方法会返回一个 java.net.Socket 类型的对象，此对象代表了客户端和服务器完成了**三次握手**后，建立的新连接。 调用 java.net.Socket 的 recv 和 send 方法和客户端进行数据收发。由于我们采用的是阻塞式 Socket 编程，java.net.ServerSocket 的 accept 方法会阻塞线程，java.net.Socket 的 recv 和 send 方法也会阻塞线程。因此，如果采用此模型，在同一时刻，服务器只能和一个客户端通信。

要想服务器同时和多个客户端进行通信，要么采用**非阻塞式 Socket 编程**，通过 **I/O 多路复用机制** 实现此目的；要么采用多线程编程模型。当然，在**非阻塞式 Socket 编程**模型下，往往也采用多线程编程。因为目前的计算机都是多核处理器，采用多线程编码模型，可以充分利用 CPU 多核的优势，最大化 CPU 资源的利用。

本节主要介绍**阻塞式 Socket 编程**中常用的两种线程模型：

* 每线程模型
* 线程池模型

## 2. Java 多线程编程方法

由于本节会涉及到 Java 多线程编程，所以需要你能预先掌握 Java 多线程编程的方法。比如，线程的创建，线程的启动，线程之间的同步和线程之间的通信。

在 Java 平台下，创建线程的方法有两种：

* 第一，是创建一个用户自定义的线程类，然后继承 java.leng.Thread 类，同时要覆写它的 run 方法，调用它的 start 方法启动线程。例如：

```java
class MyThread extends Thread
{
    @Override
    public void run() {
        super.run();
    }
}

new MyThread().start();
```

* 第二，是创建一个任务类。

首先，实现 Runnable 接口，并且重写它的 run 方法。然后，创建 java.leng.Thread 类的对象，同时将 Runnable 的实例通过 java.lang.Thread 的构造方法传入。最后，调用 java.lang.Thread 的 start 方法启动线程。例如：

```java
class MyTask implements Runnable
{
    @Override
    public void run() {

    }
}
new Thread(new MyTask()).start();
```

## 3. 每线程模型

下图展示了**每线程模型**的结构。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyr8u5zj60q60j6gsh02)

从图中可以看出，每线程模型的程序结构如下：

* 创建一个监听线程，通常会采用 Java 主线程作为监听线程。
* 创建一个 java.net.ServerSocket 实例，调用它的 accept 方法等待客户端的连接。
* 当有新的客户端和服务器建立连接，accept 方法会返回，创建一个新的线程和客户端通信。此时监听线程返回，继续调用 accept 方法，等待新的客户端连接。
* 在新线程中调用 java.net.Socket 的 recv 和 send 方法和客户端进行数据收发。
* 当数据收发完成后，调用 java.net.Socket 的 close 方法关闭连接，同时线程退出。

下来，我们通过一个简单的示例程序演示一下**每线程模型**服务器的编写方法。示例程序的基本功能如下：

* 客户端每隔 1 秒向服务器发送一个消息。
* 服务器收到客户端的消息后，向客户端发送一个响应消息。
* 客户端发送完 10 个消息后，关闭 Socket 连接，程序退出。
* 服务器检测到客户端关闭连接后，同样关闭 Socket 连接，并且负责和客户端通信的线程也退出。

客户端代码：

```java
import java.io.*;
import java.net.InetSocketAddress;
import java.net.Socket;
import java.net.SocketAddress;

public class TCPClientMultiThread {
    // 服务器监听的端口号
    private static final int PORT = 56002;
    // 连接超时时间
    private static final int TIMEOUT = 15000;
    // 客户端执行次数
    private static final int TEST_TIMES = 10;

    public static void main(String[] args) {
        Socket client = null;
        try {
            // 测试次数
            int testCount = 0;
            // 调用无参构造方法
            client = new Socket();
            // 构造服务器地址结构
            SocketAddress serverAddr = new InetSocketAddress("192.168.0.101", PORT);
            // 连接服务器，超时时间是 15 毫秒
            client.connect(serverAddr, TIMEOUT);

            System.out.println("Client start:" + client.getLocalSocketAddress().toString());
            while (true) {
                // 向服务器发送数据
                DataOutputStream out = new DataOutputStream(
                        new BufferedOutputStream(client.getOutputStream()));
                String req = "Hello Server!";
                out.writeInt(req.getBytes().length);
                out.write(req.getBytes());
                // 不能忘记 flush 方法的调用
                out.flush();
                System.out.println("Send to server:" + req);

                // 接收服务器的数据
                DataInputStream in = new DataInputStream(
                        new BufferedInputStream(client.getInputStream()));

                int msgLen = in.readInt();
                byte[] inMessage = new byte[msgLen];
                in.read(inMessage);
                System.out.println("Recv from server:" + new String(inMessage));

                // 如果执行次数已经达到上限，结束测试。
                if (++testCount >= TEST_TIMES) {
                    break;
                }

                // 等待 1 秒然后再执行
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            if (client != null){
                try {
                    client.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}
```

服务器代码：

```java
import java.io.*;
import java.net.ServerSocket;
import java.net.Socket;

public class TCPServerPerThread implements Runnable{
    private static final int PORT =56002;

    private Socket sock = null;

    TCPServerPerThread(Socket sock){
        this.sock = sock;
    }

    @Override
    public void run() {
        // 读取客户端数据
        try {
            while (true){
                // 读取客户端数据
                DataInputStream in = new DataInputStream(
                        new BufferedInputStream(sock.getInputStream()));
                int msgLen = in.readInt();
                byte[] inMessage = new byte[msgLen];
                in.read(inMessage);
                System.out.println("Recv from client:" + new String(inMessage) + "length:" + msgLen);

                // 向客户端发送数据
                String rsp = "Hello Client!\n";
                DataOutputStream out = new DataOutputStream(
                        new BufferedOutputStream(sock.getOutputStream()));
                out.writeInt(rsp.getBytes().length);
                out.write(rsp.getBytes());
                out.flush();
                System.out.println("Send to client:" + rsp + " length:" + rsp.getBytes().length);
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            if (sock != null){
                try {
                    sock.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }

    public static void main(String[] args) {
        ServerSocket ss = null;
        try {
            // 创建一个服务器 Socket
            ss = new ServerSocket(PORT);
            while (true){
                // 监听新的连接请求
                Socket conn = ss.accept();
                System.out.println("Accept a new connection:"
                        + conn.getRemoteSocketAddress().toString());
                Thread t = new Thread(new TCPServerPerThread(conn));
                t.start();
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            if (ss != null){
                try {
                    ss.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}
```

客户端采用单线程模型。服务器采用**每线程模型**，我们采用实现 Runnable 接口的方式实现多线程逻辑。从示例代码可以看出，**每线程模型**的优点就是结构简单，相比单线程模型，也没有增加复杂度。缺点就是针对每个客户端都创建线程，当和客户端通信结束后，线程要退出。频繁的创建、销毁线程，对系统的资源消耗比较大，只能用在简单的业务场景下。

## 3. 线程池模型

线程池模型的结构如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyrn44ej60pm0jeqaf02)

从图中可以看出，线程池模型的程序结构如下：

* 创建一个监听线程，通常会采用 Java 主线程作为监听线程。
* 创建一个 java.net.ServerSocket 实例，调用它的 accept 方法等待客户端的连接。
* 服务器预先创建一组线程，叫做线程池。线程池中的线程，在服务运行过程中，一直运行，不会退出。
* 当有新的客户端和服务器建立连接，accept 方法会返回 java.net.Socket 对象，表示新的连接。服务器一般会创建一个处理 java.net.Socket 逻辑的任务，并且将此任务投递给线程池去处理。然后，监听线程返回，继续调用 accept 方法，等待新的客户端连接。
* 线程池调度空闲的线程去处理任务。
* 在新新任务中调用 java.net.Socket 的 recv 和 send 方法和客户端进行数据收发。
* 当数据收发完成后，调用 java.net.Socket 的 close 方法关闭连接，任务完成。
* 线程重新回归线程池，等待调度。

下来，我们同样通过示例代码演示一下线程池模型的编写方法。程序功能和每线程模型完全一致，所以我们只编写服务端程序，客户端程序采用每线程模型的客户端。

示例代码如下：

```java

import java.io.*;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class TCPServerThreadPool{
    // 服务监听端口号
    private static final int PORT =56002;
    // 开启线程数
    private static final int THREAD_NUMS = 20;
    private static ExecutorService pool = null;

    // 创建一个 socket Task 类，处理数据收发
    private static class SockTask implements Callable<Void> {
        private Socket sock = null;

        public SockTask(Socket sock){
            this.sock = sock;
        }
        @Override
        public Void call() throws Exception {
            try {
                while (true){
                    // 读取客户端数据
                    DataInputStream in = new DataInputStream(
                            new BufferedInputStream(sock.getInputStream()));
                    int msgLen = in.readInt();
                    byte[] inMessage = new byte[msgLen];
                    in.read(inMessage);
                    System.out.println("Recv from client:" + new String(inMessage) + "length:" + msgLen);

                    // 向客户端发送数据
                    String rsp = "Hello Client!\n";
                    DataOutputStream out = new DataOutputStream(
                            new BufferedOutputStream(sock.getOutputStream()));
                    out.writeInt(rsp.getBytes().length);
                    out.write(rsp.getBytes());
                    out.flush();
                    System.out.println("Send to client:" + rsp + " length:" + rsp.getBytes().length);
                }
            } catch (IOException e) {
                e.printStackTrace();
            } finally {
                if (sock != null){
                    try {
                        sock.close();
                    } catch (IOException e) {
                        e.printStackTrace();
                    }
                }
            }
            return null;
        }
    }

    public static void main(String[] args) {
        ServerSocket ss = null;
        try {
            pool = Executors.newFixedThreadPool(THREAD_NUMS);
            // 创建一个服务器 Socket
            ss = new ServerSocket(PORT);
            while (true){
                // 监听新的连接请求
                Socket conn = ss.accept();
                System.out.println("Accept a new connection:"
                        + conn.getRemoteSocketAddress().toString());
                pool.submit(new SockTask(conn));
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            if (ss != null){
                try {
                    ss.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}

```

## 4. 小结

本节主要介绍的是 Java 服务器编程中比较简单的两种线程模型，**每线程模型**和**线程池模型**。示例程序都采用了**阻塞式** Socket 编程。编写 Java 服务器程序，通常需要采用多线程模型。对于非常简单的业务场景，可以采用**每线程模型**。对于比较复杂的业务场景，通常需要采用**线程池模型**。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

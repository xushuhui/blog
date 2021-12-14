---
title: Java从零开始（77）Executor 应用示例
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Executor 应用示例


## 1. 前言

上一节我们学习了 Executor 的基本概念和核心 API，本节带领大家实现一个具体的应用案例。从实际应用中感受一下 Executor 框架的使用，以及此框架带来的便利。

本节先描述待实现的案例内容，接着做编码实现，然后总结使用过程中的注意事项。

## 2. 案例描述

我们可以通过手工创建线程做逻辑单元的执行，但是当存在大量的需要执行的逻辑单元也是这样处理，就会出现很多麻烦的事情，且效率非常低下。手工创建线程并做线程管理，需要我们实现很多与业务无关的控制代码，另外手工不停的创建线程并做线程销毁，会浪费很多系统资源。

我们在实际项目中，常常通过使用 java 提供好的非常好用的线程框架 Executor 进行任务执行操作。

有这样一个场景：需要对某个目录下的所有文件（成百上千）进行加密并用文件的 MD5 串修改文件名称。

在开始动手实现之前，我们先做一个简单的分析。在这个案例中，我们将 “对文件进行加密、生成 MD5 串、修改文件名称” 作为待执行任务的内容。所有文件形成的列表就是我们待处理的数据范围。为了校验整个处理过程是否有文件遗漏，我们最终需要核对处理结果。为了方便演示，下面编码中部分数据采用了模拟的方式生成。

## 3. 编码实现

```java
import java.util.ArrayList;
import java.util.List;
import java.util.Random;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

public class ExecutorTest {

    // 模拟待处理的文件列表
    private static int fileListSize = new Random().nextInt(6);
    private static String[] fileList = new String[fileListSize];
    static {
        for(int i=0; i<fileListSize; i++) {
            fileList[i] = "fileName" + i;
        }
    }

    // 主线程
    public static void main(String[] args) throws Exception {
        // 创建用于处理任务的线程池
        ExecutorService executorService = Executors.newFixedThreadPool(10);
        // 任务提交，每一个任务处理一个文件
        List<FileDealTask> tasks = new ArrayList<>();
        for(int i=0; i<fileListSize; i++) {
            tasks.add(new FileDealTask(0, fileListSize, fileList[i]));
        }
        // 等待异步处理结果返回
        List<Future<Integer>> results = executorService.invokeAll(tasks);
        // 获取任务执行结果
        Integer total = 0;
        for (Future<Integer> result : results) {
            total = total + result.get();
        }
        System.out.println("预备处理的文件个数" + fileListSize + "，总共处理的文件个数：" + total);
        // 关闭线程池
        executorService.shutdown();
    }
}
```

上面代码注释已经很清楚了，我们观察下面的代码，看看任务代码。

```java
import java.util.Random;
import java.util.concurrent.Callable;

public class FileDealTask implements Callable<Integer> {

    private String fileName;

    public FileDealTask(int first, int last, String fileName) {
        this.fileName = fileName;
    }

    @Override
    public Integer call() throws Exception {
        try {
            Thread.sleep(new Random().nextInt(2000));
            System.out.println(Thread.currentThread().getName() + "：文件" + fileName + "已处理完毕");
        } catch (Exception e) {
            return 0;
        }
        return 1;
    }
}
```

我们通过在 IDE 中运行上面这个示例，看看实际的运行结果。

【补充视频】

上面代码逻辑中有随机内容，每次运行结果会有差异，运行上面的代码，我们观察运行结果：

```java
pool-1-thread-2：文件fileName1已处理完毕
pool-1-thread-3：文件fileName2已处理完毕
pool-1-thread-1：文件fileName0已处理完毕
预备处理的文件个数3，总共处理的文件个数：3
```

和我们的预期一致。

## 4. 注意事项

1. Executors 是 Executor 框架体系中的一个独立的工具类，用于快速创建各类线程池，在实际应用中，如果需要对线程池的各类参数做更多的自定义，可以参考此类的实现。
2. 做好评估权衡，当需要处理的数据量不是特别大时，没有必要使用 Executor。其底层使用多线程的方式处理任务，涉及到线程上下文的切换，当数据量不大的时候使用串行会比使用多线程快。
3. 在使用时，如果主线程不关心子任务的执行结果，请使用 Runnable 接口封装任务的执行逻辑。

## 5. 视频演示

## 6. 小结

本节通过一个实际例子的编码实现，展示了 Executor 的具体用法。当然本节中的用法相对比较简单，更多的用法还需要大家进一步学习，希望大家多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

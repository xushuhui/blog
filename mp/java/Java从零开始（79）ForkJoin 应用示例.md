---
title: Java从零开始（79）ForkJoin 应用示例
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# ForkJoin 应用示例


## 1. 前言

上一节我们学习了 ForkJoin 的基本概念和核心 API，本节带领大家实现一个具体的应用案例。从实际应用中感受一下 ForkJoin 框架的使用，以及此框架带来的便利。

本节先描述待实现的案例内容，接着做编码实现，然后总结使用过程中的注意事项。

## 2. 案例描述

我们在实际项目中，常常会碰到大数据集合的处理，如大文件、大表、内存中的大数据集合。由于数据量大，我们常常会考虑采用多线程的思路提高处理效率。

当待处理的数据集的每一部分的数据处理逻辑基本一致，且可以很好拆分成小的数据集进行处理时，使用 ForkJoin 进行处理比好合适。

我们还是采用 Executor 应用示例中的场景：需要对某个目录下的所有文件（成百上千）进行加密并用文件的 MD5 串修改文件名称。

在开始动手实现之前，我们先做一个简单的分析。在这个案例中，我们将 “对文件进行加密、生成 MD5 串、修改文件名称” 作为待执行任务的内容。所有文件形成的列表就是我们待处理的数据范围。为了校验整个处理过程是否有文件遗漏，我们最终需要核对处理结果，所以我们用 FileForkJoinTask 类（继承 RecursiveTask）封装我们的任务逻辑。为了方便演示，下面编码中部分数据采用了模拟的方式生成。

## 3. 编码实现

```java
import java.util.Random;
import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.Future;

public class ForkJoinTest {

    // 模拟待处理的文件列表
    private static int fileListSize = new Random().nextInt(15);
    private static String[] fileList = new String[fileListSize];
    static {
        for(int i=0; i<fileListSize; i++) {
            fileList[i] = "fileName" + i;
        }
    }

    // 主线程
    public static void main(String[] args) throws Exception {
        // 创建用于处理任务的线程池
        // ForkJoinPool forkJoinPool = ForkJoinPool.commonPool(); 这种创建方式可最大化使用全局系统资源
        ForkJoinPool forkJoinPool = new ForkJoinPool();
        // 提交待处理的总任务
        Future result = forkJoinPool.submit(new FileDealTask(0, fileListSize, fileList));
        // 获取任务执行结果
        System.out.println("预备处理的文件个数" + fileListSize + "，总共处理的文件个数：" + result.get());
        // 关闭线程池
        forkJoinPool.shutdown();
    }
}
```

上面代码注释已经很清楚了，我们观察下面的代码，看看任务是怎么切分的，以及子任务的结果是怎么做汇总的。

```java
import lombok.SneakyThrows;
import java.util.Random;
import java.util.concurrent.RecursiveTask;

public class FileDealTask extends RecursiveTask<Integer> {

    private String[] fileList;
    // 当子任务划分到只需要处理最多10个文件时，停止分割任务
    private final int threshold = 2;
    private int first;
    private int last;

    public FileDealTask(int first, int last, String[] fileList) {
        this.fileList = fileList;
        this.first = first;
        this.last = last;
    }

    @SneakyThrows
    @Override
    protected Integer compute() {
        // 执行结果
        int result = 0;

        // 任务足够小则直接处理（对文件进行加密、生成MD5串、修改文件名称）
        if (last - first <= threshold) {
            for (int i = first; i < last; i++) {
                result = result + 1;
                Thread.sleep(new Random().nextInt(2000));
                System.out.println(Thread.currentThread().getName() + "：文件" + fileList[i] + "已处理完毕");
            }
            System.out.println(Thread.currentThread().getName() + "：总共处理的文件数 (" + first + "," + last + ")" + result);
        } else {
            // 拆分成小任务
            int middle = first + (last - first) / 2;
            // 创建两个子任务
            FileDealTask leftTask = new FileDealTask(first, middle, fileList);
            FileDealTask rightTask = new FileDealTask(middle, last, fileList);
            // 触发两个子任务开始执行
            invokeAll(leftTask, rightTask);
            // 等待两个子任务执行结果并返回
            result = leftTask.join() + rightTask.join();
            System.out.println(Thread.currentThread().getName() + "：当前任务继续拆分 "
                    + " (" + first + "," + middle + "), (" + (middle) + "," + last + ")");
        }
        return result;
    }
}
```

我们通过在 IDE 中运行上面这个示例，看看实际的运行结果。

【补充视频】

上面代码逻辑中有随机内容，每次运行结果会有差异，运行上面的代码，我们观察某次运行结果如下：

```java
ForkJoinPool-1-worker-2：文件fileName3已处理完毕
ForkJoinPool-1-worker-2：总共处理的文件数 (3,4)1
ForkJoinPool-1-worker-1：文件fileName0已处理完毕
ForkJoinPool-1-worker-1：总共处理的文件数 (0,1)1
ForkJoinPool-1-worker-0：文件fileName4已处理完毕
ForkJoinPool-1-worker-3：文件fileName1已处理完毕
ForkJoinPool-1-worker-3：文件fileName2已处理完毕
ForkJoinPool-1-worker-3：总共处理的文件数 (1,3)2
ForkJoinPool-1-worker-1：当前任务继续拆分  (0,1), (1,3)
ForkJoinPool-1-worker-0：文件fileName5已处理完毕
ForkJoinPool-1-worker-0：总共处理的文件数 (4,6)2
ForkJoinPool-1-worker-2：当前任务继续拆分  (3,4), (4,6)
ForkJoinPool-1-worker-1：当前任务继续拆分  (0,3), (3,6)
预备处理的文件个数6，总共处理的文件个数：6
```

首先做了 (0,3)~(3,6)，之后对 (0,3) 做了 (0,1), (1,3) 的拆分，对 (3,6) 做了 (3,4), (4,6) 的拆分。和我们的预期一致。

## 4. 注意事项

1. ForkJoinPool 不是为了替代 ExecutorService，其主要用于实现 “分而治之” 的算法，最适合处理计算密集型的任务。
2. 做好评估权衡，当需要处理的数据量不是特别大时，没有必要使用 ForkJoin。其底层使用多线程的方式处理任务，涉及到线程上下文的切换，当数据量不大的时候使用串行会比使用多线程快。
3. 执行子任务时候要注意，使用 invokeAll，不能分别对子任务调用 fork。

## 5. 视频演示

## 6. 小结

本节通过一个实际例子的编码实现，展示了 ForkJoinPool 的具体用法。当然本节中的用法相对比较简单，更多的用法还需要大家进一步学习，希望大家多思考勤练习，早日掌握之。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

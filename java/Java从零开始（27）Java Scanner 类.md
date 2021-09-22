# Java Scanner 类

一直以来，我们都使用`System.out.println()`方法向屏幕打印内容，那么如何接收输入的内容呢？本小节所学习的`Scanner`类就可以实现对输入内容的接收。在本小节，我们将学习`Scanner`类的定义，如何使用`Scanner`类以及其常用方法，在学完这些基础知识后，我们会在最后学习一个比较有趣的实例程序。

## 1. 定义

`Scanner`是一个简单的文本扫描器，可以解析**基础数据类型**和**字符串**。

它位于`java.util`包下，因此如果要使用此类，必须使用`import`语句导入：

```java
import java.util.Scanner;
```

## 2. Scanner 对象创建

想要使用`Scanner`类就要了解如何创建对象，我们可以使用如下代码创建一个扫描器对象：

```java
Scanner scanner = new Scanner(System.in);
```

构造方法的参数`System.in`表示允许用户从**系统中**读取内容。本小节，我们的示例代码中都将使用这个构造方法。

> **Tips**：`System.in`是一个`InputStream`类型，`Scanner`类还有很多接收其他类型的构造方法。这里不详细介绍。

## 3. 常用方法

### 3.1 next() 及其同伴方法

想要获取用户的输入，只有对象是不行的，还要配合它的实例方法。此时配合`Scanner`类中的`next()`方法及其**同伴方法**可以获取指定类型的输入。

#### 3.1.1 next() 方法

`next()`方法的返回值是字符串类型，可以使用此方法，将用户输入的内容扫描为字符串。我们来看一个示例，获取并打印用户输入的内容：

```java
import java.util.Scanner;

public class ScannerDemo1 {
    public static void main(String[] args) {
        // 创建扫描器对象
        Scanner scanner = new Scanner(System.in);
        System.out.println("请输入一段内容，输入回车结束：");
        // 可以将用户输入的内容扫描为字符串
        String str = scanner.next();
        // 打印输出
        System.out.println("您输入的内容为：" + str);
        // 关闭扫描器
        scanner.close();
    }
}
```

在代码中我们注意到，在代码块的最后调用了`close()`方法，这个方法用于关闭当前扫描器，就和电脑的开关机一样，使用电脑前要开机，而当用不到的时候最好关掉。

编译执行代码，屏幕将会提示：

```java
请输入一段内容，输入回车结束：
```

接下来我们按照提示输入内容，然后输入回车结束输入：

![](https://xushuhui.gitee.io/image/imooc/5ec25dc10ac162a215621358.jpg)

#### 3.1.2 同伴方法

那什么是同伴方法呢？这里的同伴方法指的是`Scanner`类中以`next`单词开头的方法。我们举例来看几个同伴方法及其作用：

* `nextLine()` ：返回输入回车之前的所有字符；
* `nextInt()` ：将输入内容扫描为`int`类型；
* `nextFloat()` ：将输入内容扫描为`float`类型。

这里的`nextLine()` 方法也可以获取字符串。我们来看一下它和`next()`方法的差异：

`next()`方法只有扫描到有效字符后才会结束输入；而`nextLine()`方法可以直接使用回车结束输入。

另外，`next()`方法会自动去掉空白（例如回车、空格等），也不能得到带有空格的字符串；`nextLine()`方法可以得到空白和带有空格的字符串。

我们再来看一个示例，获取用户输入的姓名、年龄和身高，并且打印输出：

```java
import java.util.Scanner;

public class ScannerDemo2 {
    public static void main(String[] args) {
        // 创建扫描器对象
        Scanner scanner = new Scanner(System.in);
        System.out.println("请输入您的姓名：");
        // 将第一行输入扫描为字符串
        String name = scanner.nextLine();

        System.out.println("请输入您的年龄：");
        // 将第二行输入扫描为int类型
        int age = scanner.nextInt();

        System.out.println("请输入您的身高：");
        // 将第三行输入扫描为float类型
        float height = scanner.nextFloat();

        // 打印扫描器所扫描的值
        System.out.println("您的姓名为：" + name);
        System.out.println("您的年龄为：" + age);
        System.out.println("您的身高为：" + height);
        // 关闭扫描器
        scanner.close();
    }
}
```

编译执行代码，按照提示输入对应内容，直到程序完整运行：

```java
请输入您的姓名：
三井 寿
请输入您的年龄：
19
请输入您的身高：
183
您的姓名为：三井 寿
您的年龄为：19
您的身高为：183
```

> **Tips**：上面代码中，如果使用`next()`方法代替`nextLine()`方法来获取姓名字符串，是无法得到我们输入的“三井 寿”这个字符串的，这是因为`next()`方法不能获取带有空格的字符串。

要特别注意的是：Scanner 类读到的内容，只与输入顺序有关，和终端上显示的顺序无关，因此类似于下面的这种输入，是读不到空格的，执行代码的流程如下：

![](https://xushuhui.gitee.io/image/imooc/5ec25fbb0ad368f915621358.jpg)

### 3.2 hasNext() 及其同伴方法

`hasNext()`方法的返回值是一个布尔类型，如果输入中包含数据的输入，则返回`true`。否则返回`false`。通常用来做输入内容的验证。

它的同伴方法是以`hasNext`单词开头的方法，诸如`hasNextLine()`、`hasNextInt()`等方法。例如，上面的代码中，我们可以对应加入`hasNext`同伴方法结合条件判断语句，来提升代码的稳定性：

```java
int age;
if (scanner.hasNextInt()) {
    age = scanner.nextInt();
} else {
    System.out.println("不是int类型");
}

float height;
if (scanner.hasNextFloat()) {
    height = scanner.nextFloat();
} else {
    System.out.println("不是float类型");
}
```

## 4. 实例

前面我们已经对`Scanner`类的基本用法有了一定的了解，下面我们来实现一个示例程序，这个程序用于估算一个人的体脂率，这里事先给出体脂率估算公式：

```java
参数a = 腰围（cm）×0.74
参数b = 体重（kg）× 0.082 + 44.74
脂肪重量（kg）= a － b
体脂率 =（脂肪重量 ÷ 体重）× 100%。
```

从公式中我们可以看出，想要得到最终的体脂率，参数 a（腰围）和参数 b（体重）是需要用户手动输入的，公式部分只需要使用算数运算符实现即可。下面是程序代码：

```java
import java.util.Scanner;

public class GetBodyFat {
    public static void main(String[] args) {
        // 初始化腰围
        float waistline = 0f;
        // 初始化体重
        float weight = 0f;
        // 声明浮点型参数a，b，bodyFatWeight（脂肪重量）
        float a, b, bodyFatWeight;
        Scanner scanner = new Scanner(System.in);
        System.out.println("请输入您的腰围（cm）：");
        if (scanner.hasNextFloat()) {
            waistline = scanner.nextFloat();
        }
        System.out.println("请输入您的体重（kg）：");
        if (scanner.hasNextFloat()) {
            weight = scanner.nextFloat();
        }
        // 计算参数a  公式：参数a = 腰围（cm）× 0.74
        a = waistline * 0.74f;
        // 计算参数b  公式：参数b = 体重（kg）× 0.082 + 44.74
        b = weight * 0.082f + 44.74f;
        // 计算脂肪重量
        bodyFatWeight = a - b;
        // 计算体脂率 =（脂肪重量 ÷ 体重）×100%。
        float result = bodyFatWeight / weight * 100;
        System.out.println("您的体脂率为" + result + "%");
    }
}
```

编译运行代码，按照提示输入，将估算出你的体脂含量：

```java
请输入您的腰围（cm）：
70
请输入您的体重（kg）：
50
您的体脂率为5.919998%
```

执行代码的流程如下：

![](https://xushuhui.gitee.io/image/imooc/5ec260de0a39a5d415621358.jpg)

## 5. 小结

本小节我们学习了 Java 的 `Scanner`类，它是位于`java.util`包下的一个工具类，我们知道了它是一个简单的文本扫描器，可以解析基础数据类型和字符串。我们也学会了如何使用`Scanner`类来获取用户的输入，`next()`方法和`nextLine()`方法都可以扫描用户输入的字符串，要注意这两个方法的区别。我们也在最后给出了一个计算体脂率的示例代码，学习了`Scanner`类，你就可以实现比较有意思的一些小程序了。如果你想了解更多有关`Scanner`类的接口，也可翻阅官方文档。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

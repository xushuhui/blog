---
title: Java从零开始（13）Java 循环语句
zhihu-url: https://zhuanlan.zhihu.com/p/407011026
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java 循环语句


循环结构能够让程序员以最少的精力去完成大量重复的工作，它可以让计算机根据条件做循环计算，当条件成立时继续循环，当条件不成立时结束循环。

依据前面所学的知识，如果我们想在屏幕上依次打印 1 到 5，可以编写以下程序：

```java
// PrintNums.java
class PrintNums {
    public static void main(String[] args) {
      	// 依次打印1到5
        System.out.println("1");
        System.out.println("2");
        System.out.println("3");
        System.out.println("4");
        System.out.println("5");
    }
}
```

运行结果：

```java
1
2
3
4
5
```

但是这种编写代码的方案存在一些弊端：

* **不灵活**：如果我们想要更改需求，打印从 6 到 10，那么就不得不逐行更改；
* **难于维护**：如果有大量代码，更容易产生`bug`；
* **伸缩性差**：依次打印从 1 到 5 貌似很简单，如果需求变为打印从 1 到 100 呢？这意味着需要编写大量的代码。

使用循环结构，就可以解决上述弊端。下面我们打开代码编辑器，新建一个`LoopPrintNums.java`，复制如下代码：

```java
// LoopPrintNums.java
class LoopPrintNums {
    public static void main(String[] args) {
        for(int i = 1; i <= 10; i ++) {
            System.out.print(i + "\t");
        }
    }
}
```

运行结果：

```java
1	2	3	4	5	6	7	8	9	10
```

聪明的你可能发现，如果将`i <= 10`改为 `i <= 100`，屏幕将依次从 1 打印 100。

上述代码中，我们看到不需要再去编写多条打印语句，同样得到了我们想要的结果，代码量大大减少。那么上述代码具体有什么含义呢，为什么这样写就实现了多个数字的打印呢？在本小节中，我们就会进行详细介绍。

## 1. while 循环

### 1.1 语法

`while`循环是最简单的循环形式。我们先来看看它的基本用法：

```java
while (条件) {
    // 循环体，条件成立时执行
    ...
}
// 循环完成后执行
```

`while`循环在每次循环开始前，首先会判断条件是否成立，如果计算结果为`true`， 就会执行循环体内部语句。如果计算结果为`false`，会跳出循环，执行后续代码。

### 1.2 实例

对于打印从 1 到 5 的例子，可以使用`while`循环可以这样写：

```java
public class WhileLoop1 {
    public static void main(String args[]) {
        int i = 1;
        while (i <= 5) {
            System.out.println(i); // 打印i
            i ++; // i自增1
        }
    }
}
```

你也许注意到在循环体内部，除了打印变量`i`的值，还有一个语句：`i ++`，这个语句的作用是让变量`i`自增 1。如果没有这个语句，`i`的值永远不会改变，所以循环条件永远成立，程序将陷入死循环，直到将资源耗尽才会停止，所以要杜绝编写死循环代码。

> **Tips**：如果循环的条件永远不会变成 `false`（比如使用字面量`true`作为条件），则称这种循环为无限循环或死循环。循环体内的代码会无限地重复执行，直到中断。

再看一个例子，得到 1 到 100 的累加和：

```java
public class WhileLoop2 {
    public static void main(String[] args) {
        int sum = 0; // 累加和
        int i = 1;
        while (i <= 100) {
            sum = sum + i; // 使sum和i相加，并将值再次赋值给sum
            i ++; // i自增1
        }
        System.out.println("1到100的累加和为：" + sum);
    }
}
```

运行结果：

```java
1到100的累加和为：5050
```

我们来对程序进行解析：

1. 定义两个变量`sum`和`i`分别用于存放**累加和**以及**要累加的数字**。
2. 循环开始前判断条件 `i <= 100`是否成立，此时 `i = 1`，成立，执行循环体内代码。
3. 循环体内代码第一次执行，`sum = 1`，`i` 自增 1 变为 2。
4. 再次判断循环条件，成立就继续执行循环体，直到 `i = 100`，自增 1 变为 101，此时条件不成立。执行循环后面的代码。
5. 执行打印语句，打印最终的 `sum`值，程序结束。

## 2. do while 循环

### 2.1 语法

`do while`循环的功能与`while`循环类似，不同点在于：`while`循环是先判断条件，再执行循环体；而`do while` 循环则是先执行循环体，再判断条件，如果条件成立继续执行循环，条件不成立则终止循环。`do while`循环的语法如下：

```java
do {
    // 循环体
} while (条件);
```

无论条件成立与否，`do while`循环都至少执行一次。而`while`循环可能一次都不会执行。

### 2.2 实例

我们现在使用 `do while`循环实现 1 到 100 的累加求和，代码如下：

```java
public class DoWhileLoop {
    public static void main(String[] args) {
        int sum = 0; // 累加和
        int i = 1;
        do {
            sum = sum + i; // 使sum和i相加，并将值再次赋值给sum
            i ++; // i自增1
        } while (i <= 100);
        System.out.println("1到100的累加和为：" + sum);
    }
}
```

运行结果：

```java
1到100的累加和为：5050
```

## 3. for 循环

### 3.1 语法

下面是`for`循环的语法结构：

```java
for (变量声明; 条件; 迭代语句) {
    // 循环体
}
```

`for`循环是一种特殊的`while`循环，也是被使用最广泛的循环。它使用计数器来实现循环。在关键字 `for`后面的括号中，会有三个语句，第一个语句是**变量声明语句**，允许声明一个或多个**整型变量**；第二个语句是条件，条件的检测方式与`while`循环相同（每次循环开始前判断条件成立与否）；第三个语句是迭代语句，虽然可以是任何语句，但该语句通常用于递增或递减变量。

### 3.2 实例

我们使用`for`循环连续打印 5 行”真棒“的字样，代码如下：

```java
public class ForLoop1 {
    public static void main(String[] args) {
        for (int i = 1; i <= 5; i ++) {
            System.out.println("真棒");
        }
    }
}
```

运行结果：

```java
真棒
真棒
真棒
真棒
真棒
```

我们现在使用`for`来实现 1 到 100 的累加求和，代码如下：

```java
public class ForLoop2 {
    public static void main(String[] args) {
        int sum = 0; // 累加和
        for (int i = 1; i <= 100; i ++) {
            sum += i; // 加和赋值运算
        }
        System.out.println("1到100的累加和为：" + sum);
    }
}
```

运行结果：

```java
1到100的累加和为：5050
```

程序解析：

在执行循环体前，会先执行括号中的变量声明语句`int i = 1`，它定义了一个计数器变量`i`并且赋值为 1（注意这个语句只执行一次）。然后会执行条件判断语句`i <= 100`，显然条件成立，执行循环体，循环体执行完成后会才会执行括号中的第三个语句`i ++`，再次判断条件是否成立，直到当`i`的值为 100 时，最后一次执行循环体，这个时候`i`再次自增`1`为`101`，条件不成立，结束循环。

实际使用中，通常使用`for`循环来遍历一个**数组**，我们将在 Java 数组小节举例介绍。

`for`循环中第一个语句，变量声明语句可以一次声明多个整型变量，实例代码如下：

```java
public class ForLoop2 {
    public static void main(String[] args) {
        for (int i = 1, j = 10; i <= 10; i++, j--) {
            System.out.print(i + "\t");
            System.out.print(j + "\t");
        }
    }
}
```

运行结果：

```java
1	10	2	9	3	8	4	7	5	6	6	5	7	4	8	3	9	2	10	1
```

另外，`for`循环还可以缺省变量声明语句、循环条件和迭代语句，实例：

```java
// 缺省变量声明语句
for (int i = 0; ; i ++) {
    // 语句
}
```

```java
// 缺省循环条件语句和迭代语句
for (int i = 0; ; ) {
    // 语句
}
```

```java
// 缺省所有语句
for (;;) {
    // 语句
}
```

但请注意，通常情况下不建议这样写。

## 4. for each 循环

由于学习 `for each` 循环需要一些数组和集合的前置知识，我们将会在 [Java 数组](http://www.imooc.com/wiki/javalesson/arrays.html)小节举例讲解。

## 5. break 和 continue 关键字

我们可以使用`break`和`continue`语句来重定向程序执行流。

### 5.1 break

`break`关键字可以出现在一个循环中，用以**跳出当前循环**。我们来看一个实例：

```java
public class BreakKeywords1 {
    public static void main(String[] args) {
        for (int i = 1; i <= 10; i ++) {
            if(i == 3) { // 条件语句
                System.out.println("当i等于3时，终止循环！");
                break;
            }
            System.out.println(i);
        }
    }
}
```

运行结果：

```java
1
2
当i等于3时，终止循环！
```

上面是一段循环打印数字的代码，当循环到达第 3 次迭代时，执行条件分支中的语句，将终止循环。注意：`break`语句通常配合`if`语句使用。

对于多层循环的情况，`break`语句只能终止它自己所在的那一层循环，并不会影响到外层循环。请阅读如下代码：

```java
public class BreakKeywords2 {
    public static void main(String[] args) {
        for (int i = 1; i <= 5; i ++) {
            // 打印i的值
            System.out.println(i);
            for (int j = 1; j<=5; j ++) {
                if(j == 3) { // 条件语句
                    System.out.println("当j等于3时，终止循环！");
                    break;
                }
                // 打印j的值
                System.out.println(j);
            }
        }
    }
}
```

运行结果：

```java
1
1
2
当j等于3时，终止循环！
2
1
2
当j等于3时，终止循环！
3
1
2
当j等于3时，终止循环！
4
1
2
当j等于3时，终止循环！
5
1
2
当j等于3时，终止循环！
```

上面的代码有两个`for`循环嵌套，`break`语句位于内层循环，所以当表达式`j == 3`成立时，会跳出内层循环，进而继续执行外层循环。

### 5.2 continue

`continue`关键字可以跳过**当次循环**，继续执行下一次循环，我们来看一个实例：

```java
public class ContinueKeywords {
    public static void main(String[] args) {
        for (int i = 1; i <= 5; i ++) {
            if(i == 3) {
                System.out.println("当i等于3时，跳过~");
                continue;
            }
            System.out.println(i);
        }
    }
}
```

运行结果：

```java
1
2
当i等于3时，跳过~
4
5
```

在多层嵌套的循环中，`continue`语句只能跳过它自己所在位置的循环。

## 6. 小结

本节中，我们依次介绍了`Java`中的四种循环结构：`while`循环、`do while`循环、`for`循环以及`for each`循环。并且介绍了`break`和`continue`关键字在循环中的使用。

其中，`while`循环的循环体可能一次都不执行，而`do while`循环的循环体至少执行一次。

`for` 循环通过**计数器**实现复杂循环。`for each`循环可以直接遍历数组的每个元素。

在使用循环结构编写业务代码时，为了避免死循环，一定要确保循环条件能够被影响。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

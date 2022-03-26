---
title: Java 从零开始（16）Java 方法
zhihu-url: https://zhuanlan.zhihu.com/p/408075427
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java 方法

本小节我们将学习**什么是方法**、**如何自定义方法**，并按照**分类介绍**每种方法的特点，对于有参数的**方法传值**，会讲到基本数据类型作为方法参数和引用数据类型作为方法参数的区别。也会学习可变参数方法的定义语法和使用场景，**方法重载的使用和意义**也是本节的重点学习内容。

## 1. 什么是方法

在前面我们已经了解过方法的概念，Java 程序的入口`main()`就是一个方法。`System.out.println();`语句中`println()`也是一个方法。

如果你使用`IntelliJ IDEA`查看代码，可以使用`Ctrl` + 鼠标左键单击代码中的 `println()`方法，查看 JDK 源码中对于这个方法的定义：

![](https://xushuhui.gitee.io/image/imooc/5e973932090666b508300461.jpg)

上面的截图就是我们经常调用的老朋友`println()`方法的代码实现，我们可以将方法理解为一个**命名语句块**，通过其名称 + 括号运算符`()`可以调用。我们可以将需要重复编写的代码，封装到一个方法中。提高代码的复用性。

## 2. 如何定义方法

在 Java 中，定义一个方法的语法为：

```java
访问修饰符 返回类型 方法名（参数列表） {
    若干语句；
    return 方法返回值；
}
```

* **访问修饰符有 4 种情况**：`public`、`private`、`protected`，也可以省略（`default`)。由于涉及到后面的面向对象相关知识，本节统一使用`public`修饰方法；

* **返回类型**：可以是任何的数据类型或 `void`，如果方法没有返回值，返回类型设置为`void`；

* **方法名**：方法名的命名规范和变量相同；

* **参数列表**：参数是变量的一种类型，参数变量的作用域在方法内部；

* **方法体**：方法内部的一些语句。当方法返回值为 `void`时，可以省略`return`语句。

## 3. 分类

根据方法是否带有参数、是否有返回值，可以分为 4 类：

* 无参无返回值方法
* 无参带返回值方法
* 带参数无返回值方法
* 带参数带返回值方法

### 3.1 无参无返回值方法

无参无返回值方法，即参数列表为空，返回值为`void`的方法。例如，我们经常需要输出一段内容。可以将输出语句封装到方法中：

```java
class MethodDemo1 {
    // 定义无参无返回值方法
    public void printText() {
        // 三条输出语句
        System.out.println("你好新同学");
        System.out.println("Java 是就业前景最好的语言");
        System.out.println("学 Java 就来");
    }
}
```

上面代码，在类`MethodDemo1`中定义了一个方法`printText`，其中有一些打印语句。那么如何调用这个方法，让其执行呢？

对于**非静态方法**（即非`static`关键字修饰的方法），首先要使用`new`关键字实例化类，生成一个对象，再通过`对象名。方法名 ()`的方式去调动方法。

如下是 MethodDemo1 类方法调用的实例：

```java
class MethodDemo1 {
    // 定义无参无返回值方法
    public void printText() {
        // 三条输出语句
        System.out.println("你好新同学");
        System.out.println("Java 是就业前景最好的语言");
        System.out.println("学 Java 就来");
    }

    public static void main(String[] args) {
	    // 实例化 MethodDemo1 类
	    MethodDemo1 methodDemo1 = new MethodDemo1();
	    // 调用 printText 方法
	    methodDemo1.printText();
	}
}
```

运行结果：

```java
你好新同学
Java 是就业前景最好的语言
学 Java 就来
```

> **Tips**：由于我们还没有学习面向对象，此处对于对象的实例化，先记住这种固定写法即可。本节的学习重点在于如何定义方法和调用方法。

运行结果：

```java
你好新同学
Java 是就业前景最好的语言
学 Java 就来
```

### 3.2 无参带返回值方法

无参带返回值方法，即参数列表为空，返回值不为`void`的方法，方法体内部要使用`return`语句返回其声明的返回类型。

我们来自定义一个返回`Hello World!`字符串的方法：

```java
class MethodDemo2 {
  	// 定义无参带返回值方法
  	public String getHelloStr() {
        // 初始化变量 str
      	String str = "Hello World";
      	// 返回 str
      	return str;
    }

  	public static void main(String[] args) {
      	// 实例化 MethodDemo2
      	MethodDemo2 methodDemo2 = new MethodDemo2();
      	// 调用 getStr 方法，并且使用 result 变量接收返回值
      	String result = methodDemo2.getHelloStr();
      	// 打印返回值
      	System.out.println(result);
    }
}
```

运行结果：

```java
Hello World
```

### 3.3 带参数无返回值方法

带参数无返回值方法，即带有参数列表，返回值为`void`的方法。

参数列表可能有一个或多个参数，多个参数使用逗号 (`,`) 隔开。定义方式为：`DataType 参数名 1, DataType 参数名 2`，DateType 为参数的类型。

下面是一个根据长和宽求长方形面积的方法实例：

```java
class MethodDemo3 {
  	// 定义带参数无返回值方法
  	public void printArea(float a, float b) {  // 此处的 a,b 为参数变量
		float result = a * b;
      	System.out.println( "长方形面积为：" + result);
    }

  	public static void main(String[] args) {
      	// 实例化 MethodDemo3
      	MethodDemo3 methodDemo3 = new MethodDemo3();
      	// 初始化两个变量
      	float width = 12.3f;
      	float height = 20f;
      	// 调用 printArea 方法，并将 width、height 变量作为参数传入
      	methodDemo3.printArea(width, height);
      	// 也可不提前初始化变量，直接传入浮点型字面量作为参数。
      	methodDemo3.printArea(10.2f, 2.5f);
    }
}
```

运行结果：

```java
长方形面积为：246.0
长方形面积为：25.5
```

在调用带参数方法时，有两种传参方式：

1. 先初始化变量，再将变量传递给方法：

```java
float a = 12.3f;
float b = 20f;
methodDemo3.printArea(a, b);
```

1. 直接传入对应类型字面量：

```java
methodDemo3.printArea(10.2f, 2.5f);
```

当调用方把参数传递给方法时，调用时传递的值会按**参数位置一一绑定**，这也比较符合我们正常的思维逻辑。因此，对于有多个参数的方法，一定要注意每个参数的位置。

### 3.4 带参数带返回值方法

带参带返回值方法，即带有参数列表，返回值不为`void`的方法。

下面我们来看一个实例，求从 1 到 n 的平方和：

```java
class MethodDemo4 {
    // 定义带参数带返回值方法
    public int getSumOfSquares(int n) {
        int sum = 0;
        for (int i = 1; i <=n; i++) {
          	// 求 i 的平方，并加和赋值给 sum
            sum += i * i;
        }
      	// 返回计算结果
        return sum;
    }

    public static void main(String[] args) {
        // 实例化 MethodDemo4
        MethodDemo4 methodDemo4 = new MethodDemo4();
      	// 调用对象下 getSumOfSquares 方法，并用 sumOfSquares 变量接收返回结果
        int sumOfSquares = methodDemo4.getSumOfSquares(5);
        System.out.println("1 到 5 的累加平方和为：" + sumOfSquares);
    }
}
```

运行结果：

```java
1 到 5 的累加平方和为：55
```

我们也可以不定义变量接收返回结果，直接打印方法调用语句：

```java
System.out.println（"1 到 5 的累加平方和为：" + methodDemo4.getSumOfSquares(5)）;
```

上面有参方法的例子，参数都是简单的基本数据类型，当参数为**数组**时，我们再来看一个实例：

```java
class MethodDemo5 {

    /**
     * 查找数组中元素是否包含指定元素
     * @param arr 待查找数组
     * @param element 所查找的元素
     * @return boolean 查找结果，true 包含；false 不包含
     */
    public boolean contains(int[] arr, int element) {
      	// 定义返回结果，默认为 false，即未找到 element 元素
        boolean result = false;
        for (int value : arr) {
            if (value == element) {
              	// 找到了 element 元素，将变量更新为 true，并跳出循环
                result = true;
                break;
            }
        }
        // 返回计算结果
        return result;
    }

    public static void main(String[] args) {
        // 实例化 MethodDemo5
        MethodDemo5 methodDemo5 = new MethodDemo5();
        // 初始化一个数组
        int[] intArray = {1, 3, 5, 6, 7, 9};
        // 初始化待查找元素变量
        int element = 10;
        // 调用 contains 方法，并接收返回结果
        boolean contains = methodDemo5.contains(intArray, element);
        // 根据返回结果，输出内容
        if (contains) {
            System.out.println("数组 intArray 中包含元素" + element);
        } else {
            System.out.println("数组 intArray 中不包含元素" + element);
        }
    }
}
```

运行结果：

```java
数组 intArray 中不包含元素 10
```

代码解析：

上述代码定义了一个查找数组中是否包含指定元素的方法，它将**数组**和**指定元素**作为方法参数，并且返回了一个布尔型的结果，方法体内部首先声明了一个待返回的布尔变量，默认为`false`，然后对数组参数进行迭代，只要找到与元素参数相同的元素，即更新返回变量为`true`，跳出循环；如果数组迭代完成后依旧未找到指定的元素，那么待返回变量依然是`false`，最后返回这个变量。在方法调用时，以一个布尔变量来接收结果，后续代码根据这个布尔变量做了逻辑判断，以执行不同的语句块。

## 4. 方法传值

调用方和方法之间有参数的传递时，要注意方法传值问题。

### 4.1 基本类型的传值

基本类型参数的传递，是调用方值的复制。双方各自的后续修改，互不影响。简单来讲，方法内部对参数变量的任何操作，都不会影响到方法外部的变量。

我们看一个实例：

```java
class Car {
    public void speedUp(int speed) {
        System.out.println("小汽车加速前，速度为：" + speed);
        speed ++;
        System.out.println("小汽车加速后，速度为：" + speed);
    }

    public static void main(String[] args) {
        // 定义小汽车初始速度变量
        int speed = 10;
        // 实例化 Car 类，创建一个 car 对象
        Car car = new Car();
        // 调用 car 对象下的 speed 方法
        car.speedUp(speed);
        // 打印调用方法后速度参数的值
        System.out.println("调用 speedUp 方法后，调用方的 speed 参数为：" + speed);
    }
}
```

运行结果：

```java
小汽车加速前，速度为：10
小汽车加速后，速度为：11
调用 speedUp 方法后，调用方的 speed 参数为：10
```

我们根据输出的结果验证了以上给出的结论。整型参数`speed`是一个基本数据类型，在`speedUp`方法内部进行了自增，但主方法中的`speed`变量的值并不会受到其影响。

### 4.2 引用类型的传值

引用类型参数的传递，调用方的变量，和接收方的参数变量，地址指向的是同一个对象。双方任意一方对这个对象的修改，都会影响对方。

我们来看一个实例：

```java
class NBATeam {
	// 替换第一个球员方法
    public void replaceFirstPlayer(String[] players, String playerName) {
        // 替换第一个球员
        System.out.println("将第一个球员替换：");
        players[0] = playerName;
    }

    public static void main(String[] args) {
        String[] players = {"詹姆斯", "科比", "杜兰特", "乔丹"};

        System.out.println("球队中现有球员：");
        for (String player : players) {
            System.out.print(player + "\t");
        }
        System.out.println();

        // 创建 team 对象并调用其替换球员方法
        NBATeam team = new NBATeam();
        team.replaceFirstPlayer(players, "皮蓬");

        System.out.println("替换后球员：");
        for (String player : players) {
            System.out.print(player + "\t");
        }
    }
}
```

运行结果：

```java
球队中现有球员：
詹姆斯  科比    杜兰特  乔丹
将第一个球员替换：
替换后球员：
皮蓬    科比    杜兰特  乔丹
```

由运行结果可见：方法`replaceFirstPlayer`的参数`players`是数组类型，也就是引用类型。当我们在方法内部对`players`进行操作，其方法外部（即主方法中）的`players`变量也受到了影响。不只是数组，对象也是引用类型，其参数传递同样遵循以上结论。我们在日常的编码中，请注意区分基本类型和引用类型的方法传值区别。

## 5. 可变参数

null 我们可以使用可变参数列表声明方法的参数。可变参数列表的语法：

```java
参数类型。.. 参数名
```

我们可以结合可变参数来实现一个求和方法`sum`，请阅读下面实例：

```java
class VariableParameter {
    public void sum(int... n) {
        int sum = 0;
      	// 可以对可变参数进行迭代
        for (int i: n) {
            sum = sum + i;
        }
        System.out.println("sum=" + sum);
    }

    public static void main(String[] args) {
        // 创建对象
        VariableParameter variableParameter = new VariableParameter();
        // 调用方法，传递一个参数
        variableParameter.sum(1);
        // 调用方法，传递两个参数
        variableParameter.sum(2, 3);
        // 调用方法，传递三个参数
        variableParameter.sum(5, 6, 7);
    }
}
```

运行结果：

```java
sum=1
sum=5
sum=18
```

上述实例中，在主方法中给`sum`方法传参时，可选择一个或多个参数传递。方法体内对可变参数`n`进行迭代，也可以将可变参数`n`改为数组：`int[] n`，得到的效果是相同的。

当方法的参数列表有两个或两个以上参数时，可变参数一定要放在最后，请查看如下实例：

```java
public void search(int element, int... elements) {
    boolean existed  = false;
    for (int e: elements) {
        if (e == element) {
            existed = true;
            break;
        }
    }
    if (existed) {
        System.out.println("找到元素：" + element);
    } else {
        System.out.println("未找到元素：" + element);
    }
}
```

另外，也可以将数组传递给可变参数列表，例如，在主方法中调用上面的`search`方法：

```java
public class VariableParameter1 {
    public void search(int element, int... elements) {
        boolean existed  = false;
        for (int e: elements) {
            if (e == element) {
                existed = true;
                break;
            }
        }
        if (existed) {
            System.out.println("找到元素：" + element);
        } else {
            System.out.println("未找到元素：" + element);
        }
    }

    public static void main(String[] args) {
        // 创建对象
        VariableParameter1 obj = new VariableParameter1();
        // 调用方法
        obj.search(2, 1,2,3,4);
        // 定义数组参数
        int[] arr = {1,2,3,4};
        // 将数组传递给可变参数列表
        obj.search(2, arr);
    }
}
```

运行结果：

```java
找到元素：2
```

上述两种传参方式都是合法的。

## 6. 方法重载

方法重载是指在一个类中定义多个**同名**的方法，但要求每个方法具有不同的参数的类型或参数的个数。调用重载方法时，Java 编译器能通过检查调用的方法的**参数类型**和个数选择一个恰当的方法。方法重载通常用于创建完成一组任务相似但参数的类型或参数的个数或参数的顺序不同的方法。

### 6.1 自定义方法的重载

例如，在`Student`类中，有多个`study`方法：

```java
public class Student {
    public void study() {
        System.out.println("同学真好学！");
    }

    public void study(String name) {
        System.out.println(name + "同学真好学！");
    }

    public void study(String name, int age) {
        System.out.println(name + "同学真好学！" + "他今年" + age + "岁了");
    }

    public static void main(String[] args) {
        // 实例化学生对象
        Student student = new Student();
        // 调用无参数方法
        student.study();
        // 调用单参数方法
        student.study("Colorful");
        // 调用双参数方法
        student.study("小慕", 20);
    }
}
```

运行结果：

```java
同学真好学！
Colorful 同学真好学！
小慕同学真好学！他今年 20 岁了
```

代码中的三个`study`都是重载方法。通常来说，方法重载的返回值类型都是相同的。

如果我们在`Student`类中再增加一个方法：

```java
public String study() {
    return "学习 Java 语言";
}
```

注意，上述的方法不是重载方法，因为我们已经在`Student`类中定义了无参方法`study`。

判断一个方法是否是重载方法的原则：方法名相同，参数类型或参数个数不同。

### 6.2 JDK 中的方法重载

`Java`语言本身的类也定义了很多方法重载的例子，例如`String`类的`substring`方法，用于字符串截取：

``java
public String substring(int beginIndex. int endIndex); // 截取并返回从 beginIndex 位置到 endIndex-1 位置的字符串
```

如下为实际应用的实例：

```java
String hello = "Hello, Imooc";
String substring1 = hello.substring(7);
String substring2 = hello.substring(0, 5);
System.out.println(substring1);
System.out.println(substring2);
```

运行结果：

```java
Imooc
Hello
```

## 7. 小结

本小节我们学习了 Java 方法，内容较多，我们一起来回顾一下所学的内容：

理解方法是一个**命名语句块**很重要。方法提高了代码的**复用性**，优质的方法提供者（如 JDK 中的方法），提高了调用者的开发效率。

定义方法时，要理解**返回类型**、**方法名**、**参数列表**和**方法体**的概念。通过方法的名称 + 括号运算符`()`可以调用该方法。

方法内部遇到`return`时返回，返回类型为`void`表示不返回任何值。

对于方法传值问题，基本数据类型传值：方法中对参数的修改，不会影响到方法调用方传递的变量值；引用数据类型传值：方法中对参数的修改，会影响到方法调用方传递的对象。

可变参数通常用于**参数数量是不确定的、可变化的**方法。方法重载提高了程序的兼容性和易用性，为方法提供了多种可能性。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

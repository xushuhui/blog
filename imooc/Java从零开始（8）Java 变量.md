# Java 变量

本小节我们将介绍 Java 语言中也是编程语言中的一个基本概念：变量。

通过本小节的学习，你将了解到**变量是什么，如何创建变量，如何操作变量，变量的命名规范**，这些是我们本小节学习的重点。我们也会介绍常量的概念和使用。

## 1. 什么是变量

在程序设计中，变量是指一个**包含值的存储地址**以及**对应的符号名称**。

从定义上来看，变量大概可分为两部分：变量的名字和变量的值，可通过变量名来访问变量值。简单来说，变量就是存储数据的载体。

对于 Java 语言来讲，Java 是一种强类型的语言，因此变量都必须有一个类型。在 Java 代码中，变量是长这个样子的：

```java
// =左侧的age是变量名，右侧的22是变量值
int age = 22;
```

## 2. 创建变量

我们已经在上面看到了变量在代码中的样子，现在我们来动手创建变量。

在`Java`语言中，我们需要做两件事才能创建一个变量：

1. 给变量起一个名字
2. 定义变量的数据类型

创建变量的过程也叫**声明变量**，声明变量的语法如下：

```java
DataType 变量名;
```

> **Tips**：关于数据类型，我们将在下一小节介绍。此处只需知道变量有类型之分即可。

声明一个变量，实例如下：

```java
int a;
```

上述代码，声明了一个名字为`a`，类型为**整型**的变量。

我们可以再声明一个名字为`b`，类型为**字符型**的变量，实例如下：

```java
char b;
```

再声明两个整型的变量`age`、`number`，实例如下：

```java
int age;
int number;
```

## 3. 给变量赋值

变量声明操作已经告诉了编译器变量的数据类型。数据类型决定了变量中可以存放什么，不能存放什么，上述的变量`a`为`int`类型（整型），因此只能存放整数。

一旦声明了一个变量，我们就可以使用赋值语句为其赋值，实例如下：

```java
// 变量声明语句
int a;
// 赋值语句
a = 1;
```

为变量分配值的语句，就称为赋值语句。需要特别提醒的是，语句中`=`的意义不同于数学中的等号，在 Java 中，`=`是赋值符号。

对于变量的声明和赋值操作，我们也可以将这两条语句合并成一条语句，实例如下：

```java
int a = 1;
```

另外，如果要创建多个**类型相同**的变量，我们可以在一条语句中将多个变量的声明进行分组：

```java
int a, b;
a = 1;
b = 2;
```

为了精简语法，还可以将变量的声明和赋值同时分组：

```java
int a = 1, b = 2;
```

在作用域范围内，变量的值能够随时访问或重新赋值，比如：

```java
class PrintVariable {
    public static void main(String[] args) {
        int a = 20;
    	System.out.println("修改前变量a=" + a);
    	// 将变量a重新赋值为100
    	a = 100;
    	System.out.println("修改后变量a=" + a);
    }
}
```

运行结果：

```java
修改前变量a=20
修改后变量a=100
```

上述代码中，声明了一个整型变量`a`，且初始化值为`20`，可以使用`a = 100;`这种赋值语句，对变量`a`重新赋值。

注意，在 Java 中，对变量进行重新赋值，这样的写法是错误的：

```java
int a = 20;
int a = 100;
```

编译代码，将会抛出异常：

```java
Test.java:5: 错误: 已在方法 main(String[])中定义了变量 a
	int a = 2;
	    ^
1 个错误
```

编译报错的原因很简单，第二行代码再次定义了变量`a`，而变量`a`已经在内存中存在了，不能再重复定义。这与对变量重新赋值不同。

## 4. 变量命名规范

对于变量的命名，每种编程语言都有自己的规则和约定，Java 语言也不例外。命名变量的规则和约定如下：

* **变量的名称可以是任何合法的标识符，以字母，美元符号`$`或下划线`_`开头**。但是，按照约定俗成，变量应始终以字母开头，不推荐使用美元符号和下划线开头；
* **开头后续的字符可以是字母、数字、美元符号或下划线**。为变量选择名称时，推荐使用完整的英文单词，不推荐使用单词缩写，更不要使用中文拼音。这样做有利于代码的阅读和理解。另外请牢记，选择的名称不能是**关键字**或**保留字**；
* **变量命名区分大小写**；
* **变量命名应采用小驼峰命名法**。所谓小驼峰命名法，就是如果你选择的名称只包含一个单词，那么用全部小写字母拼写该单词；如果名称包含多个单词，请将第二个单词起的每个单词的第一个字母都大写；
* **如果变量存储了一个常量值，要将每个字母大写并用下划线字符分隔每个单词**。比如`static final int MAX_NUM = 100`。按照约定俗成，除了常量命名的情况，下划线字符永远不会在其他地方使用。

下面我们对照上述规则和约定来看几个实例。

实例 1：

```java
int $a;
int _b;
int age;
```

上面 3 个变量的声明都是合法的，但是根据约定俗称，不推荐`$a`这种以美元符号开头、`_b`这种以下划线开头的命名方式。

实例 2：

```java
public class VariablesNameing2 {
    public static void main(String[] args) {
        String Name = "Colorful";
        String name = "Colorful1";
        System.out.println("Name=" + Name);
        System.out.println("name=" + name);
    }
}
```

变量名区分大小写，因此上面的`Name`和`name`是两个不同的变量。由于只有一个单词组成，应将单词使用全部小写拼写，不推荐第 1 行`Name`这种命名方式。

实例 3：

```java
// 声明姓和名变量
String firstName;
String lastname;
```

上面声明了两个变量 `firstName`和`lastname`，由于它们都是由两个单词组成，按照小驼峰命名法的约定俗成：第 2 行`lastname` 变量名由两个英文单词组成，因此推荐改为`lastName`。

## 5. 常量

所谓常量，就是恒常不变的量。我们可以将常量理解成一种特殊的变量。

与变量不同的是，一旦它被赋值后，在程序的运行过程中不允许被改变。常量使用`final`关键字修饰：

```java
final DataType 常量名 = 常量值;
```

> **Tips**：常量的命名规范与普通变量有所不同，要将每个字母大写并用下划线字符分隔每个单词。

如果我们尝试在代码中修改常量的值：

```java
class ConstantTest {
    public static void main(String[] args) {
        // 声明并初始化常量 TOTAL_NUM
        final int TOTAL_NUM = 200;
        // 对 TOTAL_NUM 重新赋值
        TOTAL_NUM = 20;
    }
}
```

编译执行代码，编译器将会报错：

```java
ConstantTest.java:6: 错误: 无法为最终变量TOTAL_NUM分配值
        TOTAL_NUM = 20;
        ^
1 个错误
```

适当地使用常量可以提高代码的安全性和可维护性。例如，我们可以在程序中声明一个表示网站名称的常量： `final String SITE_NAME = "慕课网";`，这样在所有需要用到网站名称的地方，都调用这个常量即可。这样既保证了其值不可修改，提高了代码的安全性，又统一了调用方的规范。

## 6. 变量的种类

学习变量的种类，需要你对 Java 的类和对象有一定的了解。目前不要求完全理解，等到后面学完面向对象相关知识后再回过头来回味。

Java 语言定义了以下**4 种变量** ：

* 实例变量（见代码中`instanceVariable`）
* 类变量（见代码中`classVariable`）
* 局部变量（见代码中`localVariable`）
* 参数（见代码中`parameter`和`args`）

```java
public class KindsOfVariables {
    // 1.实例变量
    public int instanceVariable = 1;
    // 2.类变量
    public static int classVariable;

    public void demoMethod(int parameter) { // 3.参数
        // 4.局部变量
        int localVariable;
    }

    public static void main(String[] args) {
        // 入口方法
    }
}
```

### 6.1 实例变量

实例变量在类范围内声明。声明时不使用`static`关键字修饰，因此也称为**非静态字段**。

实例变量用于存储对象的状态，它的值对于类的每个实例都是唯一的，每个对象都拥有自己的变量副本。只要相应的对象存在于内存中，它就占用着存储空间。

下面是一个实例变量的实例，当然，我们还没有介绍类和对象的概念，如果你了解类和对象，可以阅读源码帮助理解：

```java
public class Student {
    // 实例变量 name
    public String name;
    // 实例变量 age
    public int age;

    // 构造方法
    public Student(String name, int age) {
        this.name = name;
        this.age = age;
    }

    // 打印学生基本信息的方法
    public void printInfo() {
        System.out.println("学生姓名为：" + name);
        System.out.println("学生年龄为：" + age);
    }

    public static void  main(String[] args) {
        // 实例化两个学生对象
        Student student1 = new Student("小慕", 20);
        Student student2 = new Student("小李", 18);
        // 分别调用打印学生基本信息方法，打印两个对象下的两个实例变量
        student1.printInfo();
        System.out.println("--------分割线--------");
        student2.printInfo();
    }
}
```

运行结果：

```java
学生姓名为：小慕
学生年龄为：20
--------分割线--------
学生姓名为：小李
学生年龄为：18
```

实例变量和实例对象相关，因此需要通过`new`关键字实例化出对象，通过对象才能访问其实例变量（也称实例属性）。 上述代码中，`name`和`age`就是用于存储其对象的属性，我们实例化了两个学生对象，其实例变量`name`和`age`，在每个对象实例中是唯一的。

### 6.2 类变量

类变量在类范围内使用`static`关键字修饰，因此类变量又被称为**静态字段、静态变量**。

`static`修饰符告诉编译器，无论类被实例化多少次，类变量始终只有一个变量副本。只要类被加载到内存中，它就会存在。

另外，类变量可以被声明为**常量**，通过使用`final`关键字以表示变量永远不会改变。例如：`public static final NAME = "慕慕"`，这里的`NAME`就是不会改变的常量。再次提醒，在常量的命名规范中，要将字母全部大写。

对于类变量，我们再来看一个实例：

```java
public class Course {

  	// 类变量 courseType
  	public static String courseType = "文字课程";
  	// 常量 COURSE_NAME
  	public static final String COURSE_NAME = "Java 基础教程";

	public static void main(String[] args) {
      	// 分别打印类变量和常量
      	System.out.println(Course.courseType);
      	System.out.println(Course.COURSE_NAME);
    }
}
```

运行结果：

```java
文字课程
Java 基础教程
```

类变量和类相关，因此不需要使用`new`关键字实例化对象后再调用，可以直接通过**类名 + `.`点运算符 + 类变量名**的方式调用。

上述代码中，`courseType` 和`COURSE_NAME`都使用`static`关键字修饰，它们都可以直接通过`Course.变量名`的方式调用。

### 6.3 局部变量

局部变量是在方法范围内被声明和使用的。它们没有任何关键字修饰，可以根据**变量声明的位置**来认定局部变量（即方法的左花括号和右花括号之间），因此，局部变量只可以对声明它们的方法可见。方法返回后，它们将被销毁。

在`main`方法中的局部变量实例：

```java
public static void main(String[] args) {
  	// 局部变量 name
  	String name = "小慕";
  	// 局部变量 age
  	int age = 20;
  	System.out.println("姓名：" + name);
  	System.out.println("年龄：" + age);
}
```

我们再来看一个自定义方法中的局部变量实例：

```java
public class PrintNumber {

  	public void printer() {
  	    int num = 10;
  	    for(int i = 1; i <= num; i ++) {
	      	System.out.println(i);
        }
    }

}
```

局部变量和方法相关，因此只能在方法内部局部定义和使用，在第二个实例中没有代码注释，请思考`printer`方法中有哪几个局部变量。

### 6.4 参数

参数是用于传递给方法签名的变量（例如入口方法`main`中的`args`），它们可以在方法中的任何位置被调用。在方法执行的期间位于内存中，方法返回后被销毁。

例如，上面实例变量的实例中，`Student`类的构造方法就接收两个参数，如下为代码截取：

```java
// Student 类构造方法
public Student(String name, int age) {  // name 和 age 就是传递给Student构造方法的参数
    this.name = name;
    this.age = age;
}
```

注意，方法体中的`this.name`和`this.age`指代的是**实例变量**，而 `name`和`age`是参数，它们被用于赋值给实例变量。

## 7. 小结

在本小节中，我们知道了变量是存储数据的载体，声明变量时要注意变量的命名规范，给变量赋值时要根据声明的数据类型赋值。

Java 中变量有 4 个种类，分别是：实例变量、类变量、局部变量、参数。其中变量如果使用了`final`关键字修饰，就可定义一个不可变的常量。

当然，需要再次强调的是：不要使用中文拼音来命名变量，这是很常见的但非常不好的习惯。英文基础差的童鞋可以下载一个英文词典，当你要选择一个变量的名称时，多查阅一下单词，选择出最符合变量特征的单词。好的变量命名不仅让你的代码易于阅读理解，也有助于你更好地梳理实际编码中繁琐的业务逻辑。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

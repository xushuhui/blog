---
title: Java从零开始（15）Java 数组
zhihu-url: https://zhuanlan.zhihu.com/p/408073761
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java 数组


数组在编程语言中是非常重要的数据结构。本小节我们来一起学习 Java 数组，通过本小节的学习，你将了解到数组的基本概念，**如何声明数组**以及数组的声明有哪些方式，**如何初始化数组**以及数组初始化有哪些方式，**数组的常用操作**有哪些，多维数组的声明、创建、初始化以及迭代等等。多维数组部分将以二维数组进行为例讲解，理解了二维数组，再去理解多维数组就相对容易了。

## 1. 什么是数组

数组是**相同类型的数据**按照**顺序**组成的一种**引用数据类型**。

Java 中有基本数据类型和引用数据类型两种数据类型。数组就是引用数据类型。

## 2. 数组声明

在 Java 语言中，数组声明有两种语法：

* **中括号跟在数据类型后**：`DataType[] arrayName;`

* **中括号跟在数组名称后**：`DataType arrayName[];`

其中`DataType`是数组中要存放元素的数据类型，例如`int`、`String`等，`arrayName`是数组名称，遵守变量的命名规范。

在实际的编码中，我们更常用的是第一种语法。实例如下：

```java
// 声明一个int类型的数组，名称为 firstIntArray:
int[] firstIntArray;  // 推荐写法
int firstIntArray[];
```

也可以修改数据类型，以生成其他类型的数组。实例如下：

```java
// 声明一个float类型的数组：
float[] firstFloatArray;
// 声明一个字符串类型的数组：
String[] firstStringArray;
```

## 3. 数组创建

### 3.1 先声明再创建

先声明再创建，需要编写两条语句：

```java
DataType[] arrayName;
arrayName = new DataType[数组长度];
```

数组的创建使用`new`关键字，后面跟上一个**空格**+**数据类型**+**中括号`[]`**，中括号中间是**数组的长度**，这里是必须指定的，它是一个**整型的数字**。实例如下：

```java
// 声明一个整型的数组：
int[] intArray;
// 创建数组，长度为10
intArray = new int[10];
```

### 3.2 声明时同时创建

声明时同时创建，需要编写一条语句：

```java
DataType[] arrayName = new DataType[数组长度];
```

实例如下：

```java
// 声明时同时创建
int[] intArray = new int[10];
```

数组声明用到的中括号，也可以放到数组名称的后面，实例如下：

```java
int intArray[] = new int[10];
```

所创建的数组数据类型必须与其声明的数据类型相同。对于其他类型数组的创建，实例如下：

```java
// 创建一个数据类型为字符串、长度为5的数组：
String[] stringArray = new String[5];

// 创建一个数据类型为双精度浮点型、长度为10的数组：
double[] floatArray = new double[10];

// 创建一个数据类型为字符型、长度为10的数组：
char[] charArray;
charArray = new char[10];
```

## 4. 数组初始化

数组创建后，所有元素都初始化为默认值，整型数组的元素都为 0，浮点型都为 0. 0 ，布尔型都为 `false`。例如 `int[] arr = new int[5];`语句创建了一个数组，它在结构可参考下图：

![](https://xushuhui.gitee.io/image/imooc/5e996126099e411315580538.jpg)

除了数组的默认初始化，还有另外两种常见形式：

1. 静态初始化
2. 动态初始化

下面我们将分别介绍两种数组的初始化形式。

### 4.1 静态初始化

静态初始化数组，就是声明的同时，给数组赋值。其语法为：

```java
DataType[] arrayName = {元素1, 元素2, 元素3, 元素n};
```

赋值符号右边为用中括号所包围的 n 个元素，实例如下：

```java
int[] arr = {1, 2, 3, 4, 5};
```

上面数组`arr`存放的元素为从 1 到 5 这几个整型。其的长度为 5 。 **数组的长度就是初始化时所给数组元素的个数。**

### 4.2 动态初始化

动态初始化数组，即数组的声明以及创建是与数组的赋值操作分开进行的。实例如下：

```java
// 声明并创建数组，长度为3
int[] arr = new int[3];
// 给下标位置0赋值为1
arr[0] = 1;
// 给下标位置1赋值为2
arr[1] = 2;
// 给下标位置2赋值为3
arr[2] = 3;
```

数组是有下标的，下标从 0 开始，因此长度为 3 的数组，最大下标为 2 。下图展示了上述代码的执行流程：

![](https://xushuhui.gitee.io/image/imooc/5e99613d094675a311071211.jpg)

## 5. 数组的操作

### 5.1 数组元素的引用

语法：**数组名【下标】**，下标从`0`开始。有一个整型数组，引用其第一个元素，实例如下：

```java
public class ArrayDemo {
    public static void main(String[] args) {
        // 初始化整型数组arr：
        int[] arr = {1, 2, 3};
        // 引用arr的第一个元素：
        int element1 = arr[0];
        System.out.println("数组arr中，第一个元素为：" + element1);
    }
}
```

运行结果：

```java
数组arr中，第一个元素为：1
```

注意，上面的数组下标最大为 2，也就是说，**引用元素时，不能超过其下标最大值**。下面是一个反例：

```java
public class OutOfIndexDemo {
  	public static void main(String[] args) {
        int[] arr = {1, 2, 3};
        // 引用arr的第四个元素：
        int element4 = arr[3];
    }
}
```

由于数组中只有三个元素，但引用了索引为 3 的第四个元素，编译执行代码，会抛出如下异常：

```java
Exception in thread "main" java.lang.ArrayIndexOutOfBoundsException: Index 3 out of bounds for length 3
	at OutOfIndexDemo.main(OutOfIndexDemo.java:5)
```

### 5.2 数组元素的修改

既然可以引用数组元素，就可以修改元素，实例如下：

```java
public class ArrayDemo1 {
    public static void main(String[] args) {
        // 初始化整型数组arr：
        int [] arr = {1, 2, 3};
        // 修改下标为1的元素为4
        arr[1] = 4;
        System.out.println("数组arr中，第二个元素为：" + arr[1]);
    }
}
```

运行结果：

```java
数组arr中，第二个元素为：4
```

### 5.3 获取数组长度

我们可以使用数组的内置属性`length`来获取数组的长度，使用点操作符调用其内置属性：`arrayName.length`。实例如下：

```java
public class ArrayDemo2 {
    public static void main(String[] args) {
        // 初始化整型数组arr：
        int[] arr = {1, 2, 3};
        // 打印数组arr的长度：
        System.out.println("数组arr的长度为：" + arr.length);
    }
}
```

```java
数组arr的长度为：3
```

有时我们想引用数组中最后一个元素，由于数组下标从 0 开始，所以最后一个元素的下标为数组的长度减去 1，这时可以结合`length`属性来进行引用。实例如下：

```java
public class ArrayDemo3 {
    public static void main(String[] args) {
        // 初始化一个char类型的数组：
        char[] charArr = {'I', 'm', 'o', 'o', 'c'};
        // 引用最后一个元素：
        char lastElement = charArr[charArr.length-1];
        System.out.println("charArr数组的最后一个元素为：" + lastElement);
    }
}
```

```java
charArr数组的最后一个元素为：c
```

### 5.4 数组的迭代

#### 5.4.1 循环赋值

我们知道了数组如何赋值，在实际的编码中，对整型数组循环赋值是很常用的。

实例：

运行结果：

```java
public class ArrayDemo4 {
    public static void main(String[] args) {
        // 声明一个整型数组
        int[] arr = new int[10];
        // 使用for循环对数组arr进行遍历，将数字1到10赋值给数组arr
        for (int i = 0; i < arr.length; i++) {
            arr[i] = i + 1;
        }
        // 循环赋值完成，此时 arr = {1,2,3,4,5,6,7,8,9,10}
        // 遍历数组，打印数组中每一个元素
        for (int i = 0; i < arr.length; i++) {
            System.out.print(arr[i] + "\t");
        }
    }
}
```

```java
1	2	3	4	5	6	7	8	9	10
```

#### 5.4.2 循环引用

我们可以使用循环引用依次打印数组中的每个元素。

实例如下：

```java
public class ArrayDemo5 {
    public static void main(String[] args) {
        // 初始化一个整型数组
        int[] arr = {1,2,3,4,5};
        // 使用for循环遍历数组arr
        for (int i = 0; i < arr.length; i++) {
            // 打印位置i的元素
            System.out.println("索引位置" + i + "的元素为：" + arr[i]);
        }
    }
}
```

运行结果：

```java
索引位置0的元素为：1
索引位置1的元素为：2
索引位置2的元素为：3
索引位置3的元素为：4
索引位置4的元素为：5
```

#### 5.4.3 增强 for 循环

上一小节我们提到`for each`循环可以更简单地对数组进行迭代。以下是`for each`循环的语法结构：

```java
for(变量声明: 数组或列表) {
    // 循环体
}
```

我们来看一个实例：

```java
public class ForEachLoop {
    public static void main(String[] args) {
        // 初始化字符串数组 words
        String[] words = {"Welcome ", "to ", "imooc"};
        for(String word: words) {
            System.out.print(word);
        }
    }
}
```

运行结果：

```java
Welcome to imooc
```

实例中遍历了一个单词数组，并且将它们打印成了一个句子。`for each`循环对`words`数组进行了遍历，并依次将每个索引的值赋值给了`word`变量，然后再执行循环体中的语句。显然 `for each`循环的代码看起来更加清晰。但**增强`for`循环**无法指定遍历顺序，也无法获取数组的索引。

#### 5.4.4 更多实例

对于数组的循环还有很多案例，下面我们一起来学习几个。

实例 1，求整型数组中，偶数元素的累加和：

```java
public class ArraySumDemo {
    public static void main(String[] args) {
        // 初始化变量sum，用于存放累加和
        int sum = 0;
        // 创建一个包含100个元素的整型数组
        int[] arr = new int[100];
        // 对arr进行循环赋值，将整数1~100赋值到数组中
        for(int i = 0; i < arr.length; i++) {
            arr[i] = i + 1;
        }
        // 遍历赋值好的数组
        for(int i = 0; i < arr.length; i++) {
            if(arr[i] % 2 == 0) {
                // 如果当前元素对二取余等于0，则为偶数，执行累加计算
                sum += arr[i];
            }
        }
        // 数组循环计算结束，输出最终结果
        System.out.println("数组arr中偶数元素的累加和为：" + sum);
    }
}
```

运行结果：

```java
数组arr中偶数元素的累加和为：2550
```

实例 2，求数组中最大值：

```java
public class ArrayGetMaxDemo {
    public static void main(String[] args) {
        int[] arr = {20, 30, 15, 70, 92, 11};
        int max = arr[0];
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] > max) {
                max = arr[i];
            }
        }
        System.out.println("最大值为：" + max);
    }
}
```

运行结果：

```java
最大值为：92
```

实例 3，使用`for each`循环遍历整型数组，获取数组中的最小值：

```java
public class ArrayGetMinDemo {
    public static void main(String[] args) {
        int[] arr = {20, 30, 15, 70, 92, 11};
        int min = arr[0];
        for(int n: arr) {
            if(n < min) {
                min = n;
            }
        }
        System.out.println("最小值为：" + min);
    }
}
```

运行结果：

```java
最小值为：11
```

## 6. 多维数组

在 Java 中，多维数组也是非常常用的数据结构，下面以二维数组为例，介绍二维数组的声明、创建、初始化和引用。

### 6.1 声明

二维数组有 3 种声明方式：

* **中括号跟在数据类型后**：`DataType[][] arrayName;`（最推荐写法）
* **中括号跟在数组名后**：`DataType arrayName[][];`
* **中括号一个在前，一个在后**：`DataType[] arranName[];`

实例：

```java
// 声明一个整型的二维数组
int[][] intArray;
// 声明一个单精度浮点型的二维数组
float floatArray[][];
// 声明一个char类型的二维数组
char[] charArray[];
```

### 6.2 创建

二维数组可以想象成一个存放数据的表格，二维数组的创建和一维数组相似。可以声明数组的同时创建：

```java
javaDataType[][] arrayName = new int[行数][列数];
```

也可以先声明再创建：

```java
DataType[][] arrayName;
arrayName = new DataType[行数][列数];
```

实例：

```java
// 声明并创建一个2行3列的数组
int[][] intArray = new int[2][3];
// 声明一个单精度浮点型的二维数组
float floatArray[][];
// 创建一个3行3列的二维数组
floatArray = new float[3][3];
```

在创建数组的时候，我们也可以先不指定列的数量，再在后面进行逐行创建。

实例：

```java
// 创建一个3行的二维数组
int intArray[][] = new int[3][];
// 第一行3列
intArray[0] = new int[3];
// 第二行2列
intArray[1] = new int[2];
// 第三行1列
intArray[2] = new int[1];
```

### 6.3 初始化

多维数组也有静态初始化和动态初始化之分，对于静态初始化，一维数组使用一对大括号`{}`包含其元素，而二维数组的初始化可以使用两对大括号包含行数和列数。实例如下：

```java
// 创建一个3行2列的二维数组
int[][] intArray = {{1,2}, {3,4}, {5,6}};
```

动态初始化实例如下：

```java
// 声明并创建一个数组
int intArray[][] = new int[2][3];
// 为第1行第1列的元素赋值：
intArray[0][0] = 10;
```

> **Tips**：为数组赋值时，要注意数组的下标索引是从 0 开始的。

一维数组可以通过`length`属性获取元素个数，同样二维数组的**行数和列数**也可以使用`length`属性获取。实例如下：

```java
public class ArrayDemo6 {
    public static void main(String[] args) {
        // 初始化一个3行2列的二维数组
        int[][] intArray = {{1,2,7}, {3,4}, {5,6}};
        // 获取intArray的行数
        System.out.println("intArray的行数为：" + intArray.length);
        // 获取intArray第1行的列数
        System.out.println("intArray第一行的列数为：" + intArray[0].length);
    }
}
```

运行结果：

```java
intArray的行数为：3
intArray第一行的列数为：3
```

### 6.4 二维数组的迭代

一维数组使用单层`for`循环就可以遍历，二维数组的遍历需要使用双层嵌套`for`循环。

实例：

```java
public class ArrayDemo7 {
    public static void main(String[] args) {
        // 初始化一个3行2列的二维数组
        int[][] intArray = {{1,2,7}, {3,4}, {5,6}};
        // 遍历intArray
        for(int i = 0; i < intArray.length; i++) {
            for(int j = 0; j <intArray[i].length; j++) {
                // 打印索引位置[i][j]的元素:
                System.out.println((i+1) + "行" + (j+1) + "列的元素为：" + intArray[i][j]);
            }
            // 打印一个空行
            System.out.println();
        }
    }
}
```

运行结果：

```java
1行1列的元素为：1
1行2列的元素为：2
1行3列的元素为：7

2行1列的元素为：3
2行2列的元素为：4

3行1列的元素为：5
3行2列的元素为：6
```

## 7. 小结

本小节我们学习了 Java 数组，知道了数组是**引用数据类型**。它在创建时，会在内存中开辟一个**连续的空间**；数组是**同一数据类型**的集合。可以使用数组下标来引用数组元素，数组元素的下标都是**从 0 开始**，使用下标不要超出范围，否则会抛出异常。可以使用`length`属性获取数组的长度。对于多维数组，实际开发中建议不要使用超过三维的数组，否则会让你的程序难以理解和维护。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

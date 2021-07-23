# Java 条件语句

条件语句通过判断给定条件的**真假**来控制程序的执行。在上一小节中，我们已经简要介绍过了语句和块的概念。那么本小节我们将进一步展开介绍 `Java` 中所有类型的条件语句。

## 1. if 语句

### 1.1 语法

当我们需要根据给定的条件来决定是否执行一段代码时，`if` 语句就派上用场了。`if` 块仅在与其关联的布尔表达式为 `true` 时执行。`if` 块的结构如下：

```java
if (条件) {
    // 当条件成立时执行此处代码
}
```

大括号中包含的内容就称为语句块。

### 1.2 实例

我们来看一个 `if` 语句的实例：

```java
public class IfStatement1 {
    public static void main(String args[]) {
        int age = 18;
        if(age >= 18) {
            System.out.println("在中国你已经成年");
        }
    }
}
```

运行结果：

```java
在中国你已经成年
```

当**条件**`age >= 18` 成立时，也就是布尔表达式计算结果为 `true`（真），`if` 语句块会执行。显然上面实例中 `age` 变量的值为 18，条件是成立的，执行程序，屏幕上将会打印 “在中国你已经成年”。

如果语句块内只有一条语句，我们也可以去掉花括号 `{}`：

```java
if(age >= 18)
    System.out.println("在中国你已经成年");
```

当然，即使 `if` 语句块只有一条语句，也并不推荐这种不易于阅读的写法。

## 2. if … else … 语句

### 2.1 语法

`if` 语句可以结合 `else` 来使用，当布尔表达式计算结果为 `false` （假）时，`else` 语句块将会执行。

`if .. .else` 语句用于有条件地执行两个代码块的其中一个，具体执行哪一个代码块，取决于布尔条件的结果。

```java
if (条件) {
    // 如果条件成立，执行此处代码
} else {
    // 如果条件不成立，执行此处代码
}
```

### 2.2 实例

改写上面 `if` 语句的代码：

```java
public class IfElseStatement1 {
    public static void main(String args[]) {
        int age = 15;
        if(age >= 18) {
            System.out.println("在中国你已经成年");
        } else {
            System.out.println("在中国你还未成年");
        }
    }
}
```

运行结果：

```java
在中国你还未成年
```

我们将变量 `age` 赋值为 15，并且在 `if` 语句块的后面加了一段 `else {...}` 语句块。这个时候，条件 `age >= 18` 的计算结果就为 `false` 了，程序将会执行 `else` 语句块中的代码。执行代码，屏幕将会打印 "在中国你还未成年"。

> 补充
>
> Java 支持使用**条件表达式**（又称三目运算符）`表达式1 ? 表达式2 : 表达式3` 来简化 `if else` 语句的代码。
>
> 表达式 1 的作用等同于 `if` 语句后面的条件，如果表达式 1 为 `true`， 则执行 `表达式2`， 否则执行 `表达式3`

也就是说上面的代码可以使用三目运算符来实现同样的功能。实例如下：

```java
public class IfElseStatement2 {
    public static void main(String args[]) {
        int age = 15;
        System.out.println(age >= 18 ? "在中国你已经成年" :  "在中国你还未成年");
    }
}
```

## 3. if … else if … else 语句

### 3.1 语法

> `if` 语句可以结合 `else if` 来实现更复杂的程序分支结构：

```java
if (条件1) {
    // 如果条件1成立，执行此处代码
} else if (条件2) {
    // 如果条件1不成立，并且条件2成立，执行此处代码
} else {
    // 如果条件1、条件2都不成立，执行此处代码
}
```

### 3.2 实例

`if` 语句可以搭配任意多数量的 `else if` 语句使用，但是只能有一个 `else`。下面我们来看一段代码：

```java
// 根据给定分数向屏幕打印评级
public class IfElseIfStatement {
	  public static void main(String args[]) {
        int score = 70;
        if (score >= 90) {
          	System.out.println("优秀");
        } else if(score >= 70) {
          	System.out.println("良好");
        } else if (score >= 60) {
          	System.out.println("及格");
        } else {
          	System.out.println("不及格");
        }
    }
}
```

运行结果：

```java
良好
```

> 运行结果为 "良好"，你也许会疑惑，score 为 70，同样也大于等于 60，为什么没有输入 "及格" 呢？
>
>
> 这是由于当程序遇到符合条件的分支就会执行其分支语句块中的语句，不会执行其他分支语句块的语句。

### 3.3 嵌套 if … else 语句

你也可以在另一个 `if` 或者 `else if` 语句中使用 `if` 或者 `else if` 语句：

```java
if(条件1){
   // 如果条件1为真，执行这里的语句
   if(条件2){
      ////如果条件2为真，执行这里的语句
   }
}
```

我们来看一个嵌套语句的实例：

```java
public class IfElseStatement1 {
    public static void main(String[] args) {
      	// 初始化整型变量age，值为25
        int age = 25;
				int sex = 1;  // 此处用sex变量表示性别，1：男  2：女
        if(age >= 20) {
            System.out.println("在中国你已经成年");
          	if(sex == 2) {
              	System.out.println("并且到了法定的结婚年龄");
            }
          	if(sex == 1 && age >= 22) {
              	System.out.println("并且到了法定的结婚年龄");
            }
        } else {
            System.out.println("在中国你还未成年");
        }
    }
}
```

运行结果：

```java
在中国你已经成年
并且到了法定的结婚年龄
```

## 4. switch 语句

### 4.1 语法

`switch` 条件语句可以理解为简写版的多个 `if .. else` 语句。`switch` 语句的语法如下：

```java
switch (值) {
    case 值1:
       语句1.1
       ...
	   语句n.1
       break;
    case 值2:
        语句2.1
        ...
        语句2.n
        break;
    default:
        语句n.1
        ...
        语句n.n
}
```

switch case 语句有如下规则：

* switch 语句中的变量类型可以是： `byte`、`short`、`int`、 `char` 或者 `String`；
* switch 语句可以拥有多个 case 语句。每个 case 后面跟一个要比较的值和冒号；
* case 语句中的值的数据类型必须与变量的数据类型相同，而且只能是常量或者字面常量；
* 当变量的值与 case 语句的值相等时，那么 case 语句之后的语句开始执行，直到 break 语句出现才会跳出 switch 语句；
* 当遇到 break 语句时，switch 语句终止。程序跳转到 switch 语句后面的语句执行。case 语句不一定要包含 break 语句。如果没有 break 语句出现，程序会继续执行下一条 case 语句，直到出现 break 语句；
* switch 语句可以包含一个 default 分支，该分支一般是 switch 语句的最后一个分支（可以在任何位置，但建议在最后一个）。default 在没有 case 语句的值和变量值相等的时候执行。default 分支不需要 break 语句。

### 4.2 实例

我们先来看下面的一个简单实例：

```java
public class SwitchStatement1 {
  	public static void main(String args[]) {
        int i = 2;
      	switch (i) {
          case 1:
              // i 的值不等于1，所以不执行此处代码
              System.out.println("i的值为1");
            	break;
          case 2:
              // i 的值等于2，所以执行此处代码
              System.out.println("i的值为2");
              break;
          default:
			  // case 2 分支已执行并break，所以此处代码不会执行
              System.out.println("i的值既不等于1，也不等于2");
        }
    }
}
```

运行结果：

```java
i的值为2
```

从 JDK5 开始，`switch` 语句可以与 枚举值一起使用。 例如，下面是根据数字显示对应星期几的实例代码：

```java
public class SwitchStatement2 {
   public static void main(String args[]) {
      Day day = day.MONDAY; // 假设 Day 是一种枚举类型，里面包含星期一到星期天。
      switch (day) {
         case MONDAY :
            System.out.println("星期一");
            break;
         case TUESDAY :
          	System.out.println("星期二");
            break;
         case WEDNESDAY :
            System.out.println("星期三");
            break;
         case THURSDAY :
            System.out.println("星期四");
            break;
         case FRIDAY :
            System.out.println("星期五");
            break;
         case SATURDAY :
            System.out.println("星期六");
            break;
         case SUNDAY :
            System.out.println("星期天");
      }
   }
}
```

从 JDK8 开始，`switch` 语句可以与 `String` 值一起使用：

```java
public class SwitchStatement3 {
   public static void main(String args[]) {
      String day = "TUESDAY";
      switch (day) {
         case "Wednesday" :
            System.out.println("星期一");
            break;
         case "TUESDAY" :
          	System.out.println("星期二");
            break;
         case "WEDNESDAY" :
            System.out.println("星期三");
            break;
         case "THURSDAY" :
            System.out.println("星期四");
            break;
         case "FRIDAY" :
            System.out.println("星期五");
            break;
         case "SATURDAY" :
            System.out.println("星期六");
            break;
         case "SUNDAY" :
            System.out.println("星期天");
      }
   }
}
```

```java
星期二
```

## 5. 小结

在 Java 中，条件语句主要有 `if` 语句和 `switch` 语句两种。在实际的编码中，条件语句非常常用，要根据合适的场景选择使用，例如对于多个 `==` 判断的情况下，使用 `switch` 语句就更加清晰。而对于复杂的条件表达式，选择 `if` 语句就更适合。

对于 `if` 语句的使用，建议：

1. 每个分支都要用 `{}` 括起来；
2. 当有多个 `if ... else` 的时候，要注意判断顺序，把可能性越大的分支越排在越前面；
3. `if ... else` 语句的嵌套层级不要太深。

对于 `switch` 语句的使用，建议：

1. 每个分支都不要漏写 `break`；
2. 总是写上 `default` 分支。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

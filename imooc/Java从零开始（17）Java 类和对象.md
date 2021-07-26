# Java 类和对象

经过前面一系列的学习，你可能对类和对象已经有了一定的了解，这是因为 Java 语言是纯面向对象的编程语言，类和对象在 Java 中无处不在。

在程序设计时，我们针对事物的特征和行为使用代码进行实现，就是面向对象编程。可以毫不夸张地说，面向对象编程是中大型项目代码最好的组织形式。

本小节我们将学习**类和对象的基本概念**、**类和对象的关系**，除了字段和方法，在**类中都可以定义什么**其他内容。也会讲解**什么是实例化**、**实例化过程是怎样的**，**构造方法是什么**，**如何定义以及其特点**，**`this` 关键字的含义和使用**。

## 1. 类和对象

### 1.1 什么是类

我们在前面的学习中已经知道，**类** (class) 是一个程序的基本单位，我们之前所编写的 `.java` 文件都是类。在现实生活中，学生可以是一个类、交通工具可以是一个类、NBA 球员可以是一个类。例如，如下代码声明了一个学生类：

```java
public class ImoocStudent {
    ...
}
```

### 1.2 什么是对象

你也许听到过一句话：“万物皆对象”。是的，现实世界中所有的事物，都可以是对象。例如一个学生、一辆汽车、NBA 巨星乔丹等等，都是对象。 对象是具体的一个事物。下面我们来看一下类和对象的关系。

### 1.3 类和对象的关系

类是一个抽象的概念，可以将它理解成 “模板”，可用于产生对象；而对象是一个具体的事物。

`ImoocStudent` 是一个类，毋庸置疑，它是抽象的，不能表示一个具体的事物；而屏幕面前的你作为学生，就是一个对象，因为你是实实在在的、具体的一个学生，和你一样的在学习的其他同学，都可以理解为是由 `ImoocStudent` 类产生的一个个具体的对象。

例如，很多 NBA 球星：科比、乔丹、詹姆斯等他们都是具体的球员，具体的球员就是对象，而这样一组对象在代码中可以抽象为**类**。

## 2. 定义类

定义类的语法，相信大家都已经熟练掌握，下面我们就以学生为例，初步抽象出一个简单的学生类。

> **Tips**：在本小节的剩余内容中，我们都将围绕这个例子展开学习。下面请跟上节奏认真思考。

### 2.1 学生类

在代码中，事物的**静态特征**被抽象成**属性**，事物的**动态行为**被抽象成**方法**。一个基本的学生类可以包含若干**属性**和**方法**。

学生可以有昵称、职位、城市、性别等特征，如下是个人中心的截图：

![](//img.mukewang.com/wiki/5e97ca8c09022a8c15770733.jpg)

除了这些特征，还可以有学习课程、评论、发表手记等行为。有了这些特征和行为，我们就可以抽象出一个学生类：

```java
public class ImoocStudent {
    // 定义属性（特征）
    String nickname;  // 昵称

    String position;  // 职位

    String city;      // 城市

    String sex;          // 男 | 女

    // 定义方法（行为）
    public void studyCourse() {
        System.out.println("学习课程");
    }

    public void postComment() {
        System.out.println("发表评论");
    }

    public void postArticle() {
        System.out.println("发表手记");
    }

}
```

### 2.2 组成类的元素

> **Tips**：类组成元素属于扩展内容，此处有一个整体认识即可。随着面向对象的深入学习，知识点都会涵盖。

刚刚我们通过对现实生活中**学生**的分析，抽象出了学生类，并在类中定义了一些属性和方法。那么除了属性和方法，类中还可以有什么其他元素呢？

下面我们来详细介绍，每个类都可以由以下元素组成：

* **成员属性**：也称为字段，成员变量或实例变量，属性是用以保存**每个对象的数据**的变量，例如每个学生都可能有昵称、性别等。每个对象之间的属性相互独立，互不干扰；
* **成员方法**：也称实例方法，成员方法是对对象执行的操作，例如，一个学生有一种方法来进行发表评论；
* **静态变量**：也称为类属性，它是同一个类的任何对象所共有的。例如，一个学生类中的最后一个被添加的学生 ID，可以用静态变量标记。 一个类不管被实例化了多少对象，每个静态变量在类中仅存在一次；
* **静态方法**：也称为类方法，静态方法是不影响特定对象的方法；
* **内部类**：可以将一个类包含在另一个类中，常用于该类仅提供给声明它的类使用的情况；
* **构造方法**：生成新对象的特殊方法；
* **参数化类型**：可以在定义期间将参数化类型分配给类。 参数化类型将替换为在类实例化时指定的类型。 它由编译器完成。 它类似于 C 语言宏`#define` 语句，其中预处理器评估宏。

## 3. 对象的实例化

### 3.1 创建对象

上面我们已经定义了学生类，有了类就可以产生对象了。这里的**产生对象**指的就是对象的实例化操作，在 Java 中，使用 `new` 关键字实例化对象，其语法为：

```java
类名 对象名称 = new 类名();
```

类是对象的类型，因此我们以类名作为对象类型，对象名称的命名规范与变量相同。例如，实例化一个学生对象，对象名称命名为 `student`：

```java
ImoocStudent student = new ImoocStudent();
```

关于对象的实例化过程，可以分为两部分：

* **声明对象**：在内存的栈空间中执行。`ImoocStudent student;`
* **实例化对象**：在内存的堆空间执行。`new ImoocStudent();`

### 3.2 调用属性和方法

对象实例化后，可以调用其属性和方法，其语法为：

```java
// 调用属性
对象名.属性名;
// 调用方法
对象名.方法名();
```

下面我们在 `ImoocStudent` 类的主方法中实例化一个学生对象，并依次调用其属性和方法：

```java
public class ImoocStudent {
    // 定义属性（特征）
    String nickname;  // 昵称

    String position;  // 职位

    String city;      // 城市

    String sex;          // 男 | 女

    // 定义方法（行为）
    public void studyCourse() {
        System.out.println("学习课程");
    }

    public void postComment() {
        System.out.println("发表评论");
    }

    public void postArticle() {
        System.out.println("发表手记");
    }

    public static void main(String[] args) {
        // 实例化学生对象
        ImoocStudent student = new ImoocStudent();
        // 调用并打印成员属性
        System.out.println("昵称：" + student.nickname);
        System.out.println("职位：" + student.position);
        System.out.println("城市：" + student.city);
        System.out.println("性别：" + student.sex);
        // 调用成员方法
        student.studyCourse();
        student.postComment();
        student.postArticle();
    }

}
```

运行结果：

```java
昵称：null
职位：null
城市：null
性别：null
学习课程
发表评论
发表手记
```

根据运行结果，我们看到了成员属性和成员方法都被成功调用。但所有的属性值都为 `null`，这是因为成员属性没有默认值，系统会给一个默认初值。

> **Tips**：不同数据类型的属性会有不同的初值。此例中的属性，都为字符串，默认初值为 `null`，你也可以尝试为 `ImoocStudent` 类添加其他类型的属性，来看看它们各自的默认初值。这里不再一一演示。

### 3.3 给属性赋值

给成员属性赋值的语法为：

```java
对象名.属性名 = 属性值;
```

下面我们来修改一下在主方法中的代码：

```java
public class ImoocStudent {
    // 定义属性（特征）
    String nickname;  // 昵称

    String position;  // 职位

    String city;      // 城市

    String sex;       // 男 | 女

    // 定义方法（行为）
    public void studyCourse() {
        System.out.println("学习课程");
    }

    public void postComment() {
        System.out.println("发表评论");
    }

    public void postArticle() {
        System.out.println("发表手记");
    }

    public static void main(String[] args) {
        // 实例化学生对象
        ImoocStudent student = new ImoocStudent();
        // 给成员属性赋值
        student.nickname = "Colorful";
        student.position = "服务端工程师";
        student.city = "北京";
        student.sex = "男";
        // 调用成员属性
        System.out.println("昵称：" + student.nickname);
        System.out.println("职位：" + student.position);
        System.out.println("城市：" + student.city);
        System.out.println("性别：" + student.sex);
    }
}
```

运行结果：

```java
昵称：Colorful
职位：服务端工程师
城市：北京
性别：男
```

## 4. 构造方法

### 4.1 概述

在类中，可定义一个构造方法，也称为构造函数或构造器。它用于一些初始化操作，构造方法的语法为：

```java
public 构造方法名(参数) {
  	// 代码块
}
```

需要注意的是，与普通的自定义方法不同，构造方法没有返回类型，并且方法名要**与类名同名**。

另外，如果我们没有定义构造方法，系统会有一个默认的构造方法。换句话说：任何一个类都会有一个构造方法。

构造方法只能在对象的实例化时使用，也就是说，**构造方法只能配合 `new` 关键字使用**。

根据参数的有无，构造方法可分为两种：

* **无参构造方法**
* **有参构造方法**

### 4.2 无参构造方法

在 `ImoocStudent` 类中，我们可以定义一个无参构造方法：

```java
// 定义构造方法，无返回值并且命名与类名相同
public ImoocStudent() {
  	// 执行输出语句
  	System.out.println("无参构造方法执行了...");
}
```

使用 `new` 关键字调用构造方法：

```java
ImoocStudent student = new ImoocStudent();
```

运行结果：

```java
无参构造方法执行了...
```

### 4.3 有参构造方法

当创建实例对象时，我们经常需要同时初始化这个实例的成员属性，例如：

```java
ImoocStudent student = new ImoocStudent();
student.nickname = "慕女神";
student.position = "UI设计师";
```

虽然这样能够给我们的成员属性进行赋值，但每次需要编写 2 行这样的代码。能不能在创建对象时，就对这些属性进行赋值呢？有参构造方法解决了这样的问题，我们可以在 `ImoocStudent` 类内部定义一个有参构造方法：

```java
// 有参构造方法
public ImoocStudent(String studentNickname, String studentPosition) {
    // 将参数变量赋值给实例变量
    nickname = studentNickname;
    position = studentPosition;
}
```

构造函数中，我们将参数变量的值赋值给实例变量。使用 `new` 关键字调用构造方法：

```java
ImoocStudent student1 = new ImoocStudent("慕女神", "UI设计师");
System.out.println("昵称为：" + student1.nickname);
System.out.println("职位为：" + student1.position);
```

运行结果：

```java
昵称为：慕女神
职位为：UI设计师
```

你也可以增加构造函数的参数数量，以用来给更多的实例属性赋值：

```java
// 有参构造方法
public ImoocStudent(String studentNickname, String studentPosition, String studentCity, String studentSex) {
    nickname = studentNickname;
    position = studentPosition;
    city = studentCity;
    sex = studentSex;
}
```

构造方法与普通方法相同，也会重载，因此类中允许定义多个参数列表不同构造方法。

需要特别注意的是，我们在上面两个有参构造方法中，成员属性的命名和参数变量的命名是不同的，如果参数列表中参数变量的命名和实例属性相同。将无法完成对实例属性的赋值，也就是说，下面的写法是错误的：

```java
public ImmocStudent(String nickname) {
  	nickname = nickname;
}
```

这是因为构造方法在参数的赋值过程中，会优先到同一个作用范围内的 `nickname` 进行赋值操作，这就是所谓的**就近原则**。这时我们可以使用 `this` 关键字。

## 5. this 关键字

为了解决上面无法使用与成员属性同名称的参数对成员属性进行赋值的问题，我们可以使用 `this` 关键字：

```java
public ImmocStudent(String nickname) {
  	this.nickname = nickname;
}
```

在方法内部，`this` 关键字是当前对象的默认引用，简单说，**谁调用了它谁就是 `this`**。因此，通过 `this.field` 就可以访问当前实例的字段，解决实例变量和参数变量名称冲突的问题。

如下是本小节的最终实例代码：

```java
public class ImoocStudent {
    // 定义属性（特征）
    String nickname;  // 昵称

    String position;  // 职位

    String city;      // 城市

    String sex;       // 男 | 女

    // 无参构造方法
    public ImoocStudent() {
        // 执行输出语句
        System.out.println("无参构造方法执行了...");
    }

    // 有参构造方法
    public ImoocStudent(String nickname, String position) {
        // 将参数变量赋值给实例变量
        this.nickname = nickname;
        this.position = position;
    }

    // 有参构造方法
    public ImoocStudent(String nickname, String position, String city, String sex) {
        this.nickname = nickname;
        this.position = position;
        this.city = city;
        this.sex = sex;
    }

    // 定义方法（行为）
    public void studyCourse() {
        System.out.println("学习课程");
    }

    public void postComment() {
        System.out.println("发表评论");
    }

    public void postArticle() {
        System.out.println("发表手记");
    }

    public static void main(String[] args) {
        // 实例化学生对象
        ImoocStudent student = new ImoocStudent();
        // 给成员属性赋值
        student.nickname = "Colorful";
        student.position = "服务端工程师";
        student.city = "北京";
        student.sex = "男";
        // 调用成员属性
        System.out.println("昵称：" + student.nickname);
        System.out.println("职位：" + student.position);
        System.out.println("城市：" + student.city);
        System.out.println("性别：" + student.sex);

        ImoocStudent student1 = new ImoocStudent("慕女神", "UI设计师");
        System.out.println("昵称为：" + student1.nickname);
        System.out.println("职位为：" + student1.position);
    }

}
```

```java
无参构造方法执行了...
昵称：Colorful
职位：服务端工程师
城市：北京
性别：男
昵称为：慕女神
职位为：UI设计师
```

## 6. 小结

本小节中，我们初步熟悉了 Java 中的类和对象。我们再来回顾一下本小节的重点知识：

在面向对象中，类和对象是 “模板” 和 “实例” 的关系。

类中的实例变量是每个实例独自拥有的，它们互不干扰。

关于对象的实例化，我们是使用 `new` 关键字 + 构造方法名 () 来实现的。

当类内部我们没有显示定义构造方法时，系统会自动添加无参的构造方法；当在类内部显示定义了构造方法，系统不会自动添加无参构造方法。

`this` 关键字可以解决实例变量和参数变量冲突的问题；`this` 关键字也可以调用同一类中其他的成员方法。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

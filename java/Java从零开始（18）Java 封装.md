# Java 封装

上一小节中，我们已经对类和对象有了一个基本的认识。不止于 Java，在各个面向对象语言的书籍资料中，都会提到面向对象的三大特征：**封装、继承、多态**。本小节我们就从封装开始，探讨面向对象的特征。本小节我们将学习什么是封装、为什么需要封装，最后也会以一个 NBA 球员类的案例来实现封装。

## 1. 概念和特点

类的基本作用就是封装代码。封装将类的一些特征和行为**隐藏**在类内部，不允许类外部直接访问。

封装可以被认为是一个**保护屏障**，防止该类的代码和数据被外部类定义的代码随机访问。

我们可以通过类提供的方法来实现对隐藏信息的操作和访问。**隐藏**了对象的信息，**留出**了访问的接口。

在我们日常生活中，封装与我们息息相关，智能手机就是一个拥有良好封装的例子，我们不需要关心其内部复杂的逻辑电路设计，可以通过手机的屏幕、按键、充电口、耳机接口等等外部接口来对手机进行操作和使用。复杂的逻辑电路以及模块被**封装**在手机的内部，而留出的这些必要接口，让我们更加简便地使用手机的同时也保护了手机的内部细节。

封装有两个特点：

1. 只能通过规定的方法访问数据；
2. 隐藏类的实例细节，方便修改和实现。

## 2. 为什么需要封装

封装具有以下优点：

* 封装有利于提高类的内聚性，适当的封装可以让代码更容易理解与维护；
* 良好的封装有利于降低代码的耦合度；
* 一些关键属性只允许类内部可以访问和修改，增强类的安全性；
* 隐藏实现细节，为调用方提供易于理解的接口；
* 当需求发生变动时，我们只需要修改我们封装的代码，而不需要到处修改调用处的代码。

## 3. 实现封装

在 Java 语言中，如何实现封装呢？需要 3 个步骤。

1. 修改属性的可见性为`private`；
2. 创建公开的 getter 和 setter 方法，分别用于属性的读写；
3. 在 getter 和 setter 方法中，对属性的合法性进行判断。

我们来看一个 NBA 球员类`NBAPlayer`：

```java
class NBAPlayer {
    // 姓名
    String name;
    // 年龄
    int age;
}
```

在**类内部**（即类名后面`{}`之间的区域）定义了成员属性`name`和`age`，我们知道，在类外部调用处可以对其属性进行修改：

```java
NBAPlayer player = new NBAPlayer();
player.age = -1;
```

如下是实例代码：

```java
public class NBAPlayer {
    // 姓名
    String name;
    // 年龄
    int age;

    public static void main(String[] args) {
        NBAPlayer player = new NBAPlayer();
        player.age = -1;
        System.out.println("球员年龄为：" + player.age);
    }
}
```

运行结果：

```java
球员年龄为：-1
```

我们通过**对象名。属性名**的方式对`age`赋值为 `-1`，显然，球员的年龄为`-1`是反常理的。

下面我们对`NBAPlayer`类进行封装。

1. 我们可以使用私有化访问控制符修饰类内部的属性，让其只在类的内部可以访问：

```java
// 用private修饰成员属性，限定只能在当前类内部可以访问
private String name;
private int age;
```

`private`关键字限定了其修饰的成员只能在类内部访问，这样之后就无法在类外部使用`player.age =-1`这样的赋值方式进行赋值了。

1. 创建**公开的**（public） `getter` 和 `setter`方法：

```java
// 通常以get+属性名的方式命名 getter，返回对应的私有属性
public String getName() {
  	return name;
}

// 通常以set+属性名的方式命名 setter，给对应属性进行赋值
public void setName(String name) {
  	this.name = name;
}

public int getAge() {
  	return age;
}

public void setAge(int age) {
  	this.age = age;
}
```

顾名思义，`getter`就是取属性值，`setter`就是给属性赋值，这样在类的外部就可以通过调用其方法对属性进行操作了。

1. 对属性进行逻辑判断，以`age`属性的`setter`方法为例：

```java
public void setAge(int age) {
  	// 判断参数age的合法性
  	if(age < 0) {
      	this.age = 0;
    } else {
	  	this.age = age;
  	}
}
```

在`setAge`方法中，我们将参数`age`小于 0 的情况进行了处理，如果小于 0，直接将`age`赋值为 0。除了给默认值的方式，我们也可以抛出异常，提示调用方传参不合法。

在类外部对属性进行读写：

```java
NBAPlayer player = new NBAPlayer();
// 对属性赋值：
player.setName("詹姆斯");
player.setAge(35);
// 获取属性：
System.out.println("姓名：" + player.getName());
System.out.println("年龄：" + player.getAge());
```

试想，如果在类外部，有很多地方都会操作属性值，当属性值读写逻辑发生改变时，我们只需修改类内部的逻辑。

另外，对于有参构造方法中，对属性赋值时，直接调用其`setter`方法。无需再写重复的逻辑判断，提高代码复用性：

```java
public NBAPlayer(int age) {
  	this.setAge(age);
}
```

如下是实现封装后完整实例代码：

```java
public class NBAPlayer {
    // 姓名
    private String name;
    // 年龄
    private int age;

    // 无参构造方法
    public NBAPlayer() {

    }

    // 单参构造方法
    public NBAPlayer(int age) {
        this.setAge(age);
    }

    // 全参构造方法
    public NBAPlayer(String name, int age) {
        this.setName(name);
        this.setAge(age);
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getAge() {
        return age;
    }

    public void setAge(int age) {
        // 判断参数age的合法性
        if(age < 0) {
            this.age = 0;
        }
        this.age = age;
    }

    public static void main(String[] args) {
        NBAPlayer james = new NBAPlayer();
        // 对属性赋值：
        james.setName("詹姆斯");
        james.setAge(35);
        // 打印james实例属性
        System.out.println("姓名：" + james.getName());
        System.out.println("年龄：" + james.getAge());
        System.out.println("-------------");
        // 实例化一个新的对象
        NBAPlayer jordan = new NBAPlayer("乔丹", 60);
        // 打印jordan对象实例属性
        System.out.println("姓名：" + jordan.getName());
        System.out.println("年龄：" + jordan.getAge());
    }
}
```

运行结果：

```java
姓名：詹姆斯
年龄：35
-------------
姓名：乔丹
年龄：60
```

## 4. 小结

面向对象的三大特征：**封装、继承、多态**。

封装隐藏了对象的信息，并且留出了访问的接口。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

---
title: Java 从零开始（20）Java 多态
zhihu-url: https://zhuanlan.zhihu.com/p/408511067
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b
---

# Java 多态

本小节我们来学习面向对象的最后一大特征——多态。多态是面向对象最重要的特性。我们将介绍多态的概念和特点，并带领大家实现一个多态的案例，你将了解到**多态的实现条件**、**什么是向上转型**以及**什么是向下转型**，并学会使用`instanceof`运算符来检查对象引用是否是类型的实例。

## 1. 概念和特点

多态顾名思义就是**多种形态**，是指对象能够有多种形态。在面向对象中最常用的多态性发生在当**父类引用指向子类对象**时。在面向对象编程中，所谓多态意指相同的消息给予不同的对象会引发不同的动作。换句话说：多态意味着允许不同类的对象对同一消息做出不同的响应。

例如，火车类和飞机类都继承自交通工具类，这些类下都有各自的`run()`方法，交通工具的`run()`方法输出交通工具可以运输，而火车的`run()`方法输出火车会跑，飞机的`run()`方法则输出飞机会飞，火车和飞机都继承父类的`run()`方法，但是对于不同的对象，拥有不同的操作。

任何可以通过多个`IS-A`测试的 Java 对象都被视为多态的。在 Java 中，所有 Java 对象都是多态的，因为任何对象都能够通过`IS-A`测试以获取其自身类型和 Object 类。

## 2. 实现多态

### 2.1 实现条件

在 Java 中实现多态有 3 个必要条件：

1. 满足继承关系
2. 要有重写
3. 父类引用指向子类对象

### 2.1 实例

例如，有三个类`Pet`、`Dog`、`Cat`：

父类 Pet：

```java
class Pet {
  	// 定义方法 eat
  	public void eat() {
      	System.out.println("宠物吃东西");
    }
}
```

子类 Dog 继承 Pet

```java
class Dog extends Pet { // 继承父类
  	// 重写父类方法 eat
  	public void eat() {
      	System.out.println("狗狗吃狗粮");
    }
}
子类Cat继承Pet
class Cat extends Pet { // 继承父类
  	// 重写父类方法 eat
   	public void eat() {
      	System.out.println("猫猫吃猫粮");
    }
}
```

在代码中，我们看到`Dog`和`Cat`类继承自`Pet`类，并且都重写了其`eat`方法。

现在已经满足了实现多态的前两个条件，那么如何**让父类引用指向子类对象**呢？我们在`main`方法中编写代码：

```java
public void main(String[] args) {
  	// 分别实例化三个对象，并且保持其类型为父类Pet
  	Pet pet = new Pet();
  	Pet dog = new Dog();
  	Pet cat = new Cat();
  	// 调用对象下方法
  	pet.eat();
  	dog.eat();
  	cat.eat();
}
```

运行结果：

```java
宠物吃东西
狗狗吃狗粮
猫猫吃猫粮
```

在代码中，`Pet dog = new Dog();`、`Pet cat = new Cat();`这两个语句，把`Dog`和`Cat`对象转换为`Pet`对象，这种把一个子类对象转型为父类对象的做法称为**向上转型**。父类引用指向了子类的实例。也就实现了多态。

### 2.3 向上转型

向上转型又称为自动转型、隐式转型。向上转型就是父类引用指向子类实例，也就是子类的对象可以赋值给父类对象。例如：

```java
Pet dog = new Dog();
```

这个是因为`Dog`类继承自`Pet`类，它拥有父类`Pet`的全部功能，所以如果`Pet`类型的变量指向了其子类`Dog`的实例，是不会出现问题的。

向上转型实际上是把一个子类型安全地变成了更加**抽象**的父类型，由于所有类的根类都是`Object`，我们也把子类类型转换为`Object`类型：

```java
Cat cat = new Cat();
Object o = cat;
```

### 2.4 向下转型

向上转型是父类引用指向子类实例，那么如何让**子类引用指向父类实例**呢？使用**向下转型**就可以实现。向下转型也被称为强制类型转换。例如：

```java
// 为Cat类增加run方法
class Cat extends Pet { // 继承父类
  	// 重写父类方法 eat
   	public void eat() {
      	System.out.println("猫猫吃猫粮");
    }

  	public void run() {
      	System.out.println("猫猫跑步");
    }

  	public static void main(String[] args) {
      	// 实例化子类
      	Pet cat = new Cat();
      	// 强制类型转换，只有转换为Cat对象后，才能调用其下面的run方法
      	Cat catObj = (Cat)cat;
      	catObj.run();
    }
}
```

运行结果：

```java
猫猫跑步
```

我们为`Cat`类新增了一个`run`方法，此时我们无法通过`Pet`类型的`cat`实例调用到其下面特有的`run`方法，需要向下转型，通过`(Cat)cat`将`Pet`类型的对象强制转换为`Cat`类型，这个时候就可以调用`run`方法了。

使用向下转型的时候，要注意：**不能将父类对象转换为子类类型，也不能将兄弟类对象相互转换**。以下两种都是错误的做法：

```java
// 实例化父类
Pet pet = new Pet();
// 将父类转换为子类
Cat cat = (Cat) pet;

// 实例化Dog类
Dog dog = new Dog();
// 兄弟类转换
Cat catObj = (Cat) dog;
```

不能将父类转换为子类，因为子类功能比父类多，多的功能无法凭空变出来。兄弟类之间不能转换，这就更容易理解了，兄弟类之间同样功能不尽相同，不同的功能也无法凭空变出来。

## 3. instanceof 运算符

`instanceof`运算符用来检查对象引用是否是类型的实例，或者这个类型的子类，并返回布尔值。如果是返回`true`，如果不是返回`false`。通常可以在运行时使用 `instanceof` 运算符指出某个对象是否满足一个特定类型的实例特征。其使用语法为：

```java
<对象引用> instanceof 特定类型
```

例如，在向下转型之前，可以使用`instanceof`运算符判断，这样可以提高向下转型的安全性：

```java
Pet pet = new Cat();
if (pet instanceof Cat) {
		// 将父类转换为子类
		Cat cat = (Cat) pet;
}
```

## 4. 小结

通过本小节的学习，我们知道了多态意味着一个对象有着多重特征，可以在特定的情况下，表现出不同状态，从而对应着不同的属性和方法。实现多态有 3 个必要条件，分别是**要有继承**、**要有重写**以及**父类引用指向子类对象**，通过向上转型可以使父类引用指向子类实例；通过向下转型可以使子类引用指向父类实例，使用`instanceof`运算符可以用来检查对象引用是否是类型的实例。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

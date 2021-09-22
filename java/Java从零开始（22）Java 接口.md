# Java 接口

本小节我们将学习 Java 接口（interface），通过本小节的学习，你将了解到什么是接口、为什么需要接口、如何定义和实现接口，以及接口的特点等内容。最后我们也将对比抽象类和接口的区别。

## 1. 概念

> Java 接口是一系列方法的声明，是一些方法特征的集合，一个接口只有方法的特征没有方法的实现。

在 Java 中，被关键字 `interface` 修饰的 class 就是一个接口。接口定义了一个行为协议，可以由类层次结构中任何位置的任何类实现。接口中定义了一组抽象方法，都没有具体实现，实现该接口的类必须实现该接口中定义的所有抽象方法。

## 2. 为什么需要接口

我们知道 Java 仅支持单继承，也就是说一个类只允许有一个直接父类，这样保证了数据的安全。Java 不支持下图所示的多继承：

![](https://xushuhui.gitee.io/image/imooc/5ea6a29709ece10204090294.jpg)

接口就是为了解决 Java 单继承这个弊端而产生的，虽然一个类只能有一个直接父类，但是它可以实现多个接口，没有继承关系的类也可以实现相同的接口。继承和接口的双重设计既保持了类的数据安全也变相实现了多继承。

## 3. 接口的定义和实现

### 3.1 定义接口

#### 3.1.1 接口声明

使用 `interface` 关键字声明一个接口：

```java
public interface Person {
    ...
}
```

接口声明需要两个元素：`interface` 关键字和接口名称，`public` 修饰符表示该接口可以在任何包的任何类中使用，如果为显示指定访问修饰符，则该接口只能被在同包中的类使用。

#### 3.1.2 接口主体

接口主体中，可以定义常量和方法声明：

```java
public interface Person {
  	final String NAME = "我是Person接口中的常量";
	void walk();
  	void run();
}
```

上面的 `Person` 就是一个接口，这个接口定义了一个常量 `NAME` 和两个抽象方法 `walk()`、`run()`。

接口比抽象类更加 “抽象”，它下面不能拥有具体实现的方法，必须全部都是抽象方法，所有的方法默认都是 `public abstract` 的，所以在接口主体中的方法，这两个修饰符无需显示指定。

接口除了**方法声明**外，还可以包含**常量声明**。在接口中定义的所有的常量默认都是 `public`，`static`，和 `final` 的。

接口中的成员声明不允许使用 `private` 和 `protected` 修饰符。

### 3.2 实现接口

接口定义了一些行为协议，而实现接口的类要遵循这些协议。`implements` 关键字用于实现接口，一个类可以实现一个或多个接口，当要实现多个接口时，`implements` 关键字后面是该类要实现的以逗号分割的接口名列表。其语法为：

```java
public class MyClass implements MyInterface1, MyInterface2 {
   ...
}
```

下面是实现了 `Person` 接口的 `Student` 类的示例代码：

```java
public class Student implements Person {
    @Override
    public void walk() {
      	// 打印接口中的常量
        System.out.println(Person.NAME);
        System.out.println("学生可以走路");
    }

    @Override
    public void run() {
        System.out.println("学生可以跑步");
    }
}
```

上述代码中，`Student` 类实现了 `Person` 接口。值得注意的是，可以使用**接口名。常量名**的方式调用接口中所声明的常量：

```java
String name = Person.NAME;
```

## 4. 接口继承

接口也是存在继承关系的。接口继承使用 `extends` 关键字。例如，声明两个接口 `MyInterface1` 和 `MyInterface2`，`MyInterface2` 继承自 `MyInterface1`：

```java
// MyInterface1.java
public interface MyInterface1 {
    void abstractMethod1();
}

// MyInterface2.java
public interface MyInterface2 extends MyInterface1 {
    void abstractMethod2();
}
```

当一个类实现 `MyInterface2` 接口，将会实现该接口所继承的所有抽象方法：

```java
// MyClass.java
public class MyClass implements MyInterface2 {
    @Override
    public void abstractMethod2() {
				...
    }

    @Override
    public void abstractMethod1() {
				...
    }
}
```

值得注意的是，一个接口可以继承多个父接口，接口名放在 `extends` 后面，以逗号分割，例如：

```java
// MyInterface1.java
public interface MyInterface1 {
    void abstractMethod1();
}

// MyInterface2.java
public interface MyInterface2 {
    void abstractMethod2();
}

// MyInterface3.java
public interface MyInterface3 extends MyInterface1, MyInterface2 {
    void abstractMethod3();
}
```

补充一点，当一个实现类存在 `extends` 关键字，那么 `implements` 关键字应该放在其后：

```java
public class MyClass extends SuperClass implements MyInterface {
   ...
}
```

## 5. 默认方法和静态方法

从 **JDK 1.8** 开始，接口中可以定义默认方法和静态方法。与抽象方法不同，实现类可以不实现默认方法和类方法。

### 5.1 默认方法

#### 5.1.1 声明

我们可以使用 `default` 关键字，在接口主题中实现**带方法体**的方法，例如：

```java
public interface Person {
  	void run();

  	default void eat() {
      	System.out.println("我是默认的吃方法");
    }
}
```

#### 5.1.2 调用和重写

在实现类中，可以不实现默认方法：

```java
public class Student implements Person {
  	@Override
    public void run() {
        System.out.println("学生可以跑步");
    }
}
```

我们也可以在实现类中重写默认方法，重写不需要 `default` 关键字：

```java
public class Student implements Person {
  	@Override
    public void run() {
        System.out.println("学生可以跑步");
    }

  	// 重写默认方法
  	@Override
  	public void eat() {
      	// 使用 接口名.super.方法名() 的方式调用接口中默认方法
      	Person.super.eat();
      	System.out.println("学生吃东西");
    }
}
```

如果想要在实现类中调用接口的默认方法，可以使用**接口名。super. 方法名 ()** 的方式调用。这里的 **接口名。super** 就是接口的引用。

#### 5.1.3 使用场景

当一个方法不需要所有实现类都进行实现，可以在接口中声明该方法为默认方法；使用默认方法还有一个好处，当接口新增方法时，将方法设定为默认方法，只在需要实现该方法的类中重写它，而不需要在所有实现类中实现。

### 5.2 静态方法

#### 5.2.1 声明

使用 `static` 关键字在接口中声明静态方法，例如：

```java
public interface Person {
    void walk();
    // 声明静态方法
    static void sayHello() {
        System.out.println("Hello imooc!");
    }
}
```

#### 5.2.2 调用

类中的静态方法只能被子类继承而不能被重写，同样在实现类中，**静态方法不能被重写**。如果想要调用接口中的静态方法，只需使用 **接口名。类方法名** 的方式即可调用：

```java
public class Student implements Person {
    @Override
    public void walk() {
      	// 调用接口中的类方法
        Person.sayHello();
        System.out.println("学生会走路");
    }
}
```

## 6. 接口和抽象类的区别

1. 接口的方法默认是 public ，所有方法在接口中不能有实现（Java 8 开始接口方法可以有默认实现），而抽象类可以有非抽象的方法；
2. 接口中除了 static 、final 变量，不能有其他变量，而抽象类可以；
3. 一个类可以实现多个接口，但只能实现一个抽象类。接口自己本身可以通过 extends 关键字扩展多个接口；
4. 接口方法默认修饰符是 public ，抽象方法可以有 public 、protected 和 default 这些修饰符（抽象方法就是为了被重写所以不能使用 private 关键字修饰！）；
5. 从设计层面来说，抽象是对类的抽象，是一种模板设计，而接口是对行为的抽象，是一种行为的规范。

## 7. 多个接口中的重名成员解决方法

### 7.1 多个接口存在重名默认方法

例如有两个接口 `MyInteface1.java` 和 `MyInterface2.java`，存在相同签名的默认方法：

```java
public interface MyInterface1 {
    default void defaultMethod() {
        System.out.println("我是MyInterface1接口中的默认方法");
    }
}

public interface MyInterface2 {
    default void defaultMethod() {
        System.out.println("我是MyInterface2接口中的默认方法");
    }
}
```

当实现类实现两个接口时，同名的默认方法将会发生冲突，解决办法是**在实现类中重写这个默认方法**：

```java
public class MyClass implements MyInterface1, MyInterface2 {
	public void defaultMethod() {
      	System.out.println("我是重写的默认方法");
    }
}
```

还有一种情况：实现类所继承的父类中也存在与默认方法的同名方法，此时存在三个同名方法：

```java
// 声明父类，并在父类中也定义同名方法
public class SuperClass {
  	public void defaultMethod() {
        System.out.println("我是SuperClass中的defaultMethod()方法");
    }
}

// 实现类继承父类，并实现两个接口
public class MyClass extends SuperClass implements MyInterface1, MyInterface2 {
}
```

实例化 `MyClass` 类，调用其 `defaultMethod()` 方法：

```java
MyClass myClass = new MyClass();
myClass.defaultMethod();
```

此时编译执行，不会报错：

```java
我是SuperClass中的defaultMethod()方法
```

实际上，在没有重写的情况下，它执行了实现类的父类 `SuperClass` 的 `defaultMethod()` 方法。

### 7.2 多个接口中存在重名常量

例如有两个接口，存在重名的常量：

```java
public interface MyInterface1 {
    final int NUM = 100;
}

public interface MyInterface2 {
	final int NUM = 200;
}
```

此时在实现类中，我们可以使用**接口名。常量名**的方式分别调用：

```java
public MyClass implements MyInterface1, MyInterface2 {
	  System.out.println(MyInterface1.NUM);
	  System.out.println(MyInterface2.NUM);
}
```

当实现类将入一个继承关系时：

```java
class SuperClass {
  	static int NUM = 300;
}

public MyClass extends SuperClass implements MyInterface1, MyInterface2 {
    System.out.println(NUM);
}
```

当父类中的属性或常量与接口中的常量同名时，子类无法分辨同名的 `NUM` 是哪一个。编译程序将会报错：

```java
MyClass.java:4: 错误: 对NUM的引用不明确
        System.out.println(NUM);
                           ^
  SuperClass 中的变量 NUM 和 MyInterface1 中的变量 NUM 都匹配
1 个错误
```

此时只有在子类中声明 `NUM`，才可以通过编译：

```java
public MyClass extends SuperClass implements MyInterface1, MyInterface2 {
  	int NUM = 3;
	System.out.println(NUM);
}
```

## 8. 小结

通过本小节的学习，我们知道了 Java 的接口是为了解决其单继承的弊端而产生的，可以使用 `interface` 关键字来声明一个接口，接口内部不能有具体的方法实现。可以使用 `implements` 关键字来实现接口，一个接口可以继承多个父接口，接口名放在 `extends` 后面，以逗号分割。从 Java 8 开始，接口中可以定义默认方法和静态方法。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

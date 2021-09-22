---
title: Java 从零开始（23）Java 内部类
zhihu-url: https://zhuanlan.zhihu.com/p/408512977
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Java 内部类

本节我们将介绍 Java 中的内部类。通过本节的学习，我们将了解到**什么是内部类**，内部类的**分类**和**作用**。在内部类的分类部分，我们将逐一学习各个类型的**内部类如何定义，如何实例化以及各自的特点**，要注意区分不同类型内部类的异同。有了这些基础知识之后，我们也会结合示例介绍**为什么需要内部类**。

## 1. 概念

在 Java 语言中，可以将一个类定义在另一个类里面或者一个方法里面，我们把这样的类称为内部类。

与之对应的，包含内部类的类被称为外部类。请阅读下面的代码：

```java
// 外部类 Car
public class Car {
    // 内部类 Engine
    class Engine {
        private String innerName = "发动机内部类";
    }
}
```

代码中，`Engine` 就是内部类，而 `Car` 就是外部类。

## 2. 分类

Java 中的内部类可以分为 4 种：成员内部类、静态内部类、方法内部类和匿名内部类。接下来我们按照分类一一介绍。

### 2.1 成员内部类

#### 2.1.1 定义

成员内部类也称为普通内部类，它是最常见的内部类。可以将其看作外部类的一个成员。在成员内部类中无法声明静态成员。

如下代码中声明了一个成员内部类：

```java
// 外部类 Car
public class Car {
    // 内部类 Engine
    private class Engine {
        private void run() {
            System.out.println("发动机启动了！");
        }
    }
}
```

我们在外部类 `Car` 的内部定义了一个成员内部类 `Engine`，在 `Engine` 下面有一个 `run()` 方法，其功能是打印输出一行字符串：“发动机启动了！”。

另外，需要注意的是，与普通的 Java 类不同，含有内部类的类被编译器编译后，会生成两个独立的字节码文件：

```java
Car$Engine.class
Car.class
```

内部类 `Engine` 会另外生成一个字节码文件，其文件名为：**外部类类名 $ 内部类类名。class**。

#### 2.1.2 实例化

内部类在外部使用时，无法直接实例化，需要借助外部类才能完成实例化操作。关于成员内部类的实例化，有 3 种方法：

1. 我们可以通过 `new 外部类 ().new 内部类 ()` 的方式获取内部类的实例对象：

```java
// 外部类 Car
public class Car {

    // 内部类 Engine
    private class Engine {
        private void run() {
            System.out.println("发动机启动了！");
        }
    }

    public static void main(String[] args) {
        // 1. 实例化外部类后紧接着实例化内部类
        Engine engine = new Car().new Engine();
        // 2. 调用内部类的方法
        engine.run();
    }
}
```

运行结果：

```java
发动机启动了！
```

1. 我们可通过先实例化外部类、再实例化内部类的方法获取内部类的对象实例：

```java
public static void main(String[] args) {
    // 1. 实例化外部类
    Car car = new Car();
    // 2. 通过外部类实例对象再实例化内部类
    Engine engine = car.new Engine();
    // 3. 调用内部类的方法
    engine.run();
}
```

编译执行，成功调用了内部类的 run () 方法：

```java
$javac Car.java
java Car
发动机启动了！
```

1. 我们也可以在外部类中定义一个获取内部类的方法 `getEngine()`，然后通过外部类的实例对象调用这个方法来获取内部类的实例：

```java
// 外部类 Car
public class Car {

    // 获取内部类实例的方法
    public Engine getEngine() {
        return new Engine();
    }

    // 内部类 Engine
    private class Engine {
        private void run() {
            System.out.println("发动机启动了！");
        }
    }

    public static void main(String[] args) {
		// 1. 实例化外部类
    	Car car = new Car();
    	// 2. 调用实例方法 getEngine(), 获取内部类实例
    	Engine engine = car.getEngine();
    	// 3. 调用内部类的方法
    	engine.run();
    }
}
```

运行结果：

```java
发动机启动了！
```

这种设计在是非常常见的，同样可以成功调用执行 `run()` 方法。

#### 2.1.2 成员的访问

成员内部类可以直接访问外部类的成员，例如，可以在内部类的中访问外部类的成员属性：

```java
// 外部类 Car
public class Car {

    String name;

    public Engine getEngine() {
        return new Engine();
    }

    // 内部类 Engine
    private class Engine {
        // 发动机的起动方法
        private void run() {
            System.out.println(name + "的发动机启动了！");
        }
    }

    public static void main(String[] args) {
        // 实例化外部类
        Car car = new Car();
        // 为实例属性赋值
        car.name = "大奔奔";
        // 获取内部类实例
        Engine engine = car.getEngine();
        // 调用内部类的方法
        engine.run();
    }
}
```

观察 `Engine` 的 `run()` 方法，调用了外部类的成员属性 `name`，我们在主方法实例化 `Car` 后，已经为属性 `name` 赋值。

运行结果：

```java
大奔奔的发动机启动了！
```

相同的，除了成员属性，成员方法也可以自由访问。这里不再赘述。

还存在一个同名成员的问题：如果内部类中也存在一个同名成员，那么优先访问内部类的成员。可理解为就近原则。

这种情况下如果依然希望访问外部类的属性，可以使用`外部类名。this. 成员`的方式，例如：

```java
// 外部类 Car
public class Car {

    String name;

    public Engine getEngine() {
        return new Engine();
    }
    // 汽车的跑动方法
    public void run(String name) {
        System.out.println(name + "跑起来了！");
    }

    // 内部类 Engine
    private class Engine {
        private String name = "引擎";
        // 发动机的起动方法
        private void run() {
            System.out.println("Engine 中的成员属性 name=" + name);
            System.out.println(Car.this.name + "的发动机启动了！");
            Car.this.run(Car.this.name);
        }
    }

    public static void main(String[] args) {
        // 实例化外部类
        Car car = new Car();
        // 为实例属性赋值
        car.name = "大奔奔";
        // 获取内部类实例
        Engine engine = car.getEngine();
        // 调用内部类的方法
        engine.run();
    }
}
```

运行结果：

```java
Engine 中的成员属性 name=引擎
大奔奔的发动机启动了！
大奔奔跑起来了！
```

请观察内部类 `run()` 方法中的语句：第一行语句调用了内部类自己的属性 `name`，而第二行调用了外部类 `Car` 的属性 `name`，第三行调用了外部类的方法 `run()`，并将外部类的属性 `name` 作为方法的参数。

### 2.2 静态内部类

#### 2.2.1 定义

静态内部类也称为嵌套类，是使用 `static` 关键字修饰的内部类。如下代码中定义了一个静态内部类：

```java
public class Car1 {
    // 静态内部类
    static class Engine {
        public void run() {
            System.out.println("我是静态内部类的 run() 方法");
            System.out.println("发动机启动了");
        }
    }
}
```

#### 2.2.2 实例化

静态内部类的实例化，可以不依赖外部类的对象直接创建。我们在主方法中可以这样写：

```java
// 直接创建静态内部类对象
Engine engine = new Engine();
// 调用对象下 run() 方法
engine.run();
```

运行结果：

```java
我是静态内部类的 run() 方法
发动机启动
```

#### 2.2.2 成员的访问

在静态内部类中，只能直接访问外部类的静态成员。例如：

```java
public class Car1 {

    String brand = "宝马";

    static String name = "外部类的静态属性 name";

    // 静态内部类
    static class Engine {
        public void run() {
            System.out.println(name);
        }
    }

    public static void main(String[] args) {
        Engine engine = new Engine();
        engine.run();
    }
}
```

在 `run()` 方法中，打印的 `name` 属性就是外部类中所定义的静态属性 `name`。编译执行，将会输出：

```java
外部类的静态属性 name
```

对于内外部类存在同名属性的问题，同样遵循就近原则。这种情况下依然希望调用外部类的静态成员，可以使用`外部类名。静态成员`的方式来进行调用。这里不再一一举例。

如果想要访问外部类的非静态属性，可以通过对象的方式调用，例如在 `run()` 方法中调用 `Car1` 的实例属性 `brand`：

```java
public void run() {
    // 实例化对象
    Car1 car1 = new Car1();
    System.out.println(car1.brand);
}
```

### 2.3 方法内部类

#### 2.3.1 定义

方法内部类，是定义在方法中的内部类，也称局部内部类。

如下是方法内部类的代码：

```java
public class Car2 {

	// 外部类的 run() 方法
    public void run() {
        class Engine {
            public void run() {
                System.out.println("方法内部类的 run() 方法");
                System.out.println("发动机启动了");
            }
        }
        // 在 Car2.run() 方法的内部实例化其方法内部类 Engine
        Engine engine = new Engine();
        // 调用 Engine 的 run() 方法
        engine.run();
    }

    public static void main(String[] args) {
        Car2 car2 = new Car2();
        car2.run();
    }
}
```

运行结果：

```java
方法内部类的 run() 方法
发动机启动了
```

如果我们想调用方法内部类的 `run()` 方法，必须在方法内对 `Engine` 类进行实例化，再去调用其 `run()` 方法，然后通过外部类调用自身方法的方式让内部类方法执行。

#### 2.3.2 特点

与局部变量相同，局部内部类也有以下特点：

* 方法内定义的局部内部类只能在方法内部使用；
* 方法内不能定义静态成员；
* 不能使用访问修饰符。

也就是说，`Car2.getEngine()` 方法中的 `Engine` 内部类只能在其方法内部使用；并且不能出现 `static` 关键字；也不能出现任何的访问修饰符，例如把方法内部类 `Engine` 声明为 `public` 是不合法的。

### 2.4 匿名内部类

#### 2.4.1 定义

匿名内部类就是没有名字的内部类。使用匿名内部类，通常令其实现一个抽象类或接口。请阅读如下代码：

```java
// 定义一个交通工具抽象父类，里面只有一个 run() 方法
public abstract class Transport {
    public void run() {
        System.out.println("交通工具 run() 方法");
    }

    public static void main(String[] args) {
        // 此处为匿名内部类，将对象的定义和实例化放到了一起
        Transport car = new Transport() {
            // 实现抽象父类的 run() 方法
            @Override
            public void run() {
                System.out.println("汽车跑");
            }
        };
        // 调用其方法
        car.run();

        Transport airPlain = new Transport() {
            // 实现抽象父类的 run() 方法
            @Override
            public void run() {
                System.out.println("飞机飞");
            }
        };
        airPlain.run();

    }
}
```

运行结果：

```java
汽车跑
飞机飞
```

上述代码中的抽象父类中有一个方法 `run()`，其子类必须实现，我们使用匿名内部类的方式将子类的定义和对象的实例化放到了一起，通过观察我们可以看出，代码中定义了两个匿名内部类，并且分别进行了对象的实例化，分别为 `car` 和 `airPlain`，然后成功调用了其实现的成员方法 `run()`。

#### 2.4.2 特点

* 含有匿名内部类的类被编译之后，匿名内部类会单独生成一个字节码文件，文件名的命名方式为：`外部类名称$数字。class`。例如，我们将上面含有两个匿名内部类的 `Transport.java` 编译，目录下将会生成三个字节码文件：

```java
Transport$1.class
Transport$2.class
Transport.class
```

* 匿名内部类没有类型名称和实例对象名称；
* 匿名内部类可以继承父类也可以实现接口，但二者不可兼得；
* 匿名内部类无法使用访问修饰符、`static`、`abstract` 关键字修饰；
* 匿名内部类无法编写构造方法，因为它没有类名；
* 匿名内部类中不能出现静态成员。

#### 2.4.2 使用场景

由于匿名内部类没有名称，类的定义可实例化都放到了一起，这样可以简化代码的编写，但同时也让代码变得不易阅读。当我们在代码中只用到类的一个实例、方法只调用一次，可以使用匿名内部类。

## 3. 作用

### 3.1 封装性

内部类的成员通过外部类才能访问，对成员信息有更好的隐藏，因此内部类实现了更好的封装。

### 3.2 实现多继承

我们知道 Java 不支持多继承，而接口可以实现多继承的效果，但实现接口就必须实现里面所有的方法，有时候我们的需求只是实现其中某个方法，内部类就可以解决这些问题。

下面示例中的 `SubClass`，通过两个成员内部类分别继承 `SuperClass1` 和 `SuperClass2`，并重写了方法，实现了多继承：

```java
// SuperClass1.java
public class SuperClass1 {
    public void method1() {
        System.out.println("The SuperClass1.method1");
    }
}

// SuperClass2.java
public class SuperClass2 {
    public void method2() {
        System.out.println("The SuperClass2.method2");
    }
}

// SubClass.java
public class SubClass {
	// 定义内部类 1
    class InnerClass1 extends SuperClass1 {
        // 重写父类 1 方法
        @Override
        public void method1() {
            super.method1();
        }
    }

    // 定义内部类 2
    class InnerClass2 extends SuperClass2 {
        // 重写父类 2 方法
        @Override
        public void method2() {
            super.method2();
        }
    }

    public static void main(String[] args) {
        // 实例化内部类 1
        InnerClass1 innerClass1 = new SubClass().new InnerClass1();
        // 实例化内部类 2
        InnerClass2 innerClass2 = new SubClass().new InnerClass2();
        // 分别调用内部类 1、内部类 2 的方法
        innerClass1.method1();
        innerClass2.method2();
    }
}
```

编译执行 `SubClass.java`，屏幕将会打印：

```java
$ javac SubClass.java
$ java SubClass
The SuperClass1.method1
The SuperClass1.method2
```

### 3.3 解决继承或实现接口时的方法同名问题

请阅读如下代码：

```java
// One.java
public class One {
    public void test() {
    }
}
// Two.java
public interface Two {
    void test();
}
// Demo.java
public class Demo1 extends One implements Two {
    public void test() {

    }
}
```

此时，我们无法确定 `Demo1` 类中的 `test()` 方法是父类 `One` 中的 `test` 还是接口 `Two` 中的 `test`。这时我们可以使用内部类解决这个问题：

```java
public class Demo2 extends One {

    // 重写父类方法
    @Override
    public void test() {
        System.out.println("在外部类实现了父类的 test() 方法");
    }

    // 定义内部类
    class InnerClass implements Two {
        // 重写接口方法
        @Override
        public void test() {
            System.out.println("在内部类实现了接口的 test() 方法");
        }
    }

    public static void main(String[] args) {
        // 实例化子类 Demo2
        Demo2 demo2 = new Demo2();
        // 调用子类方法
        demo2.test();
        // 实例化子类 Demo2 的内部类
        InnerClass innerClass = demo2.new InnerClass();
        // 调用内部类方法
		innerClass.test();
    }
}
```

运行结果：

```java
在外部类实现了父类的 test() 方法
在内部类实现了接口的 test() 方法
```

## 4. 小结

本小节，我们知道了什么是内部类，也知道了在 Java 中有四种内部类：成员内部类、静态内部类、方法内部类和匿名内部类。对于它们的定义和调用也做了详细讲解，理解内部类的作用是使用好内部类的关键。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

---
title: Java 从零开始（30）Java 枚举类
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b
zhihu-url: https://zhuanlan.zhihu.com/p/412913182
---

# Java 枚举类

本小节我们将一起学习 Java 枚举类，将涉及到**什么是枚举类**、**为什么需要枚举类**，**如何自定义枚举类**以及**如何使用枚举类**，**Enum 类的常用方法**等内容。理解为什么需要枚举类以及学会自定义枚举类是本小节学习的重点。

## 1. 什么是枚举类

> 在数学和计算机科学中，一个集的**枚举**是列出某些有穷序列集的所有成员的程序，或者是一种特定类型对象的技术。

枚举是一个被命名的整型常数的集合。枚举在生活中非常常见，列举如下：

* 表示星期几：`SUNDAY`、`MONDAY`、`TUESTDAY`、`WEDNESDAY`、`THURSDAY`、`FRIDAY`、`SATURDAY`就是一个枚举；
* 性别：`MALE`（男）、`FEMALE`（女）也是一个枚举；
* 订单的状态：`PAIDED`（已付款）、`UNPAIDED`（未付款）、`FINISHED`（已完成），`CANCELED`（已取消）。

知道了什么是枚举，我们就很容易理解什么是枚举类了，简单来说，**枚举类就是一个可以表示枚举的类**，当一个类的对象只有**有限个、确定个**的时候，我们就可以定义一个枚举类来存放这些对象。

## 2. 为什么需要枚举类

如果不使用枚举类，我们想在一个类中定义星期一到星期天，就可能需要在类中使用常量来表示，实例如下：

```java
public class Weekday {
    public static final int MONDAY = 1;
    public static final int TUESDAY = 2;
    public static final int WEDNESDAY = 3;
    public static final int THURSDAY = 4;
    public static final int FRIDAY = 5;
    public static final int SATURDAY = 6;
    public static final int SUNDAY = 7;
}
```

使用一组常量表示一个枚举的集合，存在一个问题，编译器无法检测每个值的范围是否合法，例如：

```java
int day = 0; // 假设 day 的值为 0
if (day == Weekday.MON) {
    System.out.println("今天星期一");
}
```

显然，0 不在这些常量值所表示的范围（1~7）内，但编译器不会给出提示，这样的编码是非常不推荐的。

当我们在开发中需要定义一组常量的时候，建议使用枚举类。接下来我们就来看如何定义枚举类。

## 3. 如何自定义枚举类

自定义枚举类有两种方式：

1. 在 Java 5.0 之前，需要通过普通 Java 类的“改装”来定义一个枚举类；
2. 在 Java 5.0 之后，可以使用 `enum`关键字来定义枚举类。

下面我们分别来看下这两种定义枚举类的方式。

### 3.1 Java 5.0 之前自定义枚举类

在 Java 5.0 之前，想要定义一个枚举类较为繁琐，通常需要以下几个步骤：

1. 定义一个 Java 普通类作为枚举类，定义枚举类的属性，使用`private final`修饰；
2. 该类不提供外部实例化操作，因此将构造方法设置为私有，并初始化属性；
3. 在类内部，提供当前枚举类的多个对象 ，使用`public static final`修饰；
4. 提供常用的`getter`、`setter`或`toString()`方法。

下面我们定义一个用于表示性别的枚举类，并演示如何调用此枚举类，其具体实例如下：

```java
/**
 * @author colorful@TaleLin
 */
public class EnumDemo1 {

    /**
     * 性别枚举类
     */
    static class Sex {

        // 定义常量
        private final String sexName;

        // 私有化构造器，不提供外部实例化
        private Sex(String sexName) {
            // 在构造器中为属性赋值
            this.sexName = sexName;
        }

        public static final Sex MALE = new Sex("男");
        public static final Sex FEMALE = new Sex("女");
        public static final Sex UNKNOWN = new Sex("保密");

        /**
         * getter
         */
        public String getSexName() {
            return sexName;
        }

        /**
         * 重写 toString 方法，方便外部打印调试
         */
        @Override
        public String toString() {
            return "Sex{" +
                    "sexName='" + sexName + '\'' +
                    '}';
        }
    }

    public static void main(String[] args) {
        System.out.println(Sex.FEMALE.getSexName());
        System.out.println(Sex.MALE.getSexName());
        System.out.println(Sex.UNKNOWN.getSexName());
    }

}
```

运行结果：

```java
女
男
保密
```

### 3.2 Java 5.0 之后自定义枚举类

在 Java 5.0 后，可以使用`eunm`关键字来定义一个枚举类，比较便捷，推荐大家使用这个方法来定义枚举类。

通常需要以下几个步骤：

1. 使用`enum`关键字定义枚举类，这个类隐式继承自`java.lang.Enum`类；
2. 在枚举类内部，提供当前枚举类的多个对象，多个对象之间使用逗号分割，最后一个对象使用分号结尾；
3. 声明枚举类的属性和构造方法，在构造方法中为属性赋值；
4. 提供 `getter` 方法，由于`Enum`类重写了 `toString()`方法，因此一般不需要我们自己来重写。

具体实例如下：

```java
/**
 * @author colorful@TaleLin
 */
public class EnumDemo2 {

    public static void main(String[] args) {
        Sex male = Sex.MALE;
        // 打印 Sex 对象
        System.out.println(male);
        // 打印 getter 方法的值
        System.out.println(male.getSexName());
        System.out.println(Sex.FEMALE.getSexName());
        System.out.println(Sex.UNKNOWN.getSexName());
    }

}

/**
 * 使用 enum 关键字定义枚举类，默认继承自 Enum 类
 */
enum Sex {
    // 1. 提供当前枚举类的多个对象，多个对象之间使用逗号分割，最后一个对象使用分号结尾
    MALE("男"),
    FEMALE("女"),
    UNKNOWN("保密");

    /**
     * 2. 声明枚举类的属性
     */
    private final String sexName;

    /**
     * 3. 编写构造方法，为属性赋值
     */
    Sex(String sexName) {
        this.sexName = sexName;
    }

    /**
     * 4. 提供 getter
     */
    public String getSexName() {
        return sexName;
    }
}
```

运行结果：

```java
MALE
男
女
保密
```

## 4. Enum 类

`java.lang.Enum`类 是 Java 语言枚举类型的公共基类，我们使用`enum`关键字定义的枚举类，是隐式继承自`Enum`类的，下面我们来看一下`Enum`类的常用方法：

* `values()`：返回枚举类型的对象数组。改方法可以很方便的遍历所有的枚举值；
* `valueOf()`：可以把一个字符串转换为对应的枚举类对象。要求字符串必须是枚举类对象的“名字”，如果不是，会抛出`IllegalArguementException`；
* `toString()`：返回当前枚举类对象常量的名称。

这 3 个方法使用起来比较简单，因此我们写在一个实例中，代码如下：

```java
/**
 * @author colorful@TaleLin
 */
public class EnumDemo3 {

    public static void main(String[] args) {
        Sex male = Sex.MALE;
        System.out.println("调用 toString() 方法：");
        System.out.println(male.toString());

        System.out.println("调用 values() 方法：");
        Sex[] values = Sex.values();
        for (Sex value : values) {
            System.out.println(value);
        }

        System.out.println("调用 valueOf() 方法：");
        Sex male1 = Sex.valueOf("MALE");
        System.out.println(male1);
    }

}

/**
 * 使用 enum 关键字定义枚举类，默认继承自 Enum 类
 */
enum Sex {
    // 1. 提供当前枚举类的多个对象，多个对象之间使用逗号分割，最后一个对象使用分号结尾
    MALE("男"),
    FEMALE("女"),
    UNKNOWN("保密");

    /**
     * 2. 声明枚举类的属性
     */
    private final String sexName;

    /**
     * 3. 编写构造方法，为属性赋值
     */
    Sex(String sexName) {
        this.sexName = sexName;
    }

    // 提供 getter 和 setter

    public String getSexName() {
        return sexName;
    }
}
```

运行结果：

```java
调用 toString() 方法：
MALE
调用 values() 方法：
MALE
FEMALE
UNKNOWN
调用 valueOf() 方法：
MALE
```

值得注意的是，当调用`valuOf()`方法时，我们传递的对象的“名字”，在枚举类中并不存在，此时会抛出运行时异常：`IllegalArgumentException`，实例如下：

```java
/**
 * @author colorful@TaleLin
 */
public class EnumDemo3 {

    public static void main(String[] args) {
        System.out.println("调用 valueOf() 方法：");
        Sex male1 = Sex.valueOf("MALE1");
        System.out.println(male1);
    }

}

/**
 * 使用 enum 关键字定义枚举类，默认继承自 Enum 类
 */
enum Sex {
    // 1. 提供当前枚举类的多个对象，多个对象之间使用逗号分割，最后一个对象使用分号结尾
    MALE("男"),
    FEMALE("女"),
    UNKNOWN("保密");

    /**
     * 2. 声明枚举类的属性
     */
    private final String sexName;

    /**
     * 3. 编写构造方法，为属性赋值
     */
    Sex(String sexName) {
        this.sexName = sexName;
    }

    // 提供 getter 和 setter

    public String getSexName() {
        return sexName;
    }
}
```

运行结果：

```java
调用 valueOf() 方法：
Exception in thread "main" java.lang.IllegalArgumentException: No enum constant Sex.MALE1
	at java.base/java.lang.Enum.valueOf(Enum.java:273)
	at Sex.valueOf(EnumDemo3.java:17)
	at EnumDemo3.main(EnumDemo3.java:8)
```

## 5. 小结

通过本小节的学习，我们知道了**枚举类就是一个可以表示枚举的类**，当一个类的对象只有**有限个、确定个**的时候，我们就可以定义一个枚举类来存放这些对象。使用枚举类可以规避编译器无法检测每个值的范围是否合法的问题。自定义枚举类可以有两种方式，更推荐使用`enum`关键字来定义枚举类。所有通过`enum`关键字定义的枚举类都继承自`java.lang.Enum`类，要了解该类的常用方法的使用。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

---
title: Java 从零开始（34）Java 注解
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
zhihu-url: https://zhuanlan.zhihu.com/p/412975832
---
# Java 注解

本小节我们将学习 `Java5` 引入的一种机制 —— 注解（Annotation）。通过本小节的学习，你将了解**什么是注解**，**注解的作用**，Java 中**内置注解**有哪些以及**注解的分类**，如何**自定义注解**，如何**处理注解**等内容。

## 1. 什么是注解

> Java 注解（Annotation）又称为 Java 标注，是 `Java5`开始支持加入源代码的特殊语法元数据。Java 语言中的类、方法、变量、参数和包等都可以被标注。Java 标注可以通过反射获取标注的内容。在编译器生成`class`文件时，标注可以被嵌入到字节码中。Java 虚拟机可以保留标注内容，在运行时可以获取到标注内容。

注解是一种用于做标注的“元数据”，什么意思呢？你可以将注解理解为一个标签，这个标签可以标记类、方法、变量、参数和包。

回想我们学习继承时，子类若重写了父类的方法，可以在子类重写的方法上使用`@Override`注解：

![](https://xushuhui.gitee.io/image/imooc/5ed7237a094258a311300323.jpg)

将`@Override` 注解标注在子类重写的方法上，可检查该方法是否正确地重写了父类的方法，如有错误将会编译报错。

## 2. 注解的作用

### 2.1 内置的注解

我们先看一下 Java 提供了哪些内置注解，以及这些注解的作用。（大致了解即可）

Java 定义了一套注解，共有 10 个，5 个在 `java.lang` 包中，剩下 5 个在 `java.lang.annotation` 包中。

#### 2.1.1 用在代码的注解

* `@Override`：检查该方法是否正确地重写了父类的方法。如果重写错误，会报编译错误；

* `@Deprecated`：标记过时方法。如果使用该方法，会报编译警告；

* `@SuppressWarnings`：指示编译器去忽略注解中声明的警告；

* `@SafeVarargs`：Java 7 开始支持，忽略任何使用参数为泛型变量的方法或构造函数调用产生的警告；

* `@FunctionalInterface`：Java 8 开始支持，标识一个匿名函数或函数式接口。

#### 2.1.2 用在其他注解的注解

此类注解也称为**元注解**（meta annotation），在下面学习定义注解的时候，我们将会详细讲解。

* `@Retention`：标识这个注解怎么保存，是只在代码中，还是编入`class`文件中，或者是在运行时可以通过反射访问；

* `@Documented`：标记这些注解是否包含在用户文档中；

* `@Target`：标记这个注解应该是哪种 Java 成员；

* `@Inherited`：标记这个注解是继承于哪个注解类；

* `@Repeatable`：Java 8 开始支持，标识某注解可以在同一个声明上使用多次。

### 2.2 分类

Java 注解可以分为 3 类：

1. **由编译器使用的注解**：如`@Override`、`@Deprecated`、`@SupressWarnings`等；

2. **由工具处理`.class`文件使用的注解**：比如有些工具会在加载`class`的时候，对`class`做动态修改，实现一些特殊的功能。这类注解会被编译进入`.class`文件，但加载结束后并不会存在于内存中。这类注解只被一些底层库使用，一般我们不必自己处理；

3. **在程序运行期间能够读取的注解**：它们在加载后一直存在于`JVM`中，这也是最常用的注解。

## 3. 定义注解

学会使用注解非常简单，很多框架都会提供丰富的注解文档（例如 Spring）。但关键的一点在于定义注解，知道如何定义注解，才能看懂别人定义的注解。

下面我们来定义一个注解。

想要定义一个注解，通常可分为 3 步：

1. 创建注解；

2. 定义注解的参数和默认值；

3. 用元注解配置注解。

关于这 3 个步骤是什么意思，如何来做，我们下面将来详细讲解。

### 3.1 创建注解

注解通过`@interface`关键字来定义。例如，我们想要定义一个可用于检查字符串长度的注解，实例如下：

```java
public @interface Length {

}
```

> Tips：通过 `@interface` 关键字定义注解，通过关键字 `interface`定义接口。注意两者不要混淆。

在`IDEA`中，我们可以在新建 Java 类的时候，选择新建注解：

![](https://xushuhui.gitee.io/image/imooc/5ed7239b09e0185c06640362.jpg)

输入我们要定义的注解的名称（遵循类命名规范），即可创建一个注解：

![](https://xushuhui.gitee.io/image/imooc/5ed723b009f747ed06280248.jpg)

### 3.2 定义参数和默认值

注解创建完成后，可以向注解添加一些要接收的参数，下面我们为`@Length`注解添加 3 个参数：

```java
public @interface Length {

    int min() default 0;

    int max() default Integer.MAX_VALUE;

    String message() default "长度不合法";

}
```

注解的参数类似无参数方法。另外参数的类型可以是基本数据类型、`String`类型、枚举类型、`Class`类型、`Annotation`类型以及这些类型的数组。

如果注解中只有一个参数，或者这个参数是最常用的参数，那么应将此参数命名为`value`。在调用注解时，如果参数名称是`value`，且只有一个参数，那么可以省略参数名称。（由于此注解没有最常用特征的参数，没有使用`value`）

可以使用`default`关键字来指定参数的默认值，推荐为每个参数都设定一个默认值。

### 3.3 用元注解配置注解

在前面学习 Java 内置的注解的时候，我们已经了解了元注解，元注解就是用于修饰其他注解的注解。

通常只需使用这些内置元注解，就可以基本满足我们自定义注解的需求。下面我们将会详解 Java 内置的 5 个元注解，你将会了解为什么需要这些元注解。

#### 3.3.1 @Retention

`Retention`译为保留。`@Retention`注解定义了一个注解的生命周期（我们前面对于 Java 注解的分类，就是通过其生命周期来划定界限的）。它可以有如下几种取值：

* `RetentionPolicy.SOURCE`：注解只在源码阶段保留，在编译器进行编译时它将被丢弃忽视；

* `RetentionPolicy.CLASS`：注解只被保留到编译进行的时候，它并不会被加载到 JVM 中；

* `RetentionPolicy.RUNTIME`：注解可以保留到程序运行的时候，它会被加载进入到 JVM 中，所以在程序运行时可以获取到它们。

下面我们使用`@Retention`注解来指定我们自定义的注解`@Length`的生命周期，实例如下：

```java
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;

@Retention(RetentionPolicy.RUNTIME)
public @interface Length {

    int min() default 0;

    int max() default Integer.MAX_VALUE;

    String message() default "长度不合法";

}
```

上面的代码中，我们指定 @Length 注解可以在程序运行期间被获取到。

#### 3.3.2 @Documented

这个元注解的作用很简单，标注了此注解的注解，能够将注解中的元素包含到 Javadoc 中去。因此不做过多解释。

#### 3.3.3 @Target

`@Target` 注解是最为常用的元注解，我们知道注解可以被应用于类、方法、变量、参数和包等处，`@Target` 注解可以指定注解能够被应用于源码中的哪些位置，它可以有如下几种取值：

* `ElementType.ANNOTATION_TYPE`：可以给一个注解进行注解；

* `ElementType.CONSTRUCTOR`：可以给构造方法进行注解；

* `ElementType.FIELD`：可以给属性进行注解；

* `ElementType.LOCAL_VARIABLE`：可以给局部变量进行注解；

* `ElementType.METHOD`：可以给方法进行注解；

* `ElementType.PACKAGE`：可以给一个包进行注解；

* `ElementType.PARAMETER`：可以给一个方法内的参数进行注解；

* `ElementType.TYPE`：可以给一个类型进行注解，比如类、接口、枚举。

例如，我们定义注解`@Length`只能用在类的属性上，可以添加一个`@Target(ElementType.FIELD)`：

```java
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.FIELD)
public @interface Length {

    int min() default 0;

    int max() default Integer.MAX_VALUE;

    String message() default "长度不合法";

}
```

`@Target`注解的参数也可以接收一个数组。例如，定义注解`@Length`可以用在属性或局部变量上：

```java
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Retention(RetentionPolicy.RUNTIME)
@Target({ElementType.FIELD, ElementType.LOCAL_VARIABLE})
public @interface Length {

    int min() default 0;

    int max() default Integer.MAX_VALUE;

    String message() default "长度不合法";

}
```

至此，我们就完成了 `@Length` 注解的定义。下面，我们再来看下剩余的两个元注解。

#### 3.3.4 @Inherited

使用`@Inherited`定义子类是否可继承父类定义的注解。`@Inherited`仅针对`@Target(ElementType.TYPE)`类型的注解有效，并且仅针对类的继承有效，对接口的继承无效：

```java
@Inherited
@Target(ElementType.TYPE)
public @interface TestAnnotation {
  	String value() default "test";
}
```

在使用的时候，如果一个类用到了`@TestAnnotation`：

```java
@TestAnnotation("测试注解")
public class Pet {
}
```

则它的子类默认也定义了该注解：

```java
public class Cat extends Pet {
}
```

#### 3.3.5 @Repeatable

使用`@Repeatable`这个元注解可以定义注解是否可重复。

例如，一个注解用于标注一个人的角色，他可以是学生，也可以是生活委员。

```java
@Target(ElementType.TYPE)
@Repeatable(Roles.class)
public @interface Role {
    String value() default "";
}

@Target(ElementType.TYPE)
public @interface Roles {
    Role[] value();
}
```

`@Repeatable` 元注解标注了`@Role`。而 `@Repeatable` 后面括号中的类相当于一个容器注解，按照规定，它里面必须要有一个 value 的属性，属性类型是一个被 @Repeatable 注解过的注解数组。

经过`@Repeatable`修饰后，在某个类型声明处，就可以添加多个`@Role`注解：

```java
@Role("学生")
@Role("生活委员")
public class Student {
}
```

## 4. 处理注解

### 4.1 尝试使用注解

我们已经完成了`@Length`注解的定义：

```java
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Retention(RetentionPolicy.RUNTIME)
@Target({ElementType.FIELD, ElementType.LOCAL_VARIABLE})
public @interface Length {

    int min() default 0;

    int max() default Integer.MAX_VALUE;

    String message() default "长度不合法";

}
```

现在就可以在字段上标注这个注解了：

```java
public class Student {

    // 标注注解
    @Length(min = 2, max = 5, message = "昵称的长度必须在 2~5 之间")
    private String nickname;

    public Student(String nickname) {
        this.setNickname(nickname);
    }

    public String getNickname() {
        return nickname;
    }

    public void setNickname(String nickname) {
        this.nickname = nickname;
    }

    public static void main(String[] args) {
        // 实例化对象
        Student student = new Student("我的名字很长很长");
        System.out.println(student.getNickname());
    }
}

```

上面代码中，我们在`Student`类中的`nickname`字段上标注了`@Length`注解，限定了其长度。

那么现在，是不是将其标注在字段上面就可以自动检查字段的长度了呢？答案是否定的。

运行过程如下，昵称长度不合法，但并没有抛出任何异常：

![](https://xushuhui.gitee.io/image/imooc/5ed736180ae3e6ee14921382.jpg)

Java 的注解本身对代码逻辑没有任何影响，它只是一个标注。想要检查字段的长度，就要读取注解，处理注解的参数，可以使用反射机制来处理注解。

### 4.2 通过反射读取注解

我们先来学习一下通过反射读取注解内容相关的 `API`。

通过反射，判断某个注解是否存在于`Class`、`Field`、`Method`或`Constructor`：

* `Class.isAnnotationPresent(Class)`

* `Field.isAnnotationPresent(Class)`

* `Method.isAnnotationPresent(Class)`

* `Constructor.isAnnotationPresent(Class)`

`isAnnotationPresent()`方法的返回值是布尔类型，例如，判断`Student`类中的`nickname`字段上，是否存在`@Length`注解：

```java
boolean isLengthPresent = Student.class.getDeclaredField("nickname").isAnnotationPresent(Length.class);
```

通过反射，获取 Annotation 对象：

使用反射 `API` 读取 Annotation：

* `Class.getAnnotation(Class)`

* `Field.getAnnotation(Class)`

* `Method.getAnnotation(Class)`

* `Constructor.getAnnotation(Class)`

例如，获取`nickname`字段标注的`@Length`注解：

```java
Length annotation = Student.class.getDeclaredField("nickname").getAnnotation(Length.class);
```

通过反射读取注解的完整实例如下：

```java
public class Student {

    // 标注注解
    @Length(min = 2, max = 5, message = "昵称的长度必须在 2~6 之间")
    private String nickname;

    public Student(String nickname) {
        this.setNickname(nickname);
    }

    public String getNickname() {
        return nickname;
    }

    public void setNickname(String nickname) {
        this.nickname = nickname;
    }

    public static void main(String[] args) throws NoSuchFieldException {
        boolean isLengthPresent = Student.class.getDeclaredField("nickname").isAnnotationPresent(Length.class);
        if (isLengthPresent) {
            Length annotation = Student.class.getDeclaredField("nickname").getAnnotation(Length.class);
            // 获取注解的参数值
            int min = annotation.min();
            int max = annotation.max();
            String message = annotation.message();
            // 打印参数值
            System.out.println("min=" + min);
            System.out.println("max=" + max);
            System.out.println("message=" + message);
        } else {
            System.out.println("没有在 nickname 字段上找到@Length 注解");
        }
    }
}
```

运行结果：

```java
min=2
max=5
message=昵称的长度必须在 2~6 之间
```

运行过程如下：

![](https://xushuhui.gitee.io/image/imooc/5ed73daa0ab9b35423221470.jpg)

### 4.3 编写校验方法

获取到了注解以及其内容，我们就可以编写一个校验方法，来校验字段长度是否合法了。我们在`Student`类中新增一个`checkFieldLength()`方法，用于检查字段长度是否合法，如果不合法则抛出异常。

完整实例如下：

```java
import java.lang.reflect.Field;

public class Student {

    // 标注注解
    @Length(min = 2, max = 5, message = "昵称的长度必须在 2~5 之间")
    private String nickname;

    public Student(String nickname) {
        this.setNickname(nickname);
    }

    public String getNickname() {
        return nickname;
    }

    public void setNickname(String nickname) {
        this.nickname = nickname;
    }

    public void checkFieldLength(Student student) throws IllegalAccessException {
        // 遍历所有 Field
        for (Field field: student.getClass().getDeclaredFields()) {
            // 获取注解
            Length annotation = field.getAnnotation(Length.class);
            if (annotation != null) {
                // 获取字段
                Object o = field.get(student);
                if (o instanceof String) {
                    String stringField = (String) o;
                    if (stringField.length() < annotation.min() || stringField.length() > annotation.max()) {
                        throw new IllegalArgumentException(field.getName() + ":" + annotation.message());
                    }
                }
            }
        }
    }

    public static void main(String[] args) throws NoSuchFieldException, IllegalAccessException {
        Student student = new Student("小");
        student.checkFieldLength(student);
    }
}
```

运行结果：

```java
Exception in thread "main" java.lang.IllegalArgumentException: nickname 昵称的长度必须在 2~5 之间
	at Student.checkFieldLength(Student.java:32)
	at Student.main(Student.java:41)
```

运行过程如下：

![](https://xushuhui.gitee.io/image/imooc/5ed73f3e0afa9a8923221470.jpg)

## 5. 小结

通过本小节的学习，我们知道了注解是 Java 语言的一种标注，Java 内置 10 个注解，要大致了解每个注解的作用。使用`@interface`关键字自定义注解；为注解定义参数的时候要注意其参数的类型，推荐为每个参数都设置默认值；想要自定义注解，必须了解 Java 中内置的 5 个元注解如何使用。

我们自定义的注解通常是用于运行时读取的，因此必须使用`@Retention(RetentionPolicy.RUNTIME)`进行标注。本小节的概念较多且较为抽象，建议读者亲手去编写几个注解，再来阅读本节内容会有更好的理解。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

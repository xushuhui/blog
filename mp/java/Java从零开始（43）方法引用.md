---
title: Java 从零开始（43）方法引用
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
zhihu-url: https://zhuanlan.zhihu.com/p/415787496
---
# 方法引用

通过前两个小节对`Lambda`表达式的学习，本小节我们来介绍一个更加深入的知识点 —— 方法引用。通过本小节的学习，你将了解到**什么是方法引用**，**方法引用的基础语法**，方法引用的**使用条件和使用场景**，**方法引用的分类**，**方法引用的使用实例**等内容。

## 1. 什么是方法引用

方法引用（`Method References`）是一种**语法糖**，它本质上就是 `Lambda` 表达式，我们知道`Lambda`表达式是函数式接口的实例，所以说方法引用也是函数式接口的实例。

> **Tips**：什么是语法糖？语法糖（Syntactic sugar），也译为糖衣语法，是由英国计算机科学家彼得·约翰·兰达（`Peter J. Landin`）发明的一个术语，指计算机语言中添加的某种语法，这种语法对语言的功能并没有影响，但是更方便程序员使用。通常来说使用语法糖能够增加程序的可读性，从而减少程序代码出错的机会。
>
> 可以将语法糖理解为汉语中的成语，用更简练的文字表达较复杂的含义。在得到广泛接受的情况下，可以提升交流的效率。

我们来回顾一个之前学过的实例：

```java
import java.util.function.Consumer;

public class MethodReferencesDemo1 {

    public static void main(String[] args) {
        Consumer<String> consumer = s -> System.out.println(s);
        consumer.accept("只消费，不返回");
    }

}
```

运行结果：

```java
只消费，不返回
```

上例是 Java 内置函数式接口中的**消费型接口**，如果你使用`idea`编写代码，`System.out.println(s)`这个表达式可以一键替换为方法引用，将鼠标光标放置到语句上，会弹出提示框，再点击`Replace lambda with method reference`按钮即可完成一键替换：

![](https://xushuhui.gitee.io/image/imooc/5f05270009e2e83622850601.jpg)

替换为方法引用后的实例代码：

```java
import java.util.function.Consumer;

public class MethodReferencesDemo1 {

    public static void main(String[] args) {
        Consumer<String> consumer = System.out::println;
        consumer.accept("只消费，不返回");
    }

}
```

运行结果：

```java
只消费，不返回
```

我们看到`System.out.println(s)`这个表达式被替换成了`System.out::println`，同样成功执行了代码，看到这里，同学们脑袋上可能全是问号，这是什么语法？我们之前怎么没提过？别着急，我们马上就来讲解语法规则。

## 2. 语法

方法引用使用一对冒号（`::`）来引用方法，格式如下：

```java
类或对象 :: 方法名
```

上面实例中方法引用的代码为：

```java
System.out::println
```

其中`System.out`就是`PrintStream`类的对象，`println`就是方法名。

## 3. 使用场景和使用条件

方法引用的使用场景为：当要传递给`Lambda`体的操作，已经有实现的方法了，就可以使用方法引用。

方法引用的使用条件为：接口中的抽象方法的形参列表和返回值类型与方法引用的方法形参列表和返回值相同。

## 4. 方法引用的分类

对于方法引用的使用，通常可以分为以下 4 种情况：

1. `对象 :: 非静态方法`：对象引用非静态方法，即实例方法；
2. `类 :: 静态方法`：类引用静态方法；
3. `类 :: 非静态方法`：类引用非静态方法；
4. `类 :: new`：构造方法引用。

下面我们根据以上几种情况来看几个实例。

### 4.1 对象引用实例方法

对象引用实例方法，我们已经在上面介绍过了，`System.out`就是对象，而`println`就是实例方法，这里不再赘述。

### 4.2 类引用静态方法

类引用静态方法，请查看以下实例：

```java
import java.util.Comparator;

public class MethodReferencesDemo2 {

    public static void main(String[] args) {
        // 使用 Lambda 表达式
        Comparator<Integer> comparator1 = (t1, t2) -> Integer.compare(t1, t2);
        System.out.println(comparator1.compare(11, 12));

        // 使用方法引用，类 :: 静态方法（ compare() 为静态方法）
        Comparator<Integer> comparator2 = Integer::compare;
        System.out.println(comparator2.compare(12, 11));
    }

}
```

运行结果：

```java
-1
1
```

查看 Java 源码，可观察到`compare()`方法是静态方法：

![](https://xushuhui.gitee.io/image/imooc/5f05281b098be68b14730840.jpg)

我们再来看一个实例：

```java
import java.util.Comparator;
import java.util.function.Function;

public class MethodReferencesDemo3 {

    public static void main(String[] args) {
        // 使用 Lambda 表达式
        Function<Double, Long> function1 = d -> Math.round(d);
        Long apply1 = function1.apply(1.0);
        System.out.println(apply1);

        // 使用方法引用，类 :: 静态方法（ round() 为静态方法）
        Function<Double, Long> function2 = Math::round;
        Long apply2 = function2.apply(2.0);
        System.out.println(apply2);
    }

}
```

运行结果：

```java
1
2
```

### 4.3 类引用实例方法

类引用实例方法，比较难以理解，请查看以下实例：

```java
import java.util.Comparator;

public class MethodReferencesDemo4 {

    public static void main(String[] args) {
        // 使用 Lambda 表达式
        Comparator<String> comparator1 = (s1, s2) -> s1.compareTo(s2);
        int compare1 = comparator1.compare("Hello", "Java");
        System.out.println(compare1);

        // 使用方法引用，类 :: 实例方法（ compareTo() 为实例方法）
        Comparator<String> comparator2 = String::compareTo;
        int compare2 = comparator2.compare("Hello", "Hello");
        System.out.println(compare2);
    }

}
```

运行结果：

```java
-2
0
```

`Comparator`接口中的`compare(T t1, T t2)`抽象方法，有两个参数，但是`String`类下的实例方法`compareTo(String anotherString)`只有 1 个参数，为什么这种情况下还能使用方法引用呢？这属于一个特殊情况，当函数式接口中的抽象方法有两个参数时，已实现方法的第 1 个参数作为方法调用者时，也可以使用方法引用。此时，就可以使用类来**引用**实例方法了（即实例中的`String::compareTo`）。

### 4.4 类引用构造方法

类引用构造方法，可以直接使用`类名 :: new`，请查看如下实例：

```java
import java.util.function.Function;
import java.util.function.Supplier;

public class MethodReferencesDemo5 {

    static class Person {
        private String name;

        public Person() {
            System.out.println("无参数构造方法执行了");
        }

        public Person(String name) {
            System.out.println("单参数构造方法执行了");
            this.name = name;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

    }

    public static void main(String[] args) {

        // 使用 Lambda 表达式，调用无参构造方法
        Supplier<Person> supplier1 = () -> new Person();
        supplier1.get();

        // 使用方法引用，引用无参构造方法
        Supplier<Person> supplier2 = Person::new;
        supplier2.get();

        // 使用 Lambda 表达式，调用单参构造方法
        Function<String, Person> function1 = name -> new Person(name);
        Person person1 = function1.apply("小慕");
        System.out.println(person1.getName());

        // 使用方法引用，引用单参构造方法
        Function<String, Person> function2 = Person::new;
        Person person2 = function1.apply("小明");
        System.out.println(person2.getName());
    }

}
```

运行结果：

```java
无参数构造方法执行了
单参数构造方法执行了
小慕
单参数构造方法执行了
小明
```

在实例中，我们使用了`Lambda`表达式和方法引用两种方式，分别调用了静态内部类`Person`的无参和单参构造方法。函数式接口中的抽象方法的形参列表与构造方法的形参列表一致，抽象方法的返回值类型就是构造方法所属的类型。

## 5. 小结

通过本小节的学习，我们知道了方法引用是一个语法糖，它本质上还是`Lambda`表达式。方法引用使用一对冒号（`::`）来引用方法。要传递给`Lambda`体的操作，已经有实现的方法了，就可以使用方法引用；想要使用方法引用，就要求接口中的抽象方法的形参列表和返回值类型与方法引用的方法形参列表和返回值相同。方法引用可能较为抽象，希望同学们课下多加练习。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

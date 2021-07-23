# Java 泛型

本小节我们将学习 Java5 以后出现的一个特性：**泛型（`Generics`）**。通过本小节的学习，你将了解到**什么是泛型**，**为什么需要泛型**，如何使用泛型，如何自定义泛型，类型通配符等知识。

## 1. 什么是泛型

泛型不只是 Java 语言所特有的特性，泛型是程序设计语言的一种特性。允许程序员在强类型的程序设计语言中编写代码时定义一些可变部分，那些部分在使用前必须做出声明。

我们在上一小节已经了解到，Java 中的集合类是支持泛型的，它在代码中是这个样子的：

![](//img.mukewang.com/wiki/5ecdb9d20944b4aa12440817.jpg)

代码中的`<Integer>`就是泛型，我们**把类型像参数一样传递**，尖括号中间就是数据类型，我们可以称之为**实际类型参数**，这里实际类型参数的数据类型只能为引用数据类型。

那么为什么需要泛型呢？我们马上就见分晓。

## 2. 为什么需要泛型

上一节中，我们在使用`ArrayList`实现类的时候，如果没有指定泛型，`IDEA`会给出警告，代码似乎也是可以顺利运行的。请看如下实例：

```java
import java.util.ArrayList;

public class GenericsDemo1 {

    public static void main(String[] args) {
        ArrayList arrayList = new ArrayList();
        arrayList.add("Hello");
        String str = (String) arrayList.get(0);
        System.out.println("str=" + str);
    }

}
```

运行结果：

```java
str=Hello
```

虽然运行时没有发生任何异常，但这样做有两个缺点：

1. **需要强制类型转换**： 由于`ArrayList`内部就是一个`Object[]`数组，在`get()`元素的时候，返回的是`Object`类型，所以在`ArrayList`外获取该对象，需要强制类型转换。其它的`Collection`、`Map`如果不使用泛型，也存在这个问题；
2. 可向集合中添加任意类型的对象，存在类型不安全风险。例如如下代码中，我们向列表中既添加了`Integer`类型，又添加了`String`类型：

```java
import java.util.ArrayList;

public class GenericsDemo2 {
    public static void main(String[] args) {
        ArrayList arrayList = new ArrayList();
        arrayList.add(123);
        arrayList.add("Hello");
        String str = (String) arrayList.get(0);
        System.out.println("element=" + str);
    }
}
```

运行结果：

```java
Exception in thread "main" java.lang.ClassCastException: class java.lang.Integer cannot be cast to class java.lang.String (java.lang.Integer and java.lang.String are in module java.base of loader 'bootstrap')
	at GenericsDemo2.main(GenericsDemo2.java:8)
```

由于我们的“疏忽”，列表第 1 个元素实际上是整型，但被我们强制转换为字符串类型，这是行不通的，因此会抛出`ClassCastException`异常。

使用泛型可以解决这些问题。泛型有如下优点：

1. 可以减少类型转换的次数，代码更加简洁；
2. 程序更加健壮：只要编译期没有警告，运行期就不会抛出`ClassCastException`异常；
3. 提高了代码的可读性：编写集合的时候，就限定了集合中能存放的类型。

## 3. 如何使用泛型

### 3.1 泛型使用

在代码中，这样使用泛型：

```java
List<String> list = new ArrayList<String>();
// Java 7 及以后的版本中，构造方法中可以省略泛型类型：
List<String> list = new ArrayList<>();
```

要注意的是，变量声明的类型必须与传递给实际对象的类型保持一致，下面是错误的例子：

```java
List<Object> list = new ArrayList<String>();
List<Number> numbers = new ArrayList(Integer);
```

### 3.2 自定义泛型类

#### 3.2.1 Java 源码中泛型的定义

在自定义泛型类之前，我们来看下`java.util.ArrayList`是如何定义的：

![](//img.mukewang.com/wiki/5ecdba2f09dcc5f216380564.jpg)

类名后面的`<E>`就是泛型的定义，`E`不是 Java 中的一个具体的类型，它是 Java 泛型的通配符（注意是大写的，实际上就是`Element`的含义），可将其理解为一个占位符，**将其定义在类上，使用时才确定类型**。此处的命名不受限制，但最好有一定含义，例如`java.lang.HashMap`的泛型定义为`HashMap<K,V>`，`K`表示`Key`，`V`表示`Value`。

#### 3.2.2 自定义泛型类实例 1

下面我们来自定义一个泛型类，自定义泛型按照约定俗成可以叫`<T>`，具有`Type`的含义，实例如下：

```java
public class NumberGeneric<T> { // 把泛型定义在类上

    private T number; // 定义在类上的泛型，在类内部可以使用

    public T getNumber() {
        return number;
    }

    public void setNumber(T number) {
        this.number = number;
    }

    public static void main(String[] args) {
        // 实例化对象，指定元素类型为整型
        NumberGeneric<Integer> integerNumberGeneric = new NumberGeneric<>();
        // 分别调用set、get方法
        integerNumberGeneric.setNumber(123);
        System.out.println("integerNumber=" + integerNumberGeneric.getNumber());

        // 实例化对象，指定元素类型为长整型
        NumberGeneric<Long> longNumberGeneric = new NumberGeneric<>();
        // 分别调用set、get方法
        longNumberGeneric.setNumber(20L);
        System.out.println("longNumber=" + longNumberGeneric.getNumber());

        // 实例化对象，指定元素类型为双精度浮点型
        NumberGeneric<Double> doubleNumberGeneric = new NumberGeneric<>();
        // 分别调用set、get方法
        doubleNumberGeneric.setNumber(4000.0);
        System.out.println("doubleNumber=" + doubleNumberGeneric.getNumber());
    }

}
```

运行结果：

```java
integerNumber=123
longNumber=20
doubleNumber=4000.0
```

我们在类的定义处也定义了泛型：`NumberGeneric<T>`；在类内部定义了一个`T`类型的`number`变量，并且为其添加了`setter`和`getter`方法。

对于泛型类的使用也很简单，在主方法中，创建对象的时候指定`T`的类型分别为`Integer`、`Long`、`Double`，类就可以自动转换成对应的类型了。

#### 3.2.3 自定义泛型类实例 2

上面我们知道了如何定义含有单个泛型的类，那么对于含有多个泛型的类，如何定义呢？

我们可以看一下`HashMap`类是如何定义的。如下是 Java 源码的截图：

![](//img.mukewang.com/wiki/5ecdba5309fe155010160388.jpg)

参照`HashMap<K,V>`类的定义，下面我们来看看如何定义含有两个泛型的类，实例如下：

```java
public class KeyValueGeneric<K,V> { // 把两个泛型K、V定义在类上

    /**
     * 类型为K的key属性
     */
    private K key;

    /**
     * 类型为V的value属性
     */
    private V value;

    public K getKey() {
        return key;
    }

    public void setKey(K key) {
        this.key = key;
    }

    public V getValue() {
        return value;
    }

    public void setValue(V value) {
        this.value = value;
    }

    public static void main(String[] args) {
        // 实例化对象，分别指定元素类型为整型、长整型
        KeyValueGeneric<Integer, Long> integerLongKeyValueGeneric = new KeyValueGeneric<>();
        // 调用setter、getter方法
        integerLongKeyValueGeneric.setKey(200);
        integerLongKeyValueGeneric.setValue(300L);
        System.out.println("key=" + integerLongKeyValueGeneric.getKey());
        System.out.println("value=" + integerLongKeyValueGeneric.getValue());

        // 实例化对象，分别指定元素类型为浮点型、字符串类型
        KeyValueGeneric<Float, String> floatStringKeyValueGeneric = new KeyValueGeneric<>();
        // 调用setter、getter方法
        floatStringKeyValueGeneric.setKey(0.5f);
        floatStringKeyValueGeneric.setValue("零点五");
        System.out.println("key=" + floatStringKeyValueGeneric.getKey());
        System.out.println("value=" + floatStringKeyValueGeneric.getValue());
    }
}
```

运行结果：

```java
key=200
value=300
key=0.5
value=零点五
```

### 3.3 自定义泛型方法

前面我们知道了如何定义泛型类，在类上定义的泛型，在方法中也可以使用。下面我们来看一下如何自定义泛型方法。

泛型方法不一定写在泛型类当中。当类的调用者总是关心类中的某个泛型方法，不关心其他属性，这个时候就没必要再整个类上定义泛型了。

请查看如下实例：

```java
public class GenericMethod {

    /**
     * 泛型方法show
     * @param t 要打印的参数
     * @param <T> T
     */
    public <T> void show(T t) {
        System.out.println(t);
    }

    public static void main(String[] args) {
        // 实例化对象
        GenericMethod genericMethod = new GenericMethod();
        // 调用泛型方法show，传入不同类型的参数
        genericMethod.show("Java");
        genericMethod.show(222);
        genericMethod.show(222.0);
        genericMethod.show(222L);
    }
}
```

运行结果：

```java
Java
222
222.0
222
```

实例中，使用`<T>`来定义`show`方法的泛型，它接收一个泛型的参数变量并在方法体打印；调用泛型方法也很简单，在主方法中实例化对象，调用对象下的泛型方法，可传入不同类型的参数。

## 4. 泛型类的子类

泛型类也是一个 Java 类，它也具有继承的特性。

泛型类的继承可分为两种情况：

1. 子类明确泛型类的类型参数变量；
2. 子类不明确泛型类的类型参数变量。

下面我们来分别看一下这两种情况。

### 4.1 明确类型参数变量

例如，有一个泛型接口：

```java
public interface GenericInterface<T> { // 在接口上定义泛型
    void show(T t);
}
```

泛型接口的实现类如下：

```java
public class GenericInterfaceImpl implements GenericInterface<String> { // 明确泛型类型为String类型
    @Override
    public void show(String s) {
        System.out.println(s);
    }
}
```

子类实现明确了泛型的参数变量为`String`类型。因此方法`show()`的重写也将`T`替换为了`String`类型。

### 4.2 不明确类型参数变量

当实现类不确定泛型类的参数变量时，实现类需要定义类型参数变量，调用者使用子类时，也需要传递类型参数变量。

如下是`GenericInterface`接口的另一个实现类：

```java
public class GenericInterfaceImpl1<T> implements GenericInterface<T> { // 实现类也需要定义泛型参数变量
    @Override
    public void show(T t) {
        System.out.println(t);
    }
}
```

在主方法中调用实现类的`show()`方法：

```java
    public static void main(String[] args) {
        GenericInterfaceImpl1<Float> floatGenericInterfaceImpl1 = new GenericInterfaceImpl1<>();
        floatGenericInterfaceImpl1.show(100.1f);
    }
```

## 5. 类型通配符

我们先来看一个泛型作为方法参数的实例：

```java
import java.util.ArrayList;
import java.util.List;

public class GenericDemo3 {
    /**
     * 遍历并打印集合中的每一个元素
     * @param list 要接收的集合
     */
    public void printListElement(List<Object> list) {
        for (Object o : list) {
            System.out.println(o);
        }
    }
}
```

观察上面的代码，参数`list`的限定的泛型类型为`Object`， 也就是说，这个方法只能接收元素为`Object`类型的集合，如果我们想传递其他元素类型的集合，是行不通的。例如，如果传递装载`Integer`元素的集合，程序在编译阶段就会报错：

![](//img.mukewang.com/wiki/5ecdba7e091a2c2613351307.jpg)

> **Tips：** 泛型中的`List<Object>`并不是`List<Integer>`的父类，它们不满足继承关系。

### 5.1 无限定通配符

想要解决这个问题，使用类型通配符即可，修改方法参数处的代码，将`<>`中间的 Object 改为`?`即可：

```java
public void printListElement(List<?> list) {
```

此处的`?`就是类型通配符，表示可以匹配任意类型，因此调用方可以传递任意泛型类型的列表。

完整实例如下：

```java
import java.util.ArrayList;
import java.util.List;

public class GenericDemo3 {
    /**
     * 遍历并打印集合中的每一个元素
     * @param list 要接收的集合
     */
    public void printListElement(List<?> list) {
        for (Object o : list) {
            System.out.println(o);
        }
    }

    public static void main(String[] args) {
        // 实例化一个整型的列表
        List<Integer> integers = new ArrayList<>();
        // 添加元素
        integers.add(1);
        integers.add(2);
        integers.add(3);
        GenericDemo3 genericDemo3 = new GenericDemo3();
        // 调用printListElement()方法
        genericDemo3.printListElement(integers);

        // 实例化一个字符串类型的列表
        List<String> strings = new ArrayList<>();
        // 添加元素
        strings.add("Hello");
        strings.add("慕课网");
        // 调用printListElement()方法
        genericDemo3.printListElement(strings);
    }
}
```

运行结果：

```java
1
2
3
Hello
慕课网
```

### 5.2 extends 通配符

`extends`通配符用来限定泛型的上限。什么意思呢？依旧以上面的实例为例，我们来看一个新的需求，我们希望方法接收的`List` 集合限定在数值类型内（float、integer、double、byte 等），不希望其他类型可以传入（比如字符串）。此时，可以改写上面的方法定义，设定上界通配符：

```java
public void printListElement(List<? extends Number> list) {
```

这样的写法的含义为：`List`集合装载的元素只能是`Number`自身或其子类（`Number`类型是所有数值类型的父类），完整实例如下：

```java
import java.util.ArrayList;
import java.util.List;

public class GenericDemo4 {
    /**
     * 遍历并打印集合中的每一个元素
     * @param list 要接收的集合
     */
    public void printListElement(List<? extends Number> list) {
        for (Object o : list) {
            System.out.println(o);
        }
    }

    public static void main(String[] args) {
        // 实例化一个整型的列表
        List<Integer> integers = new ArrayList<>();
        // 添加元素
        integers.add(1);
        integers.add(2);
        integers.add(3);
        GenericDemo4 genericDemo3 = new GenericDemo4();
        // 调用printListElement()方法
        genericDemo3.printListElement(integers);

    }
}
```

运行结果：

```java
1
2
3
```

### 5.3 super 通配符

既然已经了解了如何设定通配符上界，也就不难理解通配符的下界了，可以限定传递的参数只能是某个类型的父类。

语法如下：

```java
<? super Type>
```

## 6. 小结

通过本小节的学习，我们知道了使用泛型可以避免强制类型转换，也可以避免运行期就抛出的`ClassCastException`异常。在使用泛型时，要注意变量声明的泛型类型要匹配传递给实际对象的类型， Java 7 及以后的版本中，构造方法中可以省略泛型类型，推荐直接省略。

我们也学习了如何自定义泛型类和泛型方法，在实际的开发中，我们想要编写比较通用的代码就避免不了使用泛型，大家可以在以后的开发中慢慢体悟。

另外，泛型也是可以继承的。

最后，我们还讲解了类型通配符的概念和使用场景。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

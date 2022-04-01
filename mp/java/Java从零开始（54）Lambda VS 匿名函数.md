---
title: Java 从零开始（54）Lambda VS 匿名函数
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Lambda 表达式 VS 匿名内部类

本节我们将重点讨论下 Lambda 表达式与匿名内部类之间除了语法外还有哪些差别。再开始讲解之前我们先列出两者重要的两点区别：

* **无标识性：** 内部类会确保创建一个拥有唯一表示的对象，而 Lambda 表达式的计算结果有可能有唯一标识，也可能没有。
* **作用域规则：** 由于内部类可以从父类继承属性，Lambda 表达式却不能，所以，内部类的作用域规则比 Lambda 表达式要复杂。（关于 Lambda 表达式的作用域规则可以回看 03 节的内容）

我们来看一个例子：

```java
public class TestLambdaAndInnerClass  {
    public void test(){
        //匿名类实现
        Runnable innerRunnable = new Runnable(){
            @Override
            public void run() {
                System.out.println("call run in innerRunnable:\t"+this.getClass());
            }
        };
        //Lambda 表达式实现
        Runnable lambdaRunnable = () -> System.out.println("call run in lambdaRunnable:\t"+this.getClass());
        new Thread(innerRunnable).start();
        new Thread(lambdaRunnable).start();
    }

    public static void main(String...s){
        new TestLambdaAndInnerClass().test();
    }
}

```

返回结果为：

```java
call run in innerRunnable:  class com.github.x19990416.item.TestLambdaAndInnerClass$1
call run in lambdaRunnable: class com.github.x19990416.item.TestLambdaAndInnerClass
```

上面的例子分别在内部类和 Lambda 表达式中调用各自的 `this` 指针。我们发现 Lambda 表达式的 `this` 指针指向的是其外围类 `TestLambdaAndInnerClass`，匿名内部类的指针指向的是其本身。（对于 `super` 也是一样的结果）

我们来看下编译出来的文件：

​

![图片描述](https://xushuhui.gitee.io/image/imooc/5f0305b3083c531606150147.jpg)

​

其中 `TestLambdaAndInnerClass$1.class` 是匿名类 `innerRunnable` 编译出来的 class 文件，对于 Lambda 表达式 `lambdaRunnable` 则没有编译出具体的 class 文件。

这说明对于 Lambda 表达式而言，编译器并不认为它是一个完全的类（或者说它是一个特殊的类对象），所以也不具备一个完全类的特征。

> **Tips:** 匿名类的 `this`、`super` 指针指向的是其自身的实例，而 Lambda 表达式的 `this`、`super` 指针指向的是创建这个 Lambda 表达式的类对象的实例。

## 1. 无标识性问题

在 Lambda 表达式出现之前，一个 Java 程序的行为总是与对象关联，以标识、状态和性为为特征。然而 Lambda 表达式则违背了这个规则。

虽然 Lambda 表达式可以共享对象的一些属性，但是表示行为是其唯一的用处。由于没有状态，所以表示问题也就不那么重要了。在 Java 语言的规范中对 Lambda 表达式唯一的要求就是必须计算出其实现的相当的函数接口的实例类。如果 Java 对每个 Lambda 表达式都拥有唯一的表示，那么 Java 就没有足够的灵活性来对系统进行优化。

## 2. Lambda 表达式的作用域规则

匿名内部类与大多数类一样，由于它可以引用从父类继承下来的名字，以及声明在外部类中的名字，所以它的作用域规则非常复杂。

Lambda 表达式由于不会从父类型中继承名字，所以它的作用于规则要简单很多。除了参数以外，用在 Lambda 表达式中的名字的含义与表达式外面是一样的。

由于 Lambda 表达式的声明就是一个语句块，所以 `this` 与 `super`与表达式外围环境的含义一样，换言之它们指向的是外围对象的父类对象。

## 3. 小结

本节我们主要讨论了 Lambda 表达式与匿名内部类的本质区别，其中重要的是要记住 `this` 和 `super` 在两者之间的作用范围。记住这个作用范围可以更好的帮助我们理解 Lambda 表达式的作用域，避免我们在使用 Lambda 表达式中由于作用域引起的 bug，这一类的 bug 在实际中定位是非常困难的。

![](https://xushuhui.gitee.io/image/imooc/5f1a901009e3945909000228.jpg)

​

在 Java 里面，所有的方法参数都是有固定类型的，比如将数字 9 作为参数传递给一个方法，它的类型是 int；字符串 “9” 作为参数传递给方法，它的类型是 String。那么 Lambda 表达式的类型由是什么呢？通过本节我们学习什么是函数式接口，它与 Lambda 表达式的关系。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

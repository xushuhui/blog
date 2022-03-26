---
title: Java 从零开始（45）Optional 类
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
zhihu-url: https://zhuanlan.zhihu.com/p/415815025
---
# Optional 类

上一小节，我们接触到了`Optional`类，但没有详细展开介绍，`Optional`类也是 Java 8 新加入的类。本小节我们就来学习一下这个类，你将了解到`Optional`类的解决了什么问题，如何创建`Optioanl`类的对象，它又有哪些常用方法，如何在实际开发中应用`Optional`类等内容。

## 1. Optional 类概述

空指针异常（`NullPointerExceptions`）是 Java 最常见的异常之一，一直以来都困扰着 Java 程序员。一方面，程序员不得不在代码中写很多`null`的检查逻辑，让代码看起来非常臃肿；另一方面，由于其属于运行时异常，是非常难以预判的。

为了预防空指针异常，`Google`的`Guava`项目率先引入了`Optional`类，通过使用检查空值的方式来防止代码污染，受到`Guava`项目的启发，随后在`Java 8`中也引入了`Optional`类。

Optional 类位于`java.util`包下，是一个可以为 `null` 的**容器对象**，如果值存在则`isPresent()`方法会返回 `true` ，调用 `get()` 方法会返回该对象，可以有效避免空指针异常。下面我们来学习如何实例化这个类，以及这个类下提供了哪些常用方法。

## 2. 创建 Optional 对象

查看 `java.util.Optional`类源码，可以发现其构造方法是私有的，因此不能通过`new`关键字来实例化：

![](https://xushuhui.gitee.io/image/imooc/5f113820095ec00414660514.jpg)

我们可以通过如下几种方法，来创建 Optional 对象：

* `Optional.of(T t)`：创建一个 Optional 对象，参数 `t` 必须非空；
* `Optional.empty()`：创建一个空的`Optional`实例；
* `Optional.ofNullable(T t)`：创建一个`Optional`对象，参数`t` 可以为 `null`。

实例如下：

```java
import java.util.Optional;

public class OptionalDemo1 {

    public static void main(String[] args) {
        // 创建一个 StringBuilder 对象
        StringBuilder string = new StringBuilder("我是一个字符串");

        // 使用 Optional.of(T t) 方法，创建 Optional 对象，注意 T 不能为空：
        Optional<StringBuilder> stringBuilderOptional = Optional.of(string);
        System.out.println(stringBuilderOptional);

        // 使用 Optional.empty() 方法，创建一个空的 Optional 对象：
        Optional<Object> empty = Optional.empty();
        System.out.println(empty);

        // 使用 Optional.ofNullable(T t) 方法，创建 Optional 对象，注意 t 允许为空：
        stringBuilderOptional = null;
        Optional<Optional<StringBuilder>> stringBuilderOptional1 = Optional.ofNullable(stringBuilderOptional);
        System.out.println(stringBuilderOptional1);
    }

}
```

运行结果：

```java
Optional[我是一个字符串]
Optional.empty
```

## 3. 常用方法

`Optional<T>`类提供了如下常用方法：

* `booean isPresent()`：判断是否包换对象；
* `void ifPresent(Consumer<? super T> consumer)`：如果有值，就执行 Consumer 接口的实现代码，并且该值会作为参数传递给它；
* `T get()`：如果调用对象包含值，返回该值，否则抛出异常；
* `T orElse(T other)`：如果有值则将其返回，否则返回指定的`other` 对象；
* `T orElseGet(Supplier<? extends T other>)`：如果有值则将其返回，否则返回由`Supplier`接口实现提供的对象；
* `T orElseThrow(Supplier<? extends X> exceptionSupplier)`：如果有值则将其返回，否则抛出由`Supplier`接口实现提供的异常。

知道了如何创建`Optional`对象和常用方法，我们下面结合具体实例来看一下，`Optional`类是如何避免空指针异常的。

请查看如下实例，其在运行时会发生空指针异常：

```java
import java.util.Optional;

public class OptionalDemo2 {

    static class Category {
        private String name;

        public Category(String name) {
            this.name = name;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        @Override
        public String toString() {
            return "Category{" +
                    "name='" + name + '\'' +
                    '}';
        }
    }

    static class Goods {
        private String name;

        private Category category;

        public Goods() {

        }

        public Goods(String name, Category category) {
            this.name = name;
            this.category = category;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public Category getCategory() {
            return category;
        }

        public void setCategory(Category category) {
            this.category = category;
        }

        @Override
        public String toString() {
            return "Good{" +
                    "name='" + name + '\'' +
                    ", category=" + category +
                    '}';
        }
    }

    /**
     * 获取商品的分类名称
     * @param goods 商品
     * @return 分类名称
     */
    static String getGoodsCategoryName(Goods goods) {
        return goods.getCategory().getName();
    }

    public static void main(String[] args) {
        // 实例化一个商品类
        Goods goods = new Goods();
        // 获取商品的分类名称
        String categoryName = getGoodsCategoryName(goods);
        System.out.println(categoryName);
    }
}
```

运行结果：

```java
Exception in thread "main" java.lang.NullPointerException
	at OptionalDemo2.getGoodsCategoryName(OptionalDemo2.java:73)
	at OptionalDemo2.main(OptionalDemo2.java:80)
```

实例中，由于在实例化`Goods`类时，我们没有给其下面的`Category`类型的属性`category`赋值，它就为 `null`，在运行时， `null.getName()`就会抛出空指针异常。同理，如果`goods`实例为`null`，那么`null.getCategory()`也会抛出空指针异常。

在没有使用`Optional`类的情况下，想要优化代码，就不得不改写`getGoodsCategoryName()`方法：

```java
static String getGoodsCategoryName(Goods goods) {
    if (goods != null) {
        Category category = goods.getCategory();
        if (category != null) {
            return category.getName();
        }
    }
    return "该商品无分类";
}
```

这也就是我们上面说的`null`检查逻辑代码，此处有两层`if`嵌套，如果有更深层次的级联属性，就要嵌套更多的层级。

下面我们将`Optional`类引入实例代码：

```java
import java.util.Optional;

public class OptionalDemo3 {

    static class Category {
        private String name;

        public Category(String name) {
            this.name = name;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        @Override
        public String toString() {
            return "Category{" +
                    "name='" + name + '\'' +
                    '}';
        }
    }

    static class Goods {
        private String name;

        private Category category;

        public Goods() {

        }

        public Goods(String name, Category category) {
            this.name = name;
            this.category = category;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public Category getCategory() {
            return category;
        }

        public void setCategory(Category category) {
            this.category = category;
        }

        @Override
        public String toString() {
            return "Good{" +
                    "name='" + name + '\'' +
                    ", category=" + category +
                    '}';
        }
    }

    /**
     * 获取商品的分类名称（使用 Optional 类包装）
     * @param goods 商品
     * @return 分类名称
     */
    static String getGoodsCategoryName(Goods goods) {
        // 将商品实例包装入 Optional 类，创建 Optional<Goods> 对象
        Optional<Goods> goodsOptional = Optional.ofNullable(goods);
        Goods goods1 = goodsOptional.orElse(new Goods("默认商品", new Category("默认分类")));
        // 此时 goods1 一定是非空，不会产生空指针异常
        Category category = goods1.getCategory();

        // 将分类实例包装入 Optional 类，创建 Optional<Category> 对象
        Optional<Category> categoryOptional = Optional.ofNullable(category);
        Category category1 = categoryOptional.orElse(new Category("默认分类"));
        // 此时 category1 一定是非空，不会产生空指针异常
        return category1.getName();
    }

    public static void main(String[] args) {
        // 实例化一个商品类
        Goods goods = null;
        // 获取商品的分类名称
        String categoryName = getGoodsCategoryName(goods);
        System.out.println(categoryName);
    }
}
```

运行结果：

```java
默认分类
```

实例中，我们使用`Optional`类的 `ofNullable（T t）`方法分别包装了`goods`对象及其级联属性`category`对象，允许对象为空，然后又调用了其`ofElse(T t)`方法保证了对象一定非空。这样，空指针异常就被我们优雅地规避掉了。

## 4. 对于空指针异常的改进

Java 14 对于空指针异常有了一些改进，它提供了更明确异常堆栈打印信息，JVM 将精确地确定那个变量是`null`，不过空指针异常依然无法避免。明确的异常堆栈信息，能够帮助开发者快速定位错误发生的位置。

## 5. 小结

通过本小节的学习，我们知道了 `Optional` 类主要用于应对 Java 中的空指针异常，它是一个可以为 `null` 的容器对象，我们可以通过`Optional`类下的几个静态方法来创建对象。另外，我们也结合实例介绍了如何使用`Optional`类来规避空指针异常，实例中还有很多其他没用到的 API，希望大家可以自己研习。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

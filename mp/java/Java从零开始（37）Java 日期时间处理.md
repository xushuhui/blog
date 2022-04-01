---
title: Java 从零开始（37）Java 日期时间处理
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
zhihu-url: https://zhuanlan.zhihu.com/p/415673706
---
# Java 日期和时间

本小节我们将学习 Java 中的日期和时间，日期和时间在我们的实际开发中非常常用，例如用户的注册、数据的增删改、对敏感信息的操作等等都需要记录下日期和时间。通过本小节的学习，你将了解到什么是日期、什么是时间、什么是时区，Java 中 `Date` 类的 API 介绍，`Calendar` 日历类的使用，`LocalDateTime` 类的相关 API 介绍等内容。

## 1. 什么是日期和时间

日期指的是某一天，例如：

* **2020-10-24**：2020 年 10 月 24 日；
* **1998-6-14**：1998 年 6 月 14 日。

时间就是指某一个时刻，它分为两种，一种是带日期的时间，另外一种是不带日期的时间，例如：

* **2020-10-24 08:30:23**：2020 年 10 月 24 日 8 时 30 分 23 秒；
* **11:22:33**：11 时 22 分 33 秒。

## 2. 什么是时区

我们知道，地球上的不同地区是有时差的，因此想要准确定位一个时刻，还需要加上时区。

时区有以下 3 种表示方式：

1. **`GMT` 或者 `UTC` 加时区偏移表示**：例如：`GMT+08:00` 或者 `UTC+08:00` 表示东八区；
2. **缩写表示**：例如：`CST` 表示 `China Standard Time`（中国标准时间），但是此缩写也可表示 `Central Standard Time USA`（美国中部时间），容易混淆不推荐使用；
3. **洲 / 城市**：例如：`Asia/Shanghai` 表示上海所在地的时区。注意城市名称不是任意的城市，而是由国际标准组织规定的城市。

## 3. Date 类的使用

`java.util.Date` 类日期表示特定的时间瞬间，精度为毫秒。我们下面来看一下这个类的构造方法和常用方法。

### 3.1 构造方法

* `Date()`：创建一个对应当前时间的日期对象；
* `Date(long date)`：创建指定毫秒数的日期对象。

由于其他 4 个构造方法已经过期，这里我们不进行介绍。

### 3.2 常用方法

* `String toString()`：将此日期对象转换为以下形式的字符串：星期 月 日 时：分：秒 时区 年；
* `long getTime()`：返回此日期对象表示的自 1970 年 1 月 1 日 00:00:00 GMT 以来的毫秒数；
* `void setTime()`：将此日期对象设置为表示 1970 年 1 月 1 日 00:00:00 GMT 之后的时间点（毫秒）。

大多数其他方法都已经过期，此处不再一一列举，可翻阅 [官方文档](https://docs.oracle.com/en/java/javase/13/docs/api/java.base/java/util/Date.html) 以了解更多内容。

> **Tips**：除了 `java.util` 包下的 `Date` 类，在 `java.sql` 包下也有一个 `Date` 类。它是对应数据库字段的日期类型的类，与数据库交互的时候才会用到，由于目前我们不涉及数据库相关知识，此处做一个了解即可。我们更常用的还是 `java.util` 包下的 `Date` 类。

### 3.3 实例

创建日期对象，并打印 `toString()` 方法的结果：

```java
import java.util.Date;

/**
 * @author colorful@TaleLin
 */
public class DateDemo1 {

    public static void main(String[] args) {
        // 实例化一个 date 对象
        Date date = new Date();
      	// 调用 toString() 方法
        String s = date.toString();
        System.out.println(s);
    }

}
```

运行结果：

```java
Wed Jun 10 10:21:10 CST 2020
```

我执行代码的时间是 2020 年 06 月 10 日 10:22:10 星期三，因此得到了如上的运行结果。

调用 `getTime()` 方法获取当前日期对象自 1970 年 1 月 1 日 00:00:00 GMT 以来的毫秒数：

```java
import java.util.Date;

/**
 * @author colorful@TaleLin
 */
public class DateDemo2 {

    public static void main(String[] args) {
        // 实例化一个 date 对象
        Date date = new Date();
        // 调用 getTime() 方法
        long time = date.getTime();
        System.out.println(time);
    }

}
```

运行结果：

```java
1591755946922
```

## 4. Calendar 类的使用

Calendar 类是一个抽象类，它提供了一些方法，用于在特定的时间瞬间与一组日历字段（如年、月、月、日、小时等）之间进行转换，以及用于处理日历字段（如获取下一周的日期）。

### 4.1 实例化

由于 `Calendar` 类是一个抽象类，不能直接实例化，想要获取其实例对象通常有两种方法：

1. 使用 `Calendar.getInstance()` 方法（更常用）；
2. 调用它的子类的 `GregorianCalendar` 的构造方法。

实例如下：

```java
Calendar calendar = Calendar.getInstance();
```

### 4.3 常用方法

* `static Calendar getInstance()`：使用默认时区和区域设置获取日历；
* `int get(int field)`：返回给定日历字段的值；
* `void set(int field, int value)`：将给定的日历字段设置为给定值。（此外，`set()` 还有很多重载方法）

`get(int field)` 和 `set(int field, int value)` 方法的 `field` 参数是 `int` 类型，我们可以使用 `Calendar` 类下的一些静态字段来表示，如下是官方文档关于 `Calendar` 类的部分静态字段的截图：

![](https://xushuhui.gitee.io/image/imooc/5ee07fcd09808dc011910735.jpg)

更多常用方法和静态字段请查阅 [官方文档](https://docs.oracle.com/en/java/javase/13/docs/api/java.base/java/util/Calendar.html)。

### 4.3 实例

#### 4.3.1 get () 方法实例

```java
import java.util.Calendar;

/**
 * @author colorful@TaleLin
 */
public class CalendarDemo1 {

    public static void main(String[] args) {
        Calendar calendar = Calendar.getInstance();
        // 获取今天是这一周的第几天
        int i = calendar.get(Calendar.DAY_OF_WEEK);
        // 获取今天是这一月的第几天
        int i1 = calendar.get(Calendar.DAY_OF_MONTH);
        // 获取今天是这一月的第几周
        int i2 = calendar.get(Calendar.DAY_OF_WEEK_IN_MONTH);
        // 获取今天是这一年的第几天
        int i3 = calendar.get(Calendar.DAY_OF_YEAR);
        System.out.println("今天是这一周的第" + i + "天");
        System.out.println("今天是这一月的第" + i1 + "天");
        System.out.println("今天是这一月的第" + i2 + "周");
        System.out.println("今天是这一年的第" + i3 + "天");
    }
}
```

运行此段代码的时间是 2020 年 06 月 10 日，运行结果如下：

```java
今天是这一周的第 4 天
今天是这一月的第 10 天
今天是这一月的第 2 周
今天是这一年的第 162 天
```

通过调用 `get()` 方法，我们很方便地获取到了当前时间在日历上是第几天。要特别注意的是，获取月份，返回的值是从 0 开始的（0 ~ 11），依次表示 1 月到 12 月；获取一周的第 `n` 天，这里的返回值为 `1~7`，1 表示周日，2 表示周一，以此类推。

#### 4.3.2 set () 方法实例

```java
import java.util.Calendar;

/**
 * @author colorful@TaleLin
 */
public class CalendarDemo2 {

    public static void main(String[] args) {
        Calendar calendar = Calendar.getInstance();
        // 设置 2022 年：
        calendar.set(Calendar.YEAR, 2022);
        // 设置 9 月：(8 表示 9 月）
        calendar.set(Calendar.MONTH, 8);
        // 设置 9 日：
        calendar.set(Calendar.DATE, 9);
        // 设置时间：
        calendar.set(Calendar.HOUR_OF_DAY, 0);
        calendar.set(Calendar.MINUTE, 0);
        calendar.set(Calendar.SECOND, 0);
        System.out.println(calendar.getTime().toString());
    }
}
```

运行结果：

```java
Fri Sep 09 00:00:00 CST 2022
```

上面代码分别设置了日历的年、月、日、时、分、秒，`Calendar.getTime()` 可以将一个 `Calendar` 对象转换成 `Date` 对象，最后再打印这个对象。

## 5. Java8 后新的日期和时间 API

### 5.1 为什么要提供新的 API

在 Java8 以后，提供了新的日期和时间 API，旧 API 的如下缺点得到了解决：

* 可变性：旧的 API 日期和时间是可变的，日期和时间这样的类应该是不可变的；
* 偏移性：`Date` 中的年份是从 1990 年开始，月份是从 0 开始，星期天是用 1 表示的，不了解 API 的开发者很容易用错；
* 格式化：`SimpleDateFormat` 只能用于格式化 `Date` 类型，不能格式化 `Calendar` 类型。

### 5.2 新 API 概述

新的日期时间 API 吸取了 [Joda-Time](https://www.joda.org/joda-time/) 的精华，提供了更优秀易用的 API，位于 `java.time` 包中，包含了 `LocalTime`（本地时间）、`LocalDate`（本地日期）、`LocalDateTime`（本地日期时间）、`ZonedDateTime`（带时区的日期时间）和 `Duration`（时间间隔）类。

而 `java.util.Date` 类下面增加了 `toInstant()` 方法，用于把 `Date` 转换为新的类型。这些 API 大大简化了日期时间的运算。

对偏移性的不合理设计也有修正：月份使用 `1~12` 表示 1 月 到 12 月，星期使用 `1 ~ 7` 表示星期一到星期天。

另外，使用了新的 `DateTimeFormatter` 来取代旧的 `SimpleDateFormat`。

### 5.3 LocalDateTime 相关类的使用

```java
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

/**
 * @author colorful@TaleLin
 */
public class LocalDateTimeDemo1 {

    public static void main(String[] args) {
        // 获取当前日期
        LocalDate localDate = LocalDate.now();
        // 获取当前时间
        LocalTime localTime = LocalTime.now();
        // 获取当前日期和时间
        LocalDateTime localDateTime = LocalDateTime.now();
        // 打印
        System.out.println(localDate);
        System.out.println(localTime);
        System.out.println(localDateTime);
    }

}
```

运行结果：

```java
2020-06-10
14:17:48.618294
2020-06-10T14:17:48.618421
```

在实际开发中，`LocalDateTime` 相较于 `LocalDate` 和 `LocalTime` 更加常用。

本地日期和时间通过 `now()` 获取到的总是以当前默认时区返回的。

另外，可以使用 `of()` 方法来设置当前日期和时间：

```java
// 2020-9-30
LocalDate date = LocalDate.of(2020, 9, 30);
// 14:15:10
LocalTime time = LocalTime.of(14, 15, 10);
// 将 date 和 time 组合成一个 LocalDateTime
LocalDateTime dateTime1 = LocalDateTime.of(date, time);
// 设置 年、月、日、时、分、秒
LocalDateTime dateTime2 = LocalDateTime.of(2020, 10, 21, 14, 14);
```

## 6. 小结

通过本小节的学习，我们知道了日期、时间和时区的基本概念，在 Java 8 之前，通过 `Date` 类、`Calendar` 类以及 `SimpleDateFormat` 类来操作日期和时间，Java 8 以后，`java.time` 包下新增了一批新的日期时间 API，修复了旧 API 的一些缺点，简化了开发者操作日期和时间的难度。本小节内容比较简单，梳理了 Java 对于日期时间处理的相关 API，大家参照文档，加以练习即可。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

# Java 流式操作



流式操作，是 `Java 8` 除了`Lambda`表达式外的又一重大改变。学习流式操作，就是学习`java.util.stream`包下的`API`，我们称之为`Stream API`，它把真正的函数式编程引入到了 Java 中。



本小节我们将了解到**什么是`Stream`**，**为什么使用`Stream API`**， **流式操作的执行流程**，**如何实例化`Stream`**，**`Stream`的中间操作**、**`Stream`的终止操作**等内容。



## 1. 什么是 Stream



`Stream`是数据渠道，用于操作数据源所生成的元素序列，它可以实现对集合（`Collection`）的复杂操作，例如查找、替换、过滤和映射数据等操作。



我们这里说的`Stream`不同于`java`的输入输出流。另外，Collection 是一种静态的**数据结构**，存储在内存中，而`Stream`是用于计算的，通过`CPU`实现计算。注意不要混淆。



> **Tips**：`Stream`自己不会存储数据；`Stream`不会改变源对象，而是返回一个新的持有结果的`Stream`（不可变性）；`Stream`操作是延迟执行的（这一点将在后面介绍）。



## 2. 为什么使用 Stream API



我们在实际开发中，项目中的很多数据都来源于关系型数据库（例如 MySQL、Oracle 数据库），我们使用`SQL`的条件语句就可以实现对数据的筛选、过滤等等操作；



但也有很多数据来源于非关系型数据库（`Redis`、`MongoDB`等），想要处理这些数据，往往需要在 Java 层面去处理。



使用`Stream API`对集合中的数据进行操作，就类似于 SQL 执行的数据库查询。也可以使用`Stream API`来执行并行操作。简单来说，`Stream API`提供了一种高效且易于使用的处理数据的方式。



## 3. 流式操作的执行流程



流式操作通常分为以下 3 个步骤：



1. **创建`Stream`对象**：通过一个数据源（例如集合、数组），获取一个流；


2. **中间操作**：一个中间的链式操作，对数据源的数据进行处理（例如过滤、排序等），直到执行终止操作才执行；


3. **终止操作**：一旦执行终止操作，就执行中间的链式操作，并产生结果。


下图展示了`Stream`的执行流程：



![](//img.mukewang.com/wiki/5f113d0809d64d5f21990852.jpg)

接下来我们就按照这 3 个步骤的顺序来展开学习`Stream API`。



## 4. Stream 对象的创建



有 4 种方式来创建`Stream`对象。



### 4.1 通过集合创建 Stream



Java 8 的`java.util.Collection` 接口被扩展，提供了两个获取流的默认方法：



* `default Stream<E> stream()`：返回一个串行流（顺序流）；


* `default Stream<E> parallelStream()`：返回一个并行流。


实例如下：



```java
// 创建一个集合，并添加几个元素  
List<String> stringList = new ArrayList<>();  
stringList.add("hello");  
stringList.add("world");  
stringList.add("java");  
​  
// 通过集合获取串行 stream 对象  
Stream<String> stream = stringList.stream();  
// 通过集合获取并行 stream 对象  
Stream<String> personStream = stringList.parallelStream();
```



串行流并行流的区别是：串行流从集合中取数据是按照集合的顺序的；而并行流是并行操作的，获取到的数据是无序的。



### 4.2 通过数组创建 Stream



Java 8 中的`java.util.Arrays`的静态方法`stream()`可以获取数组流：



* `static <T> Stream<T> stream(T[] array)`：返回一个数组流。


此外，`stream()`还有几个重载方法，能够处理对应的基本数据类型的数组：



* `public static IntStream stream(int[] array)`：返回以指定数组作为其源的连续`IntStream`；


* `public static LongStream stream(long[] array)`：返回以指定数组作为其源的连续`LongStream`；


* `public static DoubleStream stream(double[] array)`：返回以指定数组作为其源的连续`DoubleStream`。


实例如下：



```java
import java.util.Arrays;
import java.util.stream.IntStream;
import java.util.stream.Stream;

public class StreamDemo1 {

    public static void main(String[] args) {
        // 初始化一个整型数组
        int[] arr = new int[]{1,2,3};
        // 通过整型数组，获取整形的 stream 对象
        IntStream stream1 = Arrays.stream(arr);

        // 通过字符串类型的数组，获取泛型类型为 String 的 stream 对象
        String[] stringArr = new String[]{"Hello", "imooc"};
        Stream<String> stream2 = Arrays.stream(stringArr);
    }
}
```



### 4.3 通过 Stream 的 `of()`方法



可以通过`Stream`类下的`of()`方法来创建 Stream 对象，实例如下：



```java
import java.util.stream.Stream;

public class StreamDemo1 {

    public static void main(String[] args) {
        // 通过 Stream 类下的 of() 方法，创建 stream 对象、
        Stream<Integer> stream = Stream.of(1, 2, 3);
    }
}
```



### 4.4 创建无限流



可以使用`Stream`类下的静态方法`iterate()`以及`generate()`创建无限流：



* `public static<T> Stream<T> iterate(final T seed, final UnaryOperator<T> f)`：遍历；


* `public static<T> Stream<T> generate(Supplier<T> s)`：生成。


创建无限流的这种方式实际使用较少，大家了解一下即可。



## 5. Stream 的中间操作



多个中间操作可以连接起来形成一个流水线，除非流水线上触发终止操作，否则中间操作不会执行任何的处理。在终止操作时会一次性全部处理这些中间操作，称为“惰性求值”。下面，我们来学习一下常用的中间操作方法。



### 5.1 筛选与切片



关于筛选和切片中间操作，有下面几个常用方法：



* `filter(Predicate p)`：接收 `Lambda`，从流中清除某些元素；


* `distinct()`：筛选，通过流生成元素的`hashCode`和`equals()`方法去除重复元素；


* `limit(long maxSize)`：截断流，使其元素不超过给定数量；


* `skip(long n)`：跳过元素，返回一个扔掉了前 `n` 个元素的流。若流中元素不足 `n` 个，则返回一个空流。与`limit(n)`互补。


我们先来看一个过滤集合元素的实例：








```java
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Stream;

public class StreamDemo2 {

    static class Person {
        private String name;
        private int age;

        public Person() { }

        public Person(String name, int age) {
            this.name = name;
            this.age = age;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public int getAge() {
            return age;
        }

        public void setAge(int age) {
            this.age = age;
        }

        @Override
        public String toString() {
            return "Person{" +
                    "name='" + name + '\'' +
                    ", age=" + age +
                    '}';
        }
    }

    /**
     * 创建一个Person的集合
     * @return List
     */
    public static List<Person> createPeople() {
        ArrayList<Person> people = new ArrayList<>();
        Person person1 = new Person("小明", 15);
        Person person2 = new Person("小芳", 20);
        Person person3 = new Person("小李", 18);
        Person person4 = new Person("小付", 23);
        Person person5 = new Person("大飞", 22);
        people.add(person1);
        people.add(person2);
        people.add(person3);
        people.add(person4);
        people.add(person5);
        return people;
    }

    public static void main(String[] args) {
        List<Person> people = createPeople();
        // 创建 Stream 对象
        Stream<Person> stream = people.stream();
        // 过滤年龄大于 20 的person
        Stream<Person> personStream = stream.filter(person -> person.getAge() >= 20);
        // 触发终止操作才能执行中间操作，遍历列表中元素并打印
        personStream.forEach(System.out::println);
    }
}
```







运行结果：



```java
Person{name='小芳', age=20}  
Person{name='小付', age=23}  
Person{name='大飞', age=22}
```



实例中，有一个静态内部类`Person`以及一个创建`Person`的集合的静态方法`createPeople()`，在主方法中，我们先调用该静态方法获取到一个`Person`列表，然后创建了`Stream`对象，再执行中间操作（即调用`fliter()`方法），这个方法的参数类型是一个**断言型的函数式接口**，接口下的抽象方法`test()`要求返回`boolean`结果，因此我们使用`Lambda`表达式，`Lambda`体为`person.getAge() >= 20`，其返回值就是一个布尔型结果，这样就实现了对年龄大于等于 20 的`person`对象的过滤。



由于必须触发终止操作才能执行中间操作，我们又调用了`forEach(System.out::println)`，在这里记住它作用是遍历该列表并打印每一个元素即可，我们下面将会讲解。另外，`filter()`等这些由于中间操作返回类型为 Stream，所以支持链式操作，我们可以将主方法中最后两行代码合并成一行：



```java
stream.filter(person -> person.getAge() >= 20).forEach(System.out::println);
```



我们再来看一个截断流的使用实例：








```java
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Stream;

public class StreamDemo3 {

    static class Person {
        private String name;
        private int age;

        public Person() { }

        public Person(String name, int age) {
            this.name = name;
            this.age = age;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public int getAge() {
            return age;
        }

        public void setAge(int age) {
            this.age = age;
        }

        @Override
        public String toString() {
            return "Person{" +
                    "name='" + name + '\'' +
                    ", age=" + age +
                    '}';
        }
    }

    /**
     * 创建一个Person的集合
     * @return List
     */
    public static List<Person> createPeople() {
        ArrayList<Person> people = new ArrayList<>();
        Person person1 = new Person("小明", 15);
        Person person2 = new Person("小芳", 20);
        Person person3 = new Person("小李", 18);
        Person person4 = new Person("小付", 23);
        Person person5 = new Person("大飞", 22);
        people.add(person1);
        people.add(person2);
        people.add(person3);
        people.add(person4);
        people.add(person5);
        return people;
    }

    public static void main(String[] args) {
        List<Person> people = createPeople();
        // 创建 Stream 对象
        Stream<Person> stream = people.stream();
        // 截断流，并调用终止操作打印集合中元素
        stream.limit(2).forEach(System.out::println);
    }
}
```







运行结果：



```java
Person{name='小明', age=15}  
Person{name='小芳', age=20}
```



根据运行结果显示，我们只打印了集合中的前两条数据。



跳过前 2 条数据的代码实例如下：



```java
// 非完整代码
public static void main(String[] args) {
    List<Person> people = createPeople();
    // 创建 Stream 对象
    Stream<Person> stream = people.stream();
    // 跳过前两个元素，并调用终止操作打印集合中元素
    stream.skip(2).forEach(System.out::println);
}
```



运行结果：



```java
Person{name='小李', age=18}  
Person{name='小付', age=23}  
Person{name='大飞', age=22}
```



`distinct()`方法会根据`equals()`和`hashCode()`方法筛选重复数据，我们在`Person`类内部重写这两个方法，并且在`createPerson()`方法中，添加几个重复的数据 ，实例如下：








```java
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;
import java.util.stream.Stream;

public class StreamDemo4 {

    static class Person {
        private String name;
        private int age;

        public Person() { }

        public Person(String name, int age) {
            this.name = name;
            this.age = age;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public int getAge() {
            return age;
        }

        public void setAge(int age) {
            this.age = age;
        }

        @Override
        public String toString() {
            return "Person{" +
                    "name='" + name + '\'' +
                    ", age=" + age +
                    '}';
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Person person = (Person) o;
            return age == person.age &&
                    Objects.equals(name, person.name);
        }

        @Override
        public int hashCode() {
            return Objects.hash(name, age);
        }
    }

    /**
     * 创建一个Person的集合
     * @return List
     */
    public static List<Person> createPeople() {
        ArrayList<Person> people = new ArrayList<>();
        people.add(new Person("小明", 15));
        people.add(new Person("小芳", 20));
        people.add(new Person("小李", 18));
        people.add(new Person("小付", 23));
        people.add(new Person("小付", 23));
        people.add(new Person("大飞", 22));
        people.add(new Person("大飞", 22));
        people.add(new Person("大飞", 22));
        return people;
    }

    public static void main(String[] args) {
        List<Person> people = createPeople();
        // 创建 Stream 对象
        Stream<Person> stream = people.stream();

        System.out.println("去重前，集合中元素有：");
        stream.forEach(System.out::println);

        System.out.println("去重后，集合中元素有：");
        // 创建一个新流
        Stream<Person> stream1 = people.stream();
        // 截断流，并调用终止操作打印集合中元素
        stream1.distinct().forEach(System.out::println);
    }
}
```







运行结果：



```java
去重前，集合中元素有：  
Person{name='小明', age=15}  
Person{name='小芳', age=20}  
Person{name='小李', age=18}  
Person{name='小付', age=23}  
Person{name='小付', age=23}  
Person{name='大飞', age=22}  
Person{name='大飞', age=22}  
Person{name='大飞', age=22}  
去重后，集合中元素有：  
Person{name='小明', age=15}  
Person{name='小芳', age=20}  
Person{name='小李', age=18}  
Person{name='小付', age=23}  
Person{name='大飞', age=22}
```



### 5.2 映射



关于映射中间操作，有下面几个常用方法：



* `map(Function f)`：接收一个方法作为参数，该方法会被应用到每个元素上，并将其映射成一个新的元素；


* `mapToDouble(ToDoubleFunction f)`：接收一个方法作为参数，该方法会被应用到每个元素上，产生一个新的`DoubleStream`；


* `mapToLong(ToLongFunction f)`：接收一个方法作为参数，该方法会被应用到每个元素上，产生一个新的`LongStream`；


* `flatMap(Function f)`：接收一个方法作为参数，将流中的每个值都换成另一个流，然后把所有流连接成一个流。


请查看如下实例：








```java
import java.util.Arrays;
import java.util.List;

public class StreamDemo5 {

    public static void main(String[] args) {
        // 创建一个包含小写字母元素的字符串列表
        List<String> stringList = Arrays.asList("php", "js", "python", "java");
        // 调用 map() 方法，将String下的toUpperCase()方法作为参数，这个方法会被应用到每个元素上，映射成一个新元素，最后打印映射后的元素
        stringList.stream().map(String::toUpperCase).forEach(System.out::println);
    }

}
```







运行结果：



```java
PHP  
JS  
PYTHON  
JAVA
```



可参考下图，理解映射的过程：



![](http://img.mukewang.com/wiki/5f08152d09e9640113921040.jpg)



### 5.3 排序



关于排序中间操作，有下面几个常用方法：



* `sorted()`：产生一个新流，其中按照自然顺序排序；


* `sorted(Comparator com)`：产生一个新流，其中按照比较器顺序排序。


请查看如下实例：








```java
import java.util.Arrays;
import java.util.List;

public class StreamDemo6 {

    public static void main(String[] args) {
        List<Integer> integers = Arrays.asList(10, 12, 9, 8, 20, 1);
        // 调用sorted()方法自然排序，并打印每个元素
        integers.stream().sorted().forEach(System.out::println);
    }

}
```







运行结果：



```java
1  
8  
9  
10  
12  
20
```



上面实例中，我们调用`sorted()`方法对集合元素进行了从小到大的自然排序，那么如果想要实现从大到小排序，任何实现呢？此时就要用到`sorted(Comparator com)`方法定制排序，查看如下实例：








```java
import java.util.Arrays;
import java.util.List;

public class StreamDemo6 {

    public static void main(String[] args) {
        List<Integer> integers = Arrays.asList(10, 12, 9, 8, 20, 1);
        // 定制排序
        integers.stream().sorted(
                (i1, i2) -> -Integer.compare(i1, i2)
        ).forEach(System.out::println);
    }

}
```







运行结果：



```java
20
12
10
9
8
1
```



实例中，`sorted()`方法接收的参数是一个函数式接口`Comparator`，因此使用`Lambda`表达式创建函数式接口实例即可，`Lambda`体调用整型的比较方法，对返回的整型值做一个取反即可。



## 6. Stream 的终止操作



执行终止操作会从流的**流水线**上生成结果，其结果可以是任何不是流的值，例如`List`、`String`、`void`。



在上面实例中，我们一直在使用`forEach()`方法来执行流的终止操作，下面我们看看还有哪些其他终止操作。



### 6.1 匹配与查找



关于匹配与查找的终止操作，有下面几个常用方法：



* `allMatch(Predicate p)`：检查是否匹配所有元素；


* `anyMatch(Predicate p)`：检查是否至少匹配一个元素；


* `noneMatch(Predicate p)`：检查是否没有匹配所有元素；


* `findFirst()`：返回第一个元素；


* `findAny()`：返回当前流中的任意元素；


* `count()`：返回流中元素总数；


* `max(Comparator c)`：返回流中最大值；


* `min(Comparator c)`：返回流中最小值；


* `forEach(Consumer c)`：内部迭代（使用 Collection 接口需要用户去做迭代，称为外部迭代；相反 `Stream API`使用内部迭代）。


如下实例，演示了几个匹配元素相关方法的使用：








```java
import java.util.Arrays;
import java.util.List;

public class StreamDemo7 {

    public static void main(String[] args) {
        // 创建一个整型列表
        List<Integer> integers = Arrays.asList(10, 12, 9, 8, 20, 1);
        // 使用 allMatch(Predicate p) 检查是否匹配所有元素，如果匹配，则返回 true；否则返回 false
        boolean b1 = integers.stream().allMatch(integer -> integer > 0);
        if (b1) {
            System.out.println(integers + "列表中所有的元素都大于0");
        } else {
            System.out.println(integers + "列表中不是所有的元素都大于0");
        }

        // 使用 anyMatch(Predicate p) 检查是否至少匹配一个元素
        boolean b2 = integers.stream().anyMatch(integer -> integer >= 20);
        if (b2) {
            System.out.println(integers + "列表中至少存在一个的元素都大于等于20");
        } else {
            System.out.println(integers + "列表中不存在任何一个大于等于20的元素");
        }

        // 使用 noneMath(Predicate p) 检查是否没有匹配所有元素
        boolean b3 = integers.stream().noneMatch(integer -> integer > 100);
        if (b3) {
            System.out.println(integers + "列表中不存在大于100的元素");
        } else {
            System.out.println(integers + "列表中存在大于100的元素");
        }
    }

}
```







运行结果：



```java
[10, 12, 9, 8, 20, 1]列表中所有的元素都大于0
[10, 12, 9, 8, 20, 1]列表中至少存在一个的元素都大于等于20
[10, 12, 9, 8, 20, 1]列表中不存在大于100的元素
```



查找元素的相关方法使用实例如下：








```java
import java.util.Arrays;
import java.util.List;
import java.util.Optional;

public class StreamDemo8 {

    public static void main(String[] args) {
        // 创建一个整型列表
        List<Integer> integers = Arrays.asList(10, 12, 9, 8, 20, 1);

        // 使用 findFirst() 获取当前流中的第一个元素
        Optional<Integer> first = integers.stream().findFirst();
        System.out.println(integers + "列表中第一个元素为：" + first);

        // 使用 findAny() 获取当前流中的任意元素
        Optional<Integer> any = integers.stream().findAny();
        System.out.println("列表中任意元素：" + any);

        // 使用 count() 获取当前流中元素总数
        long count = integers.stream().count();
        System.out.println(integers + "列表中元素总数为" + count);

        // 使用 max(Comparator c) 获取流中最大值
        Optional<Integer> max = integers.stream().max(Integer::compare);
        System.out.println(integers + "列表中最大值为" + max);

        // 使用 min(Comparator c) 获取流中最小值
        Optional<Integer> min = integers.stream().min(Integer::compare);
        System.out.println(integers + "列表中最小值为" + min);
    }

}
```







运行结果：



```java
[10, 12, 9, 8, 20, 1]列表中第一个元素为：Optional[10]
列表中任意元素：Optional[10]
[10, 12, 9, 8, 20, 1]列表中元素总数为6
[10, 12, 9, 8, 20, 1]列表中最大值为Optional[20]
[10, 12, 9, 8, 20, 1]列表中最小值为Optional[1]
```



实例中，我们观察到`findFirst()`、`findAny()`、`max()`等方法的返回值类型为`Optional`类型，关于这个`Optional`类，我们将在下一小节具体介绍。



### 6.2 归约



关于归约的终止操作，有下面几个常用方法：



* `reduce(T identity, BinaryOperator b)`：可以将流中的元素反复结合起来，得到一个值。返回 T；


* `reduce(BinaryOperator b)`：可以将流中的元素反复结合起来，得到一个值，返回 `Optional<T>`。


归约相关方法的使用实例如下：








```java
import java.util.Arrays;
import java.util.List;
import java.util.Optional;

public class StreamDemo9 {

    public static void main(String[] args) {
        // 创建一个整型列表
        List<Integer> integers = Arrays.asList(10, 12, 9, 8, 20, 1);

        // 使用 reduce(T identity, BinaryOperator b) 计算列表中所有整数和
        Integer sum = integers.stream().reduce(0, Integer::sum);
        System.out.println(sum);

        // 使用 reduce(BinaryOperator b) 计算列表中所有整数和，返回一个 Optional<T>
        Optional<Integer> reduce = integers.stream().reduce(Integer::sum);
        System.out.println(reduce);
    }

}
```







运行结果：



```java
60
Optional[60]
```



### 6.3 收集



`collect(Collector c)`：将流转换为其他形式。接收一个`Collector`接口的实现，用于给`Stream`中元素做汇总的方法。



实例如下：








```java
import java.util.Arrays;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

public class StreamDemo10 {

    public static void main(String[] args) {
        // 创建一个整型列表
        List<Integer> integers = Arrays.asList(10, 12, 9, 8, 20, 1, 10);
        Set<Integer> collect = integers.stream().collect(Collectors.toSet());
        System.out.println(collect);
    }

}
```







运行结果：



```java
[1, 20, 8, 9, 10, 12]
```



Collector 接口中的实现决定了如何对流执行收集的操作（如收集到 List、Set、Map）。`java.util.stream.Collectors` 类提供了很多静态方法，可以方便地创建常用收集器实例，常用静态方法如下：



* `static List<T> toList()`：把流中元素收集到`List`；


* `static Set<T> toSet()`：把流中元素收集到`Set`；


* `static Collection<T> toCollection()`：把流中元素收集到创建的集合。


## 7. 小结



通过本小节的学习，我们知道了`Stream`不同于`java.io`下的输入输出流，它主要用于处理数据。`Stream API`可用于处理非关系型数据库中的数据；想要使用流式操作，就要知道创建`Stream`对象的几种方式；流式操作可分为创建`Stream`对象、中间操作和终止操作三个步骤。多个中间操作可以连接起来形成一个流水线，除非流水线上触发终止操作，否则中间操作不会执行任何的处理。执行终止操作会从流的**流水线**上生成结果，其结果可以是任何不是流的值。






### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
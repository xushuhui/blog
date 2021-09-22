# 函数式数据处理

Java 8 中新增的特性其目的是为了帮助开发人员写出更好地代码，其中关键的一部分就是对核心类库的改进。流（ Stream ）和集合类库便是核心类库改进的内容。

## 1. 从外部迭代到内部迭代

对于一个集合迭代是我们常用的一种操作，通过迭代我们可以处理返回每一个操作。常用的就是 `for` 循环了。我们来看个例子：

```java
import java.util.Arrays;
import java.util.List;

public class Test{
	public static void main(String...s){
		 List<Integer> numbers = Arrays.asList(new Integer[]{1,2,3,4,5,6,7});
		 int counter = 0;
		 for(Integer integer : numbers){
			 if(integer > 5) counter++;
		 }
		 System.out.println(counter);
	}
}
```

```java
输出： 2
```

这里我们统计数组 `numbers` 中大于 5 的元素的个数，我们通过 `for` 循环对 `numbers` 数组进行迭代，随后对每一个元素进行比较。这个调用过程如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f03e86f084bb26e05620480.jpg)

在这个过程中，编译器首先会调用 `List` 的 `iterator()` 方法产生一个 `Iterator` 对象来控制迭代过程，这个过程我们称之为 **外部迭代**。 在这个过程中会显示调用 `Iterator` 对象的 `hasNext()` 和 `next()` 方法来完成迭代。

这样的外部迭代有什么问题呢？

> **Tips：** 对于循环中不同操作难以抽象。

比如我们前面的例子，假设我们要对大于 5 小于 5 和等于 5 的元素分别进行统计，那么我们所有的逻辑都要写在里面，并且只有通过阅读里面的逻辑代码才能理解其意图，这样的代码可阅读性是不是和 Lambda 表达式的可阅读性有着天壤之别呢？

Java 8 中提供了另一种通过 Stream 流来实现 **内部迭代** 的方法。我们先来看具体的例子：

```java
import java.util.List;
import java.util.Arrays;

public class Test{
	public static void main(String...s){
	    List<Integer> numbers = Lists.newArrayList(1,2,3,4,5,6,7);
	    long counter = numbers.stream().filter(e->e>5).count();
	    System.out.println(counter);
	}
}
```

```java
输出： 2
```

在这个例子中，我们调用 `stream()` 方法获取到 Stream 对象，然后调用该对象的 `filter()` 方法对元素进行过滤，最后通过 `count()` 方法来计算过滤后的 Stream 对象中包含多少个元素对象。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f03e8ae085028f405340340.jpg)

与外部迭代不同，内部迭代并不是返回控制对象 `Iterator`， 而是返回内部迭代中的相应接口 `Stream`。进而把对集合的复杂逻辑操作变成了明确的构建操作。

在这个例子中，通过内部迭代，我们把整个过程被拆分成两步：

1. 找到大于 5 的元素；
2. 统计这些元素的个数。

这样一来我们代码的可读性是不是大大提升了呢？

## 2. 实现机制

在前面的内部迭代例子中，整个操作过程被分成了过滤和计数两个简单的操作。我们再来看一下之前的例子，并做了一些改造：

```java
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.stream.Stream;

public class Test{
    public static void main(String...s){
        List<Integer> numbers = new ArrayList<Integer>();
        Collections.addAll(numbers,new Integer[]{1,2,3,4,5,6,7});
        Stream stream1 = numbers.stream();
        numbers.remove(6);
        //直接使用numbers的stream()
        long counter = numbers.stream().filter(e->e>5).count();
        System.out.println(counter);
        //调用之前的stream1
        counter = stream1.filter(ex-> (Integer)ex>5).count();
        System.out.println(counter);
    }
}
```

返回结果：

```java
1
1
```

在这个例子中，我们在获取到 Stream 对象 stream1 后删除了数组 `numbers` 中的最后一个元素，随后分别对 numbers 和 stream1 进行过滤统计操作，会发现两个结果是一样的，stream1 中的内容跟随 `numbers` 一起做相应的改变。这说明 **Stream 对象不是一个新的集合，而是创建新集合的配方**。同样，像 `filter()` 虽然返回 Stream 对象，但也只是对 Stream 的刻画，并没有产生新的集合。

我们通常对于这种不产生集合的方法叫做 **惰性求值方法**，相对应的类似于 `count()` 这种返回结果的方法我们叫做 **及早求值方法**。

我们可以把多个惰性求值方法组合起来形成一个惰性求值链，最后通过及早求值操作返回想要的结果。这类似建造者模式，使用一系列操作设置属性和配置，最后通过一个 `build` 的方法来创建对象。通过这样的一个过程我们可以让我们对集合的构建过程一目了然。这也就是 Java 8 风格的集合操作。

## 3. 常用的流操作

为了更好地理解 Stream API，我们需要掌握它的一些常用操作，接下来我们将逐个学习几种常用的操作。

### 3.1 collect

`collect` 操作是根据 Stream 里面的值生成一个列表，它是一个求值操作。

> **Tips：**`collect` 方法通常会结合 `Collectors` 接口一起使用，是一个通用的强大结构，可以满足数据转换、数据分块、字符串处理等操作。

我们来看一些例子：

1. 生成集合：

```java
import java.util.stream.Stream;
import java.util.List;
import java.util.stream.Collectors;

public class Test{
	public static void main(String...s){
		List<String> collected = Stream.of("a","b","c").collect(Collectors.toList());
		System.out.println(collected);
	}
}
```

```java
输出：[a, b, c]
```

使用 `collect(Collectors.toList())` 方法从 Stream 中生成一个列表。

1. 集合转换：

使用 `collect` 来定制集合收集元素。

```java
import java.util.List;
import java.util.stream.Stream;
import java.util.stream.Collectors;
import java.util.TreeSet;

public class Test{
	public static void main(String...s){
		List<String> collected = Stream.of("a","b","c","c").collect(Collectors.toList());
		TreeSet<String> treeSet = collected.stream().collect(Collectors.toCollection(TreeSet::new));
		System.out.println(collected);
		System.out.println(treeSet);
	}
}
```

```java
输出结果：
[a, b, c, c]
[a, b, c]
```

使用 `toCollection` 来定制集合收集元素，这样就把 `List` 集合转换成了 `TreeSet`

1. 转换成值：

使用 `collect` 来对元素求值。

```java
import java.util.List;
import java.util.stream.Stream;
import java.util.stream.Collectors;

public class Test{
	public static void main(String...s){
		List<String> collected = Stream.of("a","b","c").collect(Collectors.toList());
		String maxChar = collected.stream().collect(Collectors.maxBy(String::compareTo)).get();
		System.out.println(maxChar);
	}
}
```

```java
输出： c
```

上面我们使用 `maxBy` 接口让收集器生成一个值，通过方法引用调用了 `String` 的 `compareTo` 方法来比较元素的大小。同样还可以使用 `minBy` 来获取最小值。

1. 数据分块：

比如我们对于数据 1-7 想把他们分成两组，一组大于 5 另外一组小于等于 5，我们可以这么做：

```java
import java.util.List;
import java.util.Map;
import java.util.stream.Stream;
import java.util.stream.Collectors;

public class Test{
	public static void main(String...s){
		List<Integer> collected = Stream.of(1,2,3,4,5,6,7).collect(Collectors.toList());
		Map<Boolean,List<Integer>> divided = collected.stream().collect(Collectors.partitioningBy(e -> e>5));
		System.out.println(divided.get(true));
		System.out.println(divided.get(false));
	}
}
```

```java
输出结果：
[6, 7]
[1, 2, 3, 4, 5]
```

通过 `partitioningBy` 接口可以把数据分成两类，即满足条件的和不满足条件的，最后将其收集成为一个 `Map` 对象，其 `Key` 为 `Boolean` 类型，`Value` 为相应的集合元素。

同样我们还可以使用 `groupingBy` 方法来对数据进行分组收集，这类似于 SQL 中的 `group by` 操作。

1. 字符串处理：

`collect` 还可以来将元素组合成一个字符串。

```java
import java.util.List;
import java.util.stream.Stream;
import java.util.stream.Collectors;

public class Test{
public static void main(String...s){
    List<String> collected = Stream.of("a","b","c").collect(Collectors.toList());
    String formatted = collected.stream().collect(Collectors.joining(",","[","]"));
    System.out.println(formatted);
}
}
```

```java
输出：[a,b,c]
```

这里我们把 `collected` 数组的每个元素拼接起来，并用 `[``]` 包裹。

### 3.2 map

`map` 操作是将流中的对象换成一个新的流对象，是 Stream 上常用操作之一。 其示意图如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f03f03508ff160305730426.jpg)

​

比如我们把小写字母改成大写，通常我们会使用 `for` 循环：

```java
import java.util.List;
import java.util.ArrayList;
import java.util.Collections;

public class Test{
public static void main(String...s){
    List<String>  collected = new ArrayList<>();
     List<String>  newArr = new ArrayList<>();
    Collections.addAll(newArr,new String[]{"a","b","c"});
    for(String string : newArr){
        collected.add(string.toUpperCase());
    }
    System.out.println(collected);
}
}
```

```java
输出： [A, B, C]
```

此时，我们可以使用 `map` 操作来进行转换：

```java
import java.util.List;
import java.util.stream.Stream;
import java.util.stream.Collectors;

public class Test{
	public static void main(String...s){
	    List<String> collected = Stream.of("a","b","c").collect(Collectors.toList());
	    List<String> upperCaseList = collected.stream().map(e->e.toUpperCase()).collect(Collectors.toList());
	    System.out.println(upperCaseList);
	}
}

```

```java
输出： [A, B, C]
```

在 `map` 操作中，我们 把 `collected` 中的每一个元素转换成大写，并返回。

### 3.3 flatmap

`flatmap` 与 `map` 功能类似，只不过 `map` 对应的是一个流，而 `flatmap` 可以对应多个流。

我们来看一个例子：

```java
import java.util.List;
import java.util.Arrays;
import java.util.stream.Collectors;

public class Test{
	public static void main(String...s){
	    List<String> nameA = Arrays.asList("Mahela", "Sanga", "Dilshan");
	    List<String> nameB = Arrays.asList("Misbah", "Afridi", "Shehzad");
	    List<List<String>> nameSets = Arrays.asList(nameA,nameB);
	    List<String> flatMapList = nameSets.stream()
	            .flatMap(pList -> pList.stream())
	            .collect(Collectors.toList());
	    System.out.println(flatMapList);
	}
}
```

```java
返回结果： [Mahela, Sanga, Dilshan, Misbah, Afridi, Shehzad]
```

通过 `flatmap` 我们把集合 `nameSets` 中的字集合合并成了一个集合。

### 3.4 filter

`filter` 用来过滤元素，在元素遍历时，可以使用 `filter` 来提取我们想要的内容，这也是集合常用的方法之一。其示意图如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f03f08e08b3ddae05220441.jpg)

​

我们来看一个例子：

```java
import java.util.List;
import java.util.ArrayList;
import java.util.Collections;
import java.util.stream.Collectors;


public class Test{
	public static void main(String...s) {
	    List<Integer> numbers = new ArrayList<>();
	    Collections.addAll(numbers,new Integer[]{1,2,3,4,5,6,7});
	    List<Integer> collected = numbers.stream()
									    .filter(e->e>5).collect(Collectors.toList());
	    System.out.println(collected);
	}
}
```

```java
输出：[6, 7]
```

此时，`filter` 会遍历整个集合，将满足将满足条件的元素提取出来，并通过收集器收集成新的集合。

### 3.5 max/min

`max/min` 求最大值和最小值也是集合上常用的操作。它通常会与 `Comparator` 接口一起使用来比较元素的大小。示例如下：

```java
import java.util.List;
import java.util.Collections;

public class Test{
	public static void main(String...s) {
	    List<Integer> numbers = new ArrayList<>();
	    Collections.addAll(numbers,new Integer[]{1,2,3,4,5,6,7});
	    Integer max = numbers.stream().max(Comparator.comparing(k->k)).get();
	    Integer min = numbers.stream().min(Comparator.comparing(k->k)).get();
	    System.out.println("max:"+max);
	    System.out.println("min:"+min);
	}
}
```

```java
输出：
max:7
min:1
```

我们可以在 Comparator 接口中定制比较条件，来获得想要的结果。

### 3.6 reduce

`reduce` 操作是可以实现从流中生成一个值，我们前面提到的如 `count`、`max`、`min` 这种及早求值就是由`reduce` 提供的。我们来看一个例子：

```java
import java.util.stream.Stream;

public class Test{
	public static void main(String...s) {
	    int sum = Stream.of(1,2,3,4,5,6,7).reduce(0,(acc,e)-> acc + e);
	    System.out.println(sum);
	}
}
```

```java
输出：28
```

上面的例子是对数组元素进行求和，这个时候我们就要使用 `reduce` 方法。这个方法，接收两个参数，第一个参数相当于是一个初始值，第二参数则为具体的业务逻辑。 上面的例子中，我们给 `acc` 参数赋予一个初始值 0 ，随后将 `acc` 参数与各元素求和。

## 4. 小结

![](https://xushuhui.gitee.io/image/imooc/5f1a932e09e55a2f04980294.jpg)

以上我们学习了 Java 8 的流及常用的一些集合操作。我们需要常用的函数式接口和流操作非常熟悉才能更好地使用这些新特性。

另外，请思考一个问题，在本节关于集合的操作中都将集合通过 `stream()` 方法转换成了 Stream 对象，那么我们还有必要对外暴露一个集合对象（List 或者 Set）吗？

> **Tips：** 在编程过程中，使用 Stream 工厂比对外暴露集合对象要更好一些。仅需要暴露 Stream 接口，在实际操作中无论怎么使用都影响内部的集合。

所以，Java 8 风格不是一蹴而就的，我们可以对已有的代码进行重构来练习和强化 Java 8 的编程风格，时间长了自然就对 Stream 对象有更深的理解了。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

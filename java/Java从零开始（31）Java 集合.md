# Java 集合

在前面的小节中，我们学习了数组，本小节学习的**集合**同样用于存放一组数据，我们将学习**什么是集合**、**集合的应用场景** ，在应用场景部分我们将对比 **Java 数组与集合的区别**，还将系统介绍 Java 集合的架构，也将结合实例来讲解集合的实际应用。

## 1. 什么是集合

> 在计算机科学中，集合是一组可变数量的数据项（也可能为 0 个）的组合，这些数据可能共享某些特征，需要以某种操作方式一起进行操作。

Java 中集合主要分为`java.util.Collection`和`java.util.Map`两大接口。

下图描绘了 Java 集合的框架：

![](https://xushuhui.gitee.io/image/imooc/5ec76b7a091829fd07620566.jpg)

**Tips：** 图表最下方的`ArrayList`、`LinkedList`、`HashSet`以及`HashMap`都是常用实现类，本小节将介绍具体使用。

### 1.1 Collection

`java.util.Collection`接口的实现可用于存储 Java 对象。例如，慕课网的所有学生可以视为一个`Collection`。

`Collection`又可以分为三个子接口，分别是：

1. `List`：序列，必须按照顺序保存元素，因此它是有序的，允许重复；
2. `Queue`：队列，按照排队规则来确定对象产生的顺序，有序，允许重复；
3. `Set`：集，不能重复。

### 1.2 Map

`java.util.Map`接口的实现可用于表示“键”（key）和“值”（value）对象之间的映射。一个映射表示一组“键”对象，其中每一个“键”对象都映射到一个“值”对象。因此可以通过键来查找值。例如，慕课网的每一个学生都有他自己的账户积分，这个关联关系可以用`Map`来表示。

## 2. 集合的应用场景

### 2.1 数组与集合

在介绍集合的应用场景之前，我们先来看看数组和集合的对比。

我们知道数组和集合都用于存放一组数据，但数组的容量是固定大小的，而集合的容量是动态可变的；对于可存放的数据类型，数组既可以存放基本数据类型又可以存放引用数据类型，而集合只能存放引用数据类型，基本数据类型需要转换为对应的包装类才能存放到集合当中。

### 2.2 集合应用场景

* **无法预测存储数据的数量**：由于数组容量是固定大小，因此使用集合存储动态数量的数据更为合适；
* **同时存储具有一对一关系的数据**：例如存储慕课网学生的积分，为了方便检索对应学生的积分，可使用`Map`将慕课网学生的`uid`和对应的积分进行一对一关联；
* **数据去重**：使用数组实现需要遍历，效率低，而`Set`集合本身就具有不能重复的特性；
* **需要数据的增删**：使用数组实现增删操作需要遍历、移动数组中元素，如果操作频繁会导致效率降低。

## 3. List 集合

### 3.1 概念和特性

List 是元素有序并且可以重复的集合，称之为序列。序列可以精确地控制每个元素的插入位置或删除某个位置的元素。通过前面的学习，我们知道`List`是`Collection`的一个子接口，它有两个主要实现类，分别为`ArrayList`（动态数组）和`LinkedList`（链表）。

### 3.2 ArrayList 实现类

ArrayList 可以理解为动态数组，它的容量可以动态增长。**当添加元素时，如果发现容量已满，会自动扩容为原始大小的 1.5 倍。**

#### 3.2.1 构造方法

* `ArrayList()`：构造一个初始容量为 10 的空列表；
* `ArrayList(int initialCapacity)`：构造一个指定容量的空列表；
* `ArrayList(Collection<? extends E> c)`：构造一个包含指定集合元素的列表，其顺序由集合的迭代器返回。

在代码中，我们可以这样实例化`ArrayList`对象：

```java
// 无参构造实例化，初始容量为10
List arrayList1 = new ArrayList();
// 实例化一个初始容量为20的空列表
List arrayList2 = new ArrayList(20);
// 实例化一个集合元素为 arrayList2 的列表（由于 arrayList2 为空列表，因此其实例化的对象也为空列表）
List arrayList3 = new ArrayList(arrayList2);
```

#### 3.2.2 常用成员方法

* `void add(E e)`：将指定的元素追加到此列表的末尾；

* `void add(int index, E element)`：将指定的元素插入此列表中的指定位置；

* `E remove(int index)`：删除此列表中指定位置的元素；

* `boolean remove(Object o)`：如果存在指定元素，则从该列表中删除第一次出现的该元素；

* `void clear()`：从此列表中删除所有元素；

* `E set(int index, E element)`：用指定的元素替换此列表中指定位置的元素；

* `E get(int index)`：返回此列表中指定位置的元素；

* `boolean contains(Object o)`：如果此列表包含指定的元素，则返回 true，否则返回 false；

* `int size()`：返回该列表中元素的数量；

* `Object[] toArray()`：以正确的顺序（从第一个元素到最后一个元素）返回一个包含此列表中所有元素的数组。

更多成员方法请翻阅[官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/ArrayList.html)，下面我们将结合实例来介绍以上成员方法的使用。

### 3.3 实例

#### 3.3.1 新增元素

请查看如下实例：

```java
import java.util.ArrayList;
import java.util.List;

public class ArrayListDemo1 {

    public static void main(String[] args) {
        // 实例化一个空列表
        List arrayList = new ArrayList();
        for (int i = 0; i < 5; i ++) {
            // 将元素 i 追加到列表的末尾
            arrayList.add(i);
            // 打印列表内容
            System.out.println(arrayList);
        }
    }
}
```

运行结果：

```java
[0]
[0, 1]
[0, 1, 2]
[0, 1, 2, 3]
[0, 1, 2, 3, 4]
```

代码中，首先实例化了一个`ArrayList`对象，然后使用 for 循环语句循环 5 次，每次都向`arrayList`对象中追加变量`i`，并打印列表内容，运行结果清晰的展示了每次新增元素的过程。

> **Tips**：由于`ArrayList`的父类`AbstractCollection`重写了`toString()`方法，因此直接打印列表，可以直观地展示出列表中的元素。

#### 3.3.2 泛型初识

> **Tips**：泛型（`Genericity`）**将在下一小节详细介绍，此处我们只简要介绍一下泛型以及其使用方法。如果你比较了解泛型，可直接跳过此知识点。

如果你使用`IDEA`编写如上代码，将会有下图所示的 3 处黄色警告：

![](https://xushuhui.gitee.io/image/imooc/5ec78dbb0976ca8517040871.jpg)

既然`IDE`有了警告，我们就尝试来解决一下，将鼠标光标放置到警告处，会提示“Unchecked call to ‘add(E)’ as a member of raw type ‘java.util.List’ ”，这是`IDE`的**泛型检查**，可点击`Try to generify 'ArrayListDemo1.java'`按钮：

![](https://xushuhui.gitee.io/image/imooc/5ec78da1090e53af16390874.jpg)

此时会出现一个`Generify`的弹窗，直接点击`Refactor`按钮：

![](https://xushuhui.gitee.io/image/imooc/5ec78e3a09c402b916961052.jpg)

代码变成了下图所示的样子，那 3 处警告被成功消除了：

![](https://xushuhui.gitee.io/image/imooc/5ec78e6a09f7833012080847.jpg)

我们观察到代码第 8 行的`List`类型后面多了一对尖括号“`<>`”，`<>`里面是 Java 的包装类型`Integer`，在`ArrayList`类型后面也多了一对尖括号，这里的`<>`中承载的就是 Java 的泛型的类型参数，它表示`arrayList`对象用于存放`Integer`类型的数据。这样的目的和好处这里不详细展开讨论，本小节我们只需知道这样做就可以消除`IDEA`的警告即可。

由于前面`List`已经指定了泛型的参数类型为`Integer`，后面的`ArrayList`就不需要再重复指定了。当然你也可以这样写（但是没必要）：

```java
List<Integer> arrayList = new ArrayList<Integer>();
```

同理，如果你想向`arrayList`存放`String`类型的元素，只需将`<Integer>`改为`<String>`，我们再来看一个实例：

```java
import java.util.ArrayList;
import java.util.List;

public class ArrayListDemo2 {
    public static void main(String[] args) {
        // 实例化一个空列表
        List<String> arrayList = new ArrayList<>();
        // 将字符串元素 Hello 追加到此列表的末尾
        arrayList.add("Hello");
        // 将字符串元素 World 追加到此列表的末尾
        arrayList.add("World");
        // 打印列表
        System.out.println(arrayList);
        // 将字符串元素 Java 插入到此列表中的索引为 1 的位置
        arrayList.add(1, "Java");
        // 打印列表
        System.out.println(arrayList);
    }
}
```

运行结果：

```java
[Hello, World]
[Hello, Java, World]
```

代码中，首先实例化了一个`ArrayList`的对象，调用了两次`add(E e)`方法，依次向列表尾部插入了`Hello`和`World`元素，列表中元素为`[Hello, World]`，此时调用`add(int index, E element)`方法，将字符串元素 Java 插入到此列表中的索引为 1 的位置，因此列表中的元素为`[Hello, Java, World]`。

#### 3.3.3 删除元素

请查看如下实例：

```java
import java.util.ArrayList;
import java.util.List;

public class ArrayListDemo3 {

    public static void main(String[] args) {
        // 实例化一个空列表
        List<String> arrayList = new ArrayList<>();
        // 将字符串元素 Hello 追加到此列表的末尾
        arrayList.add("Hello");
        // 将字符串元素 World 追加到此列表的末尾
        arrayList.add("World");
        // 将字符串元素 Hello 追加到此列表的末尾
        arrayList.add("Hello");
        // 将字符串元素 Java 追加到此列表的末尾
        arrayList.add("Java");
        // 打印列表
        System.out.println(arrayList);

        // 删除此列表中索引位置为 3 的元素
        arrayList.remove(3);
        // 打印列表
        System.out.println(arrayList);

        // 删除此列表中第一次出现的 Hello 元素
        arrayList.remove("Hello");
        System.out.println(arrayList);
    }
}
```

运行结果：

```java
[Hello, World, Hello, Java]
[Hello, World, Hello]
[World, Hello]
```

代码中，我们首先添加了 4 个字符串元素，列表内容为`[Hello, World, Hello, Java]`，然后调用`remove(int index)`方法删除了索引位置为 3 的元素（即`Java`），此时列表内容为`[Hello, World, Hello]` ，再次调用`remove(Object o)`方法，删除了列表中第一次出现的`Hello`元素，此时列表内容为`[World, Hello]`。

#### 3.3.4 修改元素

可使用`set()`方法修改列表中元素，实例如下：

```java
import java.util.ArrayList;
import java.util.List;

public class ArrayListDemo4 {

    public static void main(String[] args) {
        // 实例化一个空列表
        List<String> arrayList = new ArrayList<>();
        arrayList.add("Hello");
        // 将字符串元素 World 追加到此列表的末尾
        arrayList.add("World");
        // 打印列表
        System.out.println(arrayList);
        // 用字符串元素 Hello 替换此列表中索引位置为 1 的元素
        arrayList.set(1, "Java");
        System.out.println(arrayList);
    }
}
```

运行结果：

```java
[Hello, World]
[Hello, Java]
```

#### 3.3.5 查询元素

可使用`get()`方法来获取列表中元素，实例如下：

```java
import java.util.ArrayList;
import java.util.List;

public class ArrayListDemo5 {

    public static void main(String[] args) {
        // 实例化一个空列表
        List<String> arrayList = new ArrayList<String>();
        arrayList.add("Hello");
        arrayList.add("Immoc");
        for (int i = 0; i < arrayList.size(); i ++) {
            System.out.println("索引位置" + i + "的元素为"  + arrayList.get(i));
        }
    }
}
```

运行结果：

```java
索引位置0的元素为Hello
索引位置1的元素为Immoc
```

我们在使用`for`循环遍历列表的时候，让限定条件为`i < arrayList.size();`，`size()`方法可获取该列表中元素的数量。

#### 3.2.7 自定义类的常用操作

请查看如下实例：

```java
import java.util.ArrayList;
import java.util.List;

public class ArrayListDemo6 {

    static class ImoocStudent {
        private String nickname;

        private String position;

        public ImoocStudent() {
        }

        public ImoocStudent(String nickname, String position) {
            this.setNickname(nickname);
            this.setPosition(position);
        }

        public String getNickname() {
            return nickname;
        }

        public void setNickname(String nickname) {
            this.nickname = nickname;
        }

        public String getPosition() {
            return position;
        }

        public void setPosition(String position) {
            this.position = position;
        }

        @Override
        public String toString() {
            return "ImoocStudent{" +
                    "nickname='" + nickname + '\'' +
                    ", position='" + position + '\'' +
                    '}';
        }
    }

    public static void main(String[] args) {
        // 实例化一个空列表
        List<ImoocStudent> arrayList = new ArrayList<>();
        // 实例化3个慕课网学生对象
        ImoocStudent imoocStudent1 = new ImoocStudent("Colorful", "服务端工程师");
        ImoocStudent imoocStudent2 = new ImoocStudent("Lillian", "客户端工程师");
        ImoocStudent imoocStudent3 = new ImoocStudent("小慕", "架构师");
        // 新增元素
        arrayList.add(imoocStudent1);
        arrayList.add(imoocStudent2);
        arrayList.add(imoocStudent3);
        System.out.println(arrayList);
        // 删除元素
        arrayList.remove(imoocStudent2);
        System.out.println("删除 imoocStudent2 后：arrayList 内容为：" + arrayList);
        arrayList.remove(1);
        System.out.println("删除列表中索引位置为 1 的元素后，arrayList 内容为：" + arrayList);
        // 实例化一个新的慕课网学生对象
        ImoocStudent imoocStudent4 = new ImoocStudent("小李", "UI设计师");
        // 修改元素
        arrayList.set(0, imoocStudent4);
        System.out.println("修改后：arrayList 内容为" + imoocStudent4);
        // 查询元素，将 get() 方法得到的 Object 类型强制转换为 ImoocStudent 类型
        ImoocStudent student = arrayList.get(0);
        System.out.println("索引位置 0 的学生的昵称为：" + student.getNickname());
        System.out.println("索引位置 0 的学生的职位为：" + student.getPosition());
    }
}
```

运行结果：

```java
[ImoocStudent{nickname='Colorful', position='服务端工程师'}, ImoocStudent{nickname='Lillian', position='客户端工程师'}, ImoocStudent{nickname='小慕', position='架构师'}]
删除 imoocStudent2 后：arrayList 内容为：[ImoocStudent{nickname='Colorful', position='服务端工程师'}, ImoocStudent{nickname='小慕', position='架构师'}]
删除列表中索引位置为 1 的元素后，arrayList 内容为：[ImoocStudent{nickname='Colorful', position='服务端工程师'}]
修改后：arrayList 内容为ImoocStudent{nickname='小李', position='UI设计师'}
索引位置 0 的学生的昵称为：小李
索引位置 0 的学生的职位为：UI设计师
```

为了方便演示，我们定义了一个静态内部类`ImoocStudent`，它有两个属性`nickname`和`position`，定义了属性的`getter`和`setter`，并重写了`toString()`方法。在`main()`方法中，我们实现了自定义类在`ArrayList`中的增删改查。

### 3.4 LinkedList 实现类

`LinkedList` 是一个以双向链表实现的`List`。和`ArrayList`一样，也按照索引位置排序，但它的元素是双向连接的，因此顺序访问的效率非常高，而随机访问的效率比较低。

#### 3.4.1 构造方法

* `LinkedList()`：构造一个空列表；
* `LinkedList(Collection<? extends E> c)`：构造一个包含指定集合元素的列表，其顺序由集合的迭代器返回。

#### 3.4.2 常用成员方法

* `void add(E e)`：将指定的元素追加到此列表的末尾；
* `void add(int index, E element)`：将指定的元素插入此列表中的指定位置；
* `void addFirst(E e)`：将指定的元素插入此列表的开头；
* `vod addLast(E e)`：将指定的元素添加到此列表的结尾；
* `E remove(int index)`：删除此列表中指定位置的元素；
* `boolean remove(Object o)`：如果存在指定元素，则从该列表中删除第一次出现的该元素；
* `void clear()`：从此列表中删除所有元素；
* `E set(int index, E element)`：用指定的元素替换此列表中指定位置的元素；
* `E get(int index)`：返回此列表中指定位置的元素；
* `E getFirst()`：返回此列表的第一个元素；
* `E getLast()`：返回此列表的最后一个元素；
* `boolean contains(Object o)`：如果此列表包含指定的元素，则返回 true，否则返回 false；
* `int size()`：返回该列表中元素的数量；
* `Object[] toArray()`：以正确的顺序（从第一个元素到最后一个元素）返回一个包含此列表中所有元素的数组。

更多成员方法请翻阅[官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/LinkedList.html)，对于成员方法的使用，与`ArrayList`大同小异，这里不再赘述。

## 4. Set 集合

### 4.1 概念和特性

`Set`是元素无序并且不可以重复的集合，我们称之为集。`Set`是`Collection`的一个子接口，它的主要实现类有：`HashSet`、`TreeSet`、`LinkedHashSet`、`EnumSet`等，下面我们将详细介绍最常用的`HashSet`实现类。

### 4.2 HashSet 实现类

`HashSet`类依赖于哈希表（实际上是`HashMap`实例，下面将会介绍）。`HashSet`中的元素是无序的、散列的。

#### 4.2.1 构造方法

* `HashSet()`：构造一个新的空集；默认的初始容量为 16（最常用），负载系数为 0.75；
* `HashSet(int initialCapacity)`：构造一个新的空集； 具有指定的初始容量，负载系数为 0.75；
* `HashSet(int initialCapacity, float loadFactor)`：构造一个新的空集； 支持的 HashMap 实例具有指定的初始容量和指定的负载系数；
* `HashSet(Collection<? extends E> c)`：构造一个新集合，其中包含指定集合中的元素。

#### 4.2.2 常用成员方法

`HashSet`的常用成员方法如下：

* `boolean add(E e)`：如果指定的元素尚不存在，则将其添加到该集合中；
* `boolean contains(Object o)`：如果此集合包含指定的元素，则返回 true，否则返回 false；
* `boolean isEmpty()`：如果此集合不包含任何元素，则返回 true，否则返回 false；
* `Iterator<E> iterator()`：返回此集合中元素的迭代器；
* `boolean remove(Object o)`：从该集合中删除指定的元素（如果存在）；
* `int size()`：返回此集合中的元素数量。

更多成员方法请翻阅[官方文档](https://docs.oracle.com/en/java/javase/14/docs/api/java.base/java/util/HashSet.html)，下面我们将结合实例来介绍以上成员方法的使用。

### 4.3 实例

#### 4.3.1 新增元素

可使用`add()`方法向集中添加元素，实例如下：

```java
import java.util.HashSet;
import java.util.Set;

public class HashSetDemo1 {
    public static void main(String[] args) {
        // 实例化一个新的空集
        Set<String> hashSet = new HashSet<String>();
        // 向 hashSet 集中依次添加元素：Python、Java、PHP、TypeScript、Python
        hashSet.add("Python");
        hashSet.add("Java");
        hashSet.add("PHP");
        hashSet.add("TypeScript");
        hashSet.add("Python");
        // 打印 hashSet 的内容
        System.out.println("hashSet中的内容为：" + hashSet);
    }
}
```

运行结果：

```java
hashSet中的内容为：[TypeScript, Java, PHP, Python]
```

在实例中，我们先后向`hashSet`中添加了两次`Python`元素，由于集的元素**不可重复特性**，因此集中只允许出现一个`Python`元素。我们还观察到，打印结果的元素顺序和我们添加的顺序是不同的，这验证了集的**无序特性**。

> **Tips：** 由于`HashSet`的父类`AbstractCollection`重写了`toString()`方法，因此直接打印集，可以直观地展示出集中的元素。

#### 4.3.2 删除元素

可使用`remove()`方法删除集中元素，实例如下：

```java
import java.util.HashSet;
import java.util.Set;

public class HashSetDemo2 {
    public static void main(String[] args) {
        // 实例化一个新的空集
        Set<String> hashSet = new HashSet<>();
        // 向 hashSet 集中依次添加元素：Python、Java
        hashSet.add("Python");
        hashSet.add("Java");
        // 打印 hashSet 的内容
        System.out.println(hashSet);
        // 删除 hashSet 中的 Python 元素
        hashSet.remove("Python");
        // 打印 hashSet 的内容
        System.out.println("删除 Python 元素后，hashSet中的内容为：" + hashSet);
    }
}
```

运行结果：

```java
[Java, Python]
删除 Python 元素后，hashSet中的内容为：[Java]
```

#### 4.3.3 查询元素

我们知道了`ArrayList` 通过 `get`方法来查询元素，但`HashSet`没有提供类似的`get`方法来查询元素。

这里我们介绍一个迭代器（`Iterator`）接口，所有的`Collection`都实现了`Iterator`接口，它可以以统一的方式对各种集合元素进行遍历。我们来看下`Iterator`接口的常用方法：

* `hasNaxt()` 方法检测集合中是否还有下一个元素；

* `next()`方法返回集合中的下一个元素；

* `iterator()`：返回此集合中元素的迭代器。

实例如下：

```java
import java.util.HashSet;
import java.util.Iterator;
import java.util.Set;

public class HashSetDemo3 {
    public static void main(String[] args) {
        // 实例化一个新的空集
        Set<String> hashSet = new HashSet<String>();
        // 向 hashSet 集中依次添加元素：Python、Java、PHP
        hashSet.add("Python");
        hashSet.add("Java");
        hashSet.add("PHP");
        // 打印 hashSet 的内容
        System.out.println(hashSet);

        // 获取 hashSet 中元素的迭代器
        Iterator<String> iterator = hashSet.iterator();
        System.out.println("迭代器的遍历结果为：");
        while (iterator.hasNext()) {
            System.out.println(iterator.next());
        }
    }
}
```

运行结果：

```java
[Java, PHP, Python]
迭代器的遍历结果为：
Java
PHP
Python
```

#### 4.3.4 自定义类的常用操作

请查看如下实例：

```java
import java.util.HashSet;
import java.util.Iterator;
import java.util.Set;

public class HashSetDemo4 {

    /**
     * 静态内部类：慕课网学生
     */
    static class ImoocStudent {
        private String nickname;

        private String position;

        public ImoocStudent() {
        }

        public ImoocStudent(String nickname, String position) {
            this.setNickname(nickname);
            this.setPosition(position);
        }

        public String getNickname() {
            return nickname;
        }

        public void setNickname(String nickname) {
            this.nickname = nickname;
        }

        public String getPosition() {
            return position;
        }

        public void setPosition(String position) {
            this.position = position;
        }

        @Override
        public String toString() {
            return "ImoocStudent{" +
                    "nickname='" + nickname + '\'' +
                    ", position='" + position + '\'' +
                    '}';
        }
    }

    public static void main(String[] args) {
        Set<ImoocStudent> hashSet = new HashSet<>();
        // 实例化3个慕课网学生对象
        ImoocStudent imoocStudent1 = new ImoocStudent("Colorful", "服务端工程师");
        ImoocStudent imoocStudent2 = new ImoocStudent("Lillian", "客户端工程师");
        ImoocStudent imoocStudent3 = new ImoocStudent("小慕", "架构师");
        // 新增元素
        hashSet.add(imoocStudent1);
        hashSet.add(imoocStudent2);
        hashSet.add(imoocStudent3);
        // 使用Iterator遍历hashSet
        Iterator<ImoocStudent> iterator = hashSet.iterator();
        System.out.println("迭代器的遍历结果为：");
        while (iterator.hasNext()) {
            System.out.println(iterator.next());
        }
        // 查找并删除
        if (hashSet.contains(imoocStudent1)) {
            hashSet.remove(imoocStudent1);
        }
        System.out.println("删除nickname为Colorful的对象后，集合元素为：");
        System.out.println(hashSet);
    }
}
```

运行结果：

```java
迭代器的遍历结果为：
ImoocStudent{nickname='Lillian', position='客户端工程师'}
ImoocStudent{nickname='Colorful', position='服务端工程师'}
ImoocStudent{nickname='小慕', position='架构师'}
删除nickname为Colorful的对象后，集合元素为：
[ImoocStudent{nickname='Lillian', position='客户端工程师'}, ImoocStudent{nickname='Colorful', position='服务端工程师'}, ImoocStudent{nickname='小慕', position='架构师'}]
```

为了方便演示，我们定义了一个静态内部类`ImoocStudent`，它有两个属性`nickname`和`position`，定义了属性的`getter`和`setter`，并重写了`toString()`方法。在`main()`方法中，我们实现了自定义类在`HashSet`中的增删改查，使用迭代器可以遍历元素。

## 5. Map 集合

### 5.1 概念和特性

我们已经知道`Map`是以键值对（key-value）的形式存储的对象之间的映射，`key-value`是以`java.util.Map.Entry`类型的对象实例存在。

可以使用键来查找值，一个映射中不能包含重复的键，但值是可以重复的。每个键最多只能映射到一个值。

### 5.2 HashMap 实现类

`HashMap`是`java.util.Map`接口最常用的一个实现类，前面所学的`HashSet`底层就是通过`HashMap`来实现的，`HashMap`允许使用`null`键和`null`值。

#### 5.2.1 构造方法

* `HashMap()`：构造一个新的空映射；默认的初始容量为 16（最常用），负载系数为 0.75；

* `HashMap(int initialCapacity)`：构造一个新的空映射； 具有指定的初始容量，负载系数为 0.75；

* `HashMap(int initialCapacity, float loadFactor)`：构造一个新的空映射； 支持的 `HashMap` 实例具有指定的初始容量和指定的负载系数；

* `HashSet(Map<? extends K, ? extends V> m)`：构造一个新映射，其中包含指定映射相同。

#### 5.2.2 常用成员方法

* `void clear()`：从该映射中删除所有映射；
* `Set<Map, Entry<K, V>> entrySet`：返回此映射中包含的映射的集合；
* `V get(Object key)`：返回指定键映射到的值，如果该映射不包含键的映射，则返回 null；
* `Set<K> keySet`：返回此映射中包含的**键**的结合；
* `V put(K key, V value)`：将指定值与此映射中指定键关联；
* `V remove(Object key)`：如果存在，则从此映射中删除指定键的映射。
* `Collection<V> values`：返回此映射中包含的集合。

### 5.3 实例

下面我们使用 `HashMap` 来实现一个英汉字典的例子。

```java
import java.util.HashMap;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

public class HashMapDemo1 {

    public static void main(String[] args) {
        Map<String, String> map = new HashMap<>();
        // 添加数据
        map.put("English", "英语");
        map.put("Chinese", "汉语");
        map.put("Java", "咖啡");
        // 打印 map
        System.out.println(map);
        // 删除 key 为 Java 的数据
        map.remove("Chinese");
        System.out.println("删除键为Chinese的映射后，map内容为：");
        // 打印 map
        System.out.println(map);
        // 修改元素：
        map.put("Java", "一种编程语言");
        System.out.println("修改键为Java的值后，Java=" + map.get("Java"));
        // 遍历map
        System.out.println("通过遍历entrySet方法得到 key-value 映射：");
        Set<Entry<String, String>> entries = map.entrySet();
        for (Entry<String, String> entry: entries) {
            System.out.println(entry.getKey() + " - " + entry.getValue());
        }
        // 查找集合中键为 English 对应的值
        Set<String> keySet = map.keySet();
        for (String key: keySet) {
            if (key.equals("English")) {
                System.out.println("English 键对应的值为：" + map.get(key));
                break;
            }
        }
    }
}
```

运行结果：

```java
{English=英语, Java=咖啡, Chinese=汉语}
删除键为Chinese的映射后，map内容为：
{English=英语, Java=咖啡}
修改键为Java的值后，Java=一种编程语言
通过遍历entrySet方法得到 key-value 映射：
English - 英语
Java - 一种编程语言
English 键对应的值为：英语
```

实例中，Map 的 key 是字符串类型，value 也是字符串类型。值得注意的是，我们在创建`HashMap`的时候，在`Map`类型的后面有一个`<String, String>`，分别表示映射中将要存放的 key 和 value 的类型都为 String 类型。在遍历映射的时候，我们调用了`entrySet`方法，它返回了此映射中包含的映射的集合。通过键查找值，我们可以调用`keySet`方法来获取映射中的键的集合，并且遍历这个集合即可找到对应键，通过键就可以获取值了。

## 6. 小结

本小节我们学习了 Java 的集合，它们定义在`java.util`包中，Java 中的集合主要有`Collection`和`Map`两大接口。`List`集合是元素有序并且可以重复的集合；`Set`集合是元素无序并且不可以重复的集合；`Map`是以键值对（key-value）的形式存储的对象之间的映射，它们都支持泛型。我们分别介绍了 3 个接口常用的实现类的用法。同学们要多多进行编码练习。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

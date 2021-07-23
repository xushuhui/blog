# Java 序列化与反序列化

上一小节我们学习了 Java 的输入输出流，有了这些前置知识点，我们就可以学习 Java 的序列化了。本小节将介绍什么是序列化、什么是反序列化、序列化有什么作用，如何实现序列化与反序列化，Serializable 接口介绍，常用序列化工具介绍等内容。了解序列化的用途、学会如何进行序列化和反序列化操作是本小节的重点内容。

## 1. 序列化与反序列化

> 序列化在计算机科学的数据处理中，是指将数据结构或对象状态转换成可取用格式，以留待后续在相同或另一台计算机环境中，能恢复原先状态的过程。依照序列化格式重新获取字节的结果时，可以利用它来产生与原始对象相同语义的副本。

很多编程语言自身就支持序列化操作。Java 语言提供自动序列化，序列化（`serialize`）就是将**对象**转换为**字节流**；与之相应对的，反序列化（`deserialize`）就是将字节流转换为对象。

需要注意的是，Java 序列化对象时，会把对象的**状态**保存成字节序列，对象的状态指的就是其成员变量，因此序列化的对象不会保存类的静态变量。

在 Java 中，可通过对象输出 / 输入流来实现序列化 / 反序列化操作。 `java.io`包中，提供了`ObjectInputStream`类和`ObjectOutputStream`用来序列化对象，这两个类我们将在下面介绍。下面我们来介绍一下序列化的作用。

## 2. 序列化的作用

* **序列化可以将对象的字节序列存持久化**：可以将其保存在内存、文件、数据库中（见下图）；
* **可以在网络上传输对象字节序列**；
* **可用于远端程序方法调用**。

![](//img.mukewang.com/wiki/5edee90d0933451622740802.jpg)

## 3. 实现序列化

* `ObjectOutputStream`类下的`void writeObject(Object obj)`方法用于将一个对象写入对象输出流，也就是序列化；
* `ObjectInputStream`类下的`Object readObject()`方法用于读取一个对象到输入流，也就是反序列化。

实例代码如下：

```java
import java.io.*;

public class SerializeDemo1 {

    static class Cat implements Serializable {
        private static final long serialVersionUID = 1L;

        private String nickname;

        private Integer age;

        public Cat() {}

        public Cat(String nickname, Integer age) {
            this.nickname = nickname;
            this.age = age;
        }

        @Override
        public String toString() {
            return "Cat{" +
                    "nickname='" + nickname + '\'' +
                    ", age=" + age +
                    '}';
        }
    }

    /**
     * 序列化方法
     * @param filepath 文件路径
     * @param cat 要序列化的对象
     * @throws IOException
     */
    private static void serialize(String filepath, Cat cat) throws IOException {
        // 实例化file对象
        File file = new File(filepath);
        // 实例化文件输出流
        FileOutputStream fileOutputStream = new FileOutputStream(file);
        // 实例化对象输出流
        ObjectOutputStream objectOutputStream = new ObjectOutputStream(fileOutputStream);
        // 保存cat对象
        objectOutputStream.writeObject(cat);
        // 关闭流
        fileOutputStream.close();
        objectOutputStream.close();
    }

    /**
     * 反序列化方法
     * @param filepath 文件路径
     * @throws IOException
     * @throws ClassNotFoundException
     */
    private static void deserialize(String filepath) throws IOException, ClassNotFoundException {
        // 实例化file对象
        File file = new File(filepath);
        // 实例化文件输入流
        FileInputStream fileInputStream = new FileInputStream(file);
        // 实例化对象输入流
        ObjectInputStream objectInputStream = new ObjectInputStream(fileInputStream);
        Object o = objectInputStream.readObject();
        System.out.println(o);
    }

    public static void main(String[] args) throws IOException, ClassNotFoundException {
        String filename = "C:\\Users\\Colorful\\Desktop\\imooc\\Hello.txt";
        Cat cat = new Cat("猪皮", 1);
        serialize(filename, cat);
        deserialize(filename);
    }

}
```

运行结果：

```java
Cat{nickname='猪皮', age=1}
```

上述代码中，我们定义了一个`Cat`类，它实现了`Serializable`接口，类内部有一个`private static final long serialVersionUID = 1L;`，关于这两点，我们下面紧接着就会介绍。

除了`Cat`类的定义，我们还分别封装了序列化与反序列化的方法，并在主方法中调用了这两个方法，实现了`cat`对象的序列化和反序列化操作。

在调用序列化方法后，你会发现磁盘中的`Hello.txt`文件中被`cat`对象写入了序列化后的数据：

![](//img.mukewang.com/wiki/5edee929098c53a321451333.jpg)

## 4. Seralizable 接口

被序列化的类必须是`Enum`、`Array`或`Serializable`中的任意一种类型。

如果要序列化的类不是枚举类型和数组类型的话，则必须实现`java.io.Seralizable`接口，否则直接序列化将抛出`NotSerializableException`异常。

### 4.1 serialVersionUID

`serialVersionUID` 是 Java 为每个序列化类产生的版本标识。它可以用来保证在反序列化时，发送方发送的和接受方接收的是可兼容的对象。如果接收方接收的类的 `serialVersionUID` 与发送方发送的 `serialVersionUID` 不一致，会抛出 `InvalidClassException`。

### 4.2 默认序列化机制

如果仅仅只是让某个类实现 `Serializable` 接口，而没有其它任何处理的话，那么就会使用默认序列化机制。

使用默认机制，在序列化对象时，不仅会序列化当前对象本身，还会对其父类的字段以及该对象引用的其它对象也进行序列化。同样地，这些其它对象引用的另外对象也将被序列化，以此类推。所以，如果一个对象包含的成员变量是容器类对象，而这些容器所含有的元素也是容器类对象，那么这个序列化的过程就会较复杂，开销也较大。

### 4.3 transient 关键字

在现实应用中，有些时候不能使用默认序列化机制。比如，希望在序列化过程中忽略掉敏感数据，或者简化序列化过程。下面将介绍若干影响序列化的方法。

**当某个字段被声明为 `transient` 后，默认序列化机制就会忽略该字段**。

可以尝试将实例代码中`Cat`类的成员变量`age`声明为`transient`：

```java
// 仅部分代码
static class Cat implements Serializable {
    transient private Integer age;
}
```

运行程序，我们会发现成员变量`age`没有被序列化。

## 5. 常用序列化工具

Java 官方的序列化存在很多缺点，因此，开发者们更倾向于使用优秀的第三方序列化工具来替代 Java 自身的序列化机制。

Java 官方的序列化主要体现在以下方面：

* **性能问题**：序列化后的数据相对于一些优秀的序列化的工具，还是要大不少，这大大影响存储和传输的效率；
* **繁琐的步骤**：Java 官方的序列化一定需要实现 `Serializable` 接口，略显繁琐，而且需要关注 `serialVersionUID`；
* **无法跨语言使用**：序列化的很大一个目的就是用于不同语言来读写数据。

下面列举了一些优秀的序列化工具：

* [thrift](https://github.com/apache/thrift)、[protobuf](https://github.com/protocolbuffers/protobuf) - 适用于对性能敏感，对开发体验要求不高的内部系统。
* [hessian](http://hessian.caucho.com/doc/hessian-overview.xtp) - 适用于对开发体验敏感，性能有要求的内外部系统。
* [jackson](https://github.com/FasterXML/jackson)、[gson](https://github.com/google/gson)、[fastjson](https://github.com/alibaba/fastjson) - 适用于对序列化后的数据要求有良好的可读性（转为 json 、xml 形式）。

## 6. 小结

通过本小节的学习，我们知道了序列化（`serialize`）就是将**对象**转换为**字节流**，反序列化（`deserialize`）就是将**字节流**转换为**对象**。想要实现序列化，就必须继承`Seralizable`接口，`serialVersionUID` 是 Java 为每个序列化类产生的版本标识。当某个字段被声明为 `transient` 后，默认序列化机制就会忽略该字段。学会根据自己的应用场景选择使用序列化工具。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

# Java 反射

本小节我们来学习一个 Java 语言中较为深入的概念 —— 反射（**reflection**），很多小伙伴即便参与了工作，可能也极少用到 Java 反射机制，但是如果你想要开发一个 web 框架，反射是不可或缺的知识点。本小节我们将了解到** 什么是反射**，**反射的使用场景**，不得不提的 `Class` 类，如何**通过反射访问类内部的字段、方法以及构造方法**等知识点。

## 1. 什么是反射

> Java 的反射（reflection）机制是指在程序的**运行状态中**，可以构造任意一个类的对象，可以了解任意一个对象所属的类，可以了解任意一个类的成员变量和方法，可以调用任意一个对象的属性和方法。这种**动态获取程序信息以及动态调用对象**的功能称为 Java 语言的反射机制。反射被视为动态语言的关键。

通常情况下，我们想调用一个类内部的属性或方法，需要先实例化这个类，然后通过对象去调用类内部的属性和方法；通过 Java 的反射机制，我们就可以在程序的运行状态中，动态获取类的信息，注入类内部的属性和方法，完成对象的实例化等操作。

概念可能比较抽象，我们来看一下结合示意图看一下：

![](https://xushuhui.gitee.io/image/imooc/5ed066e809e20c1022080670.jpg)

图中解释了两个问题：

1. **程序运行状态中指的是什么时刻**：`Hello.java` 源代码文件经过编译得到 `Hello.class` 字节码文件，想要运行这个程序，就要通过 JVM 的 ClassLoader （类加载器）加载 `Hello.class`，然后 JVM 来运行 Hello.class，程序的运行期间指的就是此刻；
2. **什么是反射，它有哪些功能**：在程序运行期间，可以动态获得 `Hello` 类中的属性和方法、动态完成 `Hello` 类的对象实例化等操作，这个功能就称为反射。

说到这里，大家可能觉得，在编写代码时直接通过 `new` 的方式就可以实例化一个对象，访问其属性和方法，为什么偏偏要绕个弯子，通过反射机制来进行这些操作呢？下面我们就来看一下反射的使用场景。

## 2. 反射的使用场景

Java 的反射机制，主要用来**编写一些通用性较高的代码**或者**编写框架**的时候使用。

通过反射的概念，我们可以知道，在程序的运行状态中，对于任意一个类，通过反射都可以动态获取其信息以及动态调用对象。

例如，很多框架都可以通过配置文件，来让开发者指定使用不同的类，开发者只需要关心配置，不需要关心代码的具体实现，具体实现都在框架的内部，通过反射就可以动态生成类的对象，调用这个类下面的一些方法。

下面的内容，我们将学习反射的相关 `API`，在本小节的最后，我将分享一个自己实际开发中的反射案例。

## 3. 反射常用类概述

学习反射就需要了解反射相关的一些类，下面我们来看一下如下这几个类：

* **Class**：`Class` 类的实例表示正在运行的 `Java` 应用程序中的类和接口；
* **Constructor**：关于类的单个构造方法的信息以及对它的权限访问；
* **Field**：Field 提供有关类或接口的单个字段的信息，以及对它的动态访问权限；
* **Method**：Method 提供关于类或接口上单独某个方法的信息。

字节码文件想要运行都是要被虚拟机加载的，每加载一种类，Java 虚拟机都会为其创建一个 `Class` 类型的实例，并关联起来。

例如，我们自定义了一个 `ImoocStudent.java` 类，类中包含有构造方法、成员属性、成员方法等：

```java
public class ImoocStudent {
    // 无参构造方法
    public ImoocStudent() {
    }

    // 有参构造方法
    public ImoocStudent(String nickname) {
        this.nickname = nickname;
    }

    // 昵称
    private String nickname;


    // 定义getter和setter方法
    public String getNickname() {
        return nickname;
    }

    public void setNickname(String nickname) {
        this.nickname = nickname;
    }
}
```

源码文件 `ImoocStudent.java` 会被编译器编译成字节码文件 `ImoocStudent.class`，当 Java 虚拟机加载这个 `ImoocStudent.class` 的时候，就会创建一个 `Class` 类型的实例对象：

```java
Class cls = new Class(ImoocStudent);
```

JVM 为我们自动创建了这个类的对象实例，因此就可以获取类内部的构造方法、属性和方法等 `ImoocStudent` 的构造方法就称为 `Constructor`，可以创建对象的实例，属性就称为 `Field`，可以为属性赋值，方法就称为 `Method`，可以执行方法。

## 4. Class 类

### 4.1 Class 类和 class 文件的关系

`java.lang.Class` 类用于表示一个类的字节码（.class）文件。

### 4.2 获取 Class 对象的方法

想要使用反射，就要获取某个 `class` 文件对应的 `Class` 对象，我们有 3 种方法：

1. **类名。class**：即通过一个 `Class` 的静态变量 `class` 获取，实例如下：

```java
Class cls = ImoocStudent.class;
```

1. **对象。getClass ()**：前提是有该类的对象实例，该方法由 `java.lang.Object` 类提供，实例如下：

```java
ImoocStudent imoocStudent = new ImoocStudent("小慕");
Class imoocStudent.getClass();
```

1. **Class.forName (“包名。类名”)**：如果知道一个类的完整包名，可以通过 `Class` 类的静态方法 `forName()` 获得 `Class` 对象，实例如下：

```java
class cls = Class.forName("java.util.ArrayList");
```

### 4.3 实例

```java
package com.imooc.reflect;

public class ImoocStudent {
    // 无参构造方法
    public ImoocStudent() {
    }

    // 有参构造方法
    public ImoocStudent(String nickname) {
        this.nickname = nickname;
    }

    // 昵称
    private String nickname;


    // 定义getter和setter方法
    public String getNickname() {
        return nickname;
    }

    public void setNickname(String nickname) {
        this.nickname = nickname;
    }

    public static void main(String[] args) throws ClassNotFoundException {
        // 方法1：类名.class
        Class cls1 = ImoocStudent.class;

        // 方法2：对象.getClass()
        ImoocStudent student = new ImoocStudent();
        Class cls2 = student.getClass();

        // 方法3：Class.forName("包名.类名")
        Class cls3 = Class.forName("com.imooc.reflect.ImoocStudent");
    }

}
```

代码中，我们在 `com.imooc.reflect` 包下定义了一个 `ImoocStudent` 类，并在主方法中，使用了 3 种方法获取 `Class` 的实例对象，其 `forName()` 方法会抛出一个 `ClassNotFoundException`。

### 4.4 调用构造方法

获取了 `Class` 的实例对象，我们就可以获取 `Contructor` 对象，调用其构造方法了。

那么如何获得 `Constructor` 对象？`Class` 提供了以下几个方法来获取：

* `Constructor getConstructor(Class...)`：获取某个 `public` 的构造方法；
* `Constructor getDeclaredConstructor(Class...)`：获取某个构造方法；
* `Constructor[] getConstructors()`：获取所有 `public` 的构造方法；
* `Constructor[] getDeclaredConstructors()`：获取所有构造方法。

通常我们调用类的构造方法，这样写的（以 `StringBuilder` 为例）：

```java
// 实例化StringBuilder对象
StringBuilder name = new StringBuilder("Hello Imooc");
```

通过反射，要先获取 `Constructor` 对象，再调用 `Class.newInstance()` 方法：

```java
import java.lang.reflect.Constructor;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class ReflectionDemo {
    public static void main(String[] args) throws NoSuchMethodException, InvocationTargetException, IllegalAccessException, InstantiationException {
        // 获取构造方法
        Constructor constructor = StringBuffer.class.getConstructor(String.class);
        // 调用构造方法
        Object str = constructor.newInstance("Hello Imooc");
        System.out.println(str);
    }
}
```

运行结果：

```java
Hello Imooc
```

## 5. **访问字段**

前面我们知道了如何获取 `Class` 实例，只要获取了 `Class` 实例，就可以获取它的所有信息。

### 5.1 获取字段

Field 类代表某个类中的一个成员变量，并提供动态的访问权限。`Class` 提供了以下几个方法来获取字段：

* `Field getField(name)`：根据属性名获取某个 `public` 的字段（包含父类继承）；
* `Field getDeclaredField(name)`：根据属性名获取当前类的某个字段（不包含父类继承）；
* `Field[] getFields()`：获得所有的 `public` 字段（包含父类继承）；
* `Field[] getDeclaredFields()`：获取当前类的所有字段（不包含父类继承）。

获取字段的实例如下：

```java
package com.imooc.reflect;

import java.lang.reflect.Field;

public class ImoocStudent1 {

    // 昵称 私有字段
    private String nickname;

    // 余额 私有字段
    private float balance;

    // 职位 公有字段
    public String position;

    public static void main(String[] args) throws NoSuchFieldException {
        // 类名.class 方式获取 Class 实例
        Class cls1 = ImoocStudent1.class;
        // 获取 public 的字段 position
        Field position = cls1.getField("position");
        System.out.println(position);

        // 获取字段 balance
        Field balance = cls1.getDeclaredField("balance");
        System.out.println(balance);

        // 获取所有字段
        Field[] declaredFields = cls1.getDeclaredFields();
        for (Field field: declaredFields) {
            System.out.print("name=" + field.getName());
            System.out.println("\ttype=" + field.getType());
        }
    }

}
```

运行结果：

```java
public java.lang.String com.imooc.reflect.ImoocStudent1.position
private float com.imooc.reflect.ImoocStudent1.balance
name=nickname	type=class java.lang.String
name=balance	type=float
name=position	type=class java.lang.String
```

`ImoocStudent1` 类中含有 3 个属性，其中 `position` 为公有属性，`nickname` 和 `balance` 为私有属性。我们通过`类名.class` 的方式获取了 `Class` 实例，通过调用其实例方法并打印其返回结果，验证了获取字段，获取单个字段方法，在没有找到该指定字段的情况下，会抛出一个 `NoSuchFieldException`。

调用获取所有字段方法，返回的是一个 `Field` 类型的数组。可以调用 `Field` 类下的 `getName()` 方法来获取字段名称，`getType()` 方法来获取字段类型。

### 5.2 获取字段值

既然我们已经获取到了字段，那么就理所当然地可以获取字段的值。可以通过 `Field` 类下的 `Object get(Object obj)` 方法来获取指定字段的值，方法的参数 `Object` 为对象实例，实例如下：

```java
package com.imooc.reflect;

import java.lang.reflect.Field;

public class ImoocStudent2 {

    public ImoocStudent2() {
    }

    public ImoocStudent2(String nickname, String position) {
        this.nickname = nickname;
        this.position = position;
    }

    // 昵称 私有字段
    private String nickname;

    // 职位 公有属性
    public String position;

    public static void main(String[] args) throws NoSuchFieldException, IllegalAccessException {
        // 实例化一个 ImoocStudent2 对象
        ImoocStudent2 imoocStudent2 = new ImoocStudent2("小慕", "架构师");
        Class cls = imoocStudent2.getClass();
        Field position = cls.getField("position");
        Object o = position.get(imoocStudent2);
        System.out.println(o);
    }

}
```

运行结果：

```java
架构师
```

`ImoocStudent2` 内部分别包含一个公有属性 `position` 和一个私有属性 `nickname`，我们首先实例化了一个 ImoocStudent2 对象，并且获取了与其对应的 `Class` 对象，然后调用 `getField()` 方法获取了 `position` 字段，通过调用 `Field` 类下的实例方法 `Object get(Object obj)` 来获取了 `position` 字段的值。

这里值得注意的是，如果我们想要获取 `nickname` 字段的值会稍有不同，因为它是私有属性，我们看到 `get()` 方法会抛出 `IllegalAccessException` 异常，如果直接调用 `get()` 方法获取私有属性，就会抛出此异常。

想要获取私有属性，必须调用 `Field.setAccessible(boolean flag)` 方法来设置该字段的访问权限为 `true`，表示可以访问。在 `main()` 方法中，获取私有属性 `nickname` 的值的实例如下：

```java
public static void main(String[] args) throws NoSuchFieldException, IllegalAccessException {
    // 实例化一个 ImoocStudent2 对象
    ImoocStudent2 imoocStudent2 = new ImoocStudent2("小慕", "架构师");
    Class cls = imoocStudent2.getClass();
    Field nickname = cls.getDeclaredField("nickname");
    // 设置可以访问
    nickname.setAccessible(true);
    Object o = nickname.get(imoocStudent2);
    System.out.println(o);
}
```

此时，就不会抛出异常，运行结果：

```java
小慕
```

### 5.2 为字段赋值

为字段赋值也很简单，调用 `Field.set(Object obj, Object value)` 方法即可，第一个 `Object` 参数是指定的实例，第二个 `Object` 参数是待修改的值。我们直接来看实例：

```java
package com.imooc.reflect;

import java.lang.reflect.Field;

public class ImoocStudent3 {

    public ImoocStudent3() {
    }

    public ImoocStudent3(String nickname) {
        this.nickname = nickname;
    }

    // 昵称 私有字段
    private String nickname;

    public String getNickname() {
        return nickname;
    }

    public void setNickname(String nickname) {
        this.nickname = nickname;
    }

    public static void main(String[] args) throws NoSuchFieldException, IllegalAccessException {
        // 实例化一个 ImoocStudent3 对象
        ImoocStudent3 imoocStudent3 = new ImoocStudent3("小慕");
        Class cls = imoocStudent3.getClass();
        Field nickname = cls.getDeclaredField("nickname");
        nickname.setAccessible(true);
        // 设置字段值
        nickname.set(imoocStudent3, "Colorful");
        // 打印设置后的内容
        System.out.println(imoocStudent3.getNickname());
    }

}
```

运行结果：

```java
Colorful
```

## 6. 调用方法

Method 类代表某一个类中的一个成员方法。

### 6.1 获取方法

`Class` 提供了以下几个方法来获取方法：

* `Method getMethod(name, Class...)`：获取某个 `public` 的方法（包含父类继承）；
* `Method getgetDeclaredMethod(name, Class...)`：获取当前类的某个方法（不包含父类）；
* `Method[] getMethods()`：获取所有 `public` 的方法（包含父类继承）；
* `Method[] getDeclareMethods()`：获取当前类的所有方法（不包含父类继承）。

获取方法和获取字段大同小异，只需调用以上 `API` 即可，这里不再赘述。

### 6.2 调用方法

获取方法的目的就是调用方法，调用方法也就是让方法执行。

通常情况下，我们是这样调用对象下的实例方法（以 `String` 类的 `replace()` 方法为例）：

```java
String name = new String("Colorful");
String result = name.replace("ful", "");
```

改写成通过反射方法调用：

```java
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class ReflectionDemo1 {
    public static void main(String[] args) throws NoSuchMethodException, InvocationTargetException, IllegalAccessException {
        // 实例化字符串对象
        String name = new String("Colorful");
        // 获取 method 对象
        Method method = String.class.getMethod("replace", CharSequence.class, CharSequence.class);
        // 调用 invoke() 执行方法
        String result = (String) method.invoke(name,  "ful", "");
        System.out.println(result);
    }
}
```

运行结果：

```java
Color
```

代码中，调用 `Method` 实例的 `invoke(Object obj, Object...args)` 方法，就是通过反射来调用了该方法。

其中 `invoke()` 方法的第一个参数为对象实例，紧接着的可变参数就是要调用方法的参数，参数要保持一致。

## 7. 反射应用

> **Tips：** 理解此部分内容可能需要阅读者有一定的开发经验

学习完了反射，大家可能依然非常疑惑，反射似乎离我们的实际开发非常遥远，实际情况也的确是这样的。因为我们在实际开发中基本不会用到反射。下面我来分享一个实际开发中应用反射的案例。

场景是这样的：有一个文件上传系统，文件上传系统有多种不同的方式（上传到服务器本地、上传到七牛云、阿里云 OSS 等），因此就有多个不同的文件上传实现类。系统希望通过配置文件来获取用户的配置，再去实例化对应的实现类。因此，我们一开始的思路可能是这样的（伪代码）：

```java
public class UploaderFactory {

    // 通过配置文件获取到的配置，可能为 local（上传到本地） qiniuyun（上传到七牛）
    private String uploader;

    // 创建实现类对象的方法
    public Uploader createUploader() {
        switch (uploader) {
            case "local":
                // 实例化上传到本地的实现类
                return new LocalUploader();
            case "qiniuyun":
                // 实例化上传到七牛云的实现类
                return new QiniuUploader();
            default:
                break;
        }
        return null;
    }
}
```

`createUploader()` 就是创建实现类的方法，它通过 `switch case` 结构来判断从配置文件中获取的 `uploader` 变量。

这看上去似乎没有什么问题，但试想，后续我们的实现类越来越多，就需要一直向下添加 `case` 语句，并且要约定配置文件中的字符串要和 `case` 匹配才行。这样的代码既不稳定也不健壮。

换一种思路考虑问题，我们可以通过反射机制来改写这里的代码。首先，约定配置文件的 `uploader` 配置项不再是字符串，改为类的全路径命名。因此，在 `createUploader()` 方法中不再需要 `switch case` 结构来判断，直接通过 `Class.forName(uploader)` 就可以获取 `Class` 实例，并调用其构造方法实例化对应的文件上传对象，伪代码如下：

```java
public class UploaderFactory {

    // 通过配置文件获取到的配置，实现类的包名.类名
    private String uploader;

    // 创建实现类对象的方法
    public Uploader createUploader() {
        // 获取构造方法
		Constructor constructor = Class.forName(uploader).getConstructor();
        return (Uploader) constructor.newInstance();
    }
}
```

通过反射实例化对应的实现类，我们不需要再维护 `UploaderFactory` 下的代码，其实现类的命名、放置位置也不受约束，只需要在配置文件中指定类名全路径即可。

## 8. 小结

通过本小节的学习，我们知道了反射是 Java 提供的一种机制，它可以在程序的运行状态中，动态获取类的信息，注入类内部的属性和方法，完成对象的实例化等操作。获取 `Class` 对象有 3 种方法，通过学习反射的相关接口，我们了解到通过反射可以实现一切我们想要的操作。在本小节的最后，我也分享了一个我在实际开发中应用反射的案例，希望能对大家有所启发。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

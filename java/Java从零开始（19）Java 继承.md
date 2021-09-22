# Java 继承

本小节我们将学习 Java 的继承，通过本小节的学习，你将知道**什么是继承**，**继承有什么特点**，**如何实现继承**，**方法重写**的概念和实现，**方法重写和方法重载**是比较容易混淆的概念，我们也会介绍两个概念的区别，这些都是本小节的重点，本小节的最后我们还会介绍 `super` 关键字以及 `final` 关键字。

## 1. 概念和特点

### 1.1 概念

继承是面向对象软件技术当中的一个概念。如果一个类别 B “继承自” 另一个类别 A，就把这个 B 称为 “A 的子类”，而把 A 称为 “B 的父类别” 也可以称 “A 是 B 的超类”。继承可以使得子类具有父类别的各种属性和方法，而不需要再次编写相同的代码。

Java 语言提供了类的继承机制。利用继承，新建的类可以在原有类的基础上，使用或者重写原有类的成员方法，访问原有类的成员变量。新建的类成为**子类**，而原有类为新建类的**父类**，如果 A 是 B 的父类，B 是 C 的父类，那么 C 也是 A 的子类。

### 1.2 特点

Java 中的继承为单一继承，也就是说，一个子类只能拥有一个父类，一个父类可以拥有多个子类。

另外，所有的 Java 类都继承自 `Java.lang.Object`，所以 `Object` 是所有类的祖先类，除了 Object 外，所有类都必须有一个父类。我们在定义类的时候没有显示指定其父类，它默认就继承自 `Object` 类。

子类一旦继承父类，就会继承父类所有开放的特征，不能选择性地继承父类特征。

继承体现的是类与类之间的关系，这种关系是 `is a` 的关系，也就是说满足 `A is a B` 的关系就可以形成继承关系。

下图展示了 `Object` 类、父类以及子类的树形关系：

![](https://xushuhui.gitee.io/image/imooc/5ea61b95099ac1f601500432.jpg)

紧接着我们会实现一个这样的树形关系。

## 2. 实现继承

定义父类 `SuperClass`：

```java
// 父类
class SuperClass {
  	...
}
```

在 Java 语言中，我们通过 `extends` 关键字声明一个类继承自另一个类：

```java
// 子类
class SubClass extends SuperClass {
  	...
}
```

例如，宠物猫和宠物狗都是宠物，都有昵称、年龄等属性，都有吃东西、叫喊等行为。我们可以定义一个父类：宠物类。并且宠物猫和宠物狗类都继承宠物类，继承树形图如下：

![](https://xushuhui.gitee.io/image/imooc/5ea61bef093d2bf303930400.jpg)

代码实现：

```java
public class Pet {
    private String name;  // 昵称
    private int age;      // 年龄

  	// getter 和 setter
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

    // 吃东西
    public void eat() {
        System.out.println(this.getName() + "在吃东西");
    }

  	// 叫喊
    public void shout() {
        System.out.println("宠物会叫");
    }
}

```

父类宠物类中 `name` 和 `age` 都是私有属性，而对应的 `getter`、`setter` 方法，`eat` 和 `shout` 方法都是公有方法。

宠物狗类：

```java
public class Dog extends Pet {
	// 特有属性体重
    private float weight;

  	// getter和setter
    public float getWeight() {
        return weight;
    }

    public void setWeight(float weight) {
        this.weight = weight;
    }

  	// 特有的方法 run
    public void run() {
        System.out.println("胖成了" + this.getWeight() + "斤的狗子在奔跑");
    }
}

```

宠物狗类有一个自己特有的属性 `weight`，还有一个特有的方法 `run`。

宠物猫类：

```java
public class Cat extends Pet {
    public void sleep() {
        System.out.println(this.getName() + "睡大觉zzz");
    }
}

```

宠物猫类有一个特有的方法 `sleep`，在方法中可以调用其父类 `Pet` 的 `getName` 方法。

调用类的方法：

```java
// 实例化一个宠物狗
Dog dog = new Dog();
dog.setName("欢欢");
dog.setWeight(30f);
// 调用继承自父类的公有方法
dog.eat();
// 调用其特有方法
dog.run();

// 实例化一个宠物猫
Cat cat = new Cat();
cat.setName("豆豆");
// 调用继承自父类的公有方法
cat.eat();
// 调用其特有方法
cat.sleep();
```

运行结果：

```java
欢欢在吃东西
胖成了30.0斤的狗子欢欢在奔跑
豆豆在吃东西
豆豆睡大觉zzz
```

## 3. 方法重写

### 3.1 概念

如果一个类从它的父类继承了一个方法，如果这个方法没有被标记为 `final` 或 `static`，就可以对这个方法进行重写。重写的好处是：能够定义特定于子类类型的行为，这意味着子类能够基于要求来实现父类的方法。

### 3.2 实例

在上述父类 `Pet` 中有一个 `shout` 方法，我们知道小猫和小狗的叫声是不同的，此时可以使用**方法重写**，在 `Dog` 类和 `Cat` 类中重写 `shout` 方法。

Dog 类：

```java
class Dog extends Pet{
	// 重写 shout 方法
  	public void shout() {
        System.out.println(this.getName() + "汪汪汪地叫~");
    }
}
```

Cat 类：

```java
class Cat extends Pet{
  	@Override  // 使用注解
  	public void shout() {
        System.out.println(this.getName() + "喵喵喵地叫~");
    }
}
```

> **Tips**：在要重写的方法上面，可以选择使用 `@Override` 注解，让编译器帮助检查是否进行了正确的重写。如果重写有误，编译器会提示错误。虽然这个注解不是必须的，但建议日常编码中，在所有要重写的方法上都加上 `@Override` 注解，这样可以避免我们由于马虎造成的错误。

可以使用对象实例调用其重写的方法：

```java
// 实例化一个宠物狗
Dog dog = new Dog();
dog.setName("欢欢");
// 调用重写方法
dog.shout();

// 实例化一个宠物猫
Cat cat = new Cat();
cat.setName("豆豆");
// 调用重写方法
cat.shout();
```

运行结果：

```java
欢欢汪汪汪地叫~
豆豆喵喵喵地叫~
```

### 3.3 方法重写规则

关于方法重写，有以下规则：

* 重写方法的参数列表应该与原方法完全相同；
* 返回值类型应该和原方法的返回值类型一样或者是它在父类定义时的子类型；
* 重写方法访问级别限制不能比原方法高。例如：如果父类方法声明为公有的，那么子类中的重写方法不能是私有的或是保护的。具体限制级别参考访问修饰符；
* 只有被子类继承时，方法才能被重写；
* 方法定义为 `final`，将不能被重写（`final` 关键字将在本节后面讲到）；
* 一个方法被定义为 static，将使其不能被重写，但是可以重新声明；
* 一个方法不能被继承，那么也不能被重写；
* 和父类在一个包中的子类能够重写任何没有被声明为 private 和 final 的父类方法；
* 和父类不在同一个包中的子类只能重写 non-final 方法或被声明为 public 或 protected 的方法；
* 一个重写方法能够抛出任何运行时异常，不管被重写方法是否抛出异常。然而重写方法不应该抛出比被重写方法声明的更新更广泛的已检查异常。重写方法能够抛出比被重写方法更窄或更少的异常；
* 构造方法不能重写。

### 3.4 方法重写和方法重载的区别

Java 中的方法重写（`Overriding`）是说子类重新定义了父类的方法。方法重写必须有相同的**方法名，参数列表和返回类型**。覆盖者访问修饰符的限定**大于等于**父类方法。

而方法重载（`Overloading`）发生在**同一个类**里面两个或者是多个方法的方法名相同但是参数不同的情况。

## 4. 访问修饰符

### 4.1 作用

在上一小节封装的实现中，我们使用 `private` 和 `public` 两种访问修饰符实现了对类的封装。现在终于到了详细了解访问修饰符的时候了。

为了实现对类的封装和继承，Java 提供了访问控制机制。通过访问控制机制，类的设计者可以掩盖变量和方法来达到维护类自身状态的目的，而且还可以将另外一些需要暴露的变量和方法提供给别的类进行访问和修改。

### 4.2 种类

Java 一共提供了 4 种访问修饰符：

1. **private**：私有的，只允许在本类中访问；
2. **protected**：受保护的，允许在同一个类、同一个包以及不同包的子类中访问；
3. **默认的**：允许在同一个类，同一个包中访问；
4. **public**：公共的，可以再任何地方访问。

下表按照限定能力从大到小列出了访问修饰符在不同作用域的作用范围：

|访问控制修饰符|同一个类|同一个包|不同包的子类|不同包的非子类|
|--------------|--------|--------|------------|--------------|
| private（私有的）   |✓|✕|✕|✕|
|default（默认的）    |✓|✓|✕|✕|
|protected（受保护的）|✓|✓|✓|✕|
|public（公共的）     |✓|✓|✓|✓|

## 5. super 关键字

`super` 是用在子类中的，目的是访问**直接父类**的变量或方法。注意：

* super 关键字只能调用父类的 `public` 以及 `protected` 成员；
* super 关键字可以用在子类构造方法中调用父类构造方法；
* super 关键字不能用于静态 (`static`) 方法中。

### 5.1 调用父类构造方法

父类的构造方法既不能被继承，也不能被重写。

可以使用 `super` 关键字，在子类构造方法中要调用父类的构造方法，语法为：

```java
super(参数列表)
```

例如，父类 `Pet` 中存在构造方法：

```java
public Pet(String name) {
    System.out.println("宠物实例被创建了，宠物名字为" + name);
}
```

子类 `Dog` 的构造方法中调用父类构造方法：

```java
public Dog(String name) {
  	super(name);
  	System.out.println("小狗实例被创建了");
}
```

调用 `Dog` 有参构造方法，进行实例化：

```java
new Dog("花花");
```

运行结果：

```java
宠物实例被创建了，宠物名字为花花
小狗实例被创建了
```

### 5.2 调用父类属性

子类中可以引用父类的成员变量，语法为：

```java
super.成员变量名
```

例如，在 Dog 类中调用父类的成员变量 `birthday`：

```java
class Pet {
  	protected String birthday;
}

class Dog extends Pet {
  	public Dog() {
  	    System.out.println("宠物生日：" + super.birthday);
    }
}
```

### 5.3 调用父类方法

有时候我们不想完全重写父类方法，可以使用 `super` 关键字调用父类方法，调用父类方法的语法为：

```java
super.方法名(参数列表)
```

例如，Cat 类调用父类 Pet 的 eat 方法：

```java
class Pet {
  	public void eat() {
      	System.out.println("宠物吃东西");
    }
}

class Cat extends Pet{
  	public void eat() {
      	// 在 eat 方法中调用父类 eat 方法
      	super.eat();
      	System.out.println("小猫饭量很小");
    }
}

class Test {
  	public static void main(String[] args) {
      	Cat cat = new Cat();
      	cat.eat();
    }
}
```

运行结果：

```java
宠物吃东西
小猫饭量很小
```

### 5.4 super 与 this 的对比

`this` 关键字指向**当前类对象的引用**，它的使用场景为：

* 访问当前类的成员属性和成员方法；
* 访问当前类的构造方法；
* 不能在静态方法中使用。

`super` 关键字指向**父类对象的引用**，它的使用场景为：

* 访问父类的成员属性和成员方法；
* 访问父类的构造方法；
* 不能在静态方法中使用。

另外，需要注意的是，在构造方法调用时，super 和 this 关键字不能同时出现。

## 6. final 关键字

`final` 关键字可以作用于类、方法或变量，分别具有不同的含义。在使用时，必须将其放在变量类型或者方法返回之前，建议将其放在访问修饰符和 `static` 关键字之后，例如：

```java
// 定义一个常量
public static final int MAX_NUM = 50;
```

### 6.1 final 作用于类

当 `final` 关键字用于类上面时，这个类不会被其他类继承：

```java
final class FinalClass {
  	public String name;
}

// final类不能被继承，编译会报错
public class SubClass extends FinalClass {
}
```

编译执行，将会报错：

```java
SubClass.java:1: 错误: 无法从最终FinalClass进行继承
public class SubClass extends FinalClass {
                              ^
1 个错误
```

### 6.2 final 作用于方法

当父类中方法不希望被重写时，可以将该方法标记为 `final`：

```java
class SuperClass {
  	public final void finalMethod() {
    		System.out.println("我是final方法");
    }
}

class SubClass extneds SuperClass {
  	// 被父类标记为final的方法不允许被继承，编译会报错
  	@Override
  	public void finalMethod() {
    }
}
```

编辑执行，将会报错：

```java
SubClass.java:4: 错误: SubClass中的finalMethod()无法覆盖SuperClass中的finalMethod()
    public void finalMethod() {
                ^
  被覆盖的方法为final
1 个错误
```

### 6.3 final 作用于变量

对于实例变量，可以使用 final 修饰，其修饰的变量在初始化后就不能修改：

```java
class Cat {
  	public final String name = "小花";
}
```

实例化 Cat 类，重新对 `name` 字段赋值：

```java
Cat cat = new Cat();
cat.name = "小白";
```

编译执行，将会报错：

```java
Cat.java:7: 错误: 无法为最终变量name分配值
        cat.name = "小白";
           ^
1 个错误
```

## 7. 小结

本小节我们学习了 Java 类的继承，通过类的继承，可以大大增加代码的复用性。Java 是单继承的语言，所有类的根类都是 `Object`，继承通过 `extends` 关键字实现。要注意方法重写和方法重载的区别，不要混淆。类方法和 `final` 方法不能被重写。通过 `super` 关键字可以访问父类对象成员。`final` 关键字可以作用于类、方法和变量。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

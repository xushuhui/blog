# ES6+ Class

## 1. 前言

上一节我们主要回顾了在 ES5 中使用构造函数的方式实现类了，并说明了如何实现类的继承。从上一节的讲解中，构造函数去实现类的继承还是有诸多繁琐的地方的，我们需要考虑子类和父类的关系，继承中的细节都需要自己手动处理。本节我们将要学习 ES6 中的类的基本使用和类的继承。在学习本节我们需要明确的是，ES6 中的类是基于现有语法的原型实现的，并没有引入新的面相对象的模型。它的本质还是我们上节提到的构造函数，只是让我们更加方便地使用，是基于原型的继承的语法糖。

## 2. 基本用法

### 2.1 语法

上节我们在实现类的时候说，类不能被执行只能 new 来创建实例，当时我们是在内部手动处理的。在 ES6 中天然支持这个特性：

```javascript
class Animal { }
Animal();
// Uncaught TypeError: Class constructor Animal cannot be invoked without 'new'
```

上面的代码中我们定义了一个动物类，在控制台中让其执行，会看到如上的未捕获的类型错误：意思是在没有 new 的情况下是无法调用构造函数 Animal 的。使用 class 定义的类和使用构造函数定义类，在使用上是一样的，只是创建类的方式不一样。

上节我们知道如何创建实例上的属性和原型上的属性，那么使用 class 是怎么实现的呢？class 类提供了 constructor 方法，并在 new 的时候默认执行，所以在 constructor 函数内部在 this 上绑定实例上的属性。而 原型上的方法则是在对象中直接添加属性即可，实例如下：

```javascript
class Animal {
  constructor() {
    this.type = '鸟类'
  }
  eat() {}
}
var a = new Animal();
console.log(a.hasOwnProperty('type'));	// true
console.log(a.hasOwnProperty('eat'));		// false
```

另外，在 ES7 中 class 还提供了一种方式在实例上绑定属性，这种方式不需要 this，直接使用等号在 class 类中进行赋值。

```javascript
class Animal {
  constructor() {
    this.type = '鸟类'
  }
  age='100'
  eat() {}
}
var a = new Animal();
console.log(a.hasOwnProperty('age'));	// true
```

需要注意的是，上面的等号赋值方式要在支持 ES7 的环境中才能执行。

### 2.2 get/set

当我们深入了解对象时我们就会知道属性的 getter 和 setter ，提供了 get 和 set 两个方法用于访问和设置属性。在 ES5 中有 `Object.defineProperty()` 方法可以对对象的属性进行劫持，Vue2 的底层就是使用这个 API 实现的。当然 class 类其实也是一个对象，它也可以使用 get 的方式返回属性值。如下实例：

```javascript
class Animal {
  constructor() {
    this.type = "鸟类";
    this._age = 8;
  }
  get a() {
    return this._age;
  }

  set a(newValue) {
    this._age = newValue;
  }
}

var animal = new Animal();
console.log(animal.a); // 8
animal.a = 10;
console.log(animal.a); // 10
```

上面代码中我们就使用了 get 和 set 去获取属性值和设置属性值。那我们思考一个问题，set 和 get 是自有属性还是原型上的属性呢？其实 get 和 set 还是 class 类上的一个方法，所以是原型上的方法。

```javascript
console.log(a.hasOwnProperty('a'));	// false
```

### 2.3 static

ES6 提供了用于定义静态属性和方法的关键字 `static` ，静态方法调用时不需要实例化该类，所以就不能通过实例去调用，但可以使用类直接去调用。

静态方法通常用于为一个应用程序创建工具函数，下面我们来看一个长方形类，定义一个获取长方形面积的静态方法。

```javascript
class Rectangle {
  constructor(width, height) {
    this.width = width;
    this.height = height;
  }

  static getArea(r) {
    return r.width * r.height;
  }
}

const r = new Rectangle(5, 10);

console.log(Rectangle.getArea(r));	// 50
```

## 3. 继承

### 3.1 extends

在上节构造函数中的继承我们知道，子类的构造函数中，需要我们去手动执行父构造函数并绑定 this，还需要将子类的构造函数的原型链执行父类的原型。ES6 中的继承非常简单，在创建子类时只需要使用关键字 `extends` 即可创建一个子类。

```javascript
// 父类：动物
class Animal {
  constructor(name) {
    this.name = name;
  }
  eat() {
    console.log(this.name + '会吃饭！');
  }
  static getAge() {
		console.log('获取' + this.name + '的年龄10岁了');
		return 10;
  }
}

// 子类：具体的动物——狗
class Dog extends Animal {}
```

上面的代码中子类 Owl 继承了 Animal，那这个时候我们都继承了什么呢？从上面的学习中父类中有，this 上的属性，原型上的方法和静态方法。

```javascript
var dog = new Dog('狗');

console.log('name:', dog.name);			// name: 狗
console.log('age:', Dog.getAge());	// age: 10
dog.eat();	// 狗会吃饭！
```

从上面代码打印的结果，我们知道，实例 dog 已经继承了 Animal 上的属性和方法。在父类中对 eat 方法的定义不明确，所以在子类中我们重写 eat 方法。

```javascript
class Dog extends Animal {
  eat() {
    console.log(this.name + '会吃饭！');
  }
}
var dog = new Dog('狗');
dog.eat();	// 狗喜欢吃骨头！
```

### 3.2 super

`super` 是 class 中的关键字，可以理解是父类的别名，用于调用对象的父对象上的函数。一般 `super` 有两种情况：`super` 当做函数调用；一种是 `super` 当做父类的对象使用。

第一种情况下，`super` 关键字作为函数调用，它的作用是为了绑定 this。所以子类的构造函数必须执行一次 super。默认情况下，类中不写 constructor 时，constructor 会自动执行 super， 并绑定 this 指向当前子类。

```javascript
class A {}

class B extends A {
  constructor() {
    super();
  }
}
```

上节中我们在创建子类时就去执行了父类并绑定了 this，上面代码中的 super 和 `A.call(this)` 是相同的。

第二种情况下，super 当作父类的对象来使用的，什么情况下会使用呢？当我们在子类中想使用父类的方法时可以使用 super 直接调用父类的方法即可。

```javascript
class A {
  getCount() {
    return 7;
  }
}

class B extends A {
  constructor() {
    super();
    console.log(super.getCount()); // 7
  }
}

let b = new B();
```

## 4. 小结

本节主要学习了 ES6 中 class 类的使用和相关的知识点，需要明确的是 class 类就是一个语法糖，底层还是基于现有的原型对象的继承来实现的。所以要想深入理解 ES6 的 class 就需要对 ES5 中的构造函数有深入的理解，另外，我们可以使用 babel 进行转译，得到的代码就是使用构造函数来实现的。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

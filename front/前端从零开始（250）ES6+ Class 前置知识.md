# ES6+ Class 前置知识

## 1. 前言

在早期的 JavaScript 中是没类的概念的，如果想要实现类的功能需要通过构造函数来创建，使用 prototype 来实现类的继承。对于一些高级语言如 C++、Java、python 等，都是有类的概念的，而且在这些语言中类是非常重要的。而 JavaScript 由于历史原因在设计最初就没有想要引入类的概念，随着 JavaScript 越来越多地应用到大型项目中，JavaScript 的短板就显现了。虽然，可以使用原型等方式来解决，但是还是存在各种各样的问题。

要学习 ES6 中的 class 首先要了解 ES5 中的构造函数，主要了解在构造函数中是如何实现类和继承的，了解了这些知识带你有助于后面我们更深入的理解 ES6 中的 class。

## 2. 构造函数

### 2.1 基本用法

我们知道在 ES5 中如果想创建一个实例，是通过构造函来实现的。下面我们创建一个动物的类：

```javascript
function Animal(type) {
  this.type = type || '鸟类';
}
Animal.prototype.eat = function() {
  console.log('鸟类吃虫子!')
};

var animal = new Animal();
```

上面的代码就是使用构造函数来创建一个类，这里的构造函数首字母需要大写，这是约定俗成的，不需要解释记住就行。然后使用 new 的方式来实例化一个实例。

了解构造函数后，我们要明确地知道创建的实例有两种属性，一种是自己的，一种是公用的？针对上面的代码中 type 和 eat 哪个是自己的那个是公用的呢？一般来说绑定在 this 上的是自有属性，因为在实例化一个对象后 this 是指向这个实例的；而公共属性一般认为是 prototype 上的。另外，我们可以使用 hasOwnProperty 来判断是否是自身的属性。

```javascript
console.log(animal.hasOwnProperty('type'));		// true
console.log(animal.hasOwnProperty('eat'));		// false
```

为什么要知道属性是否是自己的呢？如果能想明白这个那么就会对类的继承有个深入的理解。下面我们来看两段代码：

```javascript
var animal1 = new Animal();
var animal2 = new Animal();

console.log(animal1.type);	// 鸟类
console.log(animal2.type);	// 鸟类
animal1.type = '家禽';
console.log(animal1.type);	// 家禽
console.log(animal2.type);	// 鸟类

console.log(animal1.eat());	// 鸟类吃虫子!
console.log(animal2.eat());	// 鸟类吃虫子!
animal1.__proto__.eat = function() {
  console.log('家禽吃粮食!')
}
console.log(animal1.eat());	// 家禽吃粮食!
console.log(animal2.eat());	// 家禽吃粮食!
```

上面的代码中我们可以看出当我们对 animal1 属性 type 修改后不会影响 animal2 的 type 属性，但是我们可以通过 animal1 的原型链对原型上的 eat 方法进行修改后，这时 animal2 上的 eat 方法也被修改了。这说明在实例上修改自有属性不会影响其他实例上的属性，但是，对非自有属性进行修改时就会影响其他属性的方法。主要这样会存在一个隐患，实例可以修改类的方法，从而影响到其他继承这个类的实例。在这样的情况下我们要想实现一个完美的继承就需要考虑很多的东西了。

## 2.2 `__proto__` 、 `prototype` 、 `constructor`

在说构造函数继承之前我们需要明确几个概念： `__proto__` 、 `prototype` 、 `constructor` 这三个都是构造函数中的概念，中文的意思可以理解为 `__proto__`（原型链） 、 `prototype`（原型） 、 `constructor`（构造方法）。它们在 class 上也是存在的。想要了解它们之间的关系，我们先看下面的几段代码：

```javascript
var animal = new Animal();

animal.__proto__ === Animal.prototype;	// true
animal.__proto__.hasOwnProperty('eat');	// true

animal.constructor === animal.__proto__.constructor;	// true
```

通过上面的关系对比可以使用示意图的方式更容易理解。

![图片描述](http://img.mukewang.com/wiki/5f8dd33f092a860119200718.jpg)

通过上面的代码和示意图我们知道，原型是构造函数上的属性，实例可以通过自身的原型链查找到，并且可以修改属性。

### 2.3 继承

了解了 `__proto__` 、 `prototype` 、 `constructor` 三者的关系那么我们就要来学习一下构造函数的继承了，上面我们定义了一个动物的构造函数，但是我们不能直接去 new 一个实例，因为 new 出来的实例没有任何意义，是一个动物实例，没有具体指向。这时我们需要创建一个子类来继承它。这时可以对 Animal 类做个限制：

```javascript
function Animal(type) {
  if (new.target === Animal) {
    throw new Error('Animal 类不能被 new，只能被继承！')
  }
  this.type = type || '鸟类';
}
Animal.prototype.eat = function() {
  console.log('鸟类吃虫子!')
};

var animal = new Animal();
//VM260:3 Uncaught Error: Animal 类不能被 new，只能被继承！
```

既然不能被 new 那要怎么去继承呢？虽然不能被 new 但是我们可以去执行这个构造函数啊，比较它本质还是一个函数。执行构造函数时 this 的指向就不是当前的实例了，所以还需要对 this 进行绑定。我们定义一个子类：Owl（猫头鹰）

```javascript
function Owl() {
  Animal.call(this);
}
var owl = new Owl();
```

通过使用 call 方法在 Owl 内部绑定 this，这样实例就继承了 Animal 上 this 的属性了。但是在 Animal 的原型中还有关于 Animal 类的方法，这些方法怎么继承呢？

首先要明确的是不能使用 `Owl.prototype = Animal.prototype` 这样的方式去继承，上面也说了这会使我们对子类原型修改的方法会作用到其他子类中去。那么怎么可以实现这一继承呢？这时就需要原型链出场了，我们可以使用 Owl 原型上的原型链指向 Animal 的原型，实例 owl 根据链的查找方式是可以继承 Animal 的原型上的方法的。

```javascript
function Owl() {
  Animal.call(this);
}
Owl.prototype.__proto__ = Animal.prototype;

var owl = new Owl();
owl.eat();	// 鸟类吃虫子!
```

通过原型链的方式还是比较麻烦的，也不优雅，ES6 提供了 `setPrototypeOf()` 方法可以实现相同的效果：

```javascript
// Owl.prototype.__proto__ = Animal.prototype;
Owl.setPrototypeOf(Owl.prototype, Animal.prototype);
```

这样在子类 Owl 的原型上增加方法不会影响父类，这样也算是比较好的方式解决了子类的继承。

## 3. 小结

本节没有去学习 class 的使用，而是复习了在 ES5 中是怎么定义类的存在的，使用的是构造函数的方式来定义一个类。在类的实际应用中继承是最为关键的，通过对如何实现构造函数中的继承，复习了原型、原型链和构造方法。在构造函数的继承中，子类不能直接去 new 一个父类，因为这样没有意义。所以我们通过在子类中执行构造函数并绑定子类的 this 继承了父类的属性，再通过子类原型的原型链继承了父类原型上的属性。通过本节的学习我们更加深刻地理解构造函数在 JavaScript 中扮演什么样的角色，继而 ES6 提出了 “真正“ 意义上的类，其实本质还是通过原型的方式，下一节我们将具体学习 ES6 的 class。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

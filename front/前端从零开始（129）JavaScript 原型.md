# JavaScript 原型

> JavaScript 常被描述为一种基于原型的语言 (prototype-based language)——每个对象拥有一个原型对象，对象以其原型为模板、从原型继承方法和属性。原型对象也可能拥有原型，并从中继承方法和属性，一层一层、以此类推。这种关系常被称为原型链 (prototype chain)，它解释了为何一个对象会拥有定义在其他对象中的属性和方法。(MDN)

每个对象都有一个标签，这个标签指向他的原型对象，对象基于一种机制，可以访问到原型对象上的属性。

在标准中，一个对象的原型是使用 `[[prototype]]` 表示的，`chrome` 对其的实现是使用 `__proto__` 属性表示。

## 1. 什么是原型

### 1.1 属性的访问机制

在 JavaScript 中，除了几种基础类型，剩下的几乎都是对象。

当我们使用对象自面量创建一个对象的时候，可以访问到对象的 `toString` 方法。

```javascript
var obj = { empty: true };

console.log(obj.toString()); // 输出：[object Object]
```

在书写这个自面量的时候，并没有提供 `toString` 这个方法，却可以被成功调用。

这就涉及到了原型。

**当在访问一个对象的属性时，如果当前对象没有这个属性，就会继续往这个对象的原型对象上去找这个属性。**

**如果原型对象上没有这个属性，则继续从这个 对象 的 原型对象 的 原型对象 找这个属性。**

**这就是属性查找的机制，直到查到原型的末端，也就是 `null` ，就会停止查找，这个时候已经确定没有这个属性了，就会返回 `undefined`。**

例子中的变量 `obj` 的原型可以通过 `__proto__` 访问。

```javascript
var obj = { empty: true };

console.log(obj.__proto__);
```

在输出的原型对象中可以找到 `toString` 方法。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ed2213b09618fb219160600.jpg)

可以通过相等运算符来判断调用的 `toString` 方法是不是原型上的方法。

```javascript
var obj = { empty: true };

console.log(
  obj.toString === obj.__proto__.toString,
); // 输出：true
```

### 1.2 原型是怎么出现在一个对象上的

到这里有个问题，到底什么是原型，原型是怎么来的。

首先看一段代码：

```javascript
function Point(x, y) {
  this.x = x;
  this.y = y;
}

var point = new Point(1, 2);

console.log(point.__proto__);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ed2216e099e1d1516260354.jpg)

这样打印出来的 `point` 的原型对象，除了 `constructor` 和 `__proto__` 属性，就什么都没有了。

接下来做个改写：

```javascript
function Point(x, y) {
  this.x = x;
  this.y = y;
}

Point.prototype.info = function() {
  console.log('x: ' + this.x + ', y: ' + this.y);
};

var point = new Point(1, 2);

point.info(); // 输出："x: 1, y: 2"

console.log(point.__proto__);
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ed2218f09873d8a14380488.jpg)

这样输出的 `point` 的原型对象，就具有了一个 `info` 方法。

从这就可以看出对象的原型，和他的构造函数的 `prototype` 属性是有关的。

所有函数都具有一个 `prototype` 属性，翻译过来也是 `原型的意思`。

**当一个函数作为构造函数被调用的时候，就会把这个函数的 `prototype` 属性，作为构造函数生成的对象的原型。**

使用相等运算符，就可以验证上面这个规则：

```javascript
console.log(
  point.__proto__ === Point.prototype,
); // 输出：true
```

这就是一个对象原型的由来。

如果要知道对象由哪个构造函数生成，可以从 `constructor` 属性获取到，原型对象的 `constructor` 属性则指向这个原型所处的函数。

这一点也可以由相等运算符验证，`point` 对象的 `constructor` 属性和其原型对象下的 `constructor` 应该都指向同一个，也就是 `Point` 函数。

```javascript
console.log(
  point.constructor === point.__proto__.constructor, // 输出：true
  point.constructor === Point, // 输出：true
  point.__proto__.constructor === Point, // 输出：true
);
```

事实上对象的 `constructor` 属性就是直接从原型上继承的。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ed221c40994d01020460814.jpg)

### 1.3 原型链

前面有提到访问对象属性的机制。

```javascript
function Point(x, y) {
  this.x = x;
  this.y = y;
}

var point = new Point(1, 2);

console.log(point.toString());
```

假如要访问 `point` 对象的 `toString` 方法，首先会去 `point` 类里找，很显然是没有这个方法的。

然后回去 `point` 类的原型对象上找，也就是 `Point` 函数的 `prototype` 属性上，很显然也是没有的。

然后会再往上一层找，也就是找到了 `Point.prototype.__proto__` 上 （等同于 `point.__proto__.__proto__`），这个时候就找到了 `toString`，随后被返回并且调用。

`Point.prototype.__proto__` 其实就是 `Object.prototype`。

```javascript
console.log(
  Point.prototype.__proto__ === Object.prototype,
); // 输出：true
```

假如检查到 `Object.prototype` 还没有目标属性，则在往上就找不到了，因为 `Object.prototype.__proto__` 是 `null`。

也就是说原型查找的末端是 `null`，碰到 `null` 就会终止查找。

这些原型环环相扣，就形成了`原型链`。

有些同学会有疑问，为什么 `Point.prototype` 的原型是 `Object.prototype`。其实 `Point.prototype` 也是一个对象，可以理解成这个对象是通过 `new Object` 创建的，所以原型自然是 `Object.prototype`。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ed221d3092e0dd521041098.jpg)

## 2. **proto** 属性

在 `Chrome浏览器` 下通过访问对象的 `__proto__` 属性可以取到对象的原型对象，这是所有对象都具备的属性。

```javascript
var date = new Date();

console.log(date.__proto__);
```

`__proto__` 具有兼容性问题，因此开发中尽量不要使用到，他不在 `ES6` 之前的标准中，但是许多旧版浏览器也对他进行了实现。

在 `ES6` 中 `__proto__ 属性` 被定制成了规范。

## 3. Object.getPrototypeOf 方法

由于 `__proto__` 存在一定兼容性的问题，可以使用 `Object.getPrototypeOf` 方法代替 `__ptoto__` 属性。

```javascript
var date = new Date();

var dateProto = Object.getPrototypeOf(date);

console.log(dateProto);
console.log(dateProto === date.__proto__); // 输出：true
```

## 4. JavaScript 中没有类

在 `JavaScript` 中是没有类的概念的。

有其他面向对象开发经验的同学可能会被 `new` 关键字误导。

`JavaScript` 中采用的是原型的机制，很多文献会称其为 `原型代理`，但个人认为对于初学者使用 `原型继承` 的方式会更好理解一点，日常讨论中其实是一个意思，不需要过多纠正其说法。

`类`和`原型`是两种不同的机制。

有关于类的内容，篇幅很大，如果不熟悉但又感兴趣，可以尝试着接触一下其他面向对象的语言，如 `Python`、`Java`、`C++`。

> ES6 提供了 `class` 关键字，引入了一些类相关的概念，但其底层运行机制依然是原型这一套，所以即便是有了 `class` 关键字来帮助开发者提升开发体验，但其本质依然不是类，只是一种原型写法的语法糖。

## 5. 小结

原型的概念至关重要，利用原型的机制可以开发出更加灵活的 JavaScript 应用。

利用原型，可以很好的复用一些代码，虽然在 `JavaScript` 中没有类，但是我们可以利用原型这个特性来模拟类的实现，达到继承、多态、封装的效果，实现代码逻辑的复用，同时可以更好的组织代码结构。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

# ES6+ Reflect（二）

## 1. 前言

上一节我们学习了 Reflect 的使用和一些基本的 API，本节我们将继续学习 Reflect 的一些扩展的 API。

## 2 Reflect 扩展方法

### 2.1 Reflect.defineProperty()

`Reflect.defineProperty()` 方法会直接在一个对象上定义一个新属性，或者修改一个对象的现有属性，基本等同于 `Object.defineProperty()` 方法，唯一不同是 `Object.defineProperty()` 返回的是这个对象，`Reflect.defineProperty()` 返回的是 `Boolean` 值。

**语法：**

```javascript
Reflect.defineProperty(target, propertyKey, attributes)
```

* target：目标对象；
* propertyKey：需要定义或修改的属性的名称；
* attributes：需要定义或修改的属性的描述。

如果 target 不是 `Object`，抛出一个 `TypeError`。

```javascript
let obj = {}
Reflect.defineProperty(obj, 'a', {value: 10})  // true
obj.a;	// 10
```

`Reflect.defineProperty` 方法可以根据返回值检查属性是否被成功定义，而 `Object.defineProperty` 只能通过 `try...catch` 去捕获其中的错误，相比之下 `Reflect.defineProperty()` 方法更加方便。

```javascript
var obj = {}
var r = Reflect.defineProperty(obj, 'a', {value: 10})

if (r) {
  // 成功 todo
} else {
  // 失败 todo
}

try {
  let obj = {}
	Object.defineProperty(obj, 'a', {value: 10})
} catch(e) {
  // 如果失败，捕获的异常
}
```

### 2.2 Reflect.apply()

`**Reflect.apply()`** 通过指定的参数列表发起对目标 (target) 函数的调用。

**语法：**

```javascript
Reflect.apply(target, thisArgument, argumentsList)
```

* target：目标函数。
* thisArgument：target 函数调用时绑定的 this 对象。
* argumentsList：target 函数调用时传入的实参列表，该参数应该是一个类数组的对象。

`apply` 函数我们都知道，它可以让函数执行并可以改变 `this` 指向。

```javascript
const arr = [1, 6, 7, 10, 2, 5];
let max;
max = Math.max.apply(null, arr);
console.log(max);	// 10
```

`Reflect.apply()` 方法与

上面的代码中 `fn.apply(obj, args)` 的写法还可以写成 `Function.prototype.apply.call(func, thisArg, args)` ，`Function.prototype.apply.call(fn, obj, args)` 这和 `Reflect.apply()` 的调用时传参是一样的。都是用于绑定 `this` 对象然后执行给定函数，`Reflect` 对象则简化这种操作。

```javascript
max = Function.prototype.apply.call(Math.max, null, arr);
console.log(max);	// 10

max = Reflect.apply(Math.max, null, arr);
console.log(max);  // 10
```

`Reflect.apply()` 可以接收截取字符串的函数。

```javascript
let str = 'imooc ES6 wiki';
let newStr;

newStr = Reflect.apply(String.prototype.slice, str, [6, 9]);
console.log(newStr); // ES6

newStr = str.slice(6, 9);
console.log(newStr); // ES6
newStr = String.prototype.slice.apply(str, [6, 9]);
console.log(newStr); // ES6
```

### 2.3 Reflect.construct(target, args)

`Reflect.construct()` 和 `new` 操作符构造函数相似 ，相当于运行 `new target(...args)` ，提供了一种新的不使用 new 来调用构造函数的方法。

**语法：**

```javascript
Reflect.construct(target, argumentsList[, newTarget])
```

**参数：**

* target：被运行的目标构造函数；
* argumentsList：类数组，目标构造函数调用时的参数；
* newTarget：（可选）作为新创建对象的原型对象的 `constructor` 属性，默认值为 `target` 。

下面的两种实例化的方式是一样的。

```javascript
function Foo() {
  console.log(arguments);
}

var obj = new Foo(...args);
var obj = Reflect.construct(Foo, args);
```

`Reflect.construct()` 返回值是以 `target` 函数为构造函数，如果 `newTarget` 存在，则为 `newTarget` 。`argumentList` 为其初始化参数。

对于有没有传递第三个参数，我们可以这样理解：target 就是唯一的构造函数，但是如果传递了第三个参数，那就表示：我们的实例由两部分组成，实例上绑定在 this 上的属性部分由第一个参数的构造函数生成；不是实例上的属性部分则由第三个参数的构造函数生成。下面我们来看下具体的实例：

```javascript
class A {
  constructor(name) {
    console.log('init A class');
    this.name = name || 'Jack';
  }
  getName() {
    console.log(this.name);
    return this.name;
  }
}
class B {
	constructor(age) {
    console.log('init A class');
    this.age = age || 18;
  }
  getAge() {
    console.log(this.age);
    return this.age;
  }
}

// 使用A类作为构造函数
let a = Reflect.construct(A, ['David']);
// 使用B类作为构造函数
let b = Reflect.construct(A, ['David'], B);

console.log(a);
console.log(b);
a.getName();
b.getAge();
```

下图是上面代码的打印结果，创建实例 a 时没有第三个参数，它的原型上的 `constructor` 指向的是类 A，并且有 `getName` 方法。创建实例 b 时有第三个参数，打印的结果可以看到实例 b 原型上的 `constructor` 执行的是类 B，并且有 B 上的 `getAge` 方法。

![](https://xushuhui.gitee.io/image/imooc/5f9f6d0c0960f76d02610285.jpg)

## 3. 小结

本节主要讲解了 Reflect 扩展方法的使用

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

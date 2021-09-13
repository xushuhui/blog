# ES6+ Array.isArray()

## 1. 前言

在程序中判断数组是很常见的应用，但在 ES5 中没有能严格判断 JS 对象是否为数组，都会存在一定的问题，比较受广大认可的是借助 toString 来进行判断，很显然这样不是很简洁。ES6 提供了 `Array.isArray()` 方法更加简洁地判断 JS 对象是否为数组。

## 2. 方法详情

判断 JS 对象，如果值是 `Array`，则为 true; 否则为 false。

**语法使用：**

```javascript
Array.isArray(obj)
```

**参数解释：**

|参数|描述|
|----|----|
| obj| 需要检测的 JS 对象|

## 3. ES5 中判断数组的方法

通常使用 `typeof` 来判断变量的数据类型，但是对数组得到不一样的结果

```javascript
// 基本类型
typeof 123;  //number
typeof "123"; //string
typeof true; //boolean

// 引用类型
typeof [1,2,3]; //object
```

上面的代码中，对于基本类型的判断没有问题，但是判断数组时，返回了 object 显然不能使用 `typeof` 来作为判断数组的方法。

### 3.1 通过 instanceof 判断

`instanceof` 运算符用于检测构造函数的 `prototype` 属性是否出现在某个实例对象的原型链。

`instanceof` 可以用来判断数组是否存在，判断方式如下：

```javascript
var arr = ['a', 'b', 'c'];
console.log(arr instanceof Array); 		   // true
console.log(arr.constructor === Array;); // true
```

在解释上面的代码时，先看下数组的原型链指向示意图：

![原型链](https://img.mukewang.com/wiki/5f16615008dfbf8215400232.jpg)

数组实例的原型链指向的是 `Array.prototype` 属性，`instanceof` 运算符就是用来检测 `Array.prototype` 属性是否存在于数组的原型链上，上面代码中的 arr 变量就是一个数组，所有拥有 `Array.prototype` 属性，返回值 `true`，这样就很好的判断数组类型了。

但是，需要注意的是，`prototype` 属性是可以修改的，所以并不是最初判断为 `true` 就一定永远为真。

在我们的网站中，脚本可以拥有多个全局环境，例如 html 中拥有多个 iframe 对象，`instanceof` 的验证结果可能不会符合预期，例如：

```javascript
var iframe = document.createElement('iframe');
document.body.appendChild(iframe);

var iframeArray = window.frames[0].Array;
var arr = new iframeArray('a', 'b', 'c');

console.log(arr instanceof Array);	// false
console.log(arr)	// ["a", "b", "c"]
```

比如打开一个网站的控制台，输入上面的代码，先在 body 上创建并添加一个 iframe 对象，并把它插入到当前的网页中。这时我们可以获取 iframe 中数组构造函数。通过这个构造函数去实例化一个数组，这时再用 `instanceof` 去判断就会返回 false，但是案例中的 arr 确实是一个数组，这就是 `instanceof` 判断数组所带来的问题。

### 3.2 通过 constructor 判断

我们知道，Array 是 JavaScript 内置的构造函数，构造函数属性（prototype）的 `constructor` 指向构造函数（见下图），那么通过 `constructor` 属性也可以判断是否为一个数组。

```javascript
var arr = new Array('a', 'b', 'c');
arr.constructor === Array;	//true
```

下面我们通过构造函数的示意图来进行分析：

![图片描述](https://img.mukewang.com/wiki/5f1683a40842ccee13780616.jpg)

由上面的示意图可以知道，我们 new 出来的实例对象上的原型对象有 `constructor` 属性指向构造函数 Array，由此我们可以判断一个数组类型。

但是 `constructor` 是可以被重写，所以不能确保一定是数组，如下示例：

```javascript
var str = 'abc';
str.constructor = Array;
str.constructor === Array // true
```

上面的代码中，str 显然不是数组，但是可以把 `constructor` 指向 Array 构造函数，这样再去进行判断就是有问题的了。

`constructor` 和 `instanceof` 也存在同样问题，不同执行环境下，`constructor` 的判断也有可能不正确，可以参考 `instanceof` 的例子。

## 4. Array.isArray () 的使用

下面我们通过示例来看下 `Array.isArray()` 是怎样判断数组的。

```javascript
// 下面的函数调用都返回 true
Array.isArray([]);
Array.isArray([10]);
Array.isArray(new Array());
Array.isArray(new Array('a', 'b', 'c'))
// 鲜为人知的事实：其实 Array.prototype 也是一个数组。
Array.isArray(Array.prototype);

// 下面的函数调用都返回 false
Array.isArray();
Array.isArray({});
Array.isArray(null);
Array.isArray(undefined);
Array.isArray(17);
Array.isArray('Array');
Array.isArray(true);
Array.isArray(false);
Array.isArray(new Uint8Array(32))
Array.isArray({ __proto__: Array.prototype });
```

上面的代码中对 JavaScript 中的数据类型做验证，可以很好地区分数组类型。

## 5. 自定义 isArray

在 ES5 中比较通用的方法是使用 `Object.prototype.toString` 去判断一个值的类型，也是各大主流库的标准。在不支持 ES6 语法的环境下可以使用下面的方法给 `Array` 上添加 `isArray` 方法

```javascript
if (!Array.isArray){
  Array.isArray = function(arg){
    return Object.prototype.toString.call(arg) === '[object Array]';
  };
}
```

## 6. 小结

本节介绍了判断一个值是数组类型的方法 `Array.isArray()` 此方法可以很准确地判断数组，学习了在 ES5 中判断数组类型的几个方法的缺陷。在不支持 ES6 的情况下也可以通过 `Object.prototype.toString` 自定义 `Array.isArray()` 方法。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

# ES6+ Number 对象的方法

## 1. 前言

本节主要讲解 Number 对象下剩余的方法，Number 对象把之前在全局上的方法移植了过来，并对有缺陷的方法进行了补充和修复，上节我们已经学习了全局下的 `isFinite()` 和 `isNaN()` 两种方法存在类型转换，对于判断数值来说存在歧义。下面我们来看看，ES6 中移植的没有变的函数和新增的一些函数。

## 2. Number.parseInt()&Number.parseFloat()

为了保持方法上的统一，把全局下的 `parseInt()` 和 `parseFloat()` 移植到 ES6 的 Number 对象上。

ES6 的 Number 对象上提供的 Number.isFinite () 和 Number.isNaN () 两个函数是相同的，如何证明 Number 下的这两个方法只是移植全局的呢？可以利用 `===` 运算符来判断，如下实例：

```javascript
Number.parseInt === parseInt;				// true
Number.parseFloat === parseFloat;		// true
```

上面的代码返回的结果都为 `ture` 说明此两种函数和全局是一样的，没有发生变化。具体的使用方法可以参考 ES5 中的 `parseInt()` 和 `parseFloat()` 函数。

```javascript
// ES5的写法
parseInt('12.34') 						// 12
parseFloat('123.45#') 				// 123.45

// ES6的写法
Number.parseInt('12.34') 			// 12
Number.parseFloat('123.45#')  // 123.45
```

将这两个全局方法移植到 Number 对象上，为了逐步减少全局性方法，使语言逐步模块化。

## 3. Number.isInteger()

在学习这个函数之前，我们先来回顾一下，我们是怎么判断一个值为整数的？

### 3.1 判断一个值为整数

一种方法是：任何整数都会被 1 整除，即余数是 0。利用这个规则来判断是否是整数。就有如下函数：

```javascript
function isInteger(value) {
	return typeof value === 'number' && value%1 === 0;
}
isInteger(5) 		// true
isInteger(5.5) 	// false
isInteger('') 	// false
isInteger('8') 	// false
isInteger(true) // false
isInteger([]) 	// false
```

另一种方法是：使用 `Math.round`、`Math.ceil`、`Math.floor` 判断，因为整数取整后还是等于自己。利用这个特性来判断是否是整数，使用 `Math.floor` 示例，如下：

```javascript
function isInteger(value) {
	return Math.floor(value) === value;
}
isInteger(5) 		// true
isInteger(5.5) 	// false
isInteger('') 	// false
isInteger('8') 	// false
isInteger(true) // false
isInteger([]) 	// false
```

上面的两种方法算是比较常用的判断方式，其他的一些方式都存在一些问题，这里就不一一列举了。但是，这两种方法都不够简洁，ES6 把判断整数提升到了语言层面，下面我们来看下 `Number.isInteger()` 的使用。

### 3.2 Number.isInteger () 的用法

`Number.isInteger()` 是 ES6 新增的函数，用来判断给定的参数是否为整数。

```javascript
Number.isInteger(25) // true
Number.isInteger(25.1) // false
```

如果被检测的值是整数，则返回 `true`，否则返回 `false`。注意 `NaN` 和正负 `Infinity` 不是整数。

```javascript
Number.isInteger(0);         // true
Number.isInteger(1);         // true
Number.isInteger(-100000);   // true

Number.isInteger(0.8);       // false
Number.isInteger(Math.PI);   // false

Number.isInteger(Infinity);  // false
Number.isInteger(-Infinity); // false
Number.isInteger("100");     // false
Number.isInteger(true);      // false
Number.isInteger(false);     // false
Number.isInteger([1]);       // false
```

上面的代码基本涵盖了 JavaScript 中的值的判断，在一些不支持 ES6 语法的浏览器中可以使用上面的两种方式进行 Polyfill 处理。

## 4. Number.isSafeInteger()

**`Number.isSafeInteger()`** 是 ES6 新增的函数，用来判断传入的参数值是否是一个 “安全整数”（safe integer）在数值扩展的 [小节](https://www.imooc.com/wiki/ES6lesson/numberextend.html) 我们介绍了最大安全整数和最小安全整数，不记得的同学可以跳过去看看。

一个安全整数是一个符合下面条件的整数：

* 可以准确地表示为一个 IEEE-754 双精度数字；
* 其 IEEE-754 表示不能是舍入任何其他整数以适应 IEEE-754 表示的结果。

比如，`2e53 - 1` 是一个安全整数，它能被精确表示，在任何 IEEE-754 舍入模式（rounding mode）下，没有其他整数舍入结果为该整数。作为对比，`2e53` 就不是一个安全整数，它能够使用 IEEE-754 表示，但是 `2e53 + 1` 不能使用 IEEE-754 直接表示，在就近舍入（round-to-nearest）和向零舍入中，会被舍入为 `2e53`。

安全整数范围为 `-(2e53 - 1)到``2e53 - 1` 之间的整数，包含 `-(2e53 - 1)和``2e53 - 1`。

```javascript
Number.isSafeInteger(3);                    // true
Number.isSafeInteger(Math.pow(2, 53))       // false
Number.isSafeInteger(Math.pow(2, 53) - 1)   // true
Number.isSafeInteger(NaN);                  // false
Number.isSafeInteger(Infinity);             // false
Number.isSafeInteger("3");                  // false
Number.isSafeInteger(3.1);                  // false
Number.isSafeInteger(3.0);                  // true
```

## 5. 小结

本节学习了 Number 对象下的方法，讲解了为什么把全局的方法移植到 Number 对象下，以及对比没有 ES6 时是怎么判断数值为整数的情况，通过对 Number 对象下的方法的学习，可以看到 ES6 在收敛全局的方法，使语言逐步模块化，更加符合语言的规范。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

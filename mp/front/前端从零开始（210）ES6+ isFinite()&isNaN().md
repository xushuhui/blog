# ES6+ Number.isFinite()&Number.isNaN()

## 1. 前言

在 ES5 中，全局下的 `isFinite ()` 和 `isNaN ()` 两种方法存在类型转换，对最终的判断结果存在歧义。ES6 在 Number 对象上，新提供了 `Number.isFinite ()` 和 `Number.isNaN ()` 两个方法，更加健壮地进行数值的判断，接下来让我看看这两种方法。

## 2. Number.isFinite()

在 ES5 中有全局的 `isFinite()` 函数用来判断被传入的参数值是否为一个有限的数值，如果参数是字符串，会首先转为一个数值，然后在进行验证。

```javascript
isFinite(Infinity);  // false
isFinite(NaN);       // false
isFinite(-Infinity); // false

isFinite(0);         // true
isFinite(2e64);      // true
isFinite('2e64');    // true

isFinite("0");       // true
```

上面的代码可以看出，字符串也会被先转为数值再进行判断，而 ES6 Number 对象上提供的 `isFinite()` 更健壮，和全局的 `isFinite()` 函数相比，这个方法不会强制将一个非数值的参数转换成数值，这就意味着，**只有数值类型的值，且是有穷的（finite），才返回 `true`**。

```javascript
Number.isFinite(Infinity);  // false
Number.isFinite(NaN);       // false
Number.isFinite(-Infinity); // false

Number.isFinite(0);         // true
Number.isFinite(2e64);      // true
Number.isFinite('2e64');	// false

Number.isFinite('0');       // false
```

## 3. Number.isNaN()

在 JavaScript 中与其它的值不同，`NaN` 不能通过相等操作符（== 和 ===）来判断 ，因为 `NaN == NaN` 和 `NaN === NaN` 都会返回 `false`。 因此，判断一个值是不是 `NaN` 是有必要的。

### 3.1 NaN 值的产生

当算术运算的结果返回一个未定义的或无法表示的值时，`NaN` 就产生了。但是，`NaN` 并不一定用于表示某些值超出表示范围的情况。

* 将某些非数值强制转换为数值的时候，会得到 `NaN`。

* 0 除以 0 会返回 NaN —— 但是其他数除以 0 则不会返回 NaN。

[上一节](https://www.imooc.com/wiki/ES6lesson/numberextend.html) 我们知道可以使用 `Number()` 方法进行类型转换，下面列举被强制类型转换为 `NaN` 的例子：

```javascript
Number(undefined)				// NaN
Number('undefined')				// NaN
Number('string')				// NaN
Number({})						// NaN
Number('10,3')					// NaN
Number('123ABC')				// NaN
Number(new Date().toString())	// NaN
```

上面的例子可以看出，很多值在强制类型转换下转为 `NaN`，针对这样的值去进行判断无疑是有问题的，下面我们来看下 `isNaN ()` 的问题。

### 3.2 isNaN () 的问题

默认情况全局下存在方法 `isNaN ()` 用了判断是否为 `NaN` 值，它要求接收的是数值类型的参数，但是当参数不是 `Number` 类型， `isNaN` 函数会首先尝试将这个参数转换为数值，然后才会对转换后的结果是否是 `NaN` 进行判断。

实例：

```javascript
isNaN(NaN);       // true
isNaN(undefined); // true
isNaN('undefined')// true
isNaN({});        // true

isNaN(true);      // false
isNaN(null);      // false
isNaN(37);        // false

// strings
isNaN("37");      // false: 可以被转换成数值37
isNaN("37.37");   // false: 可以被转换成数值37.37
isNaN("37,5");    // true
isNaN('123ABC');  // true:  parseInt("123ABC")的结果是 123, 但是Number("123ABC")结果是 NaN
isNaN("");        // false: 空字符串被转换成0
isNaN(" ");       // false: 包含空格的字符串被转换成0

// dates
isNaN(new Date());                // false
isNaN(new Date().toString());     // true

isNaN("imooc")   // true: "blabla"不能转换成数值
                 // 转换成数值失败， 返回NaN
```

结合上面 NaN 是如何产生的例子的结果可以看出，使用 `isNaN` 来判断返回的是 `true`，这显然不是我们想要的结果。针对这样的问题，ES6 做了修补，下面我们看 ES6 中的 `isNaN` 方法。

### 3.3 Number.isNaN () 详情

ES6 提供了 `Number.isNaN(x)` ，通过这个方法来检测变量 `x` 是否是一个 `NaN` 将会是一种可靠的做法，它不会对所判断的值进行强制类型转换。

```javascript
Number.isNaN(NaN);        // true
Number.isNaN(Number.NaN); // true
Number.isNaN(0 / 0)       // true

// 下面这几个如果使用全局的 isNaN() 时，会返回 true。
Number.isNaN("NaN");      // false，字符串 "NaN" 不会被隐式转换成数字 NaN。
Number.isNaN(undefined);  // false
Number.isNaN('undefined');// false
Number.isNaN({});         // false
Number.isNaN("blabla");   // false

Number.isNaN(true);   	 // false
Number.isNaN(null);   	 // false
Number.isNaN(37);   	 // false
Number.isNaN("37");   	 // false
Number.isNaN("37.37");	 // false
Number.isNaN("");   	 // false
Number.isNaN(" ");   	 // false
```

通过上面的实例，基本覆盖了现有程序的所有情况，不会出现使用全局下的 `isNaN()` 多带来的问题。所有推荐使用 `Number.isNaN(x)` 方式来判断是否是 `NaN`。在不支持 `Number.isNaN` 函数情况下，可以通过表达式 `(x != x)` 来检测变量 `x` 是不是 `NaN` 会更加可靠。

## 4. 小结

本节中传统的全局方法 `isFinite()` 和 `isNaN()` 的区别在于，传统方法先调用 Number () 将非数值的值转为数值，再进行判断，而这两个新方法只对数值有效，`Number.isFinite()` 对于非数值一律返回 false，`Number.isNaN()` 只有对于 `NaN` 才返回 `true`，非 `NaN` 一律返回 `false`。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

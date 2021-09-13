# 三元运算符

> 条件（三元）运算符是 JavaScript 仅有的使用三个操作数的运算符。一个条件后面会跟一个问号（?），如果条件为 truthy ，则问号后面的表达式 A 将会执行；表达式 A 后面跟着一个冒号（:），如果条件为 falsy ，则冒号后面的表达式 B 将会执行。本运算符经常作为 if 语句的简捷形式来使用。（MDN）

三元运算符，也称条件运算符、三目运算符。

三元运算符可以代替简单的 if 语句。

## 1. 语法

```javascript
条件 ? 条件为真时执行的表达式 : 条件为假时执行的表达式
```

当条件成立或者不成立的时候，会执行对应的表达式，并将表达式的结果作为三元运算的结果。

利用三元运算符判断年龄是否成年获取对应的文案：

```javascript
var age = 19;

var str = age > 19 ? '成年了' : '没有成年';

console.log(str); // 输出："成年了"

// if 的写法2
var age = 19;

var str = '';

if (age > 19) {
  str = '成年了';
} else {
  str = '没有成年';
}

console.log(str);

// if 的写法2
var age = 19;
var str = '没有成年';

if (str > 19) {
  str = '成年了';
}

console.log(str);
```

使用三元运算符可以代替简单的 if 语句，让代码更简洁，减少分支。

其中条件为假的情况除了 `false` ，有以下几种：

* null
* undefined
* NaN
* 0（数字 0)
* 空字符串

这些值有在 `Boolean` 中提到过，他们都可以被隐式转换为 `false`。

## 2. 注意点

### 2.1 尽量不要嵌套

嵌套的三元运算符会让可读性变差

比如：如果下班回来看到卖水果的就买一斤西瓜，如果有桃子，就只买桃子。

```javascript
var hasFruit = true;
var hasPeach = false;

var buy = hasFruit ? hasPeach ? '买桃子' : '买西瓜' : '没有卖水果的';
```

这样会让逻辑变得混乱，可读性变差。

这种情况应使用 if 代替。

```javascript
var hasFruit = true;
var hasPeach = false;

var buy = '';

if (hasFruit) {
  if (hasPeach) {
    buy = '买桃子';
  } else {
    buy = '买西瓜';
  }

  // 或者这里使用三元运算符
  // buy = hasPeach ? '买桃子' : '买西瓜';
} else {
  buy = '没有卖水果的';
}
```

换成 if 可以让逻辑更清晰。

有时候不能为了让代码看起来变少而牺牲代码可读性，代码可读性是非常重要的。

### 2.2 避免使用不必要的三元运算符

如判断是否成年，结果需要布尔值：

```javascript
var age = 11;

var isAdult = age >= 18 ? true : false;

console.log(isAdult);
```

第一眼看到这串代码，可能会觉得没有问题。

在开发中经常会遇到这样的代码，有多年开发经验的老司机也可能会这样写，但其实这里没有必要使用三元运算符，因为 `age >= 18` 这里的比较运算符返回的就是布尔值。

在开发中应该避免这样不必要三元运算。

```javascript
var age = 11;

var isAdult = age >= 18;

console.log(isAdult);
```

## 3. 小结

三元运算符可以代替简单的 if 语句，但尽量不要嵌套使用，这样会降低代码的可读性。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

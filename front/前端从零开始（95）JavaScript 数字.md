# JavaScript 数字

> 基于 IEEE 754 标准的双精度 64 位二进制格式的值（-(253 -1) 到 253 -1）。——MDN

在 JavaScript 只有浮点数（可以理解成带有小数点的数）。

数字就是我们日常使用的数字，在 JavaScript 中默认是十进制的。

```javascript
10;
-1;
1.1;
10.0;
+0;
-0;
99999999;
+1;
```

正数（大于零的数）通常不需要在前面带上 `+` 号， `+1` 和 `1` 是等价的。

使用 `+` 号很多情况下是为了把字符串转换成数字：

```javascript
var num = '123';

num = +num;

console.log(num, typeof(num));
```

typeof 可以简单理解成返回数据的类型，这里返回的是 `number` ，即表示为数字。

## 1. 其他进制数字

### 1.1 二进制

使用 `0b` 开头的数字，就表示二进制。

可以在控制台直接输入内容进行调试。

```javascript
0b1012; // 报错，因为二进制只有0与1

0b1111; // 15
```

### 1.2 八进制

`0` 开头的则为八进制。

```javascript
09; // 控制台输出的9是10进制的，因为八进制中没有9，所以把09解析成了十进制中的9

010; // 8
```

### 1.3 十六进制

十六进制用到的相对会多一些，例如在 `three.js` 中就会用到大量的使用十六进制表示的颜色。

```javascript
0xffffff; // 用来表示白色


0xa; // 10
```

### 1.4 进制转化

使用 `toString` 方法，在把数字转换成字符串的同时，可以做进制的转换。

```javascript
(10).toString(2); // "1010"
(8).toString(8); // "10"
```

在调用 `toString` 的时候，可以给定一个基数作为要转换到的进制，范围是 2 到 36 ，默认为 10 。

使用 parseInt 方法，也可以实现进制转换。

如将二进制数 `1010` 转换成十进制数字：

```javascript
parseInt('1010', 2); // 10
```

parseInt 的第二个参数为进制基数，表示第一个参数是几进制的，返回值为 10 进制。

## 2. 最大值与最小值

在 `ES6` 中，提供了两个常量 `Number.MAX_SAFE_INTEGER` 与 `Number.MIN_SAFE_INTEGER` ，可以获得到在 JavaScript 可以表示的最大值与最小值（安全数）。

安全整数范围为 -（253 - 1）到 253 - 1 之间的整数，包含 -（253 - 1）和 253 - 1 。

```javascript
console.log(Number.MAX_SAFE_INTEGER); // 输出：9007199254740991
console.log(Number.MIN_SAFE_INTEGER); // 输出：-9007199254740991
```

在 ES6 中没必要自己去做比较，可以使用 `Number.isSafeInteger()` 方法判断传入的数字是不是一个安全数。

```javascript
Number.isSafeInteger(11111111111111111111); // false

Number.isSafeInteger(996); // true
```

## 3. 0.1 + 0.2 不等于 0.3 问题

观察控制台可以看到， `0.1 + 0.2 === 0.3` 是不成立的。

![图片描述](https://img.mukewang.com/wiki/5e7a2ab90a4d084915320432.jpg)

在计算机中，所有的内容都是有二进制表示的，数字也不例外。

0.1 用二进制表示就是 0.00011001100110011(0011)… 之后就是括号内 0011 的循环。

而 0.2 采用二进制表示就是 0.001100110011(0011)…只有也是 0011 的循环。

而 `JavaScript` 采用的数字规范最大只能操作 64 位的数字，也就是说虽然 `0.1` 与 `0.2` 表示成二进制是会无限循环的，但是在计算的时候终究要截断。

其中这 64 位的分布，一位为符号位，整数占据十一位，剩下的五十二位都为小数位。

这样就可以得到最后参与运算的值：

`0.1` => 2-4 * 1.1001100110011001100110011001100110011001100110011010

`0.2` => 2-3 * 1.1001100110011001100110011001100110011001100110011010

其实到这里已经可以理解为什么是不一样的了，因为精度的关系，无限循环的小数遭遇了截断，而最终结果是截断后的数据相加获得的，那必定不会再等于`0.3`。

> 注意：这是一个常考的面试题。

## 4. Infinity

Infinity 表示正无限大，也有负无限大 `-Infinity` 。

将 `0` 作为除数的时候会得到 `Infinity` 。

```javascript
10 / 0; // Infinity
-10 / 0; // -Infinity
```

比较有趣的是，当 `Infinity` 作为除数的时候，结果就会是 0 。

```javascript
10 / Infinity; // 0
10 / -Infinity; // -0
```

## 5. NaN

NaN 即 `Not A Number` ，表示这个值不是一个数字。

`NaN` 基本不会被用到，通常会在两种情况下得到：

1. 计算错误
2. 将字符串或其他类型转换成数字失败

```javascript
1 - 'a'; // NaN

parseInt('a'); // NaN
```

## 6. 其他小特性

### 6.1 -0 等于 +0

0 是唯一一个正负相等的数字。

```javascript
+0 === -0; // true
```

### 6.2 NaN 不等于 NaN

NaN 是唯一一个自己不等于自己的值。

```javascript
NaN === NaN; // false
```

根据这一个特性，要判断一个值是不是 NaN 就非常容易：

```javascript
function isNaN(val) {
  return val !== val;
}

console.log(isNaN(1)); // 输出：false
console.log(isNaN(NaN)); // 输出： true
console.log(isNaN(1 - 'a')); // 输出：true
```

### 6.3 window.isNaN 与 Number.isNaN

在 window 下有一个 `isNaN` 方法，在 Number 下也有一个 `isNaN` 方法。

这两个方法区别在于， `window.isNaN` 会将传入的参数转换成数字，转换结果为 `NaN` 就会返回 true ，而后者只有当传入的值确定为 `NaN` 的时候，才会返回 true 。

这两个方法可以根据需求选择。

### 7. 小结

在 JavaScript 中，数字都是带有小数位的数，并且范围是有限的。

因为精度问题，JavaScript 的小数计算并不精确，这是非常需要注意的一点，同时这也是面试中常见的问题。

> 在新的 ECMAScript 版本中，提供了 [BigInt](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/BigInt) 数据类型，可以表示无限大的整数。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

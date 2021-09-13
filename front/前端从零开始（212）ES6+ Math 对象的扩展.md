# ES6+ Math 对象的扩展

## 1. 前言

在编程中遇到数学问题时一般会借助 `Math` 这个 JavaScript 的内置对象，它拥有一些数学常数属性和数学函数方法。`Math` 不是一个函数对象，它主要用于 `Number` 类型。

与其他全局对象不同的是，`Math` 不是一个构造器。`Math` 的所有属性与方法都是静态的。比如：

* 圆周率的写法是 `Math.PI`，
* 调用正余弦函数的写法是 `Math.sin(x)`，`x` 是要传入的参数。

`Math` 的常量是使用 JavaScript 中的全精度浮点数来定义的。本节我们主要学习，ES6 对 Math 对象的扩展。

## 2. Math.trunc

`Math.trunc()` 方法会将数字的小数部分去掉，只保留整数部分，是一个取整操作。

`Math` 中还有三个方法： `Math.floor()`、`Math.ceil()`、`Math.round()` ，也是用于取整操作的，但是他们有一定逻辑的：

* `Math.floor()` 向下取整；
* `Math.ceil()` 向上取整；
* `Math.round()` 进行四舍五入操作。

相比之下 `Math.trunc()` 的执行逻辑就很简单，仅仅是删除掉数字的小数部分和小数点，不管参数是正数还是负数。

**使用语法：**

```javascript
Math.trunc(value)
```

**参数说明：**

|参数|描述|
|----|----|
|value|可以是任意数字，如果是非数值则会被隐式转换为数字类型|

**实例**

基本实例：

```javascript
Math.trunc(5.3) // 5
Math.trunc(5.9) // 5
Math.trunc(-5.3) // -5
Math.trunc(-5.9) // -5
Math.trunc(-0.12345) // -0
```

上面的代码中，第 5 行虽然结果是 0 但是 0 前面的负号还会被保留。

传入该方法的参数会被隐式转换成数字类型，也就是对于非数值，`Math.trunc()` 内部使用 `Number()` 方法将其先转为数值。

```javascript
Math.trunc('123.456') // 123
Math.trunc(true) //1
Math.trunc(false) // 0
Math.trunc(null) // 0
```

对于空值和无法截取整数的值，返回 NaN。

```javascript
Math.trunc(); // NaN
Math.trunc('imooc'); // NaN
Math.trunc(undefined) // NaN
Math.trunc(NaN); // NaN
```

## 3. Math.sign()

`Math.sign()` 方法用来判断一个数到底是正数、负数、还是零。对于非数值，会先将其转换为数值。

**使用语法：**

```javascript
Math.sign(value)
```

**参数说明：**

|参数|描述|
|----|----|
|value|可以是任意数字，如果是非数值则会被隐式转换为数字类型|

`Math.sign()` 会返回五种值。

* 参数为正数，返回 + 1；
* 参数为负数，返回 - 1；
* 参数为 0，返回 0；
* 参数为 - 0，返回 - 0;
* 其他值，返回 NaN。

**实例**

```javascript
Math.sign(7); // 1
Math.sign(-7.5); // -1
Math.sign("-7.5"); // -1
Math.sign(0); // 0
Math.sign(-0); // -0
```

对应 -0 转换后的结果会保留负号，如果参数是非数值，会自动转为数值。对于那些无法转换为数值的值，会返回 NaN。

```javascript
Math.sign(NaN); // NaN
Math.sign("foo"); // NaN
Math.sign(); // NaN
```

## 4 Math.cbrt()

`Math.cbrt()` 函数返回任意数字的立方根。`cbrt` 是 “cube root” 的缩写，意思是立方根。

**使用语法：**

```javascript
Math.cbrt(value)
```

**参数说明：**

|参数|描述|
|----|----|
|value|可以是任意数字，如果是非数值则会被隐式转换为数字类型|

对于非数值，`Math.cbrt()` 方法内部也是先使用 `Number` 方法将其转为数值。无法转换为数值的返回 NaN。

**实例：**

```javascript
Math.cbrt('8'); // 2
Math.cbrt('imooc'); // NaN
Math.cbrt(NaN); // NaN
Math.cbrt(-1); // -1
Math.cbrt(-0); // -0
Math.cbrt(-Infinity); // -Infinity
Math.cbrt(0); // 0
Math.cbrt(1); // 1
Math.cbrt(Infinity); // Infinity
Math.cbrt(null); // 0
Math.cbrt(2); // 1.2599210498948734
```

## 5. Math.clz32()

JavaScript 的整数使用 32 位二进制表示，`Math.clz32()` 方法返回一个数的 32 位无符号整数形式有多少个前导 0。

**使用语法：**

```javascript
Math.clz32(value)

```

**参数说明：**

|参数|描述|
|----|----|
|value|可以是任意数字，如果是非数值则会被隐式转换为数字类型|

`Math.clz32()` 函数返回一个数字在转换成 32 无符号整形数字的二进制形式后，开头的 0 的个数，比如 `1000000` 转换成 32 位无符号整形数字的二进制形式后是 `00000000000011110100001001000000`, 开头的 0 的个数是 12 个，则 `Math.clz32(1000000)` 返回 `12`.

**实例：**

```javascript
Math.clz32(-0) // 32
Math.clz32(0) // 32
Math.clz32(1) // 31
Math.clz32(1000) // 22
Math.clz32(0b01000000000000000000000000000000) // 1
Math.clz32(0b00100000000000000000000000000000) // 2
```

左移运算符（<<）与 `Math.clz32()` 方法直接相关。

```javascript
Math.clz32(0) // 32
Math.clz32(1) // 31
Math.clz32(1 << 1) // 30
Math.clz32(1 << 2) // 29
Math.clz32(1 << 29) // 2
```

对于小数，Math.clz32 方法只考虑整数部分。

```javascript
Math.clz32(3.2) // 30
Math.clz32(3.9) // 30
```

对于空值或其他类型的值，Math.clz32 方法会将它们先转为数值，然后再计算。无法转换为数值的按照 0 来算。

```javascript
Math.clz32() // 32
Math.clz32(NaN) // 32
Math.clz32(Infinity) // 32
Math.clz32(null) // 32
Math.clz32('foo') // 32
Math.clz32([]) // 32
Math.clz32({}) // 32
Math.clz32(true) // 31
```

## 6. Math.imul()

`Math.imul()` 方法返回两个数以 32 位带符号整数形式相乘的结果，返回的也是一个 32 位的带符号整数。

**使用语法：**

```javascript
Math.imul(a, b)

```

**参数说明：**

|参数|描述|
|----|----|
|a|被乘数|
|b|乘数  |

如果只考虑最后 32 位，大多数情况下，`Math.imul(a, b)` 与 `a * b` 的结果是相同的，即该方法等同于 `(a * b)|0` 的效果（超过 32 位的部分溢出）。

之所以需要部署这个方法，是因为 JavaScript 有精度限制，超过 2 的 53 次方的值无法精确表示。这就是说，对于那些很大的数的乘法，低位数值往往都是不精确的，Math.imul 方法可以返回正确的低位数值。

**实例：**

```javascript
Math.imul() // 0
Math.imul(3, 5) // 15
Math.imul(-1, 5) // -5
Math.imul(1, -5) // -5
Math.imul(-2, -2) // 4
Math.imul(2.9, 4.5) // 8
Math.imul(0.9, 5) // 0
Math.imul(true, 5) // 1
Math.imul(NaN, 5) // 0
```

这里注意的是，`Math.imul()` 方法接收的参数是整数，如果参数是浮点数则会调用 `Math.trunc()` 把小数点后面的数字去掉，然后再相乘。如果不能被转换则统一返回 0。

## 7. Math.fround()

`Math.fround()` 方法返回一个数的 32 位单精度浮点数形式。

**使用语法：**

```javascript
Math.fround(doubleFloat)
```

**参数说明：**

|参数|描述|
|----|----|
|doubleFloat|一个 `Number`。若参数为非数字类型，则会被转换成数字。无法转换时，设置成 `NaN`。|

对于 32 位单精度格式来说，数值精度是 24 个二进制位（1 位隐藏位与 23 位有效位），所以对于 -224 至 224 之间的整数（不含两个端点），返回结果与参数本身一致。

**实例：**

```javascript
Math.fround(0) // 0
Math.fround(1) // 1
Math.fround(2 ** 24 - 1) // 16777215
```

如果参数的绝对值大于 224，返回的结果便开始丢失精度。

```javascript
Math.fround(2 ** 24) // 16777216
Math.fround(2 ** 24 + 1) // 16777216
```

`Math.fround()` 方法的主要作用，是将 64 位双精度浮点数转为 32 位单精度浮点数。如果小数的精度超过 24 个二进制位，返回值就会不同于原值，否则返回值不变（即与 64 位双精度值一致）。

```javascript
// 未丢失有效精度
Math.fround(1.125) // 1.125
Math.fround(7.25) // 7.25

// 丢失精度
Math.fround(0.3) // 0.30000001192092896
Math.fround(0.7) // 0.699999988079071
Math.fround(1.0000000123) // 1
```

对于 NaN 和 Infinity，此方法返回原值。对于其它类型的非数值，`Math.fround()` 方法会先将其转为数值，再返回单精度浮点数。

```javascript
Math.fround(NaN) // NaN
Math.fround(Infinity) // Infinity
Math.fround('5') // 5
Math.fround(true) // 1
Math.fround(null) // 0
Math.fround([]) // 0
Math.fround({}) // NaN
```

## 8. Math.hypot()

Math.hypot 方法返回所有参数的平方和的平方根。

**使用语法：**

```javascript
Math.hypot([value1[,value2, ...]])
```

**参数说明：**

|参数|描述|
|----|----|
|value1, value2, …|任意数字|

如果参数不是数值，Math.hypot 方法会将其转换为数值。只要有一个参数无法转换为数值，就会返回 NaN。

**实例：**

```javascript
Math.hypot(3, 4); // 5
Math.hypot(3, 4, 5); // 7.0710678118654755
Math.hypot(); // 0
Math.hypot(NaN); // NaN
Math.hypot(3, 4, 'foo'); // NaN
Math.hypot(3, 4, '5'); // 7.0710678118654755
Math.hypot(-3); // 3
```

## 9. 对数方法

### 9.1 Math.expm1()

`Math.expm1(x)` 函数返回 `E^x - 1`，其中 `x` 是该函数的参数， `E` 是自然对数的底数 `2.718281828459045`。

**使用语法：**

```javascript
Math.expm1(x)
```

**参数说明：**

|参数|描述|
|----|----|
|x|任意数字|

**实例：**

```javascript
Math.expm1(1) // 1.718281828459045
Math.expm1(0) // 0
Math.expm1(-1) // -0.6321205588285577
Math.expm1(-37) // -0.9999999999999999
Math.expm1(-38) // -1
Math.expm1(-88) // -1
Math.expm1("-38") // -1
Math.expm1("foo") // NaN
```

### 9.2 Math.log1p()

`Math.log1p(x)` 函数返回 `1 + x` 的自然对数 （底为 `E`)，即 `Math.log(1 + x)`。如果 x 小于 -1，返回 NaN。

**使用语法：**

```javascript
Math.log1p(x)
```

**参数说明：**

|参数|描述|
|----|----|
|x|任意数字|

**实例：**

```javascript
Math.log1p(1) // 0.6931471805599453
Math.log1p(Math.E-1) // 1
Math.log1p(0) // 0
Math.log1p("0") // 0
Math.log1p(-1) // -Infinity
Math.log1p(-2) // NaN
Math.log1p("imooc") // NaN
```

### 9.3 Math.log10()

`Math.log10(x)` 函数返回以 10 为底的 x 的对数。如果 x 小于 0，则返回 NaN。

**使用语法：**

```javascript
Math.log10(x)
```

**参数说明：**

|参数|描述|
|----|----|
|x|任意数字|

**实例：**

```javascript
Math.log10(10) // 1
Math.log10(100000) // 5
Math.log10("100")// 2
Math.log10(0) // -Infinity
Math.log10(1) // 0
Math.log10(2) // 0.3010299956639812
Math.log10(-2) // NaN
Math.log10("imooc")// NaN
```

### 9.4 Math.log2()

Math.log2 (x) 返回以 2 为底的 x 的对数。如果 x 小于 0，则返回 NaN。

**使用语法：**

```javascript
Math.log2(x)
```

**参数说明：**

|参数|描述|
|----|----|
|x|任意数字|

**实例：**

```javascript
Math.log2(3) // 1.584962500721156
Math.log2(2) // 1
Math.log2(1) // 0
Math.log2(0) // -Infinity
Math.log2(-2) // NaN
Math.log2(1024) // 10
Math.log2("1024") // 10
Math.log2(1 << 29) // 29
Math.log2("imooc") // NaN
```

## 10. 幂运算符

ES2016 新增了一个幂运算符 `**` （也可以说是指数运算符）。幂运算符返回第一个操作数作底数，第二个操作数作指数的乘方。即，`var1^var2`，其中 `var1` 和 `var2` 是其两个操作数。幂运算符是右结合的。`a ** b ** c` 等同于 `a ** (b ** c)`。

```javascript
3 ** 2 // 9
3 ** 3 // 27
```

指数运算符可以与等号结合，形成一个新的赋值运算符（**=）。

```javascript
let a = 1.5;
a **= 2;
// 等同于 a = a * a;
let b = 4;
b **= 3;
// 等同于 b = b * b * b;
```

指数运算符与 `Math.pow()` 基本相同，不过使用幂运算符更加方便简洁。

```javascript
Math.pow(99, 99)
// 3.697296376497268e+197
99 ** 99
// 3.697296376497268e+197
```

## 11. 小结

本节主要介绍了 ES6 对 Math 这个对象的扩展，ES6 主要丰富了更多的数学方法，让我们更好地处理数学问题。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

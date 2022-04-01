# JavaScript Math

> Math 是一个内置对象， 它具有数学常数和函数的属性和方法。(MDN)

Math 对象提供了一些列的数学方法和常量，如三角函数、随机数、π等。

## 1. 数学方法

数学方法其实就是常用的数学中的函数，如三角函数。

其中用到较多的是使用 `Math.random()` 来产生随机数。

以下列举部分常用的方法。

### 1.1 Math.floor

> Math.floor() 返回小于或等于一个给定数字的最大整数。

Math.floor 就是对一个数进行向下取整。

```javascript
var num1 = Math.floor(2.4);
var num2 = Math.floor(2.9);

console.log(num1); // 输出：2
console.log(num2); // 输出：2
```

因为是向下取整，不会进行四舍五入，所以不论是 `2.4` 还是 `2.9` ，经过 `Math.floor` 处理后，都为 `2` 。

### 1.2 Math.ceil

> Math.ceil() 函数返回大于或等于一个给定数字的最小整数。

Math.ceil 就是对一个数进行向上取整。

```javascript
var num1 = Math.ceil(2.4);
var num2 = Math.ceil(2.9);

console.log(num1); // 输出：3
console.log(num2); // 输出：3
```

与 Math.floor 相反，Math.ceil 向上取整，也不会进行四舍五入，所以结果都为 3 。

> 许多开发者经常会弄混 `Math.floor` 与 `Math.ceil` ，其实可以根据方法名的中译来记忆，floor 可以理解成地板，ceil 可以理解成天花板，所以前者是向下取整，后者是向上取整。

### 1.3 Math.random

Math 下的 random 方法调用后返回一个`大于等于零且小于一`的随机数，即区间 `[0, 1)` 。

```javascript
var random = Math.random();

console.log(random);
```

可以尝试运行多次这段代码，`random`变量的值出现一样的概率很低。

> 事实上 Math.random() 产生的随机数并不是真正的随机数，其返回的“随机数”是由一个确定的算法得出的，这种随机数会称之为`伪随机数`。

#### 1.3.1 随机数应用

在需要随机数的需求时，大部分情况下不会是需要一个区间为 [0, 1) 的随机数，如按学号抽奖。

假设学号的范围是 1 至 100 ，那就需要产生一个 1 至 100 之间的随机数。

这样可以换个思路，随机数还是生成 [0, 1) 范围的大小，但是将这个值放大到 [1, 100] 区间的范围。

所以就可以把随机数的结果先放大 100 倍，即 `Math.random() * 100` ，这时候产生的数，区间就是 `[0, 100)` 。

因为最小值需要的是 0，而不是 1，就可以将这个随机数表达式修改成 `Math.random() * 100 + 1` ，这时候区间就变成了 `[1, 101)` 。

最后只要保证右侧的区间不大于 100 即可，因为右侧区间取不到 101 ，所以可以对结果进行向下取整。

最终表达式为：

```javascript
Math.floor(Math.random() * 100 + 1);
```

由此就可以推出一个较为通用的随机数公式：

```javascript
// 随机数区间：[下限, 上限]
Math.floor(Math.random() * 上限 + 下限);
```

## 2. 数学常量

许多数学常量被作为属性放在 `Math` 对象下，如`欧拉常数(E)`，`圆周率(PI)`。

```javascript
var e = Math.E;
var pi = Math.PI;

console.log(e); // 输出：2.718281828459045
console.log(pi); // 输出：3.141592653589793
```

> 日常开发中，大部分情况下会选择自己维护一个常量，因为可以随时的统一的修改精度。

## 3. 常量与方法列表

因为都是常量与方法，这里提供相应的列表供快速查阅，内容引用自 MDN 。

### 3.1 常量

|常量|描述|
|----|----|
|[Math.E](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/E)            |表示自然对数的底数（或称为基数），约等于 2.718。|
|[Math.LN10](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/LN10)      |表示 10 的自然对数，约为 2.302。                |
|[Math.LN2](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/LN2)        |表示 2 的自然对数，约为 0.693。                 |
|[Math.LOG10E](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/LOG10E)  |表示以 10 为底数，e 的对数，约为 0.434。        |
|[Math.LOG2E](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/LOG2E)    |表示以 2 为底数，e 的对数，约为 1.442。         |
|[Math.PI](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/PI)          |表示一个圆的周长与直径的比例，约为 3.14159。    |
|[Math.SQRT1_2](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/SQRT1_2)|属性表示 1/2 的平方根，约为 0.707。             |
|[Math.SQRT2](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/SQRT2)    |表示 2 的平方根，约为 1.414。                   |

### 3.2 方法

|方法|描述|
|----|----|
|[Math.abs](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/abs)                                                     |返回一个数的的绝对值。                                                                        |
|[Math.acos](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/acos)                                                   |返回一个数的反余弦值（单位为弧度）。                                                          |
|[Math.acosh](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/%E5%8F%8D%E5%8F%8C%E6%9B%B2%E4%BD%99%E5%BC%A6%E5%80%BC)|返回一个数字的反双曲余弦值。                                                                  |
|[Math.asin](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/asin)                                                   |返回一个数值的反正弦（单位为弧度）。                                                          |
|[Math.asinh](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/asinh)                                                 |返回给定数字的反双曲正弦值                                                                    |
|[Math.atan](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/atan)                                                   |返回一个数值的反正切（以弧度为单位）                                                          |
|[Math.atan2](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/atan2)                                                 |返回从原点(0,0)到(x,y)点的线段与x轴正方向之间的平面角度(弧度值)。                             |
|[Math.atanh](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/atanh)                                                 |函数返回一个数值反双曲正切值。                                                                |
|[Math.cbrt](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/cbrt)                                                   |返回任意数字的立方根。                                                                        |
|[Math.ceil](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/ceil)                                                   |返回大于或等于一个给定数字的最小整数。                                                        |
|[Math.clz32](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/clz32)                                                 |返回一个数字在转换成 32 无符号整形数字的二进制形式后, 开头的 0 的个数。                       |
|[Math.cos](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/cos)                                                     |返回一个数值的余弦值。                                                                        |
|[Math.cosh](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/cosh)                                                   |返回数值的双曲余弦函数。                                                                      |
|[Math.exp](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/exp)                                                     |函数返回 ex，x 表示参数，e 是欧拉常数（Euler’s constant），自然对数的底数。                  |
|[Math.expm1](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/expm1)                                                 |函数返回 Ex - 1, 其中 x 是该函数的参数, E 是自然对数的底数。                                  |
|[Math.floor](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/floor)                                                 |返回小于或等于一个给定数字的最大整数。                                                        |
|[Math.fround](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/fround)                                               |将任意的数字转换为离它最近的单精度浮点数形式的数字。                                          |
|[Math.hypot](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/hypot)                                                 |函数返回它的所有参数的平方和的平方根。                                                        |
|[Math.imul](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/imul)                                                   |该函数将两个参数分别转换为 32 位整数，相乘后返回 32 位结果，类似 C 语言的 32 位整数相乘。     |
|[Math.log](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/log)                                                     |返回一个数的自然对数。                                                                        |
|[Math.log10](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/log10)                                                 |函数返回一个数字以 10 为底的对数。                                                            |
|[Math.log1p](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/log1p)                                                 |函数返回一个数字加1后的自然对数 (底为 E), 即log(x+1)。                                        |
|[Math.log2](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/log2)                                                   |函数返回一个数字以 2 为底的对数。                                                             |
|[Math.max](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/max)                                                     |返回一组数中的最大值。                                                                        |
|[Math.min](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/min)                                                     |返回零个或更多个数值的最小值。                                                                |
|[Math.pow](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/pow)                                                     |函数返回基数（base）的指数（exponent）次幂，即 baseexponent。                                 |
|[Math.random](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/random)                                               |函数返回一个浮点, 伪随机数在范围从0到小于1，也就是说，从0（包括0）往上，但是不包括1（排除1）。|
|[Math.round](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/round)                                                 |函数返回一个数字四舍五入后最接近的整数。                                                      |
|[Math.sign](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/sign)                                                   |函数返回一个数字的符号, 指示数字是正数，负数还是零。                                          |
|[Math.sin](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/sin)                                                     |函数返回一个数值的正弦值。                                                                    |
|[Math.sinh](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/sinh)                                                   |函数返回一个数字(单位为角度)的双曲正弦值。                                                    |
|[Math.sqrt](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/sqrt)                                                   |函数返回一个数的平方根。                                                                      |
|[Math.tan](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/tan)                                                     |返回一个数值的正切值。                                                                        |
|[Math.tanh](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/tanh)                                                   |函数将会返回一个数的双曲正切函数值。                                                          |
|[Math.trunc](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Math/trunc)                                                 |方法会将数字的小数部分去掉，只保留整数部分。                                                  |

## 4. 小结

Math 对象包含了很多数学方法与常量，较常用的是用于产生伪随机数的 `Math.random`，根据特性可以推出随机数 `[下限, 上限]` 的生成公式 `Math.floor(Math.random() * 上限 + 下限);`。

尽量将 Math 对象下的方法过一遍，留住印象，避免造不必要的轮子。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

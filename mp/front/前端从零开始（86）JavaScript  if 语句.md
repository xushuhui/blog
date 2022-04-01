# JavaScript if 语句

在程序中 if 语句属于条件语句的一种。

如同 `if` 的本意，就是根据条件做不同的事情。

## 1. 基本语法

`if` 语句的最基本语法如下：

```javascript
// 方式1
if (条件) {
  // 条件满足做的事情;
}

// 方式2
if (条件) 条件满足时候做的事情;
```

> 通常方式 2 的使用场景是`在条件满足时只会写一行代码`的情况，有些代码压缩工具可能会将多个语句配合`逗号表达式`压缩成方式 2。

这是 if 语句两种最基本语法。

第一种方式用到的相对较多，当条件满足的时候，就会执行大括号内的代码，第二种方式则会在条件满足的时候执行括号后面到行尾的语句。

条件满足的情况其实就是`条件的执行结果不为false或者不能被隐式转为false`的情况。

```javascript
var score = 99;

if (score > 60) {
  console.log('及格了'); // 输出："及格了"
}

if (score > 90) {
  console.log('优秀！'); // 输出："优秀！"
}
```

`>`符号就是判断左边的数是否大于右边，如果左边大于右边则会返回 `true`，否则返回 `false`。

这里两个 `if` 语句的条件都是满足的，所以会输出`及格了`和`优秀！`。

方式二在简单函数中很常见

```javascript
// 计算圆面积
function calcArea(r) {
  if (!r) return 0;

  return r * r * 3.14;
}
```

这里判断了 r 是否有传入，如果没有传入则直接返回了 0。

## 2. 分支

if 语句可以仅有单个分支也可以有多个分支。

```javascript
// 单个分支
if (条件1) {
  // 条件1满足的时候执行的代码
} else {
  // 条件1不满足的时候执行的代码
}

// 多个分支
if (条件1) {
  // 条件1满足的时候执行的代码
} else if (条件2) {
  // 条件2满足的时候执行的代码
} else if (条件3) {
  // 条件3满足的时候执行的代码
} else {
  // 条件1、条件2、条件3都不满足的时候执行
}
```

单个分支的方式比较常用：

```javascript
var score = 77;

if (score >= 60) {
  console.log('及格了');
} else {
  console.log('不及格');
}

// 输出："及格了"
```

`>=`则表示左边的值如果大于等于右边的值，则返回 `true` 否则返回 `false`。

通过这样的分支，就可以扩展条件的场景。

多个分支的场景也非常用到，如需要判断成绩的不同区间给出不同的标准：

```javascript
var score = 88;

if (score < 60) {
  console.log('不及格');
} else if (score < 80) {
  console.log('良好');
} else if (score < 90) {
  console.log('优秀！');
} else {
  // 剩下的肯定是大于等于九十的情况
  console.log('太强了！');
}

// 输出："优秀！"
```

在多个分支的情况下，`else` 也可以不提供。

```javascript
var score = 0;

if (score < 60) {
  console.log('不及格');
} else if (score < 80) {
  console.log('良好');
} else if (score < 90) {
  console.log('优秀！');
} else if (score <= 100) {
  console.log('太强了！');
}

// 输出：不及格
```

## 3. 例子

> 注意：例子的解法不止一种，可以自己发散实现

### 3.1 判断一个数是不是偶数

```javascript
var num = 77;

if (num % 2 === 0) {
  console.log(num + '是一个偶数');
} else {
  console.log(num + '是一个奇数');
}

// 输出：77是一个奇数
```

这里就是通过将数字对 2 进行取余数操作，如果余数是 0 则说明一个数是偶数，否则就是奇数。

### 3.2 计算成人的标准体重

```javascript
var sex = '男';
var height = 178;

var offset = 0;
if (sex === '男') {
  offset = 105;
} else {
  offset = 100;
}

var weight = height - offset;

console.log('身高为' + height + 'cm的' + sex + '性标准体重为' + weight + 'KG');

// 输出：身高为178cm的男性标准体重为73KG
```

首先要知道标准体重的计算公式：

* 男：身高（cm）-105 = 标准体重（kg）
* 女：身高（cm）-100 = 标准体重（kg）

可以看到男生需要身高减去一个 `105` 的偏移量，女生需要减去 `100` 的偏移量，所以要判断性别决定这个值。

每个人的身高是不一样的，所以作为变量。

所以只需要根据性别判断是减去 105 还是减去 100，最后计算出结果即可。

## 4. 小结

if 语句很简单，但是非常常用和重要，高级语言如果失去了条件语句，又没有其他的替代品，基本就失去了活力。

使用 if 语句的时候也需要注意，碰到冗长的条件尽量思考一下是不是有更好的解决方案，防止代码可读性和可维护性变差。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

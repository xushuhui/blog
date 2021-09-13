# ES6+ repeat()

## 1. 前言

本节介绍 ES6 的字符串新增方法 `repeat`，以一个目标字符串进行声明，对该字符串进行重复操作，但不会改变原字符串。

## 2. 方法详情

**使用语法：**

```javascript
let resultString = str.repeat(count);
```

该方法构造并返回一个新字符串，表示将原字符串重复 n 次，并不会改变原字符串。

**参数说明：**

|参数|描述|
|----|----|
| count| 介于 0 和正无穷大之间的整数：[0, +∞)。表示在新构造的字符串中重复了多少遍原字符串。count 取负数的时候会报错，但是在 (-1.0] 之间不会报错，而会把 count 处理成 0|

## 3. 使用场景

1. 可以通过这个方法拷贝一个相同的字符串；
2. 取代循环拼接多个相同的字符串，会比使用 for 循环优雅，方便。

## 4. 实例

### 4.1 参数是小数

参数如果是小数，会被取整。

```javascript
"imooc".repeat(2.6)    // "imoocimooc"
```

参数 2.6 会被向下自动转换成整数，注意这里不会进位成 3。

### 4.2 负数 和 Infinity

如果 repeat 的参数是负数或者 Infinity，会报错。

```javascript
'imooc'.repeat(Infinity)  // RangeError
"imooc".repeat(-1)        // 无效的数字
```

### 4.3 0～1 和 0～- 1

如果参数是 0 到 - 1 或 0 到 1 都会先进行取整运算，所以在这两个范围内都会被 `repeat` 视同为 0。

```javascript
'imooc'.repeat(0.9)   // ''
"imooc".repeat(-0.8)  // ''
```

### 4.4 参数 NaN 等同于 0

```javascript
'imooc'.repeat(NaN)   // ''
```

### 4.4 参数是字符串

如果 repeat 的参数是字符串，则会先转换成数字。

```javascript
'imooc'.repeat('two')   // ''
'imooc'.repeat('2')     // 'imoocimooc'
```

## 5. 小结

本节讲解了字符串的 `repeat()` 方法的使用，需要注意以下几点：

* 重复次数不能为负数；
* 重复次数必须小于 infinity，且长度不会大于最长的字符串。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

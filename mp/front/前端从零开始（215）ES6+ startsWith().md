# ES6+ startsWith()

## 1. 前言

在 ES5 中用于查找字符串的方法很少，[上一节](https://www.imooc.com/wiki/ES6lesson/stringincludes.html) 我们学习了 `includes()`方法，它是针对整个字符串进行查找的，本节要介绍 ES6 的字符串新增方法 `startsWith()`，该方法用来判断当前字符串是否以给定的字符串作为开头。

字符串查找是有一定的算法的，虽然用 `includes()` 方法可以判断，无疑只查找字符串的开头算法的时间复杂度是很低的，但是使用 `includes()` 就需要对整个字符串进行查找，时间复杂度也会很高。在查找长字符串时也会比较耗费性能，虽然在测试过程中这种差别几乎可以被忽略，但是它的语义化让我们的代码可读性更高。

## 2. 方法详情

`startsWith()` 用于判断一个字符串，是否以一个指定字符串为起始的，如果是字符串的开头找到了给定的字符则返回 true，否则返回 false。

**使用语法：**

```javascript
str.startsWith(searchString[, position])
```

**参数说明：**

|参数|描述|
|----|----|
|searchString|要搜索的子字符串。                                                                  |
|position    |（可选） 在 str 中搜索 searchString 的开始位置，默认值为 0，也就是真正的字符串开头处。|

**实例：**

```javascript
const str1 = 'I love imooc.';

console.log(str1.startsWith('I'));     // true
console.log(str1.startsWith('I', 3));  // false
```

## 3. 使用场景

确定一个字符串是否以另一个字符串开头，但是使用得比较少，使用比较多的是 `includes()`。

### 3.1 一个参数

```javascript
var str = "I love imooc.";

console.log(str.startsWith("I love"));    // true
console.log(str.startsWith("imooc"));     // false
console.log(str.startsWith("eimooc"));    // false
```

### 3.2 两个参数

```javascript
var str = "I love imooc.";

console.log(str.startsWith("love", 3));   // false
console.log(str.startsWith("ove", 3));    // true
```

第二个参数的意思是，字符串的位置，上面第二个参数是 3 说明是从字符串的第三个字符开始往后，包括第三个字符。所以第一个返回的结果为 false。

## 4. 小结

在查询字符串中的开头时可以使用 `startsWith` 或者 `includes()`，在字符串少的情况下，它们的效率基本没有差别。但是如果在查询以某字符串开头的时候，使用 `startsWith` 会很有语义化，利于代码阅读。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

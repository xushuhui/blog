# endsWith()

## 1. 前言

[上一节](https://www.imooc.com/wiki/ES6lesson/startsWith.html) 我们学到了字符串的方法 `startsWith()` 用于判断是否以一个指定字符串为起始。 本节介绍 ES6 中新增的与之相反的字符串方法 `endsWith()`，该方法用来判断当前字符串是否是以另外一个给定的子字符串为结尾。

## 2. 方法详情

`endsWith()` 用于判断当前字符串，是否以一个指定字符串为结尾的，如果在字符串的结尾找到了给定的字符则返回 true，否则返回 false。

**使用语法：**

```javascript
str.endsWith(searchString[, length])
```

**参数说明：**

|参数|描述|
|----|----|
|searchString|需要查询的子字符串                         |
|length      |（可选） 作为 str 的长度。默认值为 str.length|

**实例：**

```javascript
const str1 = 'Cats are the best!';
console.log(str1.endsWith('best', 17));   // true

const str2 = 'Is this a question';
console.log(str2.endsWith('?'));          // false
```

## 3. 使用场景

查询一个字符串是否在另一个字符串的末尾。

### 3.1 没有参数

这里需要说明一下的是，当字符串调用 `endsWith()` 方法时不传参数时，默认是 `undefined` 返回的结果是 false。

```javascript
var str = "I love imooc.";
console.log(str.endsWith());   // false
console.log(str.endsWith(undefined));	// false
```

上面的代码中，第 2 行和第 3 行是等价的，因为第一个参数是必填的，所以在当我们没有传参时，默认使用 `undefined` 来填充，注意这里不是字符串类型的 ‘undefined’

### 3.2 一个参数

```javascript
var str = "I love imooc.";

console.log(str.endsWith("I love"));    // false
console.log(str.endsWith("imooc"));     // false
console.log(str.endsWith("imooc."));    // true
console.log(str.endsWith(""));    			// true
```

从例子中我们可以看出，只有结尾有最后一个字符的时候，才会返回 true，只要没有包含结尾的字符，即使查找的字符串在目标字符串里也是返回 fasle 的。在查找空字符串时，返回的结果是 true，那是因为空字符在任何字符串中都是存在的。

### 3.3 两个参数

当有第二个参数的时候，第二个参数是字符串的长度

```javascript
var str = "I love imooc.";

console.log(str.endsWith("love", 6));   // true
console.log(str.endsWith("e", 6));      // true
```

从这两个 log 打印出来的结果可以看出，第二个参数会取原字符串的指定长度作为查找的目标字符串，这里的第二个参数是 6 也就是取原字符串的 `I love`，所以 `endsWith` 判断是以 `love` 结尾的。

## 4. 小结

在查询字符串中的结尾时最好使用 `endsWith` 进行查询，它的效率要比 `includes()` 高，而且 `endsWith` 也具有语义化。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

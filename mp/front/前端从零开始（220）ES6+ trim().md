# ES6+ trim()

## 1. 前言

本节介绍 ES6 的字符串新增方法 `trim()`，该方法会从一个字符串的两端删除空白字符。在这个上下文中的空白字符是所有的空白字符 (space, tab, no-break space 等） 以及所有行终止符字符（如 LF，CR 等）。

## 2. 方法详情

`trim()` 的方法返回值是去掉两端空白字符的字符串，并不影响原字符串本身，不接收任何参数。

**使用语法：**

```javascript
str.trim()
```

在低版本浏览器中是不支持这方法的，一版我们会使用正则的方式去去除字符串两边的空格的。

```javascript
if (!String.prototype.trim) {
    String.prototype.trim = function () {
        return this.replace(/^\s+|\s+$/gm, '');
    }
}
```

## 3. 使用场景

* 去除字符串两端的空白字符。

## 4. 实例

**1. 删除前后的空白字符。**

```javascript
var str = '   foo  ';
console.log(str.trim());  // 'foo'
```

**2. 如果字符串只有一边有空白字符，则只删除一边的空白字符。**

```javascript
var str = 'foo    ';
console.log(str.trim());  // 'foo'
var str = '   foo';
console.log(str.trim());  // 'foo'
```

## 5. 拓展

通过 `trim()`方法衍生出的两个方法 `trimStart()` 和 `trimEnd()`。 `trimStart()` 是删除字符串左边的空白字符，`trimEnd()` 是删除字符串右边的空白字符。如：

```javascript
const str = '  imooc  ';
str.trim() 			// "imooc"
str.trimStart() 	// "imooc  "
str.trimEnd() 		// "  imooc"
```

`trimStart()` 把 str 左边的空白字符去掉了，`trimEnd()` 把 str 右边的空白字符去掉了。

另外在浏览器中我们可以使用 `trimStart()` 和 `trimEnd()` 的别名，也能达到同样的效果，`trimLeft()` 是 `trimStart()` 的别名，`trimRight()` 是 `trimEnd()` 的别名。

```javascript
const str = '  imooc  ';
str.trimLeft() 		// "imooc  "
str.trimRight() 	// "  imooc"
```

## 6. 小结

本节讲解了字符串的 `trim()` 方法的使用，总结以下几点：

* trim 方法会去除字符串两边的空白字符串；
* 如果只去除一般字符串时可以使用 `trimStart()` 和 `trimEnd()`。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

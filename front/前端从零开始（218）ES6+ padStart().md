# ES6+ padStart()

## 1. 前言

本节介绍 ES6 的字符串新增方法 `padStart`，该方法是字符串补全长度的方法，如果某个字符串不够指定的字符串长度，可以使用此方法在字符串左侧开始位置进行补全。

## 2. 方法详情

`padStart()` 在原字符串开头填充指定的填充字符串直到目标长度并返回一个新的字符串，不会对原字符串进行修改。

**使用语法：**

```javascript
str.padStart(targetLength [, padString])
```

**参数说明：**

|参数|描述|
|----|----|
|targetLength|当前字符串需要填充到的目标长度。如果这个数值小于当前字符串的长度，则返回当前字符串本身                                                   |
|padString   |（可选） 填充的字符串。如果字符串太长，使填充后的字符串长度超过了目标长度，则只保留最左侧的部分，其他部分会被截断。默认补全的字符串为 `' '`|

## 3. 使用场景

1. 扩展字符串长度；
2. 设置指定字符串的长度；
3. 补全日期的时候会更加方便。

## 4. 实例

**1. 拼接字符串。**

```javascript
"imooc".padStart(10, 'ilove')    // "iloveimooc"
```

**2. 如果原字符串的长度，等于或小于最大长度，则字符串补全不生效，返回原字符串。**

```javascript
'imooc'.padStart(5, 'ab') // 'imooc'
'imooc'.padEnd(2, 'ab') // 'imooc'
```

**3. 如果补全的字符串和原字符串的长度大于目标（targetLength）的长度，补全的字符串会被截取。**

```javascript
'imooc'.padStart(7, 'abc') // 'abimooc'
```

**4. 如果补全的字符串和原字符串的长度小于目标（targetLength）的长度，补全的字符串会被重复，多余的部分会被裁剪。**

```javascript
'imooc'.padStart(9, 'ab') // 'ababimooc'
'imooc'.padStart(10, 'ab') // 'ababaimooc'
```

**5. 如果省略第二个参数，默认使用空格补全长度。**

```javascript
'imooc'.padStart('7')   // '  imooc'
```

## 5. 场景实例（补全日期）

通常情况下用的比较多的就是在时间或者日期前面的补 0，比如：`2020-06-03`，但是通常我们使用时间戳获取日月时，是没有前面的 0 的，如：

```javascript
var month = new Date().getMonth() + 1;  // 6
```

这个时候获取的是 2，没有前面的 0，如果我们想在月份前面加 0 需要进行逻辑判断，我们可以写这样一个函数来统一处理实现。

```javascript
function getMonth(m) {
  return m < 10 ? `0${m}` : m;
}
```

当 m 小于 10 的时候，我们会在前面添加一个 0，否则直接返回 m 的值，虽然这样可以实现，但是这里多了一个函数，现在有了 `padStart` 就会很容易了。

```javascript
var month = String(new Date().getMonth() + 1).padStart(2, '0');	// 06
var date = String(new Date().getDate()).padStart(2, '0'); 			// 03
```

`String()` 函数对日期进行类型转换的作用，转换为字符串进行操作。

## 6. 小结

本节讲解了字符串的 `padStart()` 方法的使用，需要注意以下几点：

* 该方法不会对原字符串进行修改；
* 在没有第二个参数时，则用空格填充；
* 在实战中使用在日期前的补充会使代码变得非常简洁。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

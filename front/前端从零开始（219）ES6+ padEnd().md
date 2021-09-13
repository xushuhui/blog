# ES6+ padEnd()

## 1. 前言

本节介绍 ES6 的字符串新增方法 `padEnd` 和 `padStart` 一样也是补全字符串的长度的，但是它的补全位置是在原字符串的右侧末尾进行补全的。此方法会返回一个新的字符串，不会对原字符串进行修改。

## 2. 方法详情

`padEnd()` 在原字符串末尾填充指定的填充字符串直到目标长度并返回一个新的字符串，不会对原字符串进行修改。

**使用语法：**

```javascript
str.padEnd(targetLength [, padString])
```

**参数说明：**

|参数|描述|
|----|----|
|targetLength|当前字符串需要填充到的目标长度。如果这个数值小于当前字符串的长度，则返回当前字符串本身                                                 |
|padString   |（可选） 填充字符串。如果字符串太长，使填充后的字符串长度超过了目标长度，则只保留最左侧的部分，其他部分会被截断。默认补全的字符串为 `' '`|

## 3. 使用场景

1. 扩展字符串长度
2. 设置指定字符串的长度
3. 在实战中补全时间戳的毫秒

## 4. 实例

**1. 拼接字符串。**

```javascript
"imooc".padEnd(10, 'ilove')    // "imoocilove"
```

**2. 如果原字符串的长度，等于或小于最大长度，则字符串补全不生效，返回原字符串。**

```javascript
'imooc'.padEnd(5, 'ab') // 'imooc'
'imooc'.padEnd(2, 'ab') // 'imooc'
```

**3. 如果补全的字符串和原字符串的长度大于目标（targetLength）的长度，补全的字符串会被截取。**

```javascript
'imooc'.padEnd(7, 'abc') // 'imoocab'
```

**4. 如果补全的字符串和原字符串的长度小于目标（targetLength）的长度，补全的字符串会被重复，多余的部分会被裁剪。**

```javascript
'imooc'.padEnd(9, 'ab') // 'imoocabab'
'imooc'.padEnd(10, 'ab') // 'imoocababa'
```

**5. 如果省略第二个参数，默认使用空格补全长度。**

```javascript
'imooc'.padEnd('7')   // 'imooc  '
```

## 5. 场景实例（补全时间戳的毫秒）

有时候我们处理后端返回的时间戳数据的时候，会发现很多都是秒，是 10 位，这主要是因为数据库存储的问题。这时候我们需要补全到毫秒，可以借助 padEnd 来进行补全操作。如下：

```javascript
let timestamp = 1581828518
timestamp = String(timestamp).padEnd(13, '0'); // 1581828518000
```

`String()` 函数对时间戳 `timestamp` 进行类型转换，转换为字符串进行操作。

## 6. 小结

本节讲解了字符串的 `padEnd()` 方法的使用，需要注意以下几点：

* 该方法不会对原字符串进行修改；
* 在没有第二个参数时，则用空格填充；
* 在实战中使用在补全时间戳的毫秒中会使代码变得非常简洁。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

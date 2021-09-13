# JavaScript Date

> Date 用于创建一个 JavaScript Date 实例，该实例呈现时间中的某个时刻。

Date 对象可以用于处理日期和时间。

Date 对象使用的频率非常高，大量的业务需要对时间进行操作。

## 1. 基本使用

Date 需要实例化后使用。

```javascript
var date = new Date();
```

时间最大的两个使用场景是格式化时间与获取时间戳。

### 1.1 获取时间戳

当实例化时没有传递参数给 `Date` 的时候，则表示创建的对象为实例化时刻的时间。

使用 `getTime` 即可获取时间戳。

```javascript
var date = new Date();
var timestamp = date.getTime();

console.log(timestamp); // 输出当前时间的时间戳
```

部分开发者会利用隐式转换的规则来获取时间戳。

```javascript
var date = new Date();
var timestamp = +date;

console.log(timestamp); // 输出当前时间的时间戳
```

也可以通过 `valueOf` 来获取时间戳。

```javascript
var date = new Date();
var timestamp = date.valueOf();

console.log(timestamp); // 还是输出当前时间的时间戳
```

推荐使用 `getTime` 方法来获取时间戳，以便他人阅读代码以及避免不必要的问题。

### 1.2 格式化时间

格式化时间可以理解成把时间处理成想要的格式，如`年-月-日 时:分;秒`。

通过 Date 对象提供的一些方法，可以获得到对应的时间属性。

假如想把时间格式化成`年/月/日 时:分:秒`的形式：

```javascript
var date = new Date();

var YYYY = date.getFullYear();
var MM = date.getMonth() + 1;
var DD = date.getDate();
var hh = date.getHours();
var mm = date.getMinutes();
var ss = date.getSeconds();

console.log([YYYY, '/', MM, '/', DD, ' ', hh, ':', mm, ':', ss].join(''));
```

通过 Date 对象提供的获取年、月、日、时、分、秒的方法获取到对应的值，最后按照想要的格式拼接即可。

需要注意的是 `getMonth()` 方法返回的月份是 0 至 11 ，更像是月份的索引，实际上对应的月份还要加上 1 。

## 2. 构造函数的参数

Date 对象可以提供 4 种类型的参数，通过参数决定时间，最后对象的实例的操作都围绕这个决定的时间。

### 2.1 不传递参数

当不传递参数的时候，时间会被设置为实例化那一时刻的时间。

### 2.2 Unix 时间戳

这个方式与第一种不传递参数的方式是最常用的两种。

应用场景大部分为从服务端获取数据后，对时间戳进行格式化显示。

```javascript
var data = { _id: '', createdAt: 1482632382582, content: '' };

var date = new Date(data.createdAt);

var YYYY = date.getFullYear();
var MM = date.getMonth() + 1;
var DD = date.getDate();
var hh = date.getHours();
var mm = date.getMinutes();
var ss = date.getSeconds();

console.log([YYYY, '/', MM, '/', DD, ' ', hh, ':', mm, ':', ss].join(''));
// 输出：2016/12/25 10:19:42
```

### 2.3 时间戳字符串

这里并不是指字符串形式的 `Unix 时间戳` ，而是符合 IETF-compliant RFC 2822 timestamps 或 version of ISO8601 标准的时间字符串。

实际上只要能被 `Date.parse` 正确解析成时间戳的字符串，都可以作为参数传递过去。

```javascript
var timestamp = Date.parse('2020/02/02 11:22:33');

var date1 = new Date(timestamp);
var date2 = new Date('2020/02/02 11:22:33');
```

### 2.4 日期的每一个时间属性

这里的时间属性是指：年、月、日、时、分、秒、毫秒。

参数也按照这个顺序传递。

```javascript
// 2048年10月24日 9点9分6秒
var date = new Date(2048, 10 - 1, 24, 9, 9, 6, 0);

var YYYY = date.getFullYear();
var MM = date.getMonth() + 1;
var DD = date.getDate();
var hh = date.getHours();
var mm = date.getMinutes();
var ss = date.getSeconds();

console.log([YYYY, '/', MM, '/', DD, ' ', hh, ':', mm, ':', ss].join(''));
// 输出：2048/10/24 9:9:6
```

第二个参数之所以要减去 1 ，是因为月份是从 0 开始计算的，所以十月应该表示成 9 。

## 3. 其他常用方法

|常量|描述|
|----|----|
|[Date.UTC](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/UTC)                                                                                                              |方法接受的参数同日期构造函数接受最多参数时一样，返回从 1970-1-1 00:00:00 UTC到指定日期的的毫秒数。                                                                                                        |
|[Date.now](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/now)                                                                                                              |返回自 1970 年 1 月 1 日 00:00:00 (UTC) 到当前时间的毫秒数。                                                                                                                                              |
|[Date.parse](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/parse)                                                                                                          |解析一个表示某个日期的字符串，并返回从1970-1-1 00:00:00 UTC 到该日期对象（该日期对象的 UTC 时间）的毫秒数，如果该字符串无法识别，或者一些情况下，包含了不合法的日期数值（如：2015-02-31），则返回值为NaN。|
|[Date.getDate](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getDate)                                                                                                      |根据本地时间，返回一个指定的日期对象为一个月中的哪一日（从 1–31）。                                                                                                                                      |
|[Date.getDay](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getDay)                                                                                                        |根据本地时间，返回一个具体日期中一周的第几天，0 表示星期天。                                                                                                                                              |
|[Date.getFullYear](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getFullYear)                                                                                              |根据本地时间返回指定日期的年份。                                                                                                                                                                          |
|[Date.getHours](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getHours)                                                                                                    |方法根据本地时间，返回一个指定的日期对象的小时。                                                                                                                                                          |
|[Date.getMilliseconds](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getMilliseconds)                                                                                      |根据本地时间，返回一个指定的日期对象的毫秒数。                                                                                                                                                            |
|[Date.getMinutes](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getMinuteshttps://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getMinutes)|方法根据本地时间，返回一个指定的日期对象的分钟数。                                                                                                                                                        |
|[Date.getMonth](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getMonth)                                                                                                    |根据本地时间，返回一个指定的日期对象的月份，为基于 0 的值（0 表示一年中的第一月）。                                                                                                                       |
|[Date.getSeconds](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getSeconds)                                                                                                |方法根据本地时间，返回一个指定的日期对象的秒数。                                                                                                                                                          |

## 4. 小结

`Date` 对象用于处理日期与时间。

通常会采用`不传参`或者`传递一个 Unix 时间戳`来生成 `Date` 实例，另几种参数形式使用场景较少。

需要注意的是，`getMonth` 方法返回的月份，是从 `0` 开始计数的，对应真实月份需要加上 `1`。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

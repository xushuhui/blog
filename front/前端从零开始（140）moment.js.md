# moment

> JavaScript 日期处理类库

`moment.js` 主要用于处理时间，许多程序的第三方框架在一些时间处理上都会采用 `moment.js`。

## 1. 使用

`moment.js` 在全局下以 `moment` 作为入口，提供了一系列时间相关的方法。

```javascript
<script src="https://cdn.bootcdn.net/ajax/libs/moment.js/2.27.0/moment.min.js"></script>
<script>
  var now = moment().calendar();

  console.log(now); // 输出当前日历时间
</script>
```

![](https://img.mukewang.com/wiki/5f018114091027d911800404.jpg)

现在的相对时间差需求非常常见，如下单时间，是 `多少分钟前`，`moment.js` 提供了相对时间计算：

```javascript
moment().startOf('hour').fromNow(); // 相对这个小时过去了多少分钟

var timestamp = 1593933593236; // 2020年7曰5日下午15点20分38秒
moment(timestamp).fromNow(); // 相对时间戳多久前
```

![图片描述](https://img.mukewang.com/wiki/5f0181230965f5e012180272.jpg)

## 2. 国际化

上述例子发现输出的结果是英文的，显然是不适合在国内环境使用，`moment.js` 提供了国际化支持，在现有的库中，moment 支持的语言可以说是相对完备了。

通过引入对应的国际化资源（语言文件），来切换语言。

```javascript
<script src="https://cdn.bootcdn.net/ajax/libs/moment.js/2.27.0/moment.min.js"></script>
<script src="https://cdn.bootcdn.net/ajax/libs/moment.js/2.27.0/locale/zh-cn.min.js"></script>
<script>
  var now = moment().calendar();

  console.log(now);// 输出当前日历时间

  moment().startOf('hour').fromNow(); // 相对这个小时过去了多少分钟

  var timestamp = 1593933593236; // 2020年7曰5日下午15点20分38秒
  moment(timestamp).fromNow(); // 相对时间戳多久前
</script>
```

![图片描述](https://img.mukewang.com/wiki/5f01812e092a545e12260356.jpg)

有关国际化的更多内容可以参考[文档](http://momentjs.cn/docs/#/i18n/)。

## 3. 小结

如果项目有大量处理时间的需求，可以考虑引入 `moment.js` 来处理。



### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

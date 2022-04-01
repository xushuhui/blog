# Lodash

> 是一个一致性、模块化、高性能的 JavaScript 实用工具库。

Lodash 实现了大量实用的工具方法。([官方文档](https://www.lodashjs.com/))

```javascript
<script src="https://cdn.bootcdn.net/ajax/libs/lodash.js/4.17.15/lodash.min.js"></script>
<script>
  console.log(window._);

  var arr = [1, 2, 3, 4, 5];
  var arrChunk = _.chunk(arr, 2);

  console.log(arrChunk); // 输出：[[1, 2], [3, 4], [5]]
</script>
```

## 1. 代替原生方法使用

Lodash 提供了许多原生同名方法，如数组 `forEach`、`map`、`includes` 等。

Lodash 对这些方法增加了容错，如果是原生方法，碰到值为 `null` 或者 `undefined` 会报错，在 Lodash 中会处理掉这份错误。

Lodash 在引入后，入口为全局下的 `_`。

```javascript
<script src="https://cdn.bootcdn.net/ajax/libs/lodash.js/4.17.15/lodash.min.js"></script>
<script>
  var arr = null; // 不知道出于什么原因 本来应该是个数组 但是变成了null

  _.forEach(arr, function() {

  });

  arr.forEach(function() {

  }); // 异常：Cannot read property 'forEach' of null
</script>
```

同时 Lodash 对一些方法做了优化处理，如：假使在 `forEach` 的回调中返回了 `false` ，则不会再继续遍历，达到与 `break` 类似的效果。

```javascript
<script src="https://cdn.bootcdn.net/ajax/libs/lodash.js/4.17.15/lodash.min.js"></script>
<script>
  var arr = [1, 2, 3, 4];

  var fn = function(item, index) {
    if (index === 2) {
      return false;
    }

    console.log(item);
  };

  console.log('lodash: ');
  _.forEach(arr, fn);

  console.log('native: ')
  arr.forEach(fn);
</script>
```

对项目有强健壮性和稳定性的项目，可以考虑使用 Lodash 这样的库替代原生方法进行使用，让第三方做好兼容处理。

## 2. 小结

Lodash 是非常常用的工具库，也提供了许多 ES6 提供的同名的方法，可以充当 polyfill 使用。



### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

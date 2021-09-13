# document.cookie

> Cookie 是一个请求首部，其中含有先前由服务器通过 Set-Cookie 首部投放并存储到客户端的 HTTP cookies。

cookie 可以作为单独知识了解，这里其实介绍的是 `document.cookie` 。

## 1. document.cookie

通过 document.cookie 可以获取与设置 cookie 。

![图片描述](https://img.mukewang.com/wiki/5e7a48070aa0746917041128.jpg)

## 2. 获取 cookie

```javascript
document.cookie;
```

通过 `document.cookie` 获取到的 cookie 由 cookie 的名称和值组成，由等号`=`分隔，并且可以有多条，每条 cookie 之间用分号 ‘;’ 分隔。

## 3. 设置 cookie

```javascript
var cookie = 'cookie名称=cookie值';

document.cookie = cookie;
```

设置 cookie 采用`键值对`的形式。

对应的就是 `cookie的名称` 和 `cookie值` 。

**每次只能设置一条 cookie** ，但可以同时设置这条 `cookie的属性` 。

如果需要设置多条 cookie ，则再次给 `document.cookie` 赋一个新值即可，但如果是相同名称的 cookie ，值就会被覆盖。

## 4. 设置 cookie 属性

设置 cookie 的同时可以设置这条 cookie 的属性。

```javascript
document.cookie = '名称=值; 属性1=属性值1; 属性2=属性值2';
```

看起来是可以设置多条 cookie 一样，其实只有第一对值才是 cookie 的值，后面跟的都是这条 cookie 的属性。

可以跟随的属性有：

* `path` cookie 生效的路径
* `domain` cookie 生效的域名
* `max-age` 过期时间，单位是秒
* `expires` 过期时间，为一个 UTC 时间
* `secure` 是否只能通过 https 来传递这条 cookie

这些属性具体作用可以参考 cookie 相关的内容。

设置完属性可以通过开发者工具查看。

在开发者工具的 `Application` 面板，`cookie` 分类下，每一条 cookie 都可以看到对应的属性。

![图片描述](https://img.mukewang.com/wiki/5e7a368209f2c94932702054.jpg)

## 5. 注意点

由于设置 cookie 是具有一定格式的，所以不能有字符来干扰这个格式。

```javascript
var cookie = 'code=var a = 1; var b = 2;';

document.cookie = cookie;
```

这种情况下，cookie 就不符合预期了，被切断。

cookie 中不应该含有空格、分号、逗号这些符号。

借助 `encodeURIComponent` 方法，对 cookie 的值进行编码就可以避免这类问题。

```javascript
var cookie = 'code=' + encodeURIComponent('var a = 1; var b = 2;');

document.cookie = cookie;
```

后续需要使用到这一条 cookie 的地方，再做一次解码操作即可。

> 注意：对字符串编码还可以使用 escape 方法，但已经从标准中移除，目前浏览器虽然还支持这个方法，但无法保证永远会保留这个方法，最好避免使用 escape 方法。

## 6. 小结

随着前端存储方案的增加，前端程序员访问 `document.cookie` 相对曾经减少了很多。

给 `document.cookie` 赋值可以增加一条 cookie，同时通过 `;` 相隔，来设置这条 cookie 的属性。

当设置的 cookie 带有特殊字符的时候，如 `;` 或者 `=`，应采用 `encodeURIComponent` 对内容编码，建议所有的 cookie 都进行编码。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

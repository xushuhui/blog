# JavaScript 关键字

关键字又被称为保留字。

`JavaScript` 有许多关键字，这些关键字会被 `JavaScript` 所用到，是组成 `JavaScript` 的一部分，如 `var`、`function` 都是关键字。

关键字不能被作为变量名、函数名使用。

> 随着标准的变化，保留字的列表可能也会发生变化

## 1. 已经明确的保留字

已经明确的保留字，表示目前已经完全被纳入 `ECMAScript` 标准，必须遵循语法使用这些保留字。

||||||
||||||
|break   |extends|this      |catch   |for   |
|case    |finally|throw     |try     |class |
|function|typeof |const     |if      |var   |
|continue|import |void      |debugger|in    |
|white   |default|instanceof|with    |delete|
|net     |yield  |do        |return  |else  |
|super   |export |switch    |        |      |

## 2. 未来关键字

这些关键字目前没有特殊功能，但是未来可能会有。

其中 `enum` 关键在在严格和非严格模式下都不能使用，其余的目前只在严格模式下无法使用。

```javascript
var enum = 1; // 报错：Unexpected strict mode reserved word
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5f5a624c09daba2709420084.jpg)

```javascript
var package = 1; // 通过
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5f5a625309f7e7eb07740106.jpg)

```javascript
'use strict';

var package = 1; // 报错：Unexpected strict mode reserved word
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5f5a625b0957b80709780140.jpg)

以下是被未来关键字：

||||||
||||||
|enum      |         |      |         |       |
|implements|package  |public|interface|private|
|static    |protected|let   |         |       |

## 3. 小结

尽量不要使用未来关键字作为变量，因为将来浏览器升级，可能造成一些严重的 BUG。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

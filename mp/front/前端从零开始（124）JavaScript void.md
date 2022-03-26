# void

> void 运算符 对给定的表达式进行求值，然后返回 undefined。(MDN)

表达式前面如果带有 `void` 关键字，则表达式的结果就会被忽略，并将 `undefined` 作为结果。

从业务上来看，void 关键字并不常用。

## 1. 用于调用函数表达式

当想让一个函数立即实行的时候，需要让 JavaScript 将一个函数识别为表达式，void 关键字就能起到这个作用。

```javascript
void function() {
  alert('马上执行！冲冲冲！');
}();
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb56f510ad5d69413280476.jpg)

但有局限性，如果需要获取到函数的返回值，就不能使用 void。

```javascript
var num1 = 1;
var num2 = 2;

var result = void function(number1, number2) {
  return [number1 + number2, number1 * number2];
}(num1, num2);

result.forEach(function(res) {
  console.log(res);
});
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb56f5a0a1324c213280476.jpg)

如这个例子，函数返回了两数之和与两数之积的结果，但因为 void 关键字，实际 result 变量被赋值为 `undefined`，导致程序无法正常执行。

## 2. 内联在 HTML 中，阻止 a 标签的默认事件

`<a>` 标签的 `href` 属性，可以用来执行 `JavaScript` 代码。

通常可以这么写：

```javascript
<a href="javascript: void;">跳转！</a>
<a href="javascript: void 0;">跳转！</a>
<a href="javascript: void (0);">跳转！</a>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb56f740a1e087d13280476.jpg)

这三行代码的效果是一样的。

如果 `<a>` 标签的 `href` 属性是 `javascript:表达式;`，则会执行表达式的内容，并将页面的内容设置为表达式的结果，如果表达式的结果是 undefined，则什么都不做。

根据这个规则，`void` 就起到了作用，但其实不写表达式，依然能达到这个效果。

```javascript
<a href="javascript:;">跳转！</a>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb56f840a699bf210840432.jpg)

这样的效果和上面使用 `void` 关键字的方式是等价的，这也是常用的方式。

但碰到需要使用 `<a>` 标签执行函数的时候，void 就变得相对关键。

```javascript
<script>
  function log(who) {
    console.log('点击了：' + who);

    return who;
  }
</script>

<a href="javascript: log('add');">添加</a>
<a href="javascript: log('update');">修改</a>
<a href="javascript: log('delete');">删除</a>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb56f8e0a759d8c10840432.jpg)

这种情况如果不加 `void`，页面内容就会发生改变，因为 log 函数存在非 undefined 的返回值。

```javascript
<script>
  function log(who) {
    console.log('点击了：' + who);

    return who;
  }
</script>

<a href="javascript: void log('add');">添加</a>
<a href="javascript: void log('update');">修改</a>
<a href="javascript: void log('delete');">删除</a>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5eb56fa60a47c91110840432.jpg)

加上 void 一样，结果就符合预期了，具体的添加、删除操作，再通过绑定对应的事件来实现。

## 3. 小结

void 的使用场景有限，但在某些情况下可以提高代码的健壮性，如明确不需要结果的场景下，加上 void 关键字，这样可以避免未来表达式结果的改变带来的问题。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

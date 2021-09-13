# 异常处理

异常处理可以使程序在流程上更加完善。

在 JavaScript 中可以使用 `throw` 抛出异常，使用 `try ... catch` 捕获错误。

## 1. throw

> throw 语句用来抛出一个用户自定义的异常。(MDN)

throw 用于抛出一个异常，这种异常通常是程序出现了不符合预期的错误。

```javascript
alert('出错前');

throw '发生了一个错误！';

alert('出错后');
```

![图片描述](https://img.mukewang.com/wiki/5eafbdc50a4abd5717600678.jpg)

当出现 `throw` 时，程序将会中断执行。

如果 `throw` 发生在 `try ... catch` 中，则会执行 `catch` 中的代码块，同时将异常信息带给 catch。

## 2. try … catch

> try…catch 语句标记要尝试的语句块，并指定一个出现异常时抛出的响应。

`try ... catch` 可以用于捕获异常，当出现 throw 时，会结束 `try` 内的代码执行，直接进入到 `catch`，执行 `catch` 内的代码块。

```javascript
try {
  alert('出错前');

  throw '发生了一个错误！';

  alert('出错后');
} catch (e) { // e 是错误信息，名字随意，符合变量命名规范就行
  alert('出错了！错误是：' + e);
}
```

![图片描述](https://img.mukewang.com/wiki/5eafbdd80a7bab9517600678.jpg)

需要注意的是，以前 catch 后面的错误参数是必须接收的，否则会报错。

但 ES2019 中有一个提案，可以选择性的提供给 catch 参数，所以最新的 chrome 不传递错误参数也是可以的。

```javascript
try {
  alert('出错前');

  throw '发生了一个错误！';

  alert('出错后');
} catch {
  alert('出错了！');
}
```

![图片描述](https://img.mukewang.com/wiki/5eafbde80a8ce81c17600678.jpg)

由于是比较新的提案，所以建议没有工具参与代码编译时，还是写上错误参数的接收，避免因浏览器兼容性造成的问题。

在 try 后面还可以跟 `finally` 部分，即无论 try 中的代码块是否抛出错误，都会执行 `finally` 代码块。

```javascript
try {
  alert('开始请求数据，loading 显示');

  throw '没有网络';

  alert('请求结果是：..编不下去了，反正到不了这里');
} catch (e) {
  alert('出现错误：' + e);
} finally {
  alert('关闭 loading');
}
```

![图片描述](https://img.mukewang.com/wiki/5eafbdfc0a20c52417600678.jpg)

## 3. 可以写条件的 catch 语句

部分文献记载了如下格式的 try … catch 语法。

```javascript
try {
  throw 'error';
} catch (e if e instanceof TypeError) {
  console.log('TypeError');
} catch (e if e instanceof ReferenceError) {
  console.log('ReferenceError');
} catch (e) {
  console.log(e);
}
```

但目前主流浏览器基本都无法正常运行这种语法的 try … catch 语句，所以不要使用。

如果有类似的需求，可以使用 if 来代替。

```javascript
try {
  throw 'error';
} catch (e) {
  if (e instanceof TypeError) {
    console.log('TypeError');
  } else if (e instanceof ReferenceError) {
    console.log('ReferenceError');
  } else {
    console.log(e);
  }
}
```

## 4. Error 对象

> 通过 Error 的构造器可以创建一个错误对象。当运行时错误产生时，Error 的实例对象会被抛出。Error 对象也可用于用户自定义的异常的基础对象。(MDN)

通常在使用 throw 抛出异常时，会抛出一个 `Error` 对象的实例。

```javascript
try {
  throw new Error('主动抛出一个错误');
} catch (e) {
  console.error(e);
}
```

![图片描述](https://img.mukewang.com/wiki/5eafbe160a3acb8c17600678.jpg)

和大部分内置对象一样，Error 实例也可以不使用 `new` 关键字创建。

```javascript
try {
  throw Error('主动抛出一个错误');
} catch (e) {
  console.error(e);
}
```

![图片描述](https://img.mukewang.com/wiki/5eafbe230aca231d17600678.jpg)

抛出 Error 实例，可以得到出现异常的文件和对应的行号。

除了 `Error` ，还有几种预定义好语义的异常对象。

## 5. 其他异常对象

* `URIError` 表示以一种错误的方式使用全局 URI 处理函数而产生的错误；
* `TypeError` 值的类型非预期类型时发生的错误；
* `SyntaxError` 尝试解析语法上不合法的代码的错误；
* `ReferenceError` 当一个不存在的变量被引用时发生的错误；
* `RangeError` 当一个值不在其所允许的范围或者集合中抛出的异常；
* `InternalError` 表示出现在 JavaScript 引擎内部的错误。非标准对象，不建议使用；
* `EvalError` 本对象代表了一个关于 eval 函数的错误。此异常不再会被 JavaScript 抛出，但是 EvalError 对象仍然保持兼容性。

这些异常对象的使用和 `Error` 几乎一致。

浏览器碰到对应的异常，也会抛出。

```javascript
try {
    console.log(notDefinedVariable);
} catch (e) {
    console.error(e);
}
```

![图片描述](https://img.mukewang.com/wiki/5eafc5bb0afe160a16080692.jpg)

因为 `notDefinedVariable` 并没有定义，所以浏览器会抛出 `ReferenceError` 异常，同时提示变量没有定义。

## 6. 小结

完整的产品业务逻辑流程，基本都要 `try ... catch` 参与控制，因为出现异常时，还要有对应的动作，如网络请求异常，则提示用户重试，或主动进行超时重新请求操作。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

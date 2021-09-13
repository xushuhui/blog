# 认识输入框 input 标签

输入框是我们网页中常见的元素，登录、注册、个人资料、收货地址等都需要用到输入框。而在 HTML 中，输入框的类型和属性都有很多，我们可以根据实际需求来决定使用哪一类输入框。

## 1. 输入框的类型

在 HTML 中， input 标签表示输入框，而输入框有很多类型，比如普通文本输入框、密码框、邮箱框（只能输入邮箱格式的内容）、网址框（只能输入框网址格式的内容）、数字框（只能输入数字）、单选框、多选框等。我们可以改变 input 标签的 `type` 属性来显示不同的输入框类型。

## 2. 输入框的使用

input 的类型有很多，我们依次为大家介绍。

### 2.1. 普通输入框

普通输入框既可以输入任意内容，没有格式和内容要求。代码如下：

```javascript
<input type='text'>
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cc00090c2c5c05630051.jpg)

input 标签的 type 属性默认为`text`。

### 2.2. 密码框

把 input 的 type 属性设置为 `password`则表示密码框。密码框既输入的内容为密文显示，呈现的效果为实心黑点，不会显示具体的输入内容。代码如下：

```javascript
<input type='password'>
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cbe0095a831605330061.jpg)

### 2.3. 邮箱框

把 input 的 type 设置为 `email`则表示邮箱框，那么输入的内容会有规则限制，输入的内容必须包含 @，且 @ 后必须有内容才满足验证规则，否则会有错误提示。代码如下：

```javascript
<input type='email'>
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cc21094ab7d106060152.jpg)

![图片描述](https://img.mukewang.com/wiki/5f07cc35095b786a05010146.jpg)

![图片描述](https://img.mukewang.com/wiki/5f07cc4509bdae4d03950126.jpg)

### 2.4. 网址框

把 input 的 type 设置为 `url`则表示网址框，那么输入的内容会有规则限制，输入的内容需要以 `http://` 或者 `https://` 开头 ，且 @ 后必须有内容才满足验证规则，否则会有错误提示。代码如下：

```javascript
<input type="url">
```

![图片描述](https://img.mukewang.com/wiki/5f07cc840991f49103540120.jpg)

> **Tips**：这里的网站和我们平时输入的网站不同，前面必须加上网络协议，既 http:// 或者 https://

### 2.5. 数字框

把 input 的 type 设置为 `number`则表示数字框，那么就只能输入数字，输入其他字符无效，且输入框最右侧会有加减按钮。代码如下：

```javascript
<input type='number'>
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cd8d0929100404220159.jpg)

### 2.6. 单选框

把 input 的 type 属性设置为 `radio`则表示单选框，因为 input 标签是单标签，所有单选框的内容直接写在 input 标签后面，如果单选框有多个选项，则需要写多个 input 标签，且每个 input 标签必须添加 `name` 属性，否则不能成为一组的单选框（既可以选多个）。代码如下：

```javascript
<input type="radio" name='gender'> 男
<input type="radio" name='gender'> 女
<input type="radio" name='gender'> 保密
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cda709da7f0a03110084.jpg)

单选框可以添加 `checked` 属性，表示默认选中一项。代码如下：

```javascript
<input type="radio" checked> 男
<input type="radio"> 女
<input type="radio"> 保密
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cdc509e87dab02720093.jpg)

> **Tips**：如果给多个单选框设置 `checked` 属性，则会默认选中最后一个设置 `checked`属性的选项。

### 2.7. 多选框

把 input 的 type 属性设置为 `checkbox`，则表示多选框。多选框和单选框一样，需要设置 `name`属性，且多选框也可以设置 `checked` 属性表示默认选中，多选框的 `checked` 属性可以设置多个。代码如下：

```javascript
<input type="checkbox" name="ball" checked> 篮球
<input type="checkbox" name="ball">足球
<input type="checkbox" name="ball" checked>排球
<input type="checkbox" name="ball">乒乓球
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cdd9093118c905020147.jpg)

### 2.8. 占位符

INPUT 标签可以设置 `placeholder`属性为占位符。占位符的作用为输入提示，如果输入框没有内容，则会显示占位符的内容，一旦输入框有内容，则会显示输入框的内容，占位符的内容消失。代码如下：

```javascript
<input type="text" placeholder="请输入内容">
```

效果如下：

![图片描述](https://img.mukewang.com/wiki/5f07cdee0964352e04140123.jpg)

## 3. 注意事项

1. 单选框和多选框必须给 `name` 属性，`name` 属性的值为自定义内容；
2. 邮箱框和网站框对输入内容有限制，需按照其验证规则输入正确的内容；
3. 占位符一般为输入提示，所以占位符的内容为此输入框的提示，输入内容后消失。

## 4. 真实案例分享

京东

```javascript
 <div>
     <input id="loginname" type="text"  placeholder="邮箱/用户名/登录手机"/>
</div>

<div id="entry" class="item item-fore2">
    <input type="password" id="nloginpwd"  placeholder="密码"/>
    <span >大小写锁定已打开</span>
</div>
```

简书官网

```javascript
<div >
    <input placeholder="手机号或邮箱" type="text" />
</div>
<div >
   <input placeholder="密码" type="password" />
</div>
```

## 5. 小结

1. 改变 input 的 `type` 属性来设置输入框不同的类型。
2. 单选框和多选框需要设置 `name` 属性。
3. 邮箱框、网站框需要按照一定规则书写内容。
4. 输入框类型比较多，可以根据实际需求决定使用哪个类型。

![图片描述](https://img.mukewang.com/wiki/5f63049609b82c3d14010715.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

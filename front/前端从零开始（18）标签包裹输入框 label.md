# 认识 label 标签

label 标签的作用为 input 元素定义标注（标记）。label 元素不会向用户呈现任何特殊效果。不过，它为鼠标用户改进了可用性。如果您在 label 元素内点击文本，就会触发此控件。就是说，当用户选择该标签时，浏览器就会自动将焦点转到和标签相关的表单控件上。例如：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07ce4209f8dd3908560474.jpg)

## 1. label 标签的使用

label 需要和 input 标签搭配一起使用。LABEL 标签的 `for` 属性需要和 input 的 `id` 属性一致，这样才能点击 label 标签的内容使对应的 input 输入框自动获取焦点。**代码如下：**

```javascript
<label for="username">用户名</label>
<input type="text" placeholder="请输入内容" id='username'>
```

效果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07ce5209d5484803930114.jpg)

## 2. label 标签的作用

表单控件都是内联元素所以他俩会在一行显示。在网页中当我们点击 label 标签的内容或文本框时都会在文本框中出现光标，这个就是 label 标签的功能了。说白了 label 标签就是他所关联的表单控件的延伸，即鼠标点击了他，就会出现和点击他所关联的表单控件一样的效果。而这前提是 label 的 for 属性的属性值与想要关联的表单控件的 id 一样。

除了增强用户体验外，还为行动不便人士上网提供了便利。比如说，视力障碍者是借助“网页朗读器”发出的声音来浏览网页的，若没有 label 标签的关联，上网者就在脑海中不能想象出那种对应性，不能很好理解网站表单所想表达的内容。再比如，肢体有缺陷的上网者对于鼠标的控制是很辛苦的，运用 label 后点击的目标变大了，有利于操作。

## 3. 注意事项

1. label 标签里面需要写内容，才会在页面显示。
2. label 标签的 `for` 属性必须和 input 输入框的 `id` 一致。

## 4. 真实案例分享

京东

```javascript
 <div>
     <label for="loginname">用户名</label>
     <input id="loginname" type="text" placeholder="邮箱/用户名/登录手机"/>
</div>

<div id="entry">
    <label for="nloginpwd">密码</label>
    <input type="password" id="nloginpwd" placeholder="密码"/>
    <span class="capslock">大小写锁定已打开</span>
</div>
```

拼多多

```javascript
<div>
    <label for="user-mobile"></label>
    <input  id="user-mobile" placeholder="手机号码">
</div>
<div>
    <label for="input-code"></label>
    <input type="tel" id="input-code" placeholder="验证码">
</div>
```

## 5. 小结

1. label 标签一般和 input 一起使用。
2. label 标签的内容会和 input 在同一排显示。
3. 点击 label 标签的内容，会让 input 输入框获取焦点。
4. label 标签的 `for` 属性必须和 input 的 `id`属性一致。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f6308a909f077a414530721.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

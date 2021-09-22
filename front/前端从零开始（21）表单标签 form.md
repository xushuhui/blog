# 认识表单标签 form

表单是我们网页中常见的场景，比如登录、注册，填写个人信息，填写收货地址等。在 HTML 当中创建表单和表格一样，也是需要一组标签，而且表单的属性和表单的元素都比较多，我们可以根据实际需求来定制我们的表单内容。表单里可以嵌套各个类型的输入框，比如普通输入框、密码框等，也可以嵌套单选框、多选框以及下拉菜单。

## 1. form 表单的使用

form 标签和 ul select 标签类似，代表表单整体，而里面嵌套的元素则是表单具体的内容。我们来做一个用户名和密码的表单，这需要用到之前我们讲的 label 标签和 input 标签的知识，**代码如下：**

```javascript
<form>
    <label for="username">用户名</label>
    <input type="text" id='username'>
    <br>
    <label for="password">密码</label>
    <input type="password" id="password">
</form>
```

效果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cfcb09dc454d03030096.jpg)

表单呈现的形式和普通输入框无异，但它的作用就是我们要做提交表单的操作（既我们需要把用户输入的信息传给后台），那么普通的输入框就做不到这个功能了。那么提交表单的时候，我们可以给 form 标签加上一个 `method` 属性，这个属性表示当前提交表单的方式，一般为 `get` 或者 `post`，这个需要后台先行告知。form 标签还有一个 `action` 属性，表示表单提交的地址，这个也需要后台先行告知。

## 2. 经验分享

1. 我们在编写表单时，需要先行和后台人员沟通，以获取到表单的提交方式和提交地址；
2. 表单里面可以嵌套多个元素，这个需要根据实际需求来决定；
3. 所有表单控件（文本框、文本域、按钮、单选框、复选框等）都必须放在 form 标签之间（否则用户输入的信息可提交不到服务器上）。

## 3. 真实案例分享

京东

```javascript
<form method="post">
    <div>
        <label for="loginname"></label>
        <input id="loginname" type="text" placeholder="邮箱/用户名/登录手机">
    </div>
    <div id="entry" class="item item-fore2">
        <label class="login-label pwd-label" for="nloginpwd"></label>
        <input type="password" id="nloginpwd" placeholder="密码">
        <span><b>大小写锁定已打开</b></span>
    </div>
    <div>
      <div>
        <span>
          <a href="//passport.jd.com/uc/links?tag=safe">忘记密码</a>
          </span>
        </div>
    </div>
    <div>
      <div>
        <a href="javascript:;">登    录</a>
      </div>
    </div>
</form>
```

简书官网

```javascript
<form id="new_session" action="/sessions" method="post">
      <div>
        <input placeholder="手机号或邮箱" type="text">
      </div>
    <div>
      <input placeholder="密码" type="password">
    </div>
    <div>
      <input type="checkbox"><span>记住我</span>
    </div>
    <div>
      <a href="">登录遇到问题?</a>
      <ul>
        <li><a href="/users/password/mobile_reset">用手机号重置密码</a></li>
        <li><a href="/users/password/email_reset">用邮箱重置密码</a></li>
        <li><a target="_blank" href="/p/9058d0b8711d">无法用海外手机号登录</a></li>
        <li><a target="_blank" href="/p/498a9fa7da08">无法用 Google 帐号登录</a></li>
      </ul>
    </div>
    <button>
      <span></span>
      登录
   </button>
</form>
```

## 4. 小结

1. form 标签代表表单整体，在页面上并无特殊样式显示。
2. form 标签里面嵌套的内容是表单的内容，可以是输入框、单选框、多选框、下拉菜单、文本域等。
3. form 标签需要设置 `method`属性，为提交表单的方式。
4. form 标签需要设置 `action`属性，为提交表单的地址。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f6306970924e94514430776.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

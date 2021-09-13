# 表单校验

本篇主要介绍使用 JavaScript 进行表单验证。

表单验证并不是 JavaScript 提供的某种特性，而是结合各种特性达到的一种目的，是需求的产物。

所有线上产品的表单几乎都有验证，如注册时要求“用户名 6-16 位”，验证会由 JavaScript 来完成，通常为了安全性和准确性，服务端会再次做一遍验证。

## 1. 验证目标

表单用于收集信息，从 HTML 上讲，表单内容使用 `form` 标签进行包裹。

```javascript
<form action="/login">
  <label>
    用户名：<input type="text">
  </label>

  <label>
    密码：<input type="text">
  </label>

  <div>
    <button type="submit">登入</button>
  </div>
</form>
```

这就是一个相对简单的表单，其中包含文本框（input 标签）与按钮（button 标签），并使用 form 标签进行包裹。

利用 form 标签，再触发其 submit 事件时，会将表单内容收集后提交个体 `action` 属性配置的路径。

单其实把 form 标签去掉，在页面上展示的效果几乎是一样的。

```javascript
<label>
  用户名：<input type="text">
</label>

<label>
  密码：<input type="text">
</label>

<div>
  <button type="submit">登入</button>
</div>·
```

所以自出现 AJAX 技术后，很多开发者都不再书写 form 标签，直接使用其他元素对表单内容进行包裹，因为业务上可能不需要使用 form 标签的特性来提交表单。

其实不论是使用表单，还是不使用表单，表单的验证都是针对所有表单项的，即输入框、单选项、多选项等。

在表单提交之前，需要对写着表单项的内容做校验，然后拦截提交操作。

## 2. 获取表单内容

获取表单内容，实际上就是取到表单项对应的 DOM 节点的值。

获取 DOM 节点的方式非常多，前面的章节也有介绍。

```javascript
<style>
  h3 {margin-top: 0;color: #4caf50;}
  .login {width: 300px;padding: 32px;box-shadow: 2px 2px 10px rgba(0, 0, 0, .1);position: fixed;top: 40%;left: 50%;transform: translate(-50%, -50%);}
  .form-item {display: flex;margin-bottom: 16px;border-bottom: 1px solid #ccc;}
  .form-item .title {width: 70px;color: #666;font-size: 14px;}
  .form-item .content {flex: 1;}
  .form-item .content input {width: 100%;border: 0 none;padding: 2px 8px;outline: none;font-size: 16px;}
  .login-btn {width: 100%;border: 0 none;background-color: #4caf50;color: white;margin-top: 16px;outline: none;height: 32px;}
  .login-btn:active {background-color: #2da050;}
</style>

<form name="login-form" class="login">
  <h3>登入</h3>
  <label class="form-item">
    <div class="title">用户名</div>
    <div class="content">
      <input id="account" class="account" name="account" type="text">
    </div>
  </label>

  <label class="form-item">
    <div class="title">密码</div>
    <div class="content">
      <input name="pwd" type="password">
    </div>
  </label>

  <div>
    <button class="login-btn" type="submit">登入</button>
  </div>
</form>

<script>
  var account1 = document.getElementById('account');
  var account2 = document.getElementsByName('account');
  var account3 = document.getElementsByClassName('account');

  alert(account1 === account2[0]);
  alert(account1 === account3[0]);

  var account4 = document.forms['login-form']['account'];

  alert(account1 === account4);

  console.log(document.forms['login-form'].elements);
</script>
```

前三种获取节点的方式读者都已经熟悉了。

`account4` 的获取方式稍微有点不一样，document.forms 是文档内的表单集合，其可以通过表单的 id 和 form 的属性，取到具体的某个表单。

取到表单后，还可以直接使用表单项的 name 属性取到对应的表单项，使用 `elements` 可以取到这个表单下的所有表单项。

## 3. 校验表单项

获取到表单项后，就可以对表单项的值做判断了，如密码必须是 6-16 位：

```javascript
<style>
  h3 {margin-top: 0;color: #4caf50;}
  .login {width: 300px;padding: 32px;box-shadow: 2px 2px 10px rgba(0, 0, 0, .1);position: fixed;top: 40%;left: 50%;transform: translate(-50%, -50%);}
  .form-item {display: flex;margin-bottom: 16px;border-bottom: 1px solid #ccc;}
  .form-item .title {width: 70px;color: #666;font-size: 14px;}
  .form-item .content {flex: 1;}
  .form-item .content input {width: 100%;border: 0 none;padding: 2px 8px;outline: none;font-size: 16px;}
  .login-btn {width: 100%;border: 0 none;background-color: #4caf50;color: white;margin-top: 16px;outline: none;height: 32px;}
  .login-btn:active {background-color: #2da050;}
</style>

<form name="login-form" class="login" action="https://imooc.com">
  <h3>登入</h3>
  <label class="form-item">
    <div class="title">用户名</div>
    <div class="content">
      <input autocomplete="off" id="account" class="account" name="account" type="text">
    </div>
  </label>

  <label class="form-item">
    <div class="title">密码</div>
    <div class="content">
      <input autocomplete="off" name="pwd" type="password">
    </div>
  </label>

  <div>
    <button class="login-btn" type="submit">登入</button>
  </div>
</form>

<script>
  var loginForm = document.forms['login-form'];
  var pwdEle = loginForm.pwd;

  loginForm.onsubmit = function(e) {
    var pwd = pwdEle.value;

    if (pwd.length < 6 || pwd.length > 16) {
      alert('密码长度 6-16');
      return false; // 可以使用 return e.preventDefault() 代替
    }

    setTimeout(function() {
      alert('登入成功，马上跳转！');
    }, 1000);
  };
</script>
```

这里获取到了表单元素，同时给表单的事件处理器属性 `onsubmit` 赋值，表示在表单被提交的时候做的事情。

在事件处理器中，通过输入框的 `value` 属性获取到了输入的值，对值进行了长度判断，如果长度不正确则返回 false，表单则会终止提交。

如果正确，则会根据 form 标签的 target 属性进行提交。

需要注意的是，这里如果使用 `addEventListener` 来监听 `submit` 事件，必须使用 `preventDefault` 来阻止默认事件，即阻止表单提交，不能使用 `return false;`。

```javascript
var loginForm = document.forms['login-form'];
var pwdEle = loginForm.pwd;

loginForm.addEventListener('submit', function(e) {
  var pwd = pwdEle.value;

  if (pwd.length < 6 || pwd.length > 16) {
    alert('密码长度 6-16');
    e.preventDefault(); // 代替return false
    return;
  }

  setTimeout(function() {
    alert('登入成功，马上跳转！');
  }, 1000);
});
```

## 4. 不使用 form 提交表单

不使用 form 标签来提交表单，通常都是使用 AJAX 进行数据交互的情况。

这个时候就不需要拦截 form 的提交行为了。

```javascript
<style>
  h3 {margin-top: 0;color: #4caf50;}
  .login {width: 300px;padding: 32px;box-shadow: 2px 2px 10px rgba(0, 0, 0, .1);position: fixed;top: 40%;left: 50%;transform: translate(-50%, -50%);}
  .form-item {display: flex;margin-bottom: 16px;border-bottom: 1px solid #ccc;}
  .form-item .title {width: 70px;color: #666;font-size: 14px;}
  .form-item .content {flex: 1;}
  .form-item .content input {width: 100%;border: 0 none;padding: 2px 8px;outline: none;font-size: 16px;}
  .login-btn {width: 100%;border: 0 none;background-color: #4caf50;color: white;margin-top: 16px;outline: none;height: 32px;}
  .login-btn:active {background-color: #2da050;}
</style>

<div class="login">
  <h3>登入</h3>
  <label class="form-item">
    <div class="title">用户名</div>
    <div class="content">
      <input autocomplete="off" id="account" class="account" name="account" type="text">
    </div>
  </label>

  <label class="form-item">
    <div class="title">密码</div>
    <div class="content">
      <input autocomplete="off" name="pwd" type="password">
    </div>
  </label>

  <div>
    <button class="login-btn" type="button">登入</button>
  </div>
</div>

<script>
var loginBtn = document.querySelector('.login-btn');
var pwdEle = document.querySelector('[name="pwd"]');

function login(cb) {
  // 假装登入花了 1 秒
  setTimeout(function() {
    alert('登入成功');
    cb && cb();
  }, 1000);
}

loginBtn.addEventListener('click', function() {
  var pwd = pwdEle.value;

  if (pwd.length < 6 || pwd.length > 16) {
    alert('密码长度 6-16');
    return;
  }

  login(function() {
    window.location.href = 'https://imooc.com';
  });
});
</script>
```

使用这种方式，就可以自主控制流程，不需要再考虑 form 标签的行为。

## 5. 小结

校验表单非常常见，校验表单的场景很多时候远没有本篇介绍的这么简单，有时候数据校验的格式非常复杂，需要结合正则、校验算法等方式来解决，如严格的身份证验证就需要结合身份证算法。

但表单的校验总的来说都遵循获取表单元素、获取表单元素的值、对值进行判断、根据判断结果做下一步动作。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

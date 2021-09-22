# HTML 表单

大部分情况下，网页中展示的数据或者图片或者多媒体效果都是静态的数据，但是有时用户需要通过网页跟服务器进行动态的数据交互，例如登录、注册等操作，这时需要用到表单。表单其实是由一些输入框、单选、复选框等元素组成，用户点击提交按钮，浏览器会将这些数据发送给表单设置的 URL ，URL 对应的服务器接收到数据并处理之后，再将数据返回给浏览器。

## 1. form

form 是一个表单容器，其中包含需要提交的数据，和相应的提交按钮，例如：

```javascript
<form action="/user/login" method="post">
   <h3>登录</h3>
   <div>
      <label for="user_name">用户名</label>
      <input id="user_name" name="user_name" ><!-- 用户名表单 -->
   </div>
   <div>
      <label for="password">密码</label>
      <input id="password" name="password" type="password" ><!--密码表单-->
   </div>
    <div>
        <button  type='submit'>登录</button>
    </div>
</form>
```

代码解释：当用户点击登录按钮时，使用 post 方式把用户输入的用户名和密码这两项数据提交到，网站服务器的 /user/login 页面下，这时 /user/login 页面的程序会对用户提交过来的用户名和密码与服务器的数据库中用户注册时填写的用户名和密码进行比对，如果相同就登陆成功。

form 标签有很多的属性，下面分别来介绍。

### 1.1 accept-charset 属性

accept-charset 属性用于设定使用哪种字符集处理表单数据，常用值 utf-8、ISO-8859-1、GB2312 等。其中 utf-8、ISO-8859-1、GB2312 分别表示 Unicode 编码、拉丁字母字符编码、简体中文字符编码，如果需要使用多个字符编码，则需要使用逗号隔开。如果没有设置 accept-charset 属性，则默认使用与 HTML 文档一致的编码。

```javascript
<form accept-charset='utf-8'><!--使用Unicode编码-->
</form>
<form accept-charset='ISO-8859-1'><!--使用拉丁编码-->
</form>
<form accept-charset="utf-8,ISO-8859-1"><!--同时支持两种编码-->
</form>
```

### 1.2 action 属性

用于设定表单提交的服务器的 URL，可以是相对路径和绝对路径。例如设置：

```javascript
<form action="https://www.baidu.com" method="post">
   <h3>登录</h3>
   <div>
      <label for="user_name">用户名</label>
      <input id="user_name" name="user_name"  ><!-- 用户名表单 -->
   </div>
   <div>
       <label for="password">密码</label>
       <input id="password" name="password" type="password" ><!--密码表单-->
   </div>
   <div>
       <button  type='submit'>登录</button>
   </div>
</form>
```

代码解释：上述代码使用 post 方式向百度首页提交数据。

### 1.3 autocomplete 属性

用于设定是否启用自动完成功能，这个功能类似于历史记录的功能，之前提交过的表单下次在此输入时，浏览器会记录下来历史记录，可用值 on/off。

```javascript
<form autocomplete='on'><!--默认开启自动完成-->
<h2>开启自动完成</h2>
用户名:<input type="text" name="fname"><br>
邮箱: <input type="email" name="email"><br>
</form>

<form autocomplete='off'><!--关闭自动完成-->
<h2>关闭自动完成</h2>
用户名:<input type="text" name="fname"><br>
邮箱: <input type="email" name="email"><br>
</form>
```

**注意：除了 Opera 其他主流浏览器都支持 autocomplete。**

### 1.4 enctype 属性

enctype 用于定义表单数据提交到服务器的过程中如何对数据进行编码，可选值有：

* application/x-www-form-urlencoded；
* multipart/form-data；
* text/plain

**第一种：application/x-www-form-urlencoded**

默认方式是第一种 `application/x-www-form-urlencoded`，当使用 form 表单提交数据时，需要对数据进行编码，因为 URL 的数据是通过 `?key=value&key=value&` 的方式传输给服务器，这其中有一些作为功能性质的特殊字符例如 `& ? =`，如果数据中带有这些字符而且未进行编码的话，可能会导致传输数据发生错误，例如原本是 `?key=value` ，如果 value 是 `123&456` 的话，那么 结果服务器中接收到的 value 值变成 123，且多出一个 `key=456`，这种编码方式叫做 URI 百分号编码，其标准收录在 RFC3986 中。

当设置成 `multipart/form-data` 时，浏览器不对字符进行编码，这种编码方式通常适用于上传文件。

默认方式是第一种 `application/x-www-form-urlencoded`，对数据进行编码。

**为什么要对提交的数据进行编呢？**

当使用 form 表单提交数据时，需要对数据进行编码，因为 URL 的数据是通过 `?key=value&key=value&` 的方式传输给服务器，这其中有一些作为功能性质的特殊字符例如 `& ? =`，如果数据中带有这些字符而且未进行编码的话，可能会导致传输数据发生错误，例如原本是 `?key=value` ，如果 value 是 `123&456` 的话，那么 结果服务器中接收到的 value 值变成 123，且多出一个 `key=456`，这种编码方式叫做 URL 百分号编码，其标准收录在 RFC3986 中。

**第二种：multipart/form-data**

当设置成 `multipart/form-data` 时，浏览器不对字符进行编码，这种编码方式通常适用于上传文件；

**第三种：text/plain**

使用第三种方式 `text/plain` 时，浏览器将请求参数放入 body 作为字符串处理，这种方式用于传输原生的 HTTP 报文，不适用任何格式化处理。

### 1.5 method 属性

使用表单提交数据时，实际上只发送一个 HTTP 协议的数据请求，HTTP 协议有很多种数据请求方式，这个 method 属性用于设定 HTTP 请求的方式。常用的方式有 post、get，当未设置时默认使用 get 方式。除了常用方式之外，根据服务器 HTTP 网关的设置，还可以支持：

* options 客户端查看服务器的配置；
* head 用于获取报文头，没有 body 实体；
* delete 请求服务器删除指定页面；
* put 请求替换服务器端文档内容；
* trace 用于诊断服务器；
* connect 将连接设置成管道方式的代理服务器，用于 HTTP1.1

### 1.6 novalidate 属性

当设置时，表单不会验证输入，否则的话当点击提交按钮时，浏览器会根据表单输入框的类型来验证输入内容是否合法，例如：

```javascript
<form action="" method="post" novalidate><!--提交时不验证-->
<h2>提交时不验证</h2>
        <h3>登录</h3>
        <div>
            <label for="user_name">用户名</label>
            <input id="user_name" name="user_name"  autofocus="" required=""><!-- 用户名表单 -->
        </div>
        <div>
            <label for="password">密码</label>
            <input id="password" name="password" type="password" autocomplete="off" value="" required=""><!--密码表单-->
        </div>
        <div>
            <button  type='submit'>登录</button>
        </div>
</form>

<form action="" method="post"><!--提交时验证-->
<h2>提交时验证</h2>
        <h3>登录</h3>
        <div>
            <label for="user_name">用户名</label>
            <input id="user_name" name="user_name"  autofocus="" required=""><!-- 用户名表单 -->
        </div>
        <div>
            <label for="password">密码</label>
            <input id="password" name="password" type="password" autocomplete="off" value="" required=""><!--密码表单-->
        </div>
        <div>
            <button  type='submit'>登录</button>
        </div>
</form>
```

上述代码，当点击登陆按钮时，浏览器会验证用户名是否输入，密码是否输入，未输入的话浏览器会提示，并且提交不成功。novalidate 属性适用于：form，以及以下类型的 标签：text, search, url, telephone, email, password, date pickers, range 以及 color，复杂的规则需要结合 required 属性验证。

### 1.7 target 属性

这个属性用户设置表单提交之后浏览器的跳转地址，默认是在当前页面打开新地址，可选值有：

* _blank 新窗口；
* _self 默认，当前窗口；
* _parent 父窗口；
* _top 最顶层窗口；
* _framename 指定的框架

例如：

```javascript
<form action="https://www.baidu.com" method="post" target="_blank">
<h2>表单提交后跳转到新窗口</h2>
        <h3>	登录 	</h3>
        <div>
            <label for="user_name">用户名</label>
            <input id="user_name" name="user_name"  autofocus="" required=""><!-- 用户名表单 -->
        </div>
        <div>
            <label for="password">密码</label>
            <input id="password" name="password" type="password" autocomplete="off" value="" required=""><!--密码表单-->
        </div>
        <div>
            <button  type='submit'>登录</button>
        </div>
</form>
```

上述代码会在提交完数据之后，在一个新的浏览器界面中打开百度的首页。

## 2. input

表单内的标签，用于记录用户输入信息，可以是文本框、复选框、单选框、日期、数字、日期、按钮、文件、密码、隐藏域等等。

### 2.1 文本框

文本框是表单中最常见的表单控件，没有设置样式的话，默认是一个长方形的输入框，获取焦点时外边框变蓝色，当设置 type=text 就定义了一个文本框，例如：

```javascript
<form action="https://www.baidu.com" method="post" target="_blank">
        <h3>	登录 	</h3>
        <div>
            <label for="user_name">用户名</label>
            <input id="user_name" type='text' name="user_name"  autofocus="" required=""><!-- 用户名表单 -->
        </div>
        <div>
            <button  type='submit'>登录</button>
        </div>
</form>
```

### 2.2 复选框

复选框用于对多个选项进行复合选择，网站开发中让用户选择兴趣爱好、技能标签等可以使用复选框，使用 type=checkbox 定义复选框：

```javascript
<div class="div_title_question_all">
    <div id="divTitle5" class="div_title_question">从未使用网约车的原因是：<span class="qtypetip"> 【多选题】</span></div>
</div>
<div class="div_table_radio_question" id="divquestion5">
    <ul class="ulradiocheck" style='list-style-type:none'>
        <li style="width:99%;"><input  id="q5_1" type="checkbox" name="q5" value="不熟悉叫车程序"><label>不熟悉叫车程序</label></li>
        <li style="width:99%;"><input  id="q5_2" type="checkbox" name="q5" value="习惯小汽车出行"><label>习惯小汽车出行</label></li>
        <li style="width: 99%;"><input  id="q5_3" type="checkbox" name="q5" value="周边公共交通服务便利"><label>周边公共交通服务便利</label></li>
        <li style="width: 99%;"><input  id="q5_4" type="checkbox" name="q5" value="出租车服务便利"><label>出租车服务便利</label></li>
        <li style="width: 99%;"><input  id="q5_5" type="checkbox" name="q5" value="担忧网约车安全性"><label>担忧网约车安全性</label></li>
        <li style="width: 99%;"><input  id="q5_5" type="checkbox" name="q5" value="担忧网约车安全性"><label>担忧网约车安全性</label></li>
        <li style="width: 99%;"><button onclick='formSubmit()'>结果</button></li>
    </ul>
</div>

<script>
    function formSubmit() {//点击按钮获取复选框结果
        var obj = document.getElementsByName("q5");
        var check_arr = [];
        for (var i = 0; i < obj.length; i++) {
            if (obj[i].checked)
                check_arr.push(obj[i].value);
        }
        alert("您的选项是："+check_arr);
    }
</script>
```

上述代码示例中当用户点击结果按钮时，使用 JavaScript 获取已选中的复选框的 value 值，并打印结果。

### 2.3 单选框

和复选框比较类似，单选框也是用来作为选项的，不同的是单选框只能选择一个，但是是在一个单选框组内才行，同一个组的 name 值必须相同，定义单选框的方式是设置 type=radio：

```javascript

<div class="div_title_question_all">
<div id="divTitle2" class="div_title_question">您的年龄段：<span class="req">*</span></div>
</div>
<div class="div_table_radio_question" id="divquestion2">
    <ul class="ulradiocheck" style="
list-style-type:none">
        <li style="width: 19%;"><input  type="radio" name="q2" id="q2_1" value="1"><label for="q2_1">18岁以下</label></li>
        <li style="width:19%;"><input  type="radio" name="q2" id="q2_2" value="2"><label for="q2_2">18~25</label></li>
        <li style="width:19%;"><input  type="radio" name="q2" id="q2_3" value="3"><label for="q2_3">26~30</label></li>
        <li style="width:19%;"><input  type="radio" name="q2" id="q2_4" value="4"><label for="q2_4">31~35</label></li>
        <li style="width:19%;"><input  type="radio" name="q2" id="q2_5" value="5"><label for="q2_5">36~40</label></li>
        <li style="width: 19%;"><input  type="radio" name="q2" id="q2_6" value="6"><label for="q2_6">41~45</label></li>
        <li style="width:19%;"><input  type="radio" name="q2" id="q2_7" value="7"><label for="q2_7">46~50</label></li>
        <li style="width:19%;"><input  type="radio" name="q2" id="q2_8" value="8"><label for="q2_8">51~55</label></li>
        <li style="width:19%;"><input  type="radio" name="q2" id="q2_9" value="9"><label for="q2_9">56~60</label></li>
        <li style="width:19%;"><input  type="radio" name="q2" id="q2_10" value="10"><label for="q2_10">60岁以上</label></li>
    </ul>
</div>

```

### 2.4 密码

密码是一种特殊的文本控件，主要用于登录注册时输入的用户密码，默认显示……，设置 type=password 可以定义一个密码控件，例如：

```javascript
<div class="ui attached segment">
    <div class="required inline field ">
        <label for="user_name">用户名或邮箱</label>
        <input id="user_name" name="user_name" value="" autofocus="" required="">
    </div>
    <div class="required inline field ">
        <label for="password">密码</label>
        <input id="password" name="password" type="password" autocomplete="off" value="" required=""><!-- 密码表单 -->
    </div>
    <div class="inline field">
        <label></label>
        <button class="ui green button">登录</button>
    </div>
</div>
```

### 2.5 隐藏输入框

JavaScript 没有全局变量的功能，有时需要定义全局变量，可以用 cookie 来实现，但是浏览器可以关闭 cookie ，而且 cookie 在发送 HTTP 时会被带上，浪费数据传输，这时使用隐藏域的方式实现比较简单。例如：

```javascript
<form action="https://www.baidu.com" method="post" onclick='beforeSubmit()'>
    <input type="hidden" id='_viewstate' name="_viewstate" /><!--隐藏域-->
    <input type="submit" value="click" />
</form>
<script>
function beforeSubmit(){//提交表单前执行赋值函数
    document.getElementById("_viewstate").value = document.referrer;
}
</script>
```

上述代码通过定义隐藏域，当点击提交按钮时将该页面的上个页面的 URL 传输给表单的目标 URL，做到传输值的作用。

### 2.6 文件表单

有时需要做一个上传文件的功能，这时需要用到文件表单，通过设置 `type=file` 可以定义文件表单控件，还需要设置 `enctype=multipart/form-data` 编码方式，才能正确传输文件，例如：

```javascript
<form  action="/user/settings/avatar" method="post" enctype="multipart/form-data">
    <div>
        <label for="avatar">选择新的头像</label>
        <input name="avatar" type="file">
    </div>
</form>
```

### 2.7 重置按钮

重置按钮用于将表单内的控件的值重置为初始化状态，**并非清空数据**，表单中的初始化数据定义在 value 值中，而且重置按钮必须**包裹在 form 表单标签中**，例如：

```javascript
<form  action="/user/settings" method="post">
    <div>
        <label for="username">用户名</label>
        <input id="username" name="name" value=""  autofocus="" required="">*<!--初始化表单是空-->
    </div>
    <div>
        <label for="full_name">自定义名称</label>
        <input id="full_name" name="full_name" value=""><!--初始化表单是空-->
    </div>
    <div>
        <label for="email">邮箱</label>
        <input id="email" name="email" value="" required="">*<!--初始化表单是空-->
    </div>
    <div>
        <label for="website">个人网站</label>
        <input id="website" name="website" type="url" value=""><!--初始化表单是空-->
    </div>
    <div>
        <label for="location">所在地区</label>
        <input id="location" name="location" value=""><!--初始化表单是空-->
    </div>
    <div>
        <input type="reset" value="重置"><button class="ui green button">更新信息</button>
    </div>
</form>
```

### 2.8 提交按钮

提交按钮相当于表单 form 的开关，点击这个开关相当于将表单中的数据提交给服务器。通过设置 type=submit 可以定义一个提交表单按钮，这个按钮必须包裹在 form 标签中才能生效，例如：

```javascript
<form class="ui form" action="/user/settings" method="post">
    <div class="required field ">
        <label for="username">用户名<span class="text red hide" id="name-change-prompt"> 该操作将会影响到所有与您帐户有关的链接</span></label>
        <input id="username" name="name" value="" data-name="jdhg" autofocus="" required=""><!--初始化表单是空-->
    </div>
    <div class="field">
        <input type="submit" value="提交">
    </div>
</form>
```

## 3. 小结

本章介绍了网页中客户与服务器交互数据的方式 - 表单，通过表单用户可以将数据发送到服务器，围绕这个功能浏览器定义了一系列的提交方式、认证方式、数据类型等功能，通过复杂的交互方式网页可以实现真正的互联网属性而不是单纯的视觉展现

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

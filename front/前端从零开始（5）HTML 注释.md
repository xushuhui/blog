# HTML 文件的注释

HTML 注释标签是在 HTML 源代码中添加说明 / 注释，使用该标签注释的内容将不在浏览器中显示。

边写代码边写注释是一个好习惯，特别是代码量很大的时候，如果不写注释，你写完倒回去看可能都不知道自己写的什么了。

如果是团队合作，就更应该写注释了，如果不写每人知道你写的什么，还需要从头到尾仔细看一遍。

团队可以自行约定注释标准及规范，**例如：**

1. 所有的注释写在内容上方；
2. 如果有多句注释，每一句注释直接换行；
3. 如果是方法注释，写明方法作用和每个参数介绍；
4. 注释的语句简单明了；
5. 每个变量和方法标注作用。

## 1. 什么是注释

注释指解释字句的文字，也指用文字解释字句。当我们需要给代码做批注，我们就需要注释代码。

注释是我们提高代码可读性，代码可维护性最简单有效的方式，往往多写几个字，会给我们日后的维护，减少很大的难度。

## 2. 注释的作用

代码注释不仅方便程序员自己回忆起以前代码的用途，还可以帮助其他程序员很快的读懂你的程序的功能，方便多人合作开发网页代码。

## 3. 如何在 HTML 文件中书写注释

在 `HTML` 文件中， 我们将注释写在`<!-- -->`中，注释只是起解释作用，并不会在页面上显示出来。

我们如果使用编辑器编写代码，如 VSCode， WebStorm 等，那么在编辑器书写注释的快捷方式为 `ctrl` + `/` ， 我们只需要先写注释内容， 然后在选中注释内容，按 `ctrl` + `/` 即可。

## 4. 经验分享

1. 注释可以写在 HTML 文件的任意位置。

```javascript
<!DOCTYPE HTML>
<HTML lang="en">
  <head>
    <meta charset="UTF-8">
    <!-- HTML注释可以在head标签中 -->
    <title>HTML注释</title>
  </head>
  <body>
    <!-- HTML注释也可以在body标签中  -->
    <h1>我是一个大标题</h1> <!-- HTML注释也可以在标签后面 -->
  </body>
</HTML>
```

2. 注释内容必须写在`<!-- -->`中，否则不会生效。

```javascript
<!DOCTYPE HTML>
<HTML lang="en">
  <head>
    <meta charset="UTF-8">
    <title>HTML注释</title>
  </head>
  <body>
    <!--  --> HTML注释也可以在body标签中(此段内容没有在<!-- -->中，则不为注释，而是页面上的内容)
    <h1>我是一个大标题</h1> <!-- 只有写在规定位置，才是注释 -->
  </body>
</HTML>
```

3. 注释的内容只是解释，并不会实际的显示在页面上。

```javascript
<!DOCTYPE HTML>
<HTML lang="en">
  <head>
    <meta charset="UTF-8">
    <title>HTML注释</title>
  </head>
  <body>
    <!-- 注释的内容不会显示在页面上-->
    <h1>我是一个大标题</h1>
  </body>
</HTML>
```

4. 编写代码时，尽量把注释写的简单明了，通俗易懂，不仅方便自己维护代码，也方便与他人的协同开发。

```javascript
<!DOCTYPE HTML>
<HTML lang="en">
  <head>
    <meta charset="UTF-8">
    <title>HTML注释</title>
  </head>
  <body>
    <!-- 这是头部 -->
    <header></header>
    <!-- 这是大标题-->
    <h1>我是一个大标题</h1>
  </body>
</HTML>
```

5. 注释标签可以自定义内容，也可以注释已写的 HTML 标签。

```javascript
<!DOCTYPE HTML>
<HTML lang="en">
  <head>
    <meta charset="UTF-8">
    <title>HTML注释</title>
  </head>
  <body>
    <!-- 自定义注释内容 -->
    <p>这是一段话</p>
    <!-- 注释HTML标签， 则该标签也不会在页面上显示 -->
    <!-- <h1>我是一个大标题</h1>  -->
  </body>
</HTML>
```

6. 开发人员一定要养成书写注释的好习惯。

## 5. 真实案例分享

简书官网

```javascript
 <!-- 正常登录登录名输入框 -->
<div >
    <input placeholder="手机号或邮箱" type="text" />
</div>

<!-- 海外登录登录名输入框 -->
<div >
   <input placeholder="密码" type="password" />
</div>
```

## 6. 小结

1. 注释只是起一个批注解释作用，并不会在页面上产生实际的内容。
2. HTML 文件中注释写在 `<!-- -->` 中。

![图片描述](https://img.mukewang.com/wiki/5f62faec09c111c209660552.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

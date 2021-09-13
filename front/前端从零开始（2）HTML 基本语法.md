# HTML 基本语法

HTML 它有属于它自己的一套专属语法。我们如果想要编写一个完整的网页，那么我们就必须遵循 HTML 的语法来编写代码，HTML 都是由各种**标签**构成的，我们只需要记住这些标签的写法和意义，那么我们就可以编写网页的基本结构了。

HTML 文件都由不同的标签构成的：

```javascript
  <!DOCTYPE HTML>
    <HTML lang="en">
    <head>
      <title>HTML基本语法</title>
    </head>
    <body>
      <p>这是一段话</p>
      <h1>我是一个大标题</h1>
      <a href="https://www.baidu.com">百度</a>
    </body>
    </HTML>
```

## 1. HTML 的构成

HTML 文件都是由标签构成的，大部分标签都是**双标签**，由**头标签**和**尾标签**组成，小部分标签为**单标签**，没有尾标签。每个标签有特定的作用，例如，代表“连接”的 `a` 标签，代表“图片”的 `img` 标签等。

## 2. 标签的基本写法

### 2.1 双标签的写法

```javascript
  <p>
  这是一段话
  </p>
```

**总结**: 双标签是成对出现的， 尾标签在标签名前会多一个`/`。

### 2.2 单标签的写法

```javascript
<img src='https://www.imooc.com/static/img/index/logo.png' alt='慕课logo'>
```

**总结**：单标签没有尾标签。

## 3. 标签的内容和属性

### 3.1 标签的内容

标签的内容写在头标签和尾标签之间， 代表这段内容由特定的标签修饰。

```javascript
  <p>
  这是一段话 <!-- 这段为p标签的内容 -->
  </p>
```

> **Tips**：单标签没有内容，因为它没有尾标签，通常我们指的是双标签之间的内容。

### 3.2 标签的属性

标签的属性，如果是标签为双标签，则属性写在头标签中（头标签的<>内）， 如果是单标签，则写在标签的<>内。

```javascript
 <a href='https://www.baidu.com'>百度</a> <!-- 双标签的属性写在头标签的<>内 -->
 <img src='https://www.imooc.com/static/img/index/logo.png' alt='慕课logo'> <!-- 单标签的属性写在标签的<>内 -->
```

**总结**：标签的属性有三部分构成，**属性名**，**等号**，**属性值**，等号左边的为属性名，等号右边的为属性值，属性值必须由引号引起来，单引号和双引号都可以。标签的属性用来给标签添加属性，让标签有特点的作用。

## 4. HTML 标签的关系

### 4.1 嵌套关系

一组标签写在另外一组标签之间，充当了另外一组标签的内容。**例如：**

```javascript
<div>
  <p>
    我是一个p标签
  </p>
</div>
```

标签与标签之间是可以嵌套的，但先后顺序必须保持一致。

### 4.2 并列关系

一组标签和另外一组标签平级，没有任何的嵌套关系。**例如：**

```javascript
<div>
    我是一个div标签
</div>
<p>
    我是一个p标签
</p>
```

> **Tips**：HTML 标签只有两种关系，要么是嵌套关系，要么是并列关系。

## 5. 标签的注意事项

HTML 标签不区分大小写，`<p>` 和 `<P>` 是一样的，但建议小写，因为大部分程序员都以小写为准。

## 6. 小结

1. HTML 文件都由不同的标签构成的。

```javascript
  <!DOCTYPE HTML>
    <HTML lang="en">
    <head>
      <title>HTML基本语法</title>
    </head>
    <body>
      <p>这是一段话</p>
      <h1>我是一个大标题</h1>
      <a href="https://www.baidu.com">百度</a>
    </body>
    </HTML>
```

1. 标签分为双标签和单标签。

```javascript
  <!-- img标签为单标签，没有尾标签和内容 -->
  <img src='https://www.imooc.com/static/img/index/logo.png' alt='慕课logo'>

  <!-- p标签为双标签，有尾标签和内容 -->
  <p>这是一段话</p>
```

1. 双标签都是成对出现的。

```javascript
  <!-- p标签为双标签，有尾标签， 尾标签多一个/ -->
  <p>这是一段话</p>
```

1. 双标签具有内容和属性。

```javascript
  <!-- a标签为双标签， href为标签属性， 百度为标签的内容 -->
  <a href='https://www.baidu.com'>百度</a>
```

1. 单标签只有属性，没有内容。

```javascript
  <!-- img标签为单标签， src和alt都为属性，单标签没有内容 -->
  <img src='https://www.imooc.com/static/img/index/logo.png' alt='慕课logo'>
```

1. 标签有且仅有嵌套和并列关系。

```javascript
<!DOCTYPE HTML>
  <!-- head标签和body标签为并列关系， head和HTML、body和HTML均为嵌套关系 -->
  <HTML lang="en">
    <head>
      <title>HTML基本语法</title>
    </head>
    <body>
    </body>
  </HTML>
```

![图片描述](https://img.mukewang.com/wiki/5f62f85409b7762013250542.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

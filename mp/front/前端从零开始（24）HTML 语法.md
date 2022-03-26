# 语法简介

本章概括介绍 HTML 的标签语法，以及 HTML 语言和传统语言的差异。

## 1. HTML 来历

1969 年，IBM 的技术人员 Charles Goldfarh 和 Edward Mosher 等人一起发明了通用标记语言 GML（Generalized Marked Language）。1985 年在英国成立了国际 SGML 用户组织，在 1986 年，SGML 成为国际标准 ISO8879：信息处理标准通用标记语言（Information processing Text and office system Standard generalized markup language）。HTML 和 XML 派生于 SGML，XML 可以被认为是它的一个子集，而 HTML 是它的一个应用。为了告诉浏览器我们需要展示什么内容，HTML 定义了一整套符号标记规范，这些规范包括

* 设置文字的格式；
* 创建列表；
* 显示图片；
* 显示多媒体；
* 显示超链接；
* 等等。

## 2. 标准

HTML3.2 以前的标准是由 IETF 制定。IETF 互联网工程任务组（The Internet Engineering Task Force），成立于 1985 年底，是全球互联网最具权威的技术标准化组织，主要任务是负责互联网相关技术规范的研发和制定，当前绝大多数国际互联网技术标准出自 IETF。 IETF 的工作成果主要以 RFC 文档的途径发布。

HTML 3.2 开始，由 W3C 制定标准。W3C 万维网联盟（ The World Wide Web Consortium ）创建于 1994 年，是 Web 技术领域最具权威和影响力的国际中立性技术标准机构。到目前为止，W3C 已发布了 200 多项影响深远的 Web 技术标准及实施指南，如广为业界采用的超文本标记语言（ HTML ）、可扩展标记语言（ XML ）以及帮助残障人士有效获得 Web 内容的信息无障碍指南（ WCAG ）等，有效促进了 Web 技术的互相兼容，对互联网技术的发展和应用起到了基础性和根本性的支撑作用。

## 3.HTML 和编译型语言的区别

计算机语言分成**解释型语言**和**编译型语言**两种。

我们下面来展开讲一下**解释型语言**和**编译型语言**两者的区别：

在说两区别之前我们先来讲一下计算机怎么把代码翻译成计算机能看得懂的语言（翻译成机器码）。

众所周知，计算机 CPU 的集成电路中，除了电容、电阻、电感就是晶体管了，每个晶体管相当于一个开关，理论上 CPU 只能存储识别两个标识符，那就是 0 和 1，所以说 CPU 识别的指令集只能由 0 和 1 组合。那么所有的计算机语言想要 CPU 能看得懂，必须翻译成 0/1 代码才行，这个由 0/1 组成的代码叫做机器码。但是机器码相对于人来说过于繁琐，所以就有人发明了**高级语言**、**低级语言**等等，这些语言的分级是根据它的语法是贴近人还是贴近机器来区分的，越贴近人它就越高级，越贴近机器它就越低级，但是最终想要 CPU 可以识别都需要翻译成机器码。

典型的**低级语言**包括刚刚提到的机器码、汇编语言、c 等，**高级语言**包括 PHP、c#、JavaScript、Java、Python 等等。

什么是编译型语言和解释性语言呢？

刚刚我们提到翻译成机器码，这个翻译的过程就叫做编译或解释。编译型语言是指通过编译器翻译成完整的机器码，然后通过 CPU 去执行。

而解释型语言是指通过一个虚拟机的方式一行行的翻译，翻译一行执行一行；还有一种方式是混合型，介于两者之间。常见的编译型语言包括 c++、c、rust 等，解释型语言包括 JavaScript、PHP、HTML 等等，混合型包括 Python、Java 等。

## 4. 标签语法和属性

### 4.1 HTML 文档扩展名

HTML 以文档的形式存储，文档的后缀可以是 .html .htm .xhtml，有时也会看到 php/asp/jsp 等类型的网页后缀，这种是通过服务器的 CGI 动态解析过的网页，网页内容也是 HTML 格式，只不过网页后缀是根据服务器的 CGI 网关的不同来定义的。不同的后缀形式可能在浏览器的解释结果不相同，在此不做深入讨论。

### 4.2 HTML 标签

标签有两种定义方式：

* 闭合型标签：《标签》内容</ 标签》
* 自闭合标签： 《标签 />或者《标签》

标签不区分大小写，工作中通常使用小写，因为日常写代码的时候 IDE 通常设置小写。

```javascript
<div>这是一个闭合标签</div> <!-- 闭合标签 -->
<DIV>跟上边等价</DIV> <!-- 标签不区分大小写 -->
<input type='text' name='test' /> <!-- 自闭合标签 -->
```

自闭合标签通常是一些不需要包含具体 HTML 内容的的标签，例如表单、图片、换行符、css 样式等等。

```javascript
<input type='radio' /><!-- 表单 -->
<img src='https://www.baidu.com/img/bd_logo1.png' /><!-- 图片 -->
<br>
<link rel="stylesheet" type="text/css" href="theme.css" /><!-- 引入css样式 -->
```

从上述例子看得出来，通常自闭和标签是不包含实际的文本内容，通常是用来引入文件、图片、样式等作用，而闭合标签是包含内容，既然包含内容所以后边的闭合标签相当于起到一个结束符的作用。实际开发中，如果忘记将标签闭合，通常来说日常的浏览器（例如在 chrome v81） 中不会报语法错误，而是自动闭合，但是不建议这么做，因为浏览器并非完全的智能，有可能会引起排版错误，例如：

```javascript
<p style='color:red'>我是红色的段落</p><!-- 正确的写法 -->
<span>我是默认的文本</span>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef022d80946561302100102.jpg)

这段代码中，定义了一个段落，并设置了字体为红色，当忘记使用闭合标签时

```javascript
<p style='color:red'>我是红色的段落<!-- 忘记闭合 -->
<span>我是默认的文本</span>
```

在浏览器的展示效果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef022f70960b41f03180052.jpg)

可以看到，样式出现了错乱。

> **TIPS:** 所以在日常项目开发中，尽量选择带有语法检测的 IDE，可以避免大部分的语法错误。

### 4.3 属性

属性是控制元素的第二个维度，通过属性的设置，可以让元素展现出不同的表现风格。属性包含属性名称和属性值，通常在元素标签中用 name = value 的方式设置，多个属性之间使用空格隔开。例如表单中的是否只读，可以用

```javascript
<input readonly=true type='text'>
```

属性值通常需要使用引号隔开，并非强制性，但是一旦属性值中包含空格的话不使用引号隔开则会解析异常，例如：

```javascript
<input name="one style" type='text'><!-- 定义表单的name为 one style -->
<input name=one style type="text"><!-- 未使用引号 -->
```

上述表单定义 name 为 one style，第二种写法的解析结果则是定义表单的 name 为 one，然后定义一个空的 style 属性，显然与预期不符，但是浏览器并不会报错。所以使用规范的写法（引号隔开）则会避免这种异常情况出现。

一个标签内可以包含多个属性，属性名称不能重复，属性名称不区分大小写，例如：

```javascript
<input type='text' type='file' ><!-- 这种写法是错误的 -->
```

这样的话浏览器会解析错误，不同的浏览器会呈现不同的效果，但是一般不会报错。

除了可以在标签内定义元素的属性之外，还可以使用 JavaScript 动态控制属性，这种方式在项目开发中经常使用，例如：

```javascript
<input type='text' id = 'num1'> + <input type='text' id='num2'> = <input type='text' id='res'>
<button onclick='count()'>计算结果</button><!-- 点击按钮实现计算 -->
<script>
function count(){
    if(isNaN(document.getElementById('num1').value) || isNaN(document.getElementById('num2').value)) return alert("请输入正确的数字");//判断输入数字是否合法
    document.getElementById('res').value = parseFloat(document.getElementById("num1").value) + parseFloat(document.getElementById("num2").value);//将结果输出
}
</script>
```

上述代码，通过使用 JavaScript 设置表单元素的 value 属性，实现了一个简单的加法计算器。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef0231d09233f8608090048.jpg)

### 4.4 注释

注释内容在浏览器中不做解释，只用于开发者记录代码用途等信息，格式如下：

```javascript
< !-- 这是在HTML的注释内容，支持多行注释 -->
```

```javascript

// 这是JavaScript的单行注释
/*
这是JavaScript的多行注释
*/

```

```javascript
/* 这是css的注释，支持单行注释和多行注释 */
```

以上可以看出，css 和 HTML 都只有一种注释方式，这种设计方式可以因为 HTML 和 css 只是标记语言，没有函数和类等传统编程语言的概念，因此无需过多的注释风格。

> **TIPS:** 除以上注释，实际项目开发中可能还会针对文件、类、项目模块做注释，这种注释的设置方式跟使用的 IDE 设置有关，有些 IDE 可以设置当声明一个类或文件时自动加上定义好的注释内容，良好的注释可以提升项目整体的扩展性、维护性、可用性。在程序编译或者解释的过程中会被浏览器忽略，在开发大型项目中上线之前通过会做程序压缩合并等处理，处理的过程中通常会删掉注释代码

### 4.5 HTML 代码结构

符合 HTML 标准的网页代码结构大致格式如下

```javascript
<!DOCTYPE >  <!-- 控制w3c格式 -->
<html>
<head>
<meta charset="utf-8"> <!-- 告诉浏览器使用的是utf-8字符集 --> <title>我的网站</title><!-- 这里用来编写网站的标题 -->
</head>
<body> <!-- 只有body标签的内容，才能真正显示在浏览器的窗口 -->
</body>
</html>
```

上述内容是一个大致符合 w3c 标准的 HTML 代码框架，其中包含文档描述头标签、HTML 标签、头标签和 body 标签，但是实际开发者可能不会包含这么完整的标签框架，例如可以只声明 body 内的标签：

```javascript
<a href="http://www.baidu.com">百度搜索</a>
```

这样的话，浏览器的展示效果并未有什么变化，但是当打开浏览器调试工具可以发现：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef0233e098afcd505270192.jpg)

浏览器帮我们补齐的缺失的标签。

### 4.6 在浏览器中查看源码

开发过程中需要调试、查看代码，在浏览器中可以通过开发者工具方便的查看源码，这里以 Chrome 举例，在网页中点击右键 ->查看网页源代码 / 检查，即可查看源码。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef0235709669ff118020855.jpg)

# 总结

本章主要从标签语法、属性用途、文件格式、框架结构等几个方面大致介绍了 HTML 的规范和标准，只是预热一下，没有作深入讨论。之后的章节里，我们会根据以上内容延伸讲解。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

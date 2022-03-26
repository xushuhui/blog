# 什么是 JavaScript

> JavaScript ( JS ) 是一种具有函数优先的轻量级，解释型或即时编译型的编程语言。(MDN)

`JavaScript` 是一种编程语言，主要参与构建 Web 前端应用。

## 1. JavasScript 的由来

早期的浏览器是不具备与用户交互能力的，可以想象一下，在那个电话拨号上网的年代，带宽只有 56Kbps，也就是意味着标准最高下载速率只有 8KB/s。

在这个速度下，表单提交就是噩梦般的存在。

例如报名考试，就需要填写非常多内容，当用户花了十几分钟找各种资料填写表单后，点击提交就要等待十几秒甚至更多。

假如这时表单信息出错，如身份证没核对仔细少了一位，整个页面就会做刷新操作，表单需要重新填写。

这是一个比较典型的场景。当时最火的浏览器 `Navigator` 的开发公司 `netscape` 就因这些问题，急需一个浏览器使用的脚本语言，让运行在浏览器上的网页可以做一些交互。

`netscape` 因为有这个需求，招聘了 `Brendan Eich` ， `Brendan Eich` 进公司后就开始研究使用 `Scheme` 语言作为在网页中使用的脚本语言的可能性。

但是由于当时 Java 爆火，`netscape` 又在于开发了 Java 的 `Sun` 公司合作，就想让这个脚本语言要足够像 Java，但是又要比 Java 简单。

然后 `Brendan Eich` 就被指定开发这个“简易 Java”。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e78e45208dcdefc11530649.jpg)

Brendan Eich

一段时间之后 `JavaScript` 也就诞生了。

后续 `netscape` 将 `JavaScript` 交给了 `ECMA` 组织进行标准化，编号为 262，也就是说现在的 `JavaScript` 实际上是 `ECMA-262` 标准的实现。

## 2. 与 Java 的区别

通过`JavaScript的由来`可以知道，两个语言本质上没有太大的关系，仅仅只为了让他们像，才让 `JavaScript` 的名字中有了 `Java` ，才让他的内部的一些设计机制像 `Java`。

事实上 `JavaScript` 上在设计上还融合了`C语言`的语法，`Self语言`的原型设计等。

## 3. JavaScript 的主要应用

> 以下列举的各个场景不仅仅是需要掌握 `JavaScript` ，还需要很多知识点与技术栈来共同协作完成，但是 `JavaScript` 是必不可少的技术栈。

### 3.1 网页开发

网页开发的基本三大件为 HTML、CSS、JavaScript，如果将 HTML 比作骨架，CSS 比作皮肤，那 JavaScript 就是可以让骨架动起来，改变皮肤性状的存在。

现代的前端应用离不开 `JavaScript` ，随着浏览器的性能越来越好，产品交互越来越复杂，`JavaScript` 的地位也越来越高。

表单验证、动画效果甚至 3D 应用，均可以由 JavaScript 来完成。

![](https://xushuhui.gitee.io/image/imooc/5e78e66e0a1e6f1203360193.jpg)使用 WebGL 制作的 3D 应用，可以直接运行在现代浏览器

### 3.2 服务端应用开发

2009 年发布 `Node.js` 的发布，意味着前端程序员可以用较低的成本跨入服务端开发。

`Node.js` 提供了开发服务端所需要的特性，如 HTTP 服务、本地文件读写操作等。

开发者可以使用 `JavaScript` 语言开发 `Node.js` 应用。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e78e6a209e5fdb212240618.jpg)

Node.js

### 3.3 桌面应用开发

`Electron` 是由 `Github` 开发的，可以使用 HTML、CSS、JavaScript 来构建桌面应用的开源库。

使用`Electron`就可以让前端开发者进行桌面端应用的开发。

`Visual Studio Code`、`Atom`、`Skype` 等应用都是使用 `Electron` 开发的。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e78e6ea09b722da08000420.jpg)

Electron

### 3.4 移动端应用开发

移动端应用也可以使用 `JavaScript` 进行开发，如 `React Native` 或者 `Weex` 等框架。

![图片描述](https://xushuhui.gitee.io/image/imooc/5e78e715093f448912100590.jpg)

Weex 框架

## 4. 适合群体

本篇 Wiki 主要为 `ECMAScript262` 第五版内容，适合初学者学习或者进行知识点查阅。

## 5. 前置知识

章节中的例子可能会涉及部分 `HTML` 与 `CSS` 的知识点，所以需要了解或者掌握一些 `HTML`、`CSS` 相关的内容作为前置知识。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

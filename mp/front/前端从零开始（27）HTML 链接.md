# HTML 链接

本章节介绍 HTML 页面与页面，文档与文档之间的媒介 - 链接，链接为客户和服务器提供了主要的交互的手段。这是一个比较常见的标签类型，几乎在所有的网站中都能看到它的身影。

## 1. 样式

使用标签`<a>`设置超文本链接，链接可以是图形或者文字，当鼠标移动到链接上时，会出现一个小手的形状。

* 未访问的链接显示蓝色，带有下划线；
* 访问过的链接显示紫色，带有下划线；
* 点击时，链接显示红色，带有下划线。

以上是浏览器的链接的默认的样式，在实际的项目开发中，通常网页的风格需要重新设计，那么链接的默认样式也会随网页的风格而改变，要改变链接的样式需要用到 css 伪类，伪类是 css 的一种定义方式，可以和 css 的类搭配使用，以下是几种伪类的使用实例：

```javascript
<style>
.divcss:link{ color:#F00}/* 链接默认是红色 */
.divcss:hover{ color:#000}/* 鼠标悬停变黑色 */
.divcss:active{ color:#03F}/* 鼠标点击时蓝色 */
.divcss:visited{ color:#F0F}/* 访问过为粉红 */
</style>
<a class="divcss" href="http://www.baidu.com">百度</a>
```

从上面代码可以看到，伪类包括链接默认（:link）、鼠标悬停（:hover）、鼠标点击时（:active）、链接访问后（:visited）这几种样式定义方式，这些都是专门针对于链接。

## 2. 语法

`a` 标签的语法如下：

```javascript
<a href="url">内容</a>
```

链接 a 是一种闭合标签，一个最基础的链接定义包括链接标签 a 标签、标签内容、链接地址 href 属性，其中 href 是链接中最重要的一个属性，如果未定义 href 浏览器也不会报错，但是这就失去了标签的意义，变得跟普通文本标签没有区别了。

### 2.1 target 属性

由于链接是 HTML 中重要的交互介质，当用户点击一个链接跳转的目标界面并非是当前的界面，这时候就需要一个重要的属性 target 来定义所要跳转的目标界面。

target 属性的可选值有以下几个：

* **_blank：** 在新窗口打开链接；
* **_self ：** 默认方式，在当前窗口载入链接内容；
* **_top：** 在包含当前文档的最顶层的窗口载入链接内容。（一般用在有 frame 框架标签的页面中。）
* **_parent：** 在当前文档的父窗口载入链接内容。（一般用在有 frame 框架标签的页面中。）

其中`_top` 和 `_parent` 不太好理解，看下面一个例子：

```javascript
<iframe name="baidu"></iframe><!-- 定义一个iframe -->
<a href="http://www.baidu.com" target="baidu">搜一搜</a><!-- 定义一个链接 -->
```

以上代码实现：点击链接（搜一搜）后，在当前页面的 iframe 中嵌入搜索框页面。

> **Tips：** 该功能在 IE10 版本以下不兼容。

其中顶层窗口和父窗口针对在网页中嵌套 `iframe` 或者 `frameset` 有效，当嵌套框架时被嵌套框架是嵌套框架的 _parent，最外层的 HTML 称为 _top。

新的 HTML 标准中关于 target 属性的存在有一定的争议，主要是因为可以使用 JavaScript 的方式替代 target，例如：

```javascript
<a href="javascript:window.open('https://www.baidu.com')" target="_blank">点击在新窗口打开</a>
<a href="javascript:location.href='https://www.baidu.com'">点击在当前窗口打开</a>
<a href="javascript:top.location.href='https://www.baidu.com'">点击在顶层窗口打开</a>
<a href="javascript:parent.location.href='https://www.baidu.com'">点击在父窗口打开</a>
```

以上代码使用 JavaScript 函数的方式实现各种打开链接的方式。

### 2.1 id 属性

id 属性是 html 的通用属性，主要是用于给元素设置事件或者设置样式时用到，以下代码实现点击链接跳转之前弹框提示：

```javascript
<a href='https://www.baidu.com' id='test'>百度</a> <!-- 定义一个链接 -->
<script>
document.getElementById("test").onclick = function(){  //点击链接跳转前弹框提示
    alert("即将跳转到百度");
}
</script>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef04d38092720ad05610113.jpg)

### 2.2 charset 属性

设置编码，在 HTML5 不支持：

```javascript
<a charset="gb2312" href="http://www.baidu.com">百度</a>
```

此属性在主流浏览器中几乎不支持，所以定义了没有什么作用。

### 2.3 coords 属性

定义链接的坐标，HTML5 不支持。

### 2.4 download 属性

有时用户点击一个链接的时候，这个链接的内容有可能包含 word、pdf、png、jpg 或者 MP4，zip 等内容，用户需要的仅仅是下载而不是跳转链接，那么这时候我们只需要定义一个 download 属性，这个属性包含下载时的文件名称。

```javascript
<a href="https://img01.sogoucdn.com/app/a/07/57d9e1c69332c8da692c3ec5c0e1466b" download="图片">点击下载图片</a>
```

以上代码定义了一个下载图片链接，download 属性不支持 IE 浏览器

### 2.5 href 属性

定义超链接的跳转目标，可以是：

* 绝对 URL： 例如 `http://www.baidu.com`；
* 相对 URL： 例如 `/index.html`；
* 锚点 ： dom 的 id；
* JavaScript 表达式：例如 `javascript:void(0)` 阻止链接跳转。

```javascript
<a href="https://www.baidu.com" id="test">百度</a>
<a href="/index.html">主页</a>
<a href="javascript:void(0)">普通按钮</a>
<!-- 请注意：向下滚动页面，可看到下面这个向上跳转的链接。单击它，会让页面跳到网页的顶部-->
<a href="#test" style="margin-top:2000px;display:block">跳转到首部</a>
```

以上代码定义了几种链接的方式，其中锚点主要应用于当页面滚动条比较长时，用户可以点击跳转到首部。

### 2.6 hreflang 属性

定义被链接的文档的语言：

```javascript
<a href="http://www.baidu.com" hreflang="zh">百度</a>
```

主流的浏览器暂不支持该属性。

### 2.7 name 属性

通用属性，HTML5 不支持。

### 2.8 rel 属性

定义当前文档，主要用于搜索引擎搜集网页内链接的集合以及链接与链接之间的关系，没有实际功能用途。

`rel` 是英文单词 `relationship` 的缩写，意味着虽然这个属性本身对网页视觉效果而言没有什么本质性的用途，但是它是串联网页与搜索引擎的桥梁，搜索引擎通过这个属性可以获取到链接的一些简介，从而了解到该网页的用途，进而对网页进行归类，方便用户搜索。

### 2.9 rev 属性

定义链接文档与当前文档的关系，HTML5 不支持。

### 2.10 shape 属性

定义链接的形状，HTML5 不支持。

### 2.11 type 属性

定义链接文档的 `mime` 类型，`mime` 是指描述内容类型的因特网标准，以下罗列最新的标准项：

|扩展名|类型 / 子类型|
|------|-----------|
|       |application/octet-stream               |
|323    |text/h323                              |
|acx    |application/internet-property-stream   |
|ai     |application/postscript                 |
|aif    |audio/x-aiff                           |
|aifc   |audio/x-aiff                           |
|aiff   |audio/x-aiff                           |
|asf    |video/x-ms-asf                         |
|asr    |video/x-ms-asf                         |
|asx    |video/x-ms-asf                         |
|au     |audio/basic                            |
|avi    |video/x-msvideo                        |
|axs    |application/olescript                  |
|bas    |text/plain                             |
|bcpio  |application/x-bcpio                    |
|bin    |application/octet-stream               |
|bmp    |image/bmp                              |
|c      |text/plain                             |
|cat    |application/vnd.ms-pkiseccat           |
|cdf    |application/x-cdf                      |
|cer    |application/x-x509-ca-cert             |
|class  |application/octet-stream               |
|clp    |application/x-msclip                   |
|cmx    |image/x-cmx                            |
|cod    |image/cis-cod                          |
|cpio   |application/x-cpio                     |
|crd    |application/x-mscardfile               |
|crl    |application/pkix-crl                   |
|crt    |application/x-x509-ca-cert             |
|csh    |application/x-csh                      |
|css    |text/css                               |
|dcr    |application/x-director                 |
|der    |application/x-x509-ca-cert             |
|dir    |application/x-director                 |
|dll    |application/x-msdownload               |
|dms    |application/octet-stream               |
|doc    |application/msword                     |
|dot    |application/msword                     |
|dvi    |application/x-dvi                      |
|dxr    |application/x-director                 |
|eps    |application/postscript                 |
|etx    |text/x-setext                          |
|evy    |application/envoy                      |
|exe    |application/octet-stream               |
|fif    |application/fractals                   |
|flr    |x-world/x-vrml                         |
|gif    |image/gif                              |
|gtar   |application/x-gtar                     |
|gz     |application/x-gzip                     |
|h      |text/plain                             |
|hdf    |application/x-hdf                      |
|hlp    |application/winhlp                     |
|hqx    |application/mac-binhex40               |
|hta    |application/hta                        |
|htc    |text/x-component                       |
|htm    |text/html                              |
|html   |text/html                              |
|htt    |text/webviewhtml                       |
|ico    |image/x-icon                           |
|ief    |image/ief                              |
|iii    |application/x-iphone                   |
|ins    |application/x-internet-signup          |
|isp    |application/x-internet-signup          |
|jfif   |image/pipeg                            |
|jpe    |image/jpeg                             |
|jpeg   |image/jpeg                             |
|jpg    |image/jpeg                             |
|js     |application/x-javascript               |
|latex  |application/x-latex                    |
|lha    |application/octet-stream               |
|lsf    |video/x-la-asf                         |
|lsx    |video/x-la-asf                         |
|lzh    |application/octet-stream               |
|m13    |application/x-msmediaview              |
|m14    |application/x-msmediaview              |
|m3u    |audio/x-mpegurl                        |
|man    |application/x-troff-man                |
|mdb    |application/x-msaccess                 |
|me     |application/x-troff-me                 |
|mht    |message/rfc822                         |
|mhtml  |message/rfc822                         |
|mid    |audio/mid                              |
|mny    |application/x-msmoney                  |
|mov    |video/quicktime                        |
|movie  |video/x-sgi-movie                      |
|mp2    |video/mpeg                             |
|mp3    |audio/mpeg                             |
|mpa    |video/mpeg                             |
|mpe    |video/mpeg                             |
|mpeg   |video/mpeg                             |
|mpg    |video/mpeg                             |
|mpp    |application/vnd.ms-project             |
|mpv2   |video/mpeg                             |
|ms     |application/x-troff-ms                 |
|mvb    |application/x-msmediaview              |
|nws    |message/rfc822                         |
|oda    |application/oda                        |
|p10    |application/pkcs10                     |
|p12    |application/x-pkcs12                   |
|p7b    |application/x-pkcs7-certificates       |
|p7c    |application/x-pkcs7-mime               |
|p7m    |application/x-pkcs7-mime               |
|p7r    |application/x-pkcs7-certreqresp        |
|p7s    |application/x-pkcs7-signature          |
|pbm    |image/x-portable-bitmap                |
|pdf    |application/pdf                        |
|pfx    |application/x-pkcs12                   |
|pgm    |image/x-portable-graymap               |
|pko    |application/ynd.ms-pkipko              |
|pma    |application/x-perfmon                  |
|pmc    |application/x-perfmon                  |
|pml    |application/x-perfmon                  |
|pmr    |application/x-perfmon                  |
|pmw    |application/x-perfmon                  |
|pnm    |image/x-portable-anymap                |
|pot,   |application/vnd.ms-powerpoint          |
|ppm    |image/x-portable-pixmap                |
|pps    |application/vnd.ms-powerpoint          |
|ppt    |application/vnd.ms-powerpoint          |
|prf    |application/pics-rules                 |
|ps     |application/postscript                 |
|pub    |application/x-mspublisher              |
|qt     |video/quicktime                        |
|ra     |audio/x-pn-realaudio                   |
|ram    |audio/x-pn-realaudio                   |
|ras    |image/x-cmu-raster                     |
|rgb    |image/x-rgb                            |
|rmi    |audio/mid                              |
|roff   |application/x-troff                    |
|rtf    |application/rtf                        |
|rtx    |text/richtext                          |
|scd    |application/x-msschedule               |
|sct    |text/scriptlet                         |
|setpay |application/set-payment-initiation     |
|setreg |application/set-registration-initiation|
|sh     |application/x-sh                       |
|shar   |application/x-shar                     |
|sit    |application/x-stuffit                  |
|snd    |audio/basic                            |
|spc    |application/x-pkcs7-certificates       |
|spl    |application/futuresplash               |
|src    |application/x-wais-source              |
|sst    |application/vnd.ms-pkicertstore        |
|stl    |application/vnd.ms-pkistl              |
|stm    |text/html                              |
|svg    |image/svg+xml                          |
|sv4cpio|application/x-sv4cpio                  |
|sv4crc |application/x-sv4crc                   |
|swf    |application/x-shockwave-flash          |
|t      |application/x-troff                    |
|tar    |application/x-tar                      |
|tcl    |application/x-tcl                      |
|tex    |application/x-tex                      |
|texi   |application/x-texinfo                  |
|texinfo|application/x-texinfo                  |
|tgz    |application/x-compressed               |
|tif    |image/tiff                             |
|tiff   |image/tiff                             |
|tr     |application/x-troff                    |
|trm    |application/x-msterminal               |
|tsv    |text/tab-separated-values              |
|txt    |text/plain                             |
|uls    |text/iuls                              |
|ustar  |application/x-ustar                    |
|vcf    |text/x-vcard                           |
|vrml   |x-world/x-vrml                         |
|wav    |audio/x-wav                            |
|wcm    |application/vnd.ms-works               |
|wdb    |application/vnd.ms-works               |
|wks    |application/vnd.ms-works               |
|wmf    |application/x-msmetafile               |
|wps    |application/vnd.ms-works               |
|wri    |application/x-mswrite                  |
|wrl    |x-world/x-vrml                         |
|wrz    |x-world/x-vrml                         |
|xaf    |x-world/x-vrml                         |
|xbm    |image/x-xbitmap                        |
|xla    |application/vnd.ms-excel               |
|xlc    |application/vnd.ms-excel               |
|xlm    |application/vnd.ms-excel               |
|xls    |application/vnd.ms-excel               |
|xlt    |application/vnd.ms-excel               |
|xlw    |application/vnd.ms-excel               |
|xof    |x-world/x-vrml                         |
|xpm    |image/x-xpixmap                        |
|xwd    |image/x-xwindowdump                    |
|z      |application/x-compress                 |
|zip    |application/zip                        |

## 3. 其他链接因特网标准

除了通过`<a>`链接标签的方式，还可以通过下面三种方式：

* JavaScript
* iframe 标签
* script 标签

等方式与服务器交互。

## 3. 小结

本章节介绍了 `<a>` 标签的用途，以及各个属性的功能，其中有些已经在 HTML5 废弃了，例如，控制样式的属性。`<a>` 标签是 HTML 中用途非常广泛的标签，需要掌握。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

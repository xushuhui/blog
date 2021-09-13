# HTML 图像

在网页中经常能看到丰富多彩的图片，大家肯定很好奇这些图片是如何放到网页中的呢，其实很简单，学习了本章介绍的图片标签之后就了解了。

## 1. img 图像

img 是 英文单词 image 的缩写，意思是图像，它是 HTML 中最通用的定义图片的方式，是一个闭合标签，定义方式如下 ：

```javascript
<img src = "url" alt = "" /> <!-- 定义一个图片 -->
```

其中 src 属性是必须属性，用来指定图片所对应的 URL 地址。

### 1.1 src 属性

src 的全称是 source，表示图片的 URL 地址源。源就是表示图片的地址路径，这个路径可以是相对路径，也可以是绝对路径。绝对路径指的是一个包含网络协议头的完整路径，常用的网络协议是 HTTP 协议，例如 `https://www.baidu.com/img/bd_logo1.png；`相对路径是指这个图片文件跟当前的网页在同一个服务域，例如：`/img/bd_logo1.png`。

```javascript
<img src = "../_nuxt/img/be6bf0d.png" /> <!-- 这是一个相对路径 -->
<img src = "https://www.baidu.com/img/bd_logo1.png" /><!-- 这是一个绝对路径 -->
```

有时当图片过大时或者过多时，网页加载可能会比较慢，这时需要针对图片使用**懒加载**的方式（图片懒加载），懒加载的原理就是将图片的 src 先设置为空，网页其他内容加载完之后，再通过 JavaScript 将 src 属性赋值，例如：

```javascript
<div class="container">
  <img src="img/loading.png" alt="1" data-src="photo-1.jpeg">
  <img src="img/loading.png" alt="2" data-src="photo-2.jpeg">
  <img src="img/loading.png" alt="3" data-src="photo-3.jpeg">
</div>
<script>
window.onload = function (){//页面加载完之后再加载图片
    var a = document.getElementsByTagName("img");//获取图片DOM
    for(var i in a){
        a[i].src = a[i].getAttribute("data-src")
    }
}
</script>
```

点击下面的“运行案例”，可以试试真实的运行效果：

```javascript
<div class="container">
  <img src="https://wiki-code.oss-cn-beijing.aliyuncs.com/html5/img/loading.png" alt="1" data-src="https://wiki-code.oss-cn-beijing.aliyuncs.com/html5/img/photo-3.png">
  <img src="https://wiki-code.oss-cn-beijing.aliyuncs.com/html5/img/loading.png" alt="2" data-src="https://wiki-code.oss-cn-beijing.aliyuncs.com/html5/img/photo-3.png">
  <img src="https://wiki-code.oss-cn-beijing.aliyuncs.com/html5/img/loading.png" alt="3" data-src="https://wiki-code.oss-cn-beijing.aliyuncs.com/html5/img/photo-3.png">
</div>
<script>
window.onload = function (){//页面加载完之后再加载图片
    var a = document.getElementsByTagName("img");//获取图片DOM
    for(var i in a){
        a[i].src = a[i].getAttribute("data-src")
    }
}
</script>
```

以上代码通过定义 onload 事件，将图片延后加载。

#### 1.1.1 Base64 方式加载图片

我们都知道 exe、jpg、pdf 这些格式的文件是使用二进制方式保存的，其中包含很多无法显示和打印的字符，如果要让文本处理软件能处理二进制数据，需要将其编码成一种明文显示的安全格式，Base64 是一种最常见的二进制编码方法。有时为了方便处理，图片并非使用二进制流方式保存，而是使用 Base64 方式编码之后保存在数据库，img 标签的 src 属性可以识别 Base64 格式的编码图片格式，例如：

```javascript
<img src='data:img/jpg;base64,iVBORw0KGgoAAAANSUhEUgAAAhwAAAECCAMAAACCFP44AAAACXBIWXMAAAsTAAALEwEAmpwYAAAK
TWlDQ1BQaG90b3Nob3AgSUNDIHByb2ZpbGUAAHjanVN3WJP3Fj7f92UPVkLY8LGXbIEAIiOsCMgQ
WaIQkgBhhBASQMWFiApWFBURnEhVxILVCkidiOKgKLhnQYqIWotVXDjuH9yntX167+3t+9f7vOec
5/zOec8PgBESJpHmomoAOVKFPDrYH49PSMTJvYACFUjgBCAQ5svCZwXFAADwA3l4fnSwP/wBr28A
AgBw1S4kEsfh/4O6UCZXACCRAOAiEucLAZBSAMguVMgUAMgYALBTs2QKAJQAAGx5fEIiAKoNAOz0
ST4FANipk9wXANiiHKkIAI0BAJkoRyQCQLsAYFWBUiwCwMIAoKxAIi4EwK4BgFm2MkcCgL0FAHaO
WJAPQGAAgJlCLMwAIDgCAEMeE80DIEwDoDDSv+CpX3CFuEgBAMDLlc2XS9IzFLiV0Bp38vDg4iHi
wmyxQmEXKRBmCeQinJebIxNI5wNMzgwAABr50cH+OD+Q5+bk4eZm52zv9MWi/mvwbyI+IfHf/ryM
AgQAEE7P79pf5eXWA3DHAbB1v2upWwDaVgBo3/ldM9sJoFoK0Hr5i3k4/EAenqFQyDwdHAoLC+0l
YqG9MOOLPv8z4W/gi372/EAe/tt68ABxmkCZrcCjg/1xYW52rlKO58sEQjFu9+cj/seFf/2OKdHi
NLFcLBWK8ViJuFAiTcd5uVKRRCHJleIS6X8y8R+W/QmTdw0ArIZPwE62B7XLbMB+7gECiw5Y0nYA
QH7zLYwaC5EAEGc0Mnn3AACTv/mPQCsBAM2XpOMAALzoGFyolBdMxggAAESggSqwQQcMwRSswA6c
wR28wBcCYQZEQAwkwDwQQgbkgBwKoRiWQRlUwDrYBLWwAxqgEZrhELTBMTgN5+ASXIHrcBcGYBie
whi8hgkEQcgIE2EhOogRYo7YIs4IF5mOBCJhSDSSgKQg6YgUUSLFyHKkAqlCapFdSCPyLXIUOY1c
QPqQ28ggMor8irxHMZSBslED1AJ1QLmoHxqKxqBz0XQ0D12AlqJr0Rq0Hj2AtqKn0UvodXQAfYqO
Y4DRMQ5mjNlhXIyHRWCJWBomxxZj5Vg1Vo81Yx1YN3YVG8CeYe8IJAKLgBPsCF6EEMJsgpCQR1hM
WEOoJewjtBK6CFcJg4Qxwicik6hPtCV6EvnEeGI6sZBYRqwm7iEeIZ4lXicOE1+TSCQOyZLkTgoh
JZAySQtJa0jbSC2kU6Q+0hBpnEwm65Btyd7kCLKArCCXkbeQD5BPkvvJw+S3FDrFiOJMCaIkUqSU
Eko1ZT/lBKWfMkKZoKpRzame1AiqiDqfWkltoHZQL1OHqRM0dZolzZsWQ8ukLaPV0JppZ2n3aC/p
dLoJ3YMeRZfQl9Jr6Afp5+mD9HcMDYYNg8dIYigZaxl7GacYtxkvmUymBdOXmchUMNcyG5lnmA+Y
b1VYKvYqfBWRyhKVOpVWlX6V56pUVXNVP9V5qgtUq1UPq15WfaZGVbNQ46kJ1Bar1akdVbupNq7O
UndSj1DPUV+jvl/9gvpjDbKGhUaghkijVGO3xhmNIRbGMmXxWELWclYD6yxrmE1iW7L57Ex2Bfsb
di97TFNDc6pmrGaRZp3mcc0BDsax4PA52ZxKziHODc57LQMtPy2x1mqtZq1+rTfaetq+2mLtcu0W
7eva73VwnUCdLJ31Om0693UJuja6UbqFutt1z+o+02PreekJ9cr1Dund0Uf1bfSj9Rfq79bv0R83
MDQINpAZbDE4Y/DMkGPoa5hpuNHwhOGoEctoupHEaKPRSaMnuCbuh2fjNXgXPmasbxxirDTeZdxr
PGFiaTLbpMSkxeS+Kc2Ua5pmutG003TMzMgs3KzYrMnsjjnVnGueYb7ZvNv8jYWlRZzFSos2i8eW
2pZ8ywWWTZb3rJhWPlZ5VvVW16xJ1lzrLOtt1ldsUBtXmwybOpvLtqitm63Edptt3xTiFI8p0in1
U27aMez87ArsmuwG7Tn2YfYl9m32zx3MHBId1jt0O3xydHXMdmxwvOuk4TTDqcSpw+lXZxtnoXOd
8zUXpkuQyxKXdpcXU22niqdun3rLleUa7rrStdP1o5u7m9yt2W3U3cw9xX2r+00umxvJXcM970H0
8PdY4nHM452nm6fC85DnL152Xlle+70eT7OcJp7WMG3I28Rb4L3Le2A6Pj1l+s7pAz7GPgKfep+H
vqa+It89viN+1n6Zfgf8nvs7+sv9j/i/4XnyFvFOBWABwQHlAb2BGoGzA2sDHwSZBKUHNQWNBbsG
Lww+FUIMCQ1ZH3KTb8AX8hv5YzPcZyya0RXKCJ0VWhv6MMwmTB7WEY6GzwjfEH5vpvlM6cy2CIjg
R2yIuB9pGZkX+X0UKSoyqi7qUbRTdHF09yzWrORZ+2e9jvGPqYy5O9tqtnJ2Z6xqbFJsY+ybuIC4
qriBeIf4RfGXEnQTJAntieTE2MQ9ieNzAudsmjOc5JpUlnRjruXcorkX5unOy553PFk1WZB8OIWY
EpeyP+WDIEJQLxhP5aduTR0T8oSbhU9FvqKNolGxt7hKPJLmnVaV9jjdO31D+miGT0Z1xjMJT1Ir
eZEZkrkj801WRNberM/ZcdktOZSclJyjUg1plrQr1zC3KLdPZisrkw3keeZtyhuTh8r35CP5c/Pb
FWyFTNGjtFKuUA4WTC+oK3hbGFt4uEi9SFrUM99m/ur5IwuCFny9kLBQuLCz2Lh4WfHgIr9FuxYj
i1MXdy4xXVK6ZHhp8NJ9y2jLspb9UOJYUlXyannc8o5Sg9KlpUMrglc0lamUycturvRauWMVYZVk
Ve9ql9VbVn8qF5VfrHCsqK74sEa45uJXTl/VfPV5bdra3kq3yu3rSOuk626s91m/r0q9akHV0Ibw
Da0b8Y3lG19tSt50oXpq9Y7NtM3KzQM1YTXtW8y2rNvyoTaj9nqdf13LVv2tq7e+2Sba1r/dd3vz
DoMdFTve75TsvLUreFdrvUV99W7S7oLdjxpiG7q/5n7duEd3T8Wej3ulewf2Re/ranRvbNyvv7+y
CW1SNo0eSDpw5ZuAb9qb7Zp3tXBaKg7CQeXBJ9+mfHvjUOihzsPcw83fmX+39QjrSHkr0jq/dawt
o22gPaG97+iMo50dXh1Hvrf/fu8x42N1xzWPV56gnSg98fnkgpPjp2Snnp1OPz3Umdx590z8mWtd
UV29Z0PPnj8XdO5Mt1/3yfPe549d8Lxw9CL3Ytslt0utPa49R35w/eFIr1tv62X3y+1XPK509E3r
O9Hv03/6asDVc9f41y5dn3m978bsG7duJt0cuCW69fh29u0XdwruTNxdeo94r/y+2v3qB/oP6n+0
/rFlwG3g+GDAYM/DWQ/vDgmHnv6U/9OH4dJHzEfVI0YjjY+dHx8bDRq98mTOk+GnsqcTz8p+Vv95
63Or59/94vtLz1j82PAL+YvPv655qfNy76uprzrHI8cfvM55PfGm/K3O233vuO+638e9H5ko/ED+
UPPR+mPHp9BP9z7nfP78L/eE8/sl0p8zAAAAIGNIUk0AAHolAACAgwAA+f8AAIDpAAB1MAAA6mAA
ADqYAAAXb5JfxUYAAAMAUExURQAAAP///+EGASMZ3Pz7/vX0/fb1/fLx/fPy/fTz/TIo3kY94UxD
4l9X5XJr6Hdw6Xhx6YR+64uF7JqV75uW75yX752Y76Sf8K+r8rGt87ay88fE9snG9s/M99DN99bU
+NjW+drY+dza+eLg+uvq/Ozr/O3s/O7t/O/u/CQa3CYc3Cce3Sge3Ssh3Swi3Swj3i0k3i4l3i8m
3jAn3jMq3zQr3zUs3zYt3zcu3zgv3zkw4Dkw3zox4Dsy4Dwz4D004D414D834UA44UI54UM64UQ7
4UU84UY+4khA4kpC4k1F409H41JK41JL5FNM5FZO5FlR5VtU5VxV5V9Y5mFa5mNc5mRd5mVe52Zf
52dg52hh52li52pj52tl6G5n6G9o6HNt6XRu6Xlz6nt16n136n546n9664F764J864R/7IaB7IiD
7IuG7Y2I7Y+K7ZGM7pCL7ZKN7pGM7ZOO7pSP7paR7piU75+b8KCc8KSg8aej8aik8auo8rKv87Sx
87e09Lm29Lu49Ly59L+89cG+9cPB9sXD9svJ983L99DO+NPR+NTS+N7d+uDf+uPi++Tj++Xk++bl
++fm++no/Ojn++rp/Ono+/Dw/fb2/vn5/vr6/vz8//7+//zo5/3r6uMTDuMUD+MVEOMWEeMXEuMY
E+MaFfCFgvGHhPGJhvGOi/KRjvKUkfKWk/Shn/SmpPWpp/WsqvWurPa0sve5t/e6uPe9u/e/vffB
v/rT0vrU0/rW1frY1/va2fvc2/vd3Pvg3/zj4vvi4fzl5Pzn5uEHAuEIA+EJBOEKBeILBuIMB+IN
COIPCuIQC+QbF+QdGeQfG+QgHOUjH+UlIeUpJeUqJuYtKeYvK+YzL+c2Muc5Nec8OOg+OuhAPOlE
QelGQ+lIRelKR+lMSepPTOtWU+tZVuteW+xiX+1mY+1qZ+1tau5wbu92dO99e/B/ffCAfvCCgPOZ
l/OenPjBwPjEw/jGxfnJyPnMy/nPzv3u7v3w8P719f729v75+f/7+//9/f///4DJalUAAAEAdFJO
U///////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////wBT9wclAAAP2klEQVR42uzdaXwURRoGcKsrJEAI
i0G5A4RATEICaBQWMYiEXcGA4Y6yHIICcikKAlkuJYRICCLHLklmhMUgghwiiAguiC6IcsgRORRR
dzlFFLwRE377wWRm+qjq7pnumeme5/mYeaeH6frT3dNVXXUTQRBGbsIuQIADAQ4EOBDgQIADAQ4E
OBDgQIADAQ4EAQ4EOBDgQIADAQ4EOBDgQIADAQ4EOBAEOBDgQIADAQ4EOBDgQIADAQ4EOBDgQBDg
QIADAQ4EOBDgQIADAQ4EOBDgQIADAQ4EAQ4EOBDgQIADAQ4EOBDgQIADAQ4EOBAEOBDgQIADAQ4E
OBDgQIADAQ4EOBDgQBDgQIADAQ4EOBDgQIADAQ4EOBDgQIADAQ4EAQ4EOBDgQIADAQ4EOBDgQIAD
AQ4EOBAEOBDgQIADAQ4EOBDgQIADAQ4EOBDgQBDgQIADAQ4EOHzP0uyH7klO6TSofiEIAIcohU+0
oRVpN8UBBMDhTuz91COZ/4QC4KhMo1QqSscCHqTJjw0Zk1MEHKGRxq2oJGmLWLUNe/5R0XZ8VeAI
gcSnUVkylI8MCYPdJe1mAof98xBVyGilyvx7RDVjncBh88xUskGT5skrZ7eWFA0uBg5bJ6yLIg4a
JavMSZYVPeQADjunHmUkRvWylVI6CjjsnO4sHDXFdYs7KFY9Bxz2TSzLBk2p5lnn6Klc1S4BOGyb
p5k4aH3PultYVcOBw7Z5kI3jEY+yuBRWVVIccNg0VZLYODp61A1gl40ADpsmj3IS7yqL4VTdEQ4c
9kw2D0dDV9lAXlkD4LBnonmtPs31MzaJV/Yn4LBnhvJafUJl1TO8KpoOHPZMFq/VoyurunJxJBYB
hy3Tl9fqlb9DFlN+CoDDlunPa/SbK4pmq+CYBxy2zEAtp5VsFRxNgcOWGcZr9Fsrip5UwbEIOGwZ
brtPrygazbdxuwM4bJl6vFZvXFE0nI8jCj9l7ZmmvJ+oVSqKRvBxjAcOeyasNbvRa1QWRfNxzAGO
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
AAAASUVORK5CYII='/>
```

### 1.2 alt 属性

当出现网络或者图片路径存在问题时，图片展示可能会出问题，这时网页中图片占位的地方会出现一张裂图，如果不做任何说明时，可能会影响用户体验。这时，如果在裂图上有一段文字说明的话，会适当的增加友好度，alt 属性就是图片加载失败时的替代文本，意思是当图片加载失败时使用 alt 定义的文本显示在图片当前的位置上，例如：

```javascript
<img src=""  alt="这是一张测试图片" />
```

### 1.3 align

定义图片的排列对齐方式，在 HTML5 中推荐使用 css 替代。

```javascript
<h2>设置对齐</h2>
<p>图像 <img src ="https://www.baidu.com/img/bd_logo1.png" align="middle" border=1> 在文本中</p>
<h2>未设置对齐</h2>
<p>图像 <img src ="https://www.baidu.com/img/bd_logo1.png" border=1> 在文本中</p>
```

以上定义居中的方式可以使用 css 的方式实现相同的效果，例如：

```javascript
<p>图像<img src ="https://www.baidu.com/img/bd_logo1.png" style="vertical-align: middle;border:1px solid">在文本中</p>
```

### 1.4 border

定义图片的边框，在 HTML5 中推荐使用 css 的 border 样式替代

```javascript
<img src="https://www.baidu.com/img/bd_logo1.png" border="2" />
```

上述定义了一个 border 是 2 的图片，同样可以使用 css 的方式替代：

```javascript
<img src="https://www.baidu.com/img/bd_logo1.png" style="border:2px solid" />
```

### 1.5 height

定义图片的高度，在 HTML5 中推荐使用 css 的 height 样式替代：

```javascript
<img src="https://www.baidu.com/img/bd_logo1.png" width="50" height="50">
<br />
<img src="https://www.baidu.com/img/bd_logo1.png" width="100" height="100">
```

上述定义了一个 height 分别是 50 像素和 100 像素的图片，如果使用 CSS 定义的话，是这样的：

```javascript
<img src="https://www.baidu.com/img/bd_logo1.png" style="width:50px;height:50px;">
<br />
<img src="https://www.baidu.com/img/bd_logo1.png" style="width:100px;height:100px;">
```

### 1.6 width

定义图片的宽度，在 HTML5 中推荐使用 CSS 替代。

### 1.7 ismap

属性定义图像为服务器端图片映射，就是定义图片中可以点击的区域，并且将对标发送到服务器，需要配合 a 标签使用。

```javascript
<a href="http://www.baidu.com">
<img src="https://www.baidu.com/img/bd_logo1.png" ismap />
</a>
```

上述代码，当用户点击图片时，浏览器会将点击的坐标点以 x，y 作为参数，以 GET 的方式请求服务器，例如当用户点击图片的坐标点（8,8）时，服务器会受到一条 HTTP 请求 `http://www.baidu.com?8,8`，此坐标点是相对定位坐标点。

### 1.8 hspace

定义图片左右侧的空白，在 HTML5 中推荐使用 CSS 替代。

```javascript
<h3>带有 hspace 和 vspace 的图像：</h3>
<p>
<img border="1" src="https://www.baidu.com/img/bd_logo1.png" height=100 width=100 align="middle" hspace="30" vspace="30" />
这是并排的文字信息
</p>
```

### 1.9 vspace

定义图片上下区域的空白，在 HTML5 中推荐使用 CSS 替代。

### 1.10 longdesc

定义图片描述文档的 URL ，这个属性和 alt 类似，只不过可以描述更多的文字 - 超过 1024 字符，但是目前还没有浏览器支持这个属性。

### 1.11 usemap

定义图片在客户端的可点击区域，需要配合 map 和 area 标签使用。

```javascript
<img src="https://wiki-code.oss-cn-beijing.aliyuncs.com/html5/img/area.jpg"  usemap="#planetmap" />
<map name="planetmap">
  <area href="https://pro.jd.com/mall/active/3ZYfZKGRAhbHzJySpRriJoGWo8v6/index.html?innerAnchor=46246884002&focus=3" shape="rectangle" coords="25,290,175,337">
</map>
```

## 2. css 定义背景图

可以结合 div 和 css 的 background-image 属性来定义背景图的方式显示一张图片，例如：

```javascript
<div style="background-image: url('https://www.baidu.com/img/bd_logo1.png');width:500px;height:300px"></div>
```

上述代码的也可以实现一张图像效果。

## 3. 注意事项

由于 HTML 中的元素中，图片相对于别的文本类型的元素所占用的网络空间较大，所以图片加载可能会比较慢，而且 HTML 中的元素网络请求是同步进行的，所以当定义了一些较大的图片或者图片元素较多时，网页本身会出现卡顿的情况，所以针对图片的加载需要做出一些优化：

* 尽可能少使用图片，或者使用体积较小的图片；
* 压缩图片体积；
* 通过懒加载的方式异步加载图片；
* 通过设置 HTTP 缓存时间，防止图片重复请求服务器资源；
* 通过使用 CSS 背景图的方式，将所需图片压缩到一张图片，减少请求图片次数。

## 4. 小结

本章介绍了几种 HTML 定义图片的方式，以及图片的各种属性，最后介绍了几种加载优化方式。图片是 HTML 中使用最频繁的多媒体标签，因此需要熟悉它的使用方式

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

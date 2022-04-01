# HTML 属性

本章介绍 HTML 的标签内属性控制方式以及几种通用的属性的简单介绍。属性是 HTML 标签内提供的附加信息，一般用于描述标签或者控制标签的展现形式。

属性大部分情况下以键值对方式出现，比如 `key='value'`。

## 1. 全局属性

全局属性是大部分标签都可以使用的属性，常用到的有以下 8 个：

1. accesskey
2. class
3. id
4. lang
5. style
6. tabindex
7. dir
8. title

下面我们具体来看一下他们的用法。

### 1.1 accesskey

accessKey 属性定义了使元素获得焦点的快捷键。

例如，定义了一组表单元素 ：

```javascript
<input type="text" accesskey='h' value="按下alt+h"/><br/>
<a href='http://www.baidu.com' accesskey="a">按下alt+a</a><br/>
<button accesskey="b">按下alt+b</button><br/>
<label accesskey="c">按下alt+c</label><br/>
<legend accesskey="d">按下alt+d</legend><br/>
<textarea accesskey="g">按下alt+g</textarea>
```

在浏览器中，当你同时按下 ALT+h 键，焦点会自动切换到上边的表单标签。

目前有 `<a>``<area>``<button>``<input>``<label>``<legend>``<textarea>`元素支持 accessKey 属性。

### 1.2 class

class 属性规定了元素的类名，类名命名必须以字母为开头，内容中可以包含大小写字母和下划线 ("_") 或者横杠 ("-")，类名是区分大小写的，类名可以定义多个，以 “ ” 空格隔开，例如定义了段落元素。

```javascript
<p class="class1 class2 class3"></p>
```

相当于给 p 元素定义了 class1 class2 class3 3 个类。

给元素定义类相当于给元素打了个标签，在 JavaScript 中操作 DOM 提供了便捷，例如可以通过：

```javascript
<p class="class1 class2 class3">我是一个段落</p>
<button onclick='hideTest()'>点击隐藏</button>
<button onclick='showTest()'>点击显示</button>
<script>
function hideTest(){
    var a = document.getElementsByClassName("class1"); //这样可以获取所有 class 包含test的标签
    a[0].style.display = 'none';
}
function showTest(){
    var a = document.getElementsByClassName("class1"); //这样可以获取所有 class 包含test的标签
    a[0].style.display = 'block';
}
</script>
```

以上 demo 通过点击按钮，基于操控对应 DOM 的 class 来控制按钮的显示与隐藏。

### 1.3 id

id 属性类似于 class，不同的是 id 是唯一标签，不能重复；

```javascript
<div id='test'></div>
```

类似于 class，id 也是用于操作 dom 的标记，例如：

```javascript
<p id='test' class="class1 class2 class3">我是一个段落</p>
<button onclick='hideTest()'>点击隐藏</button>
<button onclick='showTest()'>点击显示</button>
<script>
function hideTest(){
    var a = document.getElementById("test"); //这样可以获取所有 class 包含test的标签
    a.style.display = 'none';
}
function showTest(){
    var a = document.getElementById("test"); //这样可以获取所有 class 包含test的标签
    a.style.display = 'block';
}
</script>
```

以上 demo 通过点击按钮，基于操控对应 dom 的 class 来控制按钮的显示与隐藏。

### 1.4 lang

lang 属性定义了网页或者元素的语言，是否定义 lang 属性对网页或者标签的展现影响不大，但是可以帮助搜索引擎或者浏览器区分出网页的语言。ISO 639-1 为各种语言定义了缩略词。

```javascript
<p lang="fr">Ceci est un paragraphe.</p><!-- 表示语言是法语 -->
```

|Language|ISO Code|
|--------|--------|
|Abkhazian                |ab    |
|Afar                     |aa    |
|Afrikaans                |af    |
|Albanian                 |sq    |
|Amharic                  |am    |
|Arabic                   |ar    |
|Armenian                 |hy    |
|Assamese                 |as    |
|Aymara                   |ay    |
|Azerbaijani              |az    |
|Bashkir                  |ba    |
|Basque                   |eu    |
|Bengali (Bangla)         |bn    |
|Bhutani                  |dz    |
|Bihari                   |bh    |
|Bislama                  |bi    |
|Breton                   |br    |
|Bulgarian                |bg    |
|Burmese                  |my    |
|Byelorussian (Belarusian)|be    |
|Cambodian                |km    |
|Catalan                  |ca    |
|Chinese (Simplified)     |zh    |
|Chinese (Traditional)    |zh    |
|Corsican                 |co    |
|Croatian                 |hr    |
|Czech                    |cs    |
|Danish                   |da    |
|Dutch                    |nl    |
|English                  |en    |
|Esperanto                |eo    |
|Estonian                 |et    |
|Faeroese                 |fo    |
|Farsi                    |fa    |
|Fiji                     |fj    |
|Finnish                  |fi    |
|French                   |fr    |
|Frisian                  |fy    |
|Galician                 |gl    |
|Gaelic (Scottish)        |gd    |
|Gaelic (Manx)            |gv    |
|Georgian                 |ka    |
|German                   |de    |
|Greek                    |el    |
|Greenlandic              |kl    |
|Guarani                  |gn    |
|Gujarati                 |gu    |
|Hausa                    |ha    |
|Hebrew                   |he, iw|
|Hindi                    |hi    |
|Hungarian                |hu    |
|Icelandic                |is    |
|Indonesian               |id, in|
|Interlingua              |ia    |
|Interlingue              |ie    |
|Inuktitut                |iu    |
|Inupiak                  |ik    |
|Irish                    |ga    |
|Italian                  |it    |
|Japanese                 |ja    |
|Javanese                 |jv    |
|Kannada                  |kn    |
|Kashmiri                 |ks    |
|Kazakh                   |kk    |
|Kinyarwanda (Ruanda)     |rw    |
|Kirghiz                  |ky    |
|Kirundi (Rundi)          |rn    |
|Korean                   |ko    |
|Kurdish                  |ku    |
|Laothian                 |lo    |
|Latin                    |la    |
|Latvian (Lettish)        |lv    |
|Limburgish ( Limburger)  |li    |
|Lingala                  |ln    |
|Lithuanian               |lt    |
|Macedonian               |mk    |
|Malagasy                 |mg    |
|Malay                    |ms    |
|Malayalam                |ml    |
|Maltese                  |mt    |
|Maori                    |mi    |
|Marathi                  |mr    |
|Moldavian                |mo    |
|Mongolian                |mn    |
|Nauru                    |na    |
|Nepali                   |ne    |
|Norwegian                |no    |
|Occitan                  |oc    |
|Oriya                    |or    |
|Oromo (Afan, Galla)      |om    |
|Pashto (Pushto)          |ps    |
|Polish                   |pl    |
|Portuguese               |pt    |
|Punjabi                  |pa    |
|Quechua                  |qu    |
|Rhaeto-Romance           |rm    |
|Romanian                 |ro    |
|Russian                  |ru    |
|Samoan                   |sm    |
|Sangro                   |sg    |
|Sanskrit                 |sa    |
|Serbian                  |sr    |
|Serbo-Croatian           |sh    |
|Sesotho                  |st    |
|Setswana                 |tn    |
|Shona                    |sn    |
|Sindhi                   |sd    |
|Sinhalese                |si    |
|Siswati                  |ss    |
|Slovak                   |sk    |
|Slovenian                |sl    |
|Somali                   |so    |
|Spanish                  |es    |
|Sundanese                |su    |
|Swahili (Kiswahili)      |sw    |
|Swedish                  |sv    |
|Tagalog                  |tl    |
|Tajik                    |tg    |
|Tamil                    |ta    |
|Tatar                    |tt    |
|Telugu                   |te    |
|Thai                     |th    |
|Tibetan                  |bo    |
|Tigrinya                 |ti    |
|Tonga                    |to    |
|Tsonga                   |ts    |
|Turkish                  |tr    |
|Turkmen                  |tk    |
|Twi                      |tw    |
|Uighur                   |ug    |
|Ukrainian                |uk    |
|Urdu                     |ur    |
|Uzbek                    |uz    |
|Vietnamese               |vi    |
|Volapuk                  |vo    |
|Welsh                    |cy    |
|Wolof                    |wo    |
|Xhosa                    |xh    |
|Yiddish                  |yi, ji|
|Yoruba                   |yo    |
|Zulu                     |zu    |

### 1.5 style

style 属性定义了元素的行内样式（也叫内联样式），行内样式优先于任何其他形式定义的样式。

```javascript
<p style='color:#ccc'>这是一个行内样式示例</p> <!-- 行内样式 -->
```

其他两种定义元素样式的方式是：

* style 标签；
* 引入 css 样式文件。

```javascript
<style>
.test{
}
</style><!-- 使用标签的方式 -->

<link href="/css/test.css" /> <!-- 使用引用的方式 -->
```

其中 css 样式的优先级是： **行内样式 > 标签方式 > 引用文件** 方式。

关于样式的内容本章节不做具体讨论。

### 1.6 tabindex

tabindex 属性用于定义元素的 tab 键的控制顺序，即焦点的顺序。

tabindex 的值可以在 1 到 32767 之间的任意数字。当用户使用 tab 键在网页中移动控件时，将首先移动到具有最小 tabindex 属性的控件上，然后依次按从小到大移动。

如果两个元素的 tabindex 的属性值相等，浏览器会按照代码出现的顺序移动。当一个元素的 tabindex 设置为 -1，那么这个元素会排除在 tab 的移动顺序之外。

```javascript
<a href="http://www.baidu.com.cn/" tabindex="1">焦点1</a>
<a href="http://www.baidu.com.cn/" tabindex="2">焦点2</a>
<a href="http://www.baidu.com.cn/" tabindex="3">焦点3</a>
<a href="http://www.baidu.com.cn/" tabindex="4">焦点4</a>
<a href="http://www.baidu.com.cn/" tabindex="5">焦点5</a>
```

在上述示例中，通过点击键盘的 tab 键可以看到焦点依次按顺序通过链接标签。

### 1.7 dir

dir 属性定义元素内文本的方向。参数值有 2 种：

* ltr 默认值。文字按从左到右的方向；
* rtl 文字按照从右到左的方向。

```javascript
<p dir="rtl">这是一个文字从右到左展示的方式</p>
```

### 1.8 title

title 属性用于指定元素的备注信息。这些信息通常是指，鼠标移动到元素上并且停留一段时间时，浏览器的提示框内的文本。

```javascript
<p title="test123456789" style="word-wrap:break-all;width:30px;white-space: nowrap; text-overflow: ellipsis;overflow: hidden;">test123456789</p> <!-- 定义一个段落 -->
```

上面的代码实现的功能是：当段落文字超出段落宽度时，显示省略号，并且鼠标移上去可显示全部内容：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef04965090997f602200111.jpg)

## 2. HTML5 新增全局属性

### 2.1 contentEditable

这个属性是用来将一个标签变成可编辑的类型，可能了解的人并不太多，在加入 HTML5 标准以前，浏览器也有部分支持。比较通用的创建一个可编辑标签的方式是，使用表单 input 或者 textarea 的方式，不过都有一定的局限性。

* 会跟随 form 表单提交内容，有时候你可能并不需要；
* 展现样式比较单一。

使用 div + css + contentEditable 可以解决以上的问题，例如：

```javascript
<p contenteditable="true">这是一个可编辑的段落。</p>
```

以上定义了一个段落，如果没有加 contentEditable 属性的话，段落只是展示作用，当加了 contentEditable 属性之后，鼠标点击此段落之后可以编辑段落文字。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef0499f093bf17c06570131.jpg)

### 2.2 contextmenu

contextmenu 属性用于制定 div 元素的右键单击菜单，需要配合 menu 标签使用，例如：

```javascript
<div contextmenu='test'>
<menu type="context" id="test">
<menuitem label="刷新"></menuitem> <!-- 菜单1 -->
<menuitem label="回退"></menuitem> <!-- 菜单2 -->
</menu>
</div>
```

这个功能类似于桌面应用的右键菜单功能，但是目前只有 Firefox 浏览器支持 contextmenu 属性，不支持的浏览器可以使用 div + css 模拟实现。

### 2.3 data-*

这个属性是 HTML5 中用于存储自定义属性值，自定义属性值用于方便开发者存储一些简单的数据内容，而不需从服务器端获取。

在 HTML4 中自定义属性的方式很有可能会跟系统关键字冲突，而 data-* 会被客户端忽略。

* data- 后边必须至少有一个字符，不要包含大写字符；
* JavaScript 可以用 getAttribute 函数获取自定义属性；
* HTML5 原生函数支持使用 dataset / setAttribute 来 获取 / 操作自定义属性。

下面是 JavaScrip t 使用 getAttribute 函数获取自定义属性的例子：

```javascript
<script>
function show(animal) {
    var type = animal.getAttribute("data-type");
    alert(animal.innerHTML + "是一种" + type + "。");
}
</script>
<ul>
  <li onclick="show(this)" id="owl" data-type="标记语言">html</li>
  <li onclick="show(this)" id="salmon" data-type="脚本语言">JavaScript</li>
  <li onclick="show(this)" id="tarantula" data-type="层叠样式表">css</li>
</ul>
```

以上示例通过 data-type 保存了无序列表中每个条目的类型值，通过点击列表条目可以弹出类型值。

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef049e00971056d06450130.jpg)

### 2.4 draggable

这个属性用来标识元素是否支持被拖动，如果没有被设置则按照浏览器默认的方式来执行，属性可选值有 true/false/auto 。默认情况下，只有图片、链接是可以拖动的。需要配合定义 ondragstart 事件来实现拖动之后的响应机制。

```javascript
<p id="drag1" draggable="true" >这是一段可移动的段落。</p>
```

上述标签，当鼠标点击段落且移动时，呈现出可拖动样式：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef049e00971056d06450130.jpg)

如果未定义 draggable 属性，鼠标点击段落且拖动会选中文本：

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef04a1d098a2ba805100078.jpg)

### 2.5 dropzone

dropzone 属性是指在元素上拖动数据时，是否拷贝、移动或链接被拖动数据。目前所有主流浏览器暂不支持该属性。

### 2.6 hidden

hidden 用来设置元素是否应该被隐藏，当该属性设置为 hidden 或者 true 时，浏览器将不再渲染该元素。在早期的 HTML4 中，通过设置 css 样式 `display:none` 可以实现相同的效果。

```javascript
<input type=hidden id='test' />
```

上述表单在浏览器不显示任何效果，只有当提交表单时数据才会提交到服务器端。

### 2.7 spellcheck

spellcheck 属性定义是否对元素进行拼写检查，目前该属性支持非密码的 input 标签、textarea 中的文本，可编辑元素中的文本。spellcheck 的值有 true 和 false。

```javascript
<p  spellcheck="true">这是一个段落。</p>
```

### 2.8 translate

translate 属性定义渲染元素时是否需要对内容进行翻译，目前所有主流浏览器都不支持该属性。

## 3. 总结

回顾本章，介绍了 HTML 通用的 8 个属性，以及 HTML5 新增的 8 个通用属性。通用的属性一般对大部分元素标签都支持，但是 HTML5 目前是新标准，所有有些属性不能百分之百兼容所有的浏览器。因为浏览器更新迭代速度较快，所以在此没有对浏览器兼容性做深入讲解，大家在实际使用的时候可以在 w3c 官网或者其他网站手册中查询对照。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

# HTML 多媒体

本章介绍 HTML 中的多媒体。多媒体是计算机中用于人机进行实时交互的媒介和互动方式，其中包括图片、文字、音频、视频、动画等。之前的章节已经介绍了图片和文本元素，本章主要介绍音频和视频相关的元素

## 1. object

object 元素用于定义一个网页插件，使用该元素可以在网页中嵌入多媒体，支持的插件类型有图形、音频、视频、java applets、ActiveX、PDF、flash。不过在 HTML5 标准之前，主流浏览器中只有 IE3.0 之后的版本支持

```javascript
<object classid="clsid:F08DF954-8592-11D1-B16A-00C0F0283628" id="Slider1"
width="100" height="50">
  <param name="BorderStyle" value="1" />
  <param name="MousePointer" value="0" />
  <param name="Enabled" value="1" />
  <param name="Min" value="0" />
  <param name="Max" value="10" />
</object>
```

虽然 object 标签的出现是用于取代 HTML 中的多媒体相关的标签，但是由于目前只有 IE 浏览器兼容性较好，所以未能实现初衷。虽然在手册或者文献中未能找到其他浏览器不兼容的原因，但是作者猜想是因为一方面 object 的属性太多导致入门门槛较高，一方面互联网朝向细分化方向发展而技术也在细分化，所以这种大一统的技术解决方案不太受主流大厂的接受，进而导致各大浏览器厂商认可度不高（具体原因以官方手册为准）。虽然兼容性不好，但是在 IE 低版本中 object 使用频率不低，主要用于 flash

### 1.1 archive 属性

定义对象相关的档案文件的 URL 列表

### 1.2 classid 属性

定义嵌入网页的插件的自身的 id 值，这个值可以是 flash 的对象 id 或者 Java 的类的 classid

```javascript
<object classid="clock.class"></object>
```

浏览器根据这个 classid 可以找到插件的源代码

### 1.3 codebase 属性

这是一个可选的属性，定义一个 URL 指向 classid 属性所引用的对象的代码

```javascript
<object classid="clock.class" codebase="http://www.baidu.com/classes/">
</object>
```

codebase 相当于 源代码的域名，根据 codebase+classID 浏览器可以准确定位插件的源代码。

### 1.4 codetype 属性

类似于 type 属性，用来标识 object/applet 插件的 mime 类型。

```javascript
<object codebase="clock.class" codetype="application/java">
</object>
```

上述代码定义了一个 Java 类型的插件。

### 1.5 data 属性

用于指定用于对象处理的 URL，该属性和 classID 类似，只不过是当 object 用于定义一个 多媒体（例如图片）格式的时候，data 用于表示图片的源路径，等价于 img 标签的 src 属性，不同的是 img 标签的 src 属性只能用于图像类型，data 则可用于几乎任意的多媒体文件类型。

```javascript
<object codebase="clock.class" codetype="application/jpeg" data=''>
</object>
```

### 1.6 declare 属性

通常用于当 object 插件是 flash 或者 applet 时，定义插件编程语言的前置声明，不参与展示效果。

### 1.7 form 属性

定义 object 元素对应的表单的 id，目前没有主流浏览器支持 form 属性。

### 1.8 standby 属性

定义当 object 在加载的过程中，浏览器的 object 位置上显示的文本，此属性类似于 img 标签中的 alt 属性。

### 1.9 height width 属性

定义插件的高度和宽度。

### 1.10 hspace vspace 属性

定义插件的周围的空白。

## 2. embed

这个标签的作用基本上和 object 相似，都是用来嵌入插件，不同的是 object 只有 IE 内核的浏览器支持比较完善，而非 IE 内核的浏览器则对 embed 支持度较高。embed 可以嵌入到 object 标签内，当浏览器不支持 object 时会自动加载 object 内的元素来渲染。

```javascript
<embed src="helloworld.swf" /> <!-- 定义一个flash -->
```

## 3. param

该标签用于为 object 插件定义运行时参数，控制插件的运行方式。

```javascript
<object classid="clsid:F08DF954-8592-11D1-B16A-00C0F0283628" id="test">
  <param name="BorderStyle" value="1" />
  <param name="MousePointer" value="0" />
  <param name="Enabled" value="1" />
  <param name="Min" value="0" />
  <param name="Max" value="10" />
</object>
```

上述定义的每个 param 相当于一个插件的控制参数，name 对应参数的名称，value 对应参数的值，根据具体的插件的不同而不同。

## 4. audio

audio 是 HTML5 中的新标签，用于播放音频流，格式如下：

```javascript
<audio src="test.wav">
您的浏览器不支持 audio 标签。 <!-- 定义浏览器不支持该标签时所展示的内容-->
</audio>
```

上述代码相当于在浏览器中定义了一个播放器，展示效果如下

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef0539c092cf53504310099.jpg)

### 4.1 autoplay 属性

当定义了该属性时，音频会立即播放。

### 4.2 controls 属性

该属性用于定义是否向用户展示播放按钮，如果未定义这个属性，则浏览器默认不显示任何播放器的控件（播放、音量、进度等）。控件可以通过 audio 提供的事件，通过 div+css+JavaScript 自定义实现。

### 4.3 loop 属性

该属性定义时，表示音频是否会重复循环播放。

### 4.4 muted 属性

当该属性被定义是，音频默认播放时是以静音方式播放。

### 4.5 preload 属性

该属性定义音频加载的方式，有 3 个可选值（auto 自动加载 meta 只加载元数据 none 手动加载）。当设置了 autoplay 时，该属性不生效。

### 4.6 src 属性

定义音频的 URL。

## 5. source

该标签用于定义音频 / 视频播放器的源文件，为了兼容不同的浏览器对不同的音频类型的支持，例如：

```javascript
<audio>
   <source src="test.ogg" type="audio/ogg"> <!-- 定义Ogg类型的音频 -->
   <source src="test.mp3" type="audio/mpeg"> <!-- 定义MP3类型的音频-->
</audio>
```

上述音频播放器，放到不支持 Ogg 文件的浏览器时，会自动加载 MP3 文件。

## 6. video

这个标签用于定义视频播放器，大部分属性和 audio 类似，是 HTML5 中新增的标签。

```javascript
<video src="/i/movie.ogg" controls="controls">
your browser does not support the video tag
</video>
```

视觉效果如下

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef053cf09856a6704750333.jpg)

## 7. 兼容性

HTML 中定义多媒体元素不是很复杂，但是难点在于各种浏览器的兼容性问题，例如：

* iPhone iPad 不支持 flash
* 低版本 IE 不支持 embed
* 非 IE 浏览器不完全支持 object
* 音视频标签 audio video 仅支持 HTML5 标准
* 多媒体文件格式在不同浏览器的支持程度不同
* 多媒体文件格式在不同的硬件上的支持程序也不同

所以说如果你需要大量使用多媒体技术，那么必须要非常熟悉多媒体技术在浏览器上的兼容性问题才行

## 8. 小结

本章介绍了多媒体技术在 HTML4 和 HTML5 中的使用方式，以及在不同的浏览器中的兼容性问题

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

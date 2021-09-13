# 认识 img 图片标签

在我们的网页当中，图片肯定是必不可少的元素，有了图片之后，我们的网页网站就会变得更加的丰富起来。那么我们要在 HTML 当中插入图片的话，就会用到我们的 img 图片标签了。例如：

![图片描述](https://img.mukewang.com/wiki/5f07c9f209b0a71306470190.jpg)

图文并茂的网站。

## 1. IMG 标签的使用

img 标签为单标签，所以没有尾标签。而 img 标签有一个 必填的属性： src 属性，代表图片的路径。图片的路径可以为图片的相对路径，绝对路径和网络路径。这里我们以网络路径示范：

```javascript
<img src="https://www.imooc.com/static/img/index/logo-recommended.png" alt="慕课网logo">
```

会呈现以下效果：

![图片描述](https://img.mukewang.com/wiki/5f07ca00091f636704620169.jpg)

img 标签还有一个属性为 alt 属性，alt 属性表示错误提示，如果图片的路径既 src 属性的地址出错时，会显示 alt 属性的内容，我们把上面例子稍微修改一下，把路径写错，既：

```javascript
<img src="https://www.imooc.com/static/img/index/logo-recommended123.png" alt="慕课网logo">
```

则会在页面上呈现以下效果：

![图片描述](https://img.mukewang.com/wiki/5f07ca14097838a003310060.jpg)

如果图片路径出错，那么就不会显示对应的图片，而会显示一个图片加载失败的样式。这时候我们编写的 ALT 属性里面的内容也会呈现在页面上，就表示我们当前这张加载失败的图片的错误提示。

**注意：**

alt 属性的内容只会在图片加载失败时显示，既图片路径出错时，如果图片加载成功，则 alt 属性的内容会自动隐藏。 alt 属性起到一个提示作用，如果我们在编写代码时，不小心把图片路径写错，则会提示我们该张图片为哪一张，具体内容是什么。而 alt 属性还有一个重要的作用就是网站的 SEO 作用 （搜索引擎爬取网页的内容），假设由于下列原因用户无法查看图像，alt 属性可以为图像提供替代的信息：

1. 网速太慢；
2. src 属性中的错误；
3. 浏览器禁用图像；
4. 用户使用的是屏幕阅读器。

其重要性主要有： 网页内容相关性是关键词优化的前提，搜索引擎认为，网页上的图片应该与网页主题相关。

反过来讲，当搜索引擎要判断网页的关键词时，图片的 alt - 代替属性是一个可信任的参考点。所以， 别忘了在图片的 alt - 代替属性。

img 标签还可以设置 `width` 和 `height` 属性来改变图片的宽高，**例如：**

```javascript
<img src="https://www.imooc.com/static/img/index/logo-recommended.png" alt="慕课网logo" width="100" height="100">
```

则会呈现以下效果：

![图片描述](https://img.mukewang.com/wiki/5f07ca2b097672de02580102.jpg)

图片的宽高为 100 * 100。

> **Tips**：IMG 标签的 `width` 和 `height` 属性的值为数字，一般情况下为正整数。

## 2. 注意事项

1. img 标签为单标签，没有尾标签；
2. img 标签的 `width` 和 `height` 都不要写单位，直接写数字即可；
3. src 是 img 标签的必填属性。

## 3. 真实案例分享

百度搜索（部分）

```javascript
<ul>
    <li>
      <a>
        <img
          src="https://pics0.baidu.com/feed/bd3eb13533fa828b324d63ea7f258e32970a5a0a.jpeg">
      </a>
    </li>
    <li>
      <a>
        <img
          src="https://pics4.baidu.com/feed/30adcbef76094b36eaacf1e422f6b3df8d109dba.jpeg">
      </a>
    </li>
    <li>
      <a>
        <img
          src="https://pics5.baidu.com/feed/730e0cf3d7ca7bcb8af180f33f33a465f724a821.jpeg">
      </a>
    </li>
  </ul>
```

淘宝网（部分）

```javascript
<div>
   <img src="//img.alicdn.com/imgextra/i3/180185321/O1CN01agBO561pB43nm30sJ_!!180185321-0-beehive-scenes.jpg_180x180xzq90.jpg_.webp">
</div>
<div>
   <h4>Huawei/华为 p40 pro</h4>
   <p>MUI 10.1系统，内置华为AI语音助手Celia ，支持“Hey Celia”语音唤醒。HMS服务，华为P40搭载华为HMS服务。</p>
   <p><span></span>0 人说好</p>
</div>
```

## 4. 小结

1. img 标签为单标签，没有尾标签。
2. img 标签的 src 属性为必选项，其余属性为可选项。
3. img 标签可以嵌套在任意标签里。
4. img 标签图片路径可以为图片的相对路径，绝对路径和网络路径，通常情况下我们采用相对路径。

![图片描述](https://img.mukewang.com/wiki/5f6301f409baab4b14250831.jpg)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

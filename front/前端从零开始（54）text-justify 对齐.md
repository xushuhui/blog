# text-justify

这个属性不怎么常用，因为它的兼容性不好，只兼容 IE 浏览器，它主要是给对齐属性`text-align:justify`做一个补充。

## 1. 官方定义

改变字与字之间的间距使得每行对齐。

## 2. 慕课解释

这个属性主要用来页面文字的排版，如果我们一个段落不设置任何属性，那么它的每一行有长有短，很不美观，通过这个属性，可以让每一行都能实现左右对齐。

我们首先要设置`text-align:justify`然后再设置`text-justify`去告诉浏览器使用什么样的排版方式让文字对齐。而不设置`text-justify`浏览器则使用默认的方式让其实现两端对齐。

这个属性只兼容 IE 浏览器。而其他浏览器的对齐方式仅受`text-align:justify`对齐方式的影响。

## 3. 语法

```javascript
.demo{
    text-align:justify;
    text-justify:inter-word;
}
```

属性值说明

|值|描述|
|--|----|
|auto           |浏览器自行渲染                                        |
|none           |禁用齐行。                                            |
|inter-word     |增加 / 减少单词间的间隔。                               |
|inter-ideograph|用表意文本来排齐内容。                                |
|inter-cluster  |只对不包含内部单词间隔的内容（比如亚洲语系）进行排齐。|
|distribute     |类似报纸版面，除了在东亚语系中最后一行是不齐行的。    |
|kashida        |通过拉伸字符来排齐内容。                              |

## 4. 兼容性

|IE|Edge|Firefox|Chrome|Safari|Opera|ios|android|
|--|----|-------|------|------|-----|---|-------|
|9+|no|no|no|no|no|no|no|

## 5. 实例

1. 默认文字排版。

```javascript
    <div class="demo">
        To a large degree，
        the measure of our peace of mind
        is determined by how much we are
        able to live in the present moment．
    </div>
    <div class="demo-1">
        轻轻的我走了，
        正如我轻轻的来；
        我轻轻的招手，
        作别西天的云彩。
        那河畔的金柳，
        是夕阳中的新娘；
        波光里的艳影，
        在我的心头荡漾。
        软泥上的青荇，
        油油的在水底招摇；
        在康河的柔波里，
        我甘心做一条水草！
    </div>
```

```javascript
.demo{
    background: #f2f2f2;
    margin-bottom: 10px;
}
.demo-1{
    background: #a2a2a2;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e1e709c12c2305270176.jpg)

默认文字排版效果图

说明：这两端字符第一段是英文，第二段是中文他们都没有实现两端对齐。中文还好，英文的排版最差，这是因为英文单词不像汉字，它长短不一。

下面我们通过设置`text-justify`中包含的各种属性来看看，他们都是怎么实现两端对齐的。

1. 使用浏览器默认对齐方式。

```javascript
.demo{
    background: #f2f2f2;
    margin-bottom: 10px;
    text-align:justify;
}
.demo-1{
    background: #a2a2a2;
    text-align:justify;
}
```

或

```javascript
.demo{
    background: #f2f2f2;
    margin-bottom: 10px;
    text-align:justify;
    text-justify:auto;
}
.demo-1{
    background: #a2a2a2;
    text-align:justify;
    text-justify:auto;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e1fa092bf7f905270178.jpg)

浏览器默认对齐方式效果图

说明：直接设置`text-align:justify;`就会实现文字两端对齐，对齐方式使用浏览器默认方式。

1. 使用 `inter-ideograph` 来实现文字对齐。

```javascript
.demo{
    background: #f2f2f2;
    margin-bottom: 10px;
    text-align:justify;
    text-justify: inter-ideograph;
}
.demo-1{
    background: #a2a2a2;
    text-align:justify;
    text-justify: inter-ideograph;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e2070913e27705210172.jpg)

使用 `inter-ideograph` 来实现文字对齐效果图

说明：通过设置`inter-ideograph`，让 IE 浏览器使用表意文本对齐方式对齐内容 。

1. 使用 `inter-word` 来实现文本排版对齐。

```javascript
.demo{
    background: #f2f2f2;
    margin-bottom: 10px;
    text-align:justify;
    text-justify: inter-word;
}
.demo-1{
    background: #a2a2a2;
    text-align:justify;
    text-justify: inter-word;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e214093730d605230173.jpg)

使用 `inter-word` 来实现文本排版对齐效果图

说明：如图所示，文字还是对齐了，如果和 `inter-ideograph` 的效果图对比还是有细微差别，它的对齐方式修改了单词之间的距离。所以说这只是 IE 浏览器在对齐的时候一种排版方式。

1. 使用 `inter-cluster` 来实现文本排版对齐。

```javascript
.demo{
    background: #f2f2f2;
    margin-bottom: 10px;
    text-align:justify;
    text-justify: inter-cluster;
}
.demo-1{
    background: #a2a2a2;
    text-align:justify;
    text-justify: inter-cluster;
}
```

效果图

![图片描述](https://xushuhui.gitee.io/image/imooc/5ea3e2220934d33b05190168.jpg)

使用 `inter-cluster` 来实现文本排版对齐效果图

由此可见使用其他属性`distribute`、`kashida`都只是改变 IE 浏览器的一种对齐方式。

## 6. 小结

1. 要使用这个属性一定要先设置 `text-align:justify;`
2. 且仅仅兼容 IE 浏览器。

��号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

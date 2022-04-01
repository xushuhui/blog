# HTML 表格

除了图像、多媒体这种丰富的交互效果之外，网页中还经常会用到通讯录、统计报表这种格式化的交互效果，这时候就用到了表格元素。

## 1. 创建一个简单的表格

网页中定义表格使用 table 标签，它是一个闭合标签，所有内容都放在 table 的起始标签和结束标签内。在表格中定义一行数据使用 tr 标签，表格的单元格内容放在 tr 标签内，单元格又分为表头 th 和表格单元格 td。基本表格结构如下。

```javascript
<table>
<thead>
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td>action</td><td>调用控制器类的操作</td></tr>
    <tr><td>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

代码解释：以上是 thinkPHP 手册中的关于助手函数的表格示例，其中包含了表头 thead 、内容 tbody、行元素 tr、表头单元格 th、表格单元格 td。

## 2. 相关的标签和属性介绍

### 2.1 table 标签

table 用于声明一个表格，它本身不显示任何内容，例如：

```javascript
<table border="1"></table>
```

```javascript
<table border="1"></table>
<!--只显示一个像素点。-->
```

代码解释：上述定义了一个只有边框而没有内容的表格，由于没有任何内容，它在网页中只显示一个像素点。

#### 2.1.1 bgcolor 属性

这个属性可以应用于表格的任意元素，用来定义元素的背景颜色，其内容可以用 rgb 值表示。实际项目中通常用于在表头设置背景，为了方便和单元格区分开来，例如：

```javascript
<table>
<thead bgcolor="#ccc">
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td>action</td><td>调用控制器类的操作</td></tr>
    <tr><td>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

代码解释：上述代码将表头单元格设置为灰色背景。HTML 中颜色值可以使用 3 种表示方式：

* RGB；
* 16 进制；
* 英文代码（部分颜色可用）

例如：

```javascript
<table bgcolor="#FF0000"></table><!--16进制表示法表示“纯红色”-->
<table bgcolor="rgb(255,0,0)"></table><!--RGB表示法表示“纯红色”-->
<table bgcolor="Red"></table><!--英文法表示“纯红色”-->
```

```javascript
<table bgcolor="#FF0000">
<thead>
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td>action</td><td>调用控制器类的操作</td></tr>
    <tr><td>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>

<table bgcolor="rgb(255,0,0)">
<thead>
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td>action</td><td>调用控制器类的操作</td></tr>
    <tr><td>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>

<table bgcolor="Red">
<thead>
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td>action</td><td>调用控制器类的操作</td></tr>
    <tr><td>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

注意，在实际项目中，这个属性建议使用 css 的方式替代。

#### 2.1.2 align 属性

定义表格元素的对齐方式，**一般建议使用 css 替代**，可选项有 left、center、right，例如：

```javascript
<table>
<thead bgcolor="#ccc">
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td align='center'>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td align='center'>action</td><td>调用控制器类的操作</td></tr>
    <tr><td align='center'>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

代码解释：上述代码中，将左边列设置为居中对齐。表格中的表头 th 元素默认是居中对齐，而单元格元素默认是靠左对齐（跟浏览器相关），所以实际项目中一般会将类别列设置为居中对齐，实际项目建议使用 css 的方式替代。

#### 2.1.3 border 属性

表格默认是没有边框的，通过这个属性可以给表格设置边框，但是一般情况下使用 css 设置边框样式比较常见。

```javascript
<table border=3><!--设置3个像素的边框-->
<thead bgcolor="#ccc">
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td align='center'>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td align='center'>action</td><td>调用控制器类的操作</td></tr>
    <tr><td align='center'>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

代码解释：如果给 table 元素设置 border 属性，表格内部的单元格都会继承这个属性，但是如果 border 设置值超过 1 的话，单元格的边框还是只有 1，但是表格外部边框会显示 border 值对应的宽度；如果 border 设置 0 的话表示隐藏边框。

#### 2.1.4 cellpadding 属性

通过这个属性可以设置单元格内容到边沿之间的空隙，这个值设置的越大则单元格越大，不设置的话默认是 2 个像素。如下，我们设置一个较大空隙的单元格：

```javascript
<table border=1 cellpadding=10><!--10个像素的内间隙-->
<thead bgcolor="#ccc">
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td align='center'>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td align='center'>action</td><td>调用控制器类的操作</td></tr>
    <tr><td align='center'>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

> **注意：** 请勿将该属性与下面的 cellspacing 相混淆。

#### 2.1.5 cellspacing 属性

**与上个属性不同的是，cellspacing 指的是单元格与表格外边框之间的空隙**，如果不设置的话默认是 2 个像素，设置为 0 表示没有空隙，例如：

```javascript
<table border=1 cellspacing=0><!--没有外间隙-->
<thead bgcolor="#ccc">
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td align='center'>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td align='center'>action</td><td>调用控制器类的操作</td></tr>
    <tr><td align='center'>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
<br/>
<table border=1 cellspacing=10><!--间隙为10-->
<thead bgcolor="#ccc">
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td align='center'>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td align='center'>action</td><td>调用控制器类的操作</td></tr>
    <tr><td align='center'>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

#### 2.1.6 frame 属性

定义外侧哪个边框是可见的，实际项目中很少用到。

#### 2.1.7 rules 属性

定义内侧边框哪个是可见的，实际项目中很少用到。

#### 2.1.8 summary 属性

定义表格摘要。这个属性对表格的外观没有任何作用，相当于对表格做了一个备注，可能对搜索引擎有帮助，例如：

```javascript
<table  summary="这个表格用于演示">
  <tr>
    <th>month</th>
    <th>Savings</th>
  </tr>
  <tr>
    <td>January</td>
    <td>$100</td>
  </tr>
</table>
```

### 2.2 th 标签

th 用于定义表头类型单元格，他和内容单元格的区别主要在于样式上，表头默认显示粗体居中的文本，例如：

```javascript
<table border=1>
<thead>
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td>action</td><td>调用控制器类的操作</td></tr>
    <tr><td>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

#### 2.2.1 colspan 属性

colspan 设置单元格横跨的列数，这个属性相当于 Word 表格的合并单元格，例如：

```javascript
<table border=1>
<thead>
	<tr><th colspan=2>thinkPHP</th></tr><!--定义一个横跨2列的单元格表头-->
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td>action</td><td>调用控制器类的操作</td></tr>
    <tr><td>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

> **注意：** 如果 colspan 的值超过当前单元格列数的话，默认只能横跨最大列数。

#### 2.2.2 rowspan 属性

与上述属性类似，这个是用于横跨单元行，例如：

```javascript
<table border=1>
<thead>
    <tr><th></th><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td rowspan=3>函数类别</td><td>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td>action</td><td>调用控制器类的操作</td></tr>
    <tr><td>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

#### 2.2.3 valign 属性

这个属性用于设置单元格内的内容在垂直方向的对齐方式，目前支持上对齐、下对齐、居中对齐、文字基线对齐。

### 2.3 tr 标签

tr 用于定义表格中的行，行内可以包含表头单元格 th 或者数据单元格 td, tr 标签表示一行，当中的内容可以是 th 或者 td，tr 标签支持所有的浏览器。

### 2.4 td 标签

表格中有两种类型的单元格，一种是上边介绍的表头单元格 th，一种是内容单元格 td。和 th 不同的是 td 的内容默认显示正常字体，使用左对齐方式。td 支持的属性基本上和 th 一致。td 必须包含在 tr 标签中才能生效。

### 2.5 非必须标签

除了表格、行、单元格之外，表格还有一些其他的标签可使用。这些标签是非必须的，但是可以增强表格的内容丰富度和视觉效果。

#### 2.5.1 thead tbody tfoot 标签

thead 用于定义表格的表头，同样命名为表头，它和 th 表头单元格有本质的区别。

thead 实质上是用于对表格的内容进行分组，用于告诉开发者或者搜索引擎表格的哪部分是头，哪部分是内容，哪部分是尾部。所以说 thead 需要和 tbody 、tfoot 结合使用才有效果，正常情况下定义 thead 不会影响到表格的样式和布局，除非你强制定义 thead 的 css 样式。

目前主流的浏览器大都支持 thead、tbody、tfoot 标签，例如：

```javascript
<table border=1>
    <thead>
        <tr>
            <th >Title</th>
            <th >ID</th>
            <th >Country</th>
            <th >Price</th>
            <th >Download</th>
        </tr>
    </thead>
    <tbody>
        <tr >
            <th >Tom</th>
            <td>1213456</td>
            <td>Germany</td>
            <td>$3.12</td>
            <td>Download</td>
        </tr>
        </tbody>
    <tfoot>
        <tr>
            <th >Total</th>
            <td colspan="4">4 books</td>
        </tr>
    </tfoot>
</table>
```

#### 2.5.2 col 和 colgroup 标签

col 标签用来为表格中的列统一设置属性值，使用它主要用来节省代码量。它是一个单标签，无需结束标签，colgroup 标签是 col 的升级版，不同于 col 的是 colgroup 主要以组合的方式对列设置属性样式，而且 col 可以嵌入到 colgroup 内部进行细化的设置。

```javascript
<table width="100%" >
  <col align="center" />
  <tr>
    <th>title</th>
    <th>title</th>
    <th>title</th>
  </tr>
  <tr>
    <td>css</td>
    <td>HTML</td>
    <td>JavaScript</td>
  </tr>
</table>
```

例如以上代码使用 col 对列进行居中设置，由于 colgroup 和 col 这两个标签存在严重的浏览器兼容问题，如上图所示，在 Chrome、FireFox、Safari 和 ie8+ 等浏览器中不再支持 COL 及 COLGROUP 元素的部分属性，所以建议最好不要使用。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f053ad009f9095208090468.jpg)

#### 2.5.3 caption 标签

caption 用于定义表格的标题。每个表格只能声明一个标题，默认显示在表格的正上方，仅仅起到一个展示的作用，实际上跟 table 关联不大，完全可以用一个文本类型的标签元素替代：

```javascript
<table border=1>
<caption>系统为一些常用的操作方法封装了助手函数，便于使用，包含如下：</caption>
<thead>
    <tr><th>助手函数</th><th>描述</th></tr>
</thead>
<tbody>
    <tr><td>abort</td><td>中断执行并发送HTTP状态码</td></tr>
    <tr><td>action</td><td>调用控制器类的操作</td></tr>
    <tr><td>app</td><td>快速获取容器中的实例 支持依赖注入</td></tr>
</tbody>
</table>
```

上述代码展示效果

![图片描述](https://xushuhui.gitee.io/image/imooc/5ef0517909a5dccd04670209.jpg)

## 3. 表格插件

实际项目开发中并非仅仅是单纯的展示一个表格，有可能需要动态的插入单元格数据、动态的删除行或者列、数据分页、异步加载、或者是动态的修改表格的内容等等的操作，以上这些操作需要通过 JavaScript 去操作进行。项目开发中崇尚对高度相同操作的函数进行封装打包，这也是面向对象的精髓（虽然 JavaScript 不是面向对象语言）。我们可以手动封装一个对 table 表格进行操作的类，也可以使用现成的表格插件，现在介绍一个**表格插件 layui-table**：

```javascript
<html>
<head>
<link rel="stylesheet" href="https://www.layuicdn.com/layui-v2.5.4/css/layui.css" media="all"><!--引用css文件-->
<script src="https://www.layuicdn.com/layui-v2.5.4/layui.js"></script><!--引用JavaScript文件-->
</head>
<body>
<table id="demo" lay-filter="test"></table><!--定义表格容器-->
<script>
layui.use('table', function(){//初始化表格控件
  var table = layui.table;
  table.render({//渲染表格
    elem: '#demo'//表格容器的ID
    ,height: 312//生成的表格的高度
    ,url: '/demo/table/user/' //数据接口
    ,page: true //开启分页
    ,cols: [[ //表头设置
      {field: 'id', title: 'ID', width:80, sort: true, fixed: 'left'}
      ,{field: 'username', title: '用户名', width:80}
      ,{field: 'sex', title: '性别', width:80, sort: true}
      ,{field: 'city', title: '城市', width:80}
      ,{field: 'sign', title: '签名', width: 177}
      ,{field: 'experience', title: '积分', width: 80, sort: true}
      ,{field: 'score', title: '评分', width: 80, sort: true}
      ,{field: 'classify', title: '职业', width: 80}
      ,{field: 'wealth', title: '财富', width: 135, sort: true}
    ]]
  });
});
</script>
</body>
</html>
```

## 4. 小结

本章介绍了 HTML 中用途较多的表格元素，表格的大部分属性跟 Word 中的表格类似，除了 table 、td 、th、tr 这几个标签之外，其他的标签大都是用于分组或者设置属性，不参与视觉样式的表现，而且在 HTML5 中 css 的功能已经非常强大和完善了，所以几乎无需使用 HTML 中的样式设置方式，总的来说基本上只需要掌握上述的几种标签即可。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

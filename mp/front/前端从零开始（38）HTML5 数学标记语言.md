# HTML5 数学标记语言

本章介绍一个比较专业性的 HTML 知识点 - 数学置标语言，它是一种基于 XML 的标准，用来在互联网上书写数学符号和公式的置标语言，由万维网联盟的数学工作组提出。

## 1. 发展史

### 1.1 版本

* 1.01 版于 1999 年 7 月公布
* 2.0 版于 2001 年 2 月出现。
* 万维网联盟在 2003 年 10 月发布了 MathML 2.0 的第二版
* 2010 年 10 月发布了 MathML 3.0。

### 1.2 作用

MathML 包含两个子语言：Presentation MathML 和 Content MathML。Presentation MathML 主要负责描述数学表达式的布局。Content MathML 主要负责标记表达式的某些含义或数学结构。MathML 的这一方面受到 OpenMath 语言的很大影响，在 MathML3 中，与 OpenMath 更为贴近。为什么需要使用 MathML 呢？

我们要想在网页中插入一些数学公式，早期的方式只能通过特殊符号或者图片来实现，现在利用 MathML 标签可以方便款姐的实现所有数学公式的显示。

## 2. 语法简介

介绍语法之前先看一个简单的 demo：

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
<mrow>
<mrow>
<mi>x</mi>
<mo>+</mo>
<mi>y</mi>
</mrow>
<mo>=</mo>
<mn>10</mn>
</mrow>
</math>
```

是不是看起来有点眼熟，没错，它就是基于 XML 语法格式。math 元素之最顶层的根元素，每个实例都必须包含在 math 标签内，一个 math 元素不能包含另一个 math 元素。

### 2.1 mrow

mrow 标签用于对表达式进行分组，至少由一个到多个运算符组成，此元素呈水平展示。良好的分组对数学表达式结构展示起到一定的改善效果。

### 2.2 mi

mi 标签表示运算表达式中的变量，例如：

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mrow>
      <mi>a</mi>  <!--表示变量-->
      <mo>+</mo>
      <mi>b</mi>  <!--表示变量-->
   </mrow>
</math>
```

### 2.3 mo

mo 标签用于表示运算符，例如：

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mrow>
      <mi>a</mi>
      <mo>-</mo><!--减法运算符-->
      <mi>b</mi>
   </mrow>
</math>
```

### 2.4 mn

mn 标签用于显示表达式中的数字。

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mrow>
      <mn>5</mn><!--表示数字-->
      <mo>+</mo>
      <mn>1</mn>
   </mrow>
</math>
```

### 2.5 mscarries

这个标签可用于创建基本数学中出现的进位，借位和交叉， mscarries 的子元素与 mstack 的下一行中的元素相关联。 mscarries 的每个孩子除了 mscarry 之外或者 none 被视为被 mscarry 隐含地包围，举例说明：

```javascript
<math>
<mscarries> expression <mscarry> <none/> </mscarry> </mscarries>
</math>
```

以上是简单的写法。

### 2.6 menclose

这个标签用于呈现由其表示法属性指定的封闭符号内的内容。它接受一个参数作为多个子元素的推断。它包含一个属性 notation，可选项有：

* longdiv - 精算；
* phasorangle - 激进；
* updiagonalstrike - 盒子；
* downdiagonalstrike - 圆盒；
* verticalstrike - 圆；
* horizontalstrike - 左；
* northeastarrow - 右；
* madruwb - 顶部；
* text - 底部；

### 2.7 mfenced

这个标签是一种使用 fencing 运算符 （如花括号，括号和括号） 而不是使用 标签的快捷方法。包含一个属性 expression，可选值有：

* open 指定开始分隔符，默认是 “（” ;
* close 指定结束分隔符，默认是 “）”;
* separators 指定零个或多个分隔符序列，默认是 “,”

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mrow>
      <mo>(</mo>
      <mn>1</mn>
      <mo>)</mo>
   </mrow>
</math><!--这是不适用mfenced的方式-->


<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mfenced>
      <mn>1</mn>
   </mfenced>
</math><!--使用mfenced快捷方式-->
```

通过上述代码可以看得出，使用这个标签可以减少代码量。

### 2.8 mfrac

这个标签用于绘制分数，它的子元素必须是 2 个，不然的话会出现报错，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mfrac>
      <mi>y</mi>
      <mn>10</mn>
   </mfrac>
</math>
```

以上示例中的分母是 10，分子是 y。

### 2.9 mlongdiv

这个标签用于绘制长的分区，它的简单语法格式是：

（请使用 Firefox 运行下面代码）

```javascript
<mlongdiv> divisor dividend result expression </mlongdiv>
```

### 2.10 mtable

这个标签比较类似于 HTML 中的表格元素，结合使用 mtr ，mtd 可以绘制出表格，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
      <mtable>
         <mtr>
            <mtd><mn>1</mn></mtd>
            <mtd><mn>0</mn></mtd>
            <mtd><mn>0</mn></mtd>
         </mtr>
         <mtr>
            <mtd><mn>0</mn></mtd>
            <mtd><mn>1</mn></mtd>
            <mtd><mn>0</mn></mtd>
         </mtr>
      </mtable>
</math>
```

### 2.11 msgroup

用于分组，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mstack>
      <msgroup>
         <mn>123</mn>
         <msrow>
            <mo>&#xD7;</mo>
            <mn>321</mn>
         </msrow>
      </msgroup>
      <msline/>
   </mstack>
</math>
```

### 2.12 mover

这个标签用于绘制下标，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mover accent = "true">
     <mrow>
       <mi> 1 </mi>
       <mo> + </mo>
       <mi> 2 </mi>
       <mo> + </mo>
       <mi> 3 </mi>
     </mrow>
     <mo>&#x23DE;</mo>
   </mover>
</math>
```

### 2.13 mpadded

这个标签用于在其内容周围添加填充或额外空间，它可用于调整尺寸和定位，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mrow>
      <mi>x</mi>
      <mpadded lspace = "0.2em" voffset = "0.3ex">
         <mi>y</mi>
      </mpadded>
      <mi>z</mi>
   </mrow>
</math>
```

### 2.14 mphantom

这个标签用于渲染无形中保持相同的大小和其他维度，包括基线位置，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mfrac>
      <mrow>
         <mi> x </mi>
         <mo> + </mo>
         <mi> y </mi>
         <mo> + </mo>
         <mi> z </mi>
      </mrow>

      <mrow>
         <mi> x </mi>
         <mphantom>
            <mo> + </mo>
         </mphantom>

         <mphantom>
            <mi> y </mi>
         </mphantom>
         <mo> + </mo>
         <mi> z </mi>
      </mrow>
   </mfrac>
</math>
```

### 2.15 msqrt

这个元素用于构造平方根，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <msqrt>
      <mn>256</mn>
   </msqrt>
</math>
```

### 2.16 msub

这个元素用于绘制下标表达式，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <msub>
      <mi>y</mi>
      <mn>10</mn>
   </msub>
</math>
```

### 2.17 msubsup

这个元素用于将下标和上标附加到表达式，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mrow>
      <msubsup>
         <mo> &#x222B;</mo>
         <mn> 0 </mn>
         <mn> 1 </mn>
      </msubsup>

      <mrow>
         <msup>
            <mi> e</mi>
            <mi> x </mi>
         </msup>
         <mo> &#x2062;</mo>

         <mrow>
            <mi> d</mi>
            <mi> x </mi>
         </mrow>
      </mrow>
   </mrow>
</math>
```

### 2.18 msup

这个元素用于将上标绘制到表达式，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <msup>
      <mi>y</mi>
      <mn>2</mn>
   </msup>
</math>
```

### 2.19 munder

这个元素用于绘制下标，可用于在表达式中添加重音或者限制，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <munder accent="true">
     <mrow>
       <mi>a </mi>
       <mo> + </mo>
       <mi> b </mi>
       <mo> + </mo>
       <mi> c </mi>
     </mrow>
     <mo>&#x23F;</mo>
   </munder>
</math>
```

### 2.20 munderover

这个元素用于绘制下方和上方，它支持在表达式上和下同事添加重音或者限制，例如：

（请使用 Firefox 运行下面代码）

```javascript
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <mrow>
      <munderover>
         <mo> &#x222B;</mo>
         <mn> 0 </mn>
         <mi> &#x221E;</mi>
      </munderover>
   </mrow>
</math>
```

## 3. 兼容性

![](https://imgconvert.csdnimg.cn/aHR0cHM6Ly9pbWFnZS56aGFuZ3hpbnh1LmNvbS9pbWFnZS9ibG9nLzIwMTgxMC8yMDE4LTEwLTI0XzIyNDk0Ny5wbmc?x-oss-process=image/format,png)

通过上图可以看出，基本上兼容性最好的是 Firefox 和 Safari ，其他浏览器基本上不兼容。

## 4. 第三方工具

由于 mathml 牵涉到的专业技术门槛较高，对数学知识要求较高，一般情况下如果只是项目中偶尔使用的话可以使用第三方工具降低开发成本。

### 4.1 在线转换

这个[网站](https://webdemo.myscript.com/views/main/math.html)可以在线将数学公式转换成 mathml 代码

### 4.2 第三方库

通过调用第三方库，可以使用 mathml 语法在不支持的浏览器上进行兼容模拟，例如 mathml.js

（下面例子可以使用 chorme 浏览器试试了，平方根也是可以显示出来的）

```javascript
<script src="//fred-wang.github.io/mathml.css/mspace.js"></script>
<math xmlns = "http://www.w3.org/1998/Math/MathML">
   <msqrt>
      <mn>256</mn>
   </msqrt>
</math>
```

## 5. 总结

本章介绍了小众语法 mathml，由其发展历史以及适用场景引申开来，进而简单介绍了它的语法结构，最后说明了其兼容性问题以及兼容性解决方案，由于目前只有极少数浏览器支持，所以在需要使用时需要先进行兼容性判断。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

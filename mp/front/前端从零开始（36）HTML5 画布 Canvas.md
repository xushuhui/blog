# HTML5 画布 Canvas

本章介绍 HTML 中用来绘图的元素画布。它是 HTML5 中新增的元素，通过使用 JavaScript 调用画布的函数可以控制画布中的每个像素，用来生成图形、字符或者图像。画布元素本身没有绘图功能，初始化定义的画布没有任何视觉效果，必须通过 JavaScript 拿到画布的 id，然后控制画布的绘制功能。所以想要使用画布，必须对 JavaScript 有一定的了解。画布牵涉到很多知识点，本章介绍简单的画布创建以及几种简单的基础形状绘制。

## 1. 画布的历史

画布元素最早是 Safari 浏览器在 1.3 版本引入的，为了解决在 dashboard 组件中支持脚本控制的图形，之后 Firefox1.5 和 Opera9 先后支持了画布元素，目前画布已经是 HTML5 中正式的标签元素了。

## 2. 画布基础操作

### 2.1 创建画布

通过声明 Canvas 标签可以创建一个画布元素，Canvas 支持高度、宽度属性。

```javascript
<!--如果当前你的浏览器不支持 canvas 元素，则显示 canvas 标签内的文字。如果支持什么都不会显示出来-->
<canvas id="test" width="500" height="400" style='border:1px solid #ccc'>您的浏览器不支持canvas</canvas>
```

代码说明：如果当前你的浏览器不支持 Canvas 元素，则显示 Canvas 标签内的文字。

JavaScript 可以通过 Canvas 定义的 id 来寻找 Canvas 元素，进而操控它绘图。

```javascript
var a = document.getElementById("test"); //根据id调用Canvas
```

### 2.2 坐标系

画布左上角 (0,0) 默认原点，x 坐标向右方增长，y 坐标则向下方延伸：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f0d6464095d20b202280225.jpg)

但是，Canvas 的坐标体系并不是一成不变的，原点是可改变的。坐标变换：可以对 Canvas 坐标系统进行移动 translate、旋转 rotate 和缩放等操作。坐标变换之后绘制的图形 x,y 偏移量都以新原点为准， 旋转角度，缩放比，以新坐标系角度为准。

### 2.3 获取画布 SDK 函数

声明完画布之后，画布标签本身除了高度和宽度之外基本上不再包含其他可以用于绘图的属性，所以想要操控画布必须获取到它提供的绘图 SDK 对象。

```javascript
var context = a.getContext(contextID)
```

通过 getContext 函数可以获取画布的 SDK 对象，在 HTML 中它被称为 CanvasRenderingContext2D 对象。CanvasRenderingContext2D 提供了一系列用于绘图的函数，其中包含以下几大类。

* 颜色、样式、阴影
* 线条样式
* 矩形
* 路径
* 转换
* 文本
* 图像绘制
* 像素操作
* 合成
* 其他

### 2.4 创建一个矩形

通过函数 fillRectangle 可以创建一个矩形，使用 fillStyle 属性为矩形填充颜色。

```javascript
<canvas id="test" width="500" height="400">您的浏览器不支持canvas</canvas>
<script>
var a = document.getElementById("test");
var ctx = a.getContext("2d");
ctx.fillStyle = "#eee"; //填充颜色使用rgb方式
ctx.fillRect(0,0,250,175); //定义矩形使用坐标点方式
</script>
```

### 2.5 绘制线条

使用 moveTo 函数定义线的开始坐标，lineTo 函数定义线的结束坐标，stroke 函数进行最终的绘制操作。

```javascript
<canvas id="test" width="500" height="400">您的浏览器不支持canvas</canvas>
<script>
var a = document.getElementById("test");
var ctx = a.getContext("2d");
ctx.moveTo(10, 20); //开始坐标在 (10,20)
ctx.lineTo(20, 100); //线条移动到 (20,100)
ctx.lineTo(70, 100); //线条移动到 (70,100)
ctx.strokeStyle = "grey"; //线条颜色设置为灰色
ctx.stroke(); //绘制
</script>
```

### 2.6 绘制圆形

使用 arc 可以画出一个圆形。

```javascript
<canvas id="test" width="500" height="400">您的浏览器不支持canvas</canvas>
<script>
var a = document.getElementById("test");
var ctx = a.getContext("2d");
ctx.arc(95,70,60,0,2*Math.PI); //圆心坐标是(95,70) 半径是60
ctx.stroke();
</script>
```

### 2.7 绘制文字

使用 strokeText 绘制文字。

```javascript
<!DOCTYPE html>
<html>
    <head>
        <title>A Simple Canvas Example</title>
        <style>
            body{
                background: #dddddd;
            }
            #canvas{
                margin: 10px;
	            padding: 10px;
                background: #ffffff;
                border: thin inset #aaaaaa;
            }
        </style>
    </head>
    <body>
        <canvas id="canvas" width="600" height="600">
            Canvas not supported
        </canvas>
        <script>
        var canvas=document.getElementById('canvas'),
		context=canvas.getContext('2d');
		context.font='38pt Arial';
		context.fillStyle='cornflowerblue';
		context.strokeStyle='blue';
		context.fillText('Hello imooc',canvas.width/2-150,canvas.height/2+15);
		context.strokeText('Hello imooc',canvas.width/2-150,canvas.height/2+15);
		</script>
    </body>
</html>
```

### 2.8 绘制渐变

使用 createLinearGradient 方法可以绘制线性的渐变，适用于矩形、圆形、线条、文本等。

```javascript
<canvas id="test" width="200" height="200">您的浏览器不支持canvas</canvas>
<script>
var a = document.getElementById("test");
var ctx = a.getContext("2d");
var Gradient = ctx.createLinearGradient(0,0,270,0); //圆心坐标是(95,70) 半径是60
Gradient.addColorStop(0,"red");
Gradient.addColorStop(1,"black");
ctx.fillStyle = Gradient;
ctx.fillRect(20,20,150,100);
</script>
```

绘制渐变对象，必须使用两种或两种以上的颜色。停止颜色，使用 addColorStop 方法指定颜色停止，参数为 0 - 1

### 2.9 绘制阴影

使用 shadow 系列函数可以绘制阴影，shadowBlur 表示阴影效果如何延伸 double 值。浏览器在阴影运用高斯模糊时，将会用到该值，它与像素无关，只会被用到高斯模糊方程之中，其默认值为 0。shadowColor 定义颜色值，默认值是 rgba(0,0,0,0)。shadowOffsetX 定义阴影在 X 轴方向的偏移量，以像素为单位，默认值为 0，shadowOffsetY 定义阴影在 Y 轴方向的偏移量，以像素为单位，默认值是 0。

```javascript
<canvas id="test" width="200" height="200">您的浏览器不支持canvas</canvas>
<script>
var canvas_2=document.getElementById("test");
var can2_context=canvas_2.getContext("2d");
var SHADOW_COLOR='rgba(0,0,0,0.7)'
can2_context.shadowColor=SHADOW_COLOR;
can2_context.shadowOffsetX=3;
can2_context.shadowOffsetY=3;
can2_context.shadowBlur=5
can2_context.fillStyle="red"
can2_context.fillRect(0,0,100,100)
</script>
```

### 2.10 纹理填充

填充纹理原理上是指图案的重复，通过 createPattern() 函数进行初始化。有两个参数 ，第一个是 Image 实例，第二个是形状中如何显示 repeat 图案。可以使用这个函数加载图像或者整个画布作为形状的填充图案。

```javascript
<canvas id="test" width="500" height="500">您的浏览器不支持canvas</canvas>
<script>
	var canvas = document.getElementById("test");
	canvas.width = 800;
	canvas.height = 600;
	var context = canvas.getContext("2d");

	var img = new Image();
	img.src = "https://www.easyicon.net/api/resizeApi.php?id=1183257&size=16";
	img.onload = function(){
	    var pattern = context.createPattern(img, "repeat");
	    context.fillStyle = pattern;
	    context.fillRect(0,0,500,400);
	}
</script>
```

## 3. 画布实战 - 五子棋小游戏

```javascript
<!DOCTYPE html>
<html>
<head>
<title>五子棋小游戏</title>
<meta charset="UTF-8">
</head>
<body>
<canvas id="canvas" width="600" height="600" onclick="exec(event)"  ></canvas>
<button onclick="reStart();">重新开始</button>
<button onclick="back();">悔棋</button>

</body>
<script type="text/javascript">
var c=document.getElementById("canvas");
var cxt=c.getContext("2d");
var data = [];//保存下棋的位置点
var clickCount = 0;//点击的次数
var canvasWidth = 600;//画布大小
var interval = 20;//棋盘间隔
var isEnd = false;//判断是否结束
var colorW = '#DAA520';
var colorH = '#000';
init();
function init() { //初始化棋盘
    for (var i = 0; i < canvasWidth;) {
        cxt.beginPath();
        cxt.lineWidth="1";
        cxt.strokeStyle="#8B4513";
        cxt.moveTo(i,0);
        cxt.lineTo(i,canvasWidth);
        cxt.stroke();
        cxt.beginPath();
        cxt.lineWidth="1";
        cxt.moveTo(0,i);
        cxt.lineTo(canvasWidth,i);
        cxt.stroke();
        i = i+interval;
    }
}
function exec(e) //执行下棋
{
	if(isEnd) return;

	var x1=e.clientX;
	var y1=e.clientY;
	var newX,newY;
	for (var i = 0; i < canvasWidth;) {
		if (x1>=i&&x1<i+interval/2) newX = i;
		if (x1>=i+interval/2&&x1<i+interval) newX = i+interval;
		if (y1>=i&&y1<i+interval/2) newY = i;
		if (y1>=i+interval/2&&y1<i+interval) newY = i+interval;
		i = i+interval;
	}//计算落棋位置

	if (!checkDataExists(newX,newY)) {//判断该点是否已经有棋子存在
		var isTrue = true;
		if (clickCount%2==0) {
			cxt.fillStyle=colorW;
		}else{
			cxt.fillStyle=colorH;
			isTrue = false;
		}
		cxt.beginPath();
		cxt.arc(newX,newY,interval/2,0,Math.PI*2,true);
		cxt.closePath();
		cxt.fill();
		data.push({'x':newX,'y':newY,'isTrue':isTrue});//绘制棋子
		clickCount++;
		if(isFinish(newX,newY,isTrue)){//判断是否已经结束
			isEnd = true;
			if (isTrue) alert('黄棋赢了');
			else alert('黑棋赢了');
		}
	}else{
		 alert("当前点已经存在");
	}
}
function reStart() {//比赛重新开始
    cxt.clearRect(0,0,canvasWidth,canvasWidth);
    init();
    data = [];
    clickCount=0;
    isEnd = false;
}
function back() {//执行悔棋
    cxt.clearRect(0,0,canvasWidth,canvasWidth);
    init();
    clickCount--;
    data.pop();
    isEnd = false;
    for (var i = 0; i < clickCount; i++) {
        cxt.beginPath();
        cxt.fillStyle = i%2==0 ? colorW:colorH;
        cxt.arc(data[i].x,data[i].y,10,0,Math.PI*2,true);
        cxt.closePath();
        cxt.fill();
    }
}

function checkDataExists(x,y,isTrue){//判断当前落棋点是否已经存在棋子
    for (var i = 0; i < data.length; i++) {
        if (data[i].x ==x && data[i].y == y && (typeof(isTrue) == "undefined" || data[i].isTrue == isTrue))  return true;
    }
    return false;
}
function isFinish(x1,y1,isTrue) {//判断是否结束棋局
    x2 = x3 = x4 = x5 = x1;
    y2 = y3 = y4 = y5 = y1;

    x2 = x1>=5*interval ? x1-5*interval : 0;
    lineCount = 0;
    for (var i = 0; i < 10; i++) {
        tempx = x2+interval*i;
        if (checkDataExists(tempx,y2,isTrue)) {
            lineCount++;
            if (lineCount==5) break;
        }else  lineCount=0;
    }
    if (lineCount>=5)   return true;

    if (y1>=5*interval)   y3 = y1-5*interval;
    else  y3=0;
    lineCount = 0;
    for (var i = 0; i < 10; i++) {
        tempy = y3+interval*i;
        if (checkDataExists(x3,tempy,isTrue)) {
            lineCount++;
            if (lineCount==5)  break;
        }else lineCount=0;
    }
    if (lineCount>=5) return true;

    x4 = x1-5*interval;
    y4 = y1-5*interval;
    lineCount = 0;
    for (var i = 0; i < 10; i++) {
        tempy = y4+interval*i;
        tempx = x4+interval*i;
        if (checkDataExists(tempx,tempy,isTrue)) {
            lineCount++;
            if (lineCount==5) break;
        }else  lineCount=0;
    }
    if (lineCount>=5) return true;

    x5 = x1-5*interval;
    y5 = y1+5*interval;
    lineCount = 0;
    for (var i = 0; i < 10; i++) {
        tempy = y5-interval*i;
        tempx = x5+interval*i;
        if (checkDataExists(tempx,tempy,isTrue)) {
            lineCount++;
            if (lineCount==5)  break;
        }else  lineCount=0;
    }
    if (lineCount>=5) return true;
}
</script>
</html>
```

上述代码使用 HTML 的画布功能实现了一个简单的五子棋功能，其中除了用到画布还使用到了一些简单的数据结构和算法，比如判断棋局是否结束等。

## 4. 小结

本章介绍了 HTM5 中新增的绘图工具 Canvas，Canvas 的历史，以及通过几种简单的实操方式介绍了如何实际使用画布。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

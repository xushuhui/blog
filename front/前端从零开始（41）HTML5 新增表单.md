# 新增表单

之前的教程中已经简单的介绍过了表单，早期的网页中为了实现复杂的交互效果，通常需要使用 div+css 模拟复杂的表单类型实现类似日期、滑块条、颜色选择等效果。HTML5 标准中考虑到这种情况，增加了不少的复杂表单效果。本章主要介绍 HTML5 新增的几种增强的表单类型。

## 1. email

此类型的表单跟普通 text 类型的表单类型表现方式一致，只是在输入完成之后如果不符合 email 类型浏览器会有提示，且不允许提交，定义方式如下：

```javascript
<input type=email>
```

以下实例使用 email 类型的表单实现了一个简单的注册功能：

```javascript
<script>
function beforeSend(){
	if(document.getElementById("password").value != document.getElementById("password1").value) return alert("请输入正确的密码");
}
</script>
<form method="post" action="/redister.html" onsubmit='beforeSend()'>
	<p><label for="loginName">登录名：</label><input name='loginName' type='text' pattern=".{3,20}">(3~20个字符)</p>
	<p><label for="password">登录密码：</label><input id='password' name='password' type='password' pattern=".{6,20}">(至少6位)</p>
	<p><label for="password1">重复密码：</label><input id='password1' name='password1' type='password' pattern=".{6,20}">(至少6位)</p>
	<p><label for="email">邮箱：</label><input name='email' type='email'>(请输入正确的邮箱地址)</p>
	<input type='reset' value='重置'><button type=submit>注册</button>
</form>
```

## 2. url

url 类型的表单视觉展现跟 text 类型的一致，只是在输入完成之后如果不符合 URL 类型浏览器会有提示，且不允许提交，语法如下：

```javascript
<input type='url'>
```

以下示例展示了 url 表单的实际使用场景：

```javascript
<form>
<label>姓名：</label><input type=text name='name'  pattern=".{3,20}"><!--pattern属性用于约束输入值-->
<label>电话：</label><input type=tel name='phone' ><!--使用tel类型表单-->
<label>住址：</label><input type=text name='address'  pattern=".{5,50}">
<label>个人主页：</label><input type=url name='url'><!--使用url类型表单-->
<input type='reset' value='重置'>&nbsp;&nbsp;&nbsp; <button type=submit>提交</button>
</form>
```

## 3. number

number 类型的表单也跟 text 表现形式一致，但是浏览器会强制不能输入非数字类型的字符，表单最后侧默认会有上下两个按钮，语法如下：

```javascript
<input type=number>
```

## 4. tel

tel 类型要求输入一个电话号码，但实际上它并没有特殊的验证，与 text 类型没什么区别：

```javascript
<input type=tel>
```

## 5. range

此类型将显示一个可拖动的滑块条，并可通过设定 `max/min/step` 值限定拖动范围。拖动时会反馈给 value 一个值。

```javascript
<input type=range min=20 max=100 step=2 >
```

在实际项目中可以根据动态获取滑块的 value 值，来实现一定的效果。以下展示了一个使用 range 表单实现了一个动态缩放图片的功能：

```javascript
<!DOCTYPE html>
<html>
<head>
	<title>使用滑动条缩放图片</title>
</head>
<body style="">
	<canvas id="canvas" style="display: block;margin: 0 auto;border: 1px solid #aaa;background:black">
		你的浏览器不支持canvas。
	</canvas>
	<input type="range" id="range" min="0.5" max="5.0" step="0.01" value="1.0" style="display: block;margin: 20px auto;width: 800px;"/><!--定义滑动条-->
</body>

	<script>
		var canvas = document.getElementById("canvas");
		var context = canvas.getContext("2d");
		var slider = document.getElementById("range");
		var image = new Image();
		window.onload = function(){//浏览器加载完成之后触发
			canvas.width = 600;
			canvas.height = 400;

		image.src="https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"; //加载图片

			image.onload = function() {
				slider.onmousemove = function(){//通过监听鼠标事件动态获取滑块的value值
					var scale = slider.value;
					var width = canvas.width * scale;
					var height = canvas.height * scale;
					var dx = canvas.width/2 - width/2;
					var dy = canvas.height/2 - height/2;
					context.clearRect(0, 0, canvas.width, canvas.height);
					context.drawImage(image, dx, dy, width, height);//设置图片的缩放度
				};
			};

		};
	</script>
</html>
```

## 6. color

此类型表单，可让用户通过颜色选择器选择一个颜色值，并反馈到 value 中，可以设置默认值，语法如下：

```javascript
<input type=color>
```

实际项目中，一般用来作为为画笔或者绘图选择颜色，以下示例展示了一个简单的颜色选择器表单：

```javascript
<!DOCTYPE html>
<html>
<head>
<meta charset=" utf-8">
<title>颜色选择器</title>
</head>
<body>
<form name="test" id="test" method="post" action="test.php">
选择颜色:<input type="color" form="ant" name="color"/>
<input type="submit" value="提交">
</form>

</body>
</html>

```

## 7. 时间日期系列

这个类型的表单包含几种类型，用来实现繁琐的日历控件，效果各有不同，语法如下：

```javascript
<input type=date ><!-- 日期 -->
<input type=time ><!-- 时间 -->
<input type=datetime ><!-- 日期+时间  （已经废弃）-->
<input type=datetime-local ><!-- 日期+时间 -->
<input type=month ><!-- 月份 -->
<input type=week ><!-- 星期 -->
```

可以运行下面代码，试试效果：

```javascript
<p>日期：<input type=date ></p> <!-- 日期 -->
<p>时间：<input type=time ></p>  <!-- 时间 -->
<p>日期+时间(已经废弃):<input type=datetime ></p> <!-- 日期+时间  （已经废弃）-->
<p>日期+时间:<input type=datetime-local ></p> <!-- 日期+时间 -->
<p>月份:<input type=month ></p> <!-- 月份 -->
<p>星期:<input type=week ></p> <!-- 星期 -->
```

## 8. search

此类型表示输入的将是一个搜索关键字，通过 `results=s` 或者 `x-webkit-speech` 可显示一个搜索小图标。语法如下：

```javascript
<input type=search results=s >
```

这个表单在实际项目中适用场景较少，所以没有示例可以参考。

## 9. 小结

回顾本章介绍了几种实时交互效果较强的表单控件及用法，弥补了早期 HTML 中的交互缺失的情况。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

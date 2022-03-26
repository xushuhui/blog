# HTML5 地理位置

地理定位功能是 HTML5 新增的标准，早期的 HTML 和 JavaScript 没有操控硬件和文件的权限，因为页面交互效果比较简单；但是 HTML5 之后网页已经逐渐应用于各种复杂场景包括移动设备，所以增加了各种与硬件交互的 API 接口，地理位置就是其中之一。

## 1. 早期的方式

在 HTML5 之前，获取地理位置的解决方法是在已知 IP 位置的数据库中查找访问者的 IP 地址，然后根据 IP 数据库查找到对应的位置。虽然这种方法的准确度远低于使用 GPS 设备的准确度，但是这些数据库通常能够定位到访客大致的地址范围，这对于许多应用来说是足够有用的，但是数据库需要经常维护更新。

## 2. 检测是否支持地理位置

并非所有的浏览器或者硬件设备都支持地理位置功能，所以使用之前需要进行容错判断：

```javascript

if (navigator.geolocation)    {//判断地理位置是否支持
    //业务代码
}
else{
    x.innerHTML="该浏览器不支持获取地理位置。";
}
```

获取地理位置之前需要用户点击同意按钮，因为该功能牵涉到隐私。

## 3. 地理位置 API

### 3.1 获取当前位置

使用 getCurrentPosition 函数获取用户当前的地理位置，这个函数有 3 个参数：

* 第一个参数设置成功获取的回调函数；
* 第二个参数设置失败之后的回调函数；
* 第三个参数设置一些可选参数项。

例如：

```javascript
navigator.geolocation.getCurrentPosition(function(position) {
    //TODO 成功时的处理
    var timestamp = position.timestamp;
    var coords      = position.coords;
}, function(error) {
    //TODO 失败时的处理
    console.log(error);
}, {
   //参数设置
})
```

成功获取之后的回调函数中通过参数传递的方式可以拿到地理位置的对象，它是一个

Geoposition 对象，上述示例使用 position 变量表示，这个对象包含 2 个属性：

* timestamp 时间戳
* coords 一个 coordinates 类型对象，包括    * accuracy 精度值
    * altitude 海拔
    * altitudeAccuracy 海拔的精度
    * heading 设备前进方向
    * latitude 经度
    * longitude 纬度
    * speed 前进速度

    第三个参数是一个 PositionOptions 对象，它包含 3 个用于设置的属性：

* enableHighAccuracy 是否使用最高精度表示结果
* timeout 设置超时时间
* maximumAge 表示获取多久的缓存位置

### 3.2 监视位置

使用 watchPosition 函数可以定时获取用户地理位置信息，在用户设备的地理位置发生改变的时候自动被调用。这个函数跟 getCurrentPosition 函数的使用方式基本一致。

```javascript
navigator.getlocation.watchPosition(function(pos){
    //业务代码
},function(err){
},
{}
)
```

### 3.3 清除监视

使用 clearWatch 函数删除 watchPosition 函数注册的监听器：

```javascript
var watch = navigator.geolocation.watchPosition(show_map, handle_error, {enableHighAccuracy: true,timeoout: 175000, maximumAge: 75000})
clearWatch(watch); //清除监视
```

## 4. 定位失败

由于获取地理位置功能依赖硬件信号，例如 GPS 信号、WiFi 信号等等，所以有时可能会出现获取不到位置的情况，在这里做了一下总结：

### 4.1 浏览器不支持原生定位接口

有些旧版本的浏览器不支持 HTML5，如 IE 较低版本的浏览器。这时调用定位接口会出现 error 信息，message 字段包含 Browser not Support html5 geolocation 信息。

### 4.2 用户禁用了定位权限

需要用户开启定位权限，error 信息的 message 字段包含 Geolocation permission denied。

### 4.3 浏览器禁止了非安全域的定位请求

比如 Chrome、IOS 10 已经陆续禁止，需要升级站点到 HTTPS，error 信息的 message 字段包含 Geolocation permission denied 信息。**注意：Chrome 不会禁止 localhost 域名 HTTP 协议下的定位**

### 4.4 定位超时

由于信号问题有时会出现超时问题，可以适当增加超时属性的设定值以减少这一现象。某个别浏览器本身对定位接口的友好程度较弱，也会超时返回失败，error 信息的 message 字段包含 Geolocation time out 信息。

### 4.5 定位服务问题

Chrome、Firefox 以及一些套壳浏览器接入的定位服务在国外，有较大的限制，也会造成定位失败，且失败率较高。

## 5. 总结

本章介绍了移动开发的利器 - 地理位置功能，通过这个功能可以使用 HTML 直接跟移动设备硬件交互，很大程度上从丰富了网页的交互方式，不过需要用户授权之后才能使用；地理位置相关的函数只有 3 个，使用时需要考虑浏览器兼容性

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

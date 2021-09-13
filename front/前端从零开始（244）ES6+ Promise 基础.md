# ES6+ Promise 基础

## 1. 前言

我们知道浏览器在渲染网页时，会创建一个渲染进程进行渲染页面，在渲染进程中其中有 GUI 渲染线程和 JS 引擎线程（如 V8 引擎）两个线程是互斥的。也就是说在同一时间内只能有一个线程执行。如果 JavaScript 执行一段耗时程序时会阻止页面渲染。如果要页面快速在用户面前呈现就要做一些优化处理。对于不能立马得到结果的程序，不需要等待，可以放到事件队列中，等到得到结果后再执行。

对这种不等待方式，JavaScript 提供了异步的解决方案，在 JavaScript 中常见的异步解决方案是 Callback 方式，而像 setTimeout 这样提供异步的 API，还可以使用发布订阅的来实现异步。使用回调函数存在回调地狱的问题。为了解决回调地狱，最早社区提出了 Promise 概念。最后在 ES6 时正式作为官方的解决方案，说明 Promise 有它独有的优势。本节我们将学习 Promise 的基本用法。

## 2. 回调地狱

我们都知道 JavaScript 异步使用的是回调函数，下面我们来看一个 ajax 请求的实例，下面的 ajax 方法是一个伪代码，可以看作是请求接口的方法，接口请求的库可以参考 jQuery 的 $.ajax 和 axios。

```javascript
// ajax请求的伪代码
function ajax(url, sucessCallback, failCallback) {
  // url：请求的url
  // sucessCallback：成功的回调函数
  // failCallback：失败的回调函数
}

ajax(url1, (res1) => {
  ajax(url2, (res2) => {
    ajax(url3, (res3) => {
		doSomething(res1, res2, res3)
    })
  })
})
```

上面的 ajax 请求我们可以理解为，在调用 doSomething 方法时需要前面三个请求的结果作为参数，所以只有前一个 ajax 请求得到结果后才能发起第二个请求。这样前后有依赖的嵌套被称为回调地狱。对于比较复杂逻辑的情况来说，回调地狱会使程序出问题的概率大大增加。

另外，这样做有个很严重的问题，就是接口请求的时间是三个请求的和，不能进行并发操作，当然我们也可以做一些优化操作，如下：

```javascript
let out = after(3, function (data){
  doSomething(...data)
})

ajax(url1, (res1) => {
  out(res1)
})
ajax(url2, (res2) => {
  out(res2)
})
ajax(url3, (res3) => {
  out(res3)
})

function after(times, callback) {
  const arr = [];
  return function (value){
    arr.push(value);
    if (--times==0) {
      callback(arr);
    }
  }
}
```

上面的代码很优雅地解决了回调嵌套的问题，但同时我们需要手动维护一个计数器来控制最后的回调。这无疑增加了程序的复杂度，我们更希望的是关注我的业务，而不是写更多的逻辑来优化。

针对这种情况，社区提供了很多这类优化的库，而 Promise 则是其中最亮眼的。对上面的情况，Promise 怎么解决的呢？看如下的实现方式：

```javascript
function request(url) {
  return new Promise((resolve, reject) => {
    ajax(url, (res) => {
      resolve(res)
    })
  })
}
Promise.all([request(url1), request(url1), request(url1)]).then((result) => {
  doSomething(...result)
}).catch((error) => {
  console.log(error)
})
```

上面的代码中我们封装了一个 request 请求的方法，通过 `Promise.all()` 来并发请求这些接口，当接口都正确返回才会执行 then 方法中的回调，有一个错误都会抛出异常。这种方式比较好的是，我们对请求进行了封装，不要再关注每一步请求是否完成做对应的逻辑处理，让我们在开发过程中更加关注业务逻辑，使开发效率更快。

## 3. Promise 用法

前面我们通过一个回调地狱的案例，说明了 Promise 的优点，就是为了解决异步而产生的。并且可以处理并发请求，很好地优化了程序资源。

### 3.1 实例化一个 Promise

首先需要明确 Promise 是一个类，我们在 VSCode 中输入 `new Promise()` 会给我们如下的提示：

![图片描述](http://img.mukewang.com/wiki/5f834b9e0971024c10660378.jpg)

在 `new Promise()` 时需要默认需要传入一个回调函数，这个回调函数是 executor（执行器），默认会立即执行。执行器会提供两个方法（resolve 和 reject）用于改变 promise 的状态。`resolve` 会触发成功状态，`reject` 会触发失败状态，无论成功或失败都会传入一个返回值，这个返回值会在实例调用 `then` 方法后作为响应值获取。

```javascript
var promise = new Promise((resolve, reject) => {
  ajax(url, (data) => {
    resolve(data)	// 成功
  }, (error) => {
    reject(error)	// 失败
  })
})
```

上面的代码中实例化一个 ajax 请求的 Promise， 当接口请求成功就会调用 resolve () 方法把请求的值传入进去，如果失败了就调用 reject () 方法把错误信息传入进去。在后续的链式调用中获取相应的结果。

我们需要知道的是，Promise 有三个状态：等待（padding）、成功（fulfilled），失败（rejected）。在初始化时，这个状态是等待态，在等待状态时可以转化为成功态或失败态。当状态是成功态或是失败态后不能再被改变了。

上面的代码中可以改变 Promise 状态的是执行器提供的 resolve 和 reject，resolve 会将等待态变为成功态，reject 则会将等待态变为失败态，在状态变为成功或失败的状态时就不能被更改了。

### 3.2 then

在实例化 Promise 类后，我们如何访问成功和失败的值呢？Promise 提供了链式调用的 then 方法用于访问成功的值和失败的值，then 提供了两个回调 onfulfilled（成功回调）、onrejected（失败回调）。

```javascript
var promise = new Promise((resolve, reject) => {
  resolve(123);
  // reject('error')
  // throw new Error('Error')
})

promise
	.then(
    (data) => {
      console.log(data)	// 123
      return '100'
    },
    (reason) => {
      console.log(reason)
    }
  )
  .then((data) => {
      console.log(data)	// 100
  }, null)
```

上面的代码中给我了几个测试用例，有兴趣的小伙伴可以进行测试。then 方法返回一个值而不是 Promise 实例，并且会把这个结果返回到下一个 then 的成功回调中；

如果返回的是一个 promise，下一个 then 会采用这个 promise 结果，如果返回的是失败，会传到下一个 then 中失败的回调函数中去：

```javascript
var promise = new Promise((resolve, reject) => {
  resolve(123);
})

promise
	.then(
    (data) => {
      return new Promise((resolve, reject) => {
        reject('错误内容');
      })
    },
    null
  )
  .then(null, (err) => {
      console.log('error:', err)	// error: 错误内容
  })
```

如果在失败的回调函数中返回一个普通值或成功的 promise 也会走到下一层 then 的成功回调中去。

```javascript
promise.then(null, (err) => {
	return '100';
}).then((data) => {
  console.log('data:', data);	// data: 123
}, null)
```

通过上面的例子可以知道，当前 then 中走成功与否，主要看上一层返回的结果。总结有两点。

* 当上一层返回一个普通值，或是一个成功的 Promise，则会走到下一层成功的回调函数中去；
* 如果上一层返回一个失败的 Promise，或是抛出一个异常，则会走到下一层的失败的回调函数中去。

## 4. 小结

本节主要通过 JavaScript 中回调地狱的一个案例来引出为什么使用 Promise，以及 Promise 所带来的好处。然后学习了 Promise 的基本使用和链式调用 then 方法，需要注意的是，then 中执行成功或是失败是根据它上一层的返回值，如果返回的是一个普通值或成功的 Promise 则会走 then 的成功回调；如果抛出异常或返回失败的 Promise 则走 then 的失败回调。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

# ES6+ Generator 函数应用

## 1. 前言

上一节我们注意学习了生成器的概念和基本用法，并通过两个案例来说明。但是生成器更加广泛和设计之初是为了解决异步而产生的。我们会通过一个开发中常见的问题入手来看 生成器函数到底是怎么来解决异步调用的问题，并且会实现一个简版的 co 库。

## 2. 案例

在开发过程中会遇到一个很常见的需求，我们想获取一个值，但不能直接拿到，我们只能先请求一个接口如：api_1，获取这个值的接口地址如：api_2。然后，请求 api_2 接口才能获取这个值。这是一个典型的需要异步回调才能完成的功能。

在学习 Promise 的时候我们也针对这样的情况，我们可以使用 Promise 来完成这样的功能：

```javascript
var promise = function (url) {
	return new Promise((resolve, reject) => {
    ajax(url, (data) => {
      resolve(data)	// 成功
    }, (error) => {
      reject(error)	// 失败
    })
  })
}

promise('api_1').then(res1 => {
  promise(res1).then(res2 => {
    console.log(res2)
  })
})
```

从上面的代码中可以看出，在这种情况下，使用 Promise 好像并没有解决回调地狱的问题。那如何解决这种问题呢？我们想到了 Generator 函数具有暂停功能，那是不是我们可以让请求 api_2 接口时暂停，等到 api_1 请求成功获取到地址后再去请求呢？按照这个思路我们可以有下面的代码：

```javascript
const ajax = function(api) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (api === 'api_1') {
        resolve('api_2');
      }
      if (api === 'api_2') {
        resolve(100);
      }
    }, 0)
  })
}

function * getValue() {
  const api = yield ajax('api_1');
  const value = yield ajax(api);
  return value;
}

console.log(getValue());	// Object [Generator] {}
```

上面的代码是我们模拟 ajax 请求，通过使用生成器函数写出的代码让我们感觉有了同步的感觉，但是这样去执行 getValue 函数是不会得到结果的。那么我们要怎样去获得结果呢？根据生成器函数的特点，可以这样写：

```javascript
let it = getValue();

let { value } = it.next();
value.then((data) => {
  let { value } = it.next(data);
  value.then((data) => {
    Console.log(data);
  });
});
```

从上面的代码中看出还是有嵌套，好像并没有解决问题。但如果你细心，你会发现每个回调的逻辑基本都是一样的。那么我们能不能对这样的嵌套函数进行封装呢？答案当然是可以的，有个库就专门解决了这个痛点 —— [co](https://github.com/tj/co) 库，有兴趣的可以去研究一下这个库，代码很少，下面我们就来封装一个这样的库。

先让我们看看 co 库是怎么使用的：

```javascript
co(getValue()).then(res => {
  console.log(res);
})
```

从上面的代码中可以看出，把函数传入进去，并让函数执行，然后在 then 的成功回调中可以获取 `getValue` 函数返回的最终结果。这样非常清晰地解决了上面我们需要手动执行的方法，下面我来分析一下具体的实现步骤：

1. 从上面的用法可以看出，co 返回的是一个 Promise 实例，所以我们需要返回一个 `new Promise()` 实例；
2. 传入的生成器函数执行后，我们可以调用 next () 函数拿到返回的值和是否执行完的状态，判断 done 如果是 true 时说明执行完了，可以执行 resolve；
3. 当生成器函数没有执行完时，这时我们就需要递归地去调用这个 next () 来执行下一步，因为传入的值是一个 Promise 实例，要想获取它的结果就需要链式调用 then 方法，然后拿到结果进行递归执行。

有了上面的步骤分析，不难得到下面的代码：

```javascript
function co(it) {
  return new Promise((resolve, reject) => {
    function next(data) {
      let { value, done } = it.next(data);
      if (done) {
        resolve(value);
      } else {
        Promise.resolve(value).then(data => {
          next(data);
        }, reject)
      }
    }
    next(undefined);
  })
}
```

上面的代码中需要注意的是，如果 yield 返回的不是一个 Promise 对象时，我们对 value 使用了 `Promise.resolve()` 进行了包装，这样就可以处理返回一个普通值时没有 then 方法的问题。

## 3. 小结

本节主要讲解了 Generator 函数在异步中的应用，解决了某些场景下还会产生回调地狱的问题，通过封装 co 方法让我们的代码写起来像是同步一样，但是 Generator 函数还不是我们解决异步的终极方案，下一节我们将学习 async 函数，看它是怎么来解决异步的。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

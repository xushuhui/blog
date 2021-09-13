# ES6+ 实现一个简版的 Promise

## 1. 前言

上一节我们学习了 ES6 Promise的基本用法，并且我们知道 Promise 最早出现在社区，所以ES6 中 Promise 也是遵循一个标准的规范的。这个规范就是 [Promise A+ 规范](https://promisesaplus.com/) 也就是任何人都可以遵循这个规范实现一个自己的 Promise，由于每个人实现的方式有所差异，[Promise A+ 规范](https://promisesaplus.com/) 给出了一些要求和兼容方式。

本节我们将根据 [Promise A+ 规范](https://promisesaplus.com/) 实现一个简版的 Promise API。

## 2. 实现步骤

上一节我们已经知道了 Promise 是一个类，默认接收一个参数 executor（执行器），并且会立即执行。所以首先需要创建一个 Promise 的类，然后传入一个回调函数并执行它，故有如下的初始代码：

```javascript
class Promise {
  constructor(executor) {
    executor();
  }
}
```

Promise 有三个状态：等待（padding）、成功（fulfilled），失败（rejected）。默认是等待状态，等待态可以突变为成功态或失败态，所以我们可以定义三个常量来存放这三个状态

```javascript
const PENDING = 'PENDING';
const RESOLVED = 'RESOLVED';	// 成功态
const REJECTED = 'REJECTED';	// 失败态
class Promise {
  constructor(executor) {
    this.status = PENDING;	// 默认是等待态
    executor();
  }
}
```

这样我们就知道了 Promise 的基本状态，那内部的状态是怎么突变为成功或失败的呢？这里执行器（executor）会提供两个个方法用于改变 Promise 的状态，所以我们需要在初始化时定义 resolve 和 reject 方法：在成功的时候会传入成功的值，在失败的时候会传入失败的原因。并且每个 Promise 都会提供 then 方法用于链式调用。

```javascript
class Promise {
  constructor(executor) {
    this.status = PENDING;
    const resolve = (value) => {};
    const reject = (reason) => {};
    // 执行executor时，传入成功或失败的回调
    executor(resolve, reject);
  }
  then(onfulfilled, onrejected) {

  }
}
```

这时我们就可以开始着手去更改 Promise 的状态了，由于默认情况下 Promise 的状态只能从 pending 到 fulfilled 和 rejected 的转化。

```javascript
class Promise {
  constructor(executor) {
    this.status = PENDING;
    const resolve = (value) => {
      // 只有等待态时才能更改状态
      if (this.status === PENDING) {
        this.status = RESOLVED;
      }
    };
    const reject = (reason) => {
      if (this.status === PENDING) {
        this.status = REJECTED;
      }
    };
    executor(resolve, reject);
  }
  ...
}
```

成功和失败都会返回对应的结果，所以我们需要定义成功的值和失败的原因两个全局变量，用于存放返回的结果。

```javascript
class Promise {
  constructor(executor) {
    this.status = PENDING;
    this.value = undefined;
    this.reason = undefined;
    const resolve = (value) => {
      // 只有等待态时才能更改状态
      if (this.status === PENDING) {
        this.value = value;
        this.status = RESOLVED;
      }
    };
    const reject = (reason) => {
      if (this.status === PENDING) {
        this.reason = reason;
        this.status = REJECTED;
      }
    };
    executor(resolve, reject);
  }
  ...
}
```

这时我们就已经为执行器提供了两个回调函数了，如果在执行器执行时抛出异常时，我们需要使用 try…catch 来补货一下。由于是抛出异常，所以，需要调用 reject 方法来修改为失败的状态。

```javascript
try {
  executor(resolve, reject);
} catch(e) {
  reject(e)
}
```

我们知道实例在调用 then 方法时会传入两个回调函数 onfulfilled, onrejected 去执行成功或失败的回调，所以根据状态会调用对应的函数来处理。

```javascript
then(onfulfilled, onrejected) {
  if (this.status === RESOLVED) {
    onfulfilled(this.value)
  }
  if (this.status === REJECTED) {
    onrejected(this.reason)
  }
}
```

这样我们就完了 Promise 最基本的同步功能，

```javascript
let promise = new Promise((resolve, reject) => {
  resolve('value');
  // throw new Error('错误');
  // reject('error reason')
  // setTimeout(() => {
  //   resolve('value');
  // }, 1000)
})
promise.then((data) => {
  console.log('resolve response', data);
}, (err) => {
  console.log('reject response', err);
})
```

用上面的代码对我们写的 Promise 进行验证，通过测试用例可知，我们写的 Promise 只能在同步中运行，当我们使用 setTimeout 异步去返回时，并没有预想的在 then 的成功回调中打印结果。

对于这种异步行为需要专门处理，如何处理异步的内容呢？我们知道在执行异步任务时 Promise 的状态并没有被改变，也就是并没有执行 resolve 或 reject 方法，但是 then 中的回调已经执行了，这时就需要增加当 Promise 还是等待态的逻辑，在等待态时把回调函数都存放起来，等到执行 resolve 或 reject 再依次执行之前存放的 then 的回调函数，也就是我们平时用到的发布订阅模式。实现步骤：

* 首先，需要在初始化中增加存放成功的回调函数和存放失败的回调函数；
* 然后，由于是异步执行 resolve 或 reject 所以需要在 then 方法中把回调函数存放起来；
* 最后，当执行 resolve 或 reject 时取出存放的回调函数依次执行。

根据以上的实现步骤可以得到如下的代码：

```javascript
class Promise {
  constructor(executor) {
	this.status = PENDING;
    this.value = undefined; // 成功的值
    this.reason = undefined; // 失败的原因
+   // 存放成功的回调函数
+    this.onResolvedCallbacks = [];
+    // 存放失败的回调函数
+    this.onRejectedCallbacks = [];
    let resolve = (value) => {
      if (this.status === PENDING) {
		this.value = value;
        this.status = RESOLVED;
+       // 异步时，存放在成功的回调函数依次执行
+       this.onResolvedCallbacks.forEach(fn => fn())
      }
    };
    let reject = (reason) => {
      if (this.status === PENDING) {
		this.value = reason;
        this.status = REJECTED;
+       // 异步时，存放在失败的回调函数依次执行
+       this.onRejectedCallbacks.forEach(fn => fn())
      }
    };
    try {
      executor(resolve, reject);
    } catch(e) {
      reject(e)
    }
  }
  then(onfulfilled, onrejected) {
    if (this.status === RESOLVED) {
      onfulfilled(this.value)
    }
    if (this.status === REJECTED) {
      onrejected(this.reason)
    }
+    if (this.status === PENDING) {
+      this.onResolvedCallbacks.push(() => {
+        // TODO
+        onfulfilled(this.value);
+      })
+      this.onRejectedCallbacks.push(() => {
+        // TODO
+        onrejected(this.reason);
+      })
+    }
  }
}
```

上面的代码中，在存放回调函数时把 `onfulfilled`, `onrejected` 存放在一个函数中执行，这样的好处是可以在前面增加处理问题的逻辑。这样我们就完成了处理异步的 Promise 逻辑。下面是测试用例，可以正常的执行 then 的成功回调函数。

```javascript
let promise = new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve('100');
  }, 1000)
})
promise.then((data) => {
  console.log('resolve response:', data); // resolve response: 100
}, (err) => {
  console.log('reject response:', err);
})
```

到这里我们是不是已经基本实现了 Promise 的功能呢？ES6 中的 then 方法支持链式调用，那我们写的可以吗？我们在看下面的一个测试用例：

```javascript
let promise = new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve('100');
  }, 1000)
})
promise.then((data) => {
  console.log('resolve response:', data); // resolve response: 100
  return 200
}, (err) => {
  console.log('reject response:', err);
}).then((data) => {
  console.log('data2:', data)
}, null)
// TypeError: Cannot read property 'then' of undefined
```

然而当我们在执行的时候会报错，then 是 undefined。为什么会这样呢？那我们要知道如何满足链式调用的规范，那就是在完成任务后再返回一个 Promise 实例。那如何返回一个 Promise 实例呢？在 Promise A+ 规范的 2.2.7 小节在有详细的描述，再实例化一个 promise2 来存放执行后的结果，并返回 promise2。那么我们就要改造 then 方法了。

```javascript
class Promise {
  ...
  then(onfulfilled, onrejected) {
	let promise2 = new Promise((resolve, reject) => {
      if (this.status === RESOLVED) {
        const x = onfulfilled(this.value)
		resolve(x)
      }
      if (this.status === REJECTED) {
        const x = onrejected(this.reason);
        reject(x)
      }
      if (this.status === PENDING) {
        this.onResolvedCallbacks.push(() => {
	        const x = onfulfilled(this.value)
			resolve(x)
        })
        this.onRejectedCallbacks.push(() => {
	        const x = onrejected(this.reason);
        	reject(x)
        })
      }
    })

    return promise2
  }
}
```

再使用上面的测试用例，就可以得到正确的结果：

```javascript
let promise = new Promise((resolve, reject) => {
  resolve('100');
})
promise.then((data) => {
  console.log('data1:', data);	// data1: 100
  return 200
}, null).then((data) => {
  console.log('data2:', data);	// data2: 200
  throw new Error('error')
}, null).then(null, () => {
  consol.log('程序报错...')
})
```

上面的测试用例中，当 then 的回调函数抛出异常时需要去捕获错误，传到下一个 then 的失败回调函数中。

```javascript
class Promise {
  ...
  then(onfulfilled, onrejected) {
		let promise2 = new Promise((resolve, reject) => {
      if (this.status === RESOLVED) {
        try{
			const x = onfulfilled(this.value)
			resolve(x)
        } catch(e) {
			reject(e)
        }
      }
      if (this.status === REJECTED) {
        try{
			const x = onrejected(this.reason);
        	resolve(x)
        } catch(e) {
			reject(e)
        }
      }
      if (this.status === PENDING) {
        this.onResolvedCallbacks.push(() => {
          try{
            const x = onfulfilled(this.value)
            resolve(x)
          } catch(e) {
            reject(e)
          }
        })
        this.onRejectedCallbacks.push(() => {
          try{
            const x = onrejected(this.reason);
            resolve(x)
          } catch(e) {
            reject(e)
          }
        })
      }
    })
    return promise2
  }
}
```

到这里为止我们就已经实现了一个简版的 Promise，因为Promise是一个规范，很多人都可以实现自己的 Promise 所以 Promise A+ 规范做了很多兼容处理的要求，如果想实现一个完整的 Promise 可以参考 [Promise A+ 规范](https://promisesaplus.com/) 。

## 3. 小结

本节主要按照 [Promise A+ 规范](https://promisesaplus.com/) 部分的要求实现了一个简版的 Promise API 这个 Promise 基本上满足同步异步的链式调用，对基本的异常做了处理。当然 [Promise A+ 规范](https://promisesaplus.com/) 所规定的细节比较多，剩下的都是对各种异常错误的处理，所以后面我们也没有去实现。另外官网下提供了一个测试用例来验证我们写的 Promise 是否符合 [Promise A+ 规范](https://promisesaplus.com/) ，所以可以参考 [promises-tests](https://github.com/promises-aplus/promises-tests) 这个库来完成我们的 Promise 的测试。

* [【译】 Promises/A+ 规范](https://juejin.im/post/6844903767654023182)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

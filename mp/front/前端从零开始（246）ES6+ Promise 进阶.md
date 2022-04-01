# ES6+ Promise 进阶

## 1. 前言

前两节我们学习了 Promise 的用法，并且在上一节我们动手实现了一个符合 Promise A+ 规范的简版 Promise。真正了解了 Promise 底层是怎么来实现的，更好地帮助我们理解 Promise 并对 Promise 的扩张打下了基础。对 Promise 的扩展会可以解决一些通用的问题，比如使用 `Promise.all()` 去并发请求接口。在 node 中还提供了将 `callback` 类型的 api 转换为 Promise 对象。

本节我们将继续学习 Promise 对象相关 API 的使用。这些 api 在我们的实际应用中会经常使用，并且可以很好的解决常见的问题。

## 2. Promise.resolve () 和 Promise.reject ()

前面我们已经学习了在 `new Promise()` 对象时执行器会提供两个回调函数，一个是 `resolve` 返回一个立即成功的 Promise，一个是 `reject` 返回一个立即失败的 Promise。在执行器中需要根据不同情况调 `resolve` 或 `reject` ，如果我们只想返回一个成功或失败的 Promise 怎么做呢？

Promise 对象上的提供了 `Promise.resolve(value)` 和 `Promise.reject(reason)` 语法糖，用于只返回一个成功或失败的 Promise。下面我们看下它的对比写法：

```javascript
const p1 = new Promise(function(resolve, reject){
    reslove(100)
})
const p2 = Promise.resolve(100) //和p1的写法一样

const p3 = new Promise(function(resolve, reject){
    reject('error')
})
const p4 = Promise.reject('error') //和p3的写法一样
```

通过上面的对比 `Promise.resolve(value)` 创建的实例也具有 then 方法的链式调用。这里有个概念就是：如果一个函数或对象，具有 then 方法，那么他就是 thenable 对象。

```javascript
Promise.resolve(123).then((value) => {
  console.log(value);
});

Promise.reject(new Error('error')).then(() => {
  // 这里不会走 then 的成功回调
}, (err) => {
  console.error(err);
});
```

其实，实现 `Promise.resolve(value)` 和 `Promise.reject(reason)` 的源码是很简单的。就是在 Promise 类上创建 `resolve` 和 `reject` 这个两个方法，然后去实例化一个 Promise 对象，最后分别在执行器中的 `resolve()` 和 `reject()` 函数。按照这个思路有如下实现方式：

```javascript
class Promise {
	...
  resolve(value) {
    return new Promise((resolve, reject) => {
      resolve(value)
    })
  }
  reject(reason) {
    return new Promise((resolve, reject) => {
      reject(reason)
    })
  }
}
```

通过上面的实现源码我们很容易地知道，这两个方法的用法。需要注意的是 `Promise.resolve(value)` 中的 value 是一个 Promise 对象 或者一个 thenable 对象，`Promise.reject(reason)` 传入的是一个异常的原因。

## 3. catch()

Promise 对象提供了链式调用的 catch 方法捕获上一层错误，并返回一个 Promise 对象。catch 其实就是 then 的一个别名，目的是为了更好地捕获错误。它的行为和 `Promise.prototype.then(undefined, onRejected)` 只接收 `onRejected` 回调是相同的，then 第二个参数是捕获失败的回调。所以我们可以实现一个 catch 的源码，如下：

```javascript
class Promise {
  //...
  catch(errorCallback) {
    return this.then(null, errorCallback);
  }
}
```

从上面的实现 catch 的方法我们可以知道，catch 是内部调用了 then 方法并把传入的回调传入到 then 的第二个参数中，并返回这个 Promise。这样我们就更清楚地知道 catch 的内部原理了，以后看到 catch 可以直接把它看成调用了 then 的失败的回调就行。下面我们看几个使用 catch 的例子：

```javascript
let promise = new Promise((resolve, reject) => {
  resolve('100');
})
promise.then((data) => {
  console.log('data:', data);	// data: 100
  throw new Error('error')
}, null).catch(reason => {
  console.log(reason)	// Error: error
})
```

catch 后还可以链式调用 then 方法，默认会返回 undefined。也可以返回一个普通的值或者是一个新的 Promise 实例。同样，在 catch 中如果返回的是一个普通值或者是 resolve，在下一层还是会被 then 的成功回调所捕获。如果在 catch 中抛出异常或是执行 reject 则会被下一层 then 的失败的回调所捕获。

```javascript
promise.then((data) => {
  console.log('data:', data);	// data: 100
  throw new Error('error')
}, null).catch(reason => {
  console.log(reason)	// Error: error
  return 200
}).then((value) => {
  console.log(value)	// 200
}, null)
```

## 4. finally()

finally 是 ES9 的规范，它也是 then 的一个别名，只是这个方法是一定会执行的，不像上面提到的 catch 只有在上一层抛出异常或是执行 reject 时才会走到 catch 中。

```javascript
Promise.resolve('123').finally(() => {
  console.log('100')	// 100
})
```

知道 finally 是 then 的一个别名，那我们就知道在它后面也是可以链式调用的。

```javascript
Promise.resolve('123').finally(() => {
  console.log('100')
  return 200
}).then((data) => {
  console.log(data)	// 123
})
```

需要注意的是在 finally 中返回的普通值或是返回一个 Promise 对象，是不会传到下一个链式调用的 then 中的。如果 finally 中返回的是一个异步的 Promise 对象，那么链式调用的下一层 then 是要等待 finally 有返回结果后才会执行：

```javascript
Promise.resolve('123').finally(() => {
  console.log('100')
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve(100)
    }, 3000)
  })
}).then((data) => {
  console.log(data)	// 123
})
```

执行上面的代码，在 then 中打印的结果会在 3 秒后执行。这也说明了 finally 有类似 sleep 函数的意思。

finally 是 ES9 的规范，在不兼容 ES9 的浏览器中就不能使用这个 api，所以我们可以在 Promise 对象的原型上增加一个 finally 方法。

```javascript
Promise.prototype.finally = function(callback) {
  return this.then((value) => {
    return Promise.resolve(callback()).then(() => value);
  }, (err) => {
    return Promise.reject(callback()).catch(() => {throw err});
  })
}
```

因为 finally 是一定会执行的，所以 then 中的成功和失败的回调都需要执行 finally 的回调函数。使用 `Promise.resolve(value)` 和 `Promise.reject(reason)` 去执行 finally 传入的回调函数，然后使用 then 和 catch 来返回 finally 上一层返回的结果。

## 3. Promise.all () 和 Promise.race ()

在前端面试中经常会问这两个 api 并做对比，因为它们的参数都是传入一个数组，都是做并发请求使用的。

### 3.1 Promise.all()

`Promise.all()` 特点是将多个 Promise 实例包装成一个新的 Promise 实例，只有同时成功才会返回成功的结果，如果有一个失败了就会返回失败，在使用 then 中拿到的也是一个数组，数组的顺序和传入的顺序是一致的。

```javascript
const p1 = Promse.resolve('任务1');
const p2 = Promse.resolve('任务2');
const p3 = Promse.reject('任务失败');

Promise.all([p1, p2]).then((res) => {
  console.log(res);		// ['任务1', '任务2']
}).catch((error) => {
  console.log(error)
})

Promise.all([p1, p3, p2]).then((result) => {
  console.log(result)
}).catch((error) => {
  console.log(error)      // 任务失败
})
```

`Promise.all()` 在处理多个任务时是非常有用的，比如 [Promise 基础](http://baike.imooc.com/tms/section/add?tid=3976&type=0) 一节中使用 `Promise.all()` 并发的请求接口的案例，我们希望得到所以接口请求回来的数据之后再去做一些逻辑，这样我们就不需要维护一个数据来记录接口请求有没有完成，而且这样请求的好处是最大限度地利用浏览器的并发请求，节约时间。

### 3.2 Promise.race()

`Promise.race()` 和 `Promise.all()` 一样也是包装多个 Promise 实例，返回一个新的 Promise 实例，只是返回的结果不同。`Promise.all()` 是所有的任务都处理完才会得到结果，而 `Promise.race()` 是只要任务成功就返回结果，无论结果是成功还是失败。

```javascript
const p1 = new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve('任务1成功...');
  }, 1000)
})
const p2 = new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve('任务2成功...');
  }, 1500)
})
const p3 = new Promise((resolve, reject) => {
  setTimeout(() => {
    reject('任务失败...');
  }, 500)
})

Promise.race([p1, p2]).then((res) => {
  console.log(res);   // 任务1成功...
}).catch((err) => {
  console.log(err);
})

Promise.race([p1, p2, p3]).then((res) => {
  console.log(res)
}).catch((err) => {
  console.log(err)  // 任务失败...
})
```

上面的实例代码充分的展示了 `Promise.race()` 特性，在实际的开发中很少用到这个 api，这个 api 能做什么用呢？其实这个 api 可以用在一些请求超时时的处理。

当我们浏览网页时，突然网络断开或是变得很差的情况下，可以用于提示用户网络不佳，这也是一个比较常见的情况。这个时候我们就可以使用 `Promise.race()` 来处理：

```javascript
const request = new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve('请求成功...');
  }, 3000);
})
const timeout = new Promise((resolve, reject) => {
  setTimeout(() => {
    reject('请求超时，请检查网络...');
  }, 2000);
})

Promise.race([request, timeout]).then(res => {
  console.log(res);
}, err => {
  console.log(err);   // 请求超时，请检查网络...
})
```

上面的代码中定义了两个 Promise 实例，一个是请求实例，一个是超时实例。请求实例当 3 秒的时候才会返回，而超时设置了 2 秒，所以会先返回超时的结果，这样就可以去提醒用户了。

### 3.3 实现 Promise.all ()

> 面试题：实现一个 `Promise.all()` 方法。

前面我们说到了 thenable 对象，也就是判断一个值是不是 Promise 对象，就是判断它是函数或对象，并具有 then 方法。

```javascript
const isPromise = (val) => {
  if (typeof val === "function" || (typeof val == "object" && val !== null)) {
    if (typeof val.then === "function") {
      return true;
    }
  }
  return false;
};
```

`Promise.all()` 会接收一个数组，数组的每一项都是一个 Promise 实例，并且它的返回结果也是一个 Promise，所以我们需要在内部 new 一个 Promise 对象，并返回。在执行器中我们的目标是：

1. 当有实例中有错误或抛出异常时，就要执行执行器中的 reject；
2. 没有错误时，只有所有的实例都成功时才会执行执行器中的 resolve。

基于这两点，有如下步骤：

1. 内部创建一个计数器，用于记住已经处理的实例，当计数的值和传入实例的数组长度相等时，执行执行器中的 resolve；
2. 创建一个用于存放实例返回结果的数组；
3. 处理实例的结果有两种：一种返回的是普通值、一种返回的是 Promise 对象，然后分别处理；
4. 返回普通值结果时直接存放到数组中即可；
5. 返回的是一个 Promise 对象时，就需要调用这个实例上的 then 方法得到结果后在存放到结果数组中去。

根据上面的五个步骤基本就可以把 `Promise.all()` 实现出来了，具体代码如下：

```javascript
Promise.all = function(arr) {
  return new Promise((resolve, reject) => {
    let num = 0;  // 用于计数
    const newArr = [];  // 存放最终的结果

    function processValue(index, value) {	// 处理Promise实例传入的结果
      newArr[index] = value;
      if (++num == arr.length) {	// 当计数器的值和处理的 Promise 实例的长度相当时统一返回保护所以结果的数组
        resolve(newArr);
      }
    }

    for (let i = 0; i < arr.length; i++) {
      const currentValue = arr[i];  // Promise 实例
      if (isPromise(currentValue)) {
        currentValue.then((res) => {
          processValue(i, res);
        }, reject)
      } else {
        processValue(i, currentValue);
      }
    }
  });
}
```

上面的代码已经实现了 `Promise.all()` 方法，可以使用下面的例子进行测试。

```javascript
const p1 = new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve("任务1成功...");
  }, 1000);
});

const p2 = new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve("任务2成功...");
  }, 500);
});

Promise.all([p1, p2]).then((res) => {
  console.log(res)
})
```

## 4. 小结

本节学习了根据 Promise 衍生出的相关 api 的使用，已经每个 api 基本都给出了实现源码，理解这些源码会让我们更加深刻地理解 Promise，在实际的开发过程中达到游刃有余。到此我们花了三节的时间由浅入深来介绍 Promise，花些时间来彻底弄懂这些知识点，对于我们以后学习其他的异步解决方案有更好的理解。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

## 1. 前言

本节介绍混入 (mixin) 的使用方法。包括什么是 mixin，如何定义 mixin，mixin 重名属性的合并策略，全局 mixin 的使用。其中，了解和掌握 mixin 重名属性的合并策略非常重要，属性合并问题会在 mixin 的使用中经常出现。在学完本小节之后，同学们可以尝试总结出属性的合并的规则策略，这有助于加深印象，在日后的使用中也能得心应手。

## 2. 慕课解释

> 混入 (mixin) 提供了一种非常灵活的方式，来分发 Vue 组件中的可复用功能。一个混入对象可以包含任意组件选项。当组件使用混入对象时，所有混入对象的选项将被 “混合” 进入该组件本身的选项。 – 官方定义

我们在日常开发中经常遇到多个页面或者功能模块有相同代码逻辑的情况，同学们在遇到此类情况的时候肯定会想：如果这段代码能够复用就好了！。那什么方法可以帮助我们实现复用呢？答案就是：Mixin！ `Mixin` 帮助我们抽离公共代码逻辑。一个混入对象可以包含任意组件选项。当组件使用混入对象时，所有混入对象的选项将被 “混合” 进入该组件本身的选项。

## 3. 使用 mixin

接下来我一起看看如何定义和使用一个 mixin。

对于 mixin 的使用可以分为两部分：

1. 定义 mixin
2.  混入 mixin

### 3.1 定义一个 mixin

mixin 本质上就是一个 Object 对象，它和 `vue` 实例上的属性一致，包含 data、methods、computed、watch、生命周期函数等等：

```javascript
var myMixin = {
  data(){
    return {
      //...
    }
  },
  created() {
     //...
  },
  methods: {
    //...
  },
  computed() {
    // ...
  }
}
```

### 3.2 混入 mixin

想要混入定义好的 mixin，只需要通过组件的 mixins 属性传入想要混入的 mixin 数组即可：

```javascript
var vm = new Vue({
  el: '#app',
  mixins:[myMixin]
})
```

**代码解释：**

上述代码中，我们定义了一个 Vue 实例，并在实例上混入 myMixin。

接下来我们看一个简单的示例：

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
  <div id="app">
    我是：{{name}}， 年龄：{{year}}
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  // 定义 mixin
  var myMixin = {
    data(){
      return {
        name: '句号'
      }
    },
    created: function () {
      this.mixinFun()
    },
    methods: {
      mixinFun: function () {
        console.log('mixin function')
      }
    }
  };
  var vm = new Vue({
    el: '#app',
    // 使用mixin
    mixins:[myMixin],
    data() {
    	return {
        year: '18'
      }
    }
  })
</script>
</html>

```

**代码解释：**

JS 代码第 3-17 行，定义了一个混入对象 myMixin，并定义了数据 data、钩子函数 created、方法 mixinFun。

JS 代码第 20 行，通过组件实例上的 mixins 属性引入 myMixin。

运行程序可以看到，在 myMixin 中定义的数据 name 渲染到页面上。同时打开控制台可以看到 ‘mixin function’ 被打印出来，说明 created 钩子函数被执行。

## 4. 选项合并

我们在定义 `mixin` 时会出现属性名重复的情况，例如：

```javascript
var myMixin = {
  data() {
    return {
      name: 'Imooc'
    }
  },
  create() {
    console.log('Imooc')
  }
}

var vm = new Vue({
  data() {
    return {
      name: '句号'
    }
  },
  create() {
    console.log('句号')
  }
})
```

当组件和混入对象含有同名选项时，这些选项将以恰当的方式进行 “合并”。这些重复项的合并有固定的规则，接下来我们从三个方面来详细讲解选项合并的规则。

### 4.1 data 的合并

数据对象在内部会进行递归合并，并在发生冲突时以组件数据优先。详细的合并规则如下：

1. 首先判断 mixin 中的数据类型和组件实例对象上的数据类型是否相同；
2. 如果不同，组件实例上的数据将覆盖 mixin 上的数据；
3. 如果相同，判断是否为 Object 的数据格式；
4. 如果不是 Object 的数据格式，组件实例上的数据将覆盖 mixin 上的数据；
5. 如果是 Object 的数据格式，从第一步开始循环判断 Object 的每一个属性。

具体示例：

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
  <div id="app"></div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  var myMixin = {
    data(){
      return {
        isOpen: false,
        date: '2020-02-02',
        desc: {
          title: 'Mixins 基础教程',
          desc: '本教程将讲解如何使用 mixins',
          author: {
            name: '慕课网',
            location: '北京'
          }
        }
      }
    }
  };
  var vm = new Vue({
    el: '#app',
    mixins:[myMixin],
    data() {
    	return {
        isOpen: true,
        date: new Date().toLocaleString(),
        desc: {
          author: {
            name: 'Imooc',
            age: '20'
          }
        }
      }
    },
    created() {
      console.log(this.data)
    }
  })
</script>
</html>

```

**代码解释：**

JS 代码第 3-15 行，定义了一个混入对象 myMixin，并定义了数据 data。

JS 代码第 18 行，通过组件实例上的 mixins 属性引入 myMixin。

JS 代码第 19-30 行，定义了组件实例上的数据 data。

根据之前我们学习的合并规则，得到的最终数据 data 格式如下：

```javascript
data() {
  return {
    isOpen: true,
    date: new Date().toLocaleString(),
    desc: {
      author: {
        name: 'Imooc',
        age: '20',
        location: '北京'
      }
    }
  }
}
```

### 4.2 钩子 的合并

同名钩子函数将合并为一个数组，因此都将被调用。另外，混入对象的钩子将在组件自身钩子之前调用。也就是说，如果我们在 mixin 和组件中都定义了钩子函数 created，那么 mixin 和 组件中的函数都会被执行。需要注意的是：mixin 中的钩子函数将在组件的钩子函数之前执行。

具体示例：

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
  <div id="app">
    mixin 示例
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  var myMixin = {
    created() {
      alert('mixin created 先执行')
    }
  };
  var vm = new Vue({
    el: '#app',
    mixins:[myMixin],
    created() {
      alert('组件 created 后执行')
    }
  })
</script>
</html>
```

**代码解释：**

JS 代码第 3-7 行，定义了一个混入对象 myMixin，并定义了钩子函数 created。

JS 代码第 10 行，混入定义的 myMixin。

JS 代码第 11-13 行，在组件内部定义了钩子函数 created。

所以，最终的运行结果是：

mixin created 执行

组件 created 执行

### 4.3 值为对象的选项合并

值为对象的选项，例如 methods、components 和 directives，将被合并为同一个对象。两个对象键名冲突时，取组件对象的键值对。

示例：

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
  <div id="app">
    mixin 示例
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  var myMixin = {
    methods: {
      sayName() {
        console.log('我是慕课网')
      },
      sayHello() {
        console.log('hello 大家好')
      }
    }
  };
  var vm = new Vue({
    el: '#app',
    mixins:[myMixin],
    methods: {
      sayName() {
        console.log('我是句号')
      },
      sayYear() {
        console.log('我的年龄是：18')
      }
    }
  })
  vm.sayName() // ---> 我是句号
  vm.sayHello() // ---> hello 大家好
  vm.sayYear() // ---> 我的年龄是：18
</script>
</html>

```

**代码解释：**

JS 代码第 3-12 行，定义了一个混入对象 myMixin，并定义了两个方法。

JS 代码第 15 行，混入定义的 myMixin。

JS 代码第 16-23 行，在组件内部定义了两个方法。

由于值为对象的选项合并取组件对象的键值对。所以，最终的 methods 对象是：

```javascript
methods: {
  sayName() {
    console.log('我是句号')
  },
  sayYear() {
    console.log('我的年龄是：18')
  },
  sayHello() {
    console.log('hello 大家好')
  }
}
```

## 5. 全局混入

混入也可以进行全局注册。使用时需要格外小心！一旦使用全局混入，它将影响每一个之后创建的 Vue 实例。使用恰当时，这可以用来为自定义选项注入处理逻辑。通过 Vue.mixin ({…}) 可以注册全局混入：

```javascript
Vue.mixin({
  data: {
    name: "Imooc"
  }
})
```

具体示例：

```javascript
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
  <div id="app">
    mixin 全局混入 示例
  </div>
</body>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script type="text/javascript">
  Vue.mixin({
    created() {
      console.log('全局mixin created')
    },
    methods: {
      sayHello() {
        console.log('hello 大家好')
      }
    }
  })

  var vm = new Vue({
    el: '#app'
  })
  vm.sayHello()
</script>
</html>

```

**代码解释：**

JS 代码第 3-12 行，定义了一个全局混入对象，并定义了钩子函数 created 和 sayHello 方法。

JS 代码第 14-16 行，创建了 Vue 实例。

因为全局混入会在之后创建的每一个 Vue 实例上混入，所以，控制台会输出以下数据：

全局 mixin created

hello 大家好

## 6. 小结

本节，我们带大家学习了混入 mixin 在 vue 项目中的运用。主要知识点有以下几点：

* Mixin 的定义和使用方法。
* Mixin 选项的合并策略。
* 全局 Mixin 的注册和使用。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

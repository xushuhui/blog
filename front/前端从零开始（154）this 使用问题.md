# this 使用问题

大部分开发者都会合理、巧妙的运用 `this` 关键字。

初学者容易在 `this` 指向上犯错，如下面这个 `Vue 组件`：

```javascript
<div id="app"></div>
<script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.9/vue.min.js"></script>
<script>
  // 发送post请求
  const post = (cb) => {
    // 假装发了请求并在200ms后返回了服务端响应的内容
    setTimeout(function() {
      cb([
        {
          id: 1,
          name: '小红',
        },
        {
          id: 2,
          name: '小明',
        }
      ]);
    });
  };

  new Vue({
    el: '#app',
    data: function() {
      return {
        list: [],
      };
    },
    mounted: function() {
      this.getList();
    },
    methods: {
      getList: function() {
        post(function(data) {
          this.list = data;
          console.log(this);
          this.log(); // 报错：this.log is not a function
        });
      },

      log: function() {
        console.log('输出一下 list:', this.list);
      },
    },
  });
</script>
```

这是初学 `Vue` 的同学经常碰到的问题，为什么这个 `this.log()` 会抛出异常，打印了 `this.list` 似乎也是正常的。

这其实是因为传递给 `post` 方法的回调函数，拥有自己的 this。

不光在这个场景下，其他类似的场景也要注意，在写回调函数的时候，如果在回调函数内要用到 `this`，就要特别注意一下这个 `this` 的指向。

可以使用 `ES6 的箭头函数` 或者将需要的 this 赋值给一个变量，再通过作用域链的特性访问即可：

```javascript
<div id="app"></div>
<script src="https://cdn.bootcdn.net/ajax/libs/vue/2.6.9/vue.min.js"></script>
<script>
  // 发送post请求
  const post = (cb) => {
    // 假装发了请求并在200ms后返回了服务端响应的内容
    setTimeout(function() {
      cb([
        {
          id: 1,
          name: '小红',
        },
        {
          id: 2,
          name: '小明',
        }
      ]);
    });
  };

  new Vue({
    el: '#app',
    data: function() {
      return {
        list: [],
      };
    },
    mounted: function() {
      this.getList();
    },
    methods: {
      getList: function() {

        // 传递箭头函数
        post((data) => {
          this.list = data;
          console.log(this);
          this.log(); // 报错：this.log is not a function
        });

        // 使用保留 this 的做法
        // var _this = this;
        // post(function(data) {
        //   _this.list = data;
        //   console.log(this);
        //   _this.log(); // 报错：this.log is not a function
        // });
      },

      log: function() {
        console.log('输出一下 list:', this.list);
      },
    },
  });
</script>
```

这个问题通常初学者都会碰到，之后慢慢就会形成习惯，会非常自然的规避掉这个问题。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

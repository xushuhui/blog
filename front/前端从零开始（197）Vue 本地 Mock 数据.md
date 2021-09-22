# Vue 项目本地 Mock 数据

## 1. 前言

本小节我们将带大家学习如何在 Vue-Cli3 初始化的项目中创建 Mock 数据。

## 2. 简介

在日常开发中，接口的联调是非常普遍的。然而，有些时候接口并不会及时提供，这时候就需要我们自己 Mock 数据来模拟接口的实现。

## 3. 创建 Mock 数据

首先，我们在项目的根路径下创建 vue.config.js 文件，并在文件中写如下配置：

```javascript
module.exports = {
  devServer: {
    before(app) {
      app.get("/goods/list", (req, res) => {
        res.json({
          data: [
            {name: 'Vue 基础教程'},
            {name: 'React 基础教程'}
          ]
        });
      });
    }
  }
};
```

我们通过 axios 来获取接口数据。首先需要安装 axios：

```javascript
npm install axios --save
```

在组件中使用 axios 获取 Mock 数据：

```javascript
<script>
import axios from "axios";
export default {
  name: "Home",
  created() {
    axios.get("/goods/list").then(res => {
      console.log(res);
    });
  },
  components: {}
};
</script>
```

有时候，我们需要写很多的 Mock 数据，把所有的数据都写在 vue.config.js 文件中显然是不合适的，这会使得文件变得非常大，并且难以维护。我们可以在项目中创建 Mock 文件夹，把所有数据放在 Mock 文件夹中维护。

我们在 Mock 文件夹中创建 list.json

```javascript
[
  {"name": "Vue 基础学习"},
  {"name": "React 基础学习"}
]
```

然后，在 vue.config.js 文件中加载数据：

```javascript
const list = require("./mock/list.json");
module.exports = {
  devServer: {
    before(app) {
      app.get("/goods/list", (req, res) => {
        res.json({
          data: list
        });
      });
    }
  }
};

```

## 4. 小结

本节我们带大家学习了如何在项目中使用 Mock 数据。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

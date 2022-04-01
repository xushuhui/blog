# Element 美化 TODO 项目

## 1. 前言

本小节我们将带大家利用 Element 一起优化我们的 TODO 项目。

## 2. Element 简介

Element 是一套为开发者、设计师和产品经理准备的基于 Vue 2.0 的组件库，提供了配套设计资源，帮助你的网站快速成型。由 “饿了么” 公司前端团队开源。

## 3. 编写路由

首先，我们需要定义两个路由，分别是待办事项列表页面和添加待办事项页面。我们可以打开 ‘router/index.js’ 文件进行如下配置：

```javascript
import Vue from "vue";
import VueRouter from "vue-router";
import List from "../views/List.vue";
import Add from "../views/Add.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/list",
    name: "list",
    component: List,
    alias: "/"
  },
  {
    path: "/add",
    name: "home",
    component: Add
  }
];

const router = new VueRouter({
  routes
});

export default router;
```

## 4. 入口文件

要使用 Element 首先我们需要通过 npm 安装 Element：

```javascript
npm install element-ui --save
```

安装完成之后，我们需要修改 main.js

```javascript
import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

Vue.config.productionTip = false;

import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";

Vue.use(ElementUI);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");

```

## 5. 使用 Vuex 保存数据

### 5.1 创建 mutation-types (store/types.js)

```javascript
export const INIT_TODO = "INIT_TODO";
export const ADD_TODO = "ADD_TODO";
export const DEL_TODO = "DEL_TODO";
export const COMPLETE_TODO = "COMPLETE_TODO";
```

### 5.2 创建 store (store/index.js)

```javascript
import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

import { ADD_TODO, DEL_TODO, COMPLETE_TODO, INIT_TODO } from "./types";

export default new Vuex.Store({
  state: {
    list: []
  },
  getters: {
    count: state => isComplete => {
      return state.list.filter(item => item.isComplete === isComplete).length;
    },
    todoList: state => {
      return state.list;
    }
  },
  mutations: {
    [INIT_TODO](state, payload) {
      state.list = payload.list;
    },
    [ADD_TODO](state, payload) {
      state.list.push(payload);
    },
    [DEL_TODO](state, payload) {
      const index = payload.index;
      state.list.splice(index, 1);
    },
    [COMPLETE_TODO](state, payload) {
      const index = payload.index;
      state.list[index].isComplete = 1;
    }
  },
  actions: {
    initList({ commit }) {
      axios.get("/todo/list").then(res => {
        commit(INIT_TODO, { list: res.data.data });
      });
    }
  }
});

```

## 6. 改造 App.vue

```javascript
<template>
  <div id="app">
    <el-container style="height: 100%; border: 1px solid #eee">
      <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
        <el-menu router>
          <el-menu-item index="/list">
            <i class="el-icon-menu"></i>
            <span slot="title">待办列表</span>
          </el-menu-item>
          <el-menu-item index="/add">
            <i class="el-icon-circle-plus-outline"></i>
            <span slot="title">添加待办</span>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <el-container>
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  name: "App",
  components: {},
  methods: {
    ...mapActions(["initList"])
  },
  created() {
    this.initList();
  }
};
</script>

<style lang="scss">
* {
  padding: 0;
  margin: 0;
}
#app {
  height: 100%;
}
</style>

```

## 7. 编写列表页面（views/list.vue）

```javascript
<template>
  <div>
    <el-table
      class="el-table"
      :data="todoList"
      border
      style="width: 100%"
      :row-class-name="tableRowClassName"
    >
      <el-table-column fixed prop="name" label="事项名称" width="150">
      </el-table-column>
      <el-table-column prop="date" label="事项截止时间" width="120">
      </el-table-column>
      <el-table-column prop="type" label="事项类型" width="120">
      </el-table-column>
      <el-table-column prop="urgent" label="是否紧急" width="120">
        <template slot-scope="scope">
          <span>{{ scope.row.urgent | urgentText }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="content" label="事项详情"></el-table-column>
      <el-table-column label="操作" width="160">
        <template slot-scope="scope">
          <el-button
            @click="handleDelete(scope.$index)"
            type="danger"
            size="small"
            >删除</el-button
          >
          <el-button
            type="primary"
            size="small"
            @click="handleComplete(scope.$index)"
            :disabled="scope.row.isComplete === 1"
            >完成</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <div style="text-align:left">
      总共：{{ todoList.length }} 个任务。 已完成：{{
        count(1)
      }}
      个任务。未完成：{{ count(0) }} 个任务。
    </div>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from "vuex";
import { COMPLETE_TODO, DEL_TODO } from "../store/types";
export default {
  methods: {
    ...mapMutations([COMPLETE_TODO, DEL_TODO]),
    handleDelete(index) {
      this.DEL_TODO({ index });
    },
    handleComplete(index) {
      this.COMPLETE_TODO({ index });
    },
    tableRowClassName({ row }) {
      if (row.isComplete === 1) {
        return "complete-row";
      }
      return "";
    }
  },
  computed: {
    ...mapGetters(["todoList", "count"])
  },
  filters: {
    urgentText(value) {
      if (value === 0) {
        return "紧急";
      }
      return "非紧急";
    }
  },
  data() {
    return {
      tableData: []
    };
  }
};
</script>
<style>
.el-table .complete-row {
  background: #f0f9eb;
}
</style>

```

## 8. 编写添加事项页面

```javascript
<template>
  <el-form ref="form" :model="form" :rules="rules" label-width="180px">
    <el-form-item prop="name" label="事项名称：">
      <el-input v-model="form.name"></el-input>
    </el-form-item>

    <el-form-item prop="date" label="事项截止时间：">
      <el-date-picker
        type="date"
        placeholder="选择日期"
        v-model="form.date"
        value-format="yyyy-MM-dd"
        style="width: 100%;"
      ></el-date-picker>
    </el-form-item>

    <el-form-item prop="type" label="事项类型：">
      <el-select
        v-model="form.type"
        placeholder="请选择活动区域"
        style="width: 100%;"
      >
        <el-option label="学习" value="学习"></el-option>
        <el-option label="工作" value="工作"></el-option>
        <el-option label="游戏" value="游戏"></el-option>
      </el-select>
    </el-form-item>

    <el-form-item prop="urgent" label="是否紧急：">
      <el-radio-group v-model="form.urgent">
        <el-radio :label="1">是</el-radio>
        <el-radio :label="0">否</el-radio>
      </el-radio-group>
    </el-form-item>

    <el-form-item prop="content" label="事项详情：">
      <el-input type="textarea" v-model="form.content"></el-input>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button @click="cancel">取消</el-button>
    </el-form-item>
  </el-form>
</template>
<script>
import { mapMutations } from "vuex";
import { ADD_TODO } from "../store/types";
export default {
  data() {
    return {
      form: {
        name: "",
        content: "",
        date: "",
        urgent: 1,
        type: ""
      },
      rules: {
        name: [
          { type: "string", required: true, message: "请填写待办事项的名称" }
        ],
        content: [
          {
            type: "string",
            required: true,
            message: "请填写待办事项的详情",
            trigger: "blur"
          },
          {
            type: "string",
            min: 20,
            max: 50,
            message: "长度限制在20-50个字符",
            trigger: "blur"
          }
        ],
        type: [
          { type: "string", required: true, message: "请填写待办事项的称类型" }
        ]
      }
    };
  },
  methods: {
    ...mapMutations([ADD_TODO]),
    onSubmit() {
      this.$refs.form.validate(validate => {
        if (validate) {
          this.ADD_TODO({ ...this.form, isComplete: 0 });
          this.$message({
            message: "添加成功",
            type: "success"
          });
          this.$refs.form.resetFields();
        }
      });
    },
    cancel() {
      this.$refs.form.resetFields();
    }
  }
};
</script>

```

![图片描述](https://xushuhui.gitee.io/image/imooc/5ee2de070a251d0c19020936.jpg)

## 9. 小结

本小节我们主要带大家一起使用 Element 优化了我们之前的 TODO 项目，并在项目中，把我们在之前章节中学习的知识点加以运用。其实，诸如 Element 之类的 Vue 组件库还有很多，比如： 嘀嘀团队的 Cube-UI、有赞团队的 Vant 等等，使用这些组件库可以让我们快速高效地完成项目。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

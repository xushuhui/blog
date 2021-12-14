# web 数据库

之前的章节有讨论过 web 中的存储方式，包括传统的 cookie 和新的 localstorage，这两种方式实现了 HTML 中的离线存储，但是存储方式比较简单，在有些复杂的业务场景可能不能满足条件。本章我们介绍一个计算机中一个重要的学科数据库，以及它在 HTML5 中的支持。数据库是一个内容庞大的知识体系，本章只介绍一些简单的用法以及它在 HTML 中的适用场景。

## 1. 适用场景

既然适用 localstorage 也可以做简单的数据存储，那么为什么还需要适用数据库呢？假设一个业务场景中将所有用户信息临时存储到浏览器中，这些信息包括昵称、姓名、性别等，现在需要搜索出性别是男的所有用户。如果使用 localstorage 的话，需要将所有的数据提取出来，一条条遍历，得出结果。这样的搜索算法的时间复杂度是 O(n)，性能较差。如果使用数据库存储的话，只需要给性别列加上索引，然后使用 SQL 搜索，时间复杂度是 O(lgn)，性能提升了一个等量级。关系型数据库的特点是：

* 数据模型基于关系，结构化存储，完整性约束；
* 支持事务，数据一致性；
* 支持 SQL，可以复杂查询；

缺点是：

* SQL 解析会影响性能；
* 无法适应非结构化存储；
* 横向扩展代价高；
* 入门门槛较高。

## 2. Web SQL

Web SQL 不是 HTML5 标准中的一部分，它是一个独立的规范，引入了 SQL 的 api，关于 SQL 的语法可以参考第三方的教程，在此不做解释。Web SQL 有 3 个函数

### 2.1 openDatabase

这个函数用于打开一个数据库，如果数据库不存在就创建。它有 5 个参数，分别表示：

* 数据库名称；
* 版本号；
* 数据库备注；
* 初始化数据大小；
* 创建 / 打开成功回调函数

```javascript
/**
     * 创建数据库 或者此数据库已经存在 那么就是打开数据库
     * name: 数据库名称
     * version: 版本号
     * displayName: 对数据库的描述
     * estimatedSize: 设置数据的大小
     * creationCallback: 回调函数(可省略)
     */
    var db = openDatabase("MySql", "1.0", "数据库描述", 1024 * 1024);
```

### 2.2 transaction

这个函数使用事务执行 SQL 语句，它是一个闭包，例如：

```javascript
dataBase.transaction( function(tx) {
    tx.executeSql(
    "create table if not exists test (id REAL UNIQUE, name TEXT)",
    [],
    function(tx,result){ alert('创建test表成功'); },
    function(tx, error){ alert('创建test表失败:' + error.message);
    });
});
```

### 2.3 executeSql

这个方法用于执行 SQL 语句。

```javascript
tx.executeSql(
"update stu set name = ? where id= ?",
[name, id],
function (tx, result) {
},
function (tx, error) {
alert('更新失败: ' + error.message);
});
});
```

## 3. indexedDB

IndexedDB 是 HTML5 规范里新出现的浏览器里内置的数据库。它提供了类似数据库风格的数据存储和使用方式。存储在 IndexedDB 里的数据是永久保存，不像 cookies 那样只是临时的。IndexedDB 里提供了查询数据的功能，在线和离线模式下都能使用。

### 3.1 对比 Web SQL

跟 WebSQL 不同的是，IndexedDB 更像是一个 NoSQL 数据库，而 WebSQL 更像是关系型数据库。

### 3.2 判断浏览器是否支持

```javascript

window.indexedDB = window.indexedDB || window.mozIndexedDB || window.webkitIndexedDB || window.msIndexedDB;
if(!window.indexedDB){
    console.log("你的浏览器不支持IndexedDB");
}
```

### 3.3 创建库

使用 open 方法创建数据库。

```javascript
 var request = window.indexedDB.open("testDB", 2);//第一个参数是数据库的名称，第二个参数是数据库的版本号。版本号可以在升级数据库时用来调整数据库结构和数据
 request.onsuccess = function(event){
    console.log("成功打开DB");
 }//成功之后的回调函数
```

### 3.4 添加数据

```javascript
var transaction = db.transaction(["students"],"readwrite");//先创建事务，具有读写权限
transaction.oncomplete = function(event) {
    console.log("Success");
};
transaction.onerror = function(event) {
    console.log("Error");
};
var test = transaction.objectStore("test");
test.add({rollNo: rollNo, name: name});//添加数据
```

### 3.5 查询

```javascript
var request = db.transaction(["test"],"readwrite").objectStore("test").get(rollNo);//创建具备读写功能的事务
request.onsuccess = function(event){
    console.log("结果 : "+request.result.name);
};//成功查询的回调函数
```

### 3.6 修改

```javascript

var transaction = db.transaction(["test"],"readwrite");//创建事务
var objectStore = transaction.objectStore("test");
var request = objectStore.get(rollNo);
request.onsuccess = function(event){
    console.log("Updating : "+request.result.name + " to " + name);
    request.result.name = name;//修改数据
    objectStore.put(request.result);//执行修改
};
```

### 3.7 删除

```javascript
//创建事务，并删除数据
db.transaction(["students"],"readwrite").objectStore("students").delete(rollNo);
```

## 4. 实际项目应用

```javascript
<!DOCTYPE html>
<html>
    <head lang="en">
        <meta charset="UTF-8">
        <title>离线记事本</title>
        <meta name="viewport" content="width=device-width,initial-scale=1">
        <link rel="stylesheet" href="http://code.jquery.com/mobile/1.4.5/jquery.mobile-1.4.5.min.css" />
        <script src="http://code.jquery.com/jquery-1.11.1.min.js"></script>
        <script src="http://code.jquery.com/mobile/1.4.5/jquery.mobile-1.4.5.min.js"></script><!-- 引用jQuery插件 -->
    </head>
    <script>
var datatable = null;
var db = openDatabase("note", "", "notebook", 1024 * 100);
//初始化函数方法
function init() {
    datatable = document.getElementById("datatable");
    showAllData();
}
function removeAllData() {
    for(var i = datatable.childNodes.length - 1; i >= 0; i--) {
        datatable.removeChild(datatable.childNodes[i]);
    }
    var tr = document.createElement("tr");
    var th1 = document.createElement("th");
    var th2 = document.createElement("th");
    var th3 = document.createElement("th");
    th1.innerHTML = "标题";
    th2.innerHTML = "内容";
    th3.innerHTML = "时间";
    tr.appendChild(th1);
    tr.appendChild(th2);
    tr.appendChild(th3);
    datatable.appendChild(tr);
}
//显示数据库中的数据
function showData(row) {
    var tr = document.createElement("tr");
    var td1 = document.createElement("td");
    td1.innerHTML = row.title;
    var td2 = document.createElement("td");
    td2.innerHTML = row.content;
    var td3 = document.createElement("td");
    var t = new Date();
    t.setTime(row.time);
    td3.innerHTML = t.toLocaleDateString() + " " + t.toLocaleTimeString();
    tr.appendChild(td1);
    tr.appendChild(td2);
    tr.appendChild(td3);
    datatable.appendChild(tr);
}
//显示所有的数据
function showAllData() {
    db.transaction(function(tx) {
        tx.executeSql("CREATE TABLE IF NOT EXISTS item(title TEXT,content TEXT,time INTEGER)", []);
        tx.executeSql("SELECT * FROM item", [], function(tx, rs) {
            removeAllData();
            for(var i = 0; i < rs.rows.length; i++) {
                showData(rs.rows.item(i))
            }
        })
    })
}
//添加一条记事本数据
function addData(title, content, time) {
    db.transaction(function(tx) {
        tx.executeSql("INSERT INTO item VALUES (?,?,?)", [title, content, time], function(tx, rs) {
                alert("保存成功！");
            },
            function(tx, error) {
                alert(error.source + "::" + error.message);
            }
    )
    })
}
//点击保存按钮
function saveData() {
    var title = document.getElementById("name").value;
    var content = document.getElementById("memo").value;
    var time = new Date().getTime();
    addData(title, content, time);
    showAllData();
}

    </script>
    <body onload="init()">
        <div data-role="page" id="pageone">
            <div data-role="header" data-position="fixed">
                <h1>离线记事本</h1>
            </div>
            <div data-role="main" class="ui-content">
                <p align="center">记事</p>
                <table data-role="table" class="ui-responsive">
                    <thead>
                        <tr>
                            <th>标题：</th>
                            <th>内容：</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td><input type="text" id="name"></td>
                            <td><input type="text" id="memo"></td>
                        </tr>
                    </tbody>
                </table>
                <button type="submit" onclick="saveData()">保存</button>
                <table data-role="table" data-mode="" class="ui-responsive" id="datatable">
                </table>
            </div>

        </div>
    </body>
</html>
```

上述代码通过使用 websql 实现了一个简单的离线记事本的功能，数据库中保留 3 个字段，分别是标题、内容、时间，点击保存按钮调用`insert` 豫剧将数据添加到数据库，然后通过使用 `select`语句将数据库中的数据展示出来。如果浏览器不主动清空数据的情况下离线数据将会永久保存，这样的话借助 websql 可以实现与桌面应用相差无几的功能。

## 5. indexedDB 和 websql 对比

* **访问限制：** indexdb 和 websql 一致，均是在创建数据库的域名下才能访问，且不能指定访问域名。

* **存储时间：** 这两位的存储时间也是永久，除非用户清除浏览器数据，可以用作长效的存储。

* **大小限制：** 理论上讲，这两种存储的方式是没有大小限制的。然而 indexeddb 的数据库超过 50M 的时候浏览器会弹出确认，基本上也相当于没有限制了。但是由于不同的浏览器的实现有一定的差别，实际使用中需要根据不同的浏览器做相应的容量判断容错。

* **性能测试：** indexeddb 查询少量数据花费差不多 20MS 左右。大量数据的情况下，相对耗时会变长一些，但是也就在 30MS 左右，也是相当给力了，10W 数据 +，毕竟 nosql。而 websql 的效率也不错，10w+ 数据，简单查询一下，只花费了 20MS 左右。

* **标准规范：** Web SQL 数据库是一个独立的规范，因为安全性能等问题，官方现在也已经放弃了维护；indexedDB 则属于 W3C 标准。

## 6. 小结

回顾本章，由关系数据库的优缺点及适用场景引申到 HTML5 中的数据库解决方案，以及使用方法，需要注意的是在使用 HTML 数据库的过程中需要检测浏览器是否支持数据库。实际开发项目由于考虑前端数据库的安全性以及性能等问题，如果切实需要使用需要谨慎，毕竟一般项目中数据库保存的都是敏感数据，即使保存在服务器中也需要一定的安全加密措施，所以一般前端存储的都是一些临时的数据。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

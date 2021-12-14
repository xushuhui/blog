# 认识 select option 下拉菜单标签

下拉菜单也是我们在网页中比较常见的场景。如果我们的选项过多，如果是放单选框或者多选框的话，会造成页面显示不太优雅，会铺放很多的选项。这个时候我们使用下拉菜单是最合适不过的。下拉菜单可以提供很多选项，是比较方便的一种操作。例如以下情况：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cf6f09941a3b04370404.jpg)

使用下拉菜单一方面列表框为用户输入数据提供了一种便捷方式，只要把已知的数据项列举出来，用户在列表框中选择列表项就可以；另一方面可以把需要搜集的数据规范化，防止因用户输入数据的随意性而造成后端处理数据的混乱。例如，在搜集用户个人信息时，可能需要用户输入职业信息，后端程序需要按照职业信息对用户进行职业分类。在这种情况下，采用列表框元素就是比较好的获取用户数据的方式，可以预先规划好职业信息，把规划好的职业信息以列表框方式展现给用户，用户只需要选择列表项就可以了，无需让用户输入职业信息。

## 1. select option 标签的使用

select option 标签和 ul li 标签类似，select 标签代表下拉菜单整体，而 option 则是下拉菜单的每一个选项，代码如下：

```javascript
<select>
    <option>苹果</option>
    <option>香蕉</option>
    <option>橘子</option>
</select>
```

效果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cf7c09bfe24e03230151.jpg)

我们可以给给 option 标签设置 `disabled`属性，代表当前选项是禁用项，不能选择的，代码如下：

```javascript
<select placeholder="请选择">
    <option>苹果</option>
    <option disabled>香蕉</option>
    <option>橘子</option>
</select>
```

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cf86091f153b04190222.jpg)

我们也可以给 option 标签设置 `selected`属性来表示默认选中的选项，代码如下：

```javascript
<select placeholder="请选择">
    <option>苹果</option>
    <option>香蕉</option>
    <option selected>橘子</option>
</select>
```

效果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cf93094ed78e04150132.jpg)

下拉菜单默认宽度为选项内容宽度撑开，如果想设置下拉菜单的样式，可以借助 CSS 。

## 2. 经验分享

1. select 标签里面只能嵌套 option 标签，不能嵌套其他标签；
2. 为了增强用户体验，我们一般把下拉菜单的第一项设置为请选择，然后设置为默认选项。

## 3. 真实案例分享



```javascript
<select>
    <option value="">请选择</option>
    <option value="2000">2000</option>
    <option value="2001">2001</option>
    <option value="2002">2002</option>
    <option value="2003">2003</option>
    <option value="2004">2004</option>
    <option value="2005">2005</option>
    <option value="2006">2006</option>
    <option value="2007">2007</option>
    <option value="2008">2008</option>
    <option value="2009">2009</option>
    <option value="2010">2010</option>
    <option value="2011">2011</option>
    <option value="2012">2012</option>
    <option value="2013">2013</option>
    <option value="2014">2014</option>
    <option value="2015">2015</option>
    <option value="2016">2016</option>
    <option value="2017">2017</option>
    <option value="2018">2018</option>
    <option value="2019">2019</option>
    <option value="2020">2020</option>
</select>
```

京东官网

```javascript
<select>
    <option selected="selected">请选择出生月份：</option>
    <option>1</option>
    <option>2</option>
    <option>3</option>
    <option>4</option>
    <option>5</option>
    <option>6</option>
    <option>7</option>
    <option>8</option>
    <option>9</option>
    <option>10</option>
    <option>11</option>
    <option>12</option>
</select>
```

## 4. 小结

1. select 标签表示下拉菜单整体，option 标签表示下拉菜单的每一个选项。
2. select 标签里只能嵌套 option 标签。
3. 设置 option 标签的 `disabled`属性，可以禁用该选项。
4. 设置 option 标签的 `selected` 属性，可以默认选择该选项。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f63080f0945d31613860750.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

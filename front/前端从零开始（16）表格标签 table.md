# 认识表格标签 table 标签

表格在我们的网页中是非常常见的，比如我们要展示商品信息，工作安排，产品参数等都需要用到表格。那么在 html 中，使用表格就需要用到 table 标签了。但是表格不仅是 table 一个标签，需要用到和表格相关的一组标签，这一小节我们就来学习这些标签吧。

## 1. 表格的结构

在使用表格标签之前，我们需要先认识一下在 html 中表格是由哪些结构构成的。一般情况下，表格都表头、表身构成。表头里面放每一列对应的字段，一般是一个描述，如姓名、年龄等，表身则放每一个表头对应的具体的值，如张三对应表头的姓名，20 对应表头的年龄。表格有行和列的概念，行就代表一个数据的所有属性，比如姓名：张三，年龄：20，而列则表示一个表头对应的所有数据，比如姓名这一列就只有姓名，如张三、李四等。通常情况下，我们都是先确定表头，然后再确定每行对应表头的具体数据。

## 2. 表格的使用

想要编写表格，需要用到表格的一组标签。table 标签表示表格整体，类似 ul 和 ol 表示列表整体一样。在 table 标签里， thead 标签表示表头， tbody 标签表示表示。 在 table 表头中， tr 标签代表行， th 标签代表表头的每一项。在 tbody 标签中， tr 标签代表行， td 标签代表每一个表头对应的具体数据。**代码如下：**

```javascript
 <table>
    <!-- thead 代表表头 -->
    <thead>
      <!-- tr代表表头这一行 -->
      <tr>
        <!-- th代表表头的每一项 会有加粗的效果 -->
        <th>姓名</th>
        <th>年龄</th>
        <th>性别</th>
      </tr>
    </thead>
    <!-- tbody 代表表身 表格的具体内容 -->
    <tbody>
      <!-- tr代表表身的每一行 -->
      <tr>
        <!-- td代表对应表头的具体数据 -->
        <td>小明</td>
        <td>20</td>
        <td>男</td>
      </tr>
      <tr>
        <td>小红</td>
        <td>18</td>
        <td>女</td>
      </tr>
    </tbody>
  </table>
```

**效果如下：**

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cae70994cc2306420162.jpg)

我们可以给表格添加 `border`属性给表格添加边框，`border`属性的值为正整数，默认为 0，则无边框，我们把`border` 设置为 1，**代码如下：**

```javascript
<table border='1'>
    <!-- 代码和上面演示代码一致 -->
    ...
</table>
```

则会呈现以下效果：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cb0109b47e5d04450128.jpg)

我们还可以给 table 设置`cellpadding`来使用单元格填充来创建单元格内容与其边框之间的空白，`cellpadding`值也是正整数，我们把表格的 `cellpadding`设置为 10，**代码如下：**

```javascript
<table border='1' cellpadding='10'>
    <!-- 代码和上面演示代码一致 -->
    ...
</table>
```

则效果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cb3a097b19ff04650191.jpg)

我们还可以给 table 设置`cellspacing`来设置单元格与单元格直接的距离，`cellspacing`值也是正整数，我们把表格的 `cellspacing`设置为 10，**代码如下：**

```javascript
<table border='1' cellspacing='10'>
    <!-- 代码和上面演示代码一致 -->
    ...
</table>
```

效果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cb49094591be04570160.jpg)

我们也可以为表格添加标题，那么我们就需要在 thead 标签前加上 caption 标签，而 caption 标签的内容则是表格的标题，**代码如下：**

```javascript
<table>
    <caption>学生表</caption>
    <!-- 代码和上面演示代码一致 -->
    ...
</table>
```

效果如下：

![图片描述](https://xushuhui.gitee.io/image/imooc/5f07cb570944745f04130197.jpg)

## 3. 注意事项

1. tr 标签只能嵌套 th 和 td 标签，不能嵌套其他标签；
2. tr 代表表格的每一行；
3. td 标签的内容必须和表头的信息对应。

## 4. 真实案例分享

北京大学官网

```javascript
<table>
    <thead>
        <tr>
            <th>课号</th>
            <th>课程名称</th>
            <th>开课单位</th>
            <th>学分</th>
            <th>总周数<br>(起止周)</th>
            <th>课程类型</th>
            <th>上课时间</th>
            <th>班号</th>
            <th>上课教师</th>
            <th>备注</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>01132632</td>
            <td><p style="text-align:center;padding:5px;"><a href="http://elective.pku.edu.cn/elective2008/edu/pku/stu/elective/controller/courseDetail/getCourseDetail.do?kclx=BK&course_seq_no=BZ1920301132632_15745" target="_blank">生物化学讨论课<br>Current topics on Biochemistry</a></p></td>
            <td>生命科学学院</td>
            <td>2</td>
            <td>2(1-2)</td>
            <td>A</td>
            <td><p>星期一(第10节-第12节)</br>星期二(第10节-第12节)</br>星期三(第10节-第12节)</br>星期四(第10节-第12节)</br>星期五(第10节-第12节)</br>星期六(第10节-第12节)</br></p></td>
            <td>1</td>
            <td>钟上威</td>
            <td><p>6月29-7月7日，10-12节，选修同学需先修生物化学理论课</p></td>
		</tr>
		<tr>
    <td>01132022</td>
    <td><p><a href="http://elective.pku.edu.cn/elective2008/edu/pku/stu/elective/controller/courseDetail/getCourseDetail.do?kclx=BK&course_seq_no=BZ1920301132022_18636" target="_blank">遗传学讨论<br>Current topics on Genetics</a></p></td>
    <td>生命科学学院</td>
    <td>2</td>
    <td>2(3-4)</td>
    <td>A</td>
    <td><p>星期一(第10节-第12节)</br>星期二(第10节-第12节)</br>星期三(第10节-第12节)</br>星期四(第10节-第12节)</br>星期五(第10节-第12节)</br></p></td>
	<td>1</td>
	<td>范六民</td>
	<td><p>上课时间：7月13日-7月24日，10-12节。选修同学需先修遗传学理论课</p></td>
	</tr>
  </tbody>
</table>
```

## 5. 小结

1. 表格的结果为表头和表身，可以为表格添加标题。
2. thead 代表表头，tbody 代表表身， thead 嵌套 tr 和 th， tbody 里嵌套 tr 和 td。
3. tr 里只能嵌套 th 和 td， th 一般用来表示表头，有加粗的样式。 td 代表表头对应的具体数据。
4. 原生的表格样式比较丑，我们可以通过 CSS 为表格设置样式。

![图片描述](https://xushuhui.gitee.io/image/imooc/5f6302ba0968ecb514250831.jpg)

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

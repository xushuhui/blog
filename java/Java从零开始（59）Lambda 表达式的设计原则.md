---
title: Java 从零开始（59）Lambda 表达式的设计原则
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Lambda 表达式的设计原则

Java 编程最基本的原则就是要追求高内聚和低耦合的解决方案和代码模块设计，这里我们主要讨论在 Lambda 表达式的环境下来设计我们程序时的两点原则：**单一原则** 和 **开放闭合原则**。

## 1. 单一原则

> 程序中的类或者方法只有一个改变的理由

当需要修改某个类的时候原因有且只有一个。换句话说就是让一个类只做一种类型责任，当这个类需要承当其他类型的责任的时候，就需要分解这个类。 类被修改的几率很大，因此应该专注于单一的功能。如果你把多个功能放在同一个类中，功能之间就形成了关联，改变其中一个功能，有可能中止另一个功能，这时就需要新一轮的测试来避免可能出现的问题，非常耗时耗力。

我们来看一个质数统计的例子：

```java
//这是一个违反单一原则的例子
public SRPDemo{
	public static long countPrimes(int maxNum){
		long total = 0;
		for(int i = 1; i<maxNum;i++){
			boolean isPrime = true;
			for( int j=2; j<i;j++){
				if(i%j ==0){
					isPrime = false;
				}
			}
			if(isPrime){
				total = total +1;
			}
		}
		return total;
	}
	​
	public static void main(String ...s){
		System.out.println(countPrimes(100));
	}
}
```

```java
输出结果：

26
```

上面的例子违反了单一原则，一个方法包含了两重职责：

* 计数；
* 判断是否为一个质数。

我们把上面的代码重构，将两种职责拆分到两个方法中：

```java
//符合单一原则的例子
public SRPDemo{
	//计数
	public static long countPrimes(int maxNum){
		long total = 0;
		for(int i= 1;i<maxNum;i++){
			if(isPrime(i))
				total = total+1;
		}
			return total;
	}
	//判断是否为一个质数
	public static boolean isPrime(int num){
		for(int i = 2;i<num; i++){
			if(num%i ==0){
				return false;
			}
		}
		return true;
	}
​
	public static void main(String ...s){
		System.out.println(countPrimes(100));
	}
}
```

我们现在使用集合流来重构上面代码：

```java
public SRPDemo{
	public static long countPrimes(int maxNum){
		return IntStream.range(1,maxNum).filter(MultipleInterface::isPrime).count();
	}
	public static boolean isPrime(int num){
		return IntStream.range(2,num).allMatch(x -> num%x != 0);
	}
	​
	public static void main(String ...s){
		System.out.println(countPrimes(100));
	}
}
```

可见，我们使用集合流在一定程度上可以轻松地帮我们实现单一原则。

## 2. 开放闭合原则

> 软件应该是扩展开放，修改闭合

* 通过增加代码来扩展功能，而不是修改已经存在的代码；
* 若客户模块和服务模块遵循同一个接口来设计，则客户模块可以不关心服务模块的类型，服务模块可以方便扩展服务（代码）；
* 开放闭合原则支持替换的服务，而不用修改客户模块。

我们来看一个发送消息的例子，假设我们现在有一个消息通知模块用来发送邮件和短信：

```java
//这是一个违反开放闭合原则的例子
public class OCPDemo{
	//发送邮件
	public boolean sendByEmail(String addr, String title, String content) {
		System.out.println("Send Email");
		return true;
	}
	//发送短信
	public boolean sendBySMS(String addr, String content) {
		System.out.println("Send sms");
		return true;
	}
}
```

想必很多人都会这么写，这么写有一个问题，如果哪一天需求变更要求增加微信消息通知，这个时候不仅需要增加一个 `sendWechat`的方法，还需要在调用它的地方进行修改，所以违反了 OCP 原则。现在我们来做一下修改：

```java
//一个满足开放闭合原则的例子
public class OCPDemo{
	@Data
	public static class Message{
		private String addr;
		private String title;
		private String content;
		private int type;
	}
	public boolean send(Message message){
		switch (message.getType()){
			case 0: {
				System.out.println("Send Email");
				return true;
			}
			case 1:{
				System.out.println("Send sms");
				return true;
			}
			case 2:{
				System.out.println("Send QQ");
				return true;
			}
			default:return false;
		}
	}
}
```

我们创建了一个 `Message` 对象来描述发送消息的所有信息，并增加了一个 `type` 字段用来区分发送渠道。在遇到类似的情况窝子需要在 `send` 方法中增加一个对应 渠道类型 `type` 的处理逻辑就可以了，对存量代码无需求改。满足了 OCP 原则。

现在我们再来看下使用函数式接口怎么来优化我们的程序：

```java
@Data
public class OCPDemo{
	@Data
	public static class Message{
		private String addr;
		private String title;
		private String content;
	}
​
	private Message message;
	​
	public boolean send(Function<Message , Boolean>  function){
		return function.apply(message);
	}
​
	public static void main(String ...s){
		Message message = new Message();
		message.setTitle("this is a qq msg");
		OCPDemo demo = new OCPDemo();
		demo.setMessage(message);
		demo.send((msg)->{
			System.out.println("send qq:\t"+msg.getTitle());
			return true;
		});
	}
}
```

```java
输出：

send qq:  this is a qq msg
```

此时，我们运用函数接口 `Function` 来处理 `Message`, 省去了消息类型的判断，仅当调用的时候决定使用哪种渠道发送，当然我们可以把发送逻辑都写在一个工具类里面，利用 Lambda 引用来调用。

## 3. 小结

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1a95b509dfdc8507900373.jpg)

本节主要讨论的是我们如何在我们的程序设计中来使用 Lambda 表达式时所涉及的两条原则 —— **单一原则** 和 **开放闭合原则**。

这里关注的是程序整体，而不是具体的某一个方法。其前提是对于 Lambda 表达式的深度理解和熟练运用，为了说明问题，例子大多相对简单，想了解更详细的设计原则还是需要阅读相关的专著（比如 S.O.L.I.D 原则），并在日常的编码过程中不断实践和思考。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

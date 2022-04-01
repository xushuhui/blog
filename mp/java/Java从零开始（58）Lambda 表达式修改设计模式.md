---
title: Java 从零开始（58）Lambda 表达式修改设计模式
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---
# Lambda 表达式修改设计模式

本节内容并不是讨论设计模式，而是讨论如何使用 Lambda 表达式让现有的设计模式变得更简洁，或者在某些情况是有一些不同的实现方式。我们可以从另一个角度来学习使用和理解 Lambda 表达式。

> **Tips：** 要更好地理解本节内容需要对涉及的四个设计模式有一定的了解，具体可以查阅相关资料。

## 1. 命令者模式

> 命令者模式是将操作、方法调用、命令封装成一个对象，在合适的时候让该对象进行执行。

它包五个角色：

* **客户端（Client）**：发出命令；
* **调用者（Invoker）**：调用抽象命令，还可以记录执行的命令；
* **接受者（Receiver）**：命令的实际执行者，一个命令会存在一个或多个接收者；
* **抽象命令（Command）**：定义命令执行方法；
* **具体命令（Concrete Command）**：调用接收者，执行命令的具体方法。

命令者模式被大量运用在组件化的图形界面系统、撤销功能、线城市、事务和向导中。我们来看一个例子，我们实现一个将一系列命令录制下来的功能，有点类似于 Word 中的撤销功能那样记录每一步的操作。

```java
//定义一个命令接收者，包含打开、关闭和保存三个操作
public class Editor{
	public void save(){
		System.out.println("do save")
	}
	public void open();
	public void close();
}
​
// 定名命令对象，所有操作都要实现这个接口
public interface Action{
	public void perform();
}
​
//实现保存命令操作
public Save implements Action{
	private final Editor editor;

	public Save(Editor editor){
		this.editor = editor;
	}

	public void perform(){
		editor.save();
	}
}
​
//实现打开命令操作
public class Open implements Action{
	private final Editor editor;
	​
	public Open(Editor editor){
		this.editor = editor;
	}
	public void perform(){
		editor.open();
	}
}
​
//实现关闭命令操作
public class Close implements Action{
	private final Editor editor;
	​
	public Close(Editor editor){
		this.editor = editor;
	}
	public void perform(){
		editor.close();
	}
}
​
//定义命令发起者来记录和顺序执行命令
public class Invoker{
	private final List<Action> actions = new ArrayList<>();
	​
	public void record(Action action){
		actions.add(action);
	}
	public void run(){
		for (Action action : actions) {
			action.perform();
		}
	}
}
​
//定义客户端，用来记录和执行命令
public class Client{
	public static void main(String...s){
		Invoker invoker = new Invoker();
		Editor editor = new Editor();
		//记录保存操作
		invoker.record(new Save(editor));
		//记录打开操作
		invoker.record(new Open(editor));
		//记录关闭操作
		invoker.record(new Close(editor));
		invoker.run();
	}
}
```

```java
输出结果：

do save
do open
do close
```

以上是一个完整的命令者模式的例子，我们使用 Lambda 表达式来修改客户端：

```java
public class Client{
	public static void main(String...s){
		Invoker invoker = new Invoker();
		Editor editor = new Editor();
		//记录保存操作
		invoker.record(()->editor.open());
		//记录打开操作
		invoker.record(()->editor.save());
		//记录关闭操作
		invoker.record(()->editor.close());
		invoker.run();
	}
}
```

我们使用引用方法来修改客户端：

```java
public class Client{
	public static void main(String...s){
		Invoker invoker = new Invoker();
		Editor editor = new Editor();
		//记录保存操作
		invoker.record(editor::open);
		//记录打开操作
		invoker.record(editor::save);
		//记录关闭操作
		invoker.record(editor::close);
		invoker.run();
	}
}
```

通过这样的改造，我们的代码意图更加明显了呢，一看就明白具体记录的是哪个操作。

## 2. 策略模式

> 策略模式是软件运行时，根据实际情况改变软件的算法行为。

常见的策略模式就是文件压缩软件，通常一个压缩软件可以支持多种压缩算法如 zip 、gzip、rar 等，通过策略模式可以让压缩软件根据我们具体的操作来实现不同的压缩算法。我们来看一个压缩数据的策略模式的例子：

```java
//定义压缩策略接口
public interface CompressionStrategy{
	public OutputStream compress(OutputStream data) throws IOException;
}

//gzip 压缩策略
public class GzipStrategy implements CompressionStrategy{
	@Override
	public OutputStream compress(OutputStream data) throws IOException {
		return new GZIPOutputStream(data);
	}
}

//zip 压缩策略
public class ZipStrategy implements  CompressionStrategy{
	@Override
	public OutputStream compress(OutputStream data) throws IOException {
		return new ZipOutputStream(data);
	}
}
​
//在构造类时提供压缩策略
public class Compressor{
	private final CompressionStrategy strategy;
	public Compressor(CompressionStrategy strategy){
		this.strategy = strategy;
	}
	public void compress(Path inFiles, File outputFile) throws IOException{
		try(OutputStream outputStream = new FileOutputStream(outputFile)){
			Files.copy(inFiles,strategy.compress(outputStream));
		}
	}
}
```

```java
//使用具体的策略初始化压缩策略
//gzip 策略
Compressor gzipCompressor = new Compressor(new GzipStrategy());
//zip 策略
Compressor zipCompressor = new Compressor(new ZipStrategy());
```

以上就是一个完整的 zip 和 gzip 的压缩策略。现在我们用 Lambda 表达式来优化初始化压缩策略

```java
//使用构造器引用优化初始化压缩策略
//gzip 策略
Compressor gzipCompressor = new Compressor(GzipStrategy::new);
//zip 策略
Compressor zipCompressor = new Compressor(ZipStrategy::new);
```

## 3. 观察者模式

> 观察者模式是定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。

观察者模式被适用于消息通知、触发器之类的应用场景中

观察者模式包含三个类 主题（Subject）、观察者（Observer）和客户端（Client）。Subject 对象带有绑定观察者到 Client 对象和从 Client 对象解绑观察者的方法。

我们来看一个例子：现在我们用观察者模式实现根据输入的数字，自动将输入的数字其转变成对应十六进制和二进制。

```java
//定义一个观察者
public interface Observer {
	public  void update(int num);
}
​
//创建主题
public static class Subject{
	private List<Observer> observers = new ArrayList<>();
	private int num;
	public int getNum(){
		return num;
	}
	private void setNum(int num){
		this.num = num;
		this.notifyAllObservers();
	}
private void addObserver(Observer observer) {
	observers.add(observer);
}
​
	private void notifyAllObservers(){
		for(Observer observer:observers){
			observer.update(num);
		}
	}
}
​
//创建二进制观察者
public static class BinaryObserver implements Observer{
​
	private Subject subject;
	​
	@Override
	public void update(int num) {
		System.out.println( "Binary String: "
					+ Integer.toBinaryString( num ) );
		}
}
​
//创建十六进制观察者
public static class HexObserver implements Observer{
​
	@Override
	public void update(int num) {
		System.out.println( "Hex String: "
					+ Integer.toHexString( num ) );
	}
}
​
//使用 Subject 和实体观察者对象
public class Demo{
	public static void main(String... s){
		Subject subject = new Subject();
		subject.addObserver(new BinaryObserver());
		subject.addObserver(new HexObserver());
		System.out.println("first input is：11");
		subject.setNum(11);
		System.out.println("second input is：15");
		subject.setNum(15);
	}
}
```

```java
输出结果：

first input is：11
Binary String: 1011
Hex String: b
second input is：15
Binary String: 1111
Hex String: f
```

同样我们使用 Lambda 表达式来修改 `Demo` 类：

```java
public class Demo{
	public static void main(String...s){
		Subject subject = new Subject();
		subject.addObserver( num -> System.out.println( "Binary String: " + Integer.toBinaryString( num )));
		subject.addObserver( num -> System.out.println( "Hex String: " + Integer.toHexString(num )));
		System.out.println("first input is：11");
		subject.setNum(11);
		System.out.println("second input is：15");
		subject.setNum(15);
	}
}
```

在这个例子中，我们实际上是省去了 `BinaryObserver` 和 `HexObserver` 两个类的定义，直接使用 Lambda 表达式来描述二进制和十六进制转化的逻辑。

## 4. 模板方法模式

> 模板方法模式是定义一个操作中的算法的骨架，从而将一些步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。

通常对于一些重要的复杂方法和多个子类共有的方法且逻辑相同的情况下会使用模板方法模式。比如用户第三方用户认证的时候就比较适合使用模板方法。

我们来看一个例子：假设我们现在需要用到微信、微博的第三方用户授权来获取用户的信息。

```java
//使用模板方法模式描述获取第三方用户信息的过程
public abstract class Authentication{
	public void checkUserAuthentication(){
		checkIdentity();
		fetchInfo();
	}
​
	protected abstract void checkIdentity();
	protected abstract void fetchInfo();
}
​
//微信用户
public class WechatAuthenication extends Authentication{
	@Override
	protected void checkIdentity() {
		System.out.println("获得微信用户授权");
	}
​
	@Override
	protected void fetchInfo() {
		System.out.println("获取微信用信息");
	}
}
​
//微信用户
public class WeiboAuthenication extends Authentication{
	@Override
	protected void checkIdentity() {
		System.out.println("获得微博用户授权");
	}
​
	@Override
	protected void fetchInfo() {
		System.out.println("获取微博用信息");
	}
}
​
//调用模板方法
public class Demo{
	public static void main(String...s){
		Authentication auth = new WechatAuthenication();
		auth.checkUserAuthentication();
		auth = new WeiboAuthenication();
		auth.checkUserAuthentication();
	}
}
```

```java
输出结果：

获得微信用户授权
获取微信用信信息
获得微博用户授权
获取微博用信信息
```

现在我们使用 Lambda 表达式换个角度来思考模板方法模式。如果我们用函数式接口来组织模板方法中的调用过程，相比使用继承来构建要显得灵活的多。

```java
//定义一个处理接口，用来处理一项事务，如授权或者获取信息。
public interface Processer{
	public void process();
}
​
//封装调用过程
public class Authentication{
	private final Processer identity;
	private final Processer userinfo;

	public Authentication(Criteria identity,Criteria userinfo){
		this.identity = identity;
		this.userinfo = userinfo;
	}
​
	public void checkUserAuthentication(){
		identity.process();
		userinfo.process();
	}
}
​
//使用模板方法
public class Demo{
	Authentication auth = new Authentication(()->System.out.println("获得微信用户授权"),
				()->System.out.println("获取微信用户信息"));
	auth.checkUserAuthentication();
	auth = new Authentication(()->System.out.println("获得微博用户授权"),
				()->System.out.println("获取微博用户信息"));
	auth.checkUserAuthentication();
}
```

```java
输出结果：

获得微信用户授权
获取微信用信信息
获得微博用户授权
获取微博用信信息
```

此时，我们的模板方法得到了大幅的简化，同时通过函数接口让模板方法获得了极大的灵活性。

## 5. 小结

![图片描述](https://xushuhui.gitee.io/image/imooc/5f1a9521099dd8a309870414.jpg)

本节我们讨论如何使用 Lambda 表达式让我们的设计模式变得更简单、更好用。这里我们使用了四个例子从不同的角度来。

* **命令者模式**：我们使用 Lamabda 表达式的方法引用来进行改造；
* **策略模式**：我们使用了 Lambda 表达式的构造器引用来进行改造；
* **观察者模式**：我们使用了标准的 Lambda 表达式来进行改造；
* **模板方法模式**：我们使用了函数式接口来进行改造。

目的是希望给大家一点启发，在平常的编码过程中去思考如何使用 Lambda 表达式来设计我们的程序。对于其他的设计模式如果感兴趣的话可以自己尝试下去修改它们。

### 微信公众号老徐说

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

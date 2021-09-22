# Spring Boot 第一个项目

## 1. 前言

Spring Boot 可以使用 Maven 构建，遵循 Maven 的项目结构规范，项目结构是模板化的，基本都一模一样。

模板化的东西可以自动生成，Spring 官方就提供了 Spring Initializr 。它能自动生成 Spring Boot 项目，我们直接导入到开发工具使用即可。

## 2. 生成 Spring Boot 项目

打开 Spring Initializr 网址 `http://start.spring.io` ，根据我们项目的情况填入以下信息。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2drr40j60pm0ldjtb02)

Spring Initializr 生成 Spring Boot 项目

这是第一次接触 Spring Initializr ，我们来详细了解界面上选项的作用。

**1. 构建方式选择**：此处我们选择 Maven Project 即可，表示生成的项目使用 Maven 构建。当然我们也可以发现，Spring Boot 项目亦可采用 Gradle 构建，目前 Spring Boot 主流的构建方式还是 Maven；

** 2. 编程语言选择**：此处选择 Java 即可；

** 3. Spring Boot 版本选择**： 2.x 版本与 1.x 版本还是有一些区别的，咱们学习肯定是选择 2.x 新版本。此处虽然选择了 2.2.6 版本，但是由于 2.2.6 版本刚推出没多久，国内一些 Maven 仓库尚不支持。后面我们手工改为 2.2.5 版本，便于使用国内 Maven 仓库快速构建项目；

** 4. 所属机构设置**：Group 表示项目所属的机构，就是开发项目的公司或组织。因为公司可能会重名，所以习惯上采用倒置的域名作为 Group 的值。例如的域名是 `imooc.com` , 此处写 `com.imooc` 就行了；

**5. 项目标识设置**：Artifact 是项目标识，用来区分项目。此处我们命名为 `spring-boot-hello` ，注意项目标识习惯性地采用小写英文单词，单词间加横杠的形式。比如 Spring Boot 官方提供的很多依赖，都是 `spring-boot-starter-xxx` 的形式；

**6. 项目名称设置**：Name 是项目名称，保持与 Artifact 一致即可；

** 7. 默认包名设置**：Package name 是默认包名，保持默认即可；

** 8. 打包方式选择**：此处选择将项目打包为 Jar 文件；

** 9. 添加项目依赖**：此处不必修改，我们直接在 pom.xml 中添加依赖更加方便。注意 pom.xml 就是 Maven 的配置文件，可以指定我们项目需要引入的依赖；

** 10. 生成项目**：点击 Generate 按钮，即可按我们设置的信息生成 Spring Boot 项目了。

## 3. Spring Boot 项目结构分析

我们将下载的 zip 压缩包解压后导入开发工具，此处以 Eclipse 为例，依次点击 File-Import-Existing Maven Projects ，然后选择解压后的文件夹导入。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2e1xrpj60ld0eeta402)

Eclipse 导入 Spring Boot 项目

导入后项目结构如下图，我们逐一分析下他们的用途：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo2edc0qj60lf07c3yw02)

Spring Boot 项目结构

* 最外层的 spring-boot-wikis 表示工作集（working set），可以理解为项目分类。我们将 Spring Boot 学习项目都放入该工作集下，便于集中查看；
* spring-boot-hello 是我们指定的项目名称；
* src/main/java 是 Java 源代码目录，存放我们编写的 Java 代码；
* src/main/resources 目录是静态资源目录，存放图片、脚本文件、配置文件等静态资源；
* src/test/java 目录是测试目录，存放测试类。测试是非常重要的，从目录级别跟源代码同级，就能看出来测试的重要性；
* target 目录存放我们打包生成的内容；
* pom.xml 是项目的 Maven 配置文件，指定了项目的基本信息以及依赖项，Maven 就是通过配置文件得知项目构建规则的。

> **Tips：** 此处有同学要发问了，不是说好 Spring Boot 没有配置文件吗？不要着急，Spring Boot 可以在没有配置文件时照常运行。但如果需要个性化功能的话，就会用到配置文件了。 Spring Boot 的配置文件使用非常简单，放心就是了！

## 4. pom.xml 详解

大家可能也发现了，到此刻为止，Spring Boot 项目也没啥新鲜的。都是之前了解过的东西，跟普通的 Maven 项目也没啥区别。

其实真正的变化在 pom.xml 中，我们马上打开瞧一瞧。因为 pom.xml 配置比较长，我们从头到尾分段解释下。

### 4.1 Maven 文档配置

这一段配置代码，其实是固定的格式，表示当前文档是 Maven 配置文档。

**实例：**

```java
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>
</project>
```

### 4.2 Spring Boot 版本配置

这一段配置代码，指定使用 Spring Boot 2.2.5.RELEASE 版本 。如果我们要更换 Spring Boot 版本，只需要修改 `<version></version>` 标签中间的版本号部分即可。

**实例：**

```java
	<parent>
		<groupId>org.springframework.boot</groupId>
		<artifactId>spring-boot-starter-parent</artifactId>
		<version>2.2.5.RELEASE</version>
		<relativePath/> <!-- lookup parent from repository -->
	</parent>
```

### 4.3 项目信息配置

这一段配置代码，大家看到应该比较眼熟，内容即为之前使用 Spring Initializr 指定的项目信息。其中，groupId 是机构标识、artifactId 是项目标识，version 是版本号，name 是项目名称，description 是项目的简单描述。

**实例：**

```java
	<groupId>com.imooc</groupId>
	<artifactId>spring-boot-hello</artifactId>
	<version>0.0.1-SNAPSHOT</version>
	<name>spring-boot-hello</name>
	<description>Demo project for Spring Boot</description>
```

> **Tips：** name 是项目的名称，不用特别严谨。而 artifactId 是用来区分 group 下面的子项目的，需要保证严格唯一。一般情况下将 artifactId 和 name 设置成一样的就可以了。

### 4.4 依赖配置

接下来，这一段代码配置，负责指定 Spring Boot 项目中需要的依赖。 Spring Boot 有一些起步依赖，形如 `spring-boot-starter-xxx` 的样式。起步依赖整合了很多依赖项，后续我们慢慢了解即可。

**实例：**

```java
	<dependencies>
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter</artifactId>
		</dependency>
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-test</artifactId>
			<scope>test</scope>
			<exclusions>
				<exclusion>
					<groupId>org.junit.vintage</groupId>
					<artifactId>junit-vintage-engine</artifactId>
				</exclusion>
			</exclusions>
		</dependency>
	</dependencies>
```

> **Tips：** 可以看到上面两个依赖我们并没有指定版本号，其实是因为 Spring Boot 2.2.5 已经有默认的依赖项版本号了。这是通过 Maven 父继承实现的，即 `<parent>` 标签配置部分，这个稍作了解即可。

### 4.5 插件配置

最后的这一段代码配置，指定了一个插件，用来构建、运行、打包 Spring Boot 项目。

**实例：**

```java
	<build>
		<plugins>
			<plugin>
				<groupId>org.springframework.boot</groupId>
				<artifactId>spring-boot-maven-plugin</artifactId>
			</plugin>
		</plugins>
	</build>
```

## 5. 视频演示

## 6. 小结

本章先讲了 Spring Boot 项目的构建方法，然后大体描述了 Spring Boot 项目的结构和配置文件，让大家有一个总体的感性认识。

实际使用中， Spring Boot 是高度封装的，我们开箱即用即可。在学习阶段，用不着了解很多的原理。

就像你开汽车，会挂挡就行，不需要知道变速箱是啥工作原理。

框架封装的目的就是为了傻瓜式使用， Spring Boot 就是这样的一个傻瓜式工具框架。等哪一天用得很溜了，再去研究原理也不迟。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

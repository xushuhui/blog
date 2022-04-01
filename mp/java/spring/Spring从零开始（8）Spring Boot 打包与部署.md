# Spring Boot 打包与部署

## 1. 前言

项目开发完毕后，免不了将前后端应用打包，然后部署到生产服务器上运行。本篇就演示一个标准的打包、部署过程。

## 2. 操作流程

我们以上一篇开发的前后端分离项目 `spring-boot-cors` 为例进行打包、部署演示，步骤如下：

### 2.1 服务器运行环境安装

一般服务器采用 Linux 或者 Windows Server 系统，相对而言 Linux 系统更加稳定安全。实际上 Windows Server 系统对于一般应用来说也足够了，本篇我们使用 Windows Server 系统进行演示。

推荐使用云服务器，更加稳定且易于维护，国内厂商阿里云、华为云都还不错。

> **Tips：** 云服务器的硬盘读写性能非常重要，在购买云服务器时务必关注下云硬盘的 IOPS 值（衡量硬盘读写性能的一个指标），一般建议要采用 IOPS > 3800 的云磁盘。

具备云服务器后，需要安装 JDK 以便运行 Spring Boot 应用。由于 nginx 对静态资源的负载能力非常强悍，所以我们将前端应用部署到 nginx 上。

### 2.2 Spring Boot 打包为 jar 并运行

Spring Boot 应用可以打包为 war 或者 jar ，官方和我个人都是推荐打 jar 包。可以直接运行，无需部署到 Web 服务器上。

打开命令行工具，进入 `spring-boot-cors` 项目目录后运行 `mvn clean package -Dmaven.test.skip=true` 命令，即可快速打包 Spring Boot 应用。下图中的 jar 文件，即为打包后的 Spring Boot 应用。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3nwuo5j60jh05b3z502)

打包后生成的文件内容

接下来我们将该应用拷贝至服务器，在同一目录下新建 start.bat 文件，内容如下：

```java
java -jar spring-boot-cors-0.0.1-SNAPSHOT.jar
```

双击 start.bat 文件即可启动项目，效果如下，可以看出系统已经启动成功（started）。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3o4h5ej60ig02k0t102)

Spring Boot 打包项目已启动

### 2.3 Spring Boot 打包为 war 并运行

也可以选择将 Spring Boot 打包为 war ，然后放置于 Tomcat 的 webapps 目录下加载运行，接下来我们就详细描述下打包为 war 的过程。

首先，在 pom.xml 文件中修改默认的打包方式，显式指定打包方式为 war 。

```java
	<groupId>com.imooc</groupId>
	<artifactId>spring-boot-cors</artifactId>
	<version>0.0.1-SNAPSHOT</version>
	<name>spring-boot-cors</name>
	<description>Demo project for Spring Boot</description>
    <packaging>war</packaging>
```

然后，由于 Spring Boot 内置了 Tomcat ，所以我们在打包时需要排除内置的 Tomcat ，这样可以避免内置 Tomcat 和 war 包部署运行的 Tomcat 产生冲突。在 pom.xml 中添加如下依赖即可：

```java
		<!--打war包时排除内置的tomcat-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-tomcat</artifactId>
            <scope>provided</scope>
        </dependency>
```

再然后，还需要继承 SpringBootServletInitializer 类并重写 configure 方法，这是为了告诉 Tomcat 当前应用的入口在哪。

```java
@SpringBootApplication
public class SpringBootCorsApplication extends SpringBootServletInitializer {
	@Override
    protected SpringApplicationBuilder configure(SpringApplicationBuilder application) {
        return application.sources(SpringBootCorsApplication.class);
    }
	public static void main(String[] args) {
		SpringApplication.run(SpringBootCorsApplication.class, args);
	}
}
```

最后，即可同样使用 `mvn clean package -Dmaven.test.skip=true` 命令打包应用了，运行命令后会在 target 目录下生成 war 文件，将该文件放置于 Tomcat 的 webapps 目录下运行即可。

### 2.4 前端应用部署

前端应用的部署更加简单，我们直接在云服务器上下载 nginx 然后解压。

打开网址 `http://nginx.org/en/download.html` ，点击下图中的链接下载即可。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3oe52xj60ji03ljrj02)

nginx 下载链接

下载解压后，将前端页面直接放到 nginx/html 目录下即可。当然如果有很多网页，可以先在该目录下建立子目录便于归类网页。

我们建立 shop-front 目录（表示商城系统的前端项目），然后将网页放入其中，效果如下：

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3om7n5j60jf043aaa02)

商城系统前端项目目录内容

注意还需要修改 `goods.html` 中访问的后端 URL 地址，假设云服务器的公网 IP 为 `x.x.x.x` ，则修改为：

**实例：**

```java
$.ajax({
      type: "GET",
      url: "http://x.x.x.x:8080/goods", //后端接口地址
      dataType: "json",
      contentType: "application/json; charset=utf-8",
      success: function (res) {
        $.each(res, function (i, v) {
          row = "<tr>";
          row += "<td>" + v.id + "</td>";
          row += "<td>" + v.name + "</td>";
          row += "<td>" + v.price + "</td>";
          row += "<td>" + v.pic + "</td>";
          row += "</tr>";
          $("#goodsTable").append(row);
        });
      },
      error: function (err) {
        console.log(err);
      }
    });
```

此处解释下后端地址 `http://x.x.x.x:8080/goods` ， HTTP 代表协议， x.x.x.x 代表云服务器公网地址， 8080 是我们后端项目的启动端口，由于我们没有在配置文件中设置，所以默认就是 8080 ，最后 goods 是控制器中设定的后端接口路径。

双击 nginx.exe 启动 nginx ，由于 nginx 默认启动端口是 80 ，所以此时访问 `http://x.x.x.x` ，效果如下，说明 nginx 启动成功！

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3ox78sj60ji05xdgh02)

nginx 已启动成功

### 2.5 测试

现在我们的后端 Spring Boot 应用已启动，前端项目也通过 nginx 启动起来。

我们在浏览器地址栏打开 `http://x.x.x.x/shop-front/goods.html` ，效果如下，说明我们的项目全部部署成功。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo3p864xj60jg06qjrp02)

项目部署成功后页面显示效果

## 3. 视频演示

## 4. 小结

前后端分离部署的方式，更能发挥服务器的性能，如果要进行版本升级，直接替换后端 jar 或者前端项目文件夹即可，轻松愉快。

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

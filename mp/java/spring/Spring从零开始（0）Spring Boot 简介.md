# Spring Boot 简介

## 1. 前言

每逢春暖花开的时节，我都会想起大学时代。那时候的我，在阳光明媚的日子里，坐在图书馆的落地窗前。桌子上是一叠 Java Web 书本，还有我那破破却可爱的笔记本电脑。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo21z8bbj60ji0azdi202)

你是否也偶尔怀念，大学时代的似水流年（图片来源于网络，版权归原作者所有）

那是 SSH 风华正茂的年代，Spring 如日中天，负责整合各种框架，俨然一副老大哥的样子；Hibernate 是数据持久层的不二之选，iBatis 在它面前就像个小老弟；Struts 则是 MVC 框架的形象代言，不懂点 Struts 都不好意思说在做 Web 开发。

而我却总是，被 SSH 繁琐的配置困扰。SSH 各有一大堆配置，当他们碰到一起，还需要额外互相配置。就像三个老朋友，每次再重逢，还要互相介绍。

做一个简单的项目，竟有一大半时间在配置。不是在编辑配置文件的路上，就是在修复配置错误的途中。

程序开发不应该是简单而优雅的吗？正如我们所追求的生活。

## 2. Spring 的诞生

实际上，让开发变得简单，是 Spring 诞生的原动力。

Java 官方推出的企业级开发标准是 EJB ，但 EJB 是相当臃肿、低效的，且难以测试，把当时的 Java 开发者折腾得不轻。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo227ylpj60ji06jq3902)

Spring 官网介绍：让 Java 变简单

那时候，国外有一个年轻的小伙 Rod Johnson，对 SSH 的繁琐产生了质疑。他不光质疑，还去做了他认为对的事情。

经过不断的经验总结和实践，他在 2004 年推出了经典力作《Expert one-on-one J2EE Development without EJB》。该书奠定了 Spring 框架的思想基础，把 EJB 的种种缺点逐一否定，还提出了简洁的替代方案。

从此 Rod Johnson 和 Spring 框架一炮而红，其影响之深远，恐怕连 Rod Johnson 自己都想不到吧。

有时候，不要过于迷信官方，也要敢于思考和质疑。实践是检验真理的唯一标准，编程也不外乎是。

## 3. Spring 的发展

随着 Spring 的流行，Spring 团队也深感责任重大。Spring 团队对 Spring 的优化工作也从未停歇，从 Spring1.x 到现在的 Spring5.x，每一个版本号都是进化的脚印。

最开始的时候，Spring 只支持基于 XML 的配置，后来又陆续增加了对注解配置、Java 类配置的支持。

但是无论怎么变换，都需要开发人员手工去配置，而这些配置往往千篇一律，令人乏味。

我们驾驶汽车，默认都是车窗关闭、空调关闭、仪表盘开启这样的设置。如果每次进入汽车，都要手工逐一设置一遍，其实完全没有必要。

同理，既然大多数人开发 Spring 应用，都有默认的习惯。那何不直接提供默认配置，项目启动时自动采用默认配置，只有当需要个性化功能时，再去手工配置。

所以，在 2014 年，一个叫 Spring Boot 的框架，就这么出现了。

## 4. Spring Boot 的由来

Spring Boot 为简化 Spring 应用开发而生，Spring Boot 中的 **Boot** 一词，即为快速启动的意思。Spring Boot 可以在零配置情况下一键启动，简洁而优雅。

为了让 Spring 开发者痛快到底，Spring 团队做了以下设计：

* 简化依赖，提供整合的依赖项，告别逐一添加依赖项的烦恼；
* 简化配置，提供约定俗成的默认配置，告别编写各种配置的繁琐；
* 简化部署，内置 servlet 容器，开发时一键即运行。可打包为 jar 文件，部署时一行命令即启动；
* 简化监控，提供简单方便的运行监控方式。

基于以上设计目的，Spring 团队推出了 Spring Boot 。

## 5. Spring Boot 的江湖地位

由于 Spring Boot 设计优雅，实现简单，可以节省不少开发时间。

从此，程序员们有了更多时间去陪妹子逛街买裙子。没有女朋友的小伙伴们，也有了更多时间思考追女孩的方案（一定要勇敢地行动呀）。从一定程度上讲，Spring Boot 降低了程序员群体的单身比例。

所以 Spring Boot 的火爆是必然的，据了解，Spring Boot 框架已经是 Java 企业级应用开发的主流框架了。

另外由于微服务的火爆，作为 Spring Cloud 实现基础的 Spring Boot ，更是春风得意，风头一时无两。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubo22fu5hj60ji064t9602)

从 Spring Boot 在 Spring 官网的菜单位置，可以一瞥 Spring Boot 的地位

所以不管出于哪种目的，为跳槽、为加薪、为方便、为省心、为学习、为进步、为爱情、为家庭，Spring Boot 都是 Java 开发旅途的重要风景。

而我，本系列文章的作者，愿陪你看万山红遍、层林尽染，用尽量轻松的语言，讲一些编程的故事和经验，陪你度过一段愉快的 Spring Boot 学习时光。

## 6. Spring Boot 的学习基础

Spring Boot 非常好用，但是并不是 0 基础就可以直接上手的。

Java 语言基础是必备的，这个不必赘述。

在学习 Spring Boot 之前，最好是已经对 Spring 及 Spring MVC 框架有一定的了解。Spring Boot 是一个快速开发框架，其技术基础几乎全部来源自 Spring 。

所以本系列教程的学习基础，是 Java 、 Spring 及 Spring MVC 。其中 Spring MVC 是 Spring 大家庭的非常重要的一员，所以此处单独拿出来强调下。

## 7. 小结

Spring Boot 简单易用，可以快速上手，迅速提高开发效率，值得学习！

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

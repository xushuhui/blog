# 互联网标准化组织

## 1. 前言

**互联网**（Internet）行业两个非常重要的标准化组织是 IETF（Internet Engineering Task Force）和 W3C（Internet Engineering Task Force）。IETF 主要负责 ISO 参考模型相关协议的标准化，比如 TCP、IP、OSPF、DNS 等。W3C 主要是负责 WWW 相关协议的标准化，比如 HTML、XML、SOAP 等。

## 2. IETF RFC

TCP/IP 标准都是以 **RFC**（Request for Comments）文档出版。RFC 格式诞生于 1969 年，是 ARPANET 项目的开创性成果。如今 RFC 已经是 TCP/IP 标准的的官方发行渠道。请看下图来了解 RFC 的主要来源机构。

![图片描述](https://tvax1.sinaimg.cn/large/0032zTO4gy1gubnyo7vc6j60g5095abt02)

* 互联网工程任务组 IETF（Internet Engineering Task Force）是一个开放标准组织，主要工作是开发和推广自愿的 Internet 标准，专注于工程和标准制定的短期问题研究。

* 互联网研究任务组 IRTF（Internet Research Task Force）是一个专注于和 Internet 相关的、长期问题的研究。

* 互联网体系结构委员会 IAB（Internet Architecture Board）是 IETF 的委员会，也是 ISOC 的咨询机构。其职责是对 IETF 活动进行监督。

* 互联网协会 ISOC（Internet Society）是成立于 1992 年的美国非营利性组织，领导着互联网标准的推广、教育、政策方面工作。

RFC 主要是由工程师和计算机科学家以备忘录的形式撰写的。很多 RFC 是实验性的，并不是标准，IETF 会吸收一些 RFC 建议，然后将其作为标准发布。RFC 标准都是以 "RFC xxx"的形式发布的，xxx 是编号，比如 RFC 1122。RFC 经过几十年的发展，到写作此文为止，最新编号是 **RFC 8900**。关于 RFC 的索引库，可以参考 [rfc index](https://www.rfc-editor.org/rfc-index2.html)。

IETF 主要涉及的领域：

* Internet 主要是 IP 层相关协议的开发、扩展。比如，IPv6、DHCP 等。

* Routing 负主要是 IP 路由相协议的开发。比如，MPLS、BGP、OSPF 等。

* Transport 主要是 QoS、端到端传输相关协议开发。

* Security 主要是网络安全相关协议开发。

* Network O&M（Operations And Management）

主要是网络服务、管理相关工作，比如，SNMP、MIB 等。

* User Services 主要是提供标准化过程中的文档。

* Application 主要是应用协议开发和扩展。比如，FTP、SMTP、TELNET。

## 3. W3C

W3C 也是一个互联网标准组织，是 1994 年由 Tim Berners-Lee 发起的，目前也是由此人领导的。W3C 最初的想法是建立统一、兼容的 HTML 标准，并且建议各个浏览器厂商采用 W3C 标准，从而解决各个厂商之间浏览器不兼容的问题。W3C 的主要工作范围是 WWW（World Wide Web）。W3C 的目标是用 Web 将人类以一种更高效、更公平的方式连接在一起。W3C 的成员来源于各大公司或者研究机构，个人只能以特邀专家的形式参与。

W3C 是一个非盈利性组织，由四个机构共同管理：

* 欧洲信息学和数学研究联合会（ERCM）
* 美国麻省理工学院计算机科学与人工智能实验室（MIT CSAIL）
* 日本的 Keio 大学
* 中国北京航空航天大学

W3C 目前分了很多工作组，比如 HTML 工作组、CSS 工作组，SVG 工作组。每个工作组的最终目标是发布 Web 标准规范，官方叫做 W3C 推荐（Recommendation）。W3C Recommendation 标准化过程：

* 记录（Note）

记录，也叫编辑草案，一般来源于 W3C 成员的提交，或者是内部员工的想法，或者是相关方不完善的提议。记录不一定会产生工作组，也不一定会形成 Recommendation。当某项提交被 W3C 认可，就会成立一个工作组，其中包括会员和其他相关方。工作组会发布工作草案。

* 工作草案 Working Draft (WD)

发布工作草案是标准化的第一阶段。当工作草案发布给社区，经过社区评审以后，可能会产生一些不一致的意见，这需要草案的负责人进一步修改和完善，最终达成一致意见。当工作草案经过几轮评审后，没有分歧，就会发布候选推荐。

* 候选推荐 Candidate Recommendation (CR)

候选推荐是比工作草案更接近标准的版本。进入这一阶段，意味着工作组对该标准达到其目标非常有信心。候选推荐的目的是从开发社区获得更多的帮助。因为有些标准是比较复杂，需要会员的帮助。

候选推荐可能会进一步更改，但修改范围局限在比较重要的特性。

* 提议推荐 Proposed Recommendation (PR)

提议推荐是比较完善的版本，不会进行大范围修改，只是进行 bug 的修复。到这一阶段，文档需要提交给顾问委员会批准。

* W3C 推荐 W3C recommendation (REC)

W3C 推荐是规范的最高级别了，到这一阶段，规范经过很多轮的测试和评审，经过了理论和实践的考验，被 W3C 所接受，鼓励大范围的实现和推广。

## 4. 小结

学习计算机网络，就是学习各种协议。如果不了解历史，直接研究某个协议，会有一种不适感、枯燥感，总是想知道为什么要这样设计呢？如果了解了互联网标准化历史和标准化过程，再去研究各种协议，会有一种亲切感，觉得协议的设计就是合情合理的。

TCP/IP 协议族的实现都包含在 RFC 标准中，Web 应用中的协议都包含在 W3C 标准中。在工作中如果需要学习某种协议，直接参考 RFC 和 W3C 官方标准即可。

## 5. 参考

[1]: [rfc_index] [https://www.rfc-editor.org/rfc-index2.html](https://www.rfc-editor.org/rfc-index2.html)

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)

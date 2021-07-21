spring.io
springboot 是spring下的一个模块，springboot模块下的一些模块
优先选择springboot current 带GA版本比较稳定
2.2.1.RELEASE
第一位主版本 稳定 短期不变 除非是很大调整 不向前兼容  主流springboot2
次版本 发布新特性 保证兼容
增量版本 bug修复 保证兼容
RELEASE 发布版本，里程碑版本
RC/Alpha/Beta/GA 版本发布计划，发布状态
SNAPSHOT 比较稳定，还在持续改进
General Availability（GA）：发布后，可用于生产环境。官方推荐使用该版本
微软的windows系统用RC作为正式发布版本


maven/gradle 依赖管理
maven xml组织依赖关系
grade 用Groovy 代码量少，阅读方便简洁
springboot都支持，第三方包支持maven更广泛
maven从中央仓库下载包，grade的中央仓库也是maven的中央仓库，grade基于maven，复用maven
Groovy编程 DSL语言，sql也是DSL语言，Java/Python是DPL 自定义构建打包流程，提供构建插件，提供编译插件，提供测试插件，提供打包插件，提供部署插件，xml是描述性标记，根据需要选择，新手用maven节约时间
DSL解决特定领域问题的语言 DPL通用性语言

maven安装
idea自带maven工具



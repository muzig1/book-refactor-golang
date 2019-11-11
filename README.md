# book-refactor-golang

将[《重构》](https://baike.baidu.com/item/%E9%87%8D%E6%9E%84/2182519?fr=aladdin)一书,原本用JAVA编写的代码示例,按照小生的理解编写为GO语言版本,

## 安装

> 保存到本地

```bash
git clone https://github.com/nxlsBoy/book-refactor-golang.git
```

> PDF下载

建议购买正版图书,支持作者

[重构PDF](https://github.com/nxlsBoy/book-refactor-golang/raw/master/%E9%87%8D%E6%9E%84_%E7%AC%AC%E4%B8%80%E7%89%88.pdf)

## 查阅方式

master作为最基础的分支,不包含任何代码;其他章节都由master分支切出,每个章节内容独立,互不影响

> 如何对比了解代码

举例说明，若要查看第一章的相关内容，切入第一章的分支，然后查看提交日志，每一次的提交即是书中所描述的每步的细节操作

```bash
git checkout chapter1 # 切换到第一章分支
git log -p  # 对比提交变化，了解重构步骤
```

## 其他

希望有理解不到位的地方，还请大佬多多包涵，给予指正，感谢！
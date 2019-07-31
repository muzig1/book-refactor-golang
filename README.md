# book-refactor-golang

看过《重构》这本书的小伙伴肯定知道，这本书的例子是以Java的方式描述的，小生最近准备学习这本书，但最近使用较多的是go，所以就有将书中的内容按照go的思想去理解以及重写，将按照书本的意思以及go语言的特点进行重写，使用Go语言进行举例说明。

## 0x0 准备

> 保存到本地

```bash
git clone https://github.com/nxlsBoy/book-refactor-golang.git
```

## 0x1 更新规划

整个仓库按照**分支（branch）**以及**标签（tag）**进行分类整理；分支以及标签更新的内容，都由**master**分支**check out**出来；保证每个章节更新日志的**简洁**以及**清晰性**。

---

```
\[x]: 代表更新完成
```

分支分为：

- [x] master
- [x] chapter_1:2019年07月30日00:16:14
- [x] chapter_2:2019年07月31日22:25:34
- [] chapter_3
- [] chapter_4
- [] chapter_5
- [] chapter_6
- [] chapter_7
- [] chapter_8
- [] chapter_9
- [] chapter_10

tag分为：

- chapter1
- chapter2
- ...
- chapter15

## 0x2 F&Q

> 如果了解查看代码？

举例说明，若要查看第一章的相关内容，切入第一章的分支，然后查看提交日志，每一次的提交即是书中所描述的每步的细节操作

```bash
git checkout chapter1 # 切换到第一章分支
git log               # 查看第一章重构步骤提交内容
git diff <commit_id>  # 对比提交变化，了解重构步骤
```

## 0x3 其他

初级码农，希望有理解不到位的地方，还请大佬多多包涵，给予指正，感谢！
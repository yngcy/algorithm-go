#  💭💡🎈算法竞赛学习笔记 by yngcy

## 写在前面

本仓库为个人刷题的仓库，主要语言为 C++、Java、 Go、Python 等，题目整体难度不高，偏基础， 
更多内容欢迎访问：[yngcy`s Blog（算法教程持续更新中……）](https://blog.yngcy.com)。



## 语法基础

### Go

推荐 [https://gobyexample-cn.github.io/](https://gobyexample-cn.github.io/) 快速了解 Go 的语法。

## 测试及对拍

`run(io.Reader, io.Writer)` 函数用来处理输入和输出：

- 在 `main` 中调用 `run(os.Stdin, os.Stdout)` 来执行代码；
- 测试时，将测试数据转换成 `string.Reader` 当作输入，并用 `strings.Builder` 来接收输出，然后传给 `run` 方法，实现输出的比较；
- 对拍时需要实现一个暴力算法 `runAC`，参数和 `run` 方法相同。其中通过随机数生成器来生成数据，分别传入 `runAC` 和 `run` 方法，通过对比各自输出，来检查 `run` 中的问题。

**交互题** 的对拍需要把涉及输入输出的地方抽象成接口。

## 参考
1. [灵神的算法模板](https://github.com/EndlessCheng/codeforces-go)
#  💭💡🎈算法竞赛学习笔记 by yngcy

## 写在前面

本仓库为个人刷题的仓库，主要语言为 Go/Java/C++/Python， 
更多内容欢迎访问：[yngcy`s Blog（算法教程挖坑中……）](https://blog.yngcy.com)。

刷题的 OJ 包括：
- [codeforces](https://codeforces.com)：中按照 contest ID 分类
- [力扣](https://leetcode.cn)：按照周赛、双周赛、每日一题以及题单的方式分类

本仓库涵盖以下内容，包括：

- 教程文档，包括算法的核心思想和复杂度
- 参考资料，一些好的博客或资源
- 模板代码，带有注释和使用说明，快速应用
- 练习题，包括 OJ 上的题目以及在项目中的实践

## 语法基础

### Go

- [Go by Example](https://gobyexample-cn.github.io/)：用代码示例解释语法，直观清晰。⭐⭐⭐⭐⭐

## 测试及对拍

`run(io.Reader, io.Writer)` 函数用来处理输入和输出：

- 在 `main` 中调用 `run(os.Stdin, os.Stdout)` 来执行代码；
- 测试时，将测试数据转换成 `string.Reader` 当作输入，并用 `strings.Builder` 来接收输出，然后传给 `run` 方法，实现输出的比较；
- 对拍时需要实现一个暴力算法 `runAC`，参数和 `run` 方法相同。其中通过随机数生成器来生成数据，分别传入 `runAC` 和 `run` 方法，通过对比各自输出，来检查 `run` 中的问题。

**交互题** 的对拍需要把涉及输入输出的地方抽象成接口。

## IDE配置

### GoLand

## 常用工具

### 画图

计算几何、图论或数据结构的画图工具。

- [geogebra](https://www.geogebra.org/classic)：在线的几何画板
- [Geometry Widget](https://csacademy.com/app/geometry_widget/)：用样例绘制几何图形
- [Graph Editor](https://csacademy.com/app/graph_editor/)：用样例绘制树或图

### 计算推导

- [OEIS](https://oeis.org/)：用样例推导数列
- [Wolfram|Alpha](https://www.wolframalpha.com/)：强大的计算工具

### 题单

- [ACDLadders](https://www.acodedaily.com/)：CF 的题单，可按难度和标签进行训练
- [CFTracker](https://cftracker.netlify.app/contests)：CF 比赛题目列表
- [AtCoder Problems](https://kenkoooo.com/atcoder/#/table/)：ATC 的比赛题目列表
- [LC-Rating & Training](https://huxulm.github.io/lc-rating/)：灵神的力扣题单

### 可视化统计

- [Codeforces Visualizer](https://cfviz.netlify.app/)：CF的个人刷题统计
- [Rating converter](https://silverfoxxxy.github.io/rating-converter)：ATC-CF 分数转换器

### 文档

- [StackEdit](https://stackedit.io/app#)：在线 Markdown 编辑器
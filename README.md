# Go-zh



本文档展示了一个简单Go包的开发，并介绍了用go工具来获取、 构建并安装Go包及命令的标准方式。

go 工具需要你按照指定的方式来组织代码。请仔细阅读本文档， 它说明了如何以最简单的方式来准备并运行你的Go安装。

类似的视频讲解可在此处观看。
http://www.youtube.com/watch?v=XCsL89YtqCs

https://go-zh.org/doc/code.html
引言
代码的组织
工作空间
GOPATH 环境变量
包路径
你的第一个程序
你的第一个库
包名
测试
远程包
接下来做什么
获取帮助

#实效Go编程
https://go-zh.org/doc/effective_go.html#%E5%90%8D%E7%A7%B0

#switch 的执行顺序
switch 的条件从上到下的执行，当匹配成功的时候停止。

（例如，

switch i {
case 0:
case f():
}
当 i==0 时不会调用 f 。）
###没有条件的 switch
没有条件的 switch 同 switch true 一样。

这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。

#defer
defer 语句会延迟函数的执行直到上层函数返回。

延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
###defer 栈 （先进后出）
延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

阅读博文了解更多关于 defer 语句的信息。
https://blog.go-zh.org/defer-panic-and-recover

#指针
https://tour.go-zh.org/moretypes/1
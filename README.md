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

指针
Go 具有指针。 指针保存了变量的内存地址。

类型 *T 是指向类型 T 的值的指针。其零值是 nil 。

var p *int
& 符号会生成一个指向其作用对象的指针。

i := 42
p = &i
* 符号表示指针指向的底层的值。

fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
这也就是通常所说的“间接引用”或“非直接引用”。

与 C 不同，Go 没有指针运算。

结构体指针
结构体字段可以通过结构体指针来访问。

通过指针间接的访问是透明的。


package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v //生成指针的用法
	p.X = 1e9
	fmt.Println(v)
	
    p  = &Vertex{1, 2} // 类型为 *Vertex
	fmt.Println(v1, p, v2, v3)

}

###结构体文法
结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。

---
#数组
类型 [n]T 是一个有 n 个类型为 T 的值的数组。

表达式

var a [10]int
定义变量 a 是一个有十个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是一个制约，但是请不要担心； Go 提供了更加便利的方式来使用数组。
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}


##slice
一个 slice 会指向一个序列的值，并且包含了长度信息。

[]T 是一个元素类型为 T 的 slice。

len(s) 返回 slice s 的长度。


package main

import "fmt"


###对 slice 切片
slice 可以重新切片，创建一个新的 slice 值指向相同的数组。



func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])
}

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	//e := b[:cap(b)]
	//printSlice("e", e)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

nil slice
slice 的零值是 nil 。

一个 nil 的 slice 的长度和容量是 0。

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
	
###向 slice 添加元素
向 slice 的末尾添加元素是一种常见的操作，因此 Go 提供了一个内建函数 append 。 内建函数的文档对 append 有详细介绍。

func append(s []T, vs ...T) []T
append 的第一个参数 s 是一个元素类型为 T 的 slice ，其余类型为 T 的值将会附加到该 slice 的末尾。

append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。

如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。

（了解更多关于 slice 的内容，参阅文章Go 切片：用法和本质。）
https://blog.go-zh.org/go-slices-usage-and-internals


#range
for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。

当使用 for 循环遍历一个 slice 时，每次迭代 range 将返回两个值。 第一个是当前下标（序号），第二个是该下标所对应元素的一个拷贝。

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

#map
map 映射键到值。

map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var w map[string]int

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
		test()

}

func test() {
	w = make(map[string]int)
	w["Bell Labs"] = 1234
	fmt.Println(w["Bell Labs"])
}intln(m["Bell Labs"])
}

###map 的文法
map 的文法跟结构体文法相似，不过必须有键名。

修改 map
在 map m 中插入或修改一个元素：

m[key] = elem
获得元素：

elem = m[key]
删除元素：

delete(m, key)
通过双赋值检测某个键存在：

elem, ok = m[key]
如果 key 在 m 中， ok 为 true。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。

同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
练习：
https://tour.go-zh.org/moretypes/20


#函数值
函数也是值。他们可以像其他值一样传递，比如，函数值可以作为函数的参数或者返回值。
package main

import (
	"fmt"
	"math"
)

func compute(aaa func(float64, float64) float64) float64 {
	return aaa(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}



###函数的闭包
Go 函数可以是一个闭包。闭包是一个函数值，它引用了函数体之外的变量。 这个函数可以对这个引用的变量进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

例如，函数 adder 返回一个闭包。每个返回的闭包都被绑定到其各自的 sum 变量上。

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 2; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}



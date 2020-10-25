package main

// 使用分组导入语句是更好的形式
import (
	"fmt" // Go语言标准库中的包
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	// 调用当前包的另一个函数。
	beyondHello()
}

// 函数可以在括号里加参数。
// 如果没有参数的话，也需要一个空括号。
func beyondHello() {
	var x int // 变量声明，变量必须在使用之前声明。
	x = 3     // 变量赋值。
	// 可以用:=来偷懒，它自动把变量类型、声明和赋值都搞定了。
	y := 4
	sum, prod := learnMultiple(x, y)        // 返回多个变量的函数
	fmt.Println("sum:", sum, "prod:", prod) // 简单输出
}

/* <- 快看快看我是跨行注释_(:з」∠)_
Go语言的函数可以有多个参数和 *多个* 返回值。
在这个函数中， `x`、`y` 是参数，且类型为int
`sum`、`prod` 是返回值的标识符（可以理解为名字）且类型为int
*/
func learnMultiple(x, y int) (sum int, prod int) {
	return x + y, x * y // 返回两个值
}

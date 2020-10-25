/**
* @author: JeffreyDin
* @license:
* @contact: dingjianfeng15@gmail.com
* @software:
* @date: 2020/10/26 0:26
* @file: 002functions.go
* @desc：
*/
package main

import "fmt"

// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
// x int, y int 被缩写为 x, y int
//func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

// 多值返回
//func swap(x, y string) (a, b string) {
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
//Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
// 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
//直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}


func main() {
	fmt.Println("第一个函数执行：")
	fmt.Println(add(42, 13))

	fmt.Println("第二个函数执行：")
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Printf("%s %s \n", a, b)

	fmt.Println("第三个函数执行：")
	fmt.Println(split(17))

}

/**
* @author: JeffreyDin
* @license:
* @contact: dingjianfeng15@gmail.com
* @software:
* @date: 2020/10/26 0:23
* @file: 001packages_imports_exported-names.go
* @desc：包_导入方式_导出的名称
 */
package main

import (
	"fmt"
	"math"
	"math/rand" // "math/rand" 包中的源码均以 package rand 语句开始。
)

func main() {
	/*
	在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。
	例如，Pizza 就是个已导出名，Pi 也同样，它导出自 math 包。
	pizza 和 pi 并未以大写字母开头，所以它们是未导出的。
	在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。
	*/
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	//fmt.Println(math.pi)
}

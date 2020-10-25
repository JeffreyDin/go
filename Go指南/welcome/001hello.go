// 导入包的子句在每个源文件的开头。
// Main比较特殊，它用来声明可执行文件，而不是一个库。
package main

// Import语句声明了当前文件引用的包。导入了fmt包。
import "fmt"

// 函数声明：Main是程序执行的入口。
// Go用了花括号来包住函数体。
func main()  {
	/* <- 快看快看我是跨行注释_(:з」∠)_
	Printf，Print 不会自动换行，Println 会
	Print 打印的每一项之间都没有空行，Println 打印的每一项之间都会有空行
	Printf 是格式化输出，在很多场景下比 Println 更方便
	%d 是占位符，表示数字的十进制表示。Printf 中的占位符与后面的数字变量一一对应
	*/
	a := 10
	b := 20
	c := 30
	// 往标准输出打印一行。
	// 用包名fmt限制打印函数。
	fmt.Printf("Hello, World...")
	fmt.Print("go","python","java...") // gopythonjava...
	fmt.Println("go","python","java...") // go python java...
	fmt.Println("a=", a, ",b=", b, ",c=", c)
	fmt.Printf("a=%d,b=%d,c=%d", a, b, c)
}
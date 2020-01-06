package main

import "fmt"

// Go语言中推荐使用驼峰命名

// 批量声明
var (
	name   string // ""
	age    int    // 0
	isOpen bool   // false
)

//const (
//	n1 = 100
//	n2
//	n3
//)

// 常量`n1、n2、n3`的值都是`100`

// iota
const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)

// iota声明中间插队

//const (
//		n1 = iota //0
//		n2 = 100  //100
//		n3 = iota //2
//		n4        //3
//	)
//	const n5 = iota //0

func main() {
	name = "张三"
	age = 27
	isOpen = true
	// Go语言中声明必须使用，不使用就编译不通过
	fmt.Println("----变量----")
	fmt.Print(isOpen) // 在终端中输出要打印的内容
	fmt.Println("----不换行----")
	fmt.Printf("name:%s", name) // %s 占位符 使用name这个变量的值去替换占位符
	fmt.Println("----不换行----")
	fmt.Println(age) // 打印完指定的内容之后会在后面加一个换行符
	fmt.Println()
	fmt.Println("----常量----")
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
}

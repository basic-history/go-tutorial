package main

import (
	"flag"
	"fmt"
)

// 主要分为  类型指针、切片

func main() {

	var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p\n", &cat, &str) // 输出的是这样的地址 0xc0000a0068  0xc0000881e0

	var hello string = "hello world"
	ptr := &hello
	// 打印ptr的类型
	fmt.Printf("内存地址 %p\n", ptr)

	//对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("%T\n", value) //string

	// &取指针 和 *取指针对应的的值  是一对互补操作

	// 也可以这样创建指针
	point := new(bool)
	*point = false //然后给指针赋值

	// 演示使用指针交换两个变量
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

	//演示命令行参数 可以在终端执行  go run 04_指针.go --mode=test 输出的就是test
	mode := flag.String("mode", "default value", "process mode")
	flag.Parse()

	fmt.Println(*mode) //default value
}

// 交换函数，参数都为指针
func swap(a *int, b *int) {
	t := *a //将指针a的值赋给临时变量
	*a = *b //取b指针的值，赋给a指针指向的变量，此时a已经变成b的值
	*b = t  //将a的旧值赋值给b的变量
}

package main

import "fmt" //导包

// 声明变量的几种方式，它们都有默认值，切片/函数/指针的默认值为nil，其它的为0/空字符串/false

var a int //整数类型
var b string
var c []float32   //切片类型
var d func() bool //用于回调函数
var e struct {    //结构体
	x int
}

var (
	a_ int //整数类型
	b_ string
	c_ []float32   //切片类型
	d_ func() bool //用于回调函数
	e_ struct {    //结构体
		x int
	}
)

// 标准格式
//var 变量名 类型 = 表达式
var attack = 40
var defence = 20
var damageRate float32 = 0.17
var damage = float32(attack-defence) * damageRate //需要强转

//也可以自动推导
var hp = 100

func main() {
	fmt.Println(damage) //3.4
	// 更精简的写法，这里如果不使用这个值就会编译错误，而且存在作用域的问题，比如写在外面就会访问不到；因此，短变量声明初始化时使用的较多
	age := 10
	fmt.Println(age) // 10
	name, age := 1, 2
	fmt.Println(name, age) // 1 2

	//交换变量的值
	name, age = age, name
	fmt.Println(name, age) // 2 1

	//匿名变量，不被使用也不会编译报错
	greet, _ := "hello", "world"
	fmt.Println(greet)
}

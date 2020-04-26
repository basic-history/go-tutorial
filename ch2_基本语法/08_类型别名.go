package main

import (
	"fmt"
	"time"
)

type MyDuration time.Duration

func (m MyDuration) EasySet() {

}

func main() {
	// 写法 type TypeAlias = Type
	type NewInt = int

	type NewIntType int //这是定义类型，和上面只有 =的区别

	var a NewInt

	fmt.Printf("a的类型 %T", a) // a的类型 int

}

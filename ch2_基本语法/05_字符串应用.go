package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	tip := "hello world,"
	fmt.Println(len(tip)) // 11 ASCII字符串

	tip2 := "你好，"
	fmt.Println(len(tip2)) //69 每个中文使用UTF8编码，占据3个字节

	fmt.Println(utf8.RuneCountInString(tip2)) //3 可以使用这种方式来统计中文字符

	//遍历
	for _, s := range tip2 {
		fmt.Printf("%c", s) //你好，
	}

	// 正向搜索，可以使用切片
	tip3 := "我是伟大的，中国人"
	index := strings.Index(tip3, "，")       //15
	pos := strings.Index(tip3[index:], "国") //  从，中国人中获取国的索引
	fmt.Println(pos)                        // 6

	//修改字符串
	angle := "here never die"
	angleBytes := []byte(angle)
	for i := 5; i < 10; i++ {
		angleBytes[i] = ' '
	}
	fmt.Println(string(angleBytes)) //here       die

	// 连接字符串，可以使用+的形式，也可以使用类似Stringbuilder

	left := "左侧"
	right := "，右侧"
	fmt.Println(left + right) //左侧，右侧

	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(left)
	stringBuilder.WriteString(right)
	fmt.Println(stringBuilder.String()) //左侧，右侧

}

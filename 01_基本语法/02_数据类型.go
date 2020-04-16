package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

// 整型
// 按长度分 int8、int16、int32、int64
// 还有对应的无符号整型 uint8、uint16、uint32、uint64
// 其中 uint8就是byte，int16对应C语言的short，int64对应C语言long

// 还有自动匹配类型的版本  int 和 uint，在二进制传输，读写文件的结构描述时不要使用这种版本

func main() {
	// 浮点型 注意，这是 Printf
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	name := "pleuvoir"
	fmt.Println(name)
	fmt.Println(name[0:]) // pleuvoir

	text := `一种预格式化的写法，类似html的<pre>标签
			大段文本`
	fmt.Println(text)
	// bool
	flag := true
	fmt.Println(flag)

	//切片(可以认为是list)，能动态分配的空间
	a := make([]int, 3)
	a[0] = 1
	a[1] = 3
	a[2] = 9

	for index, item := range a {
		fmt.Println(index)
		fmt.Println(item)
		fmt.Println("***")
	}

	renderImage()

}

// 画图
func renderImage() {
	const size = 300
	//根据给定大小创建灰度图 宽300长300
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//将每一个像素的灰度设置为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y), color.Gray{0})
	}
	//创建文件
	file, err := os.Create("gray.png")

	if err != nil {
		log.Fatal(err)
	}
	//使用PNG格式写入
	png.Encode(file, pic)
	file.Close()

}

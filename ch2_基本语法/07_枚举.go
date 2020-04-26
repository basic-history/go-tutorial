package main

import "fmt"

// 模拟实现枚举

type Weapon int

const (
	Arrow Weapon = iota //下面的会自动推导
	Shuriken
	Bolower
)

type ChipType int //芯片类型，无参枚举

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU啊"
	}
	return "N/A"
}

func main() {
	fmt.Println(Shuriken) //1
	fmt.Println(Bolower)  //2

	//var weapon Weapon = Bolower
	weapon := Bolower
	fmt.Println(weapon) //2

	fmt.Println(GPU) //GPU啊

}

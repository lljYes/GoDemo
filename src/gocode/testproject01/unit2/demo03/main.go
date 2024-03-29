package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//定义一个整数类型：
	var num1 int8 = 127
	fmt.Println(num1)

	var num2 uint8 = 255
	fmt.Println(num2)

	var num3 = 28
	//Printf函数的作用就是：格式化的，把num3的类型填充到%T的位置上
	fmt.Printf("num3的类型是：%T", num3) //num3的类型是：int
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num3))
}

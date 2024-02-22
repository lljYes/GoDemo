package main //声明文件所在的包，每个go文件必须有归属的包

import "fmt" //引入程序中需要用的包，为了使用包下的函数比如：Println

func main() { //main 主函数程序的入口
	fmt.Println("hello word!") //在控制台打印输出一句话，双引号中的内容会原样输出
}

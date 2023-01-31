// main函数所在包
package main

// 导入 fmt 模块
import "fmt"

// main函数，程序入口
func main() {
	// 这里是单行注释
	/* 多行注释 */

	// 打印函数，结尾处不换行
	fmt.Print("Hello Go!!!")

	// 结尾换行打印
	fmt.Println("换行打印！")

	//也可以加转义字符 \n
	fmt.Print("Hello Go!!!\n")
}

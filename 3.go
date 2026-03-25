package main

import "fmt"

func main() {
	var name string
	var number string
	var height float64

	fmt.Print("请输入姓名、学号和身高（用空格分隔，按回车结束）：")
	
	_, err := fmt.Scanf("%s %s %f", &name, &number, &height)
	if err != nil {
		fmt.Printf("输入错误：%v\n", err)
		return
	}

	
	fmt.Printf("\n📋 输入信息汇总：\n")
	fmt.Printf("姓名：%s\n", name)
	fmt.Printf("学号：%s\n", number)
	fmt.Printf("身高：%.2f 米\n", height)
}
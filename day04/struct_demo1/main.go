package main

import "fmt"

// 结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个person类型的变量p
	var p person
	// 通过字段赋值
	p.name = "周琳"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Printf("%T\n", p)

	var p2 person
	p2.name = "理想"
	p2.age = 18
	fmt.Printf("type:%T value:%v\n", p2, p2)

	// 匿名结构体 多用于临时场景
	var s struct {
		name string
		age  int
	}
	s.name = "嘿嘿嘿"
	s.age = 100
	fmt.Println(s)
}

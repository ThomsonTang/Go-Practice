package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

//比较两个人的年龄，返回年龄大的那个人，并返回年龄差
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	//定义一个person变量，并赋值初始化
	var tom person
	tom.name, tom.age = "Tom", 28

	//两个字段都写清晰的赋值方式
	lisa := person{age: 25, name: "Lisa"}

	//按照person struct定义顺序初始化值
	clover := person{"Clover", 31}

	tl_Older, tl_diff := Older(tom, lisa)
	tc_Older, tc_diff := Older(tom, clover)
	lc_Older, lc_diff := Older(lisa, clover)

	fmt.Printf("%s 和 %s 中， %s 大 %d 岁。", tom.name, lisa.name, tl_Older.name, tl_diff)
	fmt.Printf("%s 和 %s 中， %s 大 %d 岁。", tom.name, clover.name, tc_Older.name, tc_diff)
	fmt.Printf("%s 和 %s 中， %s 大 %d 岁。", lisa.name, clover.name, lc_Older.name, lc_diff)
}

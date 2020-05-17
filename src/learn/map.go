package main

import "fmt"

func test() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	ages["alice"] = 31
	ages["charlie"] = 34

	// 创建一个空map
	ages = map[string]int{}

	//使用delete函数从字典中删除一个元素
	delete(ages, "alice")
	// 即使键不存在 删除操作也是安全的
	// map使用给定的键来查找元素
	// 如果对应的元素不存在 则返回类型的零值
	//map不是一个变量 因此不能获取他的地址  原因是可能会重新散列 到新的存储位置 这就就可能导致获取的地址无效
	age, ok := ages["gfgfgfgf"]
	fmt.Println("%v  %v", age, ok)

}

var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func main() {
	test()

}

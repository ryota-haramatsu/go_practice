package testcode

import "fmt"

// 1. 簡単　なんでも使えるinterfaceについて
func main() {
	// 型を自由に使えるinterface
	var i interface{}
	i = 3
	i = 3.8
	i = "fffff"
	type Mytype uint
	i = Mytype(4)
	fmt.Println(i)
}

func checkType(i interface{}) {
	// タイプスイッチ
	switch i.(type) {
	case int:
		fmt.Println("私はintです")
	case float64:
		fmt.Println("私はfloat64です")
	case string:
		fmt.Println("私はstringです")
	default:
		fmt.Println("私はそれ以外です")
	}
}

// 2. 関数をまとめたものとして interfaceを使用できる
// type <名前> interface{
// 	func1()(string)
// 	func2()(int)
// 	func3()(float64)
// 	...
//   }

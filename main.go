package main

import (
	"simple-casdoor/object"
)

func main() {
	// 初始化 xorm 映射
	object.InitAdapter()

	r := NewRouter()

	r.Logger.Fatal(r.Start(":9825"))
}

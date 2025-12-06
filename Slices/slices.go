package main

import "fmt"

func modifyMap(m map[string]int) {
    m["age"] = 18 // 在函数内部修改
    fmt.Println("函数内修改完毕")
}

func main() {
    // 初始化 map
    users := make(map[string]int)
    users["age"] = 10

    fmt.Println("修改前:", users["age"]) // 输出 10

    // 传给函数
    modifyMap(users)

    // 外部的 map 变了！说明是指向同一块内存
    fmt.Println("修改后:", users["age"]) // 输出 18
}
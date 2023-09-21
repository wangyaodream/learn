package main

import "fmt"


func main(){
    // 创建一个长度和容量为5的切片
    slice := make([]string, 5)
    // 创建一个长度为3个元素，容量为5的切片
    slice_2 := make([]int, 3, 5)
    // 使用索引声明切片,空字符串初始化第100个元素
    slice_3 := []string{99: ""}

    // 创建一个长度为3的数组
    array := [3]string{'1', '2', '3'}

    // 创建长度和容量为3的切片
    slice_4 := []int{10, 20, 30}

    // 创建nil切片
    var slice_nil []int
}

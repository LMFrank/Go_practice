package main

import "fmt"

/*
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/

func main() {
	arr1 := []string{"I", "am", "stupid", "and", "weak"}
	for i, v := range arr1 {
		switch v {
		case "stupid":
			arr1[i] = "smart"
		case "weak":
			arr1[i] = "strong"
		}
	}
	fmt.Println(arr1)
}

package main

import "fmt"

func main() {
	fmt.Println("welcome to linear search")
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	linearSearch(items, 86)
}

func linearSearch(dataList []int, k int) {
	for i, ind := range dataList {
		fmt.Println(i)

		if k == ind {
			fmt.Println("key found  :", dataList[i], "at the index of", i)
			break
		}
	}
}

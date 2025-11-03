package main

import (
	"fmt"
)

func main() {
	//var arrayBool [3]bool = [3]bool{true, false, true}
	arrayBool := [3]bool{true, false}
	arrayInt := [3]int{-9, 11, 7}
	for i := 0; i < len(arrayBool); i++ {
		fmt.Println(i, arrayBool[i])
		fmt.Println(i, arrayInt[i])
	}
}

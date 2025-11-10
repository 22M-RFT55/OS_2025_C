package main

import "fmt"

func main() {
	subjects := []string{"Go", "", "Python"}
	subjects[0] = "Go"
	subjects[2] = "Python"

	for _, subject := range subjects {
		fmt.Println(subject)
	}
}

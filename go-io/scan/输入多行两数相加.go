package main

import "fmt"

func main() {
	a, b := 0, 0
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Println(a + b)
		}
	}
}

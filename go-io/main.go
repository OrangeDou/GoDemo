package main

import "fmt"

var (
	n       int
	like    []int
	q       int
	findSet [][]int
)

func main() {
	fmt.Scanf("%d", &n)
	like = append(like, 0)
	for i := 0; i < n; i++ {
		var c int
		fmt.Scanf("%d", &c)
		like = append(like, c)
	}
	fmt.Scanf("%d", &q)
	for j := 0; j < q; j++ {
		var line []int
		for k := 0; k < 3; k++ {
			var d int
			fmt.Scanf("%d", &d)
			line = append(line, d)
		}
		findSet = append(findSet, line)
	}

	// 寻找喜好值
	for l := 0; l < q; l++ {
		k := findSet[l][2]
		left := findSet[l][0]
		right := findSet[l][1]
		curNumber := 0
		for index := left; index < right; index++ {
			if like[index] == k {
				curNumber++
			}
		}
		fmt.Println(curNumber)
	}
}

func finalPositionOfSnake(n int, commands []string) int {
	i, j := 0, 0
	for _, c := range commands {
		switch c {
		case "UP":
			i--
		case "DOWN":
			i++
		case "RIGHT":
			j++
		case "LEFT":
			j--
		}

	}

	return (i * n) + j
}

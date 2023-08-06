package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	a := make([][]int, b)
	for i := range a {
		a[i] = make([]int, i+1)
	}

	for i := 0; i < b; i++ {
		for j := 0; j < i+1; j++ {
			if i == j || j == 0 {
				a[i][j] = 1
			} else {
				a[i][j] = a[i-1][j-1] + a[i-1][j]
			}
		}
	}

	for i := 0; i < b; i++ {
		x := b - i
		for j := 0; j <= b-i; j++ {
			if j == x {
				for k := 0; k <= i; k++ {
					fmt.Print(a[i][k], " ")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//O(n)
//O(1)

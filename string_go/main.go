package main

import "fmt"

func main() {
	var s1 string
	fmt.Scan(&s1)

	count := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] == 'g' {
			if i != len(s1)-1 && s1[i+1] == 'o' {
				count++
			}
		}
	}

	fmt.Println(count)
}

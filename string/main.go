package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	s2 := make([]byte, len(s))
	j := 0
	for i := 0; i < len(s); i++ {
		if i+1 == len(s) {
			s2[i] = s[i]
			break
		}
		if j == 0 {
			s2[i] = s[i]
			s2[i+1] = s[i+1]
			j = 1
			i++
		} else {
			s2[i] = s[i+1]
			s2[i+1] = s[i]
			j = 0
			i++
		}

	}

	fmt.Println(string(s2))
}

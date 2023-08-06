package main

import (
	"errors"
	"fmt"
)

func oddOne(odd []int) int {
	odd_sum := 0
	for i := range odd {
		odd_sum += odd[i]
	}
	return odd_sum
}

func evenOne(even []int) int {
	even_sum := 0
	for i := range even {
		if even[i] >= 5 {
			sum := 0
			mul := even[i] * 2
			for mul != 0 {
				sum += mul % 10
				mul /= 10
			}
			even_sum += sum
		} else {
			even_sum += even[i] * 2
		}
	}
	return even_sum
}

func isValid(a string) (string, error) {
	flag := 0
	for i := 0; i < len(a); i++ {
		switch a[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			flag = 1
		default:
			return "", errors.New("Invalid input")
		}
	}
	if flag == 1 {
		if a[0] == '4' || a[0] == '5' || a[0] == '6' || a[0] == '3' {
			var odd = []int{}
			var even = []int{}
			for i := range a {
				if i%2 == 0 {
					odd = append(odd, int(a[i]-'0'))
				} else {
					even = append(even, int(a[i]-'0'))
				}
			}

			odd_result := oddOne(odd)
			even_result := evenOne(even)

			result2 := odd_result + even_result
			if result2%10 == 0 {
				return "Valid", nil
			} else {
				err := errors.New("error acquired")
				return "", err
			}
		} else {
			err := errors.New("error acquired")
			return "", err
		}
	}
	return "Valid", nil
}

func main() {
	var a string
	fmt.Scan(&a)

	var result string
	var err error

	if len(a) < 13 || len(a) > 16 {
		fmt.Println("Invalid")
	} else {
		result, err = isValid(a)
	}

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

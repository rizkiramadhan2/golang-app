package main

import "fmt"

func main() {

	val := []int{1, 10, 3, 5, 7, 10}
	SumOdd(val)

}

func SumOdd(val []int) {
	total := 0
	for _, v := range val {
		if v%2 == 1 {
			total += v
		}
	}
	fmt.Println(total)
}

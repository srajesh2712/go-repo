package main

import "fmt"

func Multiple3And51(number int) int {
	sum := 0
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func Multiple3And5(number int, stop int) int {

	if number <= stop {
		if number%3 == 0 || number%5 == 0 {
			return Multiple3And5(number, number)
		}
	}
	return 0
}

func main() {
	fmt.Println(Multiple3And5(10, 10))
}

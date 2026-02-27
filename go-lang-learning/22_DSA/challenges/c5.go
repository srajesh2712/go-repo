package main

import "fmt"

func Solution(str string) []string {
	count := len(str)
	start := 0
	end := 2
	ret := make([]string, 0)
	for start < count {

		if start == count-1 {
			ret = append(ret, string(str[start])+"_")
		} else {
			ret = append(ret, str[start:end])
		}
		start = end
		end = end + 2
	}

	return ret
}

func main() {
	fmt.Println(Solution("abcde"))
}

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func Order(sentence string) string {
	var wordsarray = strings.Fields(sentence)
	re := regexp.MustCompile(`\d`)
	sort.Slice(wordsarray, func(i, j int) bool {
		return string(re.Find([]byte(wordsarray[i]))) < string(re.Find([]byte(wordsarray[j])))
	})

	return strings.Join(wordsarray, " ")
}

func main() {
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2"))

}

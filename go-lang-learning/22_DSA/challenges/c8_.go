package main

import (
	"strconv"
	"strings"
)
import "fmt"

func main() {
	fmt.Println(Is_valid_ip("177.65.-0.220"))
}

func Is_valid_ip(ip string) bool {
	var iparray []string = strings.Split(ip, ".")
	if len(iparray) != 4 {
		return false
	}
	for i := 0; i < len(iparray); i++ {
		if strings.HasPrefix(iparray[i], "-") {
			return false
		}
		val, err := strconv.Atoi(iparray[i])
		if err != nil {
			return false
		}
		if val < 0 || val > 255 {
			return false
		}
		if len(iparray[i]) > 1 && strings.HasPrefix(iparray[i], "0") {
			return false
		}
	}
	return true
}

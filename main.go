package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main()  {
	isValid := validate("353011133330000")
	fmt.Println(isValid)
}

func validate (cn string) bool {
	var ss string
	for _, r := range cn {
		if !unicode.IsSpace(r) {
			ss += string(r)
		}
	}
	var sum int64 = 0
	parity := len(ss) % 2

	cardNumWithoutChecksum := ss[:len(ss)-1]

	for i, v := range cardNumWithoutChecksum {
		item, err := strconv.Atoi(string(v))

		if err != nil {
			fmt.Println(err)
			return false
		}
		if int64(i)%2 != int64(parity) {
			sum += int64(item)
		} else if item > 4 {
			sum += int64(2 * item - 9)
		} else {
			sum += int64(2 * item)
		}
	}

	checkDigit, err := strconv.Atoi(ss[len(ss)-1:])

	if err != nil {
		fmt.Println(err)
		return false
	}
	SumMod := sum%10

	if SumMod == int64(0) {
		return SumMod == int64(checkDigit)
	}
	return int64(10)-SumMod == int64(checkDigit)
}
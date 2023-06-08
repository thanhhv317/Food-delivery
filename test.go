package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	badge := "B1"
	/*
		lastElement := arr[len(arr)-1]
			num, err := strconv.Atoi(lastElement)
			if err == nil {
				arr[len(arr)-1] = strconv.Itoa(num)
			}
	*/

	arr := strings.Split(badge, "")

	second := arr[1:]
	num2, _ := strconv.Atoi(strings.Join(second, ""))

	fmt.Println((int(badge[0])-66)*12 + num2)
}

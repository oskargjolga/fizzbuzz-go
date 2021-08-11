package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(i int) string {
	var retval string = ""
	if i%3 == 0 {
		retval = retval + "fizz"
	}
	if i%5 == 0 {
		retval = retval + "buzz"
	}
	if retval == "" {
		retval = strconv.Itoa(i)
	}
	return retval
}

func main() {
	i := 1
	for i <= 100 {
		fmt.Println(FizzBuzz(i))
		i++
	}

}

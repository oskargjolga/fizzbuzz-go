package fizzbuzzgo

func FizzBuzz(i int) string {
	var retval string = ""
	if i%3 == 0 {
		retval = retval + "fizz"
	}
	if i%6 == 0 {
		retval = retval + "buzz"
	}
	return retval
}

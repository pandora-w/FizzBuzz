package FizzBuzz

import "strconv"

func main() {
	FizzBuzz(1)
}

func FizzBuzz(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"

	} else if num%3 == 0 {
		return "Fizz"

	} else if num%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(num)
}

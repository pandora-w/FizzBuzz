package FizzBuzz

import "strconv"

func main() {
	FizzBuzz(1)
}

func FizzBuzz(num int) string {
	result := strconv.Itoa(num)

	if num == 3 {
		result = "Fizz"
	}

	if num == 5 {
		result = "Buzz"
	}

	return result
}

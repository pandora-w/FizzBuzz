package FizzBuzz

func main() {
	FizzBuzz(1)
}

func FizzBuzz(num int) string {
	result := ""

	if num == 1 {
		result = "1"
	}

	if num == 2 {
		result = "2"
	}

	return result
}

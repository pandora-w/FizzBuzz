package FizzBuzz

func main() {
	FizzBuzz(1)
}

func FizzBuzz(num int) string {
	result := ""

	if num == 1 {
		result = "1"
	}

	return result
}

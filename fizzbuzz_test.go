package FizzBuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("If input is 1 then result is 1", func(t *testing.T) {
		input := 1
		want := "1"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
}

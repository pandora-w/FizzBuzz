package FizzBuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := 1
		want := "1"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 2
		want := "2"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 3
		want := "Fizz"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 4
		want := "4"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 5
		want := "Buzz"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 6
		want := "Fizz"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 7
		want := "7"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 8
		want := "8"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 9
		want := "Fizz"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 10
		want := "10"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 11
		want := "11"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 12
		want := "Fizz"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
	t.Run("", func(t *testing.T) {
		input := 13
		want := "13"
		get := FizzBuzz(input)
		if get != want {
			t.Errorf("If input is %d then result is %s, but result is %s", input, want, get)
		}
	})
}

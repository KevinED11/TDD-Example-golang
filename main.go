package main

import (
	"fmt"
	"testing"
)

func TestFacorial(t *testing.T) {
	table := []struct {
		Fact int
		Want int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, v := range table {
		result := Factorial(v.Fact)

		if result != v.Want {
			t.Errorf("Se esperaba %d, se obtuvo: %d", v.Want, result)
		}
	}

}

func Factorial(number int) int {

	if number == 0 || number == 1 {
		return 1
	}

	return number * Factorial(number-1)

}

func main() {
	fmt.Println(Factorial(5))

}

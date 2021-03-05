package core

import (
	"errors"
	"strconv"
)

// The logic behind fizzbuzz
func FizzBuzz(string1 string, string2 string, int1 int, int2 int, limit int) (result []string, err error) {
	//params check
	if int1 == int2 {
		return nil, errors.New("InvalidParams: The 2 multiples should be different")
	}
	if limit < 1 {
		return nil, errors.New("InvalidParams: Limit should be greater than 1")
	}
	if string1 == string2 {
		return nil, errors.New("InvalidParams: The 2 strings should be different")
	}

	for i := 1; i <= limit; i++ {

		if i%int1 == 0 && i%int2 == 0 {
			result = append(result, string1+string2)
		} else if i%int1 == 0 {
			// Multiple of int1
			result = append(result, string1)
		} else if i%int2 == 0 {
			// Multiple of int2
			result = append(result, string2)
		} else {
			// None of them, so print the number itself
			result = append(result, strconv.Itoa(i))
		}
	}
	return result, nil
}

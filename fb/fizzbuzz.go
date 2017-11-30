package fb

import "strconv"

func Fizzbuzz(number int) string {

	if (number%5 == 0) && (number%3 == 0) {
		return "FizzBuzz"
	}

	if number%3 == 0 {
			return "Fizz"
			}

	if number%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(number)

}


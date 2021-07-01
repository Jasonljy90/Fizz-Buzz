package main

import "strconv"

func fizzBuzz(n int) []string {
	answer := make([]string, n)

	for i := 1; i <= n; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			answer[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			answer[i-1] = "Fizz"
		} else if i%5 == 0 {
			answer[i-1] = "Buzz"
		} else {
			t := strconv.Itoa(i)
			answer[i-1] = t
		}
	}
	return answer
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Buzzer implements the popular fizzbuzz game.
type Buzzer struct{}

// Play returns the result of playing fizzBuzz numTries times.
func (b *Buzzer) Play(numTries int) string {

	var res []string

	for i := 1; i <= numTries; i++ {
		r := ""

		if i%3 == 0 {
			r = "fizz"
		}

		if i%5 == 0 {
			r += "buzz"
		}

		if r == "" {
			r = strconv.Itoa(i)
		}
		res = append(res, r)
	}
	return strings.Join(res, " ")
}

func main() {
	b := &Buzzer{}
	fmt.Println(b.Play(20))
}

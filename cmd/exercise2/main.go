package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Buzzer implements the popular fizzbuzz game.
type Buzzer struct{}

// Play returns the result of playing fizzBuzz numTries times, with a
// modification to output 'lucky' for integers containing a 3
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

		rs := strconv.Itoa(i)

		switch {
		case strings.Contains(rs, "3"):
			r = "lucky"
		case r == "":
			r = rs
		}

		res = append(res, r)
	}
	return strings.Join(res, " ")
}

func main() {
	b := &Buzzer{}
	r := b.Play(20)
	fmt.Println(r)
}

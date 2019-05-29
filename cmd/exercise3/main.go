package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Stats represents the frequency of different results in a game of fizzbuzz.
type Stats struct {
	NumFizz, NumBuzz, NumFizzBuzz, NumLucky, NumInt int
}

// Report implements interface Stringer.
func (s *Stats) String() string {
	return fmt.Sprintf("fizz: %d buzz: %d fizzbuzz: %d lucky: %d integer: %d",
		s.NumFizz, s.NumBuzz, s.NumFizzBuzz, s.NumLucky, s.NumInt)
}

// Buzzer implements the popular fizzbuzz game.
type Buzzer struct {
}

// play returns the result of playing the lucky version of fizzbuzz, along with
// play statistics.
func (b *Buzzer) play(numTries int) (res []string, s Stats) {

	for i := 1; i <= numTries; i++ {
		r := ""
		fizzDone := false
		buzzDone := false
		luckyDone := false

		if i%3 == 0 {
			r = "fizz"
			fizzDone = true
		}

		if i%5 == 0 {
			r += "buzz"
			buzzDone = true
		}

		rs := strconv.Itoa(i)

		switch {
		case strings.Contains(rs, "3"):
			r = "lucky"
			luckyDone = true
		case r == "":
			r = rs
			s.NumInt++
		}

		switch {
		case luckyDone:
			s.NumLucky++
		case fizzDone && buzzDone:
			s.NumFizzBuzz++
		case fizzDone:
			s.NumFizz++
		case buzzDone:
			s.NumBuzz++
		}

		res = append(res, r)
	}
	return
}

// PlayAndReport returns a string representing the output of a game of fizzbuzz
// and formatted stats about the frequency of each result.
func (b *Buzzer) PlayAndReport(numTries int) string {
	r, s := b.play(numTries)
	rs := strings.Join(r, " ")
	return fmt.Sprintf("%s %s", rs, s.String())
}

func main() {
	b := &Buzzer{}
	r := b.PlayAndReport(20)
	fmt.Println(r)
}

package main

import (
	"testing"
)

func TestPlay(t *testing.T) {
	exp := `1 2 lucky 4 buzz fizz 7 8 fizz buzz 11 fizz lucky 14 fizzbuzz 16 17 fizz 19 buzz fizz: 4 buzz: 3 fizzbuzz: 1 lucky: 2 integer: 10`

	b := Buzzer{}
	r := b.PlayAndReport(20)

	if exp != r {
		t.Errorf("\nGot---: %s\n wanted %s", r, exp)
	}
}

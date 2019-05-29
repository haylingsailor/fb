package main

import (
	"testing"
)

func TestPlay(t *testing.T) {
	exp := `1 2 lucky 4 buzz fizz 7 8 fizz buzz 11 fizz lucky 14 fizzbuzz 16 17 fizz 19 buzz`

	b := Buzzer{}
	r := b.Play(20)

	if exp != r {
		t.Errorf("Got: %s, wanted %s", r, exp)
	}
}

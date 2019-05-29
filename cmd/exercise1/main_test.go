package main

import (
	"testing"
)

func TestPlay(t *testing.T) {
	exp := `1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16 17 fizz 19 buzz`

	b := Buzzer{}
	r := b.Play(20)

	if exp != r {
		t.Errorf("Got: %s, wanted %s", r, exp)
	}
}

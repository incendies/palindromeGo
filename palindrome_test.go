package main

import "testing"

func TestPalindrome(t *testing.T) {
	resultRight := isReverse("civic")
	resultWrong := isReverse("emre")

	if !resultRight || resultWrong {
		t.Error("incorrect")
	}
}

func TestReverse(t *testing.T) {
	s := "emre"
	reverseWord := "erme"
	result := Reverse(s)

	if result != reverseWord {
		t.Errorf("if incorrect: %v, give this:  %v",
			result, reverseWord)
	}
}

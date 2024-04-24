package main

import "testing"

func TestValidateGuess(t *testing.T) {
	cases := []struct {
		guess   string
		size    int
		valid   bool
		message string
	}{
		{"123", 3, true, ""},
		{"abc", 3, false, "Must enter a valid number of size 3"},
		{"1234", 3, false, "Must enter a valid number of size 3"},
		{"12", 3, false, "Must enter a valid number of size 3"},
	}

	for _, tc := range cases {
		valid, msg := validateGuess(tc.guess, tc.size)
		if valid != tc.valid {
			t.Errorf("Expected %v, got %v for guess '%v'", tc.valid, valid, tc.guess)
		}
		if msg != tc.message {
			t.Errorf("Expected message '%v', got '%v' for guess '%v'", tc.message, msg, tc.guess)
		}
	}
}

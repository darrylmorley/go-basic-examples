package main

import "testing"

// Test example using a struct
var tests = []struct {
	name     string
	dividend float64
	divisor  float64
	expected float64
	isError  bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isError {
			if err == nil {
				t.Error("expected error but none thrown")
			}
		} else {
			if err != nil {
				t.Error("got unexpected error", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

// Basic tests example
func TestDivide(t *testing.T) {
	_, err := divide(10, 1)
	if err != nil {
		t.Error("Got an error")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Error("Did not get an error, whe we should have")
	}
}

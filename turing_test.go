package turing

// No, this is not the actual 'turing test' ;)

import (
	"testing"
)

// ======
// Some instruction sets to be turned into test cases:
// ======

// From the current cell, sets that cell and the two immediately to the left, to be filled
var addition_states = map[string]State{
	"q1": State{Blank: Change{RIGHT, "q2"}, Filled: Change{ERASE, "q1"}},
	"q2": State{Blank: Change{WRITE, "q3"}, Filled: Change{RIGHT, "q2"}},
	"q3": State{Blank: Change{RIGHT, "q4"}, Filled: Change{LEFT, "q3"}},
}

// Given a tape with two blocks filled, separated with a space, adds them together
var three_states = map[string]State{
	"q1": State{Blank: Change{WRITE, "q1"}, Filled: Change{LEFT, "q2"}},
	"q2": State{Blank: Change{WRITE, "q2"}, Filled: Change{LEFT, "q3"}},
	"q3": State{Blank: Change{WRITE, "q3"}, Filled: Change{HALT, ""}},
}

var turingTests = []struct {
	States   map[string]State
	Tape     string
	Start    int
	State    string
	Expected string
}{
	{addition_states, "110111", 0, "q1", "11111"},
	{addition_states, "11110111", 0, "q1", "1111111"},
	{addition_states, "1101", 0, "q1", "111"},
	{addition_states, "101", 0, "q1", "11"},
	{three_states, "0", 0, "q1", "111"},
	{three_states, "111", 0, "q1", "11111"},
	{three_states, "1", 0, "q1", "111"},
	{three_states, "010101", 0, "q1", "11110101"},
	{three_states, "110101", 0, "q1", "11110101"},
}

func TestStatest(t *testing.T) {
	for _, tt := range turingTests {
		m := NewMachine(tt.Tape, tt.Start, tt.States, tt.State)
		m.Run()

		if m.String() != tt.Expected {
			t.Errorf("Expected %s, but tape was %s", tt.Expected, m.String())
		}
	}
}

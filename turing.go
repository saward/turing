package turing

// Go based Turing Machine, by Mark Saward

import (
	"fmt"
	"strings"
)

const (
	HALT = iota
	LEFT
	RIGHT
	ERASE
	WRITE
)

type State struct {
	// 'Change' is the action to take, given that the cell is either Blank or Filled.
	Blank  Change
	Filled Change
}

func (s State) GetChange(cell Cell) Change {
	if cell {
		return s.Filled
	} else {
		return s.Blank
	}
}

type Change struct {
	Action int    // What action to take
	State  string // What state to put the machine in after doing this.  Empty if action is HALT
}

type Cell bool

type Machine struct {
	Tape     []Cell
	position int
	States   map[string]State
	State    string
}

func (m Machine) CurrentCell() Cell {
	return m.Tape[m.position]
}

// Move the machine needle left
func (m *Machine) Left() {
	fmt.Printf(">> LEFT\n")
	// If we're at the 0th position, then we need to expand our tape array:
	if m.position == 0 {
		size := len(m.Tape)
		m.Tape = append(make([]Cell, size), m.Tape...)
		m.position += size
	}

	m.position -= 1
}

// Move the machine needle right
func (m *Machine) Right() {
	fmt.Printf(">> RIGHT\n")
	// If we're at the last position, then we need to expand our tape array:
	if m.position == (len(m.Tape) - 1) {
		size := len(m.Tape)
		m.Tape = append(m.Tape, make([]Cell, size)...)
	}

	m.position += 1
}

func (m *Machine) Erase() {
	fmt.Printf(">> ERASE\n")
	m.Tape[m.position] = false
}

func (m *Machine) Write() {
	fmt.Printf(">> WRITE\n")
	m.Tape[m.position] = true
}

// Create a new machine, setting the initial tape state and starting position
func NewMachine(tape string, start int, states map[string]State, state string) Machine {
	if len(tape) == 0 {
		panic("Tape length cannot be 0")
	}
	m := Machine{
		Tape:     make([]Cell, len(tape)),
		position: start,
		States:   states,
		State:    state,
	}

	for i, c := range tape {
		if string(c) == "1" {
			m.Tape[i] = true
		}
	}

	return m
}

func (m Machine) Print() {
	var index string
	var tape string
	for i, c := range m.Tape {
		if i == m.position {
			index += "*"
		} else {
			index += " "
		}

		if c {
			tape += "1"
		} else {
			tape += "0"
		}
	}

	fmt.Println(index)
	fmt.Println(tape)
}

// Returns the tape, with all leading and trailing 0's removed.  So for 001010, it returns 101
func (m Machine) String() string {
	var tape string
	var first_found bool

	for _, c := range m.Tape {

		if c {
			tape += "1"
			first_found = true
		} else {
			if first_found {
				tape += "0"
			}
		}
	}

	tape = strings.TrimRight(tape, "0")

	return tape
}

func (m *Machine) Run() {
	halted := false

	fmt.Println("Initial state:")
	m.Print()

	fmt.Println("Starting program:")

	for !halted {

		i := m.States[m.State]
		c := i.GetChange(m.CurrentCell())

		switch c.Action {
		case HALT:
			halted = true
			fmt.Print(">> HALT\n")
		case LEFT:
			m.Left()
		case RIGHT:
			m.Right()
		case ERASE:
			m.Erase()
		case WRITE:
			m.Write()
		}

		m.Print()
		m.State = c.State
	}

	fmt.Print("End result:\n")
	m.Print()
}

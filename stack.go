package main

// stack allows for up to 16 levels of nested subroutines
type stack16Level [16]uint16

var (
	stackPointer uint8
)

// TODO: throw error if stack overflow!!
func (s stack16Level) push() {
	s[stackPointer] = pc + 1 // ?
	stackPointer++
}

// TODO: throw error if can't go further
func (s stack16Level) pop() {
	stackPointer--
	pc = s[stackPointer]
}

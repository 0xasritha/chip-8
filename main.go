package main

import (
	"os"
)

type memoryAddress uint16 // TODO: use this?

const (
	displayPixelOff uint32 = 0x00000000
	displayPixelOn  uint32 = 0xFFFFFFFF
)

var (
	// 8 bit registers
	V0, V1, V2, V3, V4, V5, V6, V7, V8, V9, VA, VB, VC, VD, VE, VF uint8
	registers                                                      map[string]uint8 // TODO which one to use?

	// 16 bit registers
	indexRegister uint16 // store memory address
	pc            uint16

	// stack
	stack = stack16Level{}

	// timers
	delayTimer uint8
	soundTimer uint8

	// TODO: make sure its unsigned Int everywhere

	// keys
	inputKeys = map[uint8]bool{
		0x1: false,
		0x2: false, // etc.
		// TODO: type out
	}

	// display # TODO: must initialize with all off w/ for loop
	display [64][32]uint32

	// memory
	memory [4096]uint8
	/*
		memory is divided into:
		1)
		2) 16 built-in characters: 0x050-0x0A0
		3) instructions: 0x200-0xFFF
	*/

	// opcode
	opcode uint16

	// video buffer (#TODO: temporary, replace w/ the bindings)
	video [32][64]int // 64 pixels wide, 32 pixels high
	// TODO: use bool instead of int?
	// TODO: do a test printer? before using graphics?
)

const (
	// memory
	instructionStart = 0x200
	fontStart        = 0x050
)

func main() {
	// TODO: have init function?
	pc = instructionStart // initialize PC to first instruction

	// load 16 built in characters
	// better way to do this?
	for i := 0; i <= 0xF; i++ {
		for offset := 0; offset < 5; offset++ {
			memory[i+offset] = fontSet[uint8(i)][offset]
		}
	}
}

// LoadROM loads ROM into memory
func LoadROM(file string) {
	// go to address instruction start
	rom, err := os.Open(file)
	check(err)
	defer rom.Close()

	tempMemory := make([]uint8, 0xFFF-0x200) // end - start for instructions range in memory

	numOpcodesRead, err := rom.Read(tempMemory)
	check(err)

	// add to that part of the memory
	// TODO: better way to do this? fix var names. or just have giant blob array for the font set?
	for i := 0x200; i < numOpcodesRead; i++ {
		memory[i] = tempMemory[i]
	}

	// seed random generator for `RAND` instruction

}

// TODO: figure out error handling, make separate function?
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: type Debugger struct {}?

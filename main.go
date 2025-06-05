package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand/v2"
	"os"
)

type memoryAddress uint16 // TODO: use this?

const (
	displayPixelOff uint32 = 0x00000000
	displayPixelOn  uint32 = 0xFFFFFFFF
)

var (
	// 8 bit registers
	registers [16]uint8 // mapped from Vx (V0-VF): value it holds
	// VF is special register: used as a flag to hold information about result of operations

	// 16 bit registers
	indexRegister uint16 // (I) store memory address, only lowest 12 bits (rightmost) are used
	pc            uint16

	// stack
	stack = stack16Level{}

	// timers
	delayTimer uint8
	soundTimer uint8

	// TODO: make sure its unsigned Int everywhere

	// keys
	inputKeys = map[uint8]bool{
		0x0: false,
		0x1: false,
		0x2: false,
		0x3: false,
		0x4: false,
		0x5: false,
		0x6: false,
		0x7: false,
		0x8: false,
		0x9: false,
		0xA: false,
		0xB: false,
		0xC: false,
		0xD: false,
		0xE: false,
	}

	// display # TODO: must initialize with all off w/ for loop
	/*
	 (0,0)     (63, 0)
	 (0, 31)   (63, 31)
	*/
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
	// represents current opcode that CPU will execute?
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

func ogmain() {
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
func Cycle() {
}

func GenerateRandNumber() int {
	return rand.IntN(255)
}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				running = false
				break
			}
		}

		sdl.Delay(33)
	}
}

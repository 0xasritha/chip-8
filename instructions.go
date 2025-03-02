package main

// 36 instructions
// all instructions are 2 bytes long, stored MSB first (so first byte of each instruction should be located at an even address)

// sys jumps to machine code routine at nnn
func sys() {

}

// cls clears video display
func cls() {
	for rowNum := range video {
		for indexIntoRow := range rowNum {
			video[rowNum][indexIntoRow] = 0
		}
	}
}

// ret returns from subroutine
func ret() {
	stack.pop()
}

func jp() {

}

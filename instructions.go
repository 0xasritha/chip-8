package main

import "reflect"

// TODO: have consistent naming everywhere (ex. `PC`, nnn, etc.) -> figure out
// 36 instructions
// all instructions are 2 bytes long, stored MSB first (so first byte of each instruction should be located at an even address)

// sys jumps to machine code routine at nnn
func opSYS() {
	// TODO: implement?
}

// opCLS() clears video display
func opCLS() {
	for rowNum := range video {
		for indexIntoRow := range rowNum {
			video[rowNum][indexIntoRow] = 0
		}
	}
}

/*
00EE - RET
opRET() returns from a subroutine
*/
func opRET() {
	stack.pop()
}

/*
1nnn - JP addr
(nnn) is a 12 bit address that is new value of PC
in binary: `0001 nnnn nnnn nnnn`
jp() jumps to location nnn; interpreter sets PC to nnn
*/
func opJPAddr() {
	address := opcode & 0x0FFF // will mask out first 4 bits, leaving only `nnn`
	pc = address

	/*
			ex. $208
			0x 01 00 00 D0
			0x 00 0F 0F 0F
		(opcode)   0001 nnnn nnnn nnnn
		(0x0FFF) & 0000 1111 1111 1111
		--------------------------------
		           0000 nnnn nnnn nnnn  →  Extracted address `nnn`
	*/
}

/*
2nnn - CALL addr
call() calls subroutine at nnn
*/
func opCALLAddr() {
	stackPointer++
	stack[stackPointer] = pc // puts current PC on top of stack
	address := opcode & 0x0FFF
	pc = address // pc is nnn
}

/*
3xkk - SE Vx, byte
skip3xkk() skips next instruction if Vx = kk
*/
func opSEVxByte() {
	// should return error if register contains nothing -> TODO how to handle errors, program have error "opcode" or smthing?

	registerNumber := opcode & 0x0F00 // gets x

	opcodeValueToCheck := opcode & 0x00FF

	if registers[registerNumber] == uint8(opcodeValueToCheck) { // do >> 8u
		pc += 2
	}
}

/*
4xkk - SNE Vx, byte
sne() skips next instruction if Vx != kk
*/
func opSNEVxByte() {
	// should return error if register contains nothing -> TODO how to handle errors, program have error "opcode" or smthing?

	registerNumber := opcode & 0x0F00 // gets x

	opcodeValueToCheck := (opcode & 0x00FF) >> 8

	if registers[registerNumber] != uint8(opcodeValueToCheck) {
		pc += 2
	}
}

/*
5xy0 - SE Vx, Vy
skip5xy0() skips next instruction if Vx = Vy
*/
func opSEVxVy() {
	valueInVx := opcode & 0x0F00
	valueInVy := opcode & 0x00F0

	if valueInVy == valueInVx {
		pc += 2
	}
}

/*
6xkk - LD Vx, byte
opLDVxByte() puts the value kk into register Vx
*/
func opLDVxByte() {
	kk := opcode & 0x00FF
	x := (opcode & 0x0F00) >> 8

	registers[x] = uint8(kk)
}

/*
7xkk - ADD Vx, byte
add() adds values kk to value in register Vx, then stores result in Vx
*/
func opADDVxByte() {
	kk := opcode & 0x00FF
	x := opcode & 0x0F00

	kk += x
	registers[x] = uint8(kk)
}

/*
8xy0 - LD Vx, Vy
opLDVxVy() stores the value in register Vy to Vx
*/
func opLDVxVy() {
	x := opcode & 0x0F00
	y := opcode & 0x00F0

	registers[x] = registers[y]
}

/*
8xy1 - OR Vx, Vy
opORVxVy() performs a bitwise OR on the values of Vx and Vy, then stores result in Vx
*/
func opORVxVy() {
	x := opcode & 0x0F00
	y := opcode & 0x00F0

	value := (x | y) >> 8

	registers[x] = uint8(value)
}

/*
8xy2 - AND Vx, Vy
opANDVxVy() performs a bitwise AND on the values in Vx and Vy, stores result in Vx
*/
func opANDVxVy() {
	x := opcode & 0x0F00
	y := opcode & 0x00F0

	value := (x & y) >> 8

	registers[x] = uint8(value)
}

/*
8xy3 - XOR Vx, Vy
opXORVxVy() performs a XOR on values in Vx and Vy, stores result in Vx
*/
func opXORVxVy() {
	x := opcode & 0x0F00
	y := opcode & 0x00F0

	value := (x ^ y) >> 8

	registers[x] = uint8(value)
}

/*
8xy4 - ADD Vx, Vy
opADDVxVy() sets Vx = Vx + Vy, set VF as carry (1 if result is greater than 8 bits)
*/
func opADDVxVy() {
	x := opcode & 0x0F00
	y := opcode & 0x00F0

	sum := x + y
	if sum > 0xFF { // greater than 255
		registers[0xF] = 1
	} else {
		registers[0xF] = 0
	}

	registers[x] = uint8(sum & 0xFF) // store only lower 8 bits of sum
}

/*
8xy5 - SUB Vx, Vy
opSUBVxVy() sets Vx = Vx - Vy, set VF = NOT borrow
*/
func opSUBVxVy() {
	x := opcode & 0x0F00
	y := opcode & 0x00F0

	if x > y {
		registers[0xF] = 1
	}

	// better name?
	subResult := x - y
	registers[x] = uint8(subResult >> 8)
}

/*
#TODO: actually fr understand this instruction
8xy6 - SHR Vx {, Vy}
({, Vy}) means  that some interpreters ignore Vy, only use Vx
opSHRVx() sets VF to 1 if LSB of Vx is 1, otherwise 0; then divide Vx by 2
(performs bitwise right shift on value in register Vx)
*/
func opSHRVx() {
	x := opcode & 0x0F00 // TODO: also set this equal every time?? or have some other way where they are all shared?

	lsb := x & 1 // if x is odd, lsb = 1; else x = 0
	if lsb == 1 {
		registers[0xF] = 1
	} else {
		registers[0xF] = 0
	}

	registers[x] = registers[x] / 2
}

/*
8xy7 - SUBN Vx Vyu
Set Vx = Vy - Vx, set Vf = NOT borrow

VF: set to 1 when NO borrow (Vy >= Vx), set to 0 if no borrow (Vy < Vx)

in 8 bit unsigned system, numbers range from 0-255 (0x00 - 0xFF). if result from subtraction is negative, system will WRAP around by doing modulo 256 (called an "underflow" or borrow)

	-> explain this more

	ex. Vy - Vx = 50 - 100 = -50 . -50 cannot be represented in 8-bit format, so system will add 256 to wrap around, or -50 + 256 = 206, and bc underflow has occurred, so borrow happen, VF = 0.

borrow means subtraction would have required borrowing from a higher order bit (which does not exist in 8 bit system), so wraps around
*/
func opSUBNVxVy() {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 8

	if x > y { // difference will be negative
		registers[0xF] = 0 // TODO: have consistent 0xF or 0xf everywhere (for other registers as well.
	} else {
		registers[0xF] = 1 // NO borrow
	}

	registers[x] = uint8(y) - uint8(x) // confirm this works in Go w/ the overflow
}

/*
8xyE - SHL Vx {, Vy }
Set Vx = Vx SHL 1 (if MSB of Vx is 1, then Vf is = 1, otherwise 0). then Vx is * 2

E represents MSB of opcode
*/
func opSHLVx() {

	x := uint8(opcode&0x0F00) >> 8
	vx := registers[x]
	vxMSB := vx & 0b10000000

	registers[0xF] = vxMSB

	registers[x] = vx * 2
}

/*
9xy0 - SNE Vx, Vy
opSNEVxVy() skips next instruction if Vx != Vy
*/
func opSNEVxVy() {
	x := uint8(opcode&0x0F00) >> 8
	y := uint8(opcode&0x00F0) >> 8

	if x != y {
		pc += 2
	}
}

/*
Annn - LD I, addr
opLDIAddr() sets value of indexRegister to nnn
*/
func opLDIAddr() {
	nnn := (opcode & 0x0FFF) >> 8
	indexRegister = nnn
}

/*
Bnnn - JP v0, addr
opLDIAddr() sets value of indexRegister to nnn + value of V0
*/
func opJPV0Addr() {
	nnn := opcode & 0x0FFF
	v0 := registers[0]

	pc += nnn + uint16(v0)
}

/*
Cxkk - RND Vx, byte
opRNDVxByte() sets Vx = random byte & kk
*/
func opRNDVxByte() {
	kk := opcode & 0x00FF
	x := (opcode & 0x0F00) >> 8

	randNumber := x & kk

	registers[x] = uint8(randNumber)
}

/*
Dxyn - DRW Vx, Vy, nibble
opDRWVxVyNibble() displays n-byte sprite starting at
*/
func opDRWVxVyNibble() {
	//x := uint8(opcode&0x0F00) >> 8
	//y := uint8(opcode&0x00F0) >> 8
	//n := uint8(opcode&0x000F) >> 8
	//
	// address
	// TODO
}

/*
Ex9E - SKP Vx
opSKPVx() skips next instruction if key with the value of Vx is pressed
*/
func opSKPVx() {
	x := uint8(opcode&0x0F00) >> 8

	if inputKeys[x] { // key w/ value of Vx is pressed
		pc += 2
	}
}

/*
ExA1 - SKNP Vx
opSKNPVx() skips next instruction if key with the value of Vx is not pressed
*/
func opSKNPVx() {
	x := uint8(opcode&0x0F00) >> 8

	if !inputKeys[x] { // key w/ value of Vx is pressed
		pc += 2
	}
}

/*
Fx07 - LD Vx, DT
Set Vx = delay timer value
*/
func opLDVxDT() {
	x := uint8(opcode&0x0F00) >> 8
	registers[x] = delayTimer
}

/*
Fx0A - LD Vx, K
Wait for a key press (all execution stops), store the value of the key in Vx
*/
// TODO: I am not sure if this has the intended functionality
func opLDVxK() {
	// TODO: test this
	x := uint8(opcode&0x0F00) >> 8
	inputKeysSnapshot := inputKeys

	for reflect.DeepEqual(inputKeysSnapshot, inputKeys) {
	} // waits for any key press indefinitely

	// stores value of key in Vx
	pressedKey := getMapDifference(inputKeysSnapshot)
	if pressedKey != 0xFF { // TODO: HANDLE ERROR THROWING
		registers[x] = pressedKey
	}
}

// TODO: check if there is a better implementation, if not move to helper functions
func getMapDifference(inputKeysSnapshot map[uint8]bool) uint8 {
	for snapshotK, snapshotV := range inputKeysSnapshot {
		if snapshotV != inputKeys[snapshotK] {
			return snapshotK
		}
	}
	return 0xFF
}

/*
Fx15 - LD DT, Vx
opLDDTVx() sets delay timer = Vx
*/
func opLDDTVx() {
	x := uint8(opcode&0x0F00) >> 8
	delayTimer = x
}

/*
Fx18 - LD ST, Vx
opLDSTVx() sets sound timer = Vx
*/
func opLDSTVx() {
	x := uint8(opcode&0x0F00) >> 8
	soundTimer = x
}

/*
Fx1E - ADD I, Vx
opADDIVx() sets I = I + Vx
*/
func opADDIVx() {
	x := uint8(opcode&0x0F00) >> 8
	indexRegister += uint16(x) // TODO: fix this boof logic
}

/*
Fx29 - LD F, Vx
opLDFVx() sets location of sprite for digit Vx
*/
func opLDFVx() {
	// TODO
}

/*
Fx33 - LD B, Vx
opLDBVx() stores the BCD representation of Vx in memory locations I, I+1, and I+2
*/
func opLDBVx() {
	// TODO
}

/*
Fx55 - LD [I], Vx
opLDIVx() stores registers V0 - Vx in memory starting at location I
*/
func opLDIVx() {
	for _, v := range registers {
		memory[indexRegister] = v
		indexRegister++
	}
}

/*
Fx65 - LD Vx, [I]
opLDVxI() reads registers V0 - Vx from memory starting at location I
*/
func opLDVxI() {
	for k, _ := range registers {
		registers[k] = memory[indexRegister]
		indexRegister++
	}
}

// dummy function
func opNull() {}

// TODO: are my comments following proper convention

package main

// Function Pointer Table: array of function pointers, index is opcode
/*
	opcodes w/ repeating first digits use secondary tables
	- $0: needs array that can index up to $E + 1
	- $8: needs array that can index up to $E + 1
	- $E: needs array that can index up to $E + 1
	- $F: needs array that can index up to $65 + 1
*/
// TODO: come back fix the Function Pointer table so it looks like how the guy had it
var (
	// call this opcodeTable??
	functionTableOpcodeTable = map[string]func(){
		"0nnn": opSYS,
		"00E0": opCLS,
		"00EE": opRET,
		"1nnn": opJPAddr,
		"2nnn": opCALLAddr,
		"3xkk": opSEVxByte,
		"4xkk": opSNEVxByte,
		"5xy0": opSEVxVy,
		"6xkk": opLDVxByte,
		"7xkk": opADDVxByte,
		"8xy0": opLDVxVy,
		"8xy1": opORVxVy,
		"8xy2": opANDVxVy,
		"8xy3": opXORVxVy,
		"8xy4": opADDVxVy,
		"8xy5": opSUBVxVy,
		"8xy6": opSHRVx,
		"8xy7": opSUBVxVy,
		"8xyE": opSHLVx,
		"9xy0": opSNEVxVy,
		"Annn": opLDIAddr,
		"Bnnn": opJPV0Addr,
		"Cxkk": opRNDVxByte,
		"Dxyn": opDRWVxVyNibble,
		"Ex9E": opSKPVx,
		"ExA1": opSKNPVx,
		"Fx07": opLDVxDT,
		"Fx0A": opLDVxK,
		"Fx15": opLDVxDT,
		"Fx18": opLDSTVx,
		"Fx1E": opADDIVx,
		"Fx29": opLDFVx,
		"Fx33": opLDBVx,
		"Fx55": opLDIVx,
		"Fx65": opLDVxI,
	}

	// TODO: have an any type? or have funcs that become table

	// array of functions where opcode is an index into an array
	functionPointerTable = [0xF]func(){} // size: $0 - $F (so need F + 1)
	// will point to function that indexes based on relevant parts of opcode

	// secondary tables (for first digits that repeat)
	zeroTable  = [0xE]func(){}
	eightTable = [0xE]func(){}
	ETable     = [0xE]func(){}
	tableF     = [0x66]func(){} // needs to index up to $65 + 1
)

func table0() {

}

func initFunctionPointerTable() {

}

/*



 */

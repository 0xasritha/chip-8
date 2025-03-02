# Chip8 Emulator

- Chip-8 emulator: will emulate Chip-8 VM, simple interpreted language to run small games on early microcomputers
  - not actual hardware chip, interpreted language w/ small instruction set
  - 34 opcodes to emulate 

### ROM file
- ROM file: binary that contains copy of game's/system's read only memory (dump of og game's firmware from a real console)
  - emulator reads ROM file into memory, simulates hardware of original console so game can run if it were on og system
  - ROM file will contain machine instructions specific to CPU architecture of system it was made for
    - contains Header, program code, graphics data, audio data (basically the og game code)
  - so Chip-8 ROM file will contain VM code instructions (opcodes) for a single game or program that runs on the CHIP-8 VM  

### Memory 
- 4096 bytes or memory (`0x000-0xFFF`)
- 3 parts: 
    - `0x000-0x1FF`
    - `0x050-0x0A0`: storage space for 16 built in characters (0-F)
    - `0x200-0xFFF`: instructions from ROM, anything left is free to use

#### Built-in font set 
- `0x050-0x0A0`: reserved for built-in font set (characters 0-F)
  - each character is a 5x4 pixel sprite to show hex digits 
  - when CHIP-8 game wants to draw a number (0-F), it doesn't include font data in ROM itself (assumes that font data will be at `0x059-0x0A0`)
  - each digit (0-F) is presented as a 5 byte sprite, each byte represents a row of 8 pixels 

### Registers
- `V8`: holds "flags" w/ information about result of operations  
- **16-bit Index Register**: store memory addresses 
- **16-bit Program Counter**: stores address of next instruction to execute
  - must increment PC before executing instruction because some instructions will manipulate PC to control program flow
  - *pushing*: putting PC on stack 
  - *popping*: popping PC off stack 

- **8-bit Stack Pointer**: points to top of stack (index into array, need "16 indices", hence a single byte)


### Stack 
- in Chip8, stack is ONLY used for function call management, not general purpose value storage
  - stack is used to store return addresses when a subroutine `CALL`(`2NNN`) is called 
  - when `RET` (`00EE`) is executed, PC is restored from the stack
  - the registers `V0`-`VF` are used for storing temporary values, computation, etc.; values can also be stored in memory 
  - Chip-8 does not have push/pop instructions like modern CPUs (hence can't push/pop values) 

- will be represented as an array 
- `CALL` instruction: CPU will begin executing instructions in a different region 
- `RET` instruction: must go back to where it was when it hit the `CALL`
- 16 levels of stack: can hold 16 different PCs (can support 16 nested function calls)
  - pushing: putting PC onto stack 
  - popping: pulling PC off stack

- stack holds PC value when CALL instruction was executed, RET pulls that address from the stack and puts it back into PC so CPU can execute it on the next cycle 

### 8-bit Delay Timer 
- if timer = 0, stays at 0 
- if timer = non-zero value, will decrement at a rate of 60Hz

### 8-bit Sound Timer 
- if timer = non-zero value, will decrement at a rate of 60Hz (but single tone will buzz)

### Input Keys 
- 16 input keys (boolean state)

### 64x32 Monochrome Display Memory 
- `64x32` grid, each pixel has a boolean state
  - `uint32` for pixel value: 
    - `0xFFFFFFFF`: on 
    - `0x00000000`: off

- sprite: two-dimensional image or animation that's used to represent an object on the screen; drawn on top of games' background, moves independently
  - `DRAW` instruction iterates over each pixel in a sprite, XORs the sprite pixel w/ the display pixel 
    - XOR (exclusive OR): outputs 1 only when 2 inputs differ (if inputs are same, it is 0)
    - XORing allows the pixels in a sprite to toggle on and off
      - ex. to move sprite: 1) XOR out the old position w/ same bits to remove it 2) XOR w/ same bits again at the new position
    - given S = Sprite Pixel, D = Display Pixel: 
      - S: OFF, D: OFF; S XOR D = OFF 
      - S: OFF, D: ON; S XOR D = ON 
      - S: ON, D: OFF; S XOR D = ON
      - S: ON, D: ON; S XOR D = OFF

### Function Pointer Table
- to decode opcodes: use array of function pointers, where opcode is index into array of function pointers
  - types of opcodes: 
    - entire opcode is unique 
    - first digit repeats, last digit is unique 
    - first three digits are `$00E`, fourth digit is unique 
    - first digit repeats, last two digits are unique


### Fetch, Decode, Execute
- one cycle in CPU: 
1) fetch next instruction (in opcode format)
2) decode instruction to determine the operation to complete 
3) execute the instruction 


### Platform Layer (`class Platform`) 
- SDL to render and get input
  - SDL_Renderer: 2D GPU acceleration 
  - SDL_Texture: render 2D image 


### Main Loop
- main loop will continuously call a `Cycle` function
- need three CLI arguments: 1) Scale, 2) Delay, 3) ROM file to load 















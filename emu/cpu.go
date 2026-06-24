package emu

import (
	"fmt"
)

type CPU struct {
	A, F byte
	B, C byte
	D, E byte
	H, L byte

	SP uint16
	PC uint16

	IMEDelay bool
	IME      bool
	IF       byte
	IE       byte

	Bus      *Bus
	Halted   bool
	Panicked bool

	StopMode bool

	PPU   *PPU
	Timer *Timer

	Trace     bool
	LogIoRegs bool

	OAM_DMA byte // if 0xFF we ignore

	// ReadsThisCycle int
}

func (c *CPU) Init(trace bool, log_ioregs bool) {
	c.Trace = trace
	c.LogIoRegs = log_ioregs
	c.Bus = &Bus{Cpu: c}
	c.PPU.Cpu = c
	c.Bus.Init()

	c.Timer = &Timer{Cpu: c, timaCyclesLeft: 1024}

	c.OAM_DMA = 0xFF
}

func (c *CPU) AF() uint16 {
	return uint16(c.A)<<8 | uint16(c.F)
}

func (c *CPU) BC() uint16 {
	return uint16(c.B)<<8 | uint16(c.C)
}

func (c *CPU) DE() uint16 {
	return uint16(c.D)<<8 | uint16(c.E)
}

func (c *CPU) HL() uint16 {
	return uint16(c.H)<<8 | uint16(c.L)
}

func (c *CPU) SetAF(val uint16) {
	c.A = byte(val >> 8)
	c.F = byte(val & 0xF0)
}

func (c *CPU) SetBC(val uint16) {
	c.B = byte(val >> 8)
	c.C = byte(val)
}

func (c *CPU) SetDE(val uint16) {
	c.D = byte(val >> 8)
	c.E = byte(val)
}

func (c *CPU) SetHL(val uint16) {
	c.H = byte(val >> 8)
	c.L = byte(val)
}

func (c *CPU) GetFlagZ() bool {
	return c.F&(1<<7) != 0
}

func (c *CPU) GetFlagN() bool {
	return c.F&(1<<6) != 0
}

func (c *CPU) GetFlagH() bool {
	return c.F&(1<<5) != 0
}

func (c *CPU) GetFlagC() bool {
	return c.F&(1<<4) != 0
}

func (c *CPU) SetFlagZ(v bool) {
	if v {
		c.F |= 1 << 7
	} else {
		c.F &^= 1 << 7
	}
}
func (c *CPU) SetFlagN(v bool) {
	if v {
		c.F |= 1 << 6
	} else {
		c.F &^= 1 << 6
	}
}
func (c *CPU) SetFlagH(v bool) {
	if v {
		c.F |= 1 << 5
	} else {
		c.F &^= 1 << 5
	}
}
func (c *CPU) SetFlagC(v bool) {
	if v {
		c.F |= 1 << 4
	} else {
		c.F &^= 1 << 4
	}
}

func (c *CPU) PCRead() byte {
	// c.ReadsThisCycle += 1
	v := c.Bus.Read(c.PC)
	c.PC++
	return v
}

func (c *CPU) PCRead16() uint16 {
	// c.ReadsThisCycle += 2
	v := c.Bus.Read16(c.PC)
	c.PC += 2
	return v
}

func (c *CPU) Push(val uint16) {
	c.SP--
	c.Bus.Write(c.SP, byte(val>>8))
	c.SP--
	c.Bus.Write(c.SP, byte(val&0xFF))
}

func (c *CPU) Pop() uint16 {
	low := c.Bus.Read(c.SP)
	c.SP++
	hi := c.Bus.Read(c.SP)
	c.SP++
	return uint16(hi)<<8 | uint16(low)
}

func (c *CPU) Panic(msg string) {
	c.Panicked = true
	errInfo := fmt.Sprintf(
		"\n[CPU PANIC] %s\n"+
			"PC: 0x%04X\n"+
			"A: 0x%02X | F: 0x%02X\n"+
			"B: 0x%02X | C: 0x%02X\n"+
			"D: 0x%02X | E: 0x%02X\n"+
			"H: 0x%02X | L: 0x%02X\n"+
			"SP: 0x%04X\n",
		msg, c.PC, c.A, c.F, c.B, c.C, c.D, c.E, c.H, c.L, c.SP,
	)
	fmt.Print(errInfo)
}

func (c *CPU) HandleInterrupts() int {
	ifReg := c.IF
	ieReg := c.IE

	pending := ifReg & ieReg
	if pending == 0 {
		return 0
	}

	if !c.IME {
		return 0
	}

	var vector uint16
	var bit byte

	switch {
	case (pending & 0x01) != 0: // VBlank
		vector = 0x0040
		bit = 0
	case (pending & 0x02) != 0: // LCD STAT
		vector = 0x0048
		bit = 1
	case (pending & 0x04) != 0: // Timer
		vector = 0x0050
		bit = 2
	case (pending & 0x08) != 0: // Serial
		vector = 0x0058
		bit = 3
	case (pending & 0x10) != 0: // Joypad
		vector = 0x0060
		bit = 4
	default:
		return 0
	}

	if c.Trace {
		fmt.Printf("Interrupted to %#x!\n", vector)
	}

	c.IME = false

	ifReg &= ^(1 << bit)
	c.IF = ifReg

	c.Push(c.PC)

	c.PC = vector

	return 5

}

func (c *CPU) HandleDMA() int {
	if c.OAM_DMA != 0xFF {
		source := uint16((int(c.OAM_DMA) * int(0x100)))
		src_end := uint16((int(c.OAM_DMA) * int(0x100)) | 0x9F)
		dest := uint16(0xFE00)

		for source <= src_end && dest <= 0xFE9F {
			c.Bus.Write(dest, c.Bus.Read(source))
			source++
			dest++
		}
		c.OAM_DMA = 0xFF
		return 160
	}
	return 0
}

func (c *CPU) Step() int {
	// c.ReadsThisCycle = 0
	if c.StopMode {
		return 1
	}

	c.HandleDMA()

	// if dma_cycles > 0 {
	// 	c.Timer.Step(dma_cycles * 4)
	// 	return dma_cycles
	// }
	var halt_bug bool

	if c.Halted {
		if (c.IF & c.IE & 0x1F) != 0 {
			c.Halted = false

			if !c.IME {
				// panic("halt bug")
				halt_bug = true
			}
		} else {
			c.Timer.Step(4)
			return 1
		}
	}

	intCycles := c.HandleInterrupts()

	if intCycles > 0 {
		c.Timer.Step(intCycles * 4)
		return intCycles
	}

	oldpc := c.PC
	op := c.PCRead()

	var instr Instruction
	if op == 0xCB {
		op = c.PCRead()
		instr = CBOpcodes[op]
		if halt_bug {
			c.PC -= 2
		}
	} else {
		instr = Opcodes[op]
		if halt_bug {
			c.PC--
		}
	}
	if instr.Exec == nil {
		c.Panic("Instruction not implemented")
		return 0
	}
	if c.Trace {
		fmt.Printf("[%#x] Step() op: %#x SP: %#x (%s)\n", oldpc, op, c.SP, instr.Name)
	}
	res := instr.Exec(c)
	// if c.ReadsThisCycle != instr.Bytes {
	// 	c.Panic(fmt.Sprintf("Instruction %s size missmatch! %d != %d", instr.Name, c.ReadsThisCycle, instr.Bytes))
	// }
	if res == 0 && !c.Panicked {
		panic("you forgot")
	}

	if c.IMEDelay && op != 0xFB {
		c.IME = true
		c.IMEDelay = false
	}

	c.Timer.Step(res * 4)
	return res
}

func (c *CPU) Run() {
	for !c.Halted {
		cycles := c.Step()
		c.PPU.Step(cycles)
		// fmt.Printf("instr took %d cycles\n", cycles)
		// time.Sleep(2 * time.Second)
	}
}

type Timer struct {
	DIV  uint16
	TIMA byte
	TMA  byte
	TAC  byte

	timaCyclesLeft int

	Cpu *CPU
}

func (t *Timer) Step(cycles int) {
	t.DIV += uint16(cycles)

	if (t.TAC & 0x04) == 0 {
		return // TIMA disabled
	}

	var threshold int
	switch t.TAC & 0x03 {
	case 0x00:
		threshold = 1024
	case 0x01:
		threshold = 16
	case 0x02:
		threshold = 64
	case 0x03:
		threshold = 256
	}

	t.timaCyclesLeft -= cycles
	for t.timaCyclesLeft <= 0 {
		t.timaCyclesLeft += threshold

		if t.TIMA == 0xFF {
			t.TIMA = t.TMA

			ifReg := t.Cpu.IF
			t.Cpu.IF = ifReg | 0x04
		} else {
			t.TIMA++
		}
	}
}

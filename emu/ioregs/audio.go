package ioregs

import (
	"goboy/emu"
)

type CH1 struct {
	on bool
}

type CH2 struct {
	on bool
}

type CH3 struct {
	on bool
}

type CH4 struct {
	on bool
}

type Audio struct {
	on bool
}

func (r *Audio) OnRead(cpu *emu.CPU, addr uint16) byte {
	// cpu.Halt(fmt.Sprintf("Read from I/O Register %#x not implemented!", addr))
	// fmt.Printf("Audio not implemented!\n")
	return 0xFF
}

func (h *Audio) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	// cpu.Halt(fmt.Sprintf("Write to I/O Register %#x not implemented!", addr))
	// fmt.Printf("Audio not implemented!\n")
}

package ioregs

import (
	"fmt"
	"goboy/emu"
)

type KeyRegs struct{}

func (r *KeyRegs) OnRead(cpu *emu.CPU, addr uint16) byte {
	// cpu.Halt(fmt.Sprintf("Read from I/O Register %#x not implemented!", addr))
	fmt.Printf("KEY0/1 is GBC Only!\n")
	return 0xFF
}

func (r *KeyRegs) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	// cpu.Halt(fmt.Sprintf("Write to I/O Register %#x not implemented!", addr))
	fmt.Printf("KEY0/1 is GBC Only!\n")
}

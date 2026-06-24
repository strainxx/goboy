package ioregs

import (
	"fmt"
	"goboy/emu"
)

type BootRomMapped struct{}

func (r *BootRomMapped) OnRead(cpu *emu.CPU, addr uint16) byte {
	if cpu.Bus.BootRomDisabled {
		return 1
	}
	return 0
}

func (h *BootRomMapped) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	fmt.Printf("Disabled bootrom!\n")
	cpu.Bus.BootRomDisabled = true
}

package ioregs

import (
	"goboy/emu"
)

type InterruptRegs struct{}

func (r *InterruptRegs) OnRead(cpu *emu.CPU, addr uint16) byte {
	switch addr {
	case 0xFF0F:
		return cpu.IF | 0xE0
	}
	return 0x00
}

func (h *InterruptRegs) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	switch addr {
	case 0xFF0F:
		cpu.IF = val & 0x1F
	}
}

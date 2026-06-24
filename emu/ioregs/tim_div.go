package ioregs

import (
	"goboy/emu"
)

type TimDiv struct {
}

func (r *TimDiv) OnRead(cpu *emu.CPU, addr uint16) byte {
	// cpu.Halt(fmt.Sprintf("Read from I/O Register %#x not implemented!", addr))
	switch addr {
	case 0xFF04:
		return byte(cpu.Timer.DIV >> 8)
	case 0xFF05:
		return cpu.Timer.TIMA
	case 0xFF06:
		return cpu.Timer.TMA
	case 0xFF07:
		return cpu.Timer.TAC
	}
	return 0x00
}

func (h *TimDiv) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	switch addr {
	case 0xFF04:
		cpu.Timer.DIV = 0
	case 0xFF06:
		cpu.Timer.TMA = val
	case 0xFF05:
		cpu.Timer.TIMA = val
	case 0xFF07:
		cpu.Timer.TAC = val
	default:
		cpu.Panic("How????")
	}
}

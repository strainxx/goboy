package ioregs

import (
	"goboy/emu"
)

type OAMDMA struct {
}

func (r *OAMDMA) OnRead(cpu *emu.CPU, addr uint16) byte {
	return cpu.OAM_DMA
}

func (r *OAMDMA) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	cpu.OAM_DMA = val
}

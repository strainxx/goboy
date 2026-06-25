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
	NR52 byte // $FF26 Audio master control
}

func (r *Audio) OnRead(cpu *emu.CPU, addr uint16) byte {
	switch addr {
	case 0xFF26:
		return r.NR52
	}
	return 0xFF
}

func (h *Audio) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	// cpu.Halt(fmt.Sprintf("Write to I/O Register %#x not implemented!", addr))
	// fmt.Printf("Audio not implemented!\n")

}

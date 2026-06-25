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
	NR51 byte // $FF25 Sound panning
	NR50 byte // $FF24 Master volume & VIN panning
}

func (r *Audio) Step() {
	if r.NR52&0x80 == 0 {
		return
	}
}

func (r *Audio) OnRead(cpu *emu.CPU, addr uint16) byte {
	switch addr {
	case 0xFF24:
		return r.NR50
	case 0xFF25:
		return r.NR51
	case 0xFF26:
		return r.NR52
	}
	return 0xFF
}

func (r *Audio) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	switch addr {
	case 0xFF24:
		r.NR50 = val
	case 0xFF25:
		r.NR51 = val
	case 0xFF26:
		r.NR52 = val & 0x80
	}

}

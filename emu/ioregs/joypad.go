package ioregs

import (
	"fmt"
	"goboy/emu"
)

type Joypad struct {
	Register              byte
	A, B, Select, Start   bool
	Up, Down, Left, Right bool
}

func (r *Joypad) Read() byte {
	var result byte = 0x0F

	if (r.Register & 0x20) == 0 {
		var buttons byte = 0x0F
		if r.A {
			buttons &= ^byte(0x01)
		}
		if r.B {
			buttons &= ^byte(0x02)
		}
		if r.Select {
			buttons &= ^byte(0x04)
		}
		if r.Start {
			buttons &= ^byte(0x08)
		}

		result &= buttons
	}

	if (r.Register & 0x10) == 0 {
		var directions byte = 0x0F
		if r.Right {
			directions &= ^byte(0x01)
		}
		if r.Left {
			directions &= ^byte(0x02)
		}
		if r.Up {
			directions &= ^byte(0x04)
		}
		if r.Down {
			directions &= ^byte(0x08)
		}

		result &= directions
	}

	return 0xC0 | r.Register | result
}

func (r *Joypad) OnRead(cpu *emu.CPU, addr uint16) byte {
	return r.Read()
}

func (r *Joypad) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	// oldState := r.Read()

	r.Register = val & 0x30

	// newState := r.Read()

	// if (^newState & oldState & 0x0F) != 0 {
	//     cpu.IF |= 0x10
	// }
}

func (r *Joypad) UpdateAll(cpu *emu.CPU, a, b, sel, start, up, down, left, right bool) {
	oldState := r.Read()

	r.A = a
	r.B = b
	r.Select = sel
	r.Start = start
	r.Up = up
	r.Down = down
	r.Left = left
	r.Right = right

	newState := r.Read()

	if (^newState & oldState & 0x0F) != 0 {
		cpu.IF |= 0x10
		if cpu.StopMode {
			fmt.Printf("exiting stop mode\n")
			cpu.StopMode = false
		}
	}
}

package ioregs

import (
	"fmt"
	"goboy/emu"
)

type SerialDataTransfer struct {
	SerialData    byte
	SerialControl byte
}

func (r *SerialDataTransfer) OnRead(cpu *emu.CPU, addr uint16) byte {
	switch addr {
	case 0xFF01:
		return r.SerialData
	case 0xFF02:
		return r.SerialControl
	}
	return 0x00
}

func (r *SerialDataTransfer) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	switch addr {
	case 0xFF01:
		r.SerialData = val
	case 0xFF02:
		r.SerialControl = val
		if (val & 0x80) != 0 {
			fmt.Printf("%c", r.SerialData)
			r.SerialData = 0xFF
			r.SerialControl &= 0x7F

			// cpu.IF |= 0x08
		}
	}
}

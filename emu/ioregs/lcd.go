package ioregs

import (
	"fmt"
	"goboy/emu"
)

type LCD struct {
}

func (r *LCD) OnRead(cpu *emu.CPU, addr uint16) byte {
	switch addr {
	case 0xFF40:
		return cpu.PPU.LCDC
	case 0xFF41:
		return cpu.PPU.STAT | 0x80
	case 0xFF44:
		return byte(cpu.PPU.Scanline)
	case 0xFF45:
		return cpu.PPU.LYC
	case 0xFF4A:
		return cpu.PPU.WY
	case 0xFF4B:
		return cpu.PPU.WX
	default:
		cpu.Panic(fmt.Sprintf("Read from I/O Register %#x not implemented!", addr))
		return 0xFF
	}
}

func (h *LCD) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	switch addr {
	case 0xFF40:
		cpu.PPU.LCDC = val
	case 0xFF41:
		cpu.PPU.STAT = (cpu.PPU.STAT & 0x07) | (val & 0xF8)
	case 0xFF44:
		cpu.PPU.Scanline = 0
	case 0xFF45:
		cpu.PPU.LYC = val
	case 0xFF4A:
		cpu.PPU.WY = val
	case 0xFF4B:
		cpu.PPU.WX = val
	default:
		cpu.Panic(fmt.Sprintf("Write to I/O Register %#x not implemented!", addr))
	}
}

type Scrolling struct {
}

func (r *Scrolling) OnRead(cpu *emu.CPU, addr uint16) byte {
	switch addr {
	case 0xFF42:
		return cpu.PPU.SCY
	case 0xFF43:
		return cpu.PPU.SCX
	default:
		return 0xFF
	}
}

func (h *Scrolling) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	switch addr {
	case 0xFF42:
		cpu.PPU.SCY = val
	case 0xFF43:
		cpu.PPU.SCX = val
	}
}

type Palettes struct {
}

func (r *Palettes) OnRead(cpu *emu.CPU, addr uint16) byte {
	switch addr {
	case 0xFF47:
		return cpu.PPU.BGP
	case 0xFF48:
		return cpu.PPU.OBP0
	case 0xFF49:
		return cpu.PPU.OBP1
	default:
		return 0xFF
	}
}

func (h *Palettes) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	switch addr {
	case 0xFF47:
		cpu.PPU.BGP = val
	case 0xFF48:
		cpu.PPU.OBP0 = val
	case 0xFF49:
		cpu.PPU.OBP1 = val
	}
}

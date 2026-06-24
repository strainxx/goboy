package emu

import (
	"fmt"
	"goboy/cartridge"
	"reflect"
)

type IOHook interface {
	OnWrite(cpu *CPU, addr uint16, val byte)
	OnRead(cpu *CPU, addr uint16) byte
}

type IONotImpl struct{}

func (h *IONotImpl) OnRead(cpu *CPU, addr uint16) byte {
	cpu.Panic(fmt.Sprintf("Read from I/O Register %#x not implemented!", addr))
	return 0
}

func (h *IONotImpl) OnWrite(cpu *CPU, addr uint16, val byte) {
	cpu.Panic(fmt.Sprintf("Write to I/O Register %#x not implemented!", addr))
}

type IOEmpty struct{}

func (h *IOEmpty) OnRead(cpu *CPU, addr uint16) byte {
	return 0xFF
}

func (h *IOEmpty) OnWrite(cpu *CPU, addr uint16, val byte) {
}

type Bus struct {
	Cpu  *CPU
	Cart cartridge.Cartridge

	BootRomDisabled bool
	BootROM         [256]byte
	Memory          [65536]byte
	ioHooks         [128]IOHook
}

func (b *Bus) Init() {
	not_impl := &IONotImpl{}
	for i := 0; i < 128; i++ {
		b.ioHooks[i] = not_impl
	}

	// b.RegisterRange(0xFF4C, 0xFF7F, &IOEmpty{})
	b.RegisterRange(0xFF4C, 0xFF7F, &IOEmpty{})
}

func (b *Bus) RegisterHook(addr uint16, hook IOHook) {
	b.ioHooks[addr-0xFF00] = hook
}

func (b *Bus) RegisterRange(start, end uint16, hook IOHook) {
	for i := start; i <= end; i++ {
		b.RegisterHook(i, hook)
	}
}

func (b *Bus) LogIOReg(reg IOHook, otype string, addr uint16) {
	if addr == 0xFF0F {
		// avoid InterruptRegs spam
		return
	}
	fmt.Printf("[%s] %s -> %#x\n", reflect.TypeOf(reg).String(), otype, addr)
}

func (b *Bus) LoadROM(data []byte, bootrom bool, path string) error {
	if bootrom {
		copy(b.BootROM[:], data)
		return nil
	}
	// copy(b.Memory[:], data)
	cart, err := cartridge.Load(data, path)
	if err != nil {
		return err
	}
	b.Cart = cart
	return nil
}

func (b *Bus) Read(base uint16) byte {
	switch {
	case base < 0x0100:
		if !b.BootRomDisabled {
			return b.BootROM[base]
		}
		return b.Cart.Read(base)
	case base < 0x8000:
		return b.Cart.Read(base)
	case base >= 0x8000 && base <= 0x9FFF:
		return b.Cpu.PPU.VRAM[base-0x8000]
	case base >= 0xA000 && base <= 0xBFFF:
		return b.Cart.Read(base)
	case base >= 0xC000 && base <= 0xDFFF:
		return b.Memory[base]
	case base >= 0xE000 && base <= 0xFDFF:
		return b.Memory[base-0x2000] // echo ram
	case base >= 0xFE00 && base <= 0xFE9F:
		if b.Cpu.PPU.Mode == 2 || b.Cpu.PPU.Mode == 3 {
			// OAM scan or Pixel Transfer
			return 0xFF
		}
		return b.Cpu.PPU.OAM[base-0xFE00]
	case base >= 0xFEA0 && base <= 0xFEFF:
		b.Cpu.Panic("Not usable region read")
		return 0x00
	case base >= 0xFF00 && base <= 0xFF7F:
		// b.Cpu.Panic("I/O Registers not implemented")
		hook := b.ioHooks[base-0xFF00]
		if b.Cpu.LogIoRegs {
			b.LogIOReg(hook, "Read", base)
		}
		return hook.OnRead(b.Cpu, base)
	case base == 0xFFFF:
		return b.Cpu.IE
	default:
		// fmt.Printf("POCJOC %#x\n", base)
		return b.Memory[base]
	}

}

func (b *Bus) Write(addr uint16, value byte) {
	switch {
	case addr < 0x8000:
		b.Cart.Write(addr, value)
		// b.Memory[addr] = value
	case addr >= 0x8000 && addr <= 0x9FFF:
		b.Cpu.PPU.VRAM[addr-0x8000] = value
	case addr >= 0xA000 && addr <= 0xBFFF:
		b.Cart.Write(addr, value)
	case addr >= 0xC000 && addr <= 0xDFFF:
		b.Memory[addr] = value
	case addr >= 0xE000 && addr <= 0xFDFF:
		b.Cpu.Panic("Write to echo ram") // echo ram
	case addr >= 0xFE00 && addr <= 0xFE9F:
		if b.Cpu.PPU.Mode != 2 && b.Cpu.PPU.Mode != 3 {
			// OAM scan or Pixel Transfer
			b.Cpu.PPU.OAM[addr-0xFE00] = value
		}
	case addr >= 0xFEA0 && addr <= 0xFEFF:
		fmt.Printf("Not usable region write\n")
	case addr >= 0xFF00 && addr <= 0xFF7F:
		hook := b.ioHooks[addr-0xFF00]
		if b.Cpu.LogIoRegs {
			b.LogIOReg(hook, "Write", addr)
		}
		hook.OnWrite(b.Cpu, addr, value)
	case addr == 0xFFFF:
		b.Cpu.IE = value
	default:
		b.Memory[addr] = value
	}
}

func (b *Bus) Read16(addr uint16) uint16 {
	lo := uint16(b.Read(addr))
	hi := uint16(b.Read(addr + 1))
	return lo | (hi << 8)
}

func (b *Bus) Write16(addr uint16, value uint16) {
	b.Write(addr, byte(value))
	b.Write(addr+1, byte(value>>8))
}

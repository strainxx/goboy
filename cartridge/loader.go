package cartridge

import "fmt"

type Cartridge interface {
	Read(addr uint16) byte
	Write(addr uint16, val byte)
	OnExit()
}

func Load(romData []byte, path string) (Cartridge, error) {
	if len(romData) < 0x0150 {
		return nil, fmt.Errorf("ROM file too small")
	}

	cartType := romData[0x0147]
	switch cartType {
	case 0x00:
		fmt.Println("Loaded ROM ONLY (MBC0)")
		return NewMBC0(romData), nil
	case 0x01, 0x02, 0x03:
		fmt.Printf("Loaded MBC1 (Type: 0x%02X)\n", cartType)
		return NewMBC1(romData, path), nil
	case 0x0F, 0x10, 0x11, 0x12, 0x13:
		fmt.Printf("Loaded MBC3 (Type: 0x%02X)\n", cartType)
		return NewMBC3(romData, path), nil
	default:
		return nil, fmt.Errorf("unsupported cartridge type: 0x%02X", cartType)
	}
}

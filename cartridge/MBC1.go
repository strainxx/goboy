package cartridge

import (
	"fmt"
	"os"
)

type MBC1 struct {
	rom         []byte
	ram         []byte
	ramEnabled  bool
	romBank     byte
	ramBank     byte
	bankingMode bool

	hasBattery bool
	savePath   string
}

func NewMBC1(romData []byte, path string) *MBC1 {
	var ram []byte
	ram_size := romData[0x149]
	switch ram_size {
	case 0x00: // no ram
	case 0x02:
		ram = make([]byte, 8192)
	case 0x03:
		ram = make([]byte, 32768)
	case 0x04:
		ram = make([]byte, 131072)
	case 0x05:
		ram = make([]byte, 65536)
	}

	hasBattery := romData[0x0147] == 0x03
	savePath := path + ".sav"
	if hasBattery {
		data, err := os.ReadFile(savePath)
		if err == nil {
			if len(data) != len(ram) {
				fmt.Printf("[MCB1] Wrong .sav file size!\n")
			} else {
				fmt.Printf("[MCB1] Loaded %s\n", savePath)
				ram = data
			}
		} else {
			fmt.Printf("[MCB1] Couldn't read .sav file! %s\n", err.Error())
		}
	}

	return &MBC1{rom: romData, ram: ram, romBank: 1, hasBattery: hasBattery, savePath: savePath}

}

func (m *MBC1) Read(addr uint16) byte {
	switch {
	case addr <= 0x3FFF:
		if m.bankingMode == false {
			// fmt.Printf("addr %#x bank %#x\n", addr)
			return m.rom[addr]
		} else {
			bank := (m.ramBank << 5)
			// fmt.Printf("addr %#x bank %#x\n", addr, bank)
			return m.rom[(int(bank)*0x4000)+int(addr)]
		}
	case addr >= 0x4000 && addr <= 0x7FFF:
		bank := (m.ramBank << 5) | m.romBank

		if bank == 0x00 || bank == 0x20 || bank == 0x40 || bank == 0x60 {
			bank += 1
		}

		// fmt.Printf("addr %#x bank %#x\n", addr, bank)
		realAddr := (int(bank) * 0x4000) + (int(addr) - 0x4000)
		return m.rom[realAddr]
	case addr >= 0xA000 && addr <= 0xBFFF:
		if !m.ramEnabled {
			return 0xFF
		}
		bank := 0
		if m.bankingMode {
			bank = int(m.ramBank)
		}
		realRamAddr := (bank * 0x2000) + int(addr-0xA000)
		return m.ram[realRamAddr]
	}
	return 0xFF
}

func (m *MBC1) Write(addr uint16, val byte) {
	switch {
	case addr <= 0x1FFF:
		enable := (val & 0x0F) == 0x0A
		m.ramEnabled = enable
	case addr >= 0x2000 && addr <= 0x3FFF:
		bank := val & 0x1F
		if bank == 0 {
			bank = 1
		}
		m.romBank = bank
	case addr >= 0x4000 && addr <= 0x5FFF:
		bank := val & 0x03
		m.ramBank = bank
	case addr >= 0x6000 && addr <= 0x7FFF:
		m.bankingMode = (val & 0x01) == 1
	case addr >= 0xA000 && addr <= 0xBFFF:
		if !m.ramEnabled {
			return
		}
		if m.bankingMode == false { // 0
			m.ram[addr-0xA000] = val
		} else { // 1
			// todo: select bank
			bank := int(m.ramBank)
			m.ram[(bank*0x2000)+(int(addr-0xA000))] = val
		}
	}

}

func (m *MBC1) OnExit() {
	if !m.hasBattery {
		return
	}

	os.WriteFile(m.savePath, m.ram, 0644)
	fmt.Printf("[MBC1] saved ram to %s\n", m.savePath)
}

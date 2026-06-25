package cartridge

import (
	"fmt"
	"os"
)

type MBC3 struct {
	rom        []byte
	ram        []byte
	ramEnabled bool
	romBank    byte
	ramRtcBank byte

	hasBattery  bool
	savePath    string
	rtcRegs     [5]byte
	latchedTime [5]byte
	latched     bool

	latchFlag bool
}

func NewMBC3(romData []byte, path string) *MBC3 {
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

	cart_type := romData[0x0147]
	hasBattery := cart_type != 0x11 && cart_type != 0x12
	savePath := path + ".sav"
	if hasBattery {
		data, err := os.ReadFile(savePath)
		if err == nil {
			if len(data) != len(ram) {
				fmt.Printf("[MCB3] Wrong .sav file size!\n")
			} else {
				fmt.Printf("[MCB3] Loaded %s\n", savePath)
				ram = data
			}
		} else {
			fmt.Printf("[MCB3] Couldn't read .sav file! %s\n", err.Error())
		}
	}

	return &MBC3{rom: romData, ram: ram, romBank: 1, hasBattery: hasBattery, savePath: savePath}

}

func (m *MBC3) GetClockRegs() [5]byte {
	if m.latched {
		return m.latchedTime
	}
	// now := time.Now()
	return [5]byte{0x00, 0x00, 0x00, 0x00}

}

func (m *MBC3) Read(addr uint16) byte {
	switch {
	case addr <= 0x3FFF:
		return m.rom[addr]
	case addr >= 0x4000 && addr <= 0x7FFF:
		bank := m.romBank
		// fmt.Printf("addr %#x bank %#x\n", addr, bank)
		realAddr := (int(bank) * 0x4000) + (int(addr) - 0x4000)
		return m.rom[realAddr]
	case addr >= 0xA000 && addr <= 0xBFFF:
		// ram or rtc
		// if !m.ramEnabled {
		// 	return 0xFF
		// }
		bank := int(m.ramRtcBank)
		if bank <= 0x07 {
			realRamAddr := (bank * 0x2000) + int(addr-0xA000)
			return m.ram[realRamAddr]
		} else if bank <= 0x0C {
			if m.latched {
				return m.latchedTime[bank-0x08]
			} else {
				m.GetClockRegs()
				return m.rtcRegs[bank-0x08]
			}
		}
	}
	return 0xFF
}

func (m *MBC3) Write(addr uint16, val byte) {
	switch {
	case addr <= 0x1FFF:
		enable := (val & 0x0F) == 0x0A
		m.ramEnabled = enable
	case addr >= 0x2000 && addr <= 0x3FFF:
		bank := val & 0x7f
		if bank == 0 {
			bank = 1
		}
		m.romBank = bank
	case addr >= 0x4000 && addr <= 0x5FFF:
		m.ramRtcBank = val
	case addr >= 0x6000 && addr <= 0x7FFF:
		if val == 0x00 {
			m.latchFlag = true
		}
		if val == 0x01 && m.latchFlag {
			m.latchFlag = false
			if m.latched {
				m.latched = false
				m.latchedTime = [5]byte{}
			} else {
				m.latched = true
				m.latchedTime = m.GetClockRegs()
			}
		}
	case addr >= 0xA000 && addr <= 0xBFFF:
		if !m.ramEnabled {
			return
		}

		bank := int(m.ramRtcBank)
		if bank <= 0x07 {
			m.ram[(bank*0x2000)+(int(addr-0xA000))] = val
		} else if bank <= 0x0C {
			m.rtcRegs[bank-0x08] = val
		}
	}

}

func (m *MBC3) OnExit() {
	if !m.hasBattery {
		return
	}

	os.WriteFile(m.savePath, m.ram, 0644)
	fmt.Printf("[MBC3] saved ram to %s\n", m.savePath)
}

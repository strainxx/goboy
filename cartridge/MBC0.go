package cartridge

type MBC0 struct {
	rom []byte
}

func NewMBC0(romData []byte) *MBC0 {
	return &MBC0{rom: romData}
}

func (m *MBC0) Read(addr uint16) byte {
	if addr < uint16(len(m.rom)) {
		return m.rom[addr]
	}
	return 0xFF
}

func (m *MBC0) Write(addr uint16, val byte) {
}

func (m *MBC0) OnExit() {
	// TODO: maybe i can save full cpu state on exit?
}

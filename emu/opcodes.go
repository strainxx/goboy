package emu

var Opcodes [256]Instruction
var CBOpcodes [256]Instruction

func init() {

	Opcodes[0x00] = Instruction{
		Name:  "NOP ",
		Bytes: 1,
		Exec:  Impl_0x00,
	}

	Opcodes[0x01] = Instruction{
		Name:  "LD BC, n16",
		Bytes: 3,
		Exec:  Impl_0x01,
	}

	Opcodes[0x02] = Instruction{
		Name:  "LD [BC], A",
		Bytes: 1,
		Exec:  Impl_0x02,
	}

	Opcodes[0x03] = Instruction{
		Name:  "INC BC",
		Bytes: 1,
		Exec:  Impl_0x03,
	}

	Opcodes[0x04] = Instruction{
		Name:  "INC B",
		Bytes: 1,
		Exec:  Impl_0x04,
	}

	Opcodes[0x05] = Instruction{
		Name:  "DEC B",
		Bytes: 1,
		Exec:  Impl_0x05,
	}

	Opcodes[0x06] = Instruction{
		Name:  "LD B, n8",
		Bytes: 2,
		Exec:  Impl_0x06,
	}

	Opcodes[0x07] = Instruction{
		Name:  "RLCA ",
		Bytes: 1,
		Exec:  Impl_0x07,
	}

	Opcodes[0x08] = Instruction{
		Name:  "LD [a16], SP",
		Bytes: 3,
		Exec:  Impl_0x08,
	}

	Opcodes[0x09] = Instruction{
		Name:  "ADD HL, BC",
		Bytes: 1,
		Exec:  Impl_0x09,
	}

	Opcodes[0x0A] = Instruction{
		Name:  "LD A, [BC]",
		Bytes: 1,
		Exec:  Impl_0x0A,
	}

	Opcodes[0x0B] = Instruction{
		Name:  "DEC BC",
		Bytes: 1,
		Exec:  Impl_0x0B,
	}

	Opcodes[0x0C] = Instruction{
		Name:  "INC C",
		Bytes: 1,
		Exec:  Impl_0x0C,
	}

	Opcodes[0x0D] = Instruction{
		Name:  "DEC C",
		Bytes: 1,
		Exec:  Impl_0x0D,
	}

	Opcodes[0x0E] = Instruction{
		Name:  "LD C, n8",
		Bytes: 2,
		Exec:  Impl_0x0E,
	}

	Opcodes[0x0F] = Instruction{
		Name:  "RRCA ",
		Bytes: 1,
		Exec:  Impl_0x0F,
	}

	Opcodes[0x10] = Instruction{
		Name:  "STOP n8",
		Bytes: 2,
		Exec:  Impl_0x10,
	}

	Opcodes[0x11] = Instruction{
		Name:  "LD DE, n16",
		Bytes: 3,
		Exec:  Impl_0x11,
	}

	Opcodes[0x12] = Instruction{
		Name:  "LD [DE], A",
		Bytes: 1,
		Exec:  Impl_0x12,
	}

	Opcodes[0x13] = Instruction{
		Name:  "INC DE",
		Bytes: 1,
		Exec:  Impl_0x13,
	}

	Opcodes[0x14] = Instruction{
		Name:  "INC D",
		Bytes: 1,
		Exec:  Impl_0x14,
	}

	Opcodes[0x15] = Instruction{
		Name:  "DEC D",
		Bytes: 1,
		Exec:  Impl_0x15,
	}

	Opcodes[0x16] = Instruction{
		Name:  "LD D, n8",
		Bytes: 2,
		Exec:  Impl_0x16,
	}

	Opcodes[0x17] = Instruction{
		Name:  "RLA ",
		Bytes: 1,
		Exec:  Impl_0x17,
	}

	Opcodes[0x18] = Instruction{
		Name:  "JR e8",
		Bytes: 2,
		Exec:  Impl_0x18,
	}

	Opcodes[0x19] = Instruction{
		Name:  "ADD HL, DE",
		Bytes: 1,
		Exec:  Impl_0x19,
	}

	Opcodes[0x1A] = Instruction{
		Name:  "LD A, [DE]",
		Bytes: 1,
		Exec:  Impl_0x1A,
	}

	Opcodes[0x1B] = Instruction{
		Name:  "DEC DE",
		Bytes: 1,
		Exec:  Impl_0x1B,
	}

	Opcodes[0x1C] = Instruction{
		Name:  "INC E",
		Bytes: 1,
		Exec:  Impl_0x1C,
	}

	Opcodes[0x1D] = Instruction{
		Name:  "DEC E",
		Bytes: 1,
		Exec:  Impl_0x1D,
	}

	Opcodes[0x1E] = Instruction{
		Name:  "LD E, n8",
		Bytes: 2,
		Exec:  Impl_0x1E,
	}

	Opcodes[0x1F] = Instruction{
		Name:  "RRA ",
		Bytes: 1,
		Exec:  Impl_0x1F,
	}

	Opcodes[0x20] = Instruction{
		Name:  "JR NZ, e8",
		Bytes: 2,
		Exec:  Impl_0x20,
	}

	Opcodes[0x21] = Instruction{
		Name:  "LD HL, n16",
		Bytes: 3,
		Exec:  Impl_0x21,
	}

	Opcodes[0x22] = Instruction{
		Name:  "LD [HL+], A",
		Bytes: 1,
		Exec:  Impl_0x22,
	}

	Opcodes[0x23] = Instruction{
		Name:  "INC HL",
		Bytes: 1,
		Exec:  Impl_0x23,
	}

	Opcodes[0x24] = Instruction{
		Name:  "INC H",
		Bytes: 1,
		Exec:  Impl_0x24,
	}

	Opcodes[0x25] = Instruction{
		Name:  "DEC H",
		Bytes: 1,
		Exec:  Impl_0x25,
	}

	Opcodes[0x26] = Instruction{
		Name:  "LD H, n8",
		Bytes: 2,
		Exec:  Impl_0x26,
	}

	Opcodes[0x27] = Instruction{
		Name:  "DAA ",
		Bytes: 1,
		Exec:  Impl_0x27,
	}

	Opcodes[0x28] = Instruction{
		Name:  "JR Z, e8",
		Bytes: 2,
		Exec:  Impl_0x28,
	}

	Opcodes[0x29] = Instruction{
		Name:  "ADD HL, HL",
		Bytes: 1,
		Exec:  Impl_0x29,
	}

	Opcodes[0x2A] = Instruction{
		Name:  "LD A, [HL+]",
		Bytes: 1,
		Exec:  Impl_0x2A,
	}

	Opcodes[0x2B] = Instruction{
		Name:  "DEC HL",
		Bytes: 1,
		Exec:  Impl_0x2B,
	}

	Opcodes[0x2C] = Instruction{
		Name:  "INC L",
		Bytes: 1,
		Exec:  Impl_0x2C,
	}

	Opcodes[0x2D] = Instruction{
		Name:  "DEC L",
		Bytes: 1,
		Exec:  Impl_0x2D,
	}

	Opcodes[0x2E] = Instruction{
		Name:  "LD L, n8",
		Bytes: 2,
		Exec:  Impl_0x2E,
	}

	Opcodes[0x2F] = Instruction{
		Name:  "CPL ",
		Bytes: 1,
		Exec:  Impl_0x2F,
	}

	Opcodes[0x30] = Instruction{
		Name:  "JR NC, e8",
		Bytes: 2,
		Exec:  Impl_0x30,
	}

	Opcodes[0x31] = Instruction{
		Name:  "LD SP, n16",
		Bytes: 3,
		Exec:  Impl_0x31,
	}

	Opcodes[0x32] = Instruction{
		Name:  "LD [HL-], A",
		Bytes: 1,
		Exec:  Impl_0x32,
	}

	Opcodes[0x33] = Instruction{
		Name:  "INC SP",
		Bytes: 1,
		Exec:  Impl_0x33,
	}

	Opcodes[0x34] = Instruction{
		Name:  "INC [HL]",
		Bytes: 1,
		Exec:  Impl_0x34,
	}

	Opcodes[0x35] = Instruction{
		Name:  "DEC [HL]",
		Bytes: 1,
		Exec:  Impl_0x35,
	}

	Opcodes[0x36] = Instruction{
		Name:  "LD [HL], n8",
		Bytes: 2,
		Exec:  Impl_0x36,
	}

	Opcodes[0x37] = Instruction{
		Name:  "SCF ",
		Bytes: 1,
		Exec:  Impl_0x37,
	}

	Opcodes[0x38] = Instruction{
		Name:  "JR C, e8",
		Bytes: 2,
		Exec:  Impl_0x38,
	}

	Opcodes[0x39] = Instruction{
		Name:  "ADD HL, SP",
		Bytes: 1,
		Exec:  Impl_0x39,
	}

	Opcodes[0x3A] = Instruction{
		Name:  "LD A, [HL-]",
		Bytes: 1,
		Exec:  Impl_0x3A,
	}

	Opcodes[0x3B] = Instruction{
		Name:  "DEC SP",
		Bytes: 1,
		Exec:  Impl_0x3B,
	}

	Opcodes[0x3C] = Instruction{
		Name:  "INC A",
		Bytes: 1,
		Exec:  Impl_0x3C,
	}

	Opcodes[0x3D] = Instruction{
		Name:  "DEC A",
		Bytes: 1,
		Exec:  Impl_0x3D,
	}

	Opcodes[0x3E] = Instruction{
		Name:  "LD A, n8",
		Bytes: 2,
		Exec:  Impl_0x3E,
	}

	Opcodes[0x3F] = Instruction{
		Name:  "CCF ",
		Bytes: 1,
		Exec:  Impl_0x3F,
	}

	Opcodes[0x40] = Instruction{
		Name:  "LD B, B",
		Bytes: 1,
		Exec:  Impl_0x40,
	}

	Opcodes[0x41] = Instruction{
		Name:  "LD B, C",
		Bytes: 1,
		Exec:  Impl_0x41,
	}

	Opcodes[0x42] = Instruction{
		Name:  "LD B, D",
		Bytes: 1,
		Exec:  Impl_0x42,
	}

	Opcodes[0x43] = Instruction{
		Name:  "LD B, E",
		Bytes: 1,
		Exec:  Impl_0x43,
	}

	Opcodes[0x44] = Instruction{
		Name:  "LD B, H",
		Bytes: 1,
		Exec:  Impl_0x44,
	}

	Opcodes[0x45] = Instruction{
		Name:  "LD B, L",
		Bytes: 1,
		Exec:  Impl_0x45,
	}

	Opcodes[0x46] = Instruction{
		Name:  "LD B, [HL]",
		Bytes: 1,
		Exec:  Impl_0x46,
	}

	Opcodes[0x47] = Instruction{
		Name:  "LD B, A",
		Bytes: 1,
		Exec:  Impl_0x47,
	}

	Opcodes[0x48] = Instruction{
		Name:  "LD C, B",
		Bytes: 1,
		Exec:  Impl_0x48,
	}

	Opcodes[0x49] = Instruction{
		Name:  "LD C, C",
		Bytes: 1,
		Exec:  Impl_0x49,
	}

	Opcodes[0x4A] = Instruction{
		Name:  "LD C, D",
		Bytes: 1,
		Exec:  Impl_0x4A,
	}

	Opcodes[0x4B] = Instruction{
		Name:  "LD C, E",
		Bytes: 1,
		Exec:  Impl_0x4B,
	}

	Opcodes[0x4C] = Instruction{
		Name:  "LD C, H",
		Bytes: 1,
		Exec:  Impl_0x4C,
	}

	Opcodes[0x4D] = Instruction{
		Name:  "LD C, L",
		Bytes: 1,
		Exec:  Impl_0x4D,
	}

	Opcodes[0x4E] = Instruction{
		Name:  "LD C, [HL]",
		Bytes: 1,
		Exec:  Impl_0x4E,
	}

	Opcodes[0x4F] = Instruction{
		Name:  "LD C, A",
		Bytes: 1,
		Exec:  Impl_0x4F,
	}

	Opcodes[0x50] = Instruction{
		Name:  "LD D, B",
		Bytes: 1,
		Exec:  Impl_0x50,
	}

	Opcodes[0x51] = Instruction{
		Name:  "LD D, C",
		Bytes: 1,
		Exec:  Impl_0x51,
	}

	Opcodes[0x52] = Instruction{
		Name:  "LD D, D",
		Bytes: 1,
		Exec:  Impl_0x52,
	}

	Opcodes[0x53] = Instruction{
		Name:  "LD D, E",
		Bytes: 1,
		Exec:  Impl_0x53,
	}

	Opcodes[0x54] = Instruction{
		Name:  "LD D, H",
		Bytes: 1,
		Exec:  Impl_0x54,
	}

	Opcodes[0x55] = Instruction{
		Name:  "LD D, L",
		Bytes: 1,
		Exec:  Impl_0x55,
	}

	Opcodes[0x56] = Instruction{
		Name:  "LD D, [HL]",
		Bytes: 1,
		Exec:  Impl_0x56,
	}

	Opcodes[0x57] = Instruction{
		Name:  "LD D, A",
		Bytes: 1,
		Exec:  Impl_0x57,
	}

	Opcodes[0x58] = Instruction{
		Name:  "LD E, B",
		Bytes: 1,
		Exec:  Impl_0x58,
	}

	Opcodes[0x59] = Instruction{
		Name:  "LD E, C",
		Bytes: 1,
		Exec:  Impl_0x59,
	}

	Opcodes[0x5A] = Instruction{
		Name:  "LD E, D",
		Bytes: 1,
		Exec:  Impl_0x5A,
	}

	Opcodes[0x5B] = Instruction{
		Name:  "LD E, E",
		Bytes: 1,
		Exec:  Impl_0x5B,
	}

	Opcodes[0x5C] = Instruction{
		Name:  "LD E, H",
		Bytes: 1,
		Exec:  Impl_0x5C,
	}

	Opcodes[0x5D] = Instruction{
		Name:  "LD E, L",
		Bytes: 1,
		Exec:  Impl_0x5D,
	}

	Opcodes[0x5E] = Instruction{
		Name:  "LD E, [HL]",
		Bytes: 1,
		Exec:  Impl_0x5E,
	}

	Opcodes[0x5F] = Instruction{
		Name:  "LD E, A",
		Bytes: 1,
		Exec:  Impl_0x5F,
	}

	Opcodes[0x60] = Instruction{
		Name:  "LD H, B",
		Bytes: 1,
		Exec:  Impl_0x60,
	}

	Opcodes[0x61] = Instruction{
		Name:  "LD H, C",
		Bytes: 1,
		Exec:  Impl_0x61,
	}

	Opcodes[0x62] = Instruction{
		Name:  "LD H, D",
		Bytes: 1,
		Exec:  Impl_0x62,
	}

	Opcodes[0x63] = Instruction{
		Name:  "LD H, E",
		Bytes: 1,
		Exec:  Impl_0x63,
	}

	Opcodes[0x64] = Instruction{
		Name:  "LD H, H",
		Bytes: 1,
		Exec:  Impl_0x64,
	}

	Opcodes[0x65] = Instruction{
		Name:  "LD H, L",
		Bytes: 1,
		Exec:  Impl_0x65,
	}

	Opcodes[0x66] = Instruction{
		Name:  "LD H, [HL]",
		Bytes: 1,
		Exec:  Impl_0x66,
	}

	Opcodes[0x67] = Instruction{
		Name:  "LD H, A",
		Bytes: 1,
		Exec:  Impl_0x67,
	}

	Opcodes[0x68] = Instruction{
		Name:  "LD L, B",
		Bytes: 1,
		Exec:  Impl_0x68,
	}

	Opcodes[0x69] = Instruction{
		Name:  "LD L, C",
		Bytes: 1,
		Exec:  Impl_0x69,
	}

	Opcodes[0x6A] = Instruction{
		Name:  "LD L, D",
		Bytes: 1,
		Exec:  Impl_0x6A,
	}

	Opcodes[0x6B] = Instruction{
		Name:  "LD L, E",
		Bytes: 1,
		Exec:  Impl_0x6B,
	}

	Opcodes[0x6C] = Instruction{
		Name:  "LD L, H",
		Bytes: 1,
		Exec:  Impl_0x6C,
	}

	Opcodes[0x6D] = Instruction{
		Name:  "LD L, L",
		Bytes: 1,
		Exec:  Impl_0x6D,
	}

	Opcodes[0x6E] = Instruction{
		Name:  "LD L, [HL]",
		Bytes: 1,
		Exec:  Impl_0x6E,
	}

	Opcodes[0x6F] = Instruction{
		Name:  "LD L, A",
		Bytes: 1,
		Exec:  Impl_0x6F,
	}

	Opcodes[0x70] = Instruction{
		Name:  "LD [HL], B",
		Bytes: 1,
		Exec:  Impl_0x70,
	}

	Opcodes[0x71] = Instruction{
		Name:  "LD [HL], C",
		Bytes: 1,
		Exec:  Impl_0x71,
	}

	Opcodes[0x72] = Instruction{
		Name:  "LD [HL], D",
		Bytes: 1,
		Exec:  Impl_0x72,
	}

	Opcodes[0x73] = Instruction{
		Name:  "LD [HL], E",
		Bytes: 1,
		Exec:  Impl_0x73,
	}

	Opcodes[0x74] = Instruction{
		Name:  "LD [HL], H",
		Bytes: 1,
		Exec:  Impl_0x74,
	}

	Opcodes[0x75] = Instruction{
		Name:  "LD [HL], L",
		Bytes: 1,
		Exec:  Impl_0x75,
	}

	Opcodes[0x76] = Instruction{
		Name:  "HALT ",
		Bytes: 1,
		Exec:  Impl_0x76,
	}

	Opcodes[0x77] = Instruction{
		Name:  "LD [HL], A",
		Bytes: 1,
		Exec:  Impl_0x77,
	}

	Opcodes[0x78] = Instruction{
		Name:  "LD A, B",
		Bytes: 1,
		Exec:  Impl_0x78,
	}

	Opcodes[0x79] = Instruction{
		Name:  "LD A, C",
		Bytes: 1,
		Exec:  Impl_0x79,
	}

	Opcodes[0x7A] = Instruction{
		Name:  "LD A, D",
		Bytes: 1,
		Exec:  Impl_0x7A,
	}

	Opcodes[0x7B] = Instruction{
		Name:  "LD A, E",
		Bytes: 1,
		Exec:  Impl_0x7B,
	}

	Opcodes[0x7C] = Instruction{
		Name:  "LD A, H",
		Bytes: 1,
		Exec:  Impl_0x7C,
	}

	Opcodes[0x7D] = Instruction{
		Name:  "LD A, L",
		Bytes: 1,
		Exec:  Impl_0x7D,
	}

	Opcodes[0x7E] = Instruction{
		Name:  "LD A, [HL]",
		Bytes: 1,
		Exec:  Impl_0x7E,
	}

	Opcodes[0x7F] = Instruction{
		Name:  "LD A, A",
		Bytes: 1,
		Exec:  Impl_0x7F,
	}

	Opcodes[0x80] = Instruction{
		Name:  "ADD A, B",
		Bytes: 1,
		Exec:  Impl_0x80,
	}

	Opcodes[0x81] = Instruction{
		Name:  "ADD A, C",
		Bytes: 1,
		Exec:  Impl_0x81,
	}

	Opcodes[0x82] = Instruction{
		Name:  "ADD A, D",
		Bytes: 1,
		Exec:  Impl_0x82,
	}

	Opcodes[0x83] = Instruction{
		Name:  "ADD A, E",
		Bytes: 1,
		Exec:  Impl_0x83,
	}

	Opcodes[0x84] = Instruction{
		Name:  "ADD A, H",
		Bytes: 1,
		Exec:  Impl_0x84,
	}

	Opcodes[0x85] = Instruction{
		Name:  "ADD A, L",
		Bytes: 1,
		Exec:  Impl_0x85,
	}

	Opcodes[0x86] = Instruction{
		Name:  "ADD A, [HL]",
		Bytes: 1,
		Exec:  Impl_0x86,
	}

	Opcodes[0x87] = Instruction{
		Name:  "ADD A, A",
		Bytes: 1,
		Exec:  Impl_0x87,
	}

	Opcodes[0x88] = Instruction{
		Name:  "ADC A, B",
		Bytes: 1,
		Exec:  Impl_0x88,
	}

	Opcodes[0x89] = Instruction{
		Name:  "ADC A, C",
		Bytes: 1,
		Exec:  Impl_0x89,
	}

	Opcodes[0x8A] = Instruction{
		Name:  "ADC A, D",
		Bytes: 1,
		Exec:  Impl_0x8A,
	}

	Opcodes[0x8B] = Instruction{
		Name:  "ADC A, E",
		Bytes: 1,
		Exec:  Impl_0x8B,
	}

	Opcodes[0x8C] = Instruction{
		Name:  "ADC A, H",
		Bytes: 1,
		Exec:  Impl_0x8C,
	}

	Opcodes[0x8D] = Instruction{
		Name:  "ADC A, L",
		Bytes: 1,
		Exec:  Impl_0x8D,
	}

	Opcodes[0x8E] = Instruction{
		Name:  "ADC A, [HL]",
		Bytes: 1,
		Exec:  Impl_0x8E,
	}

	Opcodes[0x8F] = Instruction{
		Name:  "ADC A, A",
		Bytes: 1,
		Exec:  Impl_0x8F,
	}

	Opcodes[0x90] = Instruction{
		Name:  "SUB A, B",
		Bytes: 1,
		Exec:  Impl_0x90,
	}

	Opcodes[0x91] = Instruction{
		Name:  "SUB A, C",
		Bytes: 1,
		Exec:  Impl_0x91,
	}

	Opcodes[0x92] = Instruction{
		Name:  "SUB A, D",
		Bytes: 1,
		Exec:  Impl_0x92,
	}

	Opcodes[0x93] = Instruction{
		Name:  "SUB A, E",
		Bytes: 1,
		Exec:  Impl_0x93,
	}

	Opcodes[0x94] = Instruction{
		Name:  "SUB A, H",
		Bytes: 1,
		Exec:  Impl_0x94,
	}

	Opcodes[0x95] = Instruction{
		Name:  "SUB A, L",
		Bytes: 1,
		Exec:  Impl_0x95,
	}

	Opcodes[0x96] = Instruction{
		Name:  "SUB A, [HL]",
		Bytes: 1,
		Exec:  Impl_0x96,
	}

	Opcodes[0x97] = Instruction{
		Name:  "SUB A, A",
		Bytes: 1,
		Exec:  Impl_0x97,
	}

	Opcodes[0x98] = Instruction{
		Name:  "SBC A, B",
		Bytes: 1,
		Exec:  Impl_0x98,
	}

	Opcodes[0x99] = Instruction{
		Name:  "SBC A, C",
		Bytes: 1,
		Exec:  Impl_0x99,
	}

	Opcodes[0x9A] = Instruction{
		Name:  "SBC A, D",
		Bytes: 1,
		Exec:  Impl_0x9A,
	}

	Opcodes[0x9B] = Instruction{
		Name:  "SBC A, E",
		Bytes: 1,
		Exec:  Impl_0x9B,
	}

	Opcodes[0x9C] = Instruction{
		Name:  "SBC A, H",
		Bytes: 1,
		Exec:  Impl_0x9C,
	}

	Opcodes[0x9D] = Instruction{
		Name:  "SBC A, L",
		Bytes: 1,
		Exec:  Impl_0x9D,
	}

	Opcodes[0x9E] = Instruction{
		Name:  "SBC A, [HL]",
		Bytes: 1,
		Exec:  Impl_0x9E,
	}

	Opcodes[0x9F] = Instruction{
		Name:  "SBC A, A",
		Bytes: 1,
		Exec:  Impl_0x9F,
	}

	Opcodes[0xA0] = Instruction{
		Name:  "AND A, B",
		Bytes: 1,
		Exec:  Impl_0xA0,
	}

	Opcodes[0xA1] = Instruction{
		Name:  "AND A, C",
		Bytes: 1,
		Exec:  Impl_0xA1,
	}

	Opcodes[0xA2] = Instruction{
		Name:  "AND A, D",
		Bytes: 1,
		Exec:  Impl_0xA2,
	}

	Opcodes[0xA3] = Instruction{
		Name:  "AND A, E",
		Bytes: 1,
		Exec:  Impl_0xA3,
	}

	Opcodes[0xA4] = Instruction{
		Name:  "AND A, H",
		Bytes: 1,
		Exec:  Impl_0xA4,
	}

	Opcodes[0xA5] = Instruction{
		Name:  "AND A, L",
		Bytes: 1,
		Exec:  Impl_0xA5,
	}

	Opcodes[0xA6] = Instruction{
		Name:  "AND A, [HL]",
		Bytes: 1,
		Exec:  Impl_0xA6,
	}

	Opcodes[0xA7] = Instruction{
		Name:  "AND A, A",
		Bytes: 1,
		Exec:  Impl_0xA7,
	}

	Opcodes[0xA8] = Instruction{
		Name:  "XOR A, B",
		Bytes: 1,
		Exec:  Impl_0xA8,
	}

	Opcodes[0xA9] = Instruction{
		Name:  "XOR A, C",
		Bytes: 1,
		Exec:  Impl_0xA9,
	}

	Opcodes[0xAA] = Instruction{
		Name:  "XOR A, D",
		Bytes: 1,
		Exec:  Impl_0xAA,
	}

	Opcodes[0xAB] = Instruction{
		Name:  "XOR A, E",
		Bytes: 1,
		Exec:  Impl_0xAB,
	}

	Opcodes[0xAC] = Instruction{
		Name:  "XOR A, H",
		Bytes: 1,
		Exec:  Impl_0xAC,
	}

	Opcodes[0xAD] = Instruction{
		Name:  "XOR A, L",
		Bytes: 1,
		Exec:  Impl_0xAD,
	}

	Opcodes[0xAE] = Instruction{
		Name:  "XOR A, [HL]",
		Bytes: 1,
		Exec:  Impl_0xAE,
	}

	Opcodes[0xAF] = Instruction{
		Name:  "XOR A, A",
		Bytes: 1,
		Exec:  Impl_0xAF,
	}

	Opcodes[0xB0] = Instruction{
		Name:  "OR A, B",
		Bytes: 1,
		Exec:  Impl_0xB0,
	}

	Opcodes[0xB1] = Instruction{
		Name:  "OR A, C",
		Bytes: 1,
		Exec:  Impl_0xB1,
	}

	Opcodes[0xB2] = Instruction{
		Name:  "OR A, D",
		Bytes: 1,
		Exec:  Impl_0xB2,
	}

	Opcodes[0xB3] = Instruction{
		Name:  "OR A, E",
		Bytes: 1,
		Exec:  Impl_0xB3,
	}

	Opcodes[0xB4] = Instruction{
		Name:  "OR A, H",
		Bytes: 1,
		Exec:  Impl_0xB4,
	}

	Opcodes[0xB5] = Instruction{
		Name:  "OR A, L",
		Bytes: 1,
		Exec:  Impl_0xB5,
	}

	Opcodes[0xB6] = Instruction{
		Name:  "OR A, [HL]",
		Bytes: 1,
		Exec:  Impl_0xB6,
	}

	Opcodes[0xB7] = Instruction{
		Name:  "OR A, A",
		Bytes: 1,
		Exec:  Impl_0xB7,
	}

	Opcodes[0xB8] = Instruction{
		Name:  "CP A, B",
		Bytes: 1,
		Exec:  Impl_0xB8,
	}

	Opcodes[0xB9] = Instruction{
		Name:  "CP A, C",
		Bytes: 1,
		Exec:  Impl_0xB9,
	}

	Opcodes[0xBA] = Instruction{
		Name:  "CP A, D",
		Bytes: 1,
		Exec:  Impl_0xBA,
	}

	Opcodes[0xBB] = Instruction{
		Name:  "CP A, E",
		Bytes: 1,
		Exec:  Impl_0xBB,
	}

	Opcodes[0xBC] = Instruction{
		Name:  "CP A, H",
		Bytes: 1,
		Exec:  Impl_0xBC,
	}

	Opcodes[0xBD] = Instruction{
		Name:  "CP A, L",
		Bytes: 1,
		Exec:  Impl_0xBD,
	}

	Opcodes[0xBE] = Instruction{
		Name:  "CP A, [HL]",
		Bytes: 1,
		Exec:  Impl_0xBE,
	}

	Opcodes[0xBF] = Instruction{
		Name:  "CP A, A",
		Bytes: 1,
		Exec:  Impl_0xBF,
	}

	Opcodes[0xC0] = Instruction{
		Name:  "RET NZ",
		Bytes: 1,
		Exec:  Impl_0xC0,
	}

	Opcodes[0xC1] = Instruction{
		Name:  "POP BC",
		Bytes: 1,
		Exec:  Impl_0xC1,
	}

	Opcodes[0xC2] = Instruction{
		Name:  "JP NZ, a16",
		Bytes: 3,
		Exec:  Impl_0xC2,
	}

	Opcodes[0xC3] = Instruction{
		Name:  "JP a16",
		Bytes: 3,
		Exec:  Impl_0xC3,
	}

	Opcodes[0xC4] = Instruction{
		Name:  "CALL NZ, a16",
		Bytes: 3,
		Exec:  Impl_0xC4,
	}

	Opcodes[0xC5] = Instruction{
		Name:  "PUSH BC",
		Bytes: 1,
		Exec:  Impl_0xC5,
	}

	Opcodes[0xC6] = Instruction{
		Name:  "ADD A, n8",
		Bytes: 2,
		Exec:  Impl_0xC6,
	}

	Opcodes[0xC7] = Instruction{
		Name:  "RST $00",
		Bytes: 1,
		Exec:  Impl_0xC7,
	}

	Opcodes[0xC8] = Instruction{
		Name:  "RET Z",
		Bytes: 1,
		Exec:  Impl_0xC8,
	}

	Opcodes[0xC9] = Instruction{
		Name:  "RET ",
		Bytes: 1,
		Exec:  Impl_0xC9,
	}

	Opcodes[0xCA] = Instruction{
		Name:  "JP Z, a16",
		Bytes: 3,
		Exec:  Impl_0xCA,
	}

	Opcodes[0xCB] = Instruction{
		Name:  "PREFIX ",
		Bytes: 1,
		Exec:  Impl_0xCB,
	}

	Opcodes[0xCC] = Instruction{
		Name:  "CALL Z, a16",
		Bytes: 3,
		Exec:  Impl_0xCC,
	}

	Opcodes[0xCD] = Instruction{
		Name:  "CALL a16",
		Bytes: 3,
		Exec:  Impl_0xCD,
	}

	Opcodes[0xCE] = Instruction{
		Name:  "ADC A, n8",
		Bytes: 2,
		Exec:  Impl_0xCE,
	}

	Opcodes[0xCF] = Instruction{
		Name:  "RST $08",
		Bytes: 1,
		Exec:  Impl_0xCF,
	}

	Opcodes[0xD0] = Instruction{
		Name:  "RET NC",
		Bytes: 1,
		Exec:  Impl_0xD0,
	}

	Opcodes[0xD1] = Instruction{
		Name:  "POP DE",
		Bytes: 1,
		Exec:  Impl_0xD1,
	}

	Opcodes[0xD2] = Instruction{
		Name:  "JP NC, a16",
		Bytes: 3,
		Exec:  Impl_0xD2,
	}

	Opcodes[0xD3] = Instruction{
		Name:  "ILLEGAL_D3 ",
		Bytes: 1,
		Exec:  Impl_0xD3,
	}

	Opcodes[0xD4] = Instruction{
		Name:  "CALL NC, a16",
		Bytes: 3,
		Exec:  Impl_0xD4,
	}

	Opcodes[0xD5] = Instruction{
		Name:  "PUSH DE",
		Bytes: 1,
		Exec:  Impl_0xD5,
	}

	Opcodes[0xD6] = Instruction{
		Name:  "SUB A, n8",
		Bytes: 2,
		Exec:  Impl_0xD6,
	}

	Opcodes[0xD7] = Instruction{
		Name:  "RST $10",
		Bytes: 1,
		Exec:  Impl_0xD7,
	}

	Opcodes[0xD8] = Instruction{
		Name:  "RET C",
		Bytes: 1,
		Exec:  Impl_0xD8,
	}

	Opcodes[0xD9] = Instruction{
		Name:  "RETI ",
		Bytes: 1,
		Exec:  Impl_0xD9,
	}

	Opcodes[0xDA] = Instruction{
		Name:  "JP C, a16",
		Bytes: 3,
		Exec:  Impl_0xDA,
	}

	Opcodes[0xDB] = Instruction{
		Name:  "ILLEGAL_DB ",
		Bytes: 1,
		Exec:  Impl_0xDB,
	}

	Opcodes[0xDC] = Instruction{
		Name:  "CALL C, a16",
		Bytes: 3,
		Exec:  Impl_0xDC,
	}

	Opcodes[0xDD] = Instruction{
		Name:  "ILLEGAL_DD ",
		Bytes: 1,
		Exec:  Impl_0xDD,
	}

	Opcodes[0xDE] = Instruction{
		Name:  "SBC A, n8",
		Bytes: 2,
		Exec:  Impl_0xDE,
	}

	Opcodes[0xDF] = Instruction{
		Name:  "RST $18",
		Bytes: 1,
		Exec:  Impl_0xDF,
	}

	Opcodes[0xE0] = Instruction{
		Name:  "LDH [a8], A",
		Bytes: 2,
		Exec:  Impl_0xE0,
	}

	Opcodes[0xE1] = Instruction{
		Name:  "POP HL",
		Bytes: 1,
		Exec:  Impl_0xE1,
	}

	Opcodes[0xE2] = Instruction{
		Name:  "LDH [C], A",
		Bytes: 1,
		Exec:  Impl_0xE2,
	}

	Opcodes[0xE3] = Instruction{
		Name:  "ILLEGAL_E3 ",
		Bytes: 1,
		Exec:  Impl_0xE3,
	}

	Opcodes[0xE4] = Instruction{
		Name:  "ILLEGAL_E4 ",
		Bytes: 1,
		Exec:  Impl_0xE4,
	}

	Opcodes[0xE5] = Instruction{
		Name:  "PUSH HL",
		Bytes: 1,
		Exec:  Impl_0xE5,
	}

	Opcodes[0xE6] = Instruction{
		Name:  "AND A, n8",
		Bytes: 2,
		Exec:  Impl_0xE6,
	}

	Opcodes[0xE7] = Instruction{
		Name:  "RST $20",
		Bytes: 1,
		Exec:  Impl_0xE7,
	}

	Opcodes[0xE8] = Instruction{
		Name:  "ADD SP, e8",
		Bytes: 2,
		Exec:  Impl_0xE8,
	}

	Opcodes[0xE9] = Instruction{
		Name:  "JP HL",
		Bytes: 1,
		Exec:  Impl_0xE9,
	}

	Opcodes[0xEA] = Instruction{
		Name:  "LD [a16], A",
		Bytes: 3,
		Exec:  Impl_0xEA,
	}

	Opcodes[0xEB] = Instruction{
		Name:  "ILLEGAL_EB ",
		Bytes: 1,
		Exec:  Impl_0xEB,
	}

	Opcodes[0xEC] = Instruction{
		Name:  "ILLEGAL_EC ",
		Bytes: 1,
		Exec:  Impl_0xEC,
	}

	Opcodes[0xED] = Instruction{
		Name:  "ILLEGAL_ED ",
		Bytes: 1,
		Exec:  Impl_0xED,
	}

	Opcodes[0xEE] = Instruction{
		Name:  "XOR A, n8",
		Bytes: 2,
		Exec:  Impl_0xEE,
	}

	Opcodes[0xEF] = Instruction{
		Name:  "RST $28",
		Bytes: 1,
		Exec:  Impl_0xEF,
	}

	Opcodes[0xF0] = Instruction{
		Name:  "LDH A, [a8]",
		Bytes: 2,
		Exec:  Impl_0xF0,
	}

	Opcodes[0xF1] = Instruction{
		Name:  "POP AF",
		Bytes: 1,
		Exec:  Impl_0xF1,
	}

	Opcodes[0xF2] = Instruction{
		Name:  "LDH A, [C]",
		Bytes: 1,
		Exec:  Impl_0xF2,
	}

	Opcodes[0xF3] = Instruction{
		Name:  "DI ",
		Bytes: 1,
		Exec:  Impl_0xF3,
	}

	Opcodes[0xF4] = Instruction{
		Name:  "ILLEGAL_F4 ",
		Bytes: 1,
		Exec:  Impl_0xF4,
	}

	Opcodes[0xF5] = Instruction{
		Name:  "PUSH AF",
		Bytes: 1,
		Exec:  Impl_0xF5,
	}

	Opcodes[0xF6] = Instruction{
		Name:  "OR A, n8",
		Bytes: 2,
		Exec:  Impl_0xF6,
	}

	Opcodes[0xF7] = Instruction{
		Name:  "RST $30",
		Bytes: 1,
		Exec:  Impl_0xF7,
	}

	Opcodes[0xF8] = Instruction{
		Name:  "LD HL, SP+, e8",
		Bytes: 2,
		Exec:  Impl_0xF8,
	}

	Opcodes[0xF9] = Instruction{
		Name:  "LD SP, HL",
		Bytes: 1,
		Exec:  Impl_0xF9,
	}

	Opcodes[0xFA] = Instruction{
		Name:  "LD A, [a16]",
		Bytes: 3,
		Exec:  Impl_0xFA,
	}

	Opcodes[0xFB] = Instruction{
		Name:  "EI ",
		Bytes: 1,
		Exec:  Impl_0xFB,
	}

	Opcodes[0xFC] = Instruction{
		Name:  "ILLEGAL_FC ",
		Bytes: 1,
		Exec:  Impl_0xFC,
	}

	Opcodes[0xFD] = Instruction{
		Name:  "ILLEGAL_FD ",
		Bytes: 1,
		Exec:  Impl_0xFD,
	}

	Opcodes[0xFE] = Instruction{
		Name:  "CP A, n8",
		Bytes: 2,
		Exec:  Impl_0xFE,
	}

	Opcodes[0xFF] = Instruction{
		Name:  "RST $38",
		Bytes: 1,
		Exec:  Impl_0xFF,
	}

	CBOpcodes[0x00] = Instruction{
		Name:  "RLC B",
		Bytes: 2,
		Exec:  Impl_0xCB00,
	}

	CBOpcodes[0x01] = Instruction{
		Name:  "RLC C",
		Bytes: 2,
		Exec:  Impl_0xCB01,
	}

	CBOpcodes[0x02] = Instruction{
		Name:  "RLC D",
		Bytes: 2,
		Exec:  Impl_0xCB02,
	}

	CBOpcodes[0x03] = Instruction{
		Name:  "RLC E",
		Bytes: 2,
		Exec:  Impl_0xCB03,
	}

	CBOpcodes[0x04] = Instruction{
		Name:  "RLC H",
		Bytes: 2,
		Exec:  Impl_0xCB04,
	}

	CBOpcodes[0x05] = Instruction{
		Name:  "RLC L",
		Bytes: 2,
		Exec:  Impl_0xCB05,
	}

	CBOpcodes[0x06] = Instruction{
		Name:  "RLC [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB06,
	}

	CBOpcodes[0x07] = Instruction{
		Name:  "RLC A",
		Bytes: 2,
		Exec:  Impl_0xCB07,
	}

	CBOpcodes[0x08] = Instruction{
		Name:  "RRC B",
		Bytes: 2,
		Exec:  Impl_0xCB08,
	}

	CBOpcodes[0x09] = Instruction{
		Name:  "RRC C",
		Bytes: 2,
		Exec:  Impl_0xCB09,
	}

	CBOpcodes[0x0A] = Instruction{
		Name:  "RRC D",
		Bytes: 2,
		Exec:  Impl_0xCB0A,
	}

	CBOpcodes[0x0B] = Instruction{
		Name:  "RRC E",
		Bytes: 2,
		Exec:  Impl_0xCB0B,
	}

	CBOpcodes[0x0C] = Instruction{
		Name:  "RRC H",
		Bytes: 2,
		Exec:  Impl_0xCB0C,
	}

	CBOpcodes[0x0D] = Instruction{
		Name:  "RRC L",
		Bytes: 2,
		Exec:  Impl_0xCB0D,
	}

	CBOpcodes[0x0E] = Instruction{
		Name:  "RRC [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB0E,
	}

	CBOpcodes[0x0F] = Instruction{
		Name:  "RRC A",
		Bytes: 2,
		Exec:  Impl_0xCB0F,
	}

	CBOpcodes[0x10] = Instruction{
		Name:  "RL B",
		Bytes: 2,
		Exec:  Impl_0xCB10,
	}

	CBOpcodes[0x11] = Instruction{
		Name:  "RL C",
		Bytes: 2,
		Exec:  Impl_0xCB11,
	}

	CBOpcodes[0x12] = Instruction{
		Name:  "RL D",
		Bytes: 2,
		Exec:  Impl_0xCB12,
	}

	CBOpcodes[0x13] = Instruction{
		Name:  "RL E",
		Bytes: 2,
		Exec:  Impl_0xCB13,
	}

	CBOpcodes[0x14] = Instruction{
		Name:  "RL H",
		Bytes: 2,
		Exec:  Impl_0xCB14,
	}

	CBOpcodes[0x15] = Instruction{
		Name:  "RL L",
		Bytes: 2,
		Exec:  Impl_0xCB15,
	}

	CBOpcodes[0x16] = Instruction{
		Name:  "RL [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB16,
	}

	CBOpcodes[0x17] = Instruction{
		Name:  "RL A",
		Bytes: 2,
		Exec:  Impl_0xCB17,
	}

	CBOpcodes[0x18] = Instruction{
		Name:  "RR B",
		Bytes: 2,
		Exec:  Impl_0xCB18,
	}

	CBOpcodes[0x19] = Instruction{
		Name:  "RR C",
		Bytes: 2,
		Exec:  Impl_0xCB19,
	}

	CBOpcodes[0x1A] = Instruction{
		Name:  "RR D",
		Bytes: 2,
		Exec:  Impl_0xCB1A,
	}

	CBOpcodes[0x1B] = Instruction{
		Name:  "RR E",
		Bytes: 2,
		Exec:  Impl_0xCB1B,
	}

	CBOpcodes[0x1C] = Instruction{
		Name:  "RR H",
		Bytes: 2,
		Exec:  Impl_0xCB1C,
	}

	CBOpcodes[0x1D] = Instruction{
		Name:  "RR L",
		Bytes: 2,
		Exec:  Impl_0xCB1D,
	}

	CBOpcodes[0x1E] = Instruction{
		Name:  "RR [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB1E,
	}

	CBOpcodes[0x1F] = Instruction{
		Name:  "RR A",
		Bytes: 2,
		Exec:  Impl_0xCB1F,
	}

	CBOpcodes[0x20] = Instruction{
		Name:  "SLA B",
		Bytes: 2,
		Exec:  Impl_0xCB20,
	}

	CBOpcodes[0x21] = Instruction{
		Name:  "SLA C",
		Bytes: 2,
		Exec:  Impl_0xCB21,
	}

	CBOpcodes[0x22] = Instruction{
		Name:  "SLA D",
		Bytes: 2,
		Exec:  Impl_0xCB22,
	}

	CBOpcodes[0x23] = Instruction{
		Name:  "SLA E",
		Bytes: 2,
		Exec:  Impl_0xCB23,
	}

	CBOpcodes[0x24] = Instruction{
		Name:  "SLA H",
		Bytes: 2,
		Exec:  Impl_0xCB24,
	}

	CBOpcodes[0x25] = Instruction{
		Name:  "SLA L",
		Bytes: 2,
		Exec:  Impl_0xCB25,
	}

	CBOpcodes[0x26] = Instruction{
		Name:  "SLA [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB26,
	}

	CBOpcodes[0x27] = Instruction{
		Name:  "SLA A",
		Bytes: 2,
		Exec:  Impl_0xCB27,
	}

	CBOpcodes[0x28] = Instruction{
		Name:  "SRA B",
		Bytes: 2,
		Exec:  Impl_0xCB28,
	}

	CBOpcodes[0x29] = Instruction{
		Name:  "SRA C",
		Bytes: 2,
		Exec:  Impl_0xCB29,
	}

	CBOpcodes[0x2A] = Instruction{
		Name:  "SRA D",
		Bytes: 2,
		Exec:  Impl_0xCB2A,
	}

	CBOpcodes[0x2B] = Instruction{
		Name:  "SRA E",
		Bytes: 2,
		Exec:  Impl_0xCB2B,
	}

	CBOpcodes[0x2C] = Instruction{
		Name:  "SRA H",
		Bytes: 2,
		Exec:  Impl_0xCB2C,
	}

	CBOpcodes[0x2D] = Instruction{
		Name:  "SRA L",
		Bytes: 2,
		Exec:  Impl_0xCB2D,
	}

	CBOpcodes[0x2E] = Instruction{
		Name:  "SRA [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB2E,
	}

	CBOpcodes[0x2F] = Instruction{
		Name:  "SRA A",
		Bytes: 2,
		Exec:  Impl_0xCB2F,
	}

	CBOpcodes[0x30] = Instruction{
		Name:  "SWAP B",
		Bytes: 2,
		Exec:  Impl_0xCB30,
	}

	CBOpcodes[0x31] = Instruction{
		Name:  "SWAP C",
		Bytes: 2,
		Exec:  Impl_0xCB31,
	}

	CBOpcodes[0x32] = Instruction{
		Name:  "SWAP D",
		Bytes: 2,
		Exec:  Impl_0xCB32,
	}

	CBOpcodes[0x33] = Instruction{
		Name:  "SWAP E",
		Bytes: 2,
		Exec:  Impl_0xCB33,
	}

	CBOpcodes[0x34] = Instruction{
		Name:  "SWAP H",
		Bytes: 2,
		Exec:  Impl_0xCB34,
	}

	CBOpcodes[0x35] = Instruction{
		Name:  "SWAP L",
		Bytes: 2,
		Exec:  Impl_0xCB35,
	}

	CBOpcodes[0x36] = Instruction{
		Name:  "SWAP [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB36,
	}

	CBOpcodes[0x37] = Instruction{
		Name:  "SWAP A",
		Bytes: 2,
		Exec:  Impl_0xCB37,
	}

	CBOpcodes[0x38] = Instruction{
		Name:  "SRL B",
		Bytes: 2,
		Exec:  Impl_0xCB38,
	}

	CBOpcodes[0x39] = Instruction{
		Name:  "SRL C",
		Bytes: 2,
		Exec:  Impl_0xCB39,
	}

	CBOpcodes[0x3A] = Instruction{
		Name:  "SRL D",
		Bytes: 2,
		Exec:  Impl_0xCB3A,
	}

	CBOpcodes[0x3B] = Instruction{
		Name:  "SRL E",
		Bytes: 2,
		Exec:  Impl_0xCB3B,
	}

	CBOpcodes[0x3C] = Instruction{
		Name:  "SRL H",
		Bytes: 2,
		Exec:  Impl_0xCB3C,
	}

	CBOpcodes[0x3D] = Instruction{
		Name:  "SRL L",
		Bytes: 2,
		Exec:  Impl_0xCB3D,
	}

	CBOpcodes[0x3E] = Instruction{
		Name:  "SRL [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB3E,
	}

	CBOpcodes[0x3F] = Instruction{
		Name:  "SRL A",
		Bytes: 2,
		Exec:  Impl_0xCB3F,
	}

	CBOpcodes[0x40] = Instruction{
		Name:  "BIT 0, B",
		Bytes: 2,
		Exec:  Impl_0xCB40,
	}

	CBOpcodes[0x41] = Instruction{
		Name:  "BIT 0, C",
		Bytes: 2,
		Exec:  Impl_0xCB41,
	}

	CBOpcodes[0x42] = Instruction{
		Name:  "BIT 0, D",
		Bytes: 2,
		Exec:  Impl_0xCB42,
	}

	CBOpcodes[0x43] = Instruction{
		Name:  "BIT 0, E",
		Bytes: 2,
		Exec:  Impl_0xCB43,
	}

	CBOpcodes[0x44] = Instruction{
		Name:  "BIT 0, H",
		Bytes: 2,
		Exec:  Impl_0xCB44,
	}

	CBOpcodes[0x45] = Instruction{
		Name:  "BIT 0, L",
		Bytes: 2,
		Exec:  Impl_0xCB45,
	}

	CBOpcodes[0x46] = Instruction{
		Name:  "BIT 0, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB46,
	}

	CBOpcodes[0x47] = Instruction{
		Name:  "BIT 0, A",
		Bytes: 2,
		Exec:  Impl_0xCB47,
	}

	CBOpcodes[0x48] = Instruction{
		Name:  "BIT 1, B",
		Bytes: 2,
		Exec:  Impl_0xCB48,
	}

	CBOpcodes[0x49] = Instruction{
		Name:  "BIT 1, C",
		Bytes: 2,
		Exec:  Impl_0xCB49,
	}

	CBOpcodes[0x4A] = Instruction{
		Name:  "BIT 1, D",
		Bytes: 2,
		Exec:  Impl_0xCB4A,
	}

	CBOpcodes[0x4B] = Instruction{
		Name:  "BIT 1, E",
		Bytes: 2,
		Exec:  Impl_0xCB4B,
	}

	CBOpcodes[0x4C] = Instruction{
		Name:  "BIT 1, H",
		Bytes: 2,
		Exec:  Impl_0xCB4C,
	}

	CBOpcodes[0x4D] = Instruction{
		Name:  "BIT 1, L",
		Bytes: 2,
		Exec:  Impl_0xCB4D,
	}

	CBOpcodes[0x4E] = Instruction{
		Name:  "BIT 1, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB4E,
	}

	CBOpcodes[0x4F] = Instruction{
		Name:  "BIT 1, A",
		Bytes: 2,
		Exec:  Impl_0xCB4F,
	}

	CBOpcodes[0x50] = Instruction{
		Name:  "BIT 2, B",
		Bytes: 2,
		Exec:  Impl_0xCB50,
	}

	CBOpcodes[0x51] = Instruction{
		Name:  "BIT 2, C",
		Bytes: 2,
		Exec:  Impl_0xCB51,
	}

	CBOpcodes[0x52] = Instruction{
		Name:  "BIT 2, D",
		Bytes: 2,
		Exec:  Impl_0xCB52,
	}

	CBOpcodes[0x53] = Instruction{
		Name:  "BIT 2, E",
		Bytes: 2,
		Exec:  Impl_0xCB53,
	}

	CBOpcodes[0x54] = Instruction{
		Name:  "BIT 2, H",
		Bytes: 2,
		Exec:  Impl_0xCB54,
	}

	CBOpcodes[0x55] = Instruction{
		Name:  "BIT 2, L",
		Bytes: 2,
		Exec:  Impl_0xCB55,
	}

	CBOpcodes[0x56] = Instruction{
		Name:  "BIT 2, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB56,
	}

	CBOpcodes[0x57] = Instruction{
		Name:  "BIT 2, A",
		Bytes: 2,
		Exec:  Impl_0xCB57,
	}

	CBOpcodes[0x58] = Instruction{
		Name:  "BIT 3, B",
		Bytes: 2,
		Exec:  Impl_0xCB58,
	}

	CBOpcodes[0x59] = Instruction{
		Name:  "BIT 3, C",
		Bytes: 2,
		Exec:  Impl_0xCB59,
	}

	CBOpcodes[0x5A] = Instruction{
		Name:  "BIT 3, D",
		Bytes: 2,
		Exec:  Impl_0xCB5A,
	}

	CBOpcodes[0x5B] = Instruction{
		Name:  "BIT 3, E",
		Bytes: 2,
		Exec:  Impl_0xCB5B,
	}

	CBOpcodes[0x5C] = Instruction{
		Name:  "BIT 3, H",
		Bytes: 2,
		Exec:  Impl_0xCB5C,
	}

	CBOpcodes[0x5D] = Instruction{
		Name:  "BIT 3, L",
		Bytes: 2,
		Exec:  Impl_0xCB5D,
	}

	CBOpcodes[0x5E] = Instruction{
		Name:  "BIT 3, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB5E,
	}

	CBOpcodes[0x5F] = Instruction{
		Name:  "BIT 3, A",
		Bytes: 2,
		Exec:  Impl_0xCB5F,
	}

	CBOpcodes[0x60] = Instruction{
		Name:  "BIT 4, B",
		Bytes: 2,
		Exec:  Impl_0xCB60,
	}

	CBOpcodes[0x61] = Instruction{
		Name:  "BIT 4, C",
		Bytes: 2,
		Exec:  Impl_0xCB61,
	}

	CBOpcodes[0x62] = Instruction{
		Name:  "BIT 4, D",
		Bytes: 2,
		Exec:  Impl_0xCB62,
	}

	CBOpcodes[0x63] = Instruction{
		Name:  "BIT 4, E",
		Bytes: 2,
		Exec:  Impl_0xCB63,
	}

	CBOpcodes[0x64] = Instruction{
		Name:  "BIT 4, H",
		Bytes: 2,
		Exec:  Impl_0xCB64,
	}

	CBOpcodes[0x65] = Instruction{
		Name:  "BIT 4, L",
		Bytes: 2,
		Exec:  Impl_0xCB65,
	}

	CBOpcodes[0x66] = Instruction{
		Name:  "BIT 4, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB66,
	}

	CBOpcodes[0x67] = Instruction{
		Name:  "BIT 4, A",
		Bytes: 2,
		Exec:  Impl_0xCB67,
	}

	CBOpcodes[0x68] = Instruction{
		Name:  "BIT 5, B",
		Bytes: 2,
		Exec:  Impl_0xCB68,
	}

	CBOpcodes[0x69] = Instruction{
		Name:  "BIT 5, C",
		Bytes: 2,
		Exec:  Impl_0xCB69,
	}

	CBOpcodes[0x6A] = Instruction{
		Name:  "BIT 5, D",
		Bytes: 2,
		Exec:  Impl_0xCB6A,
	}

	CBOpcodes[0x6B] = Instruction{
		Name:  "BIT 5, E",
		Bytes: 2,
		Exec:  Impl_0xCB6B,
	}

	CBOpcodes[0x6C] = Instruction{
		Name:  "BIT 5, H",
		Bytes: 2,
		Exec:  Impl_0xCB6C,
	}

	CBOpcodes[0x6D] = Instruction{
		Name:  "BIT 5, L",
		Bytes: 2,
		Exec:  Impl_0xCB6D,
	}

	CBOpcodes[0x6E] = Instruction{
		Name:  "BIT 5, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB6E,
	}

	CBOpcodes[0x6F] = Instruction{
		Name:  "BIT 5, A",
		Bytes: 2,
		Exec:  Impl_0xCB6F,
	}

	CBOpcodes[0x70] = Instruction{
		Name:  "BIT 6, B",
		Bytes: 2,
		Exec:  Impl_0xCB70,
	}

	CBOpcodes[0x71] = Instruction{
		Name:  "BIT 6, C",
		Bytes: 2,
		Exec:  Impl_0xCB71,
	}

	CBOpcodes[0x72] = Instruction{
		Name:  "BIT 6, D",
		Bytes: 2,
		Exec:  Impl_0xCB72,
	}

	CBOpcodes[0x73] = Instruction{
		Name:  "BIT 6, E",
		Bytes: 2,
		Exec:  Impl_0xCB73,
	}

	CBOpcodes[0x74] = Instruction{
		Name:  "BIT 6, H",
		Bytes: 2,
		Exec:  Impl_0xCB74,
	}

	CBOpcodes[0x75] = Instruction{
		Name:  "BIT 6, L",
		Bytes: 2,
		Exec:  Impl_0xCB75,
	}

	CBOpcodes[0x76] = Instruction{
		Name:  "BIT 6, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB76,
	}

	CBOpcodes[0x77] = Instruction{
		Name:  "BIT 6, A",
		Bytes: 2,
		Exec:  Impl_0xCB77,
	}

	CBOpcodes[0x78] = Instruction{
		Name:  "BIT 7, B",
		Bytes: 2,
		Exec:  Impl_0xCB78,
	}

	CBOpcodes[0x79] = Instruction{
		Name:  "BIT 7, C",
		Bytes: 2,
		Exec:  Impl_0xCB79,
	}

	CBOpcodes[0x7A] = Instruction{
		Name:  "BIT 7, D",
		Bytes: 2,
		Exec:  Impl_0xCB7A,
	}

	CBOpcodes[0x7B] = Instruction{
		Name:  "BIT 7, E",
		Bytes: 2,
		Exec:  Impl_0xCB7B,
	}

	CBOpcodes[0x7C] = Instruction{
		Name:  "BIT 7, H",
		Bytes: 2,
		Exec:  Impl_0xCB7C,
	}

	CBOpcodes[0x7D] = Instruction{
		Name:  "BIT 7, L",
		Bytes: 2,
		Exec:  Impl_0xCB7D,
	}

	CBOpcodes[0x7E] = Instruction{
		Name:  "BIT 7, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB7E,
	}

	CBOpcodes[0x7F] = Instruction{
		Name:  "BIT 7, A",
		Bytes: 2,
		Exec:  Impl_0xCB7F,
	}

	CBOpcodes[0x80] = Instruction{
		Name:  "RES 0, B",
		Bytes: 2,
		Exec:  Impl_0xCB80,
	}

	CBOpcodes[0x81] = Instruction{
		Name:  "RES 0, C",
		Bytes: 2,
		Exec:  Impl_0xCB81,
	}

	CBOpcodes[0x82] = Instruction{
		Name:  "RES 0, D",
		Bytes: 2,
		Exec:  Impl_0xCB82,
	}

	CBOpcodes[0x83] = Instruction{
		Name:  "RES 0, E",
		Bytes: 2,
		Exec:  Impl_0xCB83,
	}

	CBOpcodes[0x84] = Instruction{
		Name:  "RES 0, H",
		Bytes: 2,
		Exec:  Impl_0xCB84,
	}

	CBOpcodes[0x85] = Instruction{
		Name:  "RES 0, L",
		Bytes: 2,
		Exec:  Impl_0xCB85,
	}

	CBOpcodes[0x86] = Instruction{
		Name:  "RES 0, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB86,
	}

	CBOpcodes[0x87] = Instruction{
		Name:  "RES 0, A",
		Bytes: 2,
		Exec:  Impl_0xCB87,
	}

	CBOpcodes[0x88] = Instruction{
		Name:  "RES 1, B",
		Bytes: 2,
		Exec:  Impl_0xCB88,
	}

	CBOpcodes[0x89] = Instruction{
		Name:  "RES 1, C",
		Bytes: 2,
		Exec:  Impl_0xCB89,
	}

	CBOpcodes[0x8A] = Instruction{
		Name:  "RES 1, D",
		Bytes: 2,
		Exec:  Impl_0xCB8A,
	}

	CBOpcodes[0x8B] = Instruction{
		Name:  "RES 1, E",
		Bytes: 2,
		Exec:  Impl_0xCB8B,
	}

	CBOpcodes[0x8C] = Instruction{
		Name:  "RES 1, H",
		Bytes: 2,
		Exec:  Impl_0xCB8C,
	}

	CBOpcodes[0x8D] = Instruction{
		Name:  "RES 1, L",
		Bytes: 2,
		Exec:  Impl_0xCB8D,
	}

	CBOpcodes[0x8E] = Instruction{
		Name:  "RES 1, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB8E,
	}

	CBOpcodes[0x8F] = Instruction{
		Name:  "RES 1, A",
		Bytes: 2,
		Exec:  Impl_0xCB8F,
	}

	CBOpcodes[0x90] = Instruction{
		Name:  "RES 2, B",
		Bytes: 2,
		Exec:  Impl_0xCB90,
	}

	CBOpcodes[0x91] = Instruction{
		Name:  "RES 2, C",
		Bytes: 2,
		Exec:  Impl_0xCB91,
	}

	CBOpcodes[0x92] = Instruction{
		Name:  "RES 2, D",
		Bytes: 2,
		Exec:  Impl_0xCB92,
	}

	CBOpcodes[0x93] = Instruction{
		Name:  "RES 2, E",
		Bytes: 2,
		Exec:  Impl_0xCB93,
	}

	CBOpcodes[0x94] = Instruction{
		Name:  "RES 2, H",
		Bytes: 2,
		Exec:  Impl_0xCB94,
	}

	CBOpcodes[0x95] = Instruction{
		Name:  "RES 2, L",
		Bytes: 2,
		Exec:  Impl_0xCB95,
	}

	CBOpcodes[0x96] = Instruction{
		Name:  "RES 2, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB96,
	}

	CBOpcodes[0x97] = Instruction{
		Name:  "RES 2, A",
		Bytes: 2,
		Exec:  Impl_0xCB97,
	}

	CBOpcodes[0x98] = Instruction{
		Name:  "RES 3, B",
		Bytes: 2,
		Exec:  Impl_0xCB98,
	}

	CBOpcodes[0x99] = Instruction{
		Name:  "RES 3, C",
		Bytes: 2,
		Exec:  Impl_0xCB99,
	}

	CBOpcodes[0x9A] = Instruction{
		Name:  "RES 3, D",
		Bytes: 2,
		Exec:  Impl_0xCB9A,
	}

	CBOpcodes[0x9B] = Instruction{
		Name:  "RES 3, E",
		Bytes: 2,
		Exec:  Impl_0xCB9B,
	}

	CBOpcodes[0x9C] = Instruction{
		Name:  "RES 3, H",
		Bytes: 2,
		Exec:  Impl_0xCB9C,
	}

	CBOpcodes[0x9D] = Instruction{
		Name:  "RES 3, L",
		Bytes: 2,
		Exec:  Impl_0xCB9D,
	}

	CBOpcodes[0x9E] = Instruction{
		Name:  "RES 3, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCB9E,
	}

	CBOpcodes[0x9F] = Instruction{
		Name:  "RES 3, A",
		Bytes: 2,
		Exec:  Impl_0xCB9F,
	}

	CBOpcodes[0xA0] = Instruction{
		Name:  "RES 4, B",
		Bytes: 2,
		Exec:  Impl_0xCBA0,
	}

	CBOpcodes[0xA1] = Instruction{
		Name:  "RES 4, C",
		Bytes: 2,
		Exec:  Impl_0xCBA1,
	}

	CBOpcodes[0xA2] = Instruction{
		Name:  "RES 4, D",
		Bytes: 2,
		Exec:  Impl_0xCBA2,
	}

	CBOpcodes[0xA3] = Instruction{
		Name:  "RES 4, E",
		Bytes: 2,
		Exec:  Impl_0xCBA3,
	}

	CBOpcodes[0xA4] = Instruction{
		Name:  "RES 4, H",
		Bytes: 2,
		Exec:  Impl_0xCBA4,
	}

	CBOpcodes[0xA5] = Instruction{
		Name:  "RES 4, L",
		Bytes: 2,
		Exec:  Impl_0xCBA5,
	}

	CBOpcodes[0xA6] = Instruction{
		Name:  "RES 4, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBA6,
	}

	CBOpcodes[0xA7] = Instruction{
		Name:  "RES 4, A",
		Bytes: 2,
		Exec:  Impl_0xCBA7,
	}

	CBOpcodes[0xA8] = Instruction{
		Name:  "RES 5, B",
		Bytes: 2,
		Exec:  Impl_0xCBA8,
	}

	CBOpcodes[0xA9] = Instruction{
		Name:  "RES 5, C",
		Bytes: 2,
		Exec:  Impl_0xCBA9,
	}

	CBOpcodes[0xAA] = Instruction{
		Name:  "RES 5, D",
		Bytes: 2,
		Exec:  Impl_0xCBAA,
	}

	CBOpcodes[0xAB] = Instruction{
		Name:  "RES 5, E",
		Bytes: 2,
		Exec:  Impl_0xCBAB,
	}

	CBOpcodes[0xAC] = Instruction{
		Name:  "RES 5, H",
		Bytes: 2,
		Exec:  Impl_0xCBAC,
	}

	CBOpcodes[0xAD] = Instruction{
		Name:  "RES 5, L",
		Bytes: 2,
		Exec:  Impl_0xCBAD,
	}

	CBOpcodes[0xAE] = Instruction{
		Name:  "RES 5, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBAE,
	}

	CBOpcodes[0xAF] = Instruction{
		Name:  "RES 5, A",
		Bytes: 2,
		Exec:  Impl_0xCBAF,
	}

	CBOpcodes[0xB0] = Instruction{
		Name:  "RES 6, B",
		Bytes: 2,
		Exec:  Impl_0xCBB0,
	}

	CBOpcodes[0xB1] = Instruction{
		Name:  "RES 6, C",
		Bytes: 2,
		Exec:  Impl_0xCBB1,
	}

	CBOpcodes[0xB2] = Instruction{
		Name:  "RES 6, D",
		Bytes: 2,
		Exec:  Impl_0xCBB2,
	}

	CBOpcodes[0xB3] = Instruction{
		Name:  "RES 6, E",
		Bytes: 2,
		Exec:  Impl_0xCBB3,
	}

	CBOpcodes[0xB4] = Instruction{
		Name:  "RES 6, H",
		Bytes: 2,
		Exec:  Impl_0xCBB4,
	}

	CBOpcodes[0xB5] = Instruction{
		Name:  "RES 6, L",
		Bytes: 2,
		Exec:  Impl_0xCBB5,
	}

	CBOpcodes[0xB6] = Instruction{
		Name:  "RES 6, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBB6,
	}

	CBOpcodes[0xB7] = Instruction{
		Name:  "RES 6, A",
		Bytes: 2,
		Exec:  Impl_0xCBB7,
	}

	CBOpcodes[0xB8] = Instruction{
		Name:  "RES 7, B",
		Bytes: 2,
		Exec:  Impl_0xCBB8,
	}

	CBOpcodes[0xB9] = Instruction{
		Name:  "RES 7, C",
		Bytes: 2,
		Exec:  Impl_0xCBB9,
	}

	CBOpcodes[0xBA] = Instruction{
		Name:  "RES 7, D",
		Bytes: 2,
		Exec:  Impl_0xCBBA,
	}

	CBOpcodes[0xBB] = Instruction{
		Name:  "RES 7, E",
		Bytes: 2,
		Exec:  Impl_0xCBBB,
	}

	CBOpcodes[0xBC] = Instruction{
		Name:  "RES 7, H",
		Bytes: 2,
		Exec:  Impl_0xCBBC,
	}

	CBOpcodes[0xBD] = Instruction{
		Name:  "RES 7, L",
		Bytes: 2,
		Exec:  Impl_0xCBBD,
	}

	CBOpcodes[0xBE] = Instruction{
		Name:  "RES 7, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBBE,
	}

	CBOpcodes[0xBF] = Instruction{
		Name:  "RES 7, A",
		Bytes: 2,
		Exec:  Impl_0xCBBF,
	}

	CBOpcodes[0xC0] = Instruction{
		Name:  "SET 0, B",
		Bytes: 2,
		Exec:  Impl_0xCBC0,
	}

	CBOpcodes[0xC1] = Instruction{
		Name:  "SET 0, C",
		Bytes: 2,
		Exec:  Impl_0xCBC1,
	}

	CBOpcodes[0xC2] = Instruction{
		Name:  "SET 0, D",
		Bytes: 2,
		Exec:  Impl_0xCBC2,
	}

	CBOpcodes[0xC3] = Instruction{
		Name:  "SET 0, E",
		Bytes: 2,
		Exec:  Impl_0xCBC3,
	}

	CBOpcodes[0xC4] = Instruction{
		Name:  "SET 0, H",
		Bytes: 2,
		Exec:  Impl_0xCBC4,
	}

	CBOpcodes[0xC5] = Instruction{
		Name:  "SET 0, L",
		Bytes: 2,
		Exec:  Impl_0xCBC5,
	}

	CBOpcodes[0xC6] = Instruction{
		Name:  "SET 0, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBC6,
	}

	CBOpcodes[0xC7] = Instruction{
		Name:  "SET 0, A",
		Bytes: 2,
		Exec:  Impl_0xCBC7,
	}

	CBOpcodes[0xC8] = Instruction{
		Name:  "SET 1, B",
		Bytes: 2,
		Exec:  Impl_0xCBC8,
	}

	CBOpcodes[0xC9] = Instruction{
		Name:  "SET 1, C",
		Bytes: 2,
		Exec:  Impl_0xCBC9,
	}

	CBOpcodes[0xCA] = Instruction{
		Name:  "SET 1, D",
		Bytes: 2,
		Exec:  Impl_0xCBCA,
	}

	CBOpcodes[0xCB] = Instruction{
		Name:  "SET 1, E",
		Bytes: 2,
		Exec:  Impl_0xCBCB,
	}

	CBOpcodes[0xCC] = Instruction{
		Name:  "SET 1, H",
		Bytes: 2,
		Exec:  Impl_0xCBCC,
	}

	CBOpcodes[0xCD] = Instruction{
		Name:  "SET 1, L",
		Bytes: 2,
		Exec:  Impl_0xCBCD,
	}

	CBOpcodes[0xCE] = Instruction{
		Name:  "SET 1, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBCE,
	}

	CBOpcodes[0xCF] = Instruction{
		Name:  "SET 1, A",
		Bytes: 2,
		Exec:  Impl_0xCBCF,
	}

	CBOpcodes[0xD0] = Instruction{
		Name:  "SET 2, B",
		Bytes: 2,
		Exec:  Impl_0xCBD0,
	}

	CBOpcodes[0xD1] = Instruction{
		Name:  "SET 2, C",
		Bytes: 2,
		Exec:  Impl_0xCBD1,
	}

	CBOpcodes[0xD2] = Instruction{
		Name:  "SET 2, D",
		Bytes: 2,
		Exec:  Impl_0xCBD2,
	}

	CBOpcodes[0xD3] = Instruction{
		Name:  "SET 2, E",
		Bytes: 2,
		Exec:  Impl_0xCBD3,
	}

	CBOpcodes[0xD4] = Instruction{
		Name:  "SET 2, H",
		Bytes: 2,
		Exec:  Impl_0xCBD4,
	}

	CBOpcodes[0xD5] = Instruction{
		Name:  "SET 2, L",
		Bytes: 2,
		Exec:  Impl_0xCBD5,
	}

	CBOpcodes[0xD6] = Instruction{
		Name:  "SET 2, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBD6,
	}

	CBOpcodes[0xD7] = Instruction{
		Name:  "SET 2, A",
		Bytes: 2,
		Exec:  Impl_0xCBD7,
	}

	CBOpcodes[0xD8] = Instruction{
		Name:  "SET 3, B",
		Bytes: 2,
		Exec:  Impl_0xCBD8,
	}

	CBOpcodes[0xD9] = Instruction{
		Name:  "SET 3, C",
		Bytes: 2,
		Exec:  Impl_0xCBD9,
	}

	CBOpcodes[0xDA] = Instruction{
		Name:  "SET 3, D",
		Bytes: 2,
		Exec:  Impl_0xCBDA,
	}

	CBOpcodes[0xDB] = Instruction{
		Name:  "SET 3, E",
		Bytes: 2,
		Exec:  Impl_0xCBDB,
	}

	CBOpcodes[0xDC] = Instruction{
		Name:  "SET 3, H",
		Bytes: 2,
		Exec:  Impl_0xCBDC,
	}

	CBOpcodes[0xDD] = Instruction{
		Name:  "SET 3, L",
		Bytes: 2,
		Exec:  Impl_0xCBDD,
	}

	CBOpcodes[0xDE] = Instruction{
		Name:  "SET 3, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBDE,
	}

	CBOpcodes[0xDF] = Instruction{
		Name:  "SET 3, A",
		Bytes: 2,
		Exec:  Impl_0xCBDF,
	}

	CBOpcodes[0xE0] = Instruction{
		Name:  "SET 4, B",
		Bytes: 2,
		Exec:  Impl_0xCBE0,
	}

	CBOpcodes[0xE1] = Instruction{
		Name:  "SET 4, C",
		Bytes: 2,
		Exec:  Impl_0xCBE1,
	}

	CBOpcodes[0xE2] = Instruction{
		Name:  "SET 4, D",
		Bytes: 2,
		Exec:  Impl_0xCBE2,
	}

	CBOpcodes[0xE3] = Instruction{
		Name:  "SET 4, E",
		Bytes: 2,
		Exec:  Impl_0xCBE3,
	}

	CBOpcodes[0xE4] = Instruction{
		Name:  "SET 4, H",
		Bytes: 2,
		Exec:  Impl_0xCBE4,
	}

	CBOpcodes[0xE5] = Instruction{
		Name:  "SET 4, L",
		Bytes: 2,
		Exec:  Impl_0xCBE5,
	}

	CBOpcodes[0xE6] = Instruction{
		Name:  "SET 4, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBE6,
	}

	CBOpcodes[0xE7] = Instruction{
		Name:  "SET 4, A",
		Bytes: 2,
		Exec:  Impl_0xCBE7,
	}

	CBOpcodes[0xE8] = Instruction{
		Name:  "SET 5, B",
		Bytes: 2,
		Exec:  Impl_0xCBE8,
	}

	CBOpcodes[0xE9] = Instruction{
		Name:  "SET 5, C",
		Bytes: 2,
		Exec:  Impl_0xCBE9,
	}

	CBOpcodes[0xEA] = Instruction{
		Name:  "SET 5, D",
		Bytes: 2,
		Exec:  Impl_0xCBEA,
	}

	CBOpcodes[0xEB] = Instruction{
		Name:  "SET 5, E",
		Bytes: 2,
		Exec:  Impl_0xCBEB,
	}

	CBOpcodes[0xEC] = Instruction{
		Name:  "SET 5, H",
		Bytes: 2,
		Exec:  Impl_0xCBEC,
	}

	CBOpcodes[0xED] = Instruction{
		Name:  "SET 5, L",
		Bytes: 2,
		Exec:  Impl_0xCBED,
	}

	CBOpcodes[0xEE] = Instruction{
		Name:  "SET 5, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBEE,
	}

	CBOpcodes[0xEF] = Instruction{
		Name:  "SET 5, A",
		Bytes: 2,
		Exec:  Impl_0xCBEF,
	}

	CBOpcodes[0xF0] = Instruction{
		Name:  "SET 6, B",
		Bytes: 2,
		Exec:  Impl_0xCBF0,
	}

	CBOpcodes[0xF1] = Instruction{
		Name:  "SET 6, C",
		Bytes: 2,
		Exec:  Impl_0xCBF1,
	}

	CBOpcodes[0xF2] = Instruction{
		Name:  "SET 6, D",
		Bytes: 2,
		Exec:  Impl_0xCBF2,
	}

	CBOpcodes[0xF3] = Instruction{
		Name:  "SET 6, E",
		Bytes: 2,
		Exec:  Impl_0xCBF3,
	}

	CBOpcodes[0xF4] = Instruction{
		Name:  "SET 6, H",
		Bytes: 2,
		Exec:  Impl_0xCBF4,
	}

	CBOpcodes[0xF5] = Instruction{
		Name:  "SET 6, L",
		Bytes: 2,
		Exec:  Impl_0xCBF5,
	}

	CBOpcodes[0xF6] = Instruction{
		Name:  "SET 6, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBF6,
	}

	CBOpcodes[0xF7] = Instruction{
		Name:  "SET 6, A",
		Bytes: 2,
		Exec:  Impl_0xCBF7,
	}

	CBOpcodes[0xF8] = Instruction{
		Name:  "SET 7, B",
		Bytes: 2,
		Exec:  Impl_0xCBF8,
	}

	CBOpcodes[0xF9] = Instruction{
		Name:  "SET 7, C",
		Bytes: 2,
		Exec:  Impl_0xCBF9,
	}

	CBOpcodes[0xFA] = Instruction{
		Name:  "SET 7, D",
		Bytes: 2,
		Exec:  Impl_0xCBFA,
	}

	CBOpcodes[0xFB] = Instruction{
		Name:  "SET 7, E",
		Bytes: 2,
		Exec:  Impl_0xCBFB,
	}

	CBOpcodes[0xFC] = Instruction{
		Name:  "SET 7, H",
		Bytes: 2,
		Exec:  Impl_0xCBFC,
	}

	CBOpcodes[0xFD] = Instruction{
		Name:  "SET 7, L",
		Bytes: 2,
		Exec:  Impl_0xCBFD,
	}

	CBOpcodes[0xFE] = Instruction{
		Name:  "SET 7, [HL]",
		Bytes: 2,
		Exec:  Impl_0xCBFE,
	}

	CBOpcodes[0xFF] = Instruction{
		Name:  "SET 7, A",
		Bytes: 2,
		Exec:  Impl_0xCBFF,
	}

}

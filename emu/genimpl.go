package emu

func Impl_0x00(cpu *CPU) int {
	return NOP(cpu)
}

func Impl_0x01(cpu *CPU) int {
	return LD_r16_n16(cpu, BC)
}

func Impl_0x02(cpu *CPU) int {
	return LD_r16mem_r8(cpu, BC, &cpu.A)
}

func Impl_0x03(cpu *CPU) int {
	return INC_r16(cpu, BC)
}

func Impl_0x04(cpu *CPU) int {
	return INC_r8(cpu, &cpu.B)
}

func Impl_0x05(cpu *CPU) int {
	return DEC_r8(cpu, &cpu.B)
}

func Impl_0x06(cpu *CPU) int {
	return LD_r8_n8(cpu, &cpu.B)
}

func Impl_0x07(cpu *CPU) int {
	return RLCA(cpu)
}

func Impl_0x08(cpu *CPU) int {
	return LD_a16_r16(cpu, SP)
}

func Impl_0x09(cpu *CPU) int {
	return ADD_r16_r16(cpu, HL, BC)
}

func Impl_0x0A(cpu *CPU) int {
	return LD_r8_r16mem(cpu, &cpu.A, BC)
}

func Impl_0x0B(cpu *CPU) int {
	return DEC_r16(cpu, BC)
}

func Impl_0x0C(cpu *CPU) int {
	return INC_r8(cpu, &cpu.C)
}

func Impl_0x0D(cpu *CPU) int {
	return DEC_r8(cpu, &cpu.C)
}

func Impl_0x0E(cpu *CPU) int {
	return LD_r8_n8(cpu, &cpu.C)
}

func Impl_0x0F(cpu *CPU) int {
	return RRCA(cpu)
}

func Impl_0x10(cpu *CPU) int {
	return STOP_n8(cpu)
}

func Impl_0x11(cpu *CPU) int {
	return LD_r16_n16(cpu, DE)
}

func Impl_0x12(cpu *CPU) int {
	return LD_r16mem_r8(cpu, DE, &cpu.A)
}

func Impl_0x13(cpu *CPU) int {
	return INC_r16(cpu, DE)
}

func Impl_0x14(cpu *CPU) int {
	return INC_r8(cpu, &cpu.D)
}

func Impl_0x15(cpu *CPU) int {
	return DEC_r8(cpu, &cpu.D)
}

func Impl_0x16(cpu *CPU) int {
	return LD_r8_n8(cpu, &cpu.D)
}

func Impl_0x17(cpu *CPU) int {
	return RLA(cpu)
}

func Impl_0x18(cpu *CPU) int {
	return JR_e8(cpu)
}

func Impl_0x19(cpu *CPU) int {
	return ADD_r16_r16(cpu, HL, DE)
}

func Impl_0x1A(cpu *CPU) int {
	return LD_r8_r16mem(cpu, &cpu.A, DE)
}

func Impl_0x1B(cpu *CPU) int {
	return DEC_r16(cpu, DE)
}

func Impl_0x1C(cpu *CPU) int {
	return INC_r8(cpu, &cpu.E)
}

func Impl_0x1D(cpu *CPU) int {
	return DEC_r8(cpu, &cpu.E)
}

func Impl_0x1E(cpu *CPU) int {
	return LD_r8_n8(cpu, &cpu.E)
}

func Impl_0x1F(cpu *CPU) int {
	return RRA(cpu)
}

func Impl_0x20(cpu *CPU) int {
	return JR_NZ_e8(cpu)
}

func Impl_0x21(cpu *CPU) int {
	return LD_r16_n16(cpu, HL)
}

func Impl_0x22(cpu *CPU) int {
	return LD_r16inc_r8(cpu, HL, &cpu.A)
}

func Impl_0x23(cpu *CPU) int {
	return INC_r16(cpu, HL)
}

func Impl_0x24(cpu *CPU) int {
	return INC_r8(cpu, &cpu.H)
}

func Impl_0x25(cpu *CPU) int {
	return DEC_r8(cpu, &cpu.H)
}

func Impl_0x26(cpu *CPU) int {
	return LD_r8_n8(cpu, &cpu.H)
}

func Impl_0x27(cpu *CPU) int {
	return DAA(cpu)
}

func Impl_0x28(cpu *CPU) int {
	return JR_Z_e8(cpu)
}

func Impl_0x29(cpu *CPU) int {
	return ADD_r16_r16(cpu, HL, HL)
}

func Impl_0x2A(cpu *CPU) int {
	return LD_r8_r16inc(cpu, &cpu.A, HL)
}

func Impl_0x2B(cpu *CPU) int {
	return DEC_r16(cpu, HL)
}

func Impl_0x2C(cpu *CPU) int {
	return INC_r8(cpu, &cpu.L)
}

func Impl_0x2D(cpu *CPU) int {
	return DEC_r8(cpu, &cpu.L)
}

func Impl_0x2E(cpu *CPU) int {
	return LD_r8_n8(cpu, &cpu.L)
}

func Impl_0x2F(cpu *CPU) int {
	return CPL(cpu)
}

func Impl_0x30(cpu *CPU) int {
	return JR_NC_e8(cpu)
}

func Impl_0x31(cpu *CPU) int {
	return LD_r16_n16(cpu, SP)
}

func Impl_0x32(cpu *CPU) int {
	return LD_r16dec_r8(cpu, HL, &cpu.A)
}

func Impl_0x33(cpu *CPU) int {
	return INC_r16(cpu, SP)
}

func Impl_0x34(cpu *CPU) int {
	return INC_r16mem(cpu, HL)
}

func Impl_0x35(cpu *CPU) int {
	return DEC_r16mem(cpu, HL)
}

func Impl_0x36(cpu *CPU) int {
	return LD_r16mem_n8(cpu, HL)
}

func Impl_0x37(cpu *CPU) int {
	return SCF(cpu)
}

func Impl_0x38(cpu *CPU) int {
	return JR_C_e8(cpu)
}

func Impl_0x39(cpu *CPU) int {
	return ADD_r16_r16(cpu, HL, SP)
}

func Impl_0x3A(cpu *CPU) int {
	return LD_r8_r16dec(cpu, &cpu.A, HL)
}

func Impl_0x3B(cpu *CPU) int {
	return DEC_r16(cpu, SP)
}

func Impl_0x3C(cpu *CPU) int {
	return INC_r8(cpu, &cpu.A)
}

func Impl_0x3D(cpu *CPU) int {
	return DEC_r8(cpu, &cpu.A)
}

func Impl_0x3E(cpu *CPU) int {
	return LD_r8_n8(cpu, &cpu.A)
}

func Impl_0x3F(cpu *CPU) int {
	return CCF(cpu)
}

func Impl_0x40(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.B, &cpu.B)
}

func Impl_0x41(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.B, &cpu.C)
}

func Impl_0x42(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.B, &cpu.D)
}

func Impl_0x43(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.B, &cpu.E)
}

func Impl_0x44(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.B, &cpu.H)
}

func Impl_0x45(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.B, &cpu.L)
}

func Impl_0x46(cpu *CPU) int {
	return LD_r8_r16mem(cpu, &cpu.B, HL)
}

func Impl_0x47(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.B, &cpu.A)
}

func Impl_0x48(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.C, &cpu.B)
}

func Impl_0x49(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.C, &cpu.C)
}

func Impl_0x4A(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.C, &cpu.D)
}

func Impl_0x4B(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.C, &cpu.E)
}

func Impl_0x4C(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.C, &cpu.H)
}

func Impl_0x4D(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.C, &cpu.L)
}

func Impl_0x4E(cpu *CPU) int {
	return LD_r8_r16mem(cpu, &cpu.C, HL)
}

func Impl_0x4F(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.C, &cpu.A)
}

func Impl_0x50(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.D, &cpu.B)
}

func Impl_0x51(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.D, &cpu.C)
}

func Impl_0x52(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.D, &cpu.D)
}

func Impl_0x53(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.D, &cpu.E)
}

func Impl_0x54(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.D, &cpu.H)
}

func Impl_0x55(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.D, &cpu.L)
}

func Impl_0x56(cpu *CPU) int {
	return LD_r8_r16mem(cpu, &cpu.D, HL)
}

func Impl_0x57(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.D, &cpu.A)
}

func Impl_0x58(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.E, &cpu.B)
}

func Impl_0x59(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.E, &cpu.C)
}

func Impl_0x5A(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.E, &cpu.D)
}

func Impl_0x5B(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.E, &cpu.E)
}

func Impl_0x5C(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.E, &cpu.H)
}

func Impl_0x5D(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.E, &cpu.L)
}

func Impl_0x5E(cpu *CPU) int {
	return LD_r8_r16mem(cpu, &cpu.E, HL)
}

func Impl_0x5F(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.E, &cpu.A)
}

func Impl_0x60(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.H, &cpu.B)
}

func Impl_0x61(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.H, &cpu.C)
}

func Impl_0x62(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.H, &cpu.D)
}

func Impl_0x63(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.H, &cpu.E)
}

func Impl_0x64(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.H, &cpu.H)
}

func Impl_0x65(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.H, &cpu.L)
}

func Impl_0x66(cpu *CPU) int {
	return LD_r8_r16mem(cpu, &cpu.H, HL)
}

func Impl_0x67(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.H, &cpu.A)
}

func Impl_0x68(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.L, &cpu.B)
}

func Impl_0x69(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.L, &cpu.C)
}

func Impl_0x6A(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.L, &cpu.D)
}

func Impl_0x6B(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.L, &cpu.E)
}

func Impl_0x6C(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.L, &cpu.H)
}

func Impl_0x6D(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.L, &cpu.L)
}

func Impl_0x6E(cpu *CPU) int {
	return LD_r8_r16mem(cpu, &cpu.L, HL)
}

func Impl_0x6F(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.L, &cpu.A)
}

func Impl_0x70(cpu *CPU) int {
	return LD_r16mem_r8(cpu, HL, &cpu.B)
}

func Impl_0x71(cpu *CPU) int {
	return LD_r16mem_r8(cpu, HL, &cpu.C)
}

func Impl_0x72(cpu *CPU) int {
	return LD_r16mem_r8(cpu, HL, &cpu.D)
}

func Impl_0x73(cpu *CPU) int {
	return LD_r16mem_r8(cpu, HL, &cpu.E)
}

func Impl_0x74(cpu *CPU) int {
	return LD_r16mem_r8(cpu, HL, &cpu.H)
}

func Impl_0x75(cpu *CPU) int {
	return LD_r16mem_r8(cpu, HL, &cpu.L)
}

func Impl_0x76(cpu *CPU) int {
	return HALT(cpu)
}

func Impl_0x77(cpu *CPU) int {
	return LD_r16mem_r8(cpu, HL, &cpu.A)
}

func Impl_0x78(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.A, &cpu.B)
}

func Impl_0x79(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.A, &cpu.C)
}

func Impl_0x7A(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.A, &cpu.D)
}

func Impl_0x7B(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.A, &cpu.E)
}

func Impl_0x7C(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.A, &cpu.H)
}

func Impl_0x7D(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.A, &cpu.L)
}

func Impl_0x7E(cpu *CPU) int {
	return LD_r8_r16mem(cpu, &cpu.A, HL)
}

func Impl_0x7F(cpu *CPU) int {
	return LD_r8_r8(cpu, &cpu.A, &cpu.A)
}

func Impl_0x80(cpu *CPU) int {
	return ADD_r8_r8(cpu, &cpu.A, &cpu.B)
}

func Impl_0x81(cpu *CPU) int {
	return ADD_r8_r8(cpu, &cpu.A, &cpu.C)
}

func Impl_0x82(cpu *CPU) int {
	return ADD_r8_r8(cpu, &cpu.A, &cpu.D)
}

func Impl_0x83(cpu *CPU) int {
	return ADD_r8_r8(cpu, &cpu.A, &cpu.E)
}

func Impl_0x84(cpu *CPU) int {
	return ADD_r8_r8(cpu, &cpu.A, &cpu.H)
}

func Impl_0x85(cpu *CPU) int {
	return ADD_r8_r8(cpu, &cpu.A, &cpu.L)
}

func Impl_0x86(cpu *CPU) int {
	return ADD_r8_r16mem(cpu, &cpu.A, HL)
}

func Impl_0x87(cpu *CPU) int {
	return ADD_r8_r8(cpu, &cpu.A, &cpu.A)
}

func Impl_0x88(cpu *CPU) int {
	return ADC_r8_r8(cpu, &cpu.A, &cpu.B)
}

func Impl_0x89(cpu *CPU) int {
	return ADC_r8_r8(cpu, &cpu.A, &cpu.C)
}

func Impl_0x8A(cpu *CPU) int {
	return ADC_r8_r8(cpu, &cpu.A, &cpu.D)
}

func Impl_0x8B(cpu *CPU) int {
	return ADC_r8_r8(cpu, &cpu.A, &cpu.E)
}

func Impl_0x8C(cpu *CPU) int {
	return ADC_r8_r8(cpu, &cpu.A, &cpu.H)
}

func Impl_0x8D(cpu *CPU) int {
	return ADC_r8_r8(cpu, &cpu.A, &cpu.L)
}

func Impl_0x8E(cpu *CPU) int {
	return ADC_r8_r16mem(cpu, &cpu.A, HL)
}

func Impl_0x8F(cpu *CPU) int {
	return ADC_r8_r8(cpu, &cpu.A, &cpu.A)
}

func Impl_0x90(cpu *CPU) int {
	return SUB_r8_r8(cpu, &cpu.A, &cpu.B)
}

func Impl_0x91(cpu *CPU) int {
	return SUB_r8_r8(cpu, &cpu.A, &cpu.C)
}

func Impl_0x92(cpu *CPU) int {
	return SUB_r8_r8(cpu, &cpu.A, &cpu.D)
}

func Impl_0x93(cpu *CPU) int {
	return SUB_r8_r8(cpu, &cpu.A, &cpu.E)
}

func Impl_0x94(cpu *CPU) int {
	return SUB_r8_r8(cpu, &cpu.A, &cpu.H)
}

func Impl_0x95(cpu *CPU) int {
	return SUB_r8_r8(cpu, &cpu.A, &cpu.L)
}

func Impl_0x96(cpu *CPU) int {
	return SUB_r8_r16mem(cpu, &cpu.A, HL)
}

func Impl_0x97(cpu *CPU) int {
	return SUB_r8_r8(cpu, &cpu.A, &cpu.A)
}

func Impl_0x98(cpu *CPU) int {
	return SBC_r8_r8(cpu, &cpu.A, &cpu.B)
}

func Impl_0x99(cpu *CPU) int {
	return SBC_r8_r8(cpu, &cpu.A, &cpu.C)
}

func Impl_0x9A(cpu *CPU) int {
	return SBC_r8_r8(cpu, &cpu.A, &cpu.D)
}

func Impl_0x9B(cpu *CPU) int {
	return SBC_r8_r8(cpu, &cpu.A, &cpu.E)
}

func Impl_0x9C(cpu *CPU) int {
	return SBC_r8_r8(cpu, &cpu.A, &cpu.H)
}

func Impl_0x9D(cpu *CPU) int {
	return SBC_r8_r8(cpu, &cpu.A, &cpu.L)
}

func Impl_0x9E(cpu *CPU) int {
	return SBC_r8_r16mem(cpu, &cpu.A, HL)
}

func Impl_0x9F(cpu *CPU) int {
	return SBC_r8_r8(cpu, &cpu.A, &cpu.A)
}

func Impl_0xA0(cpu *CPU) int {
	return AND_r8_r8(cpu, &cpu.A, &cpu.B)
}

func Impl_0xA1(cpu *CPU) int {
	return AND_r8_r8(cpu, &cpu.A, &cpu.C)
}

func Impl_0xA2(cpu *CPU) int {
	return AND_r8_r8(cpu, &cpu.A, &cpu.D)
}

func Impl_0xA3(cpu *CPU) int {
	return AND_r8_r8(cpu, &cpu.A, &cpu.E)
}

func Impl_0xA4(cpu *CPU) int {
	return AND_r8_r8(cpu, &cpu.A, &cpu.H)
}

func Impl_0xA5(cpu *CPU) int {
	return AND_r8_r8(cpu, &cpu.A, &cpu.L)
}

func Impl_0xA6(cpu *CPU) int {
	return AND_r8_r16mem(cpu, &cpu.A, HL)
}

func Impl_0xA7(cpu *CPU) int {
	return AND_r8_r8(cpu, &cpu.A, &cpu.A)
}

func Impl_0xA8(cpu *CPU) int {
	return XOR_r8_r8(cpu, &cpu.A, &cpu.B)
}

func Impl_0xA9(cpu *CPU) int {
	return XOR_r8_r8(cpu, &cpu.A, &cpu.C)
}

func Impl_0xAA(cpu *CPU) int {
	return XOR_r8_r8(cpu, &cpu.A, &cpu.D)
}

func Impl_0xAB(cpu *CPU) int {
	return XOR_r8_r8(cpu, &cpu.A, &cpu.E)
}

func Impl_0xAC(cpu *CPU) int {
	return XOR_r8_r8(cpu, &cpu.A, &cpu.H)
}

func Impl_0xAD(cpu *CPU) int {
	return XOR_r8_r8(cpu, &cpu.A, &cpu.L)
}

func Impl_0xAE(cpu *CPU) int {
	return XOR_r8_r16mem(cpu, &cpu.A, HL)
}

func Impl_0xAF(cpu *CPU) int {
	return XOR_r8_r8(cpu, &cpu.A, &cpu.A)
}

func Impl_0xB0(cpu *CPU) int {
	return OR_r8_r8(cpu, &cpu.A, &cpu.B)
}

func Impl_0xB1(cpu *CPU) int {
	return OR_r8_r8(cpu, &cpu.A, &cpu.C)
}

func Impl_0xB2(cpu *CPU) int {
	return OR_r8_r8(cpu, &cpu.A, &cpu.D)
}

func Impl_0xB3(cpu *CPU) int {
	return OR_r8_r8(cpu, &cpu.A, &cpu.E)
}

func Impl_0xB4(cpu *CPU) int {
	return OR_r8_r8(cpu, &cpu.A, &cpu.H)
}

func Impl_0xB5(cpu *CPU) int {
	return OR_r8_r8(cpu, &cpu.A, &cpu.L)
}

func Impl_0xB6(cpu *CPU) int {
	return OR_r8_r16mem(cpu, &cpu.A, HL)
}

func Impl_0xB7(cpu *CPU) int {
	return OR_r8_r8(cpu, &cpu.A, &cpu.A)
}

func Impl_0xB8(cpu *CPU) int {
	return CP_r8_r8(cpu, &cpu.A, &cpu.B)
}

func Impl_0xB9(cpu *CPU) int {
	return CP_r8_r8(cpu, &cpu.A, &cpu.C)
}

func Impl_0xBA(cpu *CPU) int {
	return CP_r8_r8(cpu, &cpu.A, &cpu.D)
}

func Impl_0xBB(cpu *CPU) int {
	return CP_r8_r8(cpu, &cpu.A, &cpu.E)
}

func Impl_0xBC(cpu *CPU) int {
	return CP_r8_r8(cpu, &cpu.A, &cpu.H)
}

func Impl_0xBD(cpu *CPU) int {
	return CP_r8_r8(cpu, &cpu.A, &cpu.L)
}

func Impl_0xBE(cpu *CPU) int {
	return CP_r8_r16mem(cpu, &cpu.A, HL)
}

func Impl_0xBF(cpu *CPU) int {
	return CP_r8_r8(cpu, &cpu.A, &cpu.A)
}

func Impl_0xC0(cpu *CPU) int {
	return RET_NZ(cpu)
}

func Impl_0xC1(cpu *CPU) int {
	return POP_r16(cpu, BC)
}

func Impl_0xC2(cpu *CPU) int {
	return JP_NZ_a16(cpu)
}

func Impl_0xC3(cpu *CPU) int {
	return JP_a16(cpu)
}

func Impl_0xC4(cpu *CPU) int {
	return CALL_NZ_a16(cpu)
}

func Impl_0xC5(cpu *CPU) int {
	return PUSH_r16(cpu, BC)
}

func Impl_0xC6(cpu *CPU) int {
	return ADD_r8_n8(cpu, &cpu.A)
}

func Impl_0xC7(cpu *CPU) int {
	return RST_vec(cpu, 0x00)
}

func Impl_0xC8(cpu *CPU) int {
	return RET_Z(cpu)
}

func Impl_0xC9(cpu *CPU) int {
	return RET(cpu)
}

func Impl_0xCA(cpu *CPU) int {
	return JP_Z_a16(cpu)
}

func Impl_0xCB(cpu *CPU) int {
	return PREFIX(cpu)
}

func Impl_0xCC(cpu *CPU) int {
	return CALL_Z_a16(cpu)
}

func Impl_0xCD(cpu *CPU) int {
	return CALL_a16(cpu)
}

func Impl_0xCE(cpu *CPU) int {
	return ADC_r8_n8(cpu, &cpu.A)
}

func Impl_0xCF(cpu *CPU) int {
	return RST_vec(cpu, 0x08)
}

func Impl_0xD0(cpu *CPU) int {
	return RET_NC(cpu)
}

func Impl_0xD1(cpu *CPU) int {
	return POP_r16(cpu, DE)
}

func Impl_0xD2(cpu *CPU) int {
	return JP_NC_a16(cpu)
}

func Impl_0xD3(cpu *CPU) int {
	return ILLEGAL_D3(cpu)
}

func Impl_0xD4(cpu *CPU) int {
	return CALL_NC_a16(cpu)
}

func Impl_0xD5(cpu *CPU) int {
	return PUSH_r16(cpu, DE)
}

func Impl_0xD6(cpu *CPU) int {
	return SUB_r8_n8(cpu, &cpu.A)
}

func Impl_0xD7(cpu *CPU) int {
	return RST_vec(cpu, 0x10)
}

func Impl_0xD8(cpu *CPU) int {
	return RET_C(cpu)
}

func Impl_0xD9(cpu *CPU) int {
	return RETI(cpu)
}

func Impl_0xDA(cpu *CPU) int {
	return JP_C_a16(cpu)
}

func Impl_0xDB(cpu *CPU) int {
	return ILLEGAL_DB(cpu)
}

func Impl_0xDC(cpu *CPU) int {
	return CALL_C_a16(cpu)
}

func Impl_0xDD(cpu *CPU) int {
	return ILLEGAL_DD(cpu)
}

func Impl_0xDE(cpu *CPU) int {
	return SBC_r8_n8(cpu, &cpu.A)
}

func Impl_0xDF(cpu *CPU) int {
	return RST_vec(cpu, 0x18)
}

func Impl_0xE0(cpu *CPU) int {
	return LDH_a8_r8(cpu, &cpu.A)
}

func Impl_0xE1(cpu *CPU) int {
	return POP_r16(cpu, HL)
}

func Impl_0xE2(cpu *CPU) int {
	return LDH_r8mem_r8(cpu, &cpu.C, &cpu.A)
}

func Impl_0xE3(cpu *CPU) int {
	return ILLEGAL_E3(cpu)
}

func Impl_0xE4(cpu *CPU) int {
	return ILLEGAL_E4(cpu)
}

func Impl_0xE5(cpu *CPU) int {
	return PUSH_r16(cpu, HL)
}

func Impl_0xE6(cpu *CPU) int {
	return AND_r8_n8(cpu, &cpu.A)
}

func Impl_0xE7(cpu *CPU) int {
	return RST_vec(cpu, 0x20)
}

func Impl_0xE8(cpu *CPU) int {
	return ADD_r16_e8(cpu, SP)
}

func Impl_0xE9(cpu *CPU) int {
	return JP_r16(cpu, HL)
}

func Impl_0xEA(cpu *CPU) int {
	return LD_a16_r8(cpu, &cpu.A)
}

func Impl_0xEB(cpu *CPU) int {
	return ILLEGAL_EB(cpu)
}

func Impl_0xEC(cpu *CPU) int {
	return ILLEGAL_EC(cpu)
}

func Impl_0xED(cpu *CPU) int {
	return ILLEGAL_ED(cpu)
}

func Impl_0xEE(cpu *CPU) int {
	return XOR_r8_n8(cpu, &cpu.A)
}

func Impl_0xEF(cpu *CPU) int {
	return RST_vec(cpu, 0x28)
}

func Impl_0xF0(cpu *CPU) int {
	return LDH_r8_a8(cpu, &cpu.A)
}

func Impl_0xF1(cpu *CPU) int {
	return POP_r16(cpu, AF)
}

func Impl_0xF2(cpu *CPU) int {
	return LDH_r8_r8mem(cpu, &cpu.A, &cpu.C)
}

func Impl_0xF3(cpu *CPU) int {
	return DI(cpu)
}

func Impl_0xF4(cpu *CPU) int {
	return ILLEGAL_F4(cpu)
}

func Impl_0xF5(cpu *CPU) int {
	return PUSH_r16(cpu, AF)
}

func Impl_0xF6(cpu *CPU) int {
	return OR_r8_n8(cpu, &cpu.A)
}

func Impl_0xF7(cpu *CPU) int {
	return RST_vec(cpu, 0x30)
}

func Impl_0xF8(cpu *CPU) int {
	return LD_r16_r16inc_e8(cpu, HL, SP)
}

func Impl_0xF9(cpu *CPU) int {
	return LD_r16_r16(cpu, SP, HL)
}

func Impl_0xFA(cpu *CPU) int {
	return LD_r8_a16(cpu, &cpu.A)
}

func Impl_0xFB(cpu *CPU) int {
	return EI(cpu)
}

func Impl_0xFC(cpu *CPU) int {
	return ILLEGAL_FC(cpu)
}

func Impl_0xFD(cpu *CPU) int {
	return ILLEGAL_FD(cpu)
}

func Impl_0xFE(cpu *CPU) int {
	return CP_r8_n8(cpu, &cpu.A)
}

func Impl_0xFF(cpu *CPU) int {
	return RST_vec(cpu, 0x38)
}

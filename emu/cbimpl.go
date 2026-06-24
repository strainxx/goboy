package emu

func Impl_0xCB00(cpu *CPU) int {
	return RLC_r8(cpu, &cpu.B)
}

func Impl_0xCB01(cpu *CPU) int {
	return RLC_r8(cpu, &cpu.C)
}

func Impl_0xCB02(cpu *CPU) int {
	return RLC_r8(cpu, &cpu.D)
}

func Impl_0xCB03(cpu *CPU) int {
	return RLC_r8(cpu, &cpu.E)
}

func Impl_0xCB04(cpu *CPU) int {
	return RLC_r8(cpu, &cpu.H)
}

func Impl_0xCB05(cpu *CPU) int {
	return RLC_r8(cpu, &cpu.L)
}

func Impl_0xCB06(cpu *CPU) int {
	return RLC_r16mem(cpu, HL)
}

func Impl_0xCB07(cpu *CPU) int {
	return RLC_r8(cpu, &cpu.A)
}

func Impl_0xCB08(cpu *CPU) int {
	return RRC_r8(cpu, &cpu.B)
}

func Impl_0xCB09(cpu *CPU) int {
	return RRC_r8(cpu, &cpu.C)
}

func Impl_0xCB0A(cpu *CPU) int {
	return RRC_r8(cpu, &cpu.D)
}

func Impl_0xCB0B(cpu *CPU) int {
	return RRC_r8(cpu, &cpu.E)
}

func Impl_0xCB0C(cpu *CPU) int {
	return RRC_r8(cpu, &cpu.H)
}

func Impl_0xCB0D(cpu *CPU) int {
	return RRC_r8(cpu, &cpu.L)
}

func Impl_0xCB0E(cpu *CPU) int {
	return RRC_r16mem(cpu, HL)
}

func Impl_0xCB0F(cpu *CPU) int {
	return RRC_r8(cpu, &cpu.A)
}

func Impl_0xCB10(cpu *CPU) int {
	return RL_r8(cpu, &cpu.B)
}

func Impl_0xCB11(cpu *CPU) int {
	return RL_r8(cpu, &cpu.C)
}

func Impl_0xCB12(cpu *CPU) int {
	return RL_r8(cpu, &cpu.D)
}

func Impl_0xCB13(cpu *CPU) int {
	return RL_r8(cpu, &cpu.E)
}

func Impl_0xCB14(cpu *CPU) int {
	return RL_r8(cpu, &cpu.H)
}

func Impl_0xCB15(cpu *CPU) int {
	return RL_r8(cpu, &cpu.L)
}

func Impl_0xCB16(cpu *CPU) int {
	return RL_r16mem(cpu, HL)
}

func Impl_0xCB17(cpu *CPU) int {
	return RL_r8(cpu, &cpu.A)
}

func Impl_0xCB18(cpu *CPU) int {
	return RR_r8(cpu, &cpu.B)
}

func Impl_0xCB19(cpu *CPU) int {
	return RR_r8(cpu, &cpu.C)
}

func Impl_0xCB1A(cpu *CPU) int {
	return RR_r8(cpu, &cpu.D)
}

func Impl_0xCB1B(cpu *CPU) int {
	return RR_r8(cpu, &cpu.E)
}

func Impl_0xCB1C(cpu *CPU) int {
	return RR_r8(cpu, &cpu.H)
}

func Impl_0xCB1D(cpu *CPU) int {
	return RR_r8(cpu, &cpu.L)
}

func Impl_0xCB1E(cpu *CPU) int {
	return RR_r16mem(cpu, HL)
}

func Impl_0xCB1F(cpu *CPU) int {
	return RR_r8(cpu, &cpu.A)
}

func Impl_0xCB20(cpu *CPU) int {
	return SLA_r8(cpu, &cpu.B)
}

func Impl_0xCB21(cpu *CPU) int {
	return SLA_r8(cpu, &cpu.C)
}

func Impl_0xCB22(cpu *CPU) int {
	return SLA_r8(cpu, &cpu.D)
}

func Impl_0xCB23(cpu *CPU) int {
	return SLA_r8(cpu, &cpu.E)
}

func Impl_0xCB24(cpu *CPU) int {
	return SLA_r8(cpu, &cpu.H)
}

func Impl_0xCB25(cpu *CPU) int {
	return SLA_r8(cpu, &cpu.L)
}

func Impl_0xCB26(cpu *CPU) int {
	return SLA_r16mem(cpu, HL)
}

func Impl_0xCB27(cpu *CPU) int {
	return SLA_r8(cpu, &cpu.A)
}

func Impl_0xCB28(cpu *CPU) int {
	return SRA_r8(cpu, &cpu.B)
}

func Impl_0xCB29(cpu *CPU) int {
	return SRA_r8(cpu, &cpu.C)
}

func Impl_0xCB2A(cpu *CPU) int {
	return SRA_r8(cpu, &cpu.D)
}

func Impl_0xCB2B(cpu *CPU) int {
	return SRA_r8(cpu, &cpu.E)
}

func Impl_0xCB2C(cpu *CPU) int {
	return SRA_r8(cpu, &cpu.H)
}

func Impl_0xCB2D(cpu *CPU) int {
	return SRA_r8(cpu, &cpu.L)
}

func Impl_0xCB2E(cpu *CPU) int {
	return SRA_r16mem(cpu, HL)
}

func Impl_0xCB2F(cpu *CPU) int {
	return SRA_r8(cpu, &cpu.A)
}

func Impl_0xCB30(cpu *CPU) int {
	return SWAP_r8(cpu, &cpu.B)
}

func Impl_0xCB31(cpu *CPU) int {
	return SWAP_r8(cpu, &cpu.C)
}

func Impl_0xCB32(cpu *CPU) int {
	return SWAP_r8(cpu, &cpu.D)
}

func Impl_0xCB33(cpu *CPU) int {
	return SWAP_r8(cpu, &cpu.E)
}

func Impl_0xCB34(cpu *CPU) int {
	return SWAP_r8(cpu, &cpu.H)
}

func Impl_0xCB35(cpu *CPU) int {
	return SWAP_r8(cpu, &cpu.L)
}

func Impl_0xCB36(cpu *CPU) int {
	return SWAP_r16mem(cpu, HL)
}

func Impl_0xCB37(cpu *CPU) int {
	return SWAP_r8(cpu, &cpu.A)
}

func Impl_0xCB38(cpu *CPU) int {
	return SRL_r8(cpu, &cpu.B)
}

func Impl_0xCB39(cpu *CPU) int {
	return SRL_r8(cpu, &cpu.C)
}

func Impl_0xCB3A(cpu *CPU) int {
	return SRL_r8(cpu, &cpu.D)
}

func Impl_0xCB3B(cpu *CPU) int {
	return SRL_r8(cpu, &cpu.E)
}

func Impl_0xCB3C(cpu *CPU) int {
	return SRL_r8(cpu, &cpu.H)
}

func Impl_0xCB3D(cpu *CPU) int {
	return SRL_r8(cpu, &cpu.L)
}

func Impl_0xCB3E(cpu *CPU) int {
	return SRL_r16mem(cpu, HL)
}

func Impl_0xCB3F(cpu *CPU) int {
	return SRL_r8(cpu, &cpu.A)
}

func Impl_0xCB40(cpu *CPU) int {
	return BIT_u3_r8(cpu, 0, &cpu.B)
}

func Impl_0xCB41(cpu *CPU) int {
	return BIT_u3_r8(cpu, 0, &cpu.C)
}

func Impl_0xCB42(cpu *CPU) int {
	return BIT_u3_r8(cpu, 0, &cpu.D)
}

func Impl_0xCB43(cpu *CPU) int {
	return BIT_u3_r8(cpu, 0, &cpu.E)
}

func Impl_0xCB44(cpu *CPU) int {
	return BIT_u3_r8(cpu, 0, &cpu.H)
}

func Impl_0xCB45(cpu *CPU) int {
	return BIT_u3_r8(cpu, 0, &cpu.L)
}

func Impl_0xCB46(cpu *CPU) int {
	return BIT_u3_r16mem(cpu, 0, HL)
}

func Impl_0xCB47(cpu *CPU) int {
	return BIT_u3_r8(cpu, 0, &cpu.A)
}

func Impl_0xCB48(cpu *CPU) int {
	return BIT_u3_r8(cpu, 1, &cpu.B)
}

func Impl_0xCB49(cpu *CPU) int {
	return BIT_u3_r8(cpu, 1, &cpu.C)
}

func Impl_0xCB4A(cpu *CPU) int {
	return BIT_u3_r8(cpu, 1, &cpu.D)
}

func Impl_0xCB4B(cpu *CPU) int {
	return BIT_u3_r8(cpu, 1, &cpu.E)
}

func Impl_0xCB4C(cpu *CPU) int {
	return BIT_u3_r8(cpu, 1, &cpu.H)
}

func Impl_0xCB4D(cpu *CPU) int {
	return BIT_u3_r8(cpu, 1, &cpu.L)
}

func Impl_0xCB4E(cpu *CPU) int {
	return BIT_u3_r16mem(cpu, 1, HL)
}

func Impl_0xCB4F(cpu *CPU) int {
	return BIT_u3_r8(cpu, 1, &cpu.A)
}

func Impl_0xCB50(cpu *CPU) int {
	return BIT_u3_r8(cpu, 2, &cpu.B)
}

func Impl_0xCB51(cpu *CPU) int {
	return BIT_u3_r8(cpu, 2, &cpu.C)
}

func Impl_0xCB52(cpu *CPU) int {
	return BIT_u3_r8(cpu, 2, &cpu.D)
}

func Impl_0xCB53(cpu *CPU) int {
	return BIT_u3_r8(cpu, 2, &cpu.E)
}

func Impl_0xCB54(cpu *CPU) int {
	return BIT_u3_r8(cpu, 2, &cpu.H)
}

func Impl_0xCB55(cpu *CPU) int {
	return BIT_u3_r8(cpu, 2, &cpu.L)
}

func Impl_0xCB56(cpu *CPU) int {
	return BIT_u3_r16mem(cpu, 2, HL)
}

func Impl_0xCB57(cpu *CPU) int {
	return BIT_u3_r8(cpu, 2, &cpu.A)
}

func Impl_0xCB58(cpu *CPU) int {
	return BIT_u3_r8(cpu, 3, &cpu.B)
}

func Impl_0xCB59(cpu *CPU) int {
	return BIT_u3_r8(cpu, 3, &cpu.C)
}

func Impl_0xCB5A(cpu *CPU) int {
	return BIT_u3_r8(cpu, 3, &cpu.D)
}

func Impl_0xCB5B(cpu *CPU) int {
	return BIT_u3_r8(cpu, 3, &cpu.E)
}

func Impl_0xCB5C(cpu *CPU) int {
	return BIT_u3_r8(cpu, 3, &cpu.H)
}

func Impl_0xCB5D(cpu *CPU) int {
	return BIT_u3_r8(cpu, 3, &cpu.L)
}

func Impl_0xCB5E(cpu *CPU) int {
	return BIT_u3_r16mem(cpu, 3, HL)
}

func Impl_0xCB5F(cpu *CPU) int {
	return BIT_u3_r8(cpu, 3, &cpu.A)
}

func Impl_0xCB60(cpu *CPU) int {
	return BIT_u3_r8(cpu, 4, &cpu.B)
}

func Impl_0xCB61(cpu *CPU) int {
	return BIT_u3_r8(cpu, 4, &cpu.C)
}

func Impl_0xCB62(cpu *CPU) int {
	return BIT_u3_r8(cpu, 4, &cpu.D)
}

func Impl_0xCB63(cpu *CPU) int {
	return BIT_u3_r8(cpu, 4, &cpu.E)
}

func Impl_0xCB64(cpu *CPU) int {
	return BIT_u3_r8(cpu, 4, &cpu.H)
}

func Impl_0xCB65(cpu *CPU) int {
	return BIT_u3_r8(cpu, 4, &cpu.L)
}

func Impl_0xCB66(cpu *CPU) int {
	return BIT_u3_r16mem(cpu, 4, HL)
}

func Impl_0xCB67(cpu *CPU) int {
	return BIT_u3_r8(cpu, 4, &cpu.A)
}

func Impl_0xCB68(cpu *CPU) int {
	return BIT_u3_r8(cpu, 5, &cpu.B)
}

func Impl_0xCB69(cpu *CPU) int {
	return BIT_u3_r8(cpu, 5, &cpu.C)
}

func Impl_0xCB6A(cpu *CPU) int {
	return BIT_u3_r8(cpu, 5, &cpu.D)
}

func Impl_0xCB6B(cpu *CPU) int {
	return BIT_u3_r8(cpu, 5, &cpu.E)
}

func Impl_0xCB6C(cpu *CPU) int {
	return BIT_u3_r8(cpu, 5, &cpu.H)
}

func Impl_0xCB6D(cpu *CPU) int {
	return BIT_u3_r8(cpu, 5, &cpu.L)
}

func Impl_0xCB6E(cpu *CPU) int {
	return BIT_u3_r16mem(cpu, 5, HL)
}

func Impl_0xCB6F(cpu *CPU) int {
	return BIT_u3_r8(cpu, 5, &cpu.A)
}

func Impl_0xCB70(cpu *CPU) int {
	return BIT_u3_r8(cpu, 6, &cpu.B)
}

func Impl_0xCB71(cpu *CPU) int {
	return BIT_u3_r8(cpu, 6, &cpu.C)
}

func Impl_0xCB72(cpu *CPU) int {
	return BIT_u3_r8(cpu, 6, &cpu.D)
}

func Impl_0xCB73(cpu *CPU) int {
	return BIT_u3_r8(cpu, 6, &cpu.E)
}

func Impl_0xCB74(cpu *CPU) int {
	return BIT_u3_r8(cpu, 6, &cpu.H)
}

func Impl_0xCB75(cpu *CPU) int {
	return BIT_u3_r8(cpu, 6, &cpu.L)
}

func Impl_0xCB76(cpu *CPU) int {
	return BIT_u3_r16mem(cpu, 6, HL)
}

func Impl_0xCB77(cpu *CPU) int {
	return BIT_u3_r8(cpu, 6, &cpu.A)
}

func Impl_0xCB78(cpu *CPU) int {
	return BIT_u3_r8(cpu, 7, &cpu.B)
}

func Impl_0xCB79(cpu *CPU) int {
	return BIT_u3_r8(cpu, 7, &cpu.C)
}

func Impl_0xCB7A(cpu *CPU) int {
	return BIT_u3_r8(cpu, 7, &cpu.D)
}

func Impl_0xCB7B(cpu *CPU) int {
	return BIT_u3_r8(cpu, 7, &cpu.E)
}

func Impl_0xCB7C(cpu *CPU) int {
	return BIT_u3_r8(cpu, 7, &cpu.H)
}

func Impl_0xCB7D(cpu *CPU) int {
	return BIT_u3_r8(cpu, 7, &cpu.L)
}

func Impl_0xCB7E(cpu *CPU) int {
	return BIT_u3_r16mem(cpu, 7, HL)
}

func Impl_0xCB7F(cpu *CPU) int {
	return BIT_u3_r8(cpu, 7, &cpu.A)
}

func Impl_0xCB80(cpu *CPU) int {
	return RES_u3_r8(cpu, 0, &cpu.B)
}

func Impl_0xCB81(cpu *CPU) int {
	return RES_u3_r8(cpu, 0, &cpu.C)
}

func Impl_0xCB82(cpu *CPU) int {
	return RES_u3_r8(cpu, 0, &cpu.D)
}

func Impl_0xCB83(cpu *CPU) int {
	return RES_u3_r8(cpu, 0, &cpu.E)
}

func Impl_0xCB84(cpu *CPU) int {
	return RES_u3_r8(cpu, 0, &cpu.H)
}

func Impl_0xCB85(cpu *CPU) int {
	return RES_u3_r8(cpu, 0, &cpu.L)
}

func Impl_0xCB86(cpu *CPU) int {
	return RES_u3_r16mem(cpu, 0, HL)
}

func Impl_0xCB87(cpu *CPU) int {
	return RES_u3_r8(cpu, 0, &cpu.A)
}

func Impl_0xCB88(cpu *CPU) int {
	return RES_u3_r8(cpu, 1, &cpu.B)
}

func Impl_0xCB89(cpu *CPU) int {
	return RES_u3_r8(cpu, 1, &cpu.C)
}

func Impl_0xCB8A(cpu *CPU) int {
	return RES_u3_r8(cpu, 1, &cpu.D)
}

func Impl_0xCB8B(cpu *CPU) int {
	return RES_u3_r8(cpu, 1, &cpu.E)
}

func Impl_0xCB8C(cpu *CPU) int {
	return RES_u3_r8(cpu, 1, &cpu.H)
}

func Impl_0xCB8D(cpu *CPU) int {
	return RES_u3_r8(cpu, 1, &cpu.L)
}

func Impl_0xCB8E(cpu *CPU) int {
	return RES_u3_r16mem(cpu, 1, HL)
}

func Impl_0xCB8F(cpu *CPU) int {
	return RES_u3_r8(cpu, 1, &cpu.A)
}

func Impl_0xCB90(cpu *CPU) int {
	return RES_u3_r8(cpu, 2, &cpu.B)
}

func Impl_0xCB91(cpu *CPU) int {
	return RES_u3_r8(cpu, 2, &cpu.C)
}

func Impl_0xCB92(cpu *CPU) int {
	return RES_u3_r8(cpu, 2, &cpu.D)
}

func Impl_0xCB93(cpu *CPU) int {
	return RES_u3_r8(cpu, 2, &cpu.E)
}

func Impl_0xCB94(cpu *CPU) int {
	return RES_u3_r8(cpu, 2, &cpu.H)
}

func Impl_0xCB95(cpu *CPU) int {
	return RES_u3_r8(cpu, 2, &cpu.L)
}

func Impl_0xCB96(cpu *CPU) int {
	return RES_u3_r16mem(cpu, 2, HL)
}

func Impl_0xCB97(cpu *CPU) int {
	return RES_u3_r8(cpu, 2, &cpu.A)
}

func Impl_0xCB98(cpu *CPU) int {
	return RES_u3_r8(cpu, 3, &cpu.B)
}

func Impl_0xCB99(cpu *CPU) int {
	return RES_u3_r8(cpu, 3, &cpu.C)
}

func Impl_0xCB9A(cpu *CPU) int {
	return RES_u3_r8(cpu, 3, &cpu.D)
}

func Impl_0xCB9B(cpu *CPU) int {
	return RES_u3_r8(cpu, 3, &cpu.E)
}

func Impl_0xCB9C(cpu *CPU) int {
	return RES_u3_r8(cpu, 3, &cpu.H)
}

func Impl_0xCB9D(cpu *CPU) int {
	return RES_u3_r8(cpu, 3, &cpu.L)
}

func Impl_0xCB9E(cpu *CPU) int {
	return RES_u3_r16mem(cpu, 3, HL)
}

func Impl_0xCB9F(cpu *CPU) int {
	return RES_u3_r8(cpu, 3, &cpu.A)
}

func Impl_0xCBA0(cpu *CPU) int {
	return RES_u3_r8(cpu, 4, &cpu.B)
}

func Impl_0xCBA1(cpu *CPU) int {
	return RES_u3_r8(cpu, 4, &cpu.C)
}

func Impl_0xCBA2(cpu *CPU) int {
	return RES_u3_r8(cpu, 4, &cpu.D)
}

func Impl_0xCBA3(cpu *CPU) int {
	return RES_u3_r8(cpu, 4, &cpu.E)
}

func Impl_0xCBA4(cpu *CPU) int {
	return RES_u3_r8(cpu, 4, &cpu.H)
}

func Impl_0xCBA5(cpu *CPU) int {
	return RES_u3_r8(cpu, 4, &cpu.L)
}

func Impl_0xCBA6(cpu *CPU) int {
	return RES_u3_r16mem(cpu, 4, HL)
}

func Impl_0xCBA7(cpu *CPU) int {
	return RES_u3_r8(cpu, 4, &cpu.A)
}

func Impl_0xCBA8(cpu *CPU) int {
	return RES_u3_r8(cpu, 5, &cpu.B)
}

func Impl_0xCBA9(cpu *CPU) int {
	return RES_u3_r8(cpu, 5, &cpu.C)
}

func Impl_0xCBAA(cpu *CPU) int {
	return RES_u3_r8(cpu, 5, &cpu.D)
}

func Impl_0xCBAB(cpu *CPU) int {
	return RES_u3_r8(cpu, 5, &cpu.E)
}

func Impl_0xCBAC(cpu *CPU) int {
	return RES_u3_r8(cpu, 5, &cpu.H)
}

func Impl_0xCBAD(cpu *CPU) int {
	return RES_u3_r8(cpu, 5, &cpu.L)
}

func Impl_0xCBAE(cpu *CPU) int {
	return RES_u3_r16mem(cpu, 5, HL)
}

func Impl_0xCBAF(cpu *CPU) int {
	return RES_u3_r8(cpu, 5, &cpu.A)
}

func Impl_0xCBB0(cpu *CPU) int {
	return RES_u3_r8(cpu, 6, &cpu.B)
}

func Impl_0xCBB1(cpu *CPU) int {
	return RES_u3_r8(cpu, 6, &cpu.C)
}

func Impl_0xCBB2(cpu *CPU) int {
	return RES_u3_r8(cpu, 6, &cpu.D)
}

func Impl_0xCBB3(cpu *CPU) int {
	return RES_u3_r8(cpu, 6, &cpu.E)
}

func Impl_0xCBB4(cpu *CPU) int {
	return RES_u3_r8(cpu, 6, &cpu.H)
}

func Impl_0xCBB5(cpu *CPU) int {
	return RES_u3_r8(cpu, 6, &cpu.L)
}

func Impl_0xCBB6(cpu *CPU) int {
	return RES_u3_r16mem(cpu, 6, HL)
}

func Impl_0xCBB7(cpu *CPU) int {
	return RES_u3_r8(cpu, 6, &cpu.A)
}

func Impl_0xCBB8(cpu *CPU) int {
	return RES_u3_r8(cpu, 7, &cpu.B)
}

func Impl_0xCBB9(cpu *CPU) int {
	return RES_u3_r8(cpu, 7, &cpu.C)
}

func Impl_0xCBBA(cpu *CPU) int {
	return RES_u3_r8(cpu, 7, &cpu.D)
}

func Impl_0xCBBB(cpu *CPU) int {
	return RES_u3_r8(cpu, 7, &cpu.E)
}

func Impl_0xCBBC(cpu *CPU) int {
	return RES_u3_r8(cpu, 7, &cpu.H)
}

func Impl_0xCBBD(cpu *CPU) int {
	return RES_u3_r8(cpu, 7, &cpu.L)
}

func Impl_0xCBBE(cpu *CPU) int {
	return RES_u3_r16mem(cpu, 7, HL)
}

func Impl_0xCBBF(cpu *CPU) int {
	return RES_u3_r8(cpu, 7, &cpu.A)
}

func Impl_0xCBC0(cpu *CPU) int {
	return SET_u3_r8(cpu, 0, &cpu.B)
}

func Impl_0xCBC1(cpu *CPU) int {
	return SET_u3_r8(cpu, 0, &cpu.C)
}

func Impl_0xCBC2(cpu *CPU) int {
	return SET_u3_r8(cpu, 0, &cpu.D)
}

func Impl_0xCBC3(cpu *CPU) int {
	return SET_u3_r8(cpu, 0, &cpu.E)
}

func Impl_0xCBC4(cpu *CPU) int {
	return SET_u3_r8(cpu, 0, &cpu.H)
}

func Impl_0xCBC5(cpu *CPU) int {
	return SET_u3_r8(cpu, 0, &cpu.L)
}

func Impl_0xCBC6(cpu *CPU) int {
	return SET_u3_r16mem(cpu, 0, HL)
}

func Impl_0xCBC7(cpu *CPU) int {
	return SET_u3_r8(cpu, 0, &cpu.A)
}

func Impl_0xCBC8(cpu *CPU) int {
	return SET_u3_r8(cpu, 1, &cpu.B)
}

func Impl_0xCBC9(cpu *CPU) int {
	return SET_u3_r8(cpu, 1, &cpu.C)
}

func Impl_0xCBCA(cpu *CPU) int {
	return SET_u3_r8(cpu, 1, &cpu.D)
}

func Impl_0xCBCB(cpu *CPU) int {
	return SET_u3_r8(cpu, 1, &cpu.E)
}

func Impl_0xCBCC(cpu *CPU) int {
	return SET_u3_r8(cpu, 1, &cpu.H)
}

func Impl_0xCBCD(cpu *CPU) int {
	return SET_u3_r8(cpu, 1, &cpu.L)
}

func Impl_0xCBCE(cpu *CPU) int {
	return SET_u3_r16mem(cpu, 1, HL)
}

func Impl_0xCBCF(cpu *CPU) int {
	return SET_u3_r8(cpu, 1, &cpu.A)
}

func Impl_0xCBD0(cpu *CPU) int {
	return SET_u3_r8(cpu, 2, &cpu.B)
}

func Impl_0xCBD1(cpu *CPU) int {
	return SET_u3_r8(cpu, 2, &cpu.C)
}

func Impl_0xCBD2(cpu *CPU) int {
	return SET_u3_r8(cpu, 2, &cpu.D)
}

func Impl_0xCBD3(cpu *CPU) int {
	return SET_u3_r8(cpu, 2, &cpu.E)
}

func Impl_0xCBD4(cpu *CPU) int {
	return SET_u3_r8(cpu, 2, &cpu.H)
}

func Impl_0xCBD5(cpu *CPU) int {
	return SET_u3_r8(cpu, 2, &cpu.L)
}

func Impl_0xCBD6(cpu *CPU) int {
	return SET_u3_r16mem(cpu, 2, HL)
}

func Impl_0xCBD7(cpu *CPU) int {
	return SET_u3_r8(cpu, 2, &cpu.A)
}

func Impl_0xCBD8(cpu *CPU) int {
	return SET_u3_r8(cpu, 3, &cpu.B)
}

func Impl_0xCBD9(cpu *CPU) int {
	return SET_u3_r8(cpu, 3, &cpu.C)
}

func Impl_0xCBDA(cpu *CPU) int {
	return SET_u3_r8(cpu, 3, &cpu.D)
}

func Impl_0xCBDB(cpu *CPU) int {
	return SET_u3_r8(cpu, 3, &cpu.E)
}

func Impl_0xCBDC(cpu *CPU) int {
	return SET_u3_r8(cpu, 3, &cpu.H)
}

func Impl_0xCBDD(cpu *CPU) int {
	return SET_u3_r8(cpu, 3, &cpu.L)
}

func Impl_0xCBDE(cpu *CPU) int {
	return SET_u3_r16mem(cpu, 3, HL)
}

func Impl_0xCBDF(cpu *CPU) int {
	return SET_u3_r8(cpu, 3, &cpu.A)
}

func Impl_0xCBE0(cpu *CPU) int {
	return SET_u3_r8(cpu, 4, &cpu.B)
}

func Impl_0xCBE1(cpu *CPU) int {
	return SET_u3_r8(cpu, 4, &cpu.C)
}

func Impl_0xCBE2(cpu *CPU) int {
	return SET_u3_r8(cpu, 4, &cpu.D)
}

func Impl_0xCBE3(cpu *CPU) int {
	return SET_u3_r8(cpu, 4, &cpu.E)
}

func Impl_0xCBE4(cpu *CPU) int {
	return SET_u3_r8(cpu, 4, &cpu.H)
}

func Impl_0xCBE5(cpu *CPU) int {
	return SET_u3_r8(cpu, 4, &cpu.L)
}

func Impl_0xCBE6(cpu *CPU) int {
	return SET_u3_r16mem(cpu, 4, HL)
}

func Impl_0xCBE7(cpu *CPU) int {
	return SET_u3_r8(cpu, 4, &cpu.A)
}

func Impl_0xCBE8(cpu *CPU) int {
	return SET_u3_r8(cpu, 5, &cpu.B)
}

func Impl_0xCBE9(cpu *CPU) int {
	return SET_u3_r8(cpu, 5, &cpu.C)
}

func Impl_0xCBEA(cpu *CPU) int {
	return SET_u3_r8(cpu, 5, &cpu.D)
}

func Impl_0xCBEB(cpu *CPU) int {
	return SET_u3_r8(cpu, 5, &cpu.E)
}

func Impl_0xCBEC(cpu *CPU) int {
	return SET_u3_r8(cpu, 5, &cpu.H)
}

func Impl_0xCBED(cpu *CPU) int {
	return SET_u3_r8(cpu, 5, &cpu.L)
}

func Impl_0xCBEE(cpu *CPU) int {
	return SET_u3_r16mem(cpu, 5, HL)
}

func Impl_0xCBEF(cpu *CPU) int {
	return SET_u3_r8(cpu, 5, &cpu.A)
}

func Impl_0xCBF0(cpu *CPU) int {
	return SET_u3_r8(cpu, 6, &cpu.B)
}

func Impl_0xCBF1(cpu *CPU) int {
	return SET_u3_r8(cpu, 6, &cpu.C)
}

func Impl_0xCBF2(cpu *CPU) int {
	return SET_u3_r8(cpu, 6, &cpu.D)
}

func Impl_0xCBF3(cpu *CPU) int {
	return SET_u3_r8(cpu, 6, &cpu.E)
}

func Impl_0xCBF4(cpu *CPU) int {
	return SET_u3_r8(cpu, 6, &cpu.H)
}

func Impl_0xCBF5(cpu *CPU) int {
	return SET_u3_r8(cpu, 6, &cpu.L)
}

func Impl_0xCBF6(cpu *CPU) int {
	return SET_u3_r16mem(cpu, 6, HL)
}

func Impl_0xCBF7(cpu *CPU) int {
	return SET_u3_r8(cpu, 6, &cpu.A)
}

func Impl_0xCBF8(cpu *CPU) int {
	return SET_u3_r8(cpu, 7, &cpu.B)
}

func Impl_0xCBF9(cpu *CPU) int {
	return SET_u3_r8(cpu, 7, &cpu.C)
}

func Impl_0xCBFA(cpu *CPU) int {
	return SET_u3_r8(cpu, 7, &cpu.D)
}

func Impl_0xCBFB(cpu *CPU) int {
	return SET_u3_r8(cpu, 7, &cpu.E)
}

func Impl_0xCBFC(cpu *CPU) int {
	return SET_u3_r8(cpu, 7, &cpu.H)
}

func Impl_0xCBFD(cpu *CPU) int {
	return SET_u3_r8(cpu, 7, &cpu.L)
}

func Impl_0xCBFE(cpu *CPU) int {
	return SET_u3_r16mem(cpu, 7, HL)
}

func Impl_0xCBFF(cpu *CPU) int {
	return SET_u3_r8(cpu, 7, &cpu.A)
}

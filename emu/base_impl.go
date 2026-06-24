package emu

import "fmt"

type R16Enum int

const (
	AF R16Enum = iota
	BC
	DE
	HL
	SP
)

func NOP(cpu *CPU) int {
	return 1
}

func LD_r16_n16(cpu *CPU, op0 R16Enum) int {
	val := cpu.PCRead16()

	switch op0 {
	case BC:
		cpu.SetBC(val)
	case DE:
		cpu.SetDE(val)
	case HL:
		cpu.SetHL(val)
	case SP:
		cpu.SP = val
	default:
		cpu.Panic("unreachable LD_r16_n16")
	}
	return 3
}

func LD_r16mem_r8(cpu *CPU, op0 R16Enum, op1 *byte) int {
	switch op0 {
	case BC:
		cpu.Bus.Write(cpu.BC(), *op1)
	case DE:
		cpu.Bus.Write(cpu.DE(), *op1)
	case HL:
		cpu.Bus.Write(cpu.HL(), *op1)
	default:
		cpu.Panic("unreachable LD_r16mem_r8")
	}
	return 2
}

func INC_r16(cpu *CPU, op0 R16Enum) int {
	switch op0 {
	case AF:
		cpu.SetAF(cpu.AF() + 1)
	case BC:
		cpu.SetBC(cpu.BC() + 1)
	case DE:
		cpu.SetDE(cpu.DE() + 1)
	case HL:
		cpu.SetHL(cpu.HL() + 1)
	case SP:
		cpu.SP += 1
	default:
		cpu.Panic("unreachable INC_r16")

	}
	return 2
}

func INC_r8(cpu *CPU, op0 *byte) int {
	cpu.SetFlagH((*op0 & 0x0F) == 0x0F)
	*op0++
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	return 1
}

func DEC_r8(cpu *CPU, op0 *byte) int {
	cpu.SetFlagH((*op0 & 0x0F) == 0x00)
	*op0--
	cpu.SetFlagN(true)
	cpu.SetFlagZ(*op0 == 0)
	return 1
}

func LD_r8_n8(cpu *CPU, op0 *byte) int {
	*op0 = cpu.PCRead()
	return 2
}

func RLCA(cpu *CPU) int {
	bit := cpu.A >> 7
	cpu.A = cpu.A << 1
	cpu.A = cpu.A | bit

	cpu.SetFlagZ(false)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(bit == 1)
	return 1
}

func LD_a16_r16(cpu *CPU, op1 R16Enum) int {
	if op1 != SP {
		cpu.Panic("unreachable LD_a16_r16")
	}
	cpu.Bus.Write16(cpu.PCRead16(), cpu.SP)
	return 5
}

func ADD_r16_r16(cpu *CPU, op0 R16Enum, op1 R16Enum) int {
	op := cpu.HL()
	if op0 != HL {
		cpu.Panic("unreachable ADD_r16_r16")
		return 2
	}
	var val uint16
	switch op1 {
	case BC:
		val = cpu.BC()
	case DE:
		val = cpu.DE()
	case HL:
		val = op
	case SP:
		val = cpu.SP
	default:
		cpu.Panic("unreachable ADD_r16_r16")
		return 2
	}
	cpu.SetFlagH((int(op)&0x0FFF)+(int(val)&0x0FFF) > 0x0FFF)
	cpu.SetFlagC(int(op)+int(val) > 0xFFFF)

	cpu.SetHL(op + val)

	cpu.SetFlagN(false)

	return 2
}

func LD_r8_r16mem(cpu *CPU, op0 *byte, op1 R16Enum) int {
	var mem_addr uint16
	switch op1 {
	case BC:
		mem_addr = cpu.BC()
	case DE:
		mem_addr = cpu.DE()
	case HL:
		mem_addr = cpu.HL()
	default:
		cpu.Panic("unreachable LD_r8_r16mem")
	}
	*op0 = cpu.Bus.Read(mem_addr)
	return 2
}

func DEC_r16(cpu *CPU, op0 R16Enum) int {
	switch op0 {
	case BC:
		cpu.SetBC(cpu.BC() - 1)
	case DE:
		cpu.SetDE(cpu.DE() - 1)
	case HL:
		cpu.SetHL(cpu.HL() - 1)
	case SP:
		cpu.SP -= 1
	default:
		cpu.Panic("unreachable DEC_r16")
	}
	return 2
}

func RRCA(cpu *CPU) int {
	bit := cpu.A & 0x01
	cpu.A = (cpu.A >> 1) | (bit << 7)

	cpu.SetFlagZ(false)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(bit == 1)
	return 1
}

func STOP_n8(cpu *CPU) int {
	fmt.Printf("Entering stop mode!\n")
	_ = cpu.PCRead()
	interruptPending := (cpu.IF & cpu.IE) != 0
	joyp := cpu.Bus.Read(0xFF00)

	buttonHeld := false
	if (joyp & 0x30) != 0x30 {
		if (joyp & 0x0F) != 0x0F {
			buttonHeld = true
		}
	}

	// key1 := cpu.Bus.Read(0xFF4D) // GBC
	speedSwitchRequested := false // GBC
	if buttonHeld {
		if interruptPending {
			// STOP is a 1-byte opcode
			cpu.PC--
			return 1
		} else {
			// STOP is a 2-byte opcode, HALT mode is entered
			cpu.Halted = true
			return 1
		}
	}

	if !speedSwitchRequested {
		if interruptPending {
			// STOP is a 1-byte opcode, STOP mode is entered, DIV is reset
			cpu.PC--
			// cpu.StopMode = true
			cpu.Timer.DIV = 0
			return 1
		} else {
			// STOP is a 2-byte opcode, STOP mode is entered, DIV is reset
			cpu.StopMode = true
			cpu.Timer.DIV = 0
			return 1
		}
	}

	return 1
}

func RLA(cpu *CPU) int {
	bit := (cpu.A & (1 << 7)) != 0
	cpu.A = cpu.A << 1
	if cpu.GetFlagC() {
		cpu.A = cpu.A | 0x01
	}
	cpu.SetFlagC(bit)
	cpu.SetFlagZ(false)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	return 1
}

func JR_e8(cpu *CPU) int {
	off := int8(cpu.PCRead())
	cpu.PC = uint16(int(cpu.PC) + int(off))
	return 3
}

func RRA(cpu *CPU) int {
	nextCarry := (cpu.A & 0x01) != 0
	cpu.A = cpu.A >> 1

	if cpu.GetFlagC() {
		cpu.A |= 0x80
	}

	cpu.SetFlagC(nextCarry)
	cpu.SetFlagZ(false)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)

	return 1
}

func JR_NZ_e8(cpu *CPU) int {
	off := int8(cpu.PCRead())
	if cpu.GetFlagZ() == false {
		cpu.PC = uint16(int(cpu.PC) + int(off))
		return 3
	}
	return 2
}

func LD_r16inc_r8(cpu *CPU, op0 R16Enum, op1 *byte) int {
	switch op0 {
	case HL:
		cpu.Bus.Write(cpu.HL(), *op1)
		cpu.SetHL(cpu.HL() + 1)

	default:
		cpu.Panic("unreachable LD_r16inc_r8")

	}
	return 2
}

func DAA(cpu *CPU) int {
	v := uint16(cpu.A)
	if !cpu.GetFlagN() {
		if cpu.GetFlagH() || (v&0x0F) > 0x09 {
			v += 0x6
		}
		if cpu.GetFlagC() || v > 0x9F {
			v += 0x60
			cpu.SetFlagC(true)
		}
	} else {
		if cpu.GetFlagH() {
			v -= 0x06
		}
		if cpu.GetFlagC() {
			v -= 0x60
		}
	}

	cpu.A = byte(v)

	cpu.SetFlagZ(cpu.A == 0)
	cpu.SetFlagH(false)
	return 1
}

func JR_Z_e8(cpu *CPU) int {
	off := int8(cpu.PCRead())
	if cpu.GetFlagZ() {
		cpu.PC = uint16(int(cpu.PC) + int(off))
		return 3
	}
	return 2
}

func LD_r8_r16inc(cpu *CPU, op0 *byte, op1 R16Enum) int {
	switch op1 {
	case HL:
		val := cpu.Bus.Read(cpu.HL())
		*op0 = val
		cpu.SetHL(cpu.HL() + 1)

	default:
		cpu.Panic("unreachable LD_r8_r16inc")

	}
	return 2
}

func CPL(cpu *CPU) int {
	cpu.A = ^cpu.A
	cpu.SetFlagN(true)
	cpu.SetFlagH(true)
	return 1
}

func JR_NC_e8(cpu *CPU) int {
	off := int8(cpu.PCRead())
	if cpu.GetFlagC() == false {
		cpu.PC = uint16(int(cpu.PC) + int(off))
		return 3
	}
	return 2
}

func LD_r16dec_r8(cpu *CPU, op0 R16Enum, op1 *byte) int {
	switch op0 {
	case HL:
		cpu.Bus.Write(cpu.HL(), *op1)
		cpu.SetHL(cpu.HL() - 1)

	default:
		cpu.Panic("unreachable LD_r16dec_r8")

	}
	return 2
}

func INC_r16mem(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable INC_r16mem")
	}
	byte_ := cpu.Bus.Read(cpu.HL())
	cpu.SetFlagH((byte_ & 0x0F) == 0x0F)
	byte_++
	cpu.Bus.Write(cpu.HL(), byte_)
	cpu.SetFlagZ(byte_ == 0)
	cpu.SetFlagN(false)
	return 3
}

func DEC_r16mem(cpu *CPU, op0 R16Enum) int {
	switch op0 {
	case HL:

		val := cpu.Bus.Read(cpu.HL())
		cpu.SetFlagH((val & 0x0F) == 0x00)

		val--
		cpu.Bus.Write(cpu.HL(), val)
		cpu.SetFlagZ(val == 0)
		cpu.SetFlagN(true)

	default:
		cpu.Panic("unreachable DEC_r16mem")

	}
	return 3
}

func LD_r16mem_n8(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable LD_r16mem_n8")
	}
	val := cpu.PCRead()

	cpu.Bus.Write(cpu.HL(), val)
	return 3
}

func SCF(cpu *CPU) int {
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(true)
	return 1
}

func JR_C_e8(cpu *CPU) int {
	off := int8(cpu.PCRead())
	if cpu.GetFlagC() {
		cpu.PC = uint16(int(cpu.PC) + int(off))
		return 3
	}
	return 2
}

func LD_r8_r16dec(cpu *CPU, op0 *byte, op1 R16Enum) int {
	switch op1 {
	case HL:
		val := cpu.Bus.Read(cpu.HL())
		*op0 = val
		cpu.SetHL(cpu.HL() - 1)

	default:
		cpu.Panic("unreachable LD_r8_r16dec")

	}
	return 2
}

func CCF(cpu *CPU) int {
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(!cpu.GetFlagC())
	return 1
}

func LD_r8_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	*op0 = *op1
	return 1
}

func HALT(cpu *CPU) int {
	// fmt.Printf("halted! IE = %#x\n", cpu.IE)
	cpu.Halted = true
	return 1
}

func ADD_r8_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	sum := int(*op0) + int(*op1)
	cpu.SetFlagH((int(*op0)&0x0F)+(int(*op1)&0x0F) > 0x0F)
	cpu.SetFlagC(sum > 0xFF)
	*op0 = byte(sum)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	return 1
}

func ADD_r8_r16mem(cpu *CPU, op0 *byte, op1 R16Enum) int {
	switch op1 {
	case HL:
		val := cpu.Bus.Read(cpu.HL())
		sum := int(*op0) + int(val)
		cpu.SetFlagH((int(*op0)&0x0F)+(int(val)&0x0F) > 0x0F)
		cpu.SetFlagC(sum > 0xFF)
		*op0 = byte(sum)
		cpu.SetFlagZ(*op0 == 0)
		cpu.SetFlagN(false)

	default:
		cpu.Panic("unreachable ADD_r8_r16mem")

	}
	return 2
}

func ADC_r8_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	val := *op1
	carry := 0
	if cpu.GetFlagC() {
		carry = 1
	}
	sum := int(*op0) + int(val) + carry
	halfCarry := (int(*op0) & 0x0F) + (int(val) & 0x0F) + carry
	cpu.SetFlagH(halfCarry > 0x0F)
	cpu.SetFlagC(sum > 0xFF)
	*op0 = byte(sum)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	return 1
}

func ADC_r8_r16mem(cpu *CPU, op0 *byte, op1 R16Enum) int {
	if op1 != HL {
		cpu.Panic("unreachable ADC_r8_r16mem")
	}
	memVal := cpu.Bus.Read(cpu.HL())
	carry := 0
	if cpu.GetFlagC() {
		carry = 1
	}
	sum := int(*op0) + int(memVal) + carry
	halfCarry := (int(*op0) & 0x0F) + (int(memVal) & 0x0F) + carry
	cpu.SetFlagH(halfCarry > 0x0F)
	cpu.SetFlagC(sum > 0xFF)
	*op0 = byte(sum)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	return 2
}

func SUB_r8_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	cpu.SetFlagC(*op1 > *op0)
	cpu.SetFlagH((int(*op0) & 0x0F) < (int(*op1) & 0x0F))
	*op0 = *op0 - *op1
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(true)

	return 1
}

func SUB_r8_r16mem(cpu *CPU, op0 *byte, op1 R16Enum) int {
	if op1 != HL {
		cpu.Panic("unreachable SUB_r8_r16mem")
	}
	val := cpu.Bus.Read(cpu.HL())
	cpu.SetFlagC(val > *op0)
	cpu.SetFlagH((int(*op0) & 0x0F) < (int(val) & 0x0F))
	*op0 = *op0 - val
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(true)

	return 2
}

func SBC_r8_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	val := *op1
	carry := 0
	if cpu.GetFlagC() {
		carry = 1
	}
	sum := int(*op0) - int(val) - carry
	halfBorrow := int(*op0&0x0F) - int(val&0x0F) - carry
	cpu.SetFlagH(halfBorrow < 0)
	cpu.SetFlagC(sum < 0)
	*op0 = byte(sum)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(true)
	return 1
}

func SBC_r8_r16mem(cpu *CPU, op0 *byte, op1 R16Enum) int {
	if op1 != HL {
		cpu.Panic("unreachable SBC_r8_r16mem")
	}
	val := cpu.Bus.Read(cpu.HL())
	carry := 0
	if cpu.GetFlagC() {
		carry = 1
	}
	sum := int(*op0) - int(val) - carry
	halfBorrow := int(*op0&0x0F) - int(val&0x0F) - carry
	cpu.SetFlagH(halfBorrow < 0)
	cpu.SetFlagC(sum < 0)
	*op0 = byte(sum)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(true)
	return 2
}

func AND_r8_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	*op0 &= *op1
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(true)
	cpu.SetFlagC(false)
	return 1
}

func AND_r8_r16mem(cpu *CPU, op0 *byte, op1 R16Enum) int {
	if op1 != HL {
		cpu.Panic("unreachable AND_r8_r16mem")
	}
	*op0 &= cpu.Bus.Read(cpu.HL())
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(true)
	cpu.SetFlagC(false)
	return 2
}

func XOR_r8_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	*op0 ^= *op1
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(false)
	return 1
}

func XOR_r8_r16mem(cpu *CPU, op0 *byte, op1 R16Enum) int {
	switch op1 {
	case HL:
		val := cpu.Bus.Read(cpu.HL())
		*op0 ^= val
	default:
		cpu.Panic("unreachable XOR_r8_r16mem")

	}

	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(false)
	return 2
}

func OR_r8_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	*op0 |= *op1
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(false)
	return 1
}

func OR_r8_r16mem(cpu *CPU, op0 *byte, op1 R16Enum) int {
	switch op1 {
	case HL:
		val := cpu.Bus.Read(cpu.HL())

		*op0 |= val

		cpu.SetFlagZ(*op0 == 0)
		cpu.SetFlagN(false)
		cpu.SetFlagH(false)
		cpu.SetFlagC(false)

	default:
		cpu.Panic("unreachable OR_r8_r16mem")

	}
	return 2
}

func CP_r8_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	val := *op1

	cpu.SetFlagZ(*op0 == val)
	cpu.SetFlagN(true)
	cpu.SetFlagH((int(*op0) & 0x0F) < (int(val) & 0x0F))
	cpu.SetFlagC(val > *op0)
	return 1
}

func CP_r8_r16mem(cpu *CPU, op0 *byte, op1 R16Enum) int {
	switch op1 {
	case HL:
		val := cpu.Bus.Read(cpu.HL())

		cpu.SetFlagZ(*op0 == val)
		cpu.SetFlagN(true)
		cpu.SetFlagH(((*op0) & 0x0F) < ((val) & 0x0F))
		cpu.SetFlagC(*op0 < val)

	default:
		cpu.Panic("unreachable CP_r8_r16mem")

	}
	return 2
}

func RET_NZ(cpu *CPU) int {
	if cpu.GetFlagZ() == false {
		cpu.PC = cpu.Pop()
		return 5
	}
	return 2
}

func POP_r16(cpu *CPU, op0 R16Enum) int {
	val := cpu.Pop()
	switch op0 {
	case AF:
		cpu.SetAF(val)
	case BC:
		cpu.SetBC(val)
	case DE:
		cpu.SetDE(val)
	case HL:
		cpu.SetHL(val)
	case SP:
		cpu.SP = val
	default:
		cpu.Panic("unreachable POP_r16")
	}
	return 3
}

func JP_NZ_a16(cpu *CPU) int {
	val := cpu.PCRead16()
	if cpu.GetFlagZ() == false {
		cpu.PC = val
		return 4
	}
	return 3
}

func JP_a16(cpu *CPU) int {
	cpu.PC = cpu.PCRead16()
	return 4
}

func CALL_NZ_a16(cpu *CPU) int {
	targ := cpu.PCRead16()
	if cpu.GetFlagZ() == false {
		cpu.Push(cpu.PC)
		cpu.PC = targ
		return 6
	}
	return 3
}

func PUSH_r16(cpu *CPU, op0 R16Enum) int {
	switch op0 {
	case AF:
		cpu.Push(cpu.AF())
	case BC:
		cpu.Push(cpu.BC())
	case DE:
		cpu.Push(cpu.DE())
	case HL:
		cpu.Push(cpu.HL())
	case SP:
		cpu.Push(cpu.SP)
	default:
		cpu.Panic("unreachable PUSH_r16")
	}
	return 4
}

func ADD_r8_n8(cpu *CPU, op0 *byte) int {
	val := cpu.PCRead()
	sum := int(*op0) + int(val)
	cpu.SetFlagH((int(*op0)&0x0F)+(int(val)&0x0F) > 0x0F)
	cpu.SetFlagC(sum > 0xFF)
	*op0 = byte(sum)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	return 2
}

func RST_vec(cpu *CPU, vector uint16) int {
	// cpu.Panic("ya kotori dalbaeb")
	// fmt.Printf("rst %#x %#x\n", vector, cpu.Bus.Read(vector))
	cpu.Push(cpu.PC)
	cpu.PC = vector
	return 4
}

func RET_Z(cpu *CPU) int {
	if cpu.GetFlagZ() {
		cpu.PC = cpu.Pop()
		return 5
	}
	return 2
}

func RET(cpu *CPU) int {
	cpu.PC = cpu.Pop()
	return 4
}

func JP_Z_a16(cpu *CPU) int {
	val := cpu.PCRead16()
	if cpu.GetFlagZ() {
		cpu.PC = val
		return 4
	}
	return 3
}

func PREFIX(cpu *CPU) int {
	cpu.Panic("PREFIX not implemented!")
	return 0
}

func CALL_Z_a16(cpu *CPU) int {
	targ := cpu.PCRead16()
	if cpu.GetFlagZ() {
		cpu.Push(cpu.PC)
		cpu.PC = targ
		return 6
	}
	return 3
}

func CALL_a16(cpu *CPU) int {
	targ := cpu.PCRead16()
	cpu.Push(cpu.PC)
	cpu.PC = targ
	return 6
}

func ADC_r8_n8(cpu *CPU, op0 *byte) int {
	val := cpu.PCRead()
	carry := 0
	if cpu.GetFlagC() {
		carry = 1
	}

	cpu.SetFlagH((int(cpu.A)&0x0F)+(int(val)&0x0F)+carry > 0x0F)
	cpu.SetFlagC(int(cpu.A)+int(val)+carry > 0xFF)

	res := cpu.A + val + byte(carry)
	cpu.A = res

	cpu.SetFlagZ(cpu.A == 0)
	cpu.SetFlagN(false)

	return 2
}

func RET_NC(cpu *CPU) int {
	if cpu.GetFlagC() == false {
		cpu.PC = cpu.Pop()
		return 5
	}
	return 2
}

func JP_NC_a16(cpu *CPU) int {
	val := cpu.PCRead16()
	if cpu.GetFlagC() == false {
		cpu.PC = val
		return 4
	}
	return 3
}

func ILLEGAL_D3(cpu *CPU) int {
	cpu.Panic("ILLEGAL_D3 not implemented!")
	return 0
}

func CALL_NC_a16(cpu *CPU) int {
	targ := cpu.PCRead16()
	if cpu.GetFlagC() == false {
		cpu.Push(cpu.PC)
		cpu.PC = targ
		return 6
	}
	return 3
}

func SUB_r8_n8(cpu *CPU, op0 *byte) int {
	val := cpu.PCRead()
	sum := int(*op0) - int(val)
	cpu.SetFlagH((int(*op0) & 0x0F) < (int(val) & 0x0F))
	cpu.SetFlagC(val > *op0)
	*op0 = byte(sum)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(true)
	return 2
}

func RET_C(cpu *CPU) int {
	if cpu.GetFlagC() {
		cpu.PC = cpu.Pop()
		return 5
	}
	return 2
}

func RETI(cpu *CPU) int {
	cpu.IME = true
	cpu.PC = cpu.Pop()
	return 4
}

func JP_C_a16(cpu *CPU) int {
	val := cpu.PCRead16()
	if cpu.GetFlagC() {
		cpu.PC = val
		return 4
	}
	return 3
}

func ILLEGAL_DB(cpu *CPU) int {
	cpu.Panic("ILLEGAL_DB not implemented!")
	return 0
}

func CALL_C_a16(cpu *CPU) int {
	targ := cpu.PCRead16()
	if cpu.GetFlagC() {
		cpu.Push(cpu.PC)
		cpu.PC = targ
		return 6
	}
	return 3
}

func ILLEGAL_DD(cpu *CPU) int {
	cpu.Panic("ILLEGAL_DD not implemented!")
	return 0
}

func SBC_r8_n8(cpu *CPU, op0 *byte) int {
	val := cpu.PCRead()
	carry := 0
	if cpu.GetFlagC() {
		carry = 1
	}
	sum := int(*op0) - int(val) - carry
	halfBorrow := int(*op0&0x0F) - int(val&0x0F) - carry
	cpu.SetFlagH(halfBorrow < 0)
	cpu.SetFlagC(sum < 0)
	*op0 = byte(sum)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(true)
	return 2
}

func LDH_a8_r8(cpu *CPU, op1 *byte) int {
	offset := cpu.PCRead()
	cpu.Bus.Write(0xFF00+uint16(offset), *op1)
	return 3
}

func LDH_r8_r8mem(cpu *CPU, op0 *byte, op1 *byte) int {
	*op0 = cpu.Bus.Read((0xFF00 + uint16(*op1)))
	return 2
}

func LDH_r8mem_r8(cpu *CPU, op0 *byte, op1 *byte) int {
	cpu.Bus.Write(0xFF00+uint16(*op0), *op1)
	return 2
}

func ILLEGAL_E3(cpu *CPU) int {
	cpu.Panic("ILLEGAL_E3 not implemented!")
	return 0
}

func ILLEGAL_E4(cpu *CPU) int {
	cpu.Panic("ILLEGAL_E4 not implemented!")
	return 0
}

func AND_r8_n8(cpu *CPU, op0 *byte) int {
	val := cpu.PCRead()
	*op0 &= val
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(true)
	cpu.SetFlagC(false)
	return 2
}

func ADD_r16_e8(cpu *CPU, op0 R16Enum) int {
	if op0 != SP {
		cpu.Panic("unreachable ADD_r16_e8")
	}
	rawOff := cpu.PCRead()
	off := int8(rawOff)

	spLow := byte(cpu.SP & 0xFF)

	cpu.SetFlagC(int(spLow)+int(rawOff) > 0xFF)
	cpu.SetFlagH((int(spLow)&0x0F)+(int(off)&0x0F) > 0x0F)

	cpu.SP = uint16(int(cpu.SP) + int(off))

	cpu.SetFlagZ(false)
	cpu.SetFlagN(false)
	return 4
}

func JP_r16(cpu *CPU, op0 R16Enum) int {
	switch op0 {
	case HL:
		cpu.PC = cpu.HL()
	default:
		cpu.Panic("unreachable JP_r16")

	}
	return 1
}

func LD_a16_r8(cpu *CPU, op1 *byte) int {
	addr := cpu.PCRead16()
	cpu.Bus.Write(addr, *op1)
	return 4
}

func ILLEGAL_EB(cpu *CPU) int {
	cpu.Panic("ILLEGAL_EB not implemented!")
	return 0
}

func ILLEGAL_EC(cpu *CPU) int {
	cpu.Panic("ILLEGAL_EC not implemented!")
	return 0
}

func ILLEGAL_ED(cpu *CPU) int {
	cpu.Panic("ILLEGAL_ED not implemented!")
	return 0
}

func XOR_r8_n8(cpu *CPU, op0 *byte) int {
	*op0 ^= cpu.PCRead()
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(false)
	return 2
}

func LDH_r8_a8(cpu *CPU, op0 *byte) int {
	offset := cpu.PCRead()

	*op0 = cpu.Bus.Read(0xFF00 + uint16(offset))
	return 3
}

func DI(cpu *CPU) int {
	cpu.IME = false
	cpu.IMEDelay = false
	return 1
}

func ILLEGAL_F4(cpu *CPU) int {
	cpu.Panic("ILLEGAL_F4 not implemented!")
	return 0
}

func OR_r8_n8(cpu *CPU, op0 *byte) int {
	val := cpu.PCRead()
	*op0 |= val
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(false)
	return 2
}

func LD_r16_r16inc_e8(cpu *CPU, op0 R16Enum, op1 R16Enum) int {
	// cpu.Panic("Ya dalbaeb")
	imm8 := cpu.PCRead()
	e8 := int8(imm8)
	if op1 != SP {
		cpu.Panic("unreachable LD_r16_r16inc_e8")
	}
	res := uint16(int(cpu.SP) + int(e8))
	cpu.SetHL(res)
	spLow := byte(cpu.SP & 0xFF)

	cpu.SetFlagZ(false)
	cpu.SetFlagN(false)
	cpu.SetFlagH((int(spLow)&0x0F)+(int(imm8)&0x0F) > 0x0F)
	cpu.SetFlagC(int(spLow)+int(imm8) > 0xFF)

	return 3
}

func LD_r16_r16(cpu *CPU, op0 R16Enum, op1 R16Enum) int {
	// always sp hl
	if op0 != SP {
		cpu.Panic("unreachable LD_r16_r16")
	}
	cpu.SP = cpu.HL()
	return 2
}

func LD_r8_a16(cpu *CPU, op0 *byte) int {
	addr := cpu.PCRead16()
	val := cpu.Bus.Read(addr)
	*op0 = val
	return 4
}

func EI(cpu *CPU) int {
	cpu.IMEDelay = true
	return 1
}

func ILLEGAL_FC(cpu *CPU) int {
	cpu.Panic("ILLEGAL_FC not implemented!")
	return 0
}

func ILLEGAL_FD(cpu *CPU) int {
	cpu.Panic("ILLEGAL_FD not implemented!")
	return 0
}

func CP_r8_n8(cpu *CPU, op0 *byte) int {
	val := cpu.PCRead()

	cpu.SetFlagZ(*op0 == val)
	cpu.SetFlagN(true)
	cpu.SetFlagH((int(*op0) & 0x0F) < (int(val) & 0x0F))
	cpu.SetFlagC(val > *op0)
	return 2
}

func RLC_r8(cpu *CPU, op0 *byte) int {
	bit := *op0 >> 7
	*op0 = *op0 << 1
	*op0 = *op0 | bit

	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(bit == 1)
	return 2
}

func RLC_r16mem(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable RLC_r16mem")
	}
	val := cpu.Bus.Read(cpu.HL())
	bit := val >> 7
	val = val << 1
	val = val | bit

	cpu.Bus.Write(cpu.HL(), val)

	cpu.SetFlagZ(val == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(bit == 1)
	return 4
}

func RRC_r8(cpu *CPU, op0 *byte) int {
	bit0 := *op0 & 0x01

	*op0 = (*op0 >> 1) | (bit0 << 7)

	cpu.SetFlagC(bit0 == 1)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)

	return 2
}

func RRC_r16mem(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable RRC_r16mem")
	}
	val := cpu.Bus.Read(cpu.HL())
	bit0 := val & 0x01

	val = (val >> 1) | (bit0 << 7)
	cpu.Bus.Write(cpu.HL(), val)
	cpu.SetFlagC(bit0 == 1)
	cpu.SetFlagZ(val == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	return 4
}

func RL_r8(cpu *CPU, op0 *byte) int {
	bit := (*op0 & (1 << 7)) != 0
	*op0 = *op0 << 1
	if cpu.GetFlagC() {
		*op0 = *op0 | 0x01
	}
	cpu.SetFlagC(bit)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	return 2
}

func RL_r16mem(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable RL_r16mem")
	}
	val := cpu.Bus.Read(cpu.HL())
	bit := (val & (1 << 7)) != 0
	val = val << 1
	if cpu.GetFlagC() {
		val = val | 0x01
	}

	cpu.Bus.Write(cpu.HL(), val)
	cpu.SetFlagC(bit)
	cpu.SetFlagZ(val == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	return 4
}

func RR_r8(cpu *CPU, op0 *byte) int {
	nextCarry := (*op0 & 0x01) != 0
	*op0 = *op0 >> 1
	if cpu.GetFlagC() {
		*op0 = *op0 | 0x80
	}
	cpu.SetFlagC(nextCarry)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	return 2
}

func RR_r16mem(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable RR_r16mem")
	}
	val := cpu.Bus.Read(cpu.HL())

	nextCarry := (val & 0x01) != 0
	val = val >> 1
	if cpu.GetFlagC() {
		val = val | 0x80
	}
	cpu.Bus.Write(cpu.HL(), val)
	cpu.SetFlagC(nextCarry)
	cpu.SetFlagZ(val == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	return 4
}

func SLA_r8(cpu *CPU, op0 *byte) int {
	bit := *op0 >> 7
	*op0 = *op0 << 1
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(bit == 1)
	return 2
}

func SLA_r16mem(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable SLA_r16mem")
	}
	val := cpu.Bus.Read(cpu.HL())
	bit := val >> 7
	val = val << 1
	cpu.SetFlagZ(val == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(bit == 1)
	cpu.Bus.Write(cpu.HL(), val)
	return 4
}

func SRA_r8(cpu *CPU, op0 *byte) int {
	bit0 := *op0 & 0x01
	bit7 := *op0 & 0x80
	*op0 = (*op0 >> 1) | bit7
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(bit0 == 1)
	return 2
}

func SRA_r16mem(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable SRA_r16mem")
	}

	val := cpu.Bus.Read(cpu.HL())

	bit0 := val & 0x01
	bit7 := val & 0x80
	val = (val >> 1) | bit7

	cpu.Bus.Write(cpu.HL(), val)

	cpu.SetFlagZ(val == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(bit0 == 1)
	return 4
}

func SWAP_r8(cpu *CPU, op0 *byte) int {
	up := *op0 & 0xF0
	low := *op0 & 0x0F

	*op0 = (low << 4) | (up >> 4)

	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(false)
	return 2
}

func SWAP_r16mem(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable SWAP_r16mem")
	}
	val := cpu.Bus.Read(cpu.HL())

	up := val & 0xF0
	low := val & 0x0F

	val = (low << 4) | (up >> 4)

	cpu.Bus.Write(cpu.HL(), val)

	cpu.SetFlagZ(val == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(false)
	return 4
}

func SRL_r8(cpu *CPU, op0 *byte) int {
	bit0 := (*op0 & 0x01) != 0
	*op0 = *op0 >> 1
	cpu.SetFlagC(bit0)
	cpu.SetFlagZ(*op0 == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	return 2
}

func SRL_r16mem(cpu *CPU, op0 R16Enum) int {
	if op0 != HL {
		cpu.Panic("unreachable SRL_r16mem")
	}
	val := cpu.Bus.Read(cpu.HL())

	bit0 := (val & 0x01) != 0
	val = val >> 1

	cpu.Bus.Write(cpu.HL(), val)

	cpu.SetFlagC(bit0)
	cpu.SetFlagZ(val == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	return 4
}

func BIT_u3_r8(cpu *CPU, u3 uint8, op1 *byte) int {
	cpu.SetFlagZ((*op1 & (1 << u3)) == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(true)
	return 2
}

func BIT_u3_r16mem(cpu *CPU, u3 uint8, op1 R16Enum) int {
	if op1 != HL {
		cpu.Panic("unreachable BIT_u3_r16mem")
	}
	byte_ := cpu.Bus.Read(cpu.HL())
	cpu.SetFlagZ((byte_ & (1 << u3)) == 0)
	cpu.SetFlagN(false)
	cpu.SetFlagH(true)
	return 3
}

func RES_u3_r8(cpu *CPU, u3 uint8, op1 *byte) int {
	*op1 = *op1 &^ (1 << u3)
	return 2
}

func RES_u3_r16mem(cpu *CPU, u3 uint8, op1 R16Enum) int {
	if op1 != HL {
		cpu.Panic("unreachable RES_u3_r16mem")
	}
	byte_ := cpu.Bus.Read(cpu.HL())
	byte_ = byte_ &^ (1 << u3)
	cpu.Bus.Write(cpu.HL(), byte_)
	return 4
}

func SET_u3_r8(cpu *CPU, u3 uint8, op1 *byte) int {
	*op1 = *op1 | (1 << u3)
	return 2
}

func SET_u3_r16mem(cpu *CPU, u3 uint8, op1 R16Enum) int {
	if op1 != HL {
		cpu.Panic("unreachable SET_u3_r16mem")
	}
	byte_ := cpu.Bus.Read(cpu.HL())
	byte_ = byte_ | (1 << u3)
	cpu.Bus.Write(cpu.HL(), byte_)
	return 4
}

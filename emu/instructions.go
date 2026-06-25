package emu

type Instruction struct {
	Name   string
	Bytes  int
	Cycles int
	Exec   func(*CPU) int
}

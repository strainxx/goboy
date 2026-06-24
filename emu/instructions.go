package emu

type Instruction struct {
	Name  string
	Bytes int
	Exec  func(*CPU) int
}

package main

// OpType : OpType of each instructions required
type OpType int

// Define Opcode for all instructions
const (
	LUIOpcode    = "0110111"
	AUIPCOpcode  = "0010111"
	JALOpcode    = "1101111"
	JALROpcode   = "1100111"
	BTypeOpcode  = "1100011"
	IType1Opcode = "0000011"
	IType2Opcode = "0010011"
	STypeOpcode  = "0100011"
	RTypeOpcode  = "0110011"
)

// Create enum OpType for each operation
const (
	// U-Type Instruction
	LUI OpType = iota
	AUIPC
	// J-Type Instruction
	JAL
	// B-Type Instruction
	BEQ
	BNE
	BLT
	BGE
	BLTU
	BGEU
	// S-Type Instruction
	SB
	SH
	SW
	// I-Type Instructions
	JALR
	LB
	LH
	LW
	LBU
	LHU
	SLLI
	SRLI
	SRAI
	ADDI
	SLTI
	SLTIU
	XORI
	ORI
	ANDI
	// R-Type Instructions
	ADD
	SUB
	SLL
	SLT
	SLTU
	XOR
	SRL
	SRA
	OR
	AND
)

// UTypeInstr : Store the decoded parts of U Type instruction
type UTypeInstr struct {
	instrType OpType
	imm       int
	rd        int
}

// JTypeInstr : Store the decoded parts of J Type instruction
type JTypeInstr struct {
	instrType OpType
	imm       int
	rd        int
}

// BTypeInstr : Store the decoded parts of B Type instruction
type BTypeInstr struct {
	instrType OpType
	imm1      int
	imm2      int
	r1        int
	r2        int
}

// STypeInstr : Store the decoded parts of S Type instruction
type STypeInstr struct {
	instrType OpType
	imm1      int
	imm2      int
	r1        int
	r2        int
}

// ITypeInstr : Store the decoded parts of I Type instruction
type ITypeInstr struct {
	instrType OpType
	imm1      int
	rs1       int
	rd        int
}

// RTypeInstr : Store the decoded parts of R Type instruction
type RTypeInstr struct {
	instType OpType
	r1       int
	r2       int
	rd       int
}

// DecodedInstr : Interface for all structs of decoded instructions
type DecodedInstr interface{}

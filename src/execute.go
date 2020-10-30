package main

import (
	"math"
)

// Executor : Simulator of Exec Stage in CPU
type Executor struct {
}

// Execute : Execute the decoded instruction
func (e Executor) Execute(dInstr DecodedInstr) {
	switch dInstr.(InstrType) {
	case UType:
		executeUTypeInstr(dInstr)
	case JType:
		executeJTypeInstr(dInstr)
	case IType:
		executeITypeInstr(dInstr)
	case BType:
		executeBTypeInstr(dInstr)
	case SType:
		executeSTypeInstr(dInstr)
	case RType:
		executeRTypeInstr(dInstr)
	}
}

func executeUTypeInstr(dInstr DecodedInstr) {
	UTypeDecodedInstr, _ := dInstr.(UTypeInstr)
	switch UTypeDecodedInstr.opName {
	case LUI:
		value := int64(UTypeDecodedInstr.imm << int(math.Pow(2, 12)))
		register[UTypeDecodedInstr.rd] = value
	case AUIPC:
		
	}
}

func executeJTypeInstr(dInstr DecodedInstr) {
	JTypeDecodedInstr, _ := dInstr.(JTypeInstr)
	switch JTypeDecodedInstr.opName {
	case JAL:
	}
}
func executeITypeInstr(dInstr DecodedInstr) {
	ITypeDecodedInstr, _ := dInstr.(ITypeInstr)
	switch ITypeDecodedInstr.opName {
	case JALR:
	case LB:
	case LH:
	case LW:
	case LBU:
	case LHU:
	case SLLI:
	case SRLI:
	case SRAI:
	case ADDI:
	case SLTI:
	case SLTIU:
	case XORI:
	case ORI:
	case ANDI:
	}
}

func executeBTypeInstr(dInstr DecodedInstr) {
	BTypeDecodedInstr, _ := dInstr.(BTypeInstr)
	switch BTypeDecodedInstr.opName {
	case BEQ:
	case BNE:
	case BLT:
	case BGE:
	case BLTU:
	case BGEU:
	}
}

func executeSTypeInstr(dInstr DecodedInstr) {
	STypeDecodedInstr, _ := dInstr.(STypeInstr)
	switch STypeDecodedInstr.opName {
	case SB:
	case SH:
	case SW:
	}
}

func executeRTypeInstr(dInstr DecodedInstr) {
	RTypeDecodedInstr, _ := dInstr.(RTypeInstr)
	switch RTypeDecodedInstr.opName {
	case ADD:
	case SUB:
	case SLL:
	case SLT:
	case SLTU:
	case XOR:
	case SRL:
	case SRA:
	case OR:
	case AND:
	}
}

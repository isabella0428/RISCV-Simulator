package main

import (
	"math"
)

// Decoder : Simulator for Decode Stage in CPU
type Decoder struct {
}

// Decode : decode the instructions into immediates and registers
func (d Decoder) Decode(binaryInstr []byte) DecodedInstr {
	instr := convertInstructionToBinaryString(binaryInstr)
	opcode := instr[:8]

	// var di DecodedInstr
	switch opcode {
	case LUIOpcode:
		// LUI
		return UTypeInstr{LUI, UType, biToInt(instr[12:32]), biToInt(instr[8:12])}
	case AUIPCOpcode:
		// AUIPC
		return UTypeInstr{AUIPC, UType, biToInt(instr[12:32]), biToInt(instr[8:12])}
	case JALOpcode:
		// JAL
		return JTypeInstr{JAL, JType, biToInt(instr[12:32]), biToInt(instr[8:12])}
	case JALROpcode:
		// JALR
		return ITypeInstr{JALR, IType, biToInt(instr[20:32]), biToInt(instr[15:20]), biToInt(instr[7:12])}
	case BTypeOpcode:
		// B-Type Instruction
		return selectFromBTypeInstruction(instr)
	case IType1Opcode:
		// I-Type Instruction
		return selectFromITypeInstruction(instr)
	case IType2Opcode:
		// I-Type Instruction
		return selectFromITypeInstruction(instr)
	case STypeOpcode:
		// S-Type Instruction
		return selectFromSTypeInstruction(instr)
	case RTypeOpcode:
		// R-Type Instruction
		return selectFromRTypeInstruction(instr)
	default:
		return nil
	}
}

func selectFromBTypeInstruction(instr string) DecodedInstr {
	funct3 := instr[7:12]
	bInstr := BTypeInstr{
		-1,
		BType,
		biToInt(instr[25:32]),
		biToInt(instr[7:12]),
		biToInt(instr[15:20]),
		biToInt(instr[20:25]),
	}

	switch funct3 {
	case "000":
		bInstr.opName = BEQ
		break
	case "001":
		bInstr.opName = BNE
		break
	case "100":
		bInstr.opName = BLT
		break
	case "101":
		bInstr.opName = BGE
		break
	case "110":
		bInstr.opName = BLTU
		break
	case "111":
		bInstr.opName = BGEU
		break
	}
	return bInstr
}

func selectFromITypeInstruction(instr string) DecodedInstr {
	iInstr := ITypeInstr{-1, IType, biToInt(instr[20:32]), biToInt(instr[15:20]), biToInt(instr[7:12])}
	funct3 := instr[12:15]

	opCode := instr[0:8]

	if opCode == IType1Opcode {
		switch funct3 {
		case "000":
			iInstr.opName = LB
			break
		case "001":
			iInstr.opName = LH
			break
		case "010":
			iInstr.opName = LW
			break
		case "100":
			iInstr.opName = LBU
			break
		case "101":
			iInstr.opName = LHU
			break
		}
	} else {
		switch funct3 {
		// TODO: multiple
		case "000":
			iInstr.opName = ADDI
			break
		case "001":
			iInstr.opName = SLLI
			iInstr.imm1 = -1
			break
		case "010":
			iInstr.opName = SLTI
			break
		case "011":
			iInstr.opName = SLTIU
			break
		case "100":
			iInstr.opName = XORI
			break
		case "101":
			if instr[25:32] == "0000000" {
				iInstr.opName = SRLI
			} else {
				iInstr.opName = SRAI
			}
			iInstr.imm1 = -1
			break
		case "110":
			iInstr.opName = ORI
			break
		case "111":
			iInstr.opName = ANDI
			break
		}
	}

	return iInstr
}

func selectFromSTypeInstruction(instr string) DecodedInstr {
	sInstr := STypeInstr{
		-1,
		SType,
		biToInt(instr[5:32]),
		biToInt(instr[7:12]),
		biToInt(instr[15:20]),
		biToInt(instr[20:25]),
	}

	func3 := instr[12:15]

	switch func3 {
	case "000":
		sInstr.opName = SB
		break
	case "001":
		sInstr.opName = SH
		break
	case "010":
		sInstr.opName = SW
		break
	}
	return sInstr
}

func selectFromRTypeInstruction(instr string) DecodedInstr {
	rInstr := RTypeInstr{
		-1,
		RType,
		biToInt(instr[15:20]),
		biToInt(instr[20:25]),
		biToInt(instr[7:12]),
	}

	func3 := instr[12:15]

	switch func3 {
	case "000":
		if instr[30] == '0' {
			rInstr.opName = ADD
		} else {
			rInstr.opName = SUB
		}
		break
	case "001":
		rInstr.opName = SLL
		break
	case "010":
		rInstr.opName = SLT
		break
	case "011":
		rInstr.opName = SLTU
		break
	case "100":
		rInstr.opName = XOR
		break
	case "101":
		if instr[30] == '0' {
			rInstr.opName = SRL
		} else {
			rInstr.opName = SRA
		}
		break
	case "110":
		rInstr.opName = OR
		break
	case "111":
		rInstr.opName = AND
		break
	}

	return rInstr
}

func convertInstructionToBinaryString(binaryInstr []byte) string {
	convertedString := ""
	for i := 0; i < 4; i++ {
		curByte := float64(binaryInstr[i])

		for j := 7; j >= 0; j-- {
			if curByte/(math.Pow(2, float64(j))) > 1 {
				convertedString += "1"
			} else {
				convertedString += "0"
			}
		}
	}
	return convertedString
}

func biToInt(binaryStr string) int {
	value := 0
	for _, c := range binaryStr {
		value *= 2
		if c == '1' {
			value++
		}
	}
	return value
}

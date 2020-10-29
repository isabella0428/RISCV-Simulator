package main

import (
	"math"
	"strconv"
)

// Decoder : object for decoding the instructions
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
		return UTypeInstr{LUI, biToInt(instr[12:32]), biToInt(instr[8:12])}
	case AUIPCOpcode:
		// AUIPC
		return UTypeInstr{AUIPC, biToInt(instr[12:32]), biToInt(instr[8:12])}
	case JALOpcode:
		// JAL
		return JTypeInstr{JAL, biToInt(instr[12:32]), biToInt(instr[8:12])}
	case JALROpcode:
		// JALR
		return ITypeInstr{JALR, biToInt(instr[20:32]), biToInt(instr[15:20]), biToInt(instr[7:12])}
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
		biToInt(instr[25:32]),
		biToInt(instr[7:12]),
		biToInt(instr[15:20]),
		biToInt(instr[20:25]),
	}

	switch funct3 {
	case "000":
		bInstr.instrType = BEQ
		break
	case "001":
		bInstr.instrType = BNE
		break
	case "100":
		bInstr.instrType = BLT
		break
	case "101":
		bInstr.instrType = BGE
		break
	case "110":
		bInstr.instrType = BLTU
		break
	case "111":
		bInstr.instrType = BGEU
		break
	}
	return bInstr
}

func selectFromITypeInstruction(instr string) DecodedInstr {
	iInstr := ITypeInstr{-1, biToInt(instr[20:32]), biToInt(instr[15:20]), biToInt(instr[7:12])}
	funct3 := instr[12:15]

	opCode := instr[0:8]

	if opCode == IType1Opcode {
		switch funct3 {
		case "000":
			iInstr.instrType = LB
			break
		case "001":
			iInstr.instrType = LH
			break
		case "010":
			iInstr.instrType = LW
			break
		case "100":
			iInstr.instrType = LBU
			break
		case "101":
			iInstr.instrType = LHU
			break
		}
	} else {
		switch funct3 {
		// TODO: multiple
		case "000":
			iInstr.instrType = ADDI
			break
		case "001":
			iInstr.instrType = SLLI
			iInstr.imm1 = -1
			break
		case "010":
			iInstr.instrType = SLTI
			break
		case "011":
			iInstr.instrType = SLTIU
			break
		case "100":
			iInstr.instrType = XORI
			break
		case "101":
			if instr[25:32] == "0000000" {
				iInstr.instrType = SRLI
			} else {
				iInstr.instrType = SRAI
			}
			iInstr.imm1 = -1
			break
		case "110":
			iInstr.instrType = ORI
			break
		case "111":
			iInstr.instrType = ANDI
			break
		}
	}

	return iInstr
}

func selectFromSTypeInstruction(instr string) DecodedInstr {
	sInstr := STypeInstr{
		-1,
		biToInt(instr[5:32]),
		biToInt(instr[7:12]),
		biToInt(instr[15:20]),
		biToInt(instr[20:25]),
	}

	func3 := instr[12:15]

	switch func3 {
	case "000":
		sInstr.instrType = SB
		break
	case "001":
		sInstr.instrType = SH
		break
	case "010":
		sInstr.instrType = SW
		break
	}
	return sInstr
}

func selectFromRTypeInstruction(instr string) DecodedInstr {
	rInstr := RTypeInstr{
		-1,
		biToInt(instr[15:20]),
		biToInt(instr[20:25]),
		biToInt(instr[7:12]),
	}

	func3 := instr[12:15]

	switch func3 {
	case "000":
		if instr[30] == '0' {
			rInstr.instType = ADD
		} else {
			rInstr.instType = SUB
		}
		break
	case "001":
		rInstr.instType = SLL
		break
	case "010":
		rInstr.instType = SLT
		break
	case "011":
		rInstr.instType = SLTU
		break
	case "100":
		rInstr.instType = XOR
		break
	case "101":
		if instr[30] == '0' {
			rInstr.instType = SRL
		} else {
			rInstr.instType = SRA
		}
		break
	case "110":
		rInstr.instType = OR
		break
	case "111":
		rInstr.instType = AND
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
	val, _ := strconv.ParseInt(binaryStr, 2, 8)
	return int(val)
}

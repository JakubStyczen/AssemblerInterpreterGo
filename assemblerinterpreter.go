package assemblerinterpretergo

// This repo is implementation of Assembler interpreter (part II) kata
// (https://www.codewars.com/kata/assembler-interpreter-part-ii/) for golang (not available on codewars.com)

import (
	"fmt"
)

// Main struct for whole package. Contains:
// - parsed program (with length),
// - program registers (mapped var name and value),
// - labels mapped to indexes in program
// - program output (from msg instruction)
// - cmp value (from cmp instruction)
// - pre_call_pc (pc of call instruction to return when ret is found)
// - pc (program counter)

type AssemblerInterpreter struct {
	program          string
	program_length   int
	instruction_list []string
	labels_idx       map[string]int
	registers        map[string]int
	output           string
	cmp              int
	pre_call_pc      int
	pc               int
}

type Result struct {
	output    string
	exit_code int
}

func NewAssemblerInterpreter(program string) *AssemblerInterpreter {
	fmt.Printf("Program to execute:\n%s\n", program)
	assInt := AssemblerInterpreter{}
	assInt.program = program
	assInt.registers = make(map[string]int)
	assInt.instruction_list, assInt.program_length = ParseProgram(program)
	assInt.labels_idx = ParseLabels(assInt.instruction_list)
	assInt.pre_call_pc = -1
	return &assInt
}

func (assInt *AssemblerInterpreter) Run() Result {
	for {
		if assInt.pc >= len(assInt.instruction_list) {
			return Result{output: "", exit_code: -1}
		}
		instruction, args := PrepareInstruction(assInt.instruction_list[assInt.pc])
		if instruction == "end" {
			break
		}
		assInt.ExecuteInstruction(instruction, args)
	}
	return Result{output: assInt.output, exit_code: 0}
}

func (assInt *AssemblerInterpreter) ExecuteInstruction(instruction string, args []string) {
	INSTRUCTIONS[instruction](assInt, args)
}

// func (assInt *AssemblerInterpreter) ExecuteInstruction(instruction string, args []string) {
// 	switch instruction {
// 	case "mov":
// 		assInt.Mov(args)
// 	case "inc":
// 		assInt.Inc(args)
// 	case "dec":
// 		assInt.Dec(args)
// 	case "add":
// 		assInt.Add(args)
// 	case "sub":
// 		assInt.Sub(args)
// 	case "mul":
// 		assInt.Mul(args)
// 	case "div":
// 		assInt.Div(args)
// 	case "nop":
// 		assInt.Nop(args)
// 	case "jmp":
// 		assInt.Jmp(args)
// 	case "jne":
// 		assInt.Jne(args)
// 	case "je":
// 		assInt.Je(args)
// 	case "jg":
// 		assInt.Jg(args)
// 	case "jge":
// 		assInt.Jge(args)
// 	case "jle":
// 		assInt.Jle(args)
// 	case "jl":
// 		assInt.Jl(args)
// 	case "cmp":
// 		assInt.Cmp(args)
// 	case "call":
// 		assInt.Call(args)
// 	case "msg":
// 		assInt.Msg(args)
// 	case "ret":
// 		assInt.Ret(args)
// 	case "end":
// 		assInt.End(args)
// 	default:
// 		assInt.pc++
// 	}
// }

func main() {
	var program string = "mov a, 5\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\nje kox\n ret\nkox:\nmov c, 3\ninc a\nret"

	assInt := NewAssemblerInterpreter(program)
	assInt.Run()
	fmt.Printf("Instruction list: %v\n", assInt.instruction_list)
	fmt.Printf("Program len: %v\n", assInt.program_length)
	fmt.Printf("Labels Idx: %v\n", assInt.labels_idx)
	fmt.Printf("Registers: %v\n", assInt.registers)
}

package main

import (
	"fmt"
)

type AssemblerInterpreter struct {
	program          string
	program_length   int
	instruction_list []string
	labels_idx       map[string]int
	registers        map[string]int
	output           string
	cmp              int
	pc               int
}

func NewAssemblerInterpreter(program string) *AssemblerInterpreter {
	fmt.Printf("Program to execute:\n%s\n", program)
	assInt := AssemblerInterpreter{}
	assInt.program = program
	assInt.registers = make(map[string]int)
	assInt.instruction_list, assInt.program_length = ParseProgram(program)
	assInt.labels_idx = ParseCustomSubroutines(assInt.instruction_list)
	return &assInt
}

func (assInt *AssemblerInterpreter) Run() string {
	for assInt.pc < assInt.program_length {
		instruction, args := PrepareInstruction(assInt.instruction_list[assInt.pc])
		fmt.Printf("sub: %v %v %v\n", assInt.pc, instruction, args)
		assInt.ExecuteInstruction(instruction, args)
	}
	return assInt.output
}

func (assInt *AssemblerInterpreter) ExecuteInstruction(instruction string, args []string) {
	switch instruction {
	case "mov":
		assInt.Mov(args)
	case "inc":
		assInt.Inc(args)
	case "dec":
		assInt.Dec(args)
	case "add":
		assInt.Add(args)
	case "sub":
		assInt.Sub(args)
	case "mul":
		assInt.Mul(args)
	case "div":
		assInt.Div(args)
	case "nop":
		assInt.Nop(args)
	case "jmp":
		assInt.Jmp(args)
	case "jne":
		assInt.Jne(args)
	case "je":
		assInt.Je(args)
	case "jg":
		assInt.Jg(args)
	case "jge":
		assInt.Jge(args)
	case "jle":
		assInt.Jle(args)
	case "jl":
		assInt.Jl(args)
	case "cmp":
		assInt.Cmp(args)
	case "call":
		assInt.Call(args)
	case "msg":
		assInt.Msg(args)
	case "ret":
		return
	case "end":
		assInt.End(args)
	default:
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) ExecuteSubroutine(subroutine_idx int) {
	i := subroutine_idx
	curr_pc := assInt.pc
	for {
		i++ // Skip the subroutine label
		instruction, args := PrepareInstruction(assInt.instruction_list[i])
		if instruction == "ret" {
			break
		}
		assInt.ExecuteInstruction(instruction, args)
	}
	assInt.pc = curr_pc + 1
}

// func getInstruction(instruction string) func(ai *AssemblerInterpreter, args []string) {
// 	return INSTRUCTIONS[instruction]
// }

func main() {
	// var program string = `; My first program
	// mov  a, 5
	// inc  a
	// call function
	// msg  '(5+1)/2 = ', a    ; output message
	// end

	// function:
	// 	div  a, 2
	// 	ret`
	// var program string = "mov a, 5\ninc a\nend"
	// var program string = "mov a, 5\njmp kox\nmov b, 3\nnop\nkox:\nmov c, 3\nend"
	var program string = "mov a, 5\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\nje kox\n ret\nkox:\nmov c, 3\ninc a\nret"

	assInt := NewAssemblerInterpreter(program)
	assInt.Run()
	fmt.Printf("Instruction list: %v\n", assInt.instruction_list)
	fmt.Printf("Program len: %v\n", assInt.program_length)
	fmt.Printf("Labels Idx: %v\n", assInt.labels_idx)
	fmt.Printf("Registers: %v\n", assInt.registers)
}

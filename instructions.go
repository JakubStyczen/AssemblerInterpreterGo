package main

import (
	"fmt"
	"strings"
)

func (assInt *AssemblerInterpreter) Mov(args []string) {
	// mov x, y - copy y (either an integer or the value of a register) into register x.
	assInt.registers[args[0]] = ParseArg(args[1], assInt.registers)
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Inc(args []string) {
	// inc x - increase the content of register x by one.
	assInt.registers[args[0]]++
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Dec(args []string) {
	// dec x - decrease the content of register x by one.
	assInt.registers[args[0]]--
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Add(args []string) {
	// add x, y - add the content of the register x with y (either an integer or the value of a register) and stores the result in x (i.e. register[x] += y).
	assInt.registers[args[0]] += ParseArg(args[1], assInt.registers)
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Sub(args []string) {
	// sub x, y - subtract y (either an integer or the value of a register) from the register x and stores the result in x (i.e. register[x] -= y).
	assInt.registers[args[0]] -= ParseArg(args[1], assInt.registers)
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Mul(args []string) {
	// mul x, y - same with multiply (i.e. register[x] *= y).
	assInt.registers[args[0]] *= ParseArg(args[1], assInt.registers)
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Div(args []string) {
	// div x, y - same with integer division (i.e. register[x] /= y).
	if ParseArg(args[1], assInt.registers) == 0 {
		fmt.Println("Division by zero")
		assInt.pc++
		return
	}
	assInt.registers[args[0]] /= ParseArg(args[1], assInt.registers)
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Nop(args []string) {
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Jmp(args []string) {
	lbl_idx := assInt.labels_idx[args[0]]
	if lbl_idx < assInt.program_length {
		assInt.pc = lbl_idx + 1
	} else {
		fmt.Println("Executing jump to: ", args[0])
		assInt.ExecuteSubroutine(lbl_idx)
	}
}

func (assInt *AssemblerInterpreter) Cmp(args []string) {
	x := ParseArg(args[0], assInt.registers)
	y := ParseArg(args[1], assInt.registers)
	assInt.cmp = x - y
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Jne(args []string) {
	if assInt.cmp != 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Je(args []string) {
	fmt.Printf("JUMP EQUAL: %v\n", assInt.cmp)

	if assInt.cmp == 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Jge(args []string) {
	if assInt.cmp >= 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Jg(args []string) {
	if assInt.cmp > 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Jle(args []string) {
	if assInt.cmp <= 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Jl(args []string) {
	if assInt.cmp < 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Call(args []string) {
	lbl_idx := assInt.labels_idx[args[0]]
	fmt.Println("Executing call to: ", args[0])
	assInt.ExecuteSubroutine(lbl_idx)
}

func (assInt *AssemblerInterpreter) Msg(args []string) {
	var sb strings.Builder
	for _, arg := range args {
		if arg[0] == '\'' {
			sb.WriteString(arg[1 : len(arg)-1])
		} else {
			sb.WriteString(fmt.Sprintf("%v", assInt.registers[arg]))
		}
	}
	assInt.output = sb.String()
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Ret(args []string) {

}

func (assInt *AssemblerInterpreter) End(args []string) {
	assInt.pc++
}

// var INSTRUCTIONS = map[string]func(ai *AssemblerInterpreter, args []string){
// 	"mov":  (*AssemblerInterpreter).Mov,
// 	"inc":  (*AssemblerInterpreter).Inc,
// 	"dec":  (*AssemblerInterpreter).Dec,
// 	"add":  (*AssemblerInterpreter).Add,
// 	"sub":  (*AssemblerInterpreter).Sub,
// 	"mul":  (*AssemblerInterpreter).Mul,
// 	"div":  (*AssemblerInterpreter).Div,
// 	"nop":  (*AssemblerInterpreter).Nop,
// 	"jmp":  (*AssemblerInterpreter).Jmp,
// 	"cmp":  (*AssemblerInterpreter).Cmp,
// 	"jne":  (*AssemblerInterpreter).Jne,
// 	"je":   (*AssemblerInterpreter).Je,
// 	"jge":  (*AssemblerInterpreter).Jge,
// 	"jg":   (*AssemblerInterpreter).Jg,
// 	"jle":  (*AssemblerInterpreter).Jle,
// 	"jl":   (*AssemblerInterpreter).Jl,
// 	"call": (*AssemblerInterpreter).Call,
// 	"msg":  (*AssemblerInterpreter).Msg,
// 	"ret":  (*AssemblerInterpreter).Ret,
// 	"end":  (*AssemblerInterpreter).End,
// }

//TODO zrobienie czegos na wzór slownika z subroutines tylko w obrebie programu:
//cos na wzor funckji ParseCustomSubroutines

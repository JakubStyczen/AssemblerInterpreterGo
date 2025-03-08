package assemblerinterpretergo

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
	// nop - do noting
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Jmp(args []string) {
	// jmp lbl - jumps to the label lbl.
	lbl_idx := assInt.labels_idx[args[0]]
	if lbl_idx < assInt.program_length {
		assInt.pc = lbl_idx + 1
	} else {
		fmt.Println("Executing jump to: ", args[0])
		assInt.pc = lbl_idx + 1
	}
}

func (assInt *AssemblerInterpreter) Cmp(args []string) {
	// cmp x, y -compares x (either an integer or the value of a register) and y (either an integer or the value of a register). The result is used in the conditional jumps (jne, je, jge, jg, jle and jl)
	x := ParseArg(args[0], assInt.registers)
	y := ParseArg(args[1], assInt.registers)
	assInt.cmp = x - y
	assInt.pc++
}

func (assInt *AssemblerInterpreter) Jne(args []string) {
	// jne lbl - jump to the label lbl if the values of the previous cmp command were not equal.
	if assInt.cmp != 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Je(args []string) {
	// je lbl - jump to the label lbl if the values of the previous cmp command were equal.
	if assInt.cmp == 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Jge(args []string) {
	// jge lbl - jump to the label lbl if x was greater or equal than y in the previous cmp command.
	if assInt.cmp >= 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Jg(args []string) {
	// jg lbl - jump to the label lbl if x was greater than y in the previous cmp command.
	if assInt.cmp > 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Jle(args []string) {
	// jle lbl - jump to the label lbl if x was less or equal than y in the previous cmp command.
	if assInt.cmp <= 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Jl(args []string) {
	// jl lbl - jump to the label lbl if x was less than y in the previous cmp command.
	if assInt.cmp < 0 {
		assInt.Jmp([]string{args[0]})
	} else {
		assInt.pc++
	}
}

func (assInt *AssemblerInterpreter) Call(args []string) {
	// call lbl - call to the subroutine identified by lbl. When a ret is found in a subroutine, the instruction pointer should return to the instruction next to this call command.
	if assInt.pre_call_pc == -1 {
		assInt.pre_call_pc = assInt.pc
	}
	lbl_idx := assInt.labels_idx[args[0]]
	fmt.Println("Executing call to: ", args[0])
	assInt.pc = lbl_idx + 1
}

func (assInt *AssemblerInterpreter) Msg(args []string) {
	// msg 'Register: ', x - this instruction stores the output of the program. It may contain text strings (delimited by single quotes) and registers. The number of arguments isn't limited and will vary, depending on the program.
	var sb strings.Builder
	for i, arg := range args {
		if arg == "'" && i < len(args)-1 && args[i+1] == "'" {
			sb.WriteString(", ")
			continue

		} else if arg == "'" {
			continue
		}
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
	// ret - when a ret is found in a subroutine, the instruction pointer should return to the instruction that called the current function.
	assInt.pc = assInt.pre_call_pc + 1
	assInt.pre_call_pc = -1
}

func (assInt *AssemblerInterpreter) End(args []string) {
	// end - this instruction indicates that the program ends correctly, so the stored output is returned (if the program terminates without this instruction it should return the default output: see below).
	assInt.pc++
}

var INSTRUCTIONS = map[string]func(ai *AssemblerInterpreter, args []string){
	"mov":  (*AssemblerInterpreter).Mov,
	"inc":  (*AssemblerInterpreter).Inc,
	"dec":  (*AssemblerInterpreter).Dec,
	"add":  (*AssemblerInterpreter).Add,
	"sub":  (*AssemblerInterpreter).Sub,
	"mul":  (*AssemblerInterpreter).Mul,
	"div":  (*AssemblerInterpreter).Div,
	"nop":  (*AssemblerInterpreter).Nop,
	"jmp":  (*AssemblerInterpreter).Jmp,
	"cmp":  (*AssemblerInterpreter).Cmp,
	"jne":  (*AssemblerInterpreter).Jne,
	"je":   (*AssemblerInterpreter).Je,
	"jge":  (*AssemblerInterpreter).Jge,
	"jg":   (*AssemblerInterpreter).Jg,
	"jle":  (*AssemblerInterpreter).Jle,
	"jl":   (*AssemblerInterpreter).Jl,
	"call": (*AssemblerInterpreter).Call,
	"msg":  (*AssemblerInterpreter).Msg,
	"ret":  (*AssemblerInterpreter).Ret,
	"end":  (*AssemblerInterpreter).End,
}

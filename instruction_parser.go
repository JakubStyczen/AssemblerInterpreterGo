package assemblerinterpretergo

import (
	"strconv"
	"strings"
)

// Parses the argument of an instruction
// If the argument is a register, it returns the value of the register
// If the argument is a number, it returns the number
// If the argument is not a number or a register, it returns 0
func ParseArg(arg string, reg map[string]int) int {
	if val, ok := reg[arg]; ok {
		return val
	}
	if val, err := strconv.Atoi(arg); err == nil {
		return val
	}
	return 0
}

func filterEmptyLinesAndComments(program_lines []string) []string {
	filter_comment_lines := []string{}
	for _, line := range program_lines {
		if strings.Contains(line, ";") {
			line = strings.Split(line, ";")[0]
		}
		if line == "" {
			continue
		}
		filter_comment_lines = append(filter_comment_lines, line)
	}
	return filter_comment_lines
}

// Parses the labels of the program
// Returns a map with the label as key and the index of the instruction as value
func ParseLabels(program_instruction_list []string) map[string]int {
	subroutines := make(map[string]int)
	for i, line := range program_instruction_list {
		if strings.Contains(line, ":") {
			subroutines[strings.Split(line, ":")[0]] = i
		}
	}
	return subroutines
}

func ParseProgram(program string) (instruction_list []string, program_length int) {
	raw_lines := strings.Split(program, "\n")
	filter_comment_lines := filterEmptyLinesAndComments(raw_lines)
	for i := range filter_comment_lines {
		filter_comment_lines[i] = strings.Trim(filter_comment_lines[i], " \t")
	}
	return filter_comment_lines, len(filter_comment_lines)
}

func trim_array(arr []string, delimiter string) []string {
	var trimed_array = []string{}
	for _, elem := range arr {
		if elem == "" {
			continue
		}
		trimed_array = append(trimed_array, strings.Trim(elem, delimiter))
	}
	return trimed_array
}

// Prepares the instruction to be executed
// Returns the instruction and its arguments
// If the instruction is msg, it returns the instruction and the arguments
// separated by commas
// If the instruction is a label, it returns the instruction and an empty array
// If the instruction is a 2 arguments instruction, it returns the instruction and the arguments
// If the instruction is a 3 arguments instruction, it returns the instruction and the arguments
func PrepareInstruction(instruction string) (string, []string) {
	instruction_parts := strings.Split(instruction, " ")
	if instruction_parts[0] == "msg" {
		msg_args := strings.Split(instruction[3:], ",")
		msg_args = trim_array(msg_args, " ")
		return instruction_parts[0], msg_args
	}
	instruction_parts = trim_array(instruction_parts, " ")
	//only label to subroutines or end, ret, nop and label:
	if len(instruction_parts) == 1 {
		return instruction_parts[0], []string{}
	}
	//Trim coma in case of 2 arguments instructions
	if len(instruction_parts) == 3 {
		instruction_parts[1] = strings.Trim(instruction_parts[1], ",")
	}
	args := instruction_parts[1:]
	return instruction_parts[0], args
}

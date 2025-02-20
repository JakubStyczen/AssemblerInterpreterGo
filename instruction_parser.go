package main

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

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

func ParseCustomSubroutines(program_function_defs []string) map[string]int {
	subroutines := make(map[string]int)
	for i, line := range program_function_defs {
		if strings.Contains(line, ":") {
			subroutines[strings.Split(line, ":")[0]] = i
		}
	}
	return subroutines
}

func ParseProgram(program string) (instruction_list []string, end_inx int) {
	raw_lines := strings.Split(program, "\n")
	filter_comment_lines := filterEmptyLinesAndComments(raw_lines)
	for i := range filter_comment_lines {
		filter_comment_lines[i] = strings.Trim(filter_comment_lines[i], " \t")
	}
	end_inx = slices.IndexFunc(filter_comment_lines, func(line string) bool { return strings.Contains(line, "end") })
	if end_inx == -1 {
		end_inx = len(filter_comment_lines)
		return filter_comment_lines, end_inx
	}
	return filter_comment_lines, end_inx + 1
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

func PrepareInstruction(instruction string) (string, []string) {
	instruction_parts := strings.Split(instruction, " ")
	if instruction_parts[0] == "msg" {
		msg_args := strings.Split(instruction[3:], ",")
		msg_args = trim_array(msg_args, " ")
		return instruction_parts[0], msg_args
	}
	instruction_parts = trim_array(instruction_parts, " ")
	//only label to subroutines or end and ret
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

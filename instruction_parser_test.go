package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArg(t *testing.T) {
	empty_map := make(map[string]int)
	tests := []struct {
		input_string string
		registers    map[string]int
		expected     int
	}{
		{"", empty_map, 0},
		{"5", empty_map, 5},
		{"-5", empty_map, -5},
		{"1.1", empty_map, 0},
		{"a", map[string]int{"a": 5}, 5},
		{"x", map[string]int{"a": 5}, 0},
	}

	for _, tt := range tests {
		t.Run("Arg Parsing", func(t *testing.T) {
			result := ParseArg(tt.input_string, tt.registers)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFilterEmptyLinesAndComments(t *testing.T) {
	tests := []struct {
		program  []string
		expected []string
	}{
		{[]string{""}, []string{}},
		{[]string{" "}, []string{" "}},
		{[]string{"Kox"}, []string{"Kox"}},
		{[]string{"; Test Comment"}, []string{}},
		{[]string{"mul a 6; Test Comment"}, []string{"mul a 6"}},
	}

	for _, tt := range tests {
		t.Run("Line filtering", func(t *testing.T) {
			result := filterEmptyLinesAndComments(tt.program)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseProgram(t *testing.T) {
	tests := []struct {
		program            string
		expected_inst_list []string
		expected_prog_len  int
	}{
		{
			`; My first program
		mov  a, 5
		inc  a
		call function
		msg  '(5+1)/2 = ', a    ; output message
		end`,
			[]string{
				"mov  a, 5",
				"inc  a",
				"call function",
				"msg  '(5+1)/2 = ', a",
				"end",
			},
			5,
		},
		{
			`mov   a, 5
		mov   b, a
		mov   c, a
		call  proc_fact
		call  print
		end

		proc_fact:
			dec   b
			mul   c, b
			cmp   b, 1
			jne   proc_fact
			ret`,
			[]string{
				"mov   a, 5",
				"mov   b, a",
				"mov   c, a",
				"call  proc_fact",
				"call  print",
				"end",
				"proc_fact:",
				"dec   b",
				"mul   c, b",
				"cmp   b, 1",
				"jne   proc_fact",
				"ret",
			},
			6,
		},
		{
			"mov a, 5\nend", []string{"mov a, 5", "end"}, 2,
		},
	}

	for _, tt := range tests {
		t.Run("Program parsing", func(t *testing.T) {
			inst, prog_len := ParseProgram(tt.program)
			assert.Equal(t, tt.expected_inst_list, inst)
			assert.Equal(t, tt.expected_prog_len, prog_len)
		})
	}
}

func TestParseCustomSubroutines(t *testing.T) {
	tests := []struct {
		subr_defs []string
		expected  map[string]int
	}{
		{[]string{""}, map[string]int{}},
		{[]string{"mov   a, 5", "mov   b, a"}, map[string]int{}},
		{[]string{"proc_fact:", "mov   a, 5"}, map[string]int{"proc_fact": 0}},
		{[]string{"mov b, 3", "proc_fact:", "mov   a, 5", "proc_fact_2:", "mov   a, 5"}, map[string]int{"proc_fact": 1, "proc_fact_2": 3}},
	}

	for _, tt := range tests {
		t.Run("Subroutines parsing", func(t *testing.T) {
			result := ParseCustomSubroutines(tt.subr_defs)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPrepareInstruction(t *testing.T) {
	tests := []struct {
		instruction    string
		expected_instr string
		expected_args  []string
	}{
		{"mov   a, 5", "mov", []string{"a", "5"}},
		{"mov a, 5", "mov", []string{"a", "5"}},
		{"proc_fact:", "proc_fact:", []string{}},
		{"call function", "call", []string{"function"}},
		{"msg  '(5+1)/2 = ', a", "msg", []string{"'(5+1)/2 = '", "a"}},
		{"msg a, '^', b, ' = ', c", "msg", []string{"a", "'^'", "b", "' = '", "c"}},
		{"end", "end", []string{}},
	}
	for _, tt := range tests {
		t.Run("Preparing instructions", func(t *testing.T) {
			instr, args := PrepareInstruction(tt.instruction)
			assert.Equal(t, tt.expected_instr, instr)
			assert.Equal(t, tt.expected_args, args)
		})
	}
}

func TestTrimArray(t *testing.T) {
	tests := []struct {
		arr          []string
		delimiter    string
		expected_arr []string
	}{
		{[]string{"a ", "5  "}, " ", []string{"a", "5"}},
		{[]string{"'(5+1)/2 = '"}, "'", []string{"(5+1)/2 = "}},
	}
	for _, tt := range tests {
		t.Run("Trim array", func(t *testing.T) {
			arr := trim_array(tt.arr, tt.delimiter)
			assert.Equal(t, tt.expected_arr, arr)
		})
	}
}

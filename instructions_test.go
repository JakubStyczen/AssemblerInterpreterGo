package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMov(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 5\nend", map[string]int{"a": 5}},
		{"mov a,  5\nend", map[string]int{"a": 5}},
		{"mov a, 5\nmov b, a", map[string]int{"a": 5, "b": 5}},
	}
	for _, tt := range tests {
		t.Run("Mov instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestInc(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 5\ninc a\nend", map[string]int{"a": 6}},
		{"mov a, 5\ninc a\ninc a\nend", map[string]int{"a": 7}},
	}
	for _, tt := range tests {
		t.Run("Inc instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestDec(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 5\ndec a\nend", map[string]int{"a": 4}},
		{"mov a, 5\ndec a\ndec a\nend", map[string]int{"a": 3}},
	}
	for _, tt := range tests {
		t.Run("Dec instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestNop(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"nop\nnop\nmov a, 5\ndec a\nend", map[string]int{"a": 4}},
	}
	for _, tt := range tests {
		t.Run("Nop instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestJmp(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 5\njmp kox\nmov b, 3\nnop\nkox:\nmov c, 3\nend", map[string]int{"a": 5, "c": 3}},
	}
	for _, tt := range tests {
		t.Run("Jmp instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestJne(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 5\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njne kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 5, "b": 3}},
		{"mov a, 4\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njne kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 5, "b": 3, "c": 3}},
	}
	for _, tt := range tests {
		t.Run("Jne instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestJe(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 5\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\nje kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 6, "b": 3, "c": 3}},
		{"mov a, 4\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\nje kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 4, "b": 3}},
	}
	for _, tt := range tests {
		t.Run("Je instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestJg(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 6\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njg kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 7, "b": 3, "c": 3}},
		{"mov a, 4\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njg kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 4, "b": 3}},
	}
	for _, tt := range tests {
		t.Run("Jg instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestJge(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 6\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njge kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 7, "b": 3, "c": 3}},
		{"mov a, 4\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njge kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 4, "b": 3}},
		{"mov a, 5\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njge kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 6, "b": 3, "c": 3}},
	}
	for _, tt := range tests {
		t.Run("Jge instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestJl(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 6\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njl kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 6, "b": 3}},
		{"mov a, 4\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njl kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 5, "b": 3, "c": 3}},
	}
	for _, tt := range tests {
		t.Run("Jl instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestJle(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 6\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njle kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 6, "b": 3}},
		{"mov a, 4\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njle kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 5, "b": 3, "c": 3}},
		{"mov a, 5\njmp compare\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njle kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 6, "b": 3, "c": 3}},
	}
	for _, tt := range tests {
		t.Run("Jle instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestCmp(t *testing.T) {
	tests := []struct {
		program            string
		expected_cmp       int
		expected_registers map[string]int
	}{
		{"cmp 1, 2\nend", -1, map[string]int{}},
		{"cmp 1, 1\nend", 0, map[string]int{}},
		{"cmp 1, 0\nend", 1, map[string]int{}},
		{"mov a, 1\nmov b, 2\ncmp a, b\nend", -1, map[string]int{"a": 1, "b": 2}},
		{"mov a, 1\nmov b, 1\ncmp a, b\nend", 0, map[string]int{"a": 1, "b": 1}},
		{"mov a, 1\nmov b, 0\ncmp a, b\nend", 1, map[string]int{"a": 1, "b": 0}},
		{"mov b, 2\ncmp 1, b\nend", -1, map[string]int{"b": 2}},
		{"mov b, 1\ncmp 1, b\nend", 0, map[string]int{"b": 1}},
		{"mov b, 0\ncmp 1, b\nend", 1, map[string]int{"b": 0}},
		{"mov b, 2\ncmp b, 1\nend", 1, map[string]int{"b": 2}},
		{"mov b, 1\ncmp b, 1\nend", 0, map[string]int{"b": 1}},
		{"mov b, 0\ncmp b, 1\nend", -1, map[string]int{"b": 0}},
	}
	for _, tt := range tests {
		t.Run("Cmp instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_cmp, assInt.cmp)
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestCall(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"mov a, 6\ncall kox\nmov b, 3\nnop\nend\ncompare:\ncmp a, 5\njle kox\n ret\nkox:\nmov c, 3\ninc a\nret", map[string]int{"a": 7, "b": 3, "c": 3}},
	}
	for _, tt := range tests {
		t.Run("Call instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

func TestMsg(t *testing.T) {
	tests := []struct {
		program         string
		expected_result string
	}{
		{"mov a, 3\nmsg  '(5+1)/2 = ', a\nend", "(5+1)/2 = 3"},
	}
	for _, tt := range tests {
		t.Run("Msg instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			result := assInt.Run()
			assert.Equal(t, tt.expected_result, result)
		})
	}
}

func TestEnd(t *testing.T) {
	tests := []struct {
		program            string
		expected_registers map[string]int
	}{
		{"end", map[string]int{}},
	}
	for _, tt := range tests {
		t.Run("End instruction", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			assInt.Run()
			assert.Equal(t, tt.expected_registers, assInt.registers)
		})
	}
}

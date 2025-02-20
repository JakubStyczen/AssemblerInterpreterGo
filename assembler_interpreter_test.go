package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyProgram(t *testing.T) {
	tests := []struct {
		program         string
		expected_output string
	}{
		{`
; My first program
mov  a, 5
inc  a
call function
msg  '(5+1)/2 = ', a    ; output message
end

function:
    div  a, 2
    ret`, "(5+1)/2 = 3"},
	}
	for _, tt := range tests {
		t.Run("Running Any Program", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			output := assInt.Run()
			assert.Equal(t, tt.expected_output, output)
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		program         string
		expected_output string
	}{
		{`
mov   a, 5
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
    ret

print:
    msg   a, '! = ', c ; output text
    ret`, "5! = 120"},
	}
	for _, tt := range tests {
		t.Run("Running Factorial", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			output := assInt.Run()
			assert.Equal(t, tt.expected_output, output)
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		program         string
		expected_output string
	}{
		{`
mov   a, 8            ; value
mov   b, 0            ; next
mov   c, 0            ; counter
mov   d, 0            ; first
mov   e, 1            ; second
call  proc_fib
call  print
end

proc_fib:
    cmp   c, 2
    jl    func_0
    mov   b, d
    add   b, e
    mov   d, e
    mov   e, b
    inc   c
    cmp   c, a
    jle   proc_fib
    ret

func_0:
    mov   b, c
    inc   c
    jmp   proc_fib

print:
    msg   'Term ', a, ' of Fibonacci series is: ', b        ; output text
    ret`, "Term 8 of Fibonacci series is: 21"},
	}
	for _, tt := range tests {
		t.Run("Running Fibonacci", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			output := assInt.Run()
			assert.Equal(t, tt.expected_output, output)
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		program         string
		expected_output string
	}{
		{`
mov   a, 2            ; value1
mov   b, 10           ; value2
mov   c, a            ; temp1
mov   d, b            ; temp2
call  proc_func
call  print
end

proc_func:
    cmp   d, 1
    je    continue
    mul   c, a
    dec   d
    call  proc_func

continue:
    ret

print:
    msg a, '^', b, ' = ', c
    ret`, "2^10 = 1024"},
	}
	for _, tt := range tests {
		t.Run("Running Power", func(t *testing.T) {
			assInt := NewAssemblerInterpreter(tt.program)
			output := assInt.Run()
			assert.Equal(t, tt.expected_output, output)
		})
	}
}

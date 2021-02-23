package main

import (
	// "fmt"
	"github.com/yuin/gopher-lua"
	// "log"
	// "time"
)

func square(L *lua.LState) int { //*
	i := L.ToInt(1)          // get first (1) function argument and convert to int
	ln := lua.LNumber(i * i) // make calculation and cast to LNumber
	L.Push(ln)               // Push it to the stack
	return 1                 // Notify that we pushed one value to the stack
}

func hunker(L *lua.LState) int {
	i := L.ToString(1)
	ln := lua.LString(i + " Marshall")
	L.Push(ln)
	return 1
}

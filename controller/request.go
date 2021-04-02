package controller

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

func GetRequestHeaders(L *lua.LState) int {
	fmt.Println("Jason is a TURD!")
	return 1
}

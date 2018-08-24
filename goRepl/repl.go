package main

// #cgo pkg-config: lua
/*
#cgo pkg-config: lua
#include "lua.h"
#include "lauxlib.h"
#include "lualib.h"

int my_luaL_loadbuffer(lua_State* L, const char* s, size_t sz, const char* m) {
	return luaL_loadbuffer(L, s, sz, m);
}

int my_lua_pcall(lua_State *L, int nargs, int nresults, int errfunc) {
	return lua_pcall(L, nargs, nresults, errfunc);
}

const char* my_lua_tostring(lua_State *L, int idx) {
	return lua_tostring(L, idx);
}
*/
import "C"
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var L *C.lua_State = C.luaL_newstate()

	C.luaL_openlibs(L)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")

	for scanner.Scan() {
		line := scanner.Text()
		err := C.my_luaL_loadbuffer(L, C._GoStringPtr(line), C._GoStringLen(line), C._GoStringPtr("line"))

		if err != 0 {
			fmt.Fprintf(os.Stderr, "%s\n", C.my_lua_tostring(L, -1))
		} else {
			err = C.my_lua_pcall(L, 0, 0, 0)

			if err != 0 {
				fmt.Fprintf(os.Stderr, "%s\n", C.my_lua_tostring(L, -1))
			}
		}

		fmt.Print("> ")
	}

	C.lua_close(L)
}

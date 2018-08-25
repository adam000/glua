package lua

/*
#cgo pkg-config: lua
#include "lua.h"
#include "lauxlib.h"
#include "lualib.h"

const char* my_lua_tostring(lua_State *L, int idx) {
	return lua_tostring(L, idx);
}

int my_lua_pcall(lua_State *L, int nargs, int nresults, int errfunc) {
	return lua_pcall(L, nargs, nresults, errfunc);
}

*/
import "C"
import "errors"

type State struct {
	state *C.lua_State
}

// NewState creates a new state and calls openlibs on that state.
func NewState() State {
	cState := C.luaL_newstate()
	st := State{
		state: cState,
	}

	C.luaL_openlibs(cState)

	return st
}

func (s *State) Close() {
	C.lua_close(s.state)
	s.state = nil
}

// TODO probably delete this because it exposes C. But useful
// for interim.
func (s *State) GetState() *C.lua_State {
	return s.state
}

func (s *State) ToStringTop() string {
	return C.GoString(C.my_lua_tostring(s.state, -1))
}

func (s *State) PCall() error {
	errNo := C.my_lua_pcall(s.state, 0, 0, 0)

	if errNo != 0 {
		return errors.New(s.ToStringTop())
	}

	return nil
}

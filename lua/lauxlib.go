package lua

/*
#cgo pkg-config: lua
#include "lauxlib.h"

int my_luaL_loadbuffer(lua_State* L, const char* s, size_t sz, const char* m) {
	return luaL_loadbuffer(L, s, sz, m);
}
*/
import "C"
import "errors"

// luaL_loadbuffer(self, line, len(line), name)
func (s *State) LoadBuffer(line, name string) error {
	errNo := C.my_luaL_loadbuffer(s.state, C._GoStringPtr(line), C._GoStringLen(line), C._GoStringPtr(name))

	if errNo != 0 {
		return errors.New(s.ToStringTop())
	}

	return nil
}

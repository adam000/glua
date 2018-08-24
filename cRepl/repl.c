#include "repl.h"

void process_line(char* buff, lua_State* L) {
    luaL_openlibs(L);

    int error = luaL_loadbuffer(L, buff, strlen(buff), "line") || lua_pcall(L, 0, 0, 0);
    if (error) {
        fprintf(stderr, "%s\n", lua_tostring(L, -1));
        lua_pop(L, 1);
    }
}

void read_forever() {
    char buff[256];
    lua_State *L = luaL_newstate();

    setbuf(stdout, NULL);

    printf("> ");
    while (fgets(buff, sizeof(buff), stdin) != NULL) {
        process_line(&buff[0], L);
        printf("> ");
    }

    lua_close(L);
}

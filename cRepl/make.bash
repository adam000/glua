#!/bin/bash
gcc luaRepl.c repl.c $(pkg-config --cflags lua) -llua -o repl

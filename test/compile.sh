#!/bin/bash

emcc -O3 --llvm-lto 3 \
-Wall \
-s STANDALONE_WASM=1 \
-s NO_EXIT_RUNTIME=1 \
-s NO_FILESYSTEM=1 \
-s STRICT=1 \
-s EXPORTED_FUNCTIONS='["_malloc", "_free"]' \
-s EXTRA_EXPORTED_RUNTIME_METHODS='["stackAlloc", "stackRestore", "stackSave"]' \
-o test.wasm \
_test.c

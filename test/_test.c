#include <stdint.h>
#include <stdlib.h>

#include <emscripten.h>

EMSCRIPTEN_KEEPALIVE
int32_t* int32_new() {
  return malloc(sizeof(int32_t));
}

EMSCRIPTEN_KEEPALIVE
int32_t int32_get(int32_t* ptr) {
  return *ptr;
}

EMSCRIPTEN_KEEPALIVE
void int32_set(int32_t* ptr, int32_t v) {
  *ptr = v;
}

EMSCRIPTEN_KEEPALIVE void int32_delete(int32_t* ptr) {
  free(ptr);
}

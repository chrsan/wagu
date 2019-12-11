# wagu

A WIP that converts WebAssembly to pure [Go][1] code (more docs to come).

Has been used to successfully convert [SQLite][2], [HarfBuzz][3], [ICU][4]
and others.

When using [emscripten][5], version >= `1.39.0` with the `STANDALONE_WASM`
setting enabled is preferred.

The conversion is done in two steps.

1. The WebAssembly is converted to an internal IR thats based on ([protobuf][6]).
2. The IR is converted to [Go][1] source code.

## Usage:

```console
git clone https://github.com/chrsan/wagu
cd wagu
go run cmd/*.go --help
```

[1]: https://golang.org
[2]: https://sqlite.org/index.html
[3]: http://harfbuzz.org/
[4]: http://site.icu-project.org
[5]: https://emscripten.org
[6]: https://developers.google.com/protocol-buffers

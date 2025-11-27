Wasm component model tests.

Golang service

C service.


# Compile and run locally
gcc service.c -o cservice
./cservice

go build -o goservice service.go
./goservice


Targetting wasip2

tinygo build -target wasip2

clang --target=wasm32-wasip2



----------------------------------------------------------------
# Compile Go to WASI Preview 1
GOOS=wasip1 GOARCH=wasm go build -o proxy_core.wasm proxy.go

# Convert to WASIP2 Component
wasm-tools component new proxy_core.wasm \
    -o proxy.wasm \
    --adapt wasi_snapshot_preview1=wasi_snapshot_preview1.command.wasm

    # Compile C to WASI Preview 1 (Adjust sysroot path to your wasi-sdk location)
# Example using Zig as a convenient C compiler for WASM:
zig cc -target wasm32-wasi -o server_core.wasm server.c

# OR using standard Clang with wasi-sdk installed:
# clang --target=wasm32-wasi --sysroot=/opt/wasi-sdk/share/wasi-sysroot -o server_core.wasm server.c

# Convert to WASIP2 Component
wasm-tools component new server_core.wasm \
    -o server.wasm \
    --adapt wasi_snapshot_preview1=wasi_snapshot_preview1.command.wasm

# -S common: Enables WASI system usage
# --net: Passes host network capabilities to the WASM
wasmtime run -S common --net server.wasm

# We set the PORT env var just to be safe, though code defaults to 8080
wasmtime run -S common --net --env PORT=8080 proxy.wasm

curl http://localhost:8080/fetch


# Then compose them and do it again - measure time difference :)
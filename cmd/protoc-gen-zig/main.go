package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-zig/v2/internal/runner"
	"github.com/wasilibs/go-protoc-gen-zig/v2/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-zig", os.Args[1:], wasm.ProtocGenZig, os.Stdin, os.Stdout, os.Stderr))
}

package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-zig/v5/internal/runner"
	"github.com/wasilibs/go-protoc-gen-zig/v5/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-zig", os.Args[1:], wasm.ProtocGenZig, os.Stdin, os.Stdout, os.Stderr))
}

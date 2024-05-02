package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-zig",
		LibraryRepo: "arwalk/zig-protobuf",
		GoReleaser:  true,
	})
	boot.Main()
}

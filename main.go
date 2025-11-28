package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println(Version())
}

func Version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		// Goモジュールが無効など
		return "(devel)"
	}
	return info.Main.Version
}

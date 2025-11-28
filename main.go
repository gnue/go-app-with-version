package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	version, main := Version()
	fmt.Println(version, main)
}

func Version() (string, interface{}) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		// Goモジュールが無効など
		return "(devel)", nil
	}
	return info.Main.Version, info.Main
}

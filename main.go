package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	ver, rev := Version()

	if rev == "" {
		fmt.Println(ver)
	} else {
		fmt.Sprintf("%s(%s)", ver, rev[:6])
	}
	fmt.Println()
}

func Version() (string, string) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		// Goモジュールが無効など
		return "(devel)", ""
	}
	ver := info.Main.Version

	PrintSetting(info.Settings)
	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			return ver, setting.Value
		}
	}

	return ver, ""
}

func PrintSetting(settings []debug.BuildSetting) {
	for _, setting := range settings {
		fmt.Println(setting.Key, " = ", setting.Value)
	}
}

func GetSetting(settings []debug.BuildSetting, key string) (string, bool) {
	for _, setting := range settings {
		if setting.Key == key {
			return setting.Value, true
		}
	}

	return "", false
}

package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println(Version())
}

type BuildSetting []debug.BuildSetting

func Version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		// Goモジュールが無効など
		return "(devel)"
	}
	ver := info.Main.Version
	rev, ok := GetSetting(info.Settings, "vcs.revision")
	if ok {
		return fmt.Sprintf("%s(%s)", ver, rev[:6])
	}
	return ver
}

func GetSetting(settings []debug.BuildSetting, key string) (string, bool) {
	for _, setting := range settings {
		if setting.Key == key {
			return setting.Value, true
		}
	}

	return "", false
}

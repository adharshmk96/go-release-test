package main

import (
	"fmt"
	"runtime/debug"
)

var SemVer = "v0.0.0"

func GetSemverInfo() string {
	if SemVer != "v0.0.0" {
		return SemVer
	}
	version, ok := debug.ReadBuildInfo()
	if ok && version.Main.Version != "(devel)" && version.Main.Version != "" {
		return version.Main.Version
	}
	return SemVer
}

func main() {
	fmt.Println(GetSemverInfo())
}

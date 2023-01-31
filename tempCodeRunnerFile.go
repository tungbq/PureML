package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IncrementVersion(latestVersion string) string {
	version := strings.Split(latestVersion, "v")
	versionNumber, _ := strconv.Atoi(version[1])
	newVersionNumber := versionNumber + 1
	newVersion := fmt.Sprintf("v%d", newVersionNumber)
	return newVersion
}

func main() {
	latestVersion := "v1"
	for i := 0; i < 15; i++ {
		newVersion := IncrementVersion(latestVersion)
		fmt.Println(newVersion)
		latestVersion = newVersion
	}
}
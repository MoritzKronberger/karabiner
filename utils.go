package main

import "fmt"

// https://github.com/mxstbr/karabiner/blob/main/utils.ts

func open (appName string) ToEventDefinition {
 return ToEventDefinition{
	ShellCommand: fmt.Sprintf("open %s", appName),
 }
}

func app (name string) ToEventDefinition {
	return open(fmt.Sprintf("-a '%s.app'", name))
}

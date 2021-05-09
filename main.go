// +build windows

package main

import (
	"kmschr.com/mc2brs/ansi"
	"kmschr.com/mc2brs/cli"
	"kmschr.com/mc2brs/mca"
)

func main() {
	mca.PreconvertColors()
	cli.Init()
	ansi.Help()
	for {
		ansi.BasicPrompt("\nEnter command:")
		cli.ConvertWorld()
	}
}

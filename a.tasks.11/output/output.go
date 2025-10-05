package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintCyan(s string) {
	output(color.FgCyan, s)
}

func PrintRed(s string) {
	output(color.FgRed, s)
}

func PrintYellow(s string) {
	output(color.FgYellow, s)
}

func PrintBlue(s string) {
	output(color.FgBlue, s)
}

func Print(s string) {
	output(color.FgWhite, s)
}

func output(c color.Attribute, s string) {
	output := color.New(c)
	output.Printf("%s", s)
}

func Prompt(prompt string) string {
	PrintCyan(prompt)

	var res string
	fmt.Scanln(&res)
	return res
}

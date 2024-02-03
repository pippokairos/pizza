package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(log string) {
	fmt.Printf("%s%s\n", identLevel(), log)
}

func increaseIdent()  { traceLevel = traceLevel + 1 }
func decreaseIndent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	increaseIdent()
	tracePrint("BEGIN " + msg)

	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decreaseIndent()
}

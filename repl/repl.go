package repl

import (
	"bufio"
	"fmt"
	"github.com/pippokairos/pizza/lexer"
	"github.com/pippokairos/pizza/token"
	"io"
)

const PROMPT = "pizza> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for currentToken := l.NextToken(); currentToken.Type != token.EOF; currentToken = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", currentToken)
		}
	}
}

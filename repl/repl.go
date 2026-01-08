// Package repl implements the read eval print loop
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Adam-445/go-interpreter/lexer"
	"github.com/Adam-445/go-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	// "Why not just us os.Stdin and os.Stdout?": because this makes repl testable, decoupled and reusable. Anything that implements the two interfaces can be used

	// Scanner reads from an io.Reader and splits input into tokens (by default based on new lines)
	scanner := bufio.NewScanner(in)

	for {
		if _, err := fmt.Fprint(out, PROMPT); err != nil {
			return
		}
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// %+v means print struct fields and include field names
			if _, err := fmt.Fprintf(out, "%+v\n", tok); err != nil {
				return
			}
		}
	}
}

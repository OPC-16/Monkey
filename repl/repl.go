package repl

import (
    "bufio"
    "fmt"
    "io"

    "github.com/OPC-16/Monkey/token"
    "github.com/OPC-16/Monkey/lexer"
)

const PROMPT = "[>>] "

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

        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Fprintf(out, "%+v\n", tok)
        }
    }
}

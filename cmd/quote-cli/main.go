package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/andrejacobs/go-test1/jsonquote"
	"github.com/andrejacobs/go-test1/quote"
)

func main() {
	countFlag := flag.Int("count", 1, "The number of quotes to print")
	jsonFlag := flag.Bool("json", false, "Encode the quotes in JSON")
	flag.Parse()

	if err := run(os.Stdout, *countFlag, *jsonFlag); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(out io.Writer, count int, outputJson bool) error {
	ql := &quote.QuoteList{}
	ql.Generate(count)

	if outputJson {
		jsonquote.Encode(out, ql)
	} else {
		for _, quote := range ql.Quotes {
			_, err := fmt.Fprintf(out, "%q - %s\n", quote.Line, quote.Who)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

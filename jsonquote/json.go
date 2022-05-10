package jsonquote

import (
	"encoding/json"
	"io"

	"github.com/andrejacobs/go-test1/quote"
)

func Encode(out io.Writer, ql *quote.QuoteList) {
	encoder := json.NewEncoder(out)
	encoder.Encode(ql)
}

package jsonquote_test

import (
	"bytes"
	"testing"

	"github.com/andrejacobs/go-test1/jsonquote"
	"github.com/andrejacobs/go-test1/quote"
)

func TestJsonEncoding(t *testing.T) {
	ql := &quote.QuoteList{}
	ql.Add("p1", "line1")
	ql.Add("p2", "line2")
	ql.Add("p3", "line3")

	var result bytes.Buffer
	jsonquote.Encode(&result, ql)

	expected := "{\"Quotes\":[{\"Line\":\"line1\",\"Who\":\"p1\"},{\"Line\":\"line2\",\"Who\":\"p2\"},{\"Line\":\"line3\",\"Who\":\"p3\"}]}\n"
	if result.String() != expected {
		t.Fatalf("Expected json:\n%q\n\nInstead we received:\n%q", expected, result.String())
	}
}

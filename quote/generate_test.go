package quote_test

import (
	"testing"

	"github.com/andrejacobs/go-test1/quote"
)

func TestGenerate(t *testing.T) {
	ql := &quote.QuoteList{}

	expectedCount := 5
	ql.Generate(expectedCount)
	if len(ql.Quotes) != expectedCount {
		t.Fatalf("Expected %d quotes to be generated, instead we got %d", expectedCount, len(ql.Quotes))
	}
}

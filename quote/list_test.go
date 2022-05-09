package quote_test

import (
	"testing"

	"github.com/andrejacobs/go-test1/quote"
)

func TestListAdd(t *testing.T) {
	ql := &quote.QuoteList{}

	if len(ql.Quotes) > 0 {
		t.Fatal("QuoteList not initialized correctly")
	}

	ql.Add("Person 1", "The quick brown fox")
	ql.Add("Person 2", "Jumped over the lazy dog")

	if len(ql.Quotes) != 2 {
		t.Fatalf("Expected 2 quotes in the list, instead there were: %d", len(ql.Quotes))
	}
}

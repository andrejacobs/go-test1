package quote

type Quote struct {
	Line string
	Who  string
}

type QuoteList struct {
	Quotes []Quote
}

func (ql *QuoteList) Add(who string, line string) {
	q := Quote{
		Who:  who,
		Line: line,
	}

	ql.Quotes = append(ql.Quotes, q)
}

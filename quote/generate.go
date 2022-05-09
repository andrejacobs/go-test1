package quote

import (
	"github.com/andrejacobs/go-test1/internal"
)

func (ql *QuoteList) Generate(count int) {

	for i := 0; i < count; i++ {
		person := internal.RandomPerson()
		line := internal.RandomLine()
		ql.Add(person, line)
	}
}

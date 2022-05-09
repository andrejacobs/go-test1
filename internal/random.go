package internal

import "math/rand"

func RandomPerson() string {
	var people = []string{
		"Jannie",
		"Sannie",
		"Pieter",
		"Stephan",
		"Kobus",
	}

	person := people[rand.Intn(len(people))]
	return person
}

func RandomLine() string {
	var lines = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Mauris at metus in lectus cursus placerat.",
		"Quisque gravida mauris vel ante luctus, eget congue est ornare.",
		"Sed semper mauris nec commodo pharetra.",
		"Fusce nec elit sagittis, sodales nulla a, vehicula libero.",
		"Mauris dignissim libero nec neque tristique, nec fringilla nunc pretium.",
		"Maecenas fermentum diam eget justo euismod bibendum.",
		"Nam euismod sapien vitae quam dapibus molestie.",
		"Sed ac ligula eget ipsum semper pharetra.",
		"Morbi vestibulum augue quis nulla pharetra, at finibus leo consectetur.",
	}

	line := lines[rand.Intn(len(lines))]
	return line
}

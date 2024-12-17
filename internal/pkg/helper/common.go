package helper

import "github.com/gertd/go-pluralize"

type Pluralize struct {
	pl *pluralize.Client
}

func NewPluralize() *Pluralize {
	p := pluralize.NewClient()
	return &Pluralize{pl: p}
}

// Plural 变更单词为复数，先按自然语言处理，如果处理结果和原单词一样，则按土办法强行处理
func (p *Pluralize) Plural(word string) string {
	if newWord := p.pl.Plural(word); newWord != word {
		return newWord
	}
	return addPluralize(word)
}

// 补充变更复数。当pluralize库变更的复数和单数一样时可用。
func addPluralize(word string) string {
	lastLetter := word[len(word)-1:]
	beforeLastLetter := word[len(word)-2 : len(word)-1]
	switch lastLetter {
	case "y":
		if beforeLastLetter == "a" || beforeLastLetter == "e" || beforeLastLetter == "i" || beforeLastLetter == "o" || beforeLastLetter == "u" {
			return word + "s"
		} else {
			return word[:len(word)-1] + "ies"
		}
	case "x", "s", "z", "o":
		return word + "es"
	case "h":
		if beforeLastLetter == "s" || beforeLastLetter == "c" {
			return word + "es"
		} else {
			return word + "s"
		}
	case "f":
		if beforeLastLetter == "f" {
			return word[:len(word)-2] + "ves"
		} else {
			return word[:len(word)-1] + "ves"
		}
	default:
		return word + "s"
	}
}

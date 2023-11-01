package parser

import (
	"github.com/apieat/otto/file"
	"github.com/apieat/otto/token"
)

type Fixture struct {
	Str string
}

type Fixer interface {
	Fix(string, file.Idx, token.Token) (*Fixture, error)
}

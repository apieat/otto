package otto

import (
	"github.com/apieat/otto/ast"
	"github.com/apieat/otto/file"
)

type compiler struct {
	file    *file.File
	program *ast.Program
}

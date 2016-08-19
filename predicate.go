package predicate

import (
	"github.com/pkg/errors"
	"github.com/sjhitchner/go-pred/ast"
	"github.com/sjhitchner/go-pred/lexer"
	"github.com/sjhitchner/go-pred/parser"
	"io"
)

type Predicate interface {
	Evaluate(ast.Context) (bool, error)
	Graph(io.Writer) error
	Data() interface{}
}

type SimplePredicate struct {
	root ast.Node
	data interface{}
}

func NewSimplePredicate(str string, data interface{}) (Predicate, error) {

	tree, err := NewAST(str)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create simple predicate")
	}

	return &SimplePredicate{
		tree,
		data,
	}, nil
}

func (t SimplePredicate) Evaluate(ctx ast.Context) (bool, error) {
	result, err := t.root.Evaluate(ctx)
	if err != nil {
		return false, err
	}

	abool, ok := result.(bool)
	if !ok {
		return false, errors.Errorf("Invalid result '%v' should be bool", result)
	}

	return abool, nil
}

func (t SimplePredicate) Data() interface{} {
	return t.data
}

func (t SimplePredicate) Graph(w io.Writer) error {
	return ast.Graph(w, t.root)
}

func NewAST(str string) (ast.Node, error) {

	lex := lexer.NewLexer([]byte(str))
	p := parser.NewParser()
	tree, err := p.Parse(lex)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build AST")
	}

	return tree.(ast.Node), nil
}

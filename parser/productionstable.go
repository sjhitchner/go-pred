
// generated by gocc; DO NOT EDIT.

package parser

import ( "github.com/sjhitchner/go-pred/ast" )

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Expression	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Expression "and" Expression	<< ast.NewLogicalAnd(X[0], X[2]) >>`,
		Id: "Expression",
		NTType: 1,
		Index: 1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLogicalAnd(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expression : Expression "or" Expression	<< ast.NewLogicalOr(X[0], X[2]) >>`,
		Id: "Expression",
		NTType: 1,
		Index: 2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLogicalOr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expression : "not" Expression	<< ast.NewNegation(X[1]) >>`,
		Id: "Expression",
		NTType: 1,
		Index: 3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNegation(X[1])
		},
	},
	ProdTabEntry{
		String: `Expression : Term	<<  >>`,
		Id: "Expression",
		NTType: 1,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor ">" Factor	<< ast.NewComparisonGreaterThan(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonGreaterThan(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor ">=" Factor	<< ast.NewComparisonGreaterThanEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonGreaterThanEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "<" Factor	<< ast.NewComparisonLessThan(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonLessThan(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "<=" Factor	<< ast.NewComparisonLessThanEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonLessThanEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "=" Factor	<< ast.NewComparisonEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "==" Factor	<< ast.NewComparisonEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "!=" Factor	<< ast.NewComparisonNotEquals(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonNotEquals(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "is not" Factor	<< ast.NewComparisonIsNot(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonIsNot(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "is" Factor	<< ast.NewComparisonIs(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonIs(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "contains" Factor	<< ast.NewComparisonContains(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComparisonContains(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor "matches" Factor	<< ast.NewMatches(X[0], X[2]) >>`,
		Id: "Term",
		NTType: 2,
		Index: 15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMatches(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id: "Term",
		NTType: 2,
		Index: 16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : int_lit	<< ast.NewLiteralInt(X[0]) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteralInt(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : string_lit	<< ast.NewLiteralString(X[0]) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteralString(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : float_lit	<< ast.NewLiteralFloat(X[0]) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteralFloat(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : variable	<< ast.NewResolver(X[0]) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewResolver(X[0])
		},
	},
	ProdTabEntry{
		String: `Factor : "true"	<< ast.NewLiteralBool(true) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteralBool(true)
		},
	},
	ProdTabEntry{
		String: `Factor : "false"	<< ast.NewLiteralBool(false) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteralBool(false)
		},
	},
	ProdTabEntry{
		String: `Factor : "undefined"	<< nil, nil >>`,
		Id: "Factor",
		NTType: 3,
		Index: 23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Factor : "null"	<< nil, nil >>`,
		Id: "Factor",
		NTType: 3,
		Index: 24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Factor : "empty"	<< nil, nil >>`,
		Id: "Factor",
		NTType: 3,
		Index: 25,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Factor : "(" Expression ")"	<< ast.NewClause(X[1]) >>`,
		Id: "Factor",
		NTType: 3,
		Index: 26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewClause(X[1])
		},
	},
	
}

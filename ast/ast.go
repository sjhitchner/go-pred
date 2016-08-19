package ast

import (
	"github.com/pkg/errors"
	"github.com/sjhitchner/go-pred/token"
	"regexp"
	"strconv"
)

type Attrib interface{}

// Functions used to build the AST
// They are included in the grammar.bnf so the parser
// can build the AST

func NewLogicalAnd(a, b Attrib) (*LogicalNode, error) {
	return &LogicalNode{a.(Node), b.(Node), And}, nil
}

func NewLogicalOr(a, b Attrib) (*LogicalNode, error) {
	return &LogicalNode{a.(Node), b.(Node), Or}, nil
}

func NewNegation(a Attrib) (*NegationNode, error) {
	return &NegationNode{a.(Node)}, nil
}

func NewClause(a Attrib) (*ClauseNode, error) {
	return &ClauseNode{a.(Node)}, nil
}

func NewComparisonGreaterThan(a, b Attrib) (*ComparisonNode, error) {
	return &ComparisonNode{a.(Node), b.(Node), GreaterThan}, nil
}

func NewComparisonGreaterThanEquals(a, b Attrib) (*ComparisonNode, error) {
	return &ComparisonNode{a.(Node), b.(Node), GreaterThanEquals}, nil
}

func NewComparisonLessThan(a, b Attrib) (*ComparisonNode, error) {
	return &ComparisonNode{a.(Node), b.(Node), LessThan}, nil
}

func NewComparisonLessThanEquals(a, b Attrib) (*ComparisonNode, error) {
	return &ComparisonNode{a.(Node), b.(Node), LessThanEquals}, nil
}

func NewComparisonEquals(a, b Attrib) (*ComparisonNode, error) {
	return &ComparisonNode{a.(Node), b.(Node), Equals}, nil
}

func NewComparisonNotEquals(a, b Attrib) (*ComparisonNode, error) {
	return &ComparisonNode{a.(Node), b.(Node), NotEquals}, nil
}

func NewComparisonIsNot(a, b Attrib) (*ComparisonNode, error) {
	return &ComparisonNode{a.(Node), b.(Node), IsNot}, nil
}

func NewComparisonIs(a, b Attrib) (*ComparisonNode, error) {
	return &ComparisonNode{a.(Node), b.(Node), Is}, nil
}

func NewComparisonContains(a, b Attrib) (*ComparisonNode, error) {
	return &ComparisonNode{a.(Node), b.(Node), Contains}, nil
}

func NewMatches(a, b Attrib) (*RegexNode, error) {
	bstring := string(b.(*token.Token).Lit)

	bregex, err := regexp.Compile(bstring)
	if err != nil {
		return nil, err
	}
	return &RegexNode{a.(Node), bregex}, nil
}

func NewLiteralBool(a Attrib) (*LiteralNode, error) {
	abool, ok := a.(bool)
	if !ok {
		return nil, errors.Errorf("%v not a bool", a)
	}
	return &LiteralNode{abool}, nil
}

func NewLiteralInt(a Attrib) (*LiteralNode, error) {
	aint, err := IntValue(a.(*token.Token).Lit)
	if err != nil {
		return nil, err
	}
	return &LiteralNode{aint}, nil
}

func NewLiteralFloat(a Attrib) (*LiteralNode, error) {
	afloat, err := FloatValue(a.(*token.Token).Lit)
	if err != nil {
		return nil, err
	}
	return &LiteralNode{afloat}, nil
}

func NewLiteralString(a Attrib) (*LiteralNode, error) {
	astring := string(a.(*token.Token).Lit)
	return &LiteralNode{astring}, nil
}

func NewResolver(a Attrib) (*ResolverNode, error) {
	key := string(a.(*token.Token).Lit)
	return &ResolverNode{key}, nil
}

func IntValue(lit []byte) (int64, error) {
	return strconv.ParseInt(string(lit), 10, 64)
}

func BoolValue(lit []byte) (bool, error) {
	return strconv.ParseBool(string(lit))
}

func FloatValue(lit []byte) (float64, error) {
	return strconv.ParseFloat(string(lit), 64)
}

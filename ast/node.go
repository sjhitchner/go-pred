package ast

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
	"strings"
)

type Logical int
type Comparison int

type NodeShape string

const (
	And Logical = iota
	Or

	GreaterThan Comparison = iota
	GreaterThanEquals
	LessThan
	LessThanEquals
	Equals
	NotEquals
	Is
	IsNot
	Contains
)

var LogicalError = errors.New("Logical Operation Error")
var ComparisonError = errors.New("Comparison Error")
var IncompatibleTypeError = errors.New("Incompatible Types Error")
var ContextMissingKeyError = errors.New("Key missing from context")

type Context map[string]interface{}

type Node interface {
	Evaluate(Context) (interface{}, error)
	String() string
}

// Logical Node

type LogicalNode struct {
	Left    Node
	Right   Node
	Logical Logical
}

func (t LogicalNode) Evaluate(ctx Context) (interface{}, error) {
	if t.Left == nil {
		return nil, errors.New("Left node is nil")
	}
	if t.Right == nil {
		return nil, errors.New("Right node is nil")
	}

	left, err := t.Left.Evaluate(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Left logical evaluate failed")
	}

	right, err := t.Right.Evaluate(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Right logical evaluate failed")
	}

	switch t.Logical {
	case And:
		return OperationAnd(left, right)
	case Or:
		return OperationOr(left, right)
	}

	return nil, LogicalError
}

func (t LogicalNode) String() string {
	return TwoNodeString(t.Left, t.Right, t.Logical)
}

// Comparison Node

type ComparisonNode struct {
	Left       Node
	Right      Node
	Comparison Comparison
}

func (t ComparisonNode) Evaluate(ctx Context) (interface{}, error) {

	if t.Left == nil {
		return nil, errors.New("Left node is nil")
	}

	if t.Right == nil {
		return nil, errors.New("Right node is nil")
	}

	left, err := t.Left.Evaluate(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Left comparison evaluate failed")
	}

	right, err := t.Right.Evaluate(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Right comparison evaluate failed")
	}

	switch t.Comparison {
	case GreaterThan:
		return OperationGreaterThan(left, right)
	case GreaterThanEquals:
		return OperationGreaterThanEquals(left, right)
	case LessThan:
		return OperationLessThan(left, right)
	case LessThanEquals:
		return OperationLessThanEquals(left, right)
	case Equals:
		return OperationEquals(left, right)
	case NotEquals:
		return OperationNotEquals(left, right)
	case Is:
		return OperationIs(left, right)
	case Contains:
		return OperationContains(left, right)
	}

	return nil, ComparisonError
}

func (t ComparisonNode) String() string {
	return TwoNodeString(t.Left, t.Right, t.Comparison)
}

// Regex Node

type RegexNode struct {
	Node  Node
	regex *regexp.Regexp
}

func (t RegexNode) Evaluate(ctx Context) (interface{}, error) {
	node, err := t.Node.Evaluate(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Regex evaluating node failed")
	}

	str, ok := node.(string)
	if !ok {
		return false, errors.Wrapf(IncompatibleTypeError, "REGEX %v incompatible with %v", str, reflect.TypeOf(str))
	}

	return t.regex.MatchString(node.(string)), nil
}

func (t RegexNode) String() string {
	return TwoNodeString(t.Node, nil, t.regex.String())
}

// ClauseNode

type ClauseNode struct {
	Node Node
}

func (t ClauseNode) Evaluate(ctx Context) (interface{}, error) {
	return t.Node.Evaluate(ctx)
}

func (t ClauseNode) String() string {
	return fmt.Sprintf("(%s)", t.Node)
}

// NegateNode

type NegationNode struct {
	Node Node
}

func (t NegationNode) Evaluate(ctx Context) (interface{}, error) {
	a, err := t.Node.Evaluate(ctx)
	if err != nil {
		return nil, nil
	}
	return OperationNot(a)
}

func (t NegationNode) String() string {
	return fmt.Sprintf("NOT %s", t.Node)
}

// LiteralNode

type LiteralNode struct {
	Value interface{}
}

func (t LiteralNode) Evaluate(ctx Context) (interface{}, error) {
	return t.Value, nil
}

func (t LiteralNode) String() string {
	return fmt.Sprintf("'%v'", t.Value)
}

// ResolverNode

type ResolverNode struct {
	key string
}

func (t ResolverNode) Evaluate(ctx Context) (interface{}, error) {
	// TODO arbitrarily deep context maps
	ai, ok := ctx[t.key]
	if !ok {
		return false, errors.Wrapf(ContextMissingKeyError, "key %s doesn't exist", t.key)
	}

	switch a := ai.(type) {
	case int:
		return int64(a), nil
	case float32:
		return float64(a), nil
	default:
		return a, nil
	}
}

func (t ResolverNode) String() string {
	return fmt.Sprintf("$%s", t.key)
}

// Helper methods
func TwoNodeString(left, right Node, middle interface{}) string {
	var l, r string

	if left != nil {
		l = left.String()
	} else {
		l = "empty"
	}

	if right != nil {
		r = right.String()
	} else {
		r = "empty"
	}

	return fmt.Sprintf("(%s %s %s)", l, middle, r)
}

// Operation

func OperationAnd(ai interface{}, bi interface{}) (interface{}, error) {
	a, oka := ai.(bool)
	b, okb := bi.(bool)

	if oka && okb {
		return a && b, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "AND %v incompatible with %v", a, b)
}

func OperationOr(ai interface{}, bi interface{}) (interface{}, error) {
	a, oka := ai.(bool)
	b, okb := bi.(bool)

	if oka && okb {
		return a || b, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "OR %v incompatible with %v", a, b)
}

func OperationNot(ai interface{}) (interface{}, error) {
	a, oka := ai.(bool)
	if oka {
		return !a, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "NOT %v invalid type", a)
}

func OperationGreaterThan(ai interface{}, bi interface{}) (interface{}, error) {
	switch a := ai.(type) {
	case int64:
		b, ok := bi.(int64)
		if ok {
			return a > b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "> (int) %v incompatible with %v", a, b)

	case float64:
		b, ok := bi.(float64)
		if ok {
			return a > b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "> (float) %v incompatible with %v", a, b)

	case string:
		b, ok := bi.(string)
		if ok {
			return a > b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "> (string) %v incompatible with %v", a, b)

	default:
		return false, errors.Wrapf(IncompatibleTypeError, "> %v invalid type", a)
	}
}

func OperationGreaterThanEquals(ai interface{}, bi interface{}) (interface{}, error) {
	switch a := ai.(type) {
	case int64:
		b, ok := bi.(int64)
		if ok {
			return a >= b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, ">= (int) %v incompatible with %v", a, b)

	case float64:
		b, ok := bi.(float64)
		if ok {
			return a >= b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, ">= (float) %v incompatible with %v", a, b)

	case string:
		b, ok := bi.(string)
		if ok {
			return a >= b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, ">= (string) %v incompatible with %v", a, b)

	default:
		return false, errors.Wrapf(IncompatibleTypeError, "> %v invalid type", a)
	}
}

func OperationLessThan(ai interface{}, bi interface{}) (interface{}, error) {
	switch a := ai.(type) {
	case int64:
		b, ok := bi.(int64)
		if ok {
			return a < b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "< (int) %v incompatible with %v", a, b)

	case float64:
		b, ok := bi.(float64)
		if ok {
			return a < b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "< (float) %v incompatible with %v", a, b)

	case string:
		b, ok := bi.(string)
		if ok {
			return a < b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "< (string) %v incompatible with %v", a, b)

	default:
		return false, errors.Wrapf(IncompatibleTypeError, "< %v invalid type %v", a, reflect.TypeOf(a))
	}
}

func OperationLessThanEquals(ai interface{}, bi interface{}) (interface{}, error) {
	switch a := ai.(type) {
	case int64:
		b, ok := bi.(int64)
		if ok {
			return a <= b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "<= (int) %v incompatible with %v", a, b)

	case float64:
		b, ok := bi.(float64)
		if ok {
			return a <= b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "<= (float) %v incompatible with %v", a, b)

	case string:
		b, ok := bi.(string)
		if ok {
			return a <= b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "<= (string) %v incompatible with %v", a, b)

	default:
		return false, errors.Wrapf(IncompatibleTypeError, "<= %v invalid type", a)
	}
}

func OperationEquals(ai interface{}, bi interface{}) (interface{}, error) {

	switch a := ai.(type) {
	case int64:
		b, ok := bi.(int64)
		if ok {
			return a == b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "== (int) %v incompatible with %v", a, b)

	case float64:
		b, ok := bi.(float64)
		if ok {
			return a == b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "== (float) %v incompatible with %v", a, b)

	case string:
		b, ok := bi.(string)
		if ok {
			return a == b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "== (string) %v incompatible with %v", a, b)

	case bool:
		b, ok := bi.(bool)
		if ok {
			return a == b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "== (bool) %v incompatible with %v", a, b)

	default:
		fmt.Printf("== %v invalid type %v", a, reflect.TypeOf(a))
		return false, errors.Wrapf(IncompatibleTypeError, "== %v invalid type", a)
	}
}

func OperationNotEquals(ai interface{}, bi interface{}) (interface{}, error) {
	switch a := ai.(type) {
	case int64:
		b, ok := bi.(int64)
		if ok {
			return a != b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "!= (int) %v incompatible with %v", ai, bi)

	case float64:
		b, ok := bi.(float64)
		if ok {
			return a != b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "!= (float) %v incompatible with %v", ai, bi)

	case string:
		b, ok := bi.(string)
		if ok {
			return a != b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "!= (string) %v incompatible with %v", ai, bi)

	case bool:
		b, ok := bi.(bool)
		if ok {
			return a != b, nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "!= (bool) %v incompatible with %v", a, b)

	default:
		return false, errors.Wrapf(IncompatibleTypeError, "!= %v invalid type", a)
	}
}

func OperationIs(ai interface{}, bi interface{}) (interface{}, error) {
	return false, nil
}

func OperationContains(ai interface{}, bi interface{}) (interface{}, error) {
	switch a := ai.(type) {
	case string:
		b, ok := bi.(string)
		if ok {
			return strings.Contains(a, b), nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "contains (string) %v incompatible with %v", ai, bi)
	default:
		return false, errors.Wrapf(IncompatibleTypeError, "contains %v invalid type %v", a, reflect.TypeOf(a))
	}
}

func (t Logical) String() string {
	switch t {
	case And:
		return "AND"
	case Or:
		return "OR"
	default:
		return "unknown"
	}
}

func (t Comparison) String() string {
	switch t {
	case GreaterThan:
		return ">"
	case GreaterThanEquals:
		return ">="
	case LessThan:
		return "<"
	case LessThanEquals:
		return "<="
	case Equals:
		return "=="
	case NotEquals:
		return "!="
	case Is:
		return "is"
	case Contains:
		return "contains"
	default:
		return "unknown"
	}
}

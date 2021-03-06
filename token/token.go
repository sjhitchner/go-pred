
// generated by gocc; DO NOT EDIT.

package token

import(
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const(
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap  []string
	idMap map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"and",
		"or",
		"not",
		">",
		">=",
		"<",
		"<=",
		"=",
		"==",
		"!=",
		"is not",
		"is",
		"contains",
		"matches",
		"int_lit",
		"string_lit",
		"float_lit",
		"variable",
		"true",
		"false",
		"undefined",
		"null",
		"empty",
		"(",
		")",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"and": 2,
		"or": 3,
		"not": 4,
		">": 5,
		">=": 6,
		"<": 7,
		"<=": 8,
		"=": 9,
		"==": 10,
		"!=": 11,
		"is not": 12,
		"is": 13,
		"contains": 14,
		"matches": 15,
		"int_lit": 16,
		"string_lit": 17,
		"float_lit": 18,
		"variable": 19,
		"true": 20,
		"false": 21,
		"undefined": 22,
		"null": 23,
		"empty": 24,
		"(": 25,
		")": 26,
	},
}


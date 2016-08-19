package predicate

import (
	"fmt"
	"github.com/sjhitchner/go-pred/ast"
	"testing"
)

func testEval(t *testing.T, exampleStr string, output bool) {

	tree, err := NewAST(exampleStr)
	if err != nil {
		panic(err)
	}

	ctx := make(ast.Context)

	result, err := tree.Evaluate(ctx)
	if err != nil {
		fmt.Println(err)

		t.Fatal(err)
	}

	if result != output {
		t.Fatalf("Should be %v not %v for %v", result, output, exampleStr)
	}
}

func TestAnd(t *testing.T) {
	testEval(t, "true and true", true)
	testEval(t, "true and false", false)
	testEval(t, "false and true", false)
	testEval(t, "false and false", false)
}

func TestOr(t *testing.T) {
	testEval(t, "true or true", true)
	testEval(t, "true or false", true)
	testEval(t, "false or true", true)
	testEval(t, "false or false", false)
}

func TestInt(t *testing.T) {
	testEval(t, "3 = 3", true)
	testEval(t, "1 == 1", true)
	testEval(t, "10 <= 10", true)
	testEval(t, "5 < 6", true)
	testEval(t, "7 >= 7", true)
	testEval(t, "7 > 6", true)
	testEval(t, "5 != 6", true)

	testEval(t, "3 = 2", false)
	testEval(t, "1 == 5", false)
	testEval(t, "12 <= 10", false)
	testEval(t, "8 < 6", false)
	testEval(t, "7 >= 9", false)
	testEval(t, "7 > 9", false)
	testEval(t, "6 != 6", false)
}

func TestClauses(t *testing.T) {
	testEval(t, "(3 == 4 or 2 == 2)", true)
	testEval(t, "(3 == 4 or 2 == 2) and (3 < 6)", true)
	testEval(t, "(3 == 4 or 2 == 2 and 3 < 6)", true)
}

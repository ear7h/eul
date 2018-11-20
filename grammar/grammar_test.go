package grammar

import (
	"fmt"
	"strings"
	"testing"
)

func TestGrammar(t *testing.T) {
	testStr := `1 /- (x ^ 2)`

	tree, err := ParseReader("", strings.NewReader(testStr))
	if err != nil {
		t.Fatalf("%v", err)
	}

	float := 0.0

	TreeForEach(tree.(AstNode), func(node AstNode) error {
		fmt.Printf("Node: %v (%T)\n", node, node)
		if v, ok := node.(*ScalarVar); ok {
			v.SetLoc(&float)
		}
		return nil
	})

	fmt.Println(tree.(Scalar).Val())

	for ; float <= 10; float += 1 {
		fmt.Println(tree.(Scalar).Val())
	}
}

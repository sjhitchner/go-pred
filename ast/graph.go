package ast

import (
	"fmt"
	"github.com/pkg/errors"
	graphviz "github.com/sjhitchner/go-graphviz"
	"io"
	"reflect"
)

const (
	GraphName = "Predicate"
	Root      = 0

	LogicalShape    NodeShape = "ellipse"
	NegateShape     NodeShape = "triangle"
	ComparisonShape NodeShape = "diamond"
	LiteralShape    NodeShape = "box"
	ResolverShape   NodeShape = "trapezium"
)

var count int

// Generate a Graphviz representation of the node tree
func Graph(w io.Writer, node Node) error {

	graph := graphviz.NewGraph(GraphName)

	if err := GraphWalk(graph, Root, node); err != nil {
		return err
	}

	return graph.Output(w)
}

func GraphWalk(graph *graphviz.Graph, parentId int, node Node) error {

	switch n := node.(type) {
	case *LogicalNode:
		id := AddGraphNode(graph, parentId, n.Logical.String(), LogicalShape)

		if n.Left != nil {
			if err := GraphWalk(graph, id, n.Left); err != nil {
				return err
			}
		}
		if n.Right != nil {
			if err := GraphWalk(graph, id, n.Right); err != nil {
				return err
			}
		}

	case *ComparisonNode:
		id := AddGraphNode(graph, parentId, n.Comparison.String(), ComparisonShape)

		if n.Left != nil {
			if err := GraphWalk(graph, id, n.Left); err != nil {
				return err
			}
		}
		if n.Right != nil {
			if err := GraphWalk(graph, id, n.Right); err != nil {
				return err
			}
		}

	case *NegationNode:
		id := AddGraphNode(graph, parentId, "NOT", NegateShape)
		if err := GraphWalk(graph, id, n.Node); err != nil {
			return err
		}

	case *ClauseNode:
		id := AddGraphNode(graph, parentId, "CLAUSE", NegateShape)
		if err := GraphWalk(graph, id, n.Node); err != nil {
			return err
		}

	case *LiteralNode:
		AddGraphNode(graph, parentId, n.String(), LiteralShape)

	case *ResolverNode:
		AddGraphNode(graph, parentId, n.String(), ResolverShape)

	default:
		return errors.Errorf("Unknown node type %v", reflect.TypeOf(node))
	}

	return nil
}

var nodeCount = 0

func AddGraphNode(graph *graphviz.Graph, parentId int, label string, shape NodeShape) int {
	nodeCount++

	parentName := fmt.Sprintf("Node%d", parentId)
	nodeName := fmt.Sprintf("Node%d", nodeCount)

	graph.AddNode(nodeName, map[string]string{
		"label": label,
		"shape": string(shape),
	})

	if parentId != Root {
		graph.AddEdge(parentName, nodeName, true, nil)
	}

	return nodeCount
}

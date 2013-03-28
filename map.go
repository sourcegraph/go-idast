package idast

import (
	"go/ast"
)

func Map(node ast.Node) map[ast.Node]NodeId {
	m := make(map[ast.Node]NodeId, 0)
	Inspect(node, func(node ast.Node, id NodeId) bool {
		if node != nil {
			m[node] = id.dup()
		}
		return true
	})
	return m
}

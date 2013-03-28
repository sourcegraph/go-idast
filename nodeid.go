package idast

import (
	"go/ast"
	"strings"
)

type NodeId []string

type NodeWithId struct {
	Node ast.Node
	Id   NodeId
}

func (nid *NodeId) String() string {
	return strings.Join(*nid, "/")
}

func (nid *NodeId) dup() NodeId {
	newId := make(NodeId, len(*nid))
	copy(newId, *nid)
	return newId
}

// Pushes c onto the end of nid.
func (nid *NodeId) push(c string) {
	*nid = append(*nid, c)
}

// Removes the last item in nid.
func (nid *NodeId) pop() {
	*nid = (*nid)[:len(*nid)-1]
}

// Returns a copy of nid with cs pushed onto the end.
func (nid *NodeId) pushed(cs ...string) NodeId {
	return append(nid.dup(), cs...)
}

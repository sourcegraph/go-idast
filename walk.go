package idast

import (
	"fmt"
	"go/ast"
	"reflect"
	"strconv"
)

// A Visitor's Visit method is invoked for each node encountered by
// Walk. If the result visitor w is not nil, Walk visits each of the
// children of node with the visitor w, followed by a call of
// w.Visit(nil, nil).
//
type Visitor interface {
	Visit(node ast.Node, id NodeId) (w Visitor)
}

// Helper functions for common node lists. They may be empty.

func walkIdentList(v Visitor, list []*ast.Ident, id NodeId) {
	for _, x := range list {
		id.push(x.Name)
		walk(v, x, id)
		id.pop()
	}
}

func walkExprList(v Visitor, list []ast.Expr, id NodeId) {
	for i, x := range list {
		id.push(strconv.Itoa(i))
		walk(v, x, id)
		id.pop()
	}
}

func walkStmtList(v Visitor, list []ast.Stmt, id NodeId) {
	for i, x := range list {
		id.push(strconv.Itoa(i))
		walk(v, x, id)
		id.pop()
	}
}

func walkDeclList(v Visitor, list []ast.Decl, id NodeId) {
	for i, x := range list {
		id.push(strconv.Itoa(i))
		walk(v, x, id)
		id.pop()
	}
}

// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(node, id); node must not be nil. If the visitor w returned
// by v.Visit(node, id) is not nil, Walk is invoked recursively with
// visitor w for each of the non-nil children of node, followed by a
// call of w.Visit(nil, id).
//
func Walk(v Visitor, n ast.Node) {
	id := make(NodeId, 0)
	walk(v, n, id)
}

func idComponent(node ast.Node) string {
	switch n := node.(type) {
	// Comments and fields
	case *ast.Comment:
		// nothing to do

	case *ast.CommentGroup:

	case *ast.Field:

	case *ast.FieldList:

	// Expressions
	case *ast.BadExpr:

	case *ast.Ident:

	case *ast.BasicLit:

	case *ast.Ellipsis:

	case *ast.FuncLit:

	case *ast.CompositeLit:

	case *ast.ParenExpr:

	case *ast.SelectorExpr:

	case *ast.IndexExpr:

	case *ast.SliceExpr:

	case *ast.TypeAssertExpr:

	case *ast.CallExpr:

	case *ast.StarExpr:

	case *ast.UnaryExpr:

	case *ast.BinaryExpr:

	case *ast.KeyValueExpr:

	// Types
	case *ast.ArrayType:

	case *ast.StructType:

	case *ast.FuncType:

	case *ast.InterfaceType:

	case *ast.MapType:

	case *ast.ChanType:

	// Statements
	case *ast.BadStmt:

	case *ast.DeclStmt:

	case *ast.EmptyStmt:

	case *ast.LabeledStmt:

	case *ast.ExprStmt:

	case *ast.SendStmt:

	case *ast.IncDecStmt:

	case *ast.AssignStmt:

	case *ast.GoStmt:

	case *ast.DeferStmt:

	case *ast.ReturnStmt:

	case *ast.BranchStmt:

	case *ast.BlockStmt:

	case *ast.IfStmt:

	case *ast.CaseClause:

	case *ast.SwitchStmt:

	case *ast.TypeSwitchStmt:

	case *ast.CommClause:

	case *ast.SelectStmt:

	case *ast.ForStmt:

	case *ast.RangeStmt:

	// Declarations
	case *ast.ImportSpec:

	case *ast.ValueSpec:

	case *ast.TypeSpec:

	case *ast.BadDecl:

	case *ast.GenDecl:

	case *ast.FuncDecl:

	// Files and packages
	case *ast.File:
		return n.Name.Name + ".go"

	case *ast.Package:

	default:
		fmt.Printf("ast.walk: unexpected node type %T", n)
		panic("ast.walk")
	}
	return reflect.TypeOf(node).Elem().Name()
}

func walk(v Visitor, node ast.Node, id NodeId) {
	c := idComponent(node)
	if c != "" {
		id.push(c)
		defer id.pop()
	}

	if v = v.Visit(node, id); v == nil {
		return
	}

	// walk children
	// (the order of the cases matches the order
	// of the corresponding node types in ast.go)
	switch n := node.(type) {
	// Comments and fields
	case *ast.Comment:
		// nothing to do

	case *ast.CommentGroup:
		for i, c := range n.List {
			walk(v, c, id.pushed("List", strconv.Itoa(i)))
		}
		id.pop()

	case *ast.Field:
		if n.Doc != nil {
			walk(v, n.Doc, id.pushed("Doc"))
		}
		walkIdentList(v, n.Names, id.pushed("Names"))
		walk(v, n.Type, id.pushed("Type"))
		if n.Tag != nil {
			walk(v, n.Tag, id.pushed("Tag"))
		}
		if n.Comment != nil {
			walk(v, n.Comment, id.pushed("Comment"))
		}

	case *ast.FieldList:
		for i, f := range n.List {
			walk(v, f, id.pushed("List", strconv.Itoa(i)))
		}

	// Expressions
	case *ast.BadExpr:
		// nothing to do

	case *ast.Ident:
		// nothing to do

	case *ast.BasicLit:
		// nothing to do

	case *ast.Ellipsis:
		if n.Elt != nil {
			walk(v, n.Elt, id.pushed("Elt"))
		}

	case *ast.FuncLit:
		walk(v, n.Type, id.pushed("Type"))
		walk(v, n.Body, id.pushed("Body"))

	case *ast.CompositeLit:
		if n.Type != nil {
			walk(v, n.Type, id.pushed("Type"))
		}
		walkExprList(v, n.Elts, id.pushed("Elts"))

	case *ast.ParenExpr:
		walk(v, n.X, id.pushed("X"))

	case *ast.SelectorExpr:
		walk(v, n.X, id.pushed("X"))
		walk(v, n.Sel, id.pushed("Sel"))

	case *ast.IndexExpr:
		walk(v, n.X, id.pushed("X"))
		walk(v, n.Index, id.pushed("Index"))

	case *ast.SliceExpr:
		walk(v, n.X, id.pushed("X"))
		if n.Low != nil {
			walk(v, n.Low, id.pushed("Low"))
		}
		if n.High != nil {
			walk(v, n.High, id.pushed("High"))
		}

	case *ast.TypeAssertExpr:
		walk(v, n.X, id.pushed("X"))
		if n.Type != nil {
			walk(v, n.Type, id.pushed("Type"))
		}

	case *ast.CallExpr:
		walk(v, n.Fun, id.pushed("Fun"))
		walkExprList(v, n.Args, id.pushed("Args"))

	case *ast.StarExpr:
		walk(v, n.X, id.pushed("X"))

	case *ast.UnaryExpr:
		walk(v, n.X, id.pushed("X"))

	case *ast.BinaryExpr:
		walk(v, n.X, id.pushed("X"))
		walk(v, n.Y, id.pushed("Y"))

	case *ast.KeyValueExpr:
		walk(v, n.Key, id.pushed("Key"))
		walk(v, n.Value, id.pushed("Value"))

	// Types
	case *ast.ArrayType:
		if n.Len != nil {
			walk(v, n.Len, id.pushed("Len"))
		}
		walk(v, n.Elt, id.pushed("Elt"))

	case *ast.StructType:
		walk(v, n.Fields, id.pushed("Fields"))

	case *ast.FuncType:
		if n.Params != nil {
			walk(v, n.Params, id.pushed("Params"))
		}
		if n.Results != nil {
			walk(v, n.Results, id.pushed("Results"))
		}

	case *ast.InterfaceType:
		walk(v, n.Methods, id.pushed("Methods"))

	case *ast.MapType:
		walk(v, n.Key, id.pushed("Key"))
		walk(v, n.Value, id.pushed("Value"))

	case *ast.ChanType:
		walk(v, n.Value, id.pushed("Value"))

	// Statements
	case *ast.BadStmt:
		// nothing to do

	case *ast.DeclStmt:
		walk(v, n.Decl, id.pushed("Decl"))

	case *ast.EmptyStmt:
		// nothing to do

	case *ast.LabeledStmt:
		walk(v, n.Label, id.pushed("Label"))
		walk(v, n.Stmt, id.pushed("Stmt"))

	case *ast.ExprStmt:
		walk(v, n.X, id.pushed("X"))

	case *ast.SendStmt:
		walk(v, n.Chan, id.pushed("Chan"))
		walk(v, n.Value, id.pushed("Value"))

	case *ast.IncDecStmt:
		walk(v, n.X, id.pushed("X"))

	case *ast.AssignStmt:
		walkExprList(v, n.Lhs, id.pushed("Lhs"))
		walkExprList(v, n.Rhs, id.pushed("Rhs"))

	case *ast.GoStmt:
		walk(v, n.Call, id.pushed("Call"))

	case *ast.DeferStmt:
		walk(v, n.Call, id.pushed("Call"))

	case *ast.ReturnStmt:
		walkExprList(v, n.Results, id.pushed("Results"))

	case *ast.BranchStmt:
		if n.Label != nil {
			walk(v, n.Label, id.pushed("Label"))
		}

	case *ast.BlockStmt:
		walkStmtList(v, n.List, id.pushed("List"))

	case *ast.IfStmt:
		if n.Init != nil {
			walk(v, n.Init, id.pushed("Init"))
		}
		walk(v, n.Cond, id.pushed("Cond"))
		walk(v, n.Body, id.pushed("Body"))
		if n.Else != nil {
			walk(v, n.Else, id.pushed("Else"))
		}

	case *ast.CaseClause:
		walkExprList(v, n.List, id.pushed("List"))
		walkStmtList(v, n.Body, id.pushed("Body"))

	case *ast.SwitchStmt:
		if n.Init != nil {
			walk(v, n.Init, id.pushed("Init"))
		}
		if n.Tag != nil {
			walk(v, n.Tag, id.pushed("Tag"))
		}
		walk(v, n.Body, id.pushed("Body"))

	case *ast.TypeSwitchStmt:
		if n.Init != nil {
			walk(v, n.Init, id.pushed("Init"))
		}
		walk(v, n.Assign, id.pushed("Assign"))
		walk(v, n.Body, id.pushed("Body"))

	case *ast.CommClause:
		if n.Comm != nil {
			walk(v, n.Comm, id.pushed("Comm"))
		}
		walkStmtList(v, n.Body, id.pushed("Body"))

	case *ast.SelectStmt:
		walk(v, n.Body, id.pushed("Body"))

	case *ast.ForStmt:
		if n.Init != nil {
			walk(v, n.Init, id.pushed("Init"))
		}
		if n.Cond != nil {
			walk(v, n.Cond, id.pushed("Cond"))
		}
		if n.Post != nil {
			walk(v, n.Post, id.pushed("Post"))
		}
		walk(v, n.Body, id.pushed("Body"))

	case *ast.RangeStmt:
		walk(v, n.Key, id.pushed("Key"))
		if n.Value != nil {
			walk(v, n.Value, id.pushed("Value"))
		}
		walk(v, n.X, id.pushed("X"))
		walk(v, n.Body, id.pushed("Body"))

	// Declarations
	case *ast.ImportSpec:
		if n.Doc != nil {
			walk(v, n.Doc, id.pushed("Doc"))
		}
		if n.Name != nil {
			walk(v, n.Name, id.pushed("Name"))
		}
		walk(v, n.Path, id.pushed("Path"))
		if n.Comment != nil {
			walk(v, n.Comment, id.pushed("Comment"))
		}

	case *ast.ValueSpec:
		if n.Doc != nil {
			walk(v, n.Doc, id.pushed("Doc"))
		}
		walkIdentList(v, n.Names, id.pushed("Names"))
		if n.Type != nil {
			walk(v, n.Type, id.pushed("Type"))
		}
		walkExprList(v, n.Values, id.pushed("Values"))
		if n.Comment != nil {
			walk(v, n.Comment, id.pushed("Comment"))
		}

	case *ast.TypeSpec:
		if n.Doc != nil {
			walk(v, n.Doc, id.pushed("Doc"))
		}
		walk(v, n.Name, id.pushed("Name"))
		walk(v, n.Type, id.pushed("Type"))
		if n.Comment != nil {
			walk(v, n.Comment, id.pushed("Comment"))
		}

	case *ast.BadDecl:
		// nothing to do

	case *ast.GenDecl:
		if n.Doc != nil {
			walk(v, n.Doc, id.pushed("Doc"))
		}
		for i, s := range n.Specs {
			walk(v, s, id.pushed("Specs", strconv.Itoa(i)))
		}

	case *ast.FuncDecl:
		if n.Doc != nil {
			walk(v, n.Doc, id.pushed("Doc"))
		}
		if n.Recv != nil {
			walk(v, n.Recv, id.pushed("Recv"))
		}
		walk(v, n.Name, id.pushed("Name"))
		walk(v, n.Type, id.pushed("Type"))
		if n.Body != nil {
			walk(v, n.Body, id.pushed("Body"))
		}

	// Files and packages
	case *ast.File:
		if n.Doc != nil {
			walk(v, n.Doc, id.pushed("Doc"))
		}
		walk(v, n.Name, id.pushed("Name"))
		walkDeclList(v, n.Decls, id.pushed("Decls"))
		// don't walk n.Comments - they have been
		// visited already through the individual
		// nodes

	case *ast.Package:
		for _, f := range n.Files {
			walk(v, f, id.pushed("Files", f.Name.Name))
		}

	default:
		fmt.Printf("ast.walk: unexpected node type %T", n)
		panic("ast.walk")
	}

	v.Visit(nil, id)
}

type inspector func(ast.Node, NodeId) bool

func (f inspector) Visit(node ast.Node, id NodeId) Visitor {
	if f(node, id) {
		return f
	}
	return nil
}

// Inspect traverses an AST in depth-first order: It starts by calling
// f(node, id); node must not be nil. If f returns true, Inspect invokes f
// for all the non-nil children of node, recursively.
//
func Inspect(node ast.Node, f func(ast.Node, NodeId) bool) {
	Walk(inspector(f), node)
}

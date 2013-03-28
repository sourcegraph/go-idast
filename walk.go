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
		walk(v, x, id.pushed(x.Name))
	}
}

func walkExprList(v Visitor, list []ast.Expr, id NodeId) {
	for i, x := range list {
		walk(v, x, id.pushed(strconv.Itoa(i)))
	}
}

func walkStmtList(v Visitor, list []ast.Stmt, id NodeId) {
	for i, x := range list {
		walk(v, x, id.pushed(strconv.Itoa(i)))
	}
}

func walkDeclList(v Visitor, list []ast.Decl, id NodeId) {
	for i, x := range list {
		walk(v, x, id.pushed(strconv.Itoa(i)))
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
			walk(v, c, id.pushed(strconv.Itoa(i)))
		}
		id.pop()

	case *ast.Field:
		if n.Doc != nil {
			walk(v, n.Doc, id)
		}
		walkIdentList(v, n.Names, id)
		walk(v, n.Type, id)
		if n.Tag != nil {
			walk(v, n.Tag, id)
		}
		if n.Comment != nil {
			walk(v, n.Comment, id)
		}

	case *ast.FieldList:
		for _, f := range n.List {
			walk(v, f, id)
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
			walk(v, n.Elt, id)
		}

	case *ast.FuncLit:
		walk(v, n.Type, id)
		walk(v, n.Body, id)

	case *ast.CompositeLit:
		if n.Type != nil {
			walk(v, n.Type, id)
		}
		walkExprList(v, n.Elts, id)

	case *ast.ParenExpr:
		walk(v, n.X, id)

	case *ast.SelectorExpr:
		walk(v, n.X, id)
		walk(v, n.Sel, id)

	case *ast.IndexExpr:
		walk(v, n.X, id)
		walk(v, n.Index, id)

	case *ast.SliceExpr:
		walk(v, n.X, id)
		if n.Low != nil {
			walk(v, n.Low, id)
		}
		if n.High != nil {
			walk(v, n.High, id)
		}

	case *ast.TypeAssertExpr:
		walk(v, n.X, id)
		if n.Type != nil {
			walk(v, n.Type, id)
		}

	case *ast.CallExpr:
		walk(v, n.Fun, id)
		walkExprList(v, n.Args, id)

	case *ast.StarExpr:
		walk(v, n.X, id)

	case *ast.UnaryExpr:
		walk(v, n.X, id)

	case *ast.BinaryExpr:
		walk(v, n.X, id)
		walk(v, n.Y, id)

	case *ast.KeyValueExpr:
		walk(v, n.Key, id)
		walk(v, n.Value, id)

	// Types
	case *ast.ArrayType:
		if n.Len != nil {
			walk(v, n.Len, id)
		}
		walk(v, n.Elt, id)

	case *ast.StructType:
		walk(v, n.Fields, id)

	case *ast.FuncType:
		if n.Params != nil {
			walk(v, n.Params, id)
		}
		if n.Results != nil {
			walk(v, n.Results, id)
		}

	case *ast.InterfaceType:
		walk(v, n.Methods, id)

	case *ast.MapType:
		walk(v, n.Key, id)
		walk(v, n.Value, id)

	case *ast.ChanType:
		walk(v, n.Value, id)

	// Statements
	case *ast.BadStmt:
		// nothing to do

	case *ast.DeclStmt:
		walk(v, n.Decl, id)

	case *ast.EmptyStmt:
		// nothing to do

	case *ast.LabeledStmt:
		walk(v, n.Label, id)
		walk(v, n.Stmt, id)

	case *ast.ExprStmt:
		walk(v, n.X, id)

	case *ast.SendStmt:
		walk(v, n.Chan, id)
		walk(v, n.Value, id)

	case *ast.IncDecStmt:
		walk(v, n.X, id)

	case *ast.AssignStmt:
		walkExprList(v, n.Lhs, id.pushed("lhs"))
		walkExprList(v, n.Rhs, id.pushed("rhs"))

	case *ast.GoStmt:
		walk(v, n.Call, id)

	case *ast.DeferStmt:
		walk(v, n.Call, id)

	case *ast.ReturnStmt:
		walkExprList(v, n.Results, id)

	case *ast.BranchStmt:
		if n.Label != nil {
			walk(v, n.Label, id)
		}

	case *ast.BlockStmt:
		walkStmtList(v, n.List, id)

	case *ast.IfStmt:
		if n.Init != nil {
			walk(v, n.Init, id)
		}
		walk(v, n.Cond, id)
		walk(v, n.Body, id)
		if n.Else != nil {
			walk(v, n.Else, id)
		}

	case *ast.CaseClause:
		walkExprList(v, n.List, id)
		walkStmtList(v, n.Body, id)

	case *ast.SwitchStmt:
		if n.Init != nil {
			walk(v, n.Init, id)
		}
		if n.Tag != nil {
			walk(v, n.Tag, id)
		}
		walk(v, n.Body, id)

	case *ast.TypeSwitchStmt:
		if n.Init != nil {
			walk(v, n.Init, id)
		}
		walk(v, n.Assign, id)
		walk(v, n.Body, id)

	case *ast.CommClause:
		if n.Comm != nil {
			walk(v, n.Comm, id)
		}
		walkStmtList(v, n.Body, id)

	case *ast.SelectStmt:
		walk(v, n.Body, id)

	case *ast.ForStmt:
		if n.Init != nil {
			walk(v, n.Init, id)
		}
		if n.Cond != nil {
			walk(v, n.Cond, id)
		}
		if n.Post != nil {
			walk(v, n.Post, id)
		}
		walk(v, n.Body, id)

	case *ast.RangeStmt:
		walk(v, n.Key, id)
		if n.Value != nil {
			walk(v, n.Value, id)
		}
		walk(v, n.X, id)
		walk(v, n.Body, id)

	// Declarations
	case *ast.ImportSpec:
		if n.Doc != nil {
			walk(v, n.Doc, id)
		}
		if n.Name != nil {
			walk(v, n.Name, id)
		}
		walk(v, n.Path, id)
		if n.Comment != nil {
			walk(v, n.Comment, id)
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
			walk(v, n.Doc, id)
		}
		walk(v, n.Name, id)
		walk(v, n.Type, id)
		if n.Comment != nil {
			walk(v, n.Comment, id)
		}

	case *ast.BadDecl:
		// nothing to do

	case *ast.GenDecl:
		if n.Doc != nil {
			walk(v, n.Doc, id)
		}
		for _, s := range n.Specs {
			walk(v, s, id)
		}

	case *ast.FuncDecl:
		if n.Doc != nil {
			walk(v, n.Doc, id)
		}
		if n.Recv != nil {
			walk(v, n.Recv, id)
		}
		walk(v, n.Name, id)
		walk(v, n.Type, id)
		if n.Body != nil {
			walk(v, n.Body, id)
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
			walk(v, f, id.pushed(f.Name.Name))
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

package idast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path"
	"reflect"
	"strings"
	"testing"
)

func TestXref(t *testing.T) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "testdata", goFilesOnly, parser.ParseComments)
	if err != nil {
		t.Errorf("Error parsing testdata dir: %v", err)
		return
	}

	for _, pkg := range pkgs {
		for filename, file := range pkg.Files {
			ns := collect(file)
			checkUnique(filename, ns, t)
			checkOutput(filename, ns, t)
		}
	}
}

func goFilesOnly(file os.FileInfo) bool {
	return file.Mode().IsRegular() && path.Ext(file.Name()) == ".go"
}

func collect(node ast.Node) (nodes []NodeWithId) {
	nodes = make([]NodeWithId, 0)
	Inspect(node, func(node ast.Node, id NodeId) bool {
		if node != nil {
			nodes = append(nodes, NodeWithId{node, id.dup()})
		}
		return true
	})
	return
}

func checkUnique(srcFilename string, ns []NodeWithId, t *testing.T) {
	ids := make(map[string]NodeWithId, 0)
	for _, n := range ns {
		if existing, hasExisting := ids[n.Id.String()]; hasExisting {
			t.Errorf("duplicate NodeId '%s' for nodes:\n%v\n\n-- and --\n\n%v", n.Id.String(), existing.Node, n.Node)
		}
	}
}

func checkOutput(srcFilename string, ns []NodeWithId, t *testing.T) {
	actualFilename := srcFilename + "_actual.json"
	expectedFilename := srcFilename + "_expected.json"

	// write actual output
	writeJson(actualFilename, ns)

	// diff
	cmd := exec.Command("diff", "-u", expectedFilename, actualFilename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Wait()
	if !cmd.ProcessState.Success() {
		t.Errorf("%s: actual output did not match expected output", srcFilename)
	}
}

func writeJson(filename string, ns []NodeWithId) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic("Error opening file: " + err.Error())
	}
	defer f.Close()

	for _, n := range ns {
		io.WriteString(f, fmt.Sprintf(" %-15s | %-31.31s | %s\n", reflect.TypeOf(n.Node).Elem().Name(), strings.Replace(pretty(n.Node), "\n", "\\n", -1), n.Id.String()))
	}
}

var emptyFileSet = token.NewFileSet()

func pretty(n ast.Node) string {
	var b bytes.Buffer
	printer.Fprint(&b, emptyFileSet, n)
	return b.String()
}

package tagseq

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `report inconsistent struct tags`

var Analyzer = &analysis.Analyzer{
	Name:     "tagseq",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return
		}
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return
		}
		for _, field := range st.Fields.List {
			if field.Tag == nil {
				continue
			}
			tagVal := strings.Trim(field.Tag.Value, "`")
			// using reflect.StructTag(tagValue) is too
			// limited for what we need :/
			knownNames := make(map[string]bool)
			for _, tag := range strings.Split(tagVal, " ") {
				l := strings.SplitN(tag, ":", 2)
				if len(l) < 2 {
					pass.Reportf(n.Pos(), "no : in %q for %q", tag, tagVal)
					continue
				}
				tagName := l[1]
				if len(knownNames) == 0 {
					knownNames[tagName] = true
					continue
				}
				if !knownNames[tagName] {
					pass.Reportf(n.Pos(), "inconsistent struct tags found in %q", tagVal)
				}
			}
		}
	})
	return nil, nil
}

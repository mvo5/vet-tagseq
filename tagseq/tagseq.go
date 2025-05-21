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
		knownNames := make(map[string]bool)
		for _, field := range st.Fields.List {
			if field.Tag == nil {
				continue
			}
			tagVal := strings.Trim(field.Tag.Value, "`")
			// Compare subsequent names with the first, all have to match
			var tagName string
			for _, tag := range strings.Split(tagVal, " ") {
				l := strings.SplitN(tag, ":", 2)
				if len(l) < 2 {
					pass.Reportf(n.Pos(), "no : in %q for %q", tag, tagVal)
					continue
				}
				if tagName == "" {
					tagName = l[1]
					if knownNames[tagName] {
						pass.Reportf(n.Pos(), "duplicate struct tags found in %q", tagVal)
					}
					knownNames[tagName] = true
					continue
				}

				if tagName != l[1] {
					pass.Reportf(n.Pos(), "inconsistent struct tags found in %q", tagVal)
				}
			}
		}
	})
	return nil, nil
}

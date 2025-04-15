// Package analyzer defines the unsignedlint analysis logic.
package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `checks for potentially unsafe unsigned integer subtractions

The linter checks for binary expressions of the form 'x - y' where 'y'
is an unsigned integer type (uint, uint8, uint16, uint32, uint64, uintptr).
Such operations can lead to unexpected integer wraps (underflow) if y > x.`

// Analyzer defines the unsignedlint analysis tool.
var Analyzer = &analysis.Analyzer{
	Name:     "unsignedlint",
	Doc:      Doc,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// Use the alias for the Analyzer and the type assertion
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}

	ins.Preorder(nodeFilter, func(n ast.Node) {
		expr := n.(*ast.BinaryExpr)

		// Check if the operation is subtraction
		if expr.Op != token.SUB {
			return
		}

		// Check the type of the right-hand operand (Y)
		yType := pass.TypesInfo.TypeOf(expr.Y)
		xType := pass.TypesInfo.TypeOf(expr.X)
		if yType == nil || xType == nil {
			return // Skip if type information is unavailable
		}

		// Only warn if both operands are unsigned, and the right is not a constant identifier
		xBasic, xOk := xType.Underlying().(*types.Basic)
		yBasic, yOk := yType.Underlying().(*types.Basic)
		if xOk && yOk {
			switch xBasic.Kind() {
			case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64, types.Uintptr:
				switch yBasic.Kind() {
				case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64, types.Uintptr:
					// Check the right operand Y

					// Report if Y is a basic literal (e.g., u8 - 10)
					if _, isLit := expr.Y.(*ast.BasicLit); isLit {
						pass.Reportf(expr.Pos(), "subtraction with unsigned integer operand '%s' may underflow", types.ExprString(expr.Y))
						return // Report and exit for this node
					}

					// Report if Y is an identifier that is NOT a compile-time constant
					if ident, isIdent := expr.Y.(*ast.Ident); isIdent {
						// Look up the object the identifier refers to
						if obj, ok := pass.TypesInfo.Uses[ident]; ok {
							// Skip if the object is a constant
							if _, isConst := obj.(*types.Const); isConst {
								return
							}
							pass.Reportf(expr.Pos(), "subtraction with unsigned integer operand '%s' may underflow", types.ExprString(ident))
						}
						// If the object lookup fails (shouldn't typically happen for valid code), we don't report.
					}
					// Other cases for expr.Y (like function calls, etc.) are currently ignored.
				}
			}
		}
	})

	return nil, nil
}

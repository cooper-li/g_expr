package g_expr

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

const (
	funcInData  = "in_data"
	funcInArray = "in_array"
)

func matchFunc(expr *ast.CallExpr, sourceData map[string]interface{}) (bool, error) {
	funIdent, ok := expr.Fun.(*ast.Ident)
	if !ok {
		return false, fmt.Errorf("CallExpr node error, node info: %v", expr)
	}

	switch funIdent.Name {
	case funcInData:
		return InOp(expr.Args, sourceData)
	default:
		return false, fmt.Errorf("not support func: %s", funIdent.Name)
	}
}

func InOp(args []ast.Expr, sourceData map[string]interface{}) (bool, error) {
	if len(args) != 2 {
		return false, fmt.Errorf("in_data args length error")
	}
	key, ok := args[0].(*ast.BasicLit)
	if !ok || key.Kind != token.STRING {
		return false, fmt.Errorf("in_data args key type error, node: %v", args)
	}
	data, ok := sourceData[trimQuotes(key.Value)]
	if !ok {
		return false, fmt.Errorf("in_data args key not exists in source data, node: %v", args)
	}
	dataString, err := convToString(data)
	if err != nil {
		return false, err
	}

	valBasicLit, ok := args[1].(*ast.BasicLit)
	if !ok {
		return false, fmt.Errorf("in_data args val type error, node: %#v", args)
	}
	valList := strings.Split(trimQuotes(valBasicLit.Value), "|")
	for _, val := range valList {
		if val == dataString {
			return true, nil
		}
	}
	return false, nil
}

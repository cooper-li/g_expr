package g_expr

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Match(exprRule string, sourceData map[string]interface{}) (bool, error) {
	if len(exprRule) == 0 {
		return false, errors.New("empty exprRule")
	}

	// 解析表达式
	exprAst, err := parser.ParseExpr(exprRule)
	if err != nil {
		return false, fmt.Errorf("parse exprRule err: %w", err)
	}
	//fset := token.NewFileSet()
	//ast.Print(fset, exprAst)

	res, err := judge(exprAst, sourceData)
	if err != nil {
		return false, fmt.Errorf("judge err: %w", err)
	}
	return res, nil
}

// 递归解析ast
func judge(expr ast.Expr, sourceData map[string]interface{}) (bool, error) {

	switch t := expr.(type) {
	case *ast.BinaryExpr:
		if isBinaryLeaf(t) {
			return cmpBinary(t, sourceData)
		}
		// 递归比较
		lRes, err := judge(t.X, sourceData)
		// fmt.Printf("-- lr: %v, op: %v \n", lRes, t.Op)
		if err != nil {
			return false, err
		}
		rRes, err := judge(t.Y, sourceData)
		//fmt.Printf("## lr: %v, rr: %v, op: %v \n", lRes, rRes, t.Op)
		if err != nil {
			return false, err
		}

		switch t.Op {
		case token.LAND:
			return lRes && rRes, nil
		case token.LOR:
			return lRes || rRes, nil
		}
		return false, fmt.Errorf("not support op xx")
	case *ast.CallExpr: // 匹配到函数
		return matchFunc(t, sourceData)
	case *ast.ParenExpr:
		return judge(t.X, sourceData)
	default:
		return false, errors.New(fmt.Sprintf("%#v type is not support", expr))
	}
}

func isBinaryLeaf(expr *ast.BinaryExpr) bool {
	// 二元表达式的最小单位，左节点是Ident，右节点是BasicLit
	_, lType := expr.X.(*ast.Ident)
	_, rType := expr.Y.(*ast.BasicLit)
	return lType && rType
}

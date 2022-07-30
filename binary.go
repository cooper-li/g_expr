package g_expr

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

func cmpBinary(expr *ast.BinaryExpr, sourceData map[string]interface{}) (bool, error) {
	var (
		xName string
		xVal  interface{}
	)

	xName = expr.X.(*ast.Ident).Name
	xVal, ok := sourceData[xName]
	if !ok {
		return false, fmt.Errorf("source data key: %s not exists", xName)
	}
	y := expr.Y.(*ast.BasicLit)

	// 根据不同类型比较
	switch y.Kind {
	case token.INT: // 转换为int64比较
		xValInt64, err := convToInt64(xVal)
		if err != nil {
			return false, err
		}
		yInt64, err := strconv.ParseInt(y.Value, 10, 64)
		if err != nil {
			return false, err
		}
		return cmpInt64(xValInt64, yInt64, expr.Op)
	case token.STRING:
		xValString, err := convToString(xVal)
		if err != nil {
			return false, err
		}
		return cmpString(xValString, strings.Trim(y.Value, "\""), expr.Op)
	}

	return false, fmt.Errorf("cmpBinary error")
}

func cmpString(x, y string, op token.Token) (bool, error) {
	switch op {
	case token.EQL:
		return x == y, nil
	default:
		return false, fmt.Errorf("not support string op: %v", op)
	}
}

func cmpInt64(x, y int64, op token.Token) (bool, error) {
	switch op {
	case token.EQL:
		return x == y, nil
	case token.LSS:
		return x < y, nil
	case token.GTR:
		return x > y, nil
	case token.NEQ:
		return x != y, nil
	case token.LEQ:
		return x <= y, nil
	case token.GEQ:
		return x >= y, nil
	default:
		return false, fmt.Errorf("not support number op: %v", op)
	}
}

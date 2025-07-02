package resolvers

import (
	"memory/src/util"
	"strings"
)

func ResolveCalculsInCondition(expr string) string {
	expr = strings.TrimSpace(expr)
	level := 0
	for i := 0; i < len(expr)-1; i++ {
		if expr[i] == '(' {
			level++
		} else if expr[i] == ')' {
			level--
		} else if level == 0 && expr[i] == '&' && expr[i+1] == '&' {
			left := ResolveCalculsInCondition(expr[:i])
			right := ResolveCalculsInCondition(expr[i+2:])
			return left + "&&" + right
		} else if level == 0 && expr[i] == '|' && expr[i+1] == '|' {
			left := ResolveCalculsInCondition(expr[:i])
			right := ResolveCalculsInCondition(expr[i+2:])
			return left + "||" + right
		}
	}
	ops := []string{"==", ">=", "<=", "!=", ">", "<"}
	for _, op := range ops {
		idx := strings.Index(expr, op)
		if idx != -1 {
			left := strings.TrimSpace(expr[:idx])
			right := strings.TrimSpace(expr[idx+len(op):])
			left = ReplaceVariablesInExpr(left)
			if util.LooksLikeCalcul(left) || util.IsNumber(left) {
				left = ResolveCalculs(left)
			}
			right = ReplaceVariablesInExpr(right)
			if util.LooksLikeCalcul(right) || util.IsNumber(right) {
				right = ResolveCalculs(right)
			}
			return left + op + right
		}
	}
	return expr
}
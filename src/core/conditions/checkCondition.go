package conditions

import (
	"fmt"
	"memory/src/core/resolvers"
	"os"
	"strconv"
	"strings"
    "memory/src/util"
)

func CheckCondition(condition string) bool {
    condition = strings.TrimSpace(condition)
    if condition == "" {
        fmt.Printf("[73402] - Invalid condition: empty string: '%s'\n", condition)
        os.Exit(1)
    }

    if condition == "true" {
        return true
    }
    if condition == "false" {
        return false
    }

    ops := []string{"==", ">=", "<=", "!=", ">", "<"}
    var op string
    var idx int
    for _, o := range ops {
        if i := strings.Index(condition, o); i != -1 {
            op = o
            idx = i
            break
        }
    }
    if op == "" {
        fmt.Printf("[73402] - No valid comparison operator found in: '%s'\n", condition)
        os.Exit(1)
    }

    left := strings.TrimSpace(condition[:idx])
    right := strings.TrimSpace(condition[idx+len(op):])

    if left == "" || right == "" {
        fmt.Printf("[73402] - One side of the comparison is empty in: '%s'\n", condition)
        os.Exit(1)
    }

    leftVal, _ := resolvers.ResolveValue(left)
    leftVal = resolvers.ReplaceVariablesInExpr(leftVal)
    if util.LooksLikeCalcul(leftVal) || util.IsNumber(leftVal) {
        leftVal = resolvers.ResolveCalculs(leftVal)
    }
    rightVal, _ := resolvers.ResolveValue(right)
    rightVal = resolvers.ReplaceVariablesInExpr(rightVal)
    if util.LooksLikeCalcul(rightVal) || util.IsNumber(rightVal) {
        rightVal = resolvers.ResolveCalculs(rightVal)
    }

    leftNum, leftErr := strconv.ParseFloat(leftVal, 64)
    rightNum, rightErr := strconv.ParseFloat(rightVal, 64)

    if leftErr == nil && rightErr == nil {
        switch op {
        case "==":
            return leftNum == rightNum
        case "!=":
            return leftNum != rightNum
        case ">":
            return leftNum > rightNum
        case "<":
            return leftNum < rightNum
        case ">=":
            return leftNum >= rightNum
        case "<=":
            return leftNum <= rightNum
        }
    }

    if (leftVal == "true" || leftVal == "false") && (rightVal == "true" || rightVal == "false") {
        leftBool := leftVal == "true"
        rightBool := rightVal == "true"
        switch op {
        case "==":
            return leftBool == rightBool
        case "!=":
            return leftBool != rightBool
        default:
            fmt.Printf("[73402] - Invalid operator for booleans in: '%s'\n", condition)
            os.Exit(1)
        }
    }

    if strings.HasPrefix(leftVal, "\"") && strings.HasSuffix(leftVal, "\"") &&
        strings.HasPrefix(rightVal, "\"") && strings.HasSuffix(rightVal, "\"") {
        leftStr := leftVal[1 : len(leftVal)-1]
        rightStr := rightVal[1 : len(rightVal)-1]
        switch op {
        case "==":
            return leftStr == rightStr
        case "!=":
            return leftStr != rightStr
        default:
            fmt.Printf("[73402] - Invalid operator for strings in: '%s'\n", condition)
            os.Exit(1)
        }
    }

    if leftVal == "nil" && rightVal == "nil" {
        switch op {
        case "==":
            return true
        case "!=":
            return false
        default:
            fmt.Printf("[73402] - Invalid operator for nil in: '%s'\n", condition)
            os.Exit(1)
        }
    }

    fmt.Printf("[73402] - Type mismatch or unsupported types in: '%s'\n", condition)
    os.Exit(1)
    return false
}
package conditions

import (
    "fmt"
    "os"
    "strings"
)

func EvaluateConditions(expression string) bool {
    expression = strings.TrimSpace(expression)
    if expression == "" {
        fmt.Printf("[73402] - Empty condition expression: '%s'\n", expression)
        os.Exit(1)
    }

    result, err := evalCondition(expression)
    if err != nil {
        fmt.Printf("[73402] - Invalid condition: '%s' (%v)\n", expression, err)
        os.Exit(1)
    }
    return result
}

func evalCondition(expr string) (bool, error) {
    expr = strings.TrimSpace(expr)
    if expr == "" {
        return false, fmt.Errorf("empty condition")
    }

    for strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") && isBalanced(expr[1:len(expr)-1]) {
        expr = strings.TrimSpace(expr[1 : len(expr)-1])
    }

    level := 0
    for i := 0; i < len(expr)-1; i++ {
        if expr[i] == '(' {
            level++
        } else if expr[i] == ')' {
            level--
        } else if level == 0 && expr[i] == '|' && expr[i+1] == '|' {
            left, err := evalCondition(expr[:i])
            if err != nil {
                return false, err
            }
            right, err := evalCondition(expr[i+2:])
            if err != nil {
                return false, err
            }
            return left || right, nil
        }
    }

    level = 0
    for i := 0; i < len(expr)-1; i++ {
        if expr[i] == '(' {
            level++
        } else if expr[i] == ')' {
            level--
        } else if level == 0 && expr[i] == '&' && expr[i+1] == '&' {
            left, err := evalCondition(expr[:i])
            if err != nil {
                return false, err
            }
            right, err := evalCondition(expr[i+2:])
            if err != nil {
                return false, err
            }
            return left && right, nil
        }
    }

    return CheckCondition(expr), nil
}
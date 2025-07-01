package conditions

import (
    "fmt"
    "regexp"
    "strings"
)

func IsCondition(expression string) bool {
    expression = strings.TrimSpace(expression)
    if expression == "" {
        return false
    }

    count := 0
    for _, c := range expression {
        if c == '(' {
            count++
        } else if c == ')' {
            count--
            if count < 0 {
                return false
            }
        }
    }
    if count != 0 {
        return  false
    }

    valid := regexp.MustCompile(`^[\w\s\(\)\|\&\!\=\>\<\.\$]+$`).MatchString
    if !valid(expression) {
        return false
    }

    _, err := parseCondition(expression)
    return err == nil
}

func parseCondition(expr string) (any, error) {
    expr = strings.TrimSpace(expr)
    if expr == "" {
        return nil, fmt.Errorf("empty condition")
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
            left, err := parseCondition(expr[:i])
            if err != nil {
                return nil, err
            }
            right, err := parseCondition(expr[i+2:])
            if err != nil {
                return nil, err
            }
            return []any{"||", left, right}, nil
        }
    }

    level = 0
    for i := 0; i < len(expr)-1; i++ {
        if expr[i] == '(' {
            level++
        } else if expr[i] == ')' {
            level--
        } else if level == 0 && expr[i] == '&' && expr[i+1] == '&' {
            left, err := parseCondition(expr[:i])
            if err != nil {
                return nil, err
            }
            right, err := parseCondition(expr[i+2:])
            if err != nil {
                return nil, err
            }
            return []any{"&&", left, right}, nil
        }
    }

    if isComparison(expr) {
        return expr, nil
    }
    return nil, fmt.Errorf("not a valid comparison: '%s'", expr)
}

func isBalanced(expr string) bool {
    count := 0
    for _, c := range expr {
        if c == '(' {
            count++
        } else if c == ')' {
            count--
            if count < 0 {
                return false
            }
        }
    }
    return count == 0
}

func isComparison(expr string) bool {
    ops := []string{"==", ">=", "<=", "!=", ">", "<"}
    for _, op := range ops {
        if strings.Contains(expr, op) {
            return true
        }
    }
    trim := strings.TrimSpace(expr)
    return trim == "true" || trim == "false"
}
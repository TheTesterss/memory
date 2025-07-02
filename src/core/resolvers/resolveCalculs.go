package resolvers

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ResolveCalculs(s string) string {
    s = strings.TrimSpace(s)
    if s == "" {
        return ""
    }
    val, err := evalExpr(s)
    if err != nil {
        fmt.Printf("[73402] - Invalid calculation: '%s' (%v)\n", s, err)
        os.Exit(1)
    }
    return fmt.Sprintf("%v", val)
}

func evalExpr(expr string) (float64, error) {
    tokens := tokenize(expr)
    rpn, err := shuntingYard(tokens)
    if err != nil {
        return 0, err
    }
    return evalRPN(rpn)
}

func tokenize(expr string) []string {
    var tokens []string
    var num strings.Builder
    for i := 0; i < len(expr); i++ {
        c := expr[i]
        if c == ' ' {
            continue
        }
        if (c >= '0' && c <= '9') || c == '.' {
            num.WriteByte(c)
        } else {
            if num.Len() > 0 {
                tokens = append(tokens, num.String())
                num.Reset()
            }
            if c == '(' || c == ')' {
                tokens = append(tokens, string(c))
            } else if strings.ContainsRune("+-*/%^", rune(c)) {
                tokens = append(tokens, string(c))
            }
        }
    }
    if num.Len() > 0 {
        tokens = append(tokens, num.String())
    }
    return tokens
}

func precedence(op string) int {
    switch op {
    case "^":
        return 4
    case "*", "/", "%":
        return 3
    case "+", "-":
        return 2
    default:
        return 0
    }
}

func isRightAssoc(op string) bool {
    return op == "^"
}

func shuntingYard(tokens []string) ([]string, error) {
    var output []string
    var stack []string
    for _, token := range tokens {
        if _, err := strconv.ParseFloat(token, 64); err == nil {
            output = append(output, token)
        } else if token == "(" {
            stack = append(stack, token)
        } else if token == ")" {
            for len(stack) > 0 && stack[len(stack)-1] != "(" {
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            if len(stack) == 0 {
                return nil, fmt.Errorf("mismatched parentheses")
            }
            stack = stack[:len(stack)-1]
        } else if strings.Contains("+-*/%^", token) {
            for len(stack) > 0 {
                top := stack[len(stack)-1]
                if top == "(" {
                    break
                }
                if (precedence(token) < precedence(top)) ||
                    (precedence(token) == precedence(top) && !isRightAssoc(token)) {
                    output = append(output, top)
                    stack = stack[:len(stack)-1]
                } else {
                    break
                }
            }
            stack = append(stack, token)
        }
    }
    for len(stack) > 0 {
        if stack[len(stack)-1] == "(" {
            return nil, fmt.Errorf("mismatched parentheses")
        }
        output = append(output, stack[len(stack)-1])
        stack = stack[:len(stack)-1]
    }
    return output, nil
}

func evalRPN(tokens []string) (float64, error) {
    var stack []float64
    for _, token := range tokens {
        if v, err := strconv.ParseFloat(token, 64); err == nil {
            stack = append(stack, v)
        } else {
            if len(stack) < 2 {
                return 0, fmt.Errorf("not enough operands for '%s'", token)
            }
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            var res float64
            switch token {
            case "+":
                res = a + b
            case "-":
                res = a - b
            case "*":
                res = a * b
            case "/":
                res = a / b
            case "%":
                res = math.Mod(a, b)
            case "^":
                res = math.Pow(a, b)
            default:
                return 0, fmt.Errorf("unknown operator '%s'", token)
            }
            stack = append(stack, res)
        }
    }
    if len(stack) != 1 {
        return 0, fmt.Errorf("invalid expression")
    }
    return stack[0], nil
}
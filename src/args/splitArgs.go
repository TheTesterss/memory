package args

import (
	"fmt"
	"memory/src/core/conditions"
	"memory/src/core/resolvers"
	"memory/src/types"
	"memory/src/util"
	"os"
	"strings"
)

func Split(f types.Item) []types.Arg {
	if f.Listed_args == "" {
		return []types.Arg{}
	}
	var r []types.Arg = []types.Arg{
		{T: "", Value: ""},
	}
	var inString bool = false

	for i, char := range f.Listed_args {

		switch string(char) {

		case "\"":
    		arg := &r[len(r)-1]
    		if i > 0 && string(f.Listed_args[i-1]) == "\\" {
    		    arg.Value += "\""
    		    continue
    		}
    		inString = !inString
    		if !inString {
    		    isLast := i == len(f.Listed_args)-1
    		    if !isLast {
    				next := string(f.Listed_args[i+1])
    				next2 := ""
    				if i+2 < len(f.Listed_args) {
        				next2 = string(f.Listed_args[i+1 : i+3])
    				}
    				allowed := next == ";" || next == "]" ||
        			next2 == "==" || next2 == "!=" || next2 == ">=" || next2 == "<=" ||
        			next == ">" || next == "<" ||
        			next2 == "&&" || next2 == "||"
    				if !allowed {
        				fmt.Printf("[73402] - At line %d - Closing a string but not the argument.\n", f.Line)
        				os.Exit(1)
    				}
				}
    		}
    		if inString {
        		arg.T = "str"
    		}

		case ";":
			if inString {
				arg := &r[len(r)-1]
				arg.Value+=";"
			}
			v := &r[len(r)-1]
			v.Value = strings.TrimSpace(v.Value)
			if !isBalancedParentheses(v.Value) {
    			for strings.HasSuffix(v.Value, ")") && !strings.HasPrefix(v.Value, "\"") {
        			v.Value = strings.TrimSuffix(v.Value, ")")
        			v.Value = strings.TrimSpace(v.Value)
    			}
			}
			if v.T != "str" {
    			value, t := resolvers.ResolveValue(v.Value)
    			v.Value = value
    			if t != "" {
    			    v.T = t
    			}
			}

			if v.T != "str" && !inString {
	    		v.Value = strings.TrimSpace(v.Value)
	    		v.Value = resolvers.ReplaceVariablesInExpr(v.Value)

 	   			if conditions.IsCondition(v.Value) {
	    	    	result := conditions.EvaluateConditions(v.Value)
	    		    v.Value = fmt.Sprintf("%v", result)
	    		    v.T = "bool"
 	    		} else if util.LooksLikeCalcul(v.Value) {
 	    		    v.Value = resolvers.ResolveCalculs(v.Value)
 	    		    v.T = "int"
 	    		} else if util.IsNumber(v.Value) {
 	    		    v.T = "int"
 	    		} else if v.Value == "nil" {
        			v.T = "nil"
    			} else if util.IsBoolean(v.Value) {
    			    v.T = "bool"
    			}
			}
			if v.T == "str" {
    			v.Value = conditions.ReplaceBracedVariablesInString(v.Value)
    			if inString {
        			fmt.Printf("[73402] - At line %d - %s is an opened string but never closed.\n", f.Line, v.Value)
        			os.Exit(1)
    			}
			}
			if v.T == "" {
				fmt.Printf("[73402] - At line %d - %s is not accorded to any working type (int/nil/bool/str).\n", f.Line, v.Value)
				os.Exit(1)
			}

			if i != len(f.Listed_args)-1 {
    			r = append(r, types.Arg{})
			}
			inString = false

		default:
			arg := &r[len(r)-1]
			arg.Value+=string(char)
		}
	}
	for len(r) > 0 && strings.TrimSpace(r[len(r)-1].Value) == "" {
    	r = r[:len(r)-1]
	}
	if len(r) == 0 {
	    return r
	}

	v := &r[len(r)-1]
	v.Value = strings.TrimSpace(v.Value)
	if !isBalancedParentheses(v.Value) {
    	for strings.HasSuffix(v.Value, ")") && !strings.HasPrefix(v.Value, "\"") {
        	v.Value = strings.TrimSuffix(v.Value, ")")
        	v.Value = strings.TrimSpace(v.Value)
    	}
	}
	if v.Value == "" {
	    r = r[:len(r)-1]
    	return r
	}
	if v.T != "str" {
    	value, t := resolvers.ResolveValue(v.Value)
    	v.Value = value
    	if t != "" {
    	    v.T = t
    	}
	}

	if v.T != "str" && !inString {
	    v.Value = strings.TrimSpace(v.Value)
	    v.Value = resolvers.ReplaceVariablesInExpr(v.Value)

 	    if conditions.IsCondition(v.Value) {
	        result := conditions.EvaluateConditions(v.Value)
	        v.Value = fmt.Sprintf("%v", result)
	        v.T = "bool"
 	    } else if util.LooksLikeCalcul(v.Value) {
 	        v.Value = resolvers.ResolveCalculs(v.Value)
 	        v.T = "int"
 	    } else if util.IsNumber(v.Value) {
 	        v.T = "int"
 	    } else if v.Value == "nil" {
        	v.T = "nil"
    	} else if util.IsBoolean(v.Value) {
    	    v.T = "bool"
    	}
	}
	if v.T == "str" {
    	v.Value = conditions.ReplaceBracedVariablesInString(v.Value)
    	if inString {
    	    fmt.Printf("[73402] - At line %d - %s is an opened string but never closed.\n", f.Line, v.Value)
    	    os.Exit(1)
    	}
	}
	if v.T == "" {
	    fmt.Printf("[73402] - At line %d - %s is not accorded to any working type (int/nil/bool/str).\n", f.Line, v.Value)
	    os.Exit(1)
	}
	return r
}

func isBalancedParentheses(s string) bool {
    count := 0
    for _, c := range s {
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
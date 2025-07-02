package conditions

import (
	"fmt"
	"memory/src/core/resolvers"
	"memory/src/util"
	"memory/src/vars"
	"strings"
)

func ReplaceBracedVariablesInString(s string) string {
	variables := vars.GetAvailableVariables()
	var result strings.Builder
	inBraces := false
	varName := strings.Builder{}
	escaped := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		if escaped {
			result.WriteByte(c)
			escaped = false
			continue
		}
		if c == '\\' {
			escaped = true
			continue
		}
		if inBraces {
			if c == '}' {
				name := varName.String()
				trimmed := strings.TrimSpace(name)
				if val, ok := variables[trimmed]; ok {
					result.WriteString(val.Value)
				} else {
					if IsCondition(trimmed) {
						b := EvaluateConditions(trimmed)
						result.WriteString(fmt.Sprintf("%v", b))
					} else if util.LooksLikeCalcul(trimmed) || util.IsNumber(trimmed) {
						result.WriteString(resolvers.ResolveCalculs(trimmed))
					} else {
						result.WriteString("{" + name + "}")
					}
				}
				varName.Reset()
				inBraces = false
			} else {
				varName.WriteByte(c)
			}
		} else {
			if c == '{' {
				inBraces = true
			} else {
				result.WriteByte(c)
			}
		}
	}
	if inBraces {
		result.WriteString("{" + varName.String())
	}
	return result.String()
}
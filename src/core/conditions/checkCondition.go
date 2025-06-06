package conditions;

func EvaluateCondition(expression string) bool {
	var r bool = false
	for _, char := range expression {

		switch string(char) {

		case ">":

		case "<":

		case "=":

		case "!":

		case "&":

		case "|":

		default:

		}
	}

	// Evaluates once parsed and much easily comprehensible by the system.

	return r
}
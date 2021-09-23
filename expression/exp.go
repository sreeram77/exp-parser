package expression

import (
	"strconv"
	"strings"
)

func ParseExp(exp string, json map[string]interface{}) bool {
	if strings.ContainsRune(exp, '(') {
		// remove bracket
		return false
		// call recursive

	}

	if strings.Contains(exp, evalOr) {
		orExps := strings.Split(exp, evalOr)

		orRes := false

		for i := range orExps {
			if strings.Contains(orExps[i], evalAnd) {
				andExps := strings.Split(orExps[i], evalAnd)

				res := evaluateSubExpression(andExps[0], json)

				for i := range andExps {
					res = res && evaluateSubExpression(andExps[i], json)
				}

				orExps[i] = strconv.FormatBool(res)
			}

			if i == 0 {
				orRes = evaluateSubExpression(orExps[0], json)
			} else {
				orRes = orRes || evaluateSubExpression(orExps[i], json)
			}
		}

		return orRes
	} else if strings.Contains(exp, evalAnd) {
		andExps := strings.Split(exp, evalAnd)

		res := evaluateSubExpression(andExps[0], json)
		for i := range andExps {
			res = res && evaluateSubExpression(andExps[i], json)
		}

		return res
	}

	return evaluateSubExpression(exp, json)
}

func evaluateSubExpression(subExp string, j map[string]interface{}) bool {
	if subExp == evalTrue {
		return true
	}

	if subExp == evalFalse {
		return false
	}

	switch Type(subExp) {
	case EqualTo:
		var e ExpEqual
		return e.Evaluate(subExp, j)
	case Exists:
		var e ExpExists
		return e.Evaluate(subExp, j)
	case NotExists:
		var e ExpNotExists
		return e.Evaluate(subExp, j)
	}

	return false
}

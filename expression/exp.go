package expression

import (
	"strconv"
	"strings"
)

func ParseExp(exp string, json map[string]interface{}) bool {
	if strings.ContainsRune(exp, evalOpenBracket) {
		for strings.ContainsRune(exp, evalOpenBracket) {
			start := strings.LastIndexByte(exp, evalOpenBracket)
			end := strings.IndexByte(exp, evalClosingBracket)

			pre := exp[:start]
			cur := exp[start+1 : end]
			pos := exp[end+1:]

			cur = strconv.FormatBool(ParseExp(cur, json))
			exp = pre + cur + pos
		}
	}

	if strings.Contains(exp, evalOr) {
		return parseOR(exp, json)
	} else if strings.Contains(exp, evalAnd) {
		return parseAND(exp, json)
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

func parseOR(exp string, json map[string]interface{}) bool {
	orExps := strings.Split(exp, evalOr)
	orRes := false

	for i := range orExps {
		if strings.Contains(orExps[i], evalAnd) {
			orExps[i] = strconv.FormatBool(parseAND(orExps[i], json))
		}

		if i == 0 {
			orRes = evaluateSubExpression(orExps[0], json)
		} else {
			orRes = orRes || evaluateSubExpression(orExps[i], json)
		}
	}

	return orRes
}

func parseAND(exp string, json map[string]interface{}) bool {
	andExps := strings.Split(exp, evalAnd)

	res := evaluateSubExpression(andExps[0], json)
	for i := range andExps {
		res = res && evaluateSubExpression(andExps[i], json)
	}

	return res
}

package expression

import (
	"strconv"
	"strings"
)

// ParseExp parses an expression and evaluates against JSON input.
func ParseExp(exp string, json map[string]interface{}) bool {
	// Check for brackets and evaluate first
	if strings.ContainsRune(exp, evalOpenBracket) {
		for strings.ContainsRune(exp, evalOpenBracket) {
			// Get indices of inner most brackets
			start := strings.LastIndexByte(exp, evalOpenBracket)
			end := strings.IndexByte(exp, evalClosingBracket)

			// pre contains everything before inner opening bracket
			pre := exp[:start]
			// cur contains sub expression inside inner bracket
			cur := exp[start+1 : end]
			// pos contains everything after inner closing bracket
			pos := exp[end+1:]

			// Parse sub expression and store result as string
			cur = strconv.FormatBool(ParseExp(cur, json))
			// Join string to simplified expression
			exp = pre + cur + pos
		}
	}

	// Evaluate OR and AND
	if strings.Contains(exp, evalOr) {
		return parseOR(exp, json)
	} else if strings.Contains(exp, evalAnd) {
		return parseAND(exp, json)
	}

	// Evaluate simple sub expression
	return evaluateSubExpression(exp, json)
}

// evaluateSubExpression evaluates sub expressions such as "==", "EXISTS"
// "NOT EXISTS" etc.
func evaluateSubExpression(subExp string, j map[string]interface{}) bool {
	if subExp == evalTrue {
		return true
	}

	if subExp == evalFalse {
		return false
	}

	if strings.Contains(subExp, evalEqual) {
		return evaluateEqual(subExp, j)
	}

	if strings.Contains(subExp, evalNotExists) {
		return evaluateNotExists(subExp, j)
	}

	if strings.Contains(subExp, evalExists) {
		return evaluateExists(subExp, j)
	}

	return false
}

// Parse expression container OR and AND
func parseOR(exp string, json map[string]interface{}) bool {
	// Split into sub expressions
	orExps := strings.Split(exp, evalOr)
	orRes := false

	// Iterate over each sub expression
	for i := range orExps {
		// If sub exp contains AND, execute it
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

// Parse expression with only AND
func parseAND(exp string, json map[string]interface{}) bool {
	// Split expression by AND
	andExps := strings.Split(exp, evalAnd)

	res := evaluateSubExpression(andExps[0], json)
	for i := range andExps {
		res = res && evaluateSubExpression(andExps[i], json)
	}

	return res
}

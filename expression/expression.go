package expression

import (
	"strconv"
	"strings"
)

const (
	evalEqual          = " == "
	evalNotExists      = "NOT EXISTS "
	evalExists         = "EXISTS "
	evalVar            = "$"
	evalStrVal         = "'"
	evalNested         = "."
	evalAnd            = " AND "
	evalOr             = " OR "
	evalTrue           = "true"
	evalFalse          = "false"
	evalOpenBracket    = '('
	evalClosingBracket = ')'
)

func evaluateEqual(s string, json map[string]interface{}) bool {
	ops := strings.Split(s, evalEqual)
	key := strings.Trim(ops[0], evalVar)
	value := ops[1]

	return checkEqual(key, value, json)
}

func checkEqual(key, value string, json interface{}) bool {
	if strings.Contains(key, evalNested) {
		nestedKey := strings.Split(key, evalNested)
		switch t := json.(type) {
		case map[interface{}]interface{}:
			val, ok := t[nestedKey[0]]
			if !ok {
				return false
			}

			return checkEqual(nestedKey[1], value, val)
		case map[string]interface{}:
			val, ok := t[nestedKey[0]]
			if !ok {
				return false
			}

			return checkEqual(nestedKey[1], value, val)
		}
	}

	switch t := json.(type) {
	case map[interface{}]interface{}:
		val, ok := t[key]
		if !ok {
			return false
		}

		return compare(value, val)
	case map[string]interface{}:
		val, ok := t[key]
		if !ok {
			return false
		}

		return compare(value, val)
	}

	return false
}

func compare(op string, value interface{}) bool {
	switch v := value.(type) {
	case string:
		v = strings.Trim(v, evalStrVal)
		op = strings.Trim(op, evalStrVal)
		return v == op
	case int:
		val, err := strconv.Atoi(op)
		if err != nil {
			return false
		}
		return v == val
	case bool:
		val, err := strconv.ParseBool(op)
		if err != nil {
			return false
		}
		return v == val
	case float64:
		val, err := strconv.ParseFloat(op, 64)
		if err != nil {
			return false
		}
		return v == val
	}

	return false
}

func evaluateNotExists(s string, json map[string]interface{}) bool {
	ops := strings.Split(s, evalNotExists)
	op := strings.Trim(ops[1], evalVar)

	return checkNotExists(op, json)
}

func checkNotExists(key string, json interface{}) bool {
	if strings.Contains(key, evalNested) {
		nestedKey := strings.Split(key, evalNested)
		switch v := json.(type) {
		case map[interface{}]interface{}:
			val, ok := v[nestedKey[0]]
			if !ok {
				return true
			}

			return checkExists(nestedKey[1], val)
		case map[string]interface{}:
			val, ok := v[nestedKey[0]]
			if !ok {
				return true
			}

			return checkExists(nestedKey[1], val)
		}
	}

	switch v := json.(type) {
	case map[interface{}]interface{}:
		_, ok := v[key]
		if !ok {
			return true
		}

		return false
	case map[string]interface{}:
		_, ok := v[key]
		if !ok {
			return true
		}

		return false
	}

	return true
}

func evaluateExists(s string, json map[string]interface{}) bool {
	ops := strings.Split(s, evalExists)
	op := strings.Trim(ops[1], evalVar)

	return checkExists(op, json)
}

func checkExists(key string, json interface{}) bool {
	if strings.Contains(key, evalNested) {
		nestedKey := strings.Split(key, evalNested)
		switch v := json.(type) {
		case map[interface{}]interface{}:
			val, ok := v[nestedKey[0]]
			if !ok {
				return false
			}

			return checkExists(nestedKey[1], val)
		case map[string]interface{}:
			val, ok := v[nestedKey[0]]
			if !ok {
				return false
			}

			return checkExists(nestedKey[1], val)
		}
	}

	switch v := json.(type) {
	case map[interface{}]interface{}:
		_, ok := v[key]
		if !ok {
			return false
		}

		return true
	case map[string]interface{}:
		_, ok := v[key]
		if !ok {
			return false
		}

		return true
	}

	return false
}

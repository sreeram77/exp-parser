package expression

import (
	"strconv"
	"strings"
)

type Exp int

type Expression interface {
	Evaluate(exp string, json map[string]interface{}) bool
}

const (
	Unknown   Exp = iota
	EqualTo   Exp = 1
	Exists    Exp = 2
	NotExists Exp = 3
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

func Type(s string) Exp {
	if strings.Contains(s, evalEqual) {
		return EqualTo
	}

	if strings.Contains(s, evalNotExists) {
		return NotExists
	}

	if strings.Contains(s, evalExists) {
		return Exists
	}

	return Unknown
}

type ExpEqual struct{}

func (e ExpEqual) Evaluate(s string, json map[string]interface{}) bool {
	ops := strings.Split(s, evalEqual)

	op := strings.Trim(ops[0], evalVar)
	oper := ops[1]

	if strings.Contains(op, evalNested) {
		nested := strings.Split(op, evalNested)
		val, ok := json[nested[0]]
		if !ok {
			return false
		}

		oper = getNestedValue(nested[1], val)
	}

	return compare(oper, ops[1])
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

func getNestedValue(key string, json interface{}) string {
	if strings.Contains(key, evalNested) {
		nestedKey := strings.Split(key, evalNested)
		switch v := json.(type) {
		case map[interface{}]interface{}:
			val, ok := v[key]
			if !ok {
				return ""
			}

			return getNestedValue(nestedKey[1], val)
		}
		return ""
	}

	switch v := json.(type) {
	case map[interface{}]interface{}:
		val, ok := v[key]
		if !ok {
			return ""
		}

		return val.(string)
	}

	return ""
}

type ExpNotExists struct{}

func (e ExpNotExists) Evaluate(s string, json map[string]interface{}) bool {
	ops := strings.Split(s, evalNotExists)
	op := strings.Trim(ops[1], evalVar)

	return checkNotExists(op, json)
}

func checkNotExists(key string, json interface{}) bool {
	if strings.Contains(key, evalNested) {
		nestedKey := strings.Split(key, evalNested)
		switch v := json.(type) {
		case map[interface{}]interface{}:
			val, ok := v[key]
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

type ExpExists struct{}

func (e ExpExists) Evaluate(s string, json map[string]interface{}) bool {
	ops := strings.Split(s, evalExists)
	op := strings.Trim(ops[1], evalVar)

	return checkExists(op, json)
}

func checkExists(key string, json interface{}) bool {
	if strings.Contains(key, evalNested) {
		nestedKey := strings.Split(key, evalNested)
		switch v := json.(type) {
		case map[interface{}]interface{}:
			val, ok := v[key]
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

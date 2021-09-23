package parser

import (
	"github.com/sreeram77/exp-parser/expression"
	"github.com/sreeram77/exp-parser/model"
)

type parser struct {
}

func New() Parser {
	return parser{}
}

func (p parser) Parse(data model.TestCases) (model.TestCases, error) {
	for i := range data.Testcase {
		// Parse Expression
		res := expression.ParseExp(data.Testcase[i].Expression, data.Testcase[i].Json)
		data.Testcase[i].ActualOutput = res
	}

	return model.TestCases{}, nil
}

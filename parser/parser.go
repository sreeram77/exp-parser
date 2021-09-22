package parser

import "github.com/sreeram77/exp-parser/model"

type parser struct {
}

func New() Parser {
	return parser{}
}

func (p parser) Parse(data []model.TestCase) ([]model.TestCase, error) {
	for i := range data {
		// Parse Expression

		// Evaluate Expression

	}

	return []model.TestCase{}, nil
}

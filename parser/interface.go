package parser

import "github.com/sreeram77/exp-parser/model"

type Parser interface {
	Parse(data model.TestCases) (model.TestCases, error)
}

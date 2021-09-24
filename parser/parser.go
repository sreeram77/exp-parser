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

type response struct {
	output bool
	index  int
}

func (p parser) Parse(data model.TestCases) (model.TestCases, error) {
	respCh := make(chan response)

	defer func() {
		close(respCh)
	}()

	for i, v := range data.Testcase {
		go func(in int, tc model.TestCase) {
			result := expression.ParseExp(tc.Expression, tc.Json)
			r := response{
				output: result,
				index:  in,
			}
			respCh <- r
		}(i, v)
	}

	tCases := make([]model.TestCase, len(data.Testcase))

	for i := 0; i < len(data.Testcase); i++ {
		select {
		case out := <-respCh:
			tCase := data.Testcase[out.index]
			tCase.ActualOutput = out.output
			tCases[out.index] = tCase
		}
	}

	return model.TestCases{
		Testcase: tCases,
	}, nil
}

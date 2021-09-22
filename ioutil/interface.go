package ioutil

import "github.com/sreeram77/exp-parser/model"

type IOUtil interface {
	Read() (model.TestCases, error)
	Write(data model.TestCases) error
}

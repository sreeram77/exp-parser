package ioutil

import "github.com/sreeram77/exp-parser/model"

type IOUtil interface {
	Read() ([]model.TestCase, error)
	Write(data []model.TestCase) error
}

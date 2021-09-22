package ioutil

import (
	"github.com/sreeram77/exp-parser/model"
	"gopkg.in/yaml.v2"
)

type ioutil struct{}

func New() IOUtil {
	return ioutil{}
}

func (i ioutil) Read() ([]model.TestCase, error) {
	var data = `
testcases:
 - expression: "$color == 'red'"
   json: {"color":"red","size":10,"cost":100.0,"mattress":{"name":"king"},"big":true,"legs":[{"length":4}]}
   expected_output: true
 - expression: "$mattress.name == 'king' AND $cost == 100.0"
   json: {"color":"red","size":10,"cost":100.0,"mattress":{"name":"king"},"big":true,"legs":[{"length":4}]}
   expected_output: true`

	var t model.TestCases
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return []model.TestCase{}, err
	}

	return t.Testcase, nil
}

func (i ioutil) Write([]model.TestCase) error {
	panic("implment me")
}

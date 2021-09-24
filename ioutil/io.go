package ioutil

import (
	"fmt"
	"os"

	"github.com/sreeram77/exp-parser/model"
	"gopkg.in/yaml.v2"
)

type ioutil struct {
	ipFilePath string
	opFilePath string
}

func New(ip, op string) IOUtil {
	return ioutil{ipFilePath: ip, opFilePath: op}
}

func (i ioutil) Read() (model.TestCases, error) {
	data, err := os.ReadFile(i.ipFilePath)
	if err != nil {
		return model.TestCases{}, err
	}

	var dat = `
testcases:
 - expression: "$color == 'red'"
   json: {"color":"red","size":10,"cost":100.0,"mattress":{"name":"king"},"big":true,"legs":[{"length":4}]}
   expected_output: true
 - expression: "$mattress.name == 'king' AND $cost == 100.0"
   json: {"color":"red","size":10,"cost":100.0,"mattress":{"name":"king"},"big":true,"legs":[{"length":4}]}
   expected_output: true
 - expression: "NOT EXISTS $color"
   json: {"color":"red","size":10,"cost":100.0,"mattress":{"name":"king"},"big":true,"legs":[{"length":4}]}
   expected_output: false
 - expression: "( $cost == 100.0 AND ( $mattress.big == false ) ) OR $size == 100"
   json: {"color":"red","size":10,"cost":100.0,"mattress":{"name":"king"},"big":true,"legs":[{"length":4}]}
   expected_output: false`

	fmt.Println(len(dat))

	var t model.TestCases
	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		return model.TestCases{}, err
	}

	return t, nil
}

func (i ioutil) Write(t model.TestCases) error {
	data, err := yaml.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(i.opFilePath, data, 0777)
}

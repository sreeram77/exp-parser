package ioutil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

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
	var result []bool

	for _, v := range t.Testcase {
		result = append(result, v.ActualOutput)
	}

	lines, err := linesFromFile(i.ipFilePath)
	if err != nil {
		return err
	}

	lines = append(lines, "")

	content := ""
	j := 0

	for i, line := range lines {
		if i > 3 && (i-1)%3 == 0 {
			res := fmt.Sprintf("   actual_output: %s\n", strconv.FormatBool(result[j]))
			content += res
			j++
		}
		content += line
		content += "\n"
	}

	return os.WriteFile(i.opFilePath, []byte(content), 0644)
}

func linesFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

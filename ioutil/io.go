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
	content    [][]byte
}

func New(ip, op string) IOUtil {
	return &ioutil{ipFilePath: ip, opFilePath: op}
}

func (io *ioutil) Read() (model.TestCases, error) {
	fmt.Println("reading from file ", io.ipFilePath)
	var err error

	io.content, err = linesFromFile(io.ipFilePath)
	if err != nil {
		return model.TestCases{}, err
	}

	var ts []model.TestCase

	for i := 1; i < len(io.content); i = i + 3 {
		var tmp []byte
		var t model.TestCases

		tmp = append(tmp, io.content[0]...)
		tmp = append(tmp, byte('\r'), byte('\n'))
		tmp = append(tmp, io.content[i]...)
		tmp = append(tmp, byte('\r'), byte('\n'))
		tmp = append(tmp, io.content[i+1]...)

		err := yaml.Unmarshal(tmp, &t)
		if err != nil {
			return model.TestCases{}, err
		}

		ts = append(ts, t.Testcase...)
	}

	return model.TestCases{Testcase: ts}, nil
}

func (io *ioutil) Write(t model.TestCases) error {
	fmt.Println("writing to file ", io.opFilePath)
	var result []bool

	for _, v := range t.Testcase {
		result = append(result, v.ActualOutput)
	}

	io.content = append(io.content, []byte("\r\n"))

	content := ""
	j := 0

	for i, line := range io.content {
		if i > 3 && (i-1)%3 == 0 {
			res := fmt.Sprintf("   actual_output: %s\n", strconv.FormatBool(result[j]))
			content += res
			j++
		}
		content += string(line)
		content += "\n"
	}

	fmt.Println("output : \n", string(content))

	return os.WriteFile(io.opFilePath, []byte(content), 0644)
}

func linesFromFile(path string) ([][]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([][]byte, error) {
	var lines [][]byte
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

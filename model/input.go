package model

type TestCase struct {
	Expression     string                 `yaml:"expression"`
	Json           map[string]interface{} `yaml:"json"`
	ExpectedOutput bool                   `yaml:"expected_output"`
}

type TestCases struct {
	Testcase []TestCase `yaml:"testcases"`
}

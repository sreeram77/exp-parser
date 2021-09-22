package parser

type parser struct {
}

func New() Parser {
	return parser{}
}

func (p parser) Parse(data []byte) ([]byte, error) {
	panic("implment me")
}

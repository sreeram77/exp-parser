package parser

type Parser interface {
	Parse(data []byte) ([]byte, error)
}

package ioutil

type IOUtil interface {
	Read() ([]byte, error)
	Write(data []byte) error
}

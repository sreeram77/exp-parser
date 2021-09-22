package ioutil

type ioutil struct{}

func New() IOUtil {
	return ioutil{}
}

func (i ioutil) Read() ([]byte, error) {
	panic("implment me")
}

func (i ioutil) Write(data []byte) error {
	panic("implment me")
}

package reader

type Reader struct {
	ID   int
	Name string
}

func NewReader(id int, name string) *Reader {
	return &Reader{
		ID:   id,
		Name: name,
	}
}

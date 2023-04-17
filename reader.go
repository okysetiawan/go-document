package document

// Reader is interface implemented to read document
type Reader interface {
	Bytes() ([]byte, error)
}

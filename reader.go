package document

// Reader is interface implemented to read document
type Reader interface {
	Scan(dest interface{}) error
}

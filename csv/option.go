package csv

// WriterOption is function to update option to initialize csv.NewBuilder() and csv.NewWriterByType()
type WriterOption func(csv *writer)

func WithWriterCommaDelimiter() WriterOption { return func(csv *writer) { csv.delimiter = Comma } }

func WithWriterSemicolonDelimiter() WriterOption {
	return func(csv *writer) { csv.delimiter = Semicolon }
}

func WithWriterHeader(headers []string) WriterOption {
	return func(csv *writer) { csv.header = headers }
}

// ReaderOption is function to update option to initialize csv.NewReader()
type ReaderOption func(csv *reader)

func WithReaderBytes(bytes []byte) ReaderOption {
	return func(csv *reader) { csv.raw = bytes }
}

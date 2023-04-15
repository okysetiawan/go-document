package csv

type WriterOption func(csv *writer)

func WithWriterCommaDelimiter() WriterOption { return func(csv *writer) { csv.delimiter = Comma } }
func WithWriterSemicolonDelimiter() WriterOption {
	return func(csv *writer) { csv.delimiter = Semicolon }
}
func WithWriterHeader(headers []string) WriterOption {
	return func(csv *writer) { csv.header = headers }
}

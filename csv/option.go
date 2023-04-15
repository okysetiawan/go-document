package csv

type WriterOption func(csv *writer)

func WithCommaDelimiter() WriterOption         { return func(csv *writer) { csv.delimiter = Comma } }
func WithSemicolonDelimiter() WriterOption     { return func(csv *writer) { csv.delimiter = Semicolon } }
func WithHeader(headers []string) WriterOption { return func(csv *writer) { csv.header = headers } }

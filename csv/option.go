package csv

type Option func(csv *writer)

func WithCommaDelimiter() Option     { return func(csv *writer) { csv.delimiter = Comma } }
func WithSemicolonDelimiter() Option { return func(csv *writer) { csv.delimiter = Semicolon } }

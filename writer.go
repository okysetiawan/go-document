package document

import "bytes"

type (
	Builder interface {
		CreateSheet(sheetName string, headers ...string) Builder
		AddRow(cells ...any) Builder
		Writer
	}

	Writer interface {
		Buffer() (*bytes.Buffer, error)
		Save(path string) (err error)
	}
)

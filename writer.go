package document

import "bytes"

type (
	// Builder is interface implemented by an object that can handle build Document before process write file
	Builder interface {
		CreateSheet(sheetName string, headers ...string) Builder
		AddRow(cells ...any) Builder
		Writer
	}

	// Writer is interface implemented by an object that can handle Write Document.
	Writer interface {
		Buffer() (*bytes.Buffer, error)
		Save(path string) (err error)
	}
)

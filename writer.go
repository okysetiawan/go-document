package document

import "bytes"

type (
	Writer interface {
		CreateSheet(sheetName string, headers ...string) Writer
		AddRow(cells []interface{}) Writer
		Buffer() (*bytes.Buffer, error)
		Save(path string) (err error)
	}
)

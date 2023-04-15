package csv

import (
	"github.com/okysetiawan/go-document"
	"github.com/okysetiawan/go-document/internal"
)

func NewByType(types any, filename string, options ...WriterOption) (document.Writer, error) {
	var (
		header []string
		rows   [][]string
		err    error
	)

	w := New(options...)

	header, err = internal.GetHeaderFromAny(types, "document")
	if err != nil {
		return nil, err
	}

	w.CreateSheet(filename, header...)

	rows, err = internal.GetRowsFromAny(types)
	if err != nil {
		return nil, err
	}

	for i := range rows {
		w.AddRow(rows[i])
	}

	return w, nil
}

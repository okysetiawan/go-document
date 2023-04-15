package csv

import (
	"bytes"
	"encoding/csv"
	"github.com/okysetiawan/go-document"
	"github.com/okysetiawan/go-document/errors"
	"github.com/okysetiawan/go-document/internal"
	"io"
	"os"
	"strings"
)

type (
	writer struct {
		// fileName attribute is document file name
		fileName string
		// store header list
		// used to validate rows length
		header []string
		// store rows data
		// used to temporary store rows before Save() or Buffer()
		rows [][]string
		// store custom delimiter
		delimiter rune
	}
)

func New(opts ...Option) document.Builder {
	builder := &writer{
		delimiter: Semicolon,
		header:    make([]string, 0),
		rows:      make([][]string, 0),
	}

	for i := range opts {
		opts[i](builder)
	}

	return builder
}

// CreateSheet will store file name and create headers if exists, it will replace existing file name and header if its already exists.
// CSV will only have 1 sheet.
func (w *writer) CreateSheet(sheetName string, headers ...string) document.Builder {
	w.fileName = sheetName
	w.header = headers
	return w
}

// AddRow will store rows on impl.rows temporary before Save() or Buffer()
func (w *writer) AddRow(cells ...any) document.Builder {
	w.rows = append(w.rows, internal.SliceAnyToString(cells))
	return w
}

func (w *writer) write(wr io.Writer) error {
	var (
		err error
	)

	csvWriter := csv.NewWriter(wr)
	csvWriter.Comma = w.delimiter

	if err = csvWriter.Write(w.header); err != nil {
		return errors.WithCode(err, errors.FailedWriteHeader)
	}

	if err = csvWriter.WriteAll(w.rows); err != nil {
		return errors.WithCode(err, errors.FailedWriteRows)
	}

	return nil
}

// Buffer will write csv writer and return data on *bytes.Buffer format
func (w *writer) Buffer() (*bytes.Buffer, error) {
	buff := &bytes.Buffer{}
	err := w.write(buff)
	if err != nil {
		return nil, err
	}

	return buff, nil
}

func (w *writer) Save(folderPath string) error {
	path := strings.ReplaceAll(folderPath+"/"+w.fileName, "//", "/")
	file, err := os.Create(path)
	if err != nil {
		return errors.WithCode(err, errors.FailedWriteToDevice)
	}
	defer file.Close()

	err = w.write(file)
	if err != nil {
		return err
	}

	return nil
}

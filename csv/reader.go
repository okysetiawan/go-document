package csv

import (
	"bytes"
	"encoding/csv"
	"github.com/okysetiawan/go-document"
	"github.com/okysetiawan/go-document/errors"
	"io"
	"os"
)

// reader is object represent to read csv file
type reader struct {
	// raw containing document raw
	raw []byte
	csv [][]string
}

func (r reader) Bytes() ([]byte, error) {
	return r.raw, nil
}

// NewReaderIO to initialize document.Reader with io.Reader types
func NewReaderIO(ioReader io.Reader, options ...ReaderOption) (document.Reader, error) {

	bytes, err := io.ReadAll(ioReader)
	if err != nil {
		return nil, errors.WithCode(err, errors.CodeReadReader)
	}

	if bytes == nil || len(bytes) == 0 {
		return nil, errors.New("io.Reader is nil or empty").WithCode(errors.CodeReadEmpty)
	}

	options = append(options, WithReaderBytes(bytes))
	r, err := newReader(options...)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// NewReaderOpenFile to initialize document.Reader with open file from system
func NewReaderOpenFile(path string, options ...ReaderOption) (document.Reader, error) {

	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.WithCode(err, errors.CodeReadFileNotExist)
		}
		return nil, errors.WithCode(err, errors.CodeReadFileFailed)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.WithCode(err, errors.CodeReadFileFailed)
	}
	defer file.Close()

	return NewReaderIO(file, options...)
}

func newReader(options ...ReaderOption) (*reader, error) {
	var (
		r   = &reader{}
		err error
	)

	for i := range options {
		options[i](r)
	}

	read := csv.NewReader(bytes.NewReader(r.raw))
	r.csv, err = read.ReadAll()
	if err != nil {
		return nil, err
	}

	return r, nil
}

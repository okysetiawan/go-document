package csv

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"github.com/okysetiawan/go-document"
	"github.com/okysetiawan/go-document/errors"
	"io"
	"os"
)

// reader is object represent to read csv file
type reader struct {
	// parser is attribute to Unmarshal content into destination object
	parser jsoniter.API
	// content containing document content
	content []byte
}

func (r *reader) Scan(dest interface{}) error {

	err := r.parser.Unmarshal(r.content, dest)
	if err != nil {
		return errors.WithCode(err, errors.CodeReadUnmarshal)
	}

	return nil
}

// NewReaderIO to initialize document.Reader with io.Reader types
func NewReaderIO(ioReader io.Reader, options ...ReaderOption) (document.Reader, error) {

	bytes, err := io.ReadAll(ioReader)
	if err != nil {
		return nil, errors.WithCode(err, errors.CodeReadReader)
	}

	r := newReader(options...)
	r.content = bytes
	return r, nil
}

// NewReaderBuffer to initialize document.Reader with *bytes.Buffer types
func NewReaderBuffer(buff *bytes.Buffer, options ...ReaderOption) (document.Reader, error) {

	if buff == nil || buff.Len() == 0 {
		return nil, errors.New("buff is nil or empty").WithCode(errors.CodeReadBufferEmpty)
	}

	r := newReader(options...)
	r.content = buff.Bytes()
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

func newReader(options ...ReaderOption) *reader {
	r := &reader{parser: jsoniter.Config{TagKey: "json"}.Froze()}

	for i := range options {
		options[i](r)
	}

	return r
}

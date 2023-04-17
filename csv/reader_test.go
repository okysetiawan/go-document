package csv

import (
	"bytes"
	"github.com/okysetiawan/go-document/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReaderIO(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		buff := &bytes.Buffer{}

		r, err := NewReaderIO(buff)
		assert.Error(t, err)
		assert.Nil(t, r)
		assert.True(t, errors.IsCode(err, errors.CodeReadEmpty))
	})

	t.Run("Read Buffer", func(t *testing.T) {
		data := "Name;Birth Date\nAlex;2000-10-23\nJustin;1999-03-30\n"
		buff := &bytes.Buffer{}
		buff.WriteString(data)
		defer buff.Reset()

		r, err := NewReaderIO(buff)
		assert.NoError(t, err)
		assert.NotNil(t, r)

		var raw []byte
		raw, err = r.Bytes()
		assert.NoError(t, err)
		assert.Len(t, raw, 50)
		assert.EqualValues(t, string(raw), data)
		buff.Reset()
	})

	t.Run("Read File", func(t *testing.T) {
		r, err := NewReaderOpenFile("../tests/test_reader.csv")
		assert.NoError(t, err)
		assert.NotNil(t, r)

		var raw []byte
		raw, err = r.Bytes()
		assert.NoError(t, err)
		assert.Len(t, raw, 47)
		assert.EqualValues(t, string(raw), "Number;Name\n1;Ashish Verma\n2;Justin\n3;Christine")
	})
}

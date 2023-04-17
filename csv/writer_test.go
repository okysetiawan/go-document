package csv

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestWriter_Buffer(t *testing.T) {
	w := NewBuilder(WithWriterSemicolonDelimiter())

	var (
		headers = []string{"Number", "Name"}
		rows    = [][]any{
			{1, "Alex"},
			{2, "Jonathan"},
			{3, "Abby"},
		}
		expected = "Number;Name\n1;Alex\n2;Jonathan\n3;Abby\n4;Evelyn\n"
	)

	w.CreateSheet("test.csv", headers...)

	for i := range rows {
		w.AddRow(rows[i]...)
	}

	w = w.AddRow(4, "Evelyn")

	buff, err := w.Buffer()
	assert.NoError(t, err)
	assert.EqualValues(t, expected, buff.String())
}

func TestWriter_Save(t *testing.T) {
	w := NewBuilder(WithWriterSemicolonDelimiter())

	var (
		headers = []string{"Number", "Name"}
		rows    = [][]any{
			{1, "Alex"},
			{2, "Jonathan"},
			{3, "Abby"},
		}
		expected = "Number;Name\n1;Alex\n2;Jonathan\n3;Abby\n4;Evelyn\n"
		path     = "../tests/writer.csv"
	)

	w.CreateSheet("writer.csv", headers...)

	for i := range rows {
		w.AddRow(rows[i]...)
	}

	w = w.AddRow(4, "Evelyn")

	// do save file
	err := w.Save("../tests")
	defer os.Remove(path)
	assert.NoError(t, err)

	// open file
	file, err := os.Open(path)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	defer file.Close()

	// read actual raw
	actualContent, err := io.ReadAll(file)
	assert.NoError(t, err)
	assert.EqualValues(t, expected, actualContent)

}

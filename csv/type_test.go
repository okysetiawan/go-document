package csv

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
	"time"
)

func TestNewByType(t *testing.T) {
	type Table struct {
		Name       string    `document:"Name"`
		BirthPlace string    `document:"Birth Place"`
		BirthDate  time.Time `document:"Birth Date"`
	}
	table := []Table{
		{
			Name:       "Alexander",
			BirthPlace: "Jakarta",
			BirthDate:  time.Date(2000, 10, 1, 10, 0, 0, 0, time.UTC),
		},
		{
			Name:       "Justin",
			BirthPlace: "Bali",
			BirthDate:  time.Date(1989, 4, 25, 4, 30, 0, 0, time.UTC),
		},
	}

	w, err := NewByType(table, "test_type.csv")
	assert.NoError(t, err)
	assert.NotNil(t, w)

	expected := "Name;Birth Place;Birth Date\n[Alexander Jakarta 2000-10-01 10:00:00 +0000 UTC]\n[Justin Bali 1989-04-25 04:30:00 +0000 UTC]\n"

	// verify Buffer()
	buff, err := w.Buffer()
	assert.NoError(t, err)
	assert.NotNil(t, buff)
	assert.EqualValues(t, expected, buff.String())

	// verify Save()
	path := "../tests/test_type.csv"
	err = w.Save("../tests")
	defer os.Remove(path)
	assert.NoError(t, err)

	// open file
	file, err := os.Open(path)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	defer file.Close()

	// read actual content
	actualContent, err := io.ReadAll(file)
	assert.NoError(t, err)
	assert.EqualValues(t, expected, string(actualContent))
}

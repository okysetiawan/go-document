package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type iMarshaler string

func (i iMarshaler) MarshalText() ([]byte, error) { return []byte(i + " iMarshaler"), nil }

func TestGetHeaderFromAny(t *testing.T) {
	type args struct {
		in  any
		tag string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				in: struct {
					Name       string    `document:"Name"`
					BirthPlace string    `document:"Birth Place"`
					BirthDate  time.Time `document:"Birth Date"`
				}{},
				tag: "document",
			},
			want: []string{
				"Name", "Birth Place", "Birth Date",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHeaderFromAny(tt.args.in, tt.args.tag)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestGetRowsFromAny(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				in: []struct {
					Name       string    `document:"Name"`
					BirthPlace string    `document:"Birth Place"`
					BirthDate  time.Time `document:"Birth Date"`
				}{
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
				},
			},
			want: [][]string{
				{
					"Alexander",
					"Jakarta",
					"2000-10-01 10:00:00 +0000 UTC",
				},
				{
					"Justin",
					"Bali",
					"1989-04-25 04:30:00 +0000 UTC",
				},
			},
			wantErr: false,
		},
		{
			name: "Success Time with Format",
			args: args{
				in: []struct {
					Name       string    `document:"Name"`
					BirthPlace string    `document:"Birth Place"`
					BirthDate  time.Time `document:"Birth Date" format:"2006-01-02"`
				}{
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
				},
			},
			want: [][]string{
				{
					"Alexander",
					"Jakarta",
					"2000-10-01",
				},
				{
					"Justin",
					"Bali",
					"1989-04-25",
				},
			},
			wantErr: false,
		},
		{
			name: "Success Time with json.Marshaler",
			args: args{
				in: []struct {
					Name       string     `document:"Name"`
					BirthPlace string     `document:"Birth Place"`
					BirthDate  time.Time  `document:"Birth Date" format:"2006-01-02"`
					UniqueKey  iMarshaler `document:"Unique Key"`
				}{
					{
						Name:       "Alexander",
						BirthPlace: "Jakarta",
						BirthDate:  time.Date(2000, 10, 1, 10, 0, 0, 0, time.UTC),
						UniqueKey:  "1",
					},
					{
						Name:       "Justin",
						BirthPlace: "Bali",
						BirthDate:  time.Date(1989, 4, 25, 4, 30, 0, 0, time.UTC),
						UniqueKey:  "2",
					},
				},
			},
			want: [][]string{
				{
					"Alexander",
					"Jakarta",
					"2000-10-01",
					"1 iMarshaler",
				},
				{
					"Justin",
					"Bali",
					"1989-04-25",
					"2 iMarshaler",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRowsFromAny(tt.args.in)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}

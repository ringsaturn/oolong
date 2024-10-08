package oolong_test

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
	"testing"

	"github.com/ringsaturn/oolong"
)

func newReaderFromSection(v any) io.Reader {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v)
	return bytes.NewReader(buf.Bytes())
}

func TestNewSection0FromBytes(t *testing.T) {
	tests := []struct {
		name    string
		want    *oolong.Section0Indicator
		wantErr bool
	}{
		{
			name: "Test Section 0",
			want: &oolong.Section0Indicator{
				Identifier:    [4]byte{0x47, 0x52, 0x49, 0x42},
				Reserved:      0,
				Discipline:    0,
				EditionNumber: 2,
				TotalLength:   12,
			},
			wantErr: false,
		},
		{
			name: "Test Section 0",
			want: &oolong.Section0Indicator{
				Identifier:    [4]byte{0x47, 0x52, 0x49, 0x42},
				Reserved:      0,
				Discipline:    0,
				EditionNumber: 121,
				TotalLength:   12,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newReaderFromSection(tt.want)
			got, err := oolong.NewSection0FromBytes(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSection0FromBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSection0FromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

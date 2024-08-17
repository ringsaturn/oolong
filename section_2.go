package oolong

import (
	"encoding/binary"
	"io"
)

// https://codes.ecmwf.int/grib/format/grib2/sections/2/
type Section2LocalUse struct {
	Section2Length  uint32 `byteRange:"1-4"`  // Length of section in octets (nn)
	NumberOfSection uint8  `byteRange:"5"`    // Number of section (2)
	LocalUse        []byte `byteRange:"6-nn"` // Local use
}

func NewSection2FromBytes(r io.Reader) (*Section2LocalUse, error) {
	sec := &Section2LocalUse{}
	if err := binary.Read(r, binary.BigEndian, sec); err != nil {
		return nil, err
	}

	return sec, nil
}

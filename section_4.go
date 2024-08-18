package oolong

import (
	"encoding/binary"
	"io"
)

// https://codes.ecmwf.int/grib/format/grib2/sections/4/
type Section4ProductDefinition struct {
	Section4Length                  uint32 `byteRange:"1-4"`     // Length of section in octets (nn)
	NumberOfSection                 uint8  `byteRange:"5"`       // Number of section (4)
	NV                              uint16 `byteRange:"6-7"`     // Number of coordinate values after template or number of information according to 3D vertical coordinate GRIB2 message
	ProductDefinitionTemplateNumber uint16 `byteRange:"8-9"`     // Product Definition Template Number (see Code Table 4.0)
	ProductDefinitionTemplate       []byte `byteRange:"10-xx"`   // Product Definition Template (see Template 4.X, where X is the Product Definition Template Number given in octets 8-9)
	OptionalList                    []byte `byteRange:"xx+1-nn"` // Optional list of coordinate values or vertical grid information
}

func NewSection4FromBytes(r io.Reader) (*Section4ProductDefinition, error) {
	sec := &Section4ProductDefinition{}
	if err := binary.Read(r, binary.BigEndian, sec); err != nil {
		return nil, err
	}

	return sec, nil
}

package oolong

import (
	"bytes"
	"encoding/binary"
	"io"
)

// GRIB Message Section 0, the Indicator Section.
// This section should always be 16 octets long.
//
// - ECMWF: https://codes.ecmwf.int/grib/format/grib2/sections/0/
// - NOAA: https://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_sect0.shtml
type Section0Indicator struct {
	Identifier    [4]byte    // octec: "1-4", 'GRIB' (Coded according to the International Alphabet Number 5)
	Reserved      uint16     // octec: "5-6", Reserved
	Discipline    Discipline // octec: "7", [Discipline]
	EditionNumber byte       // octec: "8", Edition number - 2 for GRIB2

	// octec: "9-16", Total length of GRIB message in octets (All sections);
	// Please note that a GRIB file can contain multiple GRIB messages.
	TotalLength uint64
}

func (s *Section0Indicator) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, s)
	return buf.Bytes()
}

func NewSection0FromBytes(r io.Reader) (*Section0Indicator, error) {
	sec := &Section0Indicator{}
	if err := binary.Read(r, binary.BigEndian, sec); err != nil {
		return nil, err
	}
	if !bytes.Equal(sec.Identifier[:], _GRIB_HEADER[:]) {
		return sec, ErrInvalidHeader
	}
	if sec.EditionNumber != _GRIB_EDITION {
		return sec, ErrInvalidEditionNumber
	}

	return sec, nil
}

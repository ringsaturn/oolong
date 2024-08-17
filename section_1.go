package oolong

import (
	"encoding/binary"
	"errors"
	"io"
)

var ErrInvalidTablesVersion = errors.New("oolong: invalid tables version") // Error message for invalid tables version

// https://codes.ecmwf.int/grib/format/grib2/sections/1/
type Section1Identification struct {
	Section1Length     uint32 `byteRange:"1-4"`   // Length of section in octets (21 or nn)
	NumberOfSection    uint8  `byteRange:"5"`     // Number of section (1)
	Centre             uint16 `byteRange:"6-7"`   // Identification of originating/generating centre (see Common Code Table C-11). Refer https://github.com/wmo-im/CCT/blob/master/C11.csv
	SubCentre          uint16 `byteRange:"8-9"`   // Identification of originating/generating sub-centre (allocated by originating/generating Centre)
	TablesVersion      uint8  `byteRange:"10"`    // GRIB Master Tables Version Number (see Code Table 1.0 and Note 1)
	LocalTablesVersion uint8  `byteRange:"11"`    // Version number of GRIB Local Tables used to augment Master Tables (see Code Table 1.1 and Note 2)
	SignificanceOfRef  uint8  `byteRange:"12"`    // Significance of Reference Time (see Code Table 1.2)
	Year               uint16 `byteRange:"13-14"` // Year
	Month              uint8  `byteRange:"15"`    // Month
	Day                uint8  `byteRange:"16"`    // Day | Reference time of data
	Hour               uint8  `byteRange:"17"`    // Hour
	Minute             uint8  `byteRange:"18"`    // Minute
	Second             uint8  `byteRange:"19"`    // Second
	ProductionStatus   uint8  `byteRange:"20"`    // Production status of processed data in this GRIB message (see Code Table 1.3)
	TypeOfData         uint8  `byteRange:"21"`    // Type of processed data in this GRIB message (see Code Table 1.4)

	// IdentificationTemplateNumber uint16 `byteRange:"22-23"` // Identification template number (optional, see Code table 1.5)
	// IdentificationTemplate       []byte `byteRange:"24-nn"` // Identification template (optional, see template 1.X, where X is the identification template number given in octets 22-23)
}

func NewSection1FromBytes(r io.Reader) (*Section1Identification, error) {
	sec := &Section1Identification{}
	if err := binary.Read(r, binary.BigEndian, sec); err != nil {
		return nil, err
	}

	// (2) If octet 10 contains 255 then only local tables are in use,
	// the local tables version number (Octet 11) must not be zero nor missing,
	// and local tables may include entries from the entire range of the tables.
	if sec.TablesVersion == 255 && sec.LocalTablesVersion == 0 {
		return sec, ErrInvalidTablesVersion
	}
	// (3) If Octet 11 is zero, Octet 10 must contain a valid master tables
	// version number and only those parts of the tables not reserved for local
	// use may be used.
	if sec.LocalTablesVersion == 0 && sec.TablesVersion == 0 {
		return sec, ErrInvalidTablesVersion
	}

	return sec, nil
}

package oolong

// https://codes.ecmwf.int/grib/format/grib2/sections/7/
type Section7Data struct {
	Section7Length  uint32 `byteRange:"1-4"`  // Length of section in octets (nn)
	NumberOfSection uint8  `byteRange:"5"`    // Number of section (7)
	Data            []byte `byteRange:"6-nn"` // Data in a format described by Data Template 7.x, where x is the Data Representation Template number given in octets 10-11 of Section 5.
}

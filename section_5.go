package oolong

// https://codes.ecmwf.int/grib/format/grib2/sections/5/
type Section5DataRepresentation struct {
	Section5Length                   uint32 `byteRange:"1-4"`   // Length of section in octets (nn)
	NumberOfSection                  uint8  `byteRange:"5"`     // Number of section (5)
	NumberOfValues                   uint32 `byteRange:"6-9"`   // Number of data points where one or more values are specified in Section 7 when a bit map is present, total number of data points when a bit map is absent.
	DataRepresentationTemplateNumber uint16 `byteRange:"10-11"` // Data Representation Template Number (see Code Table 5.0)
	DataRepresentationTemplate       []byte `byteRange:"12-nn"` // Data Representation Template (see Template 5.x, where x is the Data Representation Template Number given in octets 10-11)
}

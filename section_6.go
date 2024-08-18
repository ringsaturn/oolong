package oolong

// Octets	Key	Type	Content
// 1-4	section6Length	unsigned	Length of section in octets (nn)
// 5	numberOfSection	unsigned	Number of section (6)
// 6	bitMapIndicator	codetable	Bit-map indicator (see Code Table 6.0 and Note 1)
// 7-nn			Bit-map - Contiguous bits with a bit to data point correspondence, ordered as defined in Section 3. A bit set equal to 1 implies the presence of a data value at the corresponding data point, whereas a value of 0 implies the absence of such a value.

// https://codes.ecmwf.int/grib/format/grib2/sections/6/
type Section6BitMap struct {
	Section6Length  uint32 `byteRange:"1-4"`  // Length of section in octets (nn)
	NumberOfSection uint8  `byteRange:"5"`    // Number of section (6)
	BitMapIndicator uint8  `byteRange:"6"`    // Bit-map indicator (see Code Table 6.0)
	BitMap          []byte `byteRange:"7-nn"` // Bit-map - Contiguous bits with a bit to data point correspondence, ordered as defined in Section 3. A bit set equal to 1 implies the presence of a data value at the corresponding data point, whereas a value of 0 implies the absence of such a value.
}

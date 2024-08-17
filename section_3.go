package oolong

type Section3GridDefinition struct {
	Section3Length                   uint32 `byteRange:"1-4"`     // Length of section in octets (nn)
	NumberOfSection                  uint8  `byteRange:"5"`       // Number of section (3)
	SourceOfGridDefinition           uint8  `byteRange:"6"`       // Source of grid definition (see Code Table 3.0 and Note 1)
	NumberOfDataPoints               uint32 `byteRange:"7-10"`    // Number of data points
	NumberOfOctectsForNumberOfPoints uint8  `byteRange:"11"`      // Number of octets for optional list of numbers (see Note 2)
	InterpretationOfNumberOfPoints   uint8  `byteRange:"12"`      // Interpretation of list of numbers (see Code Table 3.11)
	GridDefinitionTemplateNumber     uint16 `byteRange:"13-14"`   // Grid Definition Template Number (= N) (see Code Table 3.1)
	GridDefinitionTemplate           []byte `byteRange:"15-xx"`   // Grid Definition Template (see Template 3.N, where N is the Grid Definition Template Number given in octets 13-14)
	OptionalList                     []byte `byteRange:"xx+1-nn"` // Optional list of numbers defining number of points (see Notes 2, 3 and 4)
}

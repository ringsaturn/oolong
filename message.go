package oolong

type Message struct {
	Indicator          Section0Indicator
	Identification     Section1Identification
	LocalUse           Section2LocalUse
	GridDefinition     Section3GridDefinition
	ProductDefinition  Section4ProductDefinition
	DataRepresentation Section5DataRepresentation
	BitMap             Section6BitMap
	Data               Section7Data
	End                Section8End
}

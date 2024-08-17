package oolong

type Message struct {
	Indicator      Section0Indicator
	Identification Section1Identification
	LocalUse       Section2LocalUse
}

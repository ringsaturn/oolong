package oolong

// https://codes.ecmwf.int/grib/format/grib2/sections/8/
type Section8End struct {
	END7777 string `byteRange:"1-4"` // "7777" (coded according to the International Alphabet No. 5.)
}

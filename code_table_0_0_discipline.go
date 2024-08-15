package oolong

// Discipline of processed data in the GRIB message, number of GRIB Master table
//
// - NOAA: https://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_table4-1.shtml
// - ECMWF: https://codes.ecmwf.int/grib/format/grib2/ctables/0/0/
type Discipline uint8

const (
	DisciplineMeteorological         Discipline = 0
	DisciplineHydrological           Discipline = 1
	DisciplineLandSurface            Discipline = 2
	DisciplineSatelliteRemoteSensing Discipline = 3 // formerly "Space Products"
	DisciplineSpaceWeather           Discipline = 4
	DisciplineOceanographic          Discipline = 10
	DisciplineHealthAndSocioeconomic Discipline = 20
	DisciplineMissing                Discipline = 255
)

func (d Discipline) String() string {
	switch d {
	case DisciplineMeteorological:
		return "Meteorological"
	case DisciplineHydrological:
		return "Hydrological"
	case DisciplineLandSurface:
		return "Land Surface"
	case DisciplineSatelliteRemoteSensing:
		return "Satellite Remote Sensing"
	case DisciplineSpaceWeather:
		return "Space Weather"
	case DisciplineOceanographic:
		return "Oceanographic"
	case DisciplineHealthAndSocioeconomic:
		return "Health and Socioeconomic"
	case DisciplineMissing:
		return "Missing"
	default:
		return "Unknown"
	}
}

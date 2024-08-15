package oolong

// Production status of processed data in this GRIB message
//
// - NOAA https://www.nco.ncep.noaa.gov/pmb/docs/grib2/grib2_doc/grib2_table1-3.shtml
//   - Original "Revised": "07/12/2024"
//   - Last Updated: 2024-08-15
//
// - ECMWF https://codes.ecmwf.int/grib/format/grib2/ctables/1/3/
type ProductionStatus uint8

const (
	ProductionStatusOperational            ProductionStatus = 0
	ProductionStatusOperationalTest        ProductionStatus = 1
	ProductionStatusResearch               ProductionStatus = 2
	ProductionStatusReAnalysis             ProductionStatus = 3
	ProductionStatusTIGGE                  ProductionStatus = 4
	ProductionStatusTIGGE_LAM              ProductionStatus = 5
	ProductionStatusS2SOperational         ProductionStatus = 6
	ProductionStatusS2STest                ProductionStatus = 7
	ProductionStatusUERRA                  ProductionStatus = 8
	ProductionStatusUERRATest              ProductionStatus = 9
	ProductionStatusCopernicusRegional     ProductionStatus = 10
	ProductionStatusCopernicusRegionalTest ProductionStatus = 11
	ProductionStatusDestinationEarth       ProductionStatus = 12
	ProductionStatusDestinationEarthTest   ProductionStatus = 13
	ProductionStatusMissing                ProductionStatus = 255
)

func (p ProductionStatus) String() string {
	switch p {
	case ProductionStatusOperational:
		return "Operational Products"
	case ProductionStatusOperationalTest:
		return "Operational Test Products"
	case ProductionStatusResearch:
		return "Research Products"
	case ProductionStatusReAnalysis:
		return "Re-Analysis Products"
	case ProductionStatusTIGGE:
		return "THORPEX Interactive Grand Global Ensemble (TIGGE)"
	case ProductionStatusTIGGE_LAM:
		return "THORPEX Interactive Grand Global Ensemble (TIGGE) test"
	case ProductionStatusS2SOperational:
		return "S2S Operational Products"
	case ProductionStatusS2STest:
		return "S2S Test Products"
	case ProductionStatusUERRA:
		return "Uncertainties in ensembles of regional reanalysis project (UERRA)"
	case ProductionStatusUERRATest:
		return "Uncertainties in ensembles of regional reanalysis project (UERRA) Test"
	case ProductionStatusCopernicusRegional:
		return "Copernicus Regional Reanalysis"
	case ProductionStatusCopernicusRegionalTest:
		return "Copernicus Regional Reanalysis Test"
	case ProductionStatusDestinationEarth:
		return "Destination Earth"
	case ProductionStatusDestinationEarthTest:
		return "Destination Earth test"
	case ProductionStatusMissing:
		return "Missing"
	}

	if p >= 14 && p <= 191 {
		return "Reserved"
	}
	if p >= 192 && p <= 254 {
		return "Reserved for Local Use"
	}

	return "Unknown"
}

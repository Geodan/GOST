package entities

import (
	"errors"
)

// EncodingType holds the information on a EncodingType
type EncodingType struct {
	Code  int
	Value string
}

// List of supported EncodingTypes, do not change!!
var (
	EncodingUnknown         = EncodingType{0, "unknown"}
	EncodingGeoJSON         = EncodingType{1, "application/vnd.geo+json"}
	EncodingPDF             = EncodingType{2, "application/pdf"}
	EncodingSensorML        = EncodingType{3, "http://www.opengis.net/doc/IS/SensorML/2.0"}
	EncodingTextHTML        = EncodingType{4, "text/html"}
	EncodingLocationType    = EncodingType{5, "http://example.org/location_types#GeoJSON"}
	EncodingTypeDescription = EncodingType{6, "http://schema.org/description"}
)

// EncodingValues is a list of names mapped to their EncodingValue
var EncodingValues = []EncodingType{
	EncodingUnknown,
	EncodingGeoJSON,
	EncodingPDF,
	EncodingSensorML,
	EncodingTextHTML,
	EncodingLocationType,
	EncodingTypeDescription,
}

//GetSupportedEncodings returns a list of supported encodings
func GetSupportedEncodings() string {
	var supportedEncodings string
	for _, k := range EncodingValues {
		supportedEncodings += k.Value + ", "
	}
	return supportedEncodings
}

// CreateEncodingType returns the int representation for a given encoding, returns an error when encoding is not supported
func CreateEncodingType(encoding string) (EncodingType, error) {
	for _, k := range EncodingValues {
		if k.Value == encoding {
			return k, nil
		}
	}
	supportedEncodings := GetSupportedEncodings()
	return EncodingUnknown, errors.New("Encoding not supported. Supported encodings:" + supportedEncodings)
}

package vin

import (
	"errors"
	"regexp"
	"strings"
)

type InvalidVINError error
type RegionNotFoundError error

var regionMap = map[string]string{
	"1": "United States",
	"2": "Canada",
	"3": "Mexico",
	"4": "United States",
	"5": "United States",
	"6": "Australia",
	"7": "New Zealand",
	"8": "Argentina",
	"9": "Brazil",
	"A": "South Africa",
	"B": "Angola",
	"C": "Benin",
	"D": "Egypt",
	"E": "Ethiopia",
	"F": "Kenya",
	"G": "Ivory Coast",
	"H": "Ghana",
	"J": "Japan",
	"K": "South Korea",
	"L": "China (Mainland)",
	"M": "India",
	"N": "Turkey",
	"P": "Philippines",
	"R": "Taiwan",
	"S": "United Kingdom",
	"T": "Switzerland",
	"U": "Germany",
	"V": "France",
	"W": "Germany",
	"X": "Russia",
	"Y": "Belgium",
	"Z": "Italy",
}

type VIN struct {
	value  string
	region string
}

func FromString(vin string) (*VIN, error) {
	return parse(vin)
}

func parse(vin string) (*VIN, error) {
	if !isValidVIN(vin) {
		return nil, InvalidVINError(errors.New("invalid length"))
	} else {
		regionFromVIN, err := getRegionFromVIN(vin)
		if err != nil {
			return nil, err
		}
		return &VIN{
			value:  vin,
			region: regionFromVIN,
		}, nil
	}
}

func isValidVIN(vin string) bool {
	if len(vin) != 17 {
		return false
	}

	vinPattern := regexp.MustCompile(`^[A-HJ-NPR-Z0-9]{17}$`)

	if !vinPattern.MatchString(vin) {
		return false
	}

	return true
}

func getRegionFromVIN(vin string) (string, error) {
	firstChar := strings.ToUpper(string(vin[0]))

	region, found := regionMap[firstChar]
	if !found {
		return "", RegionNotFoundError(errors.New("region not found"))
	}

	return region, nil
}

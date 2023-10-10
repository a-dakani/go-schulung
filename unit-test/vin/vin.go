package vin

import "errors"

type ErrInvalidLength error

type VIN struct {
	value  string
	region string
}

func FromString(vin string) (*VIN, error) {
	return parse(vin)
}

func parse(vin string) (*VIN, error) {
	if len(vin) != 17 {
		return nil, ErrInvalidLength(errors.New("invalid length"))
	} else {

		return &VIN{
			value:  vin,
			region: vin[0:3],
		}, nil
	}
}

package autoRepository

import "github.com/a-dakani/go-schulung/http-server-gin/ginserver"

type AutoRepository struct {
	autos []ginserver.Auto
}

func (ar *AutoRepository) AddAuto(auto ginserver.Auto) error {
	ar.autos = append(ar.autos, auto)
	return nil
}

func (ar *AutoRepository) GetAllAutos() ([]ginserver.Auto, error) {
	return ar.autos, nil
}

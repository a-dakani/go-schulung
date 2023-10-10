package autoRepository

import "github.com/a-dakani/go-schulung/http-server-gin/v2/ginserver"

type AutoRepository struct {
	Autos []ginserver.Auto
}

func (ar *AutoRepository) AddAuto(auto ginserver.Auto) error {
	ar.Autos = append(ar.Autos, auto)
	return nil
}

func (ar *AutoRepository) GetAllAutos() ([]ginserver.Auto, error) {
	return ar.Autos, nil
}

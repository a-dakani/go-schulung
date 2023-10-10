package ginserver

import "context"

type (
	Kennenzeichen  string
	AutoRepository interface {
		AddAuto(ctx context.Context, auto Auto) error
		GetAllAutos(ctx context.Context) ([]Auto, error)
	}

	Audi struct {
		MotorGestartet     bool `json:"_"`
		GeschwindigkeitKmH int  `json:"_"`
		Kennenzeichen      `json:"kennzeichen"`
		*Getriebe          `json:"getriebe,omitempty"`
	}

	Auto interface {
		MotorStarten() error
		Beschleunigen()
		Bremsen()
	}

	Getriebe struct {
		EingangsDrehzahl int
		EingelegterGang  int
	}
)

func (g *Getriebe) initialisieren() error {
	return nil
}

func (a Audi) MotorStarten() error {
	a.MotorGestartet = true
	return nil
}
func (a Audi) Beschleunigen() {
	a.GeschwindigkeitKmH += 10
}
func (a Audi) Bremsen() {
	a.GeschwindigkeitKmH -= 10
}

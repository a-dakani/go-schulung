package ginserver

type (
	Kennzeichen    string
	AutoRepository interface {
		AddAuto(auto Auto) error
		GetAllAutos() ([]Auto, error)
	}

	Audi struct {
		MotorGestartet     bool `json:"_"`
		GeschwindigkeitKmH int  `json:"_"`
		Kennzeichen        `json:"kennzeichen"`
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

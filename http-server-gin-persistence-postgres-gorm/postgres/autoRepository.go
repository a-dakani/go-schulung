package postgres

import (
	"context"
	"github.com/a-dakani/go-schulung/http-server-gin-persistence-postgres-gorm/ginserver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AutoRepository struct {
	db *gorm.DB
}

type Auto struct {
	gorm.Model
	Kennzeichen     string
	Geschwindigkeit int
	MotorGestartet  bool
	GetriebeID      int
	Getriebe        Getriebe `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Getriebe struct {
	gorm.Model
	EingelegterGang  int
	EingangsDrehzahl int
}

func NewAutoRepository(ctx context.Context) (*AutoRepository, error) {
	dsn := "host=localhost user=user password=pass dbname=db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)

	err = db.AutoMigrate(&Auto{}, &Getriebe{})
	if err != nil {
		return nil, err
	}
	return &AutoRepository{db: db}, nil
}

func (ar *AutoRepository) AddAuto(ctx context.Context, auto ginserver.Auto) error {
	internalAuto := mapAuto(auto)
	return ar.db.WithContext(ctx).Create(&internalAuto).Error
}

func (ar *AutoRepository) GetAllAutos(ctx context.Context) ([]ginserver.Auto, error) {
	var autos []Auto
	err := ar.db.WithContext(ctx).Preload("Getriebe").Find(&autos).Error
	if err != nil {
		return nil, err
	}
	return mapAutoList(autos), nil
}

func mapAutoList(autos []Auto) []ginserver.Auto {
	var internalAutos []ginserver.Auto
	for _, auto := range autos {
		audi := ginserver.Audi{
			Kennzeichen:        ginserver.Kennzeichen(auto.Kennzeichen),
			GeschwindigkeitKmH: auto.Geschwindigkeit,
			MotorGestartet:     auto.MotorGestartet,
			Getriebe: &ginserver.Getriebe{
				EingangsDrehzahl: auto.Getriebe.EingangsDrehzahl,
				EingelegterGang:  auto.Getriebe.EingelegterGang,
			},
		}
		internalAutos = append(internalAutos, audi)
	}
	return internalAutos
}

func mapAuto(a ginserver.Auto) Auto {

	internalAuto := Auto{}
	if audi, ok := a.(ginserver.Audi); ok {
		internalAuto = Auto{
			Kennzeichen:     string(audi.Kennzeichen),
			MotorGestartet:  audi.MotorGestartet,
			Geschwindigkeit: audi.GeschwindigkeitKmH,
			Getriebe: Getriebe{
				EingangsDrehzahl: audi.Getriebe.EingangsDrehzahl,
				EingelegterGang:  audi.Getriebe.EingelegterGang,
			},
		}
	}
	return internalAuto
}

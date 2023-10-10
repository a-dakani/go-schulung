package autoRepository

import (
	"context"
	"fmt"
	Logging "github.com/a-dakani/go-schulung/contextAndLogging"
	"github.com/a-dakani/go-schulung/contextAndLogging/ginserver"
	"time"
)

type AutoRepository struct {
	Autos []ginserver.Auto
}

func (ar *AutoRepository) AddAuto(ctx context.Context, auto ginserver.Auto) error {
	ar.Autos = append(ar.Autos, auto)
	return nil
}

func (ar *AutoRepository) GetAllAutos(ctx context.Context) ([]ginserver.Auto, error) {
	logger := Logging.Logger(ctx)
	logger.Info("GetAllAutos abgerufen")

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("weiter")
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return ar.Autos, nil
}

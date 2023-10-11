package ginserver

import "golang.org/x/net/context"

type AutoService struct {
	notifier AutoNotifier
	repo     AutoRepository
}

func NewAutoService(notifier AutoNotifier, repo AutoRepository) *AutoService {

	return &AutoService{notifier: notifier, repo: repo}
}

func (as *AutoService) AddAuto(ctx context.Context, auto Auto) error {
	err := as.repo.AddAuto(ctx, auto)
	if err != nil {
		return err
	}
	return as.notifier.NewAutoCreated(ctx, auto)
}

func (as *AutoService) GetAllAutos(ctx context.Context) ([]Auto, error) {
	return as.repo.GetAllAutos(ctx)
}

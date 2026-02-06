package service

import (
	"context"
	"github.com/arthurrso/statectl/internal/controller"
	"github.com/arthurrso/statectl/internal/domain"
)

type InstanceService struct {
	repo       domain.InstanceRepository
	controller *controller.InstanceController
}

func NewInstanceService(
	repo domain.InstanceRepository,
	controller *controller.InstanceController,
) *InstanceService {
	return &InstanceService{
		repo:       repo,
		controller: controller,
	}
}

// CREATE
func (s *InstanceService) Create(ctx context.Context, name string, spec domain.InstanceSpec) error {
	inst := &domain.Instance{
		Name: name,
		Spec: domain.InstanceSpec{
			DesiredState: domain.DesiredRunning,
		},
	}

	if err := s.repo.Create(ctx, inst); err != nil {
		return err
	}

	return s.controller.Reconcile(ctx, name)
}

// DELETE
func (s *InstanceService) Delete(ctx context.Context, name string) error {
	inst, err := s.repo.Get(ctx, name)

	if err != nil {
		return err
	}

	inst.Spec.DesiredState = domain.DesiredDeleted

	if err := s.repo.UpdateSpec(ctx, inst); err != nil {
		return err
	}

	return s.controller.Reconcile(ctx, name)
}

// LIST
func (s *InstanceService) List(ctx context.Context) ([]*domain.Instance, error) {
	return s.repo.List(ctx)
}

// GET
func (s *InstanceService) Get(ctx context.Context, name string) (*domain.Instance, error) {
	return s.repo.Get(ctx, name)
}

// UPDATESPEC
func (s *InstanceService) UpdateSpec(ctx context.Context, name string, spec domain.InstanceSpec) error {
	inst, err := s.repo.Get(ctx, name)

	if err != nil {
		return err
	}

	inst.Spec = spec

	if err := s.repo.UpdateSpec(ctx, inst); err != nil {
		return err
	}

	return s.controller.Reconcile(ctx, name)
}

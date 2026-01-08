package domain

import "context"

type InstanceRepository interface {
	Get(ctx context.Context, name string) (*Instance, error)
	List(ctx context.Context) ([]*Instance, error)

	Create(ctx context.Context, instance *Instance) error
	UpdateSpec(ctx context.Context, instance *Instance) error
	UpdateStatus(ctx context.Context, instance *Instance) error
	Delete(ctx context.Context, name string) error
}

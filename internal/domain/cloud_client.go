package domain

import "context"

type CloudInstanceClient interface {
	Get(ctx context.Context, name string) (*CloudInstance, error)
	Create(ctx context.Context, spec InstanceSpec) (providerID string, err error)
	Start(ctx context.Context, providerID string) error
	Stop(ctx context.Context, providerID string) error
	Delete(ctx context.Context, providerID string) error
}

type CloudInstance struct {
	ProviderID string
	State      InstanceState
}

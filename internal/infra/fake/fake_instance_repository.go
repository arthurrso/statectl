package fake

import (
	"context"
	"errors"
	"sync"

	"gdch-cli/internal/domain"
)

type FakeInstanceRepository struct {
	mu        sync.RWMutex
	instances map[string]*domain.Instance
}

func NewFakeInstanceRepository() *FakeInstanceRepository {
	return &FakeInstanceRepository{
		instances: make(map[string]*domain.Instance),
	}
}

func (r *FakeInstanceRepository) Get(ctx context.Context, name string) (*domain.Instance, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	inst, ok := r.instances[name]
	if !ok {
		return nil, errors.New("instance not found")
	}

	return cloneInstance(inst), nil
}

func (r *FakeInstanceRepository) Create(ctx context.Context, inst *domain.Instance) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.instances[inst.Name]; exists {
		return errors.New("instance already exists")
	}

	r.instances[inst.Name] = cloneInstance(inst)
	return nil
}

func (r *FakeInstanceRepository) UpdateSpec(ctx context.Context, inst *domain.Instance) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, ok := r.instances[inst.Name]
	if !ok {
		return errors.New("instance not found")
	}

	existing.Spec = inst.Spec
	return nil
}

func (r *FakeInstanceRepository) UpdateStatus(ctx context.Context, inst *domain.Instance) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, ok := r.instances[inst.Name]
	if !ok {
		return errors.New("instance not found")
	}

	existing.Status = inst.Status
	return nil
}

func (r *FakeInstanceRepository) Delete(ctx context.Context, name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.instances, name)
	return nil
}

func cloneInstance(in *domain.Instance) *domain.Instance {
	if in == nil {
		return nil
	}

	out := *in
	return &out
}
func (r *FakeInstanceRepository) List(ctx context.Context) ([]*domain.Instance, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*domain.Instance, 0, len(r.instances))

	for _, inst := range r.instances {
		result = append(result, cloneInstance(inst))
	}

	return result, nil
}

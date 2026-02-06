package fake

import (
	"context"
	"errors"
	"fmt"
	"github.com/arthurrso/statectl/internal/domain"
	"sync"
)

type FakeCloudInstanceClient struct {
	mu sync.Mutex

	instances map[string]*domain.CloudInstance
	counter   int

	FailCreate bool
	FailStart  bool
	FailStop   bool
	FailDelete bool
}

func NewFakeCloudInstanceClient() *FakeCloudInstanceClient {
	return &FakeCloudInstanceClient{
		instances: make(map[string]*domain.CloudInstance),
	}
}

func (c *FakeCloudInstanceClient) Create(ctx context.Context, spec domain.InstanceSpec) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.FailCreate {
		return "", errors.New("cloud provider failed to create instance")
	}

	c.counter++
	id := fmt.Sprintf("inst-%d", c.counter)

	c.instances[id] = &domain.CloudInstance{
		ProviderID: id,
		State:      domain.StateProvisioning,
	}

	return id, nil
}

func (c *FakeCloudInstanceClient) Get(ctx context.Context, providerID string) (*domain.CloudInstance, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	inst, ok := c.instances[providerID]
	if !ok {
		return nil, errors.New("cloud instance not found")
	}

	copy := *inst
	return &copy, nil
}

func (c *FakeCloudInstanceClient) Start(ctx context.Context, providerID string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.FailStart {
		return errors.New("cloud provider failed to start instance")
	}

	inst, ok := c.instances[providerID]
	if !ok {
		return errors.New("cloud instance not found")
	}

	inst.State = domain.StateRunning
	return nil
}

func (c *FakeCloudInstanceClient) Stop(ctx context.Context, providerID string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.FailStop {
		return errors.New("cloud provider failed to stop instance")
	}

	inst, ok := c.instances[providerID]
	if !ok {
		return errors.New("cloud instance not found")
	}

	inst.State = domain.StateStopped
	return nil
}

func (c *FakeCloudInstanceClient) Delete(ctx context.Context, providerID string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.FailDelete {
		return errors.New("cloud provider failed to delete instance")
	}

	delete(c.instances, providerID)
	return nil
}

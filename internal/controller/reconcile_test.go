package controller

import (
	"context"
	"github.com/arthurrso/statectl/internal/domain"
	"github.com/arthurrso/statectl/internal/infra/fake"
	"testing"
)

func TestReconcile_CreateInstance(t *testing.T) {
	ctx := context.Background()

	repo := fake.NewFakeInstanceRepository()
	cloud := fake.NewFakeCloudInstanceClient()

	controller := &InstanceController{
		repo:   repo,
		client: cloud,
	}

	inst := &domain.Instance{
		Name: "vm-1",
		Spec: domain.InstanceSpec{
			DesiredState: domain.DesiredRunning,
			MachineType:  "e2-medium",
			Region:       "us-central1",
			DiskSizeGB:   10,
		},
	}

	if err := repo.Create(ctx, inst); err != nil {
		t.Fatalf("failed to create instance in repo: %v", err)
	}

	if err := controller.Reconcile(ctx, "vm-1"); err != nil {
		t.Fatalf("reconcile returned error: %v", err)
	}

	updated, err := repo.Get(ctx, "vm-1")
	if err != nil {
		t.Fatalf("failed to get instance: %v", err)
	}

	if updated.Status.ProviderID == "" {
		t.Fatalf("expected ProviderID to be set")
	}

	if updated.Status.State != domain.StateProvisioning && updated.Status.State != domain.StateRunning {
		t.Fatalf("unexpected state: %s", updated.Status.State)
	}
}

package controller

import (
	"context"
	"statectl/internal/domain"
)

func (c *InstanceController) Reconcile(ctx context.Context, name string) error {
	inst, err := c.repo.Get(ctx, name)
	if err != nil {
		// não existe mais no repositório → nada a fazer
		return nil
	}

	switch inst.Spec.DesiredState {

	case domain.DesiredRunning:
		return c.reconcileRunning(ctx, inst)

	case domain.DesiredStopped:
		return c.reconcileStopped(ctx, inst)

	case domain.DesiredDeleted:
		return c.reconcileDeleted(ctx, inst)

	default:
		return nil
	}
}

func (c *InstanceController) reconcileRunning(ctx context.Context, inst *domain.Instance) error {

	// Caso 1: ainda não existe no cloud → criar
	if inst.Status.ProviderID == "" {
		id, err := c.client.Create(ctx, inst.Spec)
		if err != nil {
			inst.Status.State = domain.StateError
			inst.Status.LastError = err.Error()
			_ = c.repo.UpdateStatus(ctx, inst)
			return err
		}

		inst.Status.ProviderID = id
		inst.Status.State = domain.StateProvisioning
		return c.repo.UpdateStatus(ctx, inst)
	}

	// Caso 2: já existe → olhar estado real
	cloudInst, err := c.client.Get(ctx, inst.Status.ProviderID)
	if err != nil {
		inst.Status.State = domain.StateError
		inst.Status.LastError = err.Error()
		_ = c.repo.UpdateStatus(ctx, inst)
		return err
	}

	switch cloudInst.State {

	case domain.StateProvisioning:
		// ainda criando → só refletir
		inst.Status.State = domain.StateProvisioning
		return c.repo.UpdateStatus(ctx, inst)

	case domain.StateStopped:
		// existe, mas parado → start
		if err := c.client.Start(ctx, inst.Status.ProviderID); err != nil {
			inst.Status.State = domain.StateError
			inst.Status.LastError = err.Error()
			_ = c.repo.UpdateStatus(ctx, inst)
			return err
		}

		inst.Status.State = domain.StateProvisioning
		return c.repo.UpdateStatus(ctx, inst)

	case domain.StateRunning:
		// estado desejado atingido
		inst.Status.State = domain.StateRunning
		inst.Status.LastError = ""
		return c.repo.UpdateStatus(ctx, inst)

	default:
		return nil
	}
}
func (c *InstanceController) reconcileStopped(ctx context.Context, inst *domain.Instance) error {

	if inst.Status.ProviderID == "" {
		// nunca foi criada → já está parada
		inst.Status.State = domain.StateStopped
		return c.repo.UpdateStatus(ctx, inst)
	}

	cloudInst, err := c.client.Get(ctx, inst.Status.ProviderID)
	if err != nil {
		inst.Status.State = domain.StateError
		inst.Status.LastError = err.Error()
		_ = c.repo.UpdateStatus(ctx, inst)
		return err
	}

	if cloudInst.State == domain.StateRunning {
		if err := c.client.Stop(ctx, inst.Status.ProviderID); err != nil {
			inst.Status.State = domain.StateError
			inst.Status.LastError = err.Error()
			_ = c.repo.UpdateStatus(ctx, inst)
			return err
		}

		inst.Status.State = domain.StateStopping
		return c.repo.UpdateStatus(ctx, inst)
	}

	inst.Status.State = domain.StateStopped
	return c.repo.UpdateStatus(ctx, inst)
}

func (c *InstanceController) reconcileDeleted(ctx context.Context, inst *domain.Instance) error {

	if inst.Status.ProviderID == "" {
		_ = c.repo.Delete(ctx, inst.Name)
		return nil
	}

	if err := c.client.Delete(ctx, inst.Status.ProviderID); err != nil {
		inst.Status.State = domain.StateError
		inst.Status.LastError = err.Error()
		_ = c.repo.UpdateStatus(ctx, inst)
		return err
	}

	_ = c.repo.Delete(ctx, inst.Name)
	return nil
}

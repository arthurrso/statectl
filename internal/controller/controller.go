package controller

import (
	"gdch-cli/internal/domain"
)

type InstanceController struct {
	repo   domain.InstanceRepository
	client domain.CloudInstanceClient
}

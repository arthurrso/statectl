package controller

import (
	"statectl/internal/domain"
)

type InstanceController struct {
	repo   domain.InstanceRepository
	client domain.CloudInstanceClient
}

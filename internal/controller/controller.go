package controller

import (
	"github.com/arthurrso/statectl/internal/domain"
)

type InstanceController struct {
	repo   domain.InstanceRepository
	client domain.CloudInstanceClient
}
